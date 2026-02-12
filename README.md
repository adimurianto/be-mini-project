# BE Mini Project

## Setup Project

- Clone project

```bash
  git clone https://github.com/adimurianto/be-mini-project.git
  cd be-mini-project
```


- Get all dependency

```bash
  go get .
```


- Create file **.env** base on file **.env.example** and adjust the contents of the variable
  

-  Init for generate Swagger API documentation

```bash
  swag init
```


-  Init for live reload

```bash
  air init
```


- Running project

Choose one
  
```bash
  // if using live reload
  air

  // test locally on your system
  go run main.go

  // build locally on your system then make sure run /main.exe or /main
  go build main.go
```


- Access url http://127.0.0.1:5000/docs/index.html


  <img width="auto" alt="Screenshot 2025-02-04 at 14 18 01" src="https://github.com/user-attachments/assets/a071689b-f935-4d13-8042-0f2a386cd288" />



## Folder Structure
Here are the folders to consider when adding a new endpoint

```
├── controllers/
│   ├── user_controller.go
│   └── ...
├── models/
│   ├── user_model.go
│   └── ...
├── routers/
|   ├── groups/
│       ├── user.go
|       └── ...
|   └── ...
```
