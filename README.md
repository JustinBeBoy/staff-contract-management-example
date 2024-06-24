
# Staff and Contract Management System

This project is a staff and contract management system with APIs only. It provides functionalities for two user roles: Admin and Staff. Admin can manage all staff and their contracts, while staff can only manage their own contracts. All APIs require JWT token authentication.

## Setup

1. **Install Dependencies**
   Make sure you have Go installed. Then, run:
   ```sh
   go mod tidy
   ```

2. **Run the Application**
   ```sh
   go run cmd/main.go
   ```

## APIs

### Authentication

- **Login**
  - **POST** `/auth/login`
  - Input: `{ "username": "string", "password": "string" }`
  - Output: `{ "token": "string" }`

### Contracts (Admin and Staff)

- **Create Contract**
  - **POST** `/contracts/`
  - Input: `{ "Title": "string", "Content": "string" }`
  - Output: `Contract Object`

- **Get Contracts**
  - **GET** `/contracts/`
  - Output: `List of Contract Objects`

- **Get Contract Detail**
  - **GET** `/contracts/:id`
  - Output: `Contract Object`

- **Update Contract**
  - **PUT** `/contracts/:id`
  - Input: `{ "Title": "string", "Content": "string" }`
  - Output: `Updated Contract Object`

- **Delete Contract**
  - **DELETE** `/contracts/:id`

### Staff (Admin Only)

- **Create Staff**
  - **POST** `/staffs/`
  - Input: `{ "Username": "string", "Password": "string", "Role": "string" }`
  - Output: `Staff Object`

- **Get Staffs**
  - **GET** `/staffs/`
  - Output: `List of Staff Objects`

- **Get Staff Detail**
  - **GET** `/staffs/:id`
  - Output: `Staff Object`

- **Update Staff**
  - **PUT** `/staffs/:id`
  - Input: `{ "Username": "string", "Password": "string", "Role": "string" }`
  - Output: `Updated Staff Object`

- **Delete Staff**
  - **DELETE** `/staffs/:id`

## Running Tests

To run the tests, navigate to the `tests/scripts` directory and run the test script:

```sh
cd tests/scripts
go run run_tests.go
```

This will run all the unit tests and output the results.

## Background Tasks

Background tasks are handled using NATS. The `delete.contract` and `delete.staff` subjects are used to process delete tasks in the background.

1. **Initialize NATS**
   ```go
   nc := utils.InitNATS()
   ```

2. **Subscribe to Delete Tasks**
   ```go
   go tasks.SubscribeToDeleteTasks(nc, db)
   ```

3. **Publish to NATS for Background Task**
   ```go
   nc.Publish("delete.contract", []byte(contractID))
   nc.Publish("delete.staff", []byte(staffID))
   ```

## Utils

### Database Initialization
**utils/db.go**
```go
func InitDB() *gorm.DB {
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect database")
    }
    db.AutoMigrate(&models.Staff{}, &models.Contract{})
    return db
}
```

### NATS Initialization
**utils/nats.go**
```go
func InitNATS() *nats.Conn {
    nc, err := nats.Connect(nats.DefaultURL)
    if err != nil {
        log.Fatal(err)
    }
    return nc
}
```

### Response Helpers
**utils/respond.go**
```go
func RespondError(w http.ResponseWriter, code int, message string) {
    RespondJSON(w, code, map[string]string{"error": message})
}

func RespondJSON(w http.ResponseWriter, code int, payload interface{}) {
    response, _ := json.Marshal(payload)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(response)
}
```