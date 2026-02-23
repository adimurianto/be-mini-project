package repository

import (
	"be-mini-project/infra/database"
	"fmt"
	"math"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

func Get(model interface{}) error {
	database.DB = database.DB.Debug()
	err := database.DB.Find(model).Error
	if err != nil {
		return err
	}
	return nil
}

func Save(model interface{}) interface{} {
	err := database.DB.Debug().Create(model).Error
	if err != nil {
		return err
	}
	return nil
}

func GetWithFilter(model interface{}, sort string, filter string, page int, perPage int, preload ...string) (interface{}, int, int, error) {
	database.DB = database.DB.Debug()
	db := database.DB

	// Don't include is_deleted
	db = db.Where("status = ?", true)

	// Apply sorting
	db = applySortingNew(db, sort)

	// Apply filtering
	var err error
	db, err = applyFilteringNew(db, filter)

	if err != nil {
		return nil, 0, 0, err
	}

	// Get total count before applying pagination
	var totalCount int64
	if err := db.Model(model).Count(&totalCount).Error; err != nil {
		return nil, 0, 0, err
	}
	totalPages := int(math.Ceil(float64(totalCount) / float64(perPage)))

	// Apply pagination
	if page > 0 && perPage > 0 {
		offset := (page - 1) * perPage
		db = db.Offset(offset).Limit(perPage)
	}

	// Apply preloading
	if len(preload) > 0 {
		for i := 0; i < len(preload); i++ {
			db = db.Preload(preload[i])
		}
	} else {
		db = db.Preload(clause.Associations)
	}

	// Find records
	err = db.Find(model).Error
	if err != nil {
		return nil, 0, 0, err
	}

	database.DB = database.DB.Session(&gorm.Session{Logger: database.DB.Logger.LogMode(logger.Silent)})

	return model, int(totalCount), totalPages, nil
}

func applySortingNew(db *gorm.DB, sort string) *gorm.DB {
	if sort != "" {
		// Split the sort parameter into field and order
		sortParts := strings.Split(sort, ",")
		if len(sortParts) == 2 {
			field, order := sortParts[0], sortParts[1]
			// Check for valid order values (asc or desc)
			if order == "asc" || order == "desc" {
				// Apply sorting
				return db.Order(field + " " + order)
			}
		} else {
			return db.Order(sort + " asc")
		}
	}

	// Default sorting if no specific sort is provided
	return db.Order("COALESCE(id) desc")
}

func applyFilteringNew(db *gorm.DB, filter string) (*gorm.DB, error) {
	if filter == "" {
		return db, nil
	}

	filters, err := parseFilterString(filter)
	if err != nil {
		return nil, err
	}

	if len(filters) > 0 {
		for column, value := range filters {
			column = db.NamingStrategy.ColumnName("", column)
			operator := ""
			values := ""

			if strings.Contains(value, ",") {
				parts := strings.Split(value, ",")
				value = parts[0]
				operator = parts[len(parts)-1]
				if len(parts) > 1 {
					values = strings.Join(parts[:len(parts)-1], ",")
				}
			}

			switch operator {
			case ">":
				db = db.Where(column+" > ?", value)
			case "<":
				db = db.Where(column+" < ?", value)
			case ">=":
				db = db.Where(column+" >= ?", value)
			case "<=":
				db = db.Where(column+" <= ?", value)
			case "!=":
				db = db.Where(column+" != ?", value)
			case "in":
				inValues := strings.Split(values, ",")
				for i, v := range inValues {
					inValues[i] = strings.TrimSpace(v)
				}
				db = db.Where(column+" IN (?)", inValues)
			case "like":
				db = db.Where(column+" LIKE ?", "%"+value+"%")
			case "ilike":
				db = db.Where(column+" ILIKE ?", "%"+value+"%")
			default:
				db = db.Where(column+" = ?", value)
			}
		}
	}

	return db, nil
}

func parseFilterString(filterString string) (map[string]string, error) {
	if filterString == "" {
		return nil, nil
	}

	filterString = strings.TrimPrefix(filterString, "[")
	filterString = strings.TrimSuffix(filterString, "]")

	filters := make(map[string]string)
	for _, filterPart := range strings.Split(filterString, ";") {
		components := strings.Split(filterPart, ",")

		// Validate minimum components
		if len(components) < 3 {
			return nil, fmt.Errorf("invalid filter format: %s", filterPart)
		}

		column := components[0]
		operator := components[1]

		if operator == "in" {
			value := strings.Join(components[2:], ",")
			filters[column] = value + "," + operator
		} else {
			if len(components) != 3 {
				return nil, fmt.Errorf("invalid filter format: %s", filterPart)
			}
			value := components[2]
			filters[column] = value + "," + operator
		}
	}
	return filters, nil
}

func GetById(model interface{}, id any) error {
	if err := database.DB.Where("id = ?", id).First(model).Error; err != nil {
		return err
	}
	return nil
}

func Update(model interface{}) interface{} {
	err := database.DB.Save(model).Error
	return err
}
