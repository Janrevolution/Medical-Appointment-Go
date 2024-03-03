package main

import (
	"bufio"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/MasterDimmy/go-cls"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

func secretary() {
	fmt.Println("Welcom Sec!")
}

func doctor() {
	fmt.Println("Welcome Doc!")
}

func adminFunction() {
	scanner := bufio.NewScanner(os.Stdin)
	var choice int
	var err error
OuterLoop:
	for {
		fmt.Println("\nAdmin Menu:")
		fmt.Println("1. Rooms")
		fmt.Println("2. Employee function")
		fmt.Println("3. Assign Doctor")
		fmt.Println("4. Create Account")
		fmt.Println("5. Go back to Main Menu")
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			for {
				err = printRooms()
				if err != nil {
					fmt.Println("Error reading room data:", err)
				}

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
						cls.CLS()
						fmt.Println("Error creating user:", err)
					} else {
						cls.CLS()
						fmt.Println("User created successfully")
					}

				case 2:
					fmt.Println("To be edited soon")
				case 3:
					var roomNumber string
					fmt.Print("Enter the room number to be deleted: ")
					fmt.Scanln(&roomNumber)

					err := deleteRecord(roomNumber, "room")
					if err != nil {
						cls.CLS()
						fmt.Println("Error deleting room:", err)
					} else {
						cls.CLS()
						fmt.Println("Room deleted successfully")
					}

				case 4:
					fmt.Println("Going back to Admin Menu...")
					continue OuterLoop
				default:
					fmt.Println("Invalid choice. Please try again.")
				}
			}
		case 2:
			for {
				err = printEmployees()
				if err != nil {
					fmt.Println("Error reading employee data:", err)
				}

				fmt.Println("\nEmployee Menu:")
				fmt.Println("1. Add Employee")
				fmt.Println("2. Edit Employee")
				fmt.Println("3. Delete Employee")
				fmt.Println("4. Go back to Admin Menu")
				fmt.Print("Enter your choice: ")
				fmt.Scanln(&choice)

				switch choice {
				case 1:

					// To read the whole line, use standard input scanner
					var lastName, firstName, middleName, profession, specialization string

					fmt.Print("Enter Last Name: ")
					scanner.Scan()
					lastName = scanner.Text()

					fmt.Print("Enter First Name: ")
					scanner.Scan()
					firstName = scanner.Text()

					fmt.Print("Enter Middle Name: ")
					scanner.Scan()
					middleName = scanner.Text()

					fmt.Print("Enter Profession: ")
					scanner.Scan()
					profession = scanner.Text()

					fmt.Print("Enter Specialization(N/A for non-doctors): ")
					scanner.Scan()
					specialization = scanner.Text()

					err := addEmployee(lastName, firstName, middleName, profession, specialization)
					if err != nil {
						cls.CLS()
						fmt.Println("Error creating user:", err)
					} else {
						cls.CLS()
						fmt.Println("User created successfully")
					}

				case 2:
					fmt.Println("To be edited soon")
				case 3:
					var hp_id string

					fmt.Print("Enter Employee ID to be deleted: ")
					scanner.Scan()
					hp_id = scanner.Text()

					err := deleteRecord(hp_id, "employee")
					if err != nil {
						cls.CLS()
						fmt.Println("Error removing Employee:", err)
					} else {
						cls.CLS()
						fmt.Println("Employee removed successfully")
					}

				case 4:
					fmt.Println("Going back to Admin Menu...")
					continue OuterLoop
				default:
					fmt.Println("Invalid choice. Please try again.")
				}
			}
		case 3:
			for {
				fmt.Println("\nAssign Menu:")
				fmt.Println("1. Assign Doctor")
				fmt.Println("2. Edit Doctor")
				fmt.Println("3. Remove Doctor assignment")
				fmt.Println("4. Go back to Admin Menu")
				fmt.Print("Enter your choice: ")
				fmt.Scanln(&choice)

				switch choice {
				case 1:
					fmt.Println("Room Data:")
					err = printRooms()
					if err != nil {
						fmt.Println("Error reading room data:", err)
					}

					fmt.Println("\nDoctor Data:")
					err = printDoctors()
					if err != nil {
						fmt.Println("Error reading doctor data:", err)
					}

					fmt.Println("\nAssigned Doctors: ")
					err = printAssignedDoctor()
					if err != nil {
						fmt.Println("Error deleting doctor & room data:", err)
					}

					// To read the whole line, use standard input scanner
					var roomNumber string
					fmt.Print("Enter room number: ")
					scanner.Scan()
					roomNumber = scanner.Text()

					var doctorId string
					fmt.Print("Enter doctor ID: ")
					scanner.Scan()
					doctorId = scanner.Text()

					// fmt.Printf("You entered room type: %s and room number: %d\n", roomType, roomNumber)

					err := assignDoctor(roomNumber, doctorId)
					if err != nil {
						cls.CLS()
						fmt.Println("Error assigning a doctor:", err)
					} else {
						cls.CLS()
						fmt.Println("Successfully assigned doctor to a room!")
					}

				case 2:
					fmt.Println("To be edited soon")
				case 3:
					var roomNumber string
					fmt.Print("Enter the room number to be deleted: ")
					fmt.Scanln(&roomNumber)

					// err := deleteRoom(roomNumber)
					if err != nil {
						cls.CLS()
						fmt.Println("Error deleting room:", err)
					} else {
						cls.CLS()
						fmt.Println("Room deleted successfully")
					}

				case 4:
					fmt.Println("Going back to Admin Menu...")
					continue OuterLoop
				default:
					fmt.Println("Invalid choice. Please try again.")
				}
			}
		case 4:
			for {
				fmt.Println("\nAccount Creation")
				fmt.Println("1. Create Account")
				fmt.Println("2. Update Account")
				fmt.Println("3. Delete Account")
				fmt.Println("4. Go back to Admin Menu")
				fmt.Print("Enter your choice: ")
				fmt.Scanln(&choice)

				switch choice {
				case 1:
					fmt.Println("\nList of Employees: ")
					err = printEmployees()
					if err != nil {
						fmt.Println("Error reading employee data:", err)
					}

					// To read the whole line, use standard input scanner
					var hp_id string
					fmt.Print("Enter ID number: ")
					scanner.Scan()
					hp_id = scanner.Text()

					var username string
					fmt.Print("Enter username: ")
					scanner.Scan()
					username = scanner.Text()

					var password string
					fmt.Print("Enter password: ")
					scanner.Scan()
					password = scanner.Text()

					// fmt.Printf("You entered room type: %s and room number: %d\n", roomType, roomNumber)

					err := addAccount(hp_id, username, password)
					if err != nil {
						cls.CLS()
						fmt.Println("Error creation:", err)
					} else {
						cls.CLS()
						fmt.Println("Successfully created an account!")
					}

				case 2:
					fmt.Println("To be edited soon")
				case 3:
					var accountId string
					fmt.Print("Enter the ID to be deleted: ")
					fmt.Scanln(&accountId)

					err := deleteRecord(accountId, "account")
					if err != nil {
						cls.CLS()
						fmt.Println("Error deleting room:", err)
					} else {
						cls.CLS()
						fmt.Println("Account deleted successfully")
					}

				case 4:
					fmt.Println("Going back to Admin Menu...")
					continue OuterLoop
				default:
					fmt.Println("Invalid choice. Please try again.")
				}
			}
		case 5:
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
		fmt.Println("1. Secretary")
		fmt.Println("2. Doctor")
		fmt.Println("3. Admin")
		fmt.Println("4. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			secretary()
		case 2:
			doctor()
		case 3:
			fmt.Print("Enter username: ")
			fmt.Scanln(&username)
			fmt.Print("Enter password: ")
			fmt.Scanln(&password)
			if username == "admin" && password == "admin" {
				adminFunction()
			} else {
				fmt.Println("Incorrect username or password. Try again.")
			}
		case 4:
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

func printRooms() error {
	db, err := connectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	rows, err := db.Query("SELECT room_number, room_type FROM tbl_rooms")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var roomNumber string
		var capacity string

		err := rows.Scan(&roomNumber, &capacity)
		if err != nil {
			return err
		}

		fmt.Printf("Room Number: %s, Capacity: %s\n", roomNumber, capacity)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return nil
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
func printEmployees() error {
	db, err := connectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	rows, err := db.Query("SELECT hp_id, last_name, first_name, middle_name, profession, specialization from tbl_employees")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var hpid, lastName, firstName, middleName, profession, specialization string

		err := rows.Scan(&hpid, &lastName, &firstName, &middleName, &profession, &specialization)
		if err != nil {
			return err
		}

		fmt.Printf("ID: %s | Employees: %s, %s %s | Profession: %s | Specialization: %s\n", hpid, lastName, firstName, middleName, profession, specialization)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}

func printDoctors() error {
	db, err := connectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	rows, err := db.Query("SELECT hp_id, last_name, first_name, middle_name, specialization from tbl_employees where profession='Doctor'")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var hpid, lastName, firstName, middleName, specialization string

		err := rows.Scan(&hpid, &lastName, &firstName, &middleName, &specialization)
		if err != nil {
			return err
		}

		fmt.Printf("ID: %s | Doctor: %s, %s %s Specialization: %s\n", hpid, lastName, firstName, middleName, specialization)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}

func generateMiliSec() string {
	// Step 1: Declare Variables
	var id string

	// Step 3: Generate Timestamp
	currentTime := time.Now().UnixNano() / int64(time.Millisecond)

	// Step 4: Format Timestamp
	formattedTime := fmt.Sprintf("%06d", currentTime) // Padding with zeroes to ensure consistent length

	// Step 5: Concatenate with Prefix
	id = "HPID-" + formattedTime

	return id
}

func addEmployee(lastName string, firstName string, middleName string, profession string, specialization string) error {
	db, err := connectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	hpid := generateMiliSec()
	uuid := uuid.New().String()

	_, err = db.Exec("INSERT INTO tbl_employees (emp_id, hp_id, last_name, first_name, middle_name, profession, specialization) VALUES (?, ?, ?, ?, ?, ?, ?)", uuid, hpid, lastName, firstName, middleName, profession, specialization)
	if err != nil {
		return err
	}

	return nil
}

func deleteRecord(identifier string, table string) error {
	db, err := connectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	var query string

	switch table {
	case "room":
		query = "DELETE FROM tbl_rooms WHERE room_number=?"
	case "employee":
		query = "DELETE FROM tbl_employees WHERE hp_id=?"
	case "account":
		query = "DELETE FROM tbl_accounts WHERE emp_id=?"
	}

	result, err := db.Exec(query, identifier)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New(identifier + " doesn't exist")
	}

	return nil
}

func assignDoctor(roomNumber string, doctorID string) error {
	db, err := connectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	// Getting of ID's before insertion
	roomID, err := getId(roomNumber, "")
	if err != nil {
		return err
	}

	doctorID, err = getId("", doctorID)
	if err != nil {
		return err
	}

	rdID := uuid.New().String()

	_, err = db.Exec("INSERT INTO tbl_room_doctor (rd_id, doctor_id_fk, room_id_fk) VALUES (?, ?, ?)", rdID, doctorID, roomID)
	if err != nil {
		return err
	}

	return nil
}

func getId(roomNumber string, doctorID string) (string, error) {
	db, err := connectDB()
	if err != nil {
		return "", err
	}
	defer db.Close()

	var id string
	var query string

	if doctorID != "" {
		query = "SELECT emp_id FROM tbl_employees WHERE hp_id = ?"
		db.QueryRow(query, doctorID).Scan(&id)
	} else if roomNumber != "" {
		query = "SELECT room_id FROM tbl_rooms WHERE room_number = ?"
		db.QueryRow(query, roomNumber).Scan(&id)
	} else {
		return "", errors.New("invalid arguments")
	}

	return id, nil
}

func printAssignedDoctor() error {
	db, err := connectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	rows, err := db.Query("SELECT CONCAT(e.last_name, ', ', e.first_name, ' ', e.middle_name) AS doctor_full_name, e.specialization, r.room_number FROM tbl_room_doctor rd JOIN tbl_employees e ON rd.doctor_id_fk = e.emp_id JOIN tbl_rooms r ON rd.room_id_fk = r.room_id;")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var drName string
		var roomNumber string
		var specialization string

		err := rows.Scan(&drName, &specialization, &roomNumber)
		if err != nil {
			return err
		}

		fmt.Printf("Doctor Name: %s| Specialization: %s| Room Number: %s\n", drName, specialization, roomNumber)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}

func addAccount(hp_id string, username string, password string) error {
	db, err := connectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	emp_id, err := getId("", hp_id)
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO tbl_accounts (emp_id, username, password) VALUES (?, ?, ?)", emp_id, username, password)
	if err != nil {
		return err
	}

	return nil
}
