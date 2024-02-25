package main

import (
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "github.com/google/uuid"
    "os"
    "bufio"
)

func customerFunction() {
    fmt.Println("Hello, patient!")
}

func adminFunction() {
    scanner := bufio.NewScanner(os.Stdin)
    var choice int
OuterLoop:
    for {
        fmt.Println("\nAdmin Menu:")
        fmt.Println("1. Rooms")
        fmt.Println("2. Doctor function")
        fmt.Println("3. Go back to Main Menu")
        fmt.Print("Enter your choice: ")
        fmt.Scanln(&choice)

        switch choice {
        case 1:
            for {
                fmt.Println("\nRooms Menu:")
                fmt.Println("1. Add room")
                fmt.Println("2. Edit room")
                fmt.Println("3. Delete room")
                fmt.Println("4. Go back to Admin Menu")
                fmt.Print("Enter your choice: ")
                fmt.Scanln(&choice)

                switch choice {
                case 1:
                    // To read the whole line, use standard input scanner
                    var roomType string
                    fmt.Print("Enter the room type: ")
                    scanner.Scan()
                    roomType = scanner.Text()                    

                    var roomNumber int
                    fmt.Print("Enter the room number: ")
                    fmt.Scanln(&roomNumber)

                    // fmt.Printf("You entered room type: %s and room number: %d\n", roomType, roomNumber)

                    err := addRoom(roomType, roomNumber)
                    if err != nil {
                        fmt.Println("Error creating user:", err)
                    } else {
                        fmt.Println("User created successfully")
                    }

                case 2:
                    fmt.Println("To be edited soon")
                case 3:
                    fmt.Println("To be deleted soon")
                case 4:
                    fmt.Println("Going back to Admin Menu...")
                    continue OuterLoop 
                default:
                    fmt.Println("Invalid choice. Please try again.")
                }
            }
        case 2:
            fmt.Println("Doctor function is not implemented yet.")
            // Implement Doctor function logic here
        case 3:
            fmt.Println("Going back to Main Menu...")
            return
        default:
            fmt.Println("Invalid choice. Please try again.")
        }
    }
}



func main() {
    var choice int
    var username, password string

    for {
        fmt.Println("Main Menu:")
        fmt.Println("1. Customer")
        fmt.Println("2. Admin")
        fmt.Println("3. Exit")
        fmt.Print("Enter your choice: ")
        fmt.Scanln(&choice)

        switch choice {
        case 1:
            customerFunction()
        case 2:
            fmt.Print("Enter username: ")
            fmt.Scanln(&username)
            fmt.Print("Enter password: ")
            fmt.Scanln(&password)
            if username == "admin" && password == "admin" {
                adminFunction()
            } else {
                fmt.Println("Incorrect username or password. Try again.")
            }
        case 3:
            fmt.Println("Exiting program...")
            return
        default:
            fmt.Println("Invalid choice. Please try again.")
        }
    }
}

func connectDB() (*sql.DB, error) {
    db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/mydb")
    if err != nil {
        return nil, err
    }
    return db, nil
}

func addRoom(roomType string, roomNumber int) error {
    db, err := connectDB()
    if err != nil {
        return err
    }
    defer db.Close()

    uuid := uuid.New().String()

    _, err = db.Exec("INSERT INTO tbl_rooms (room_id, room_type, room_number) VALUES (?, ?, ?)", uuid, roomType, roomNumber)
    if err != nil {
        return err
    }

    return nil
}
