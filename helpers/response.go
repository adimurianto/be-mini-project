package helpers

type Response struct {
	Code    int         `json:"code"`
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Info    interface{} `json:"info"`
	Data    interface{} `json:"data,omitempty"`
}

type Info struct {
	Page       int `json:"page,omitempty"`
	PerPage    int `json:"perPage,omitempty"`
	TotalPages int `json:"totalPages,omitempty"`
	TotalData  int `json:"totalData,omitempty"`
}

type AuthResponse struct {
	Code    int         `json:"code"`
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
