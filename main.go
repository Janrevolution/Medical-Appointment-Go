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

func customerFunction() {
	fmt.Println("Hello, patient!")
}

func adminFunction() {
	scanner := bufio.NewScanner(os.Stdin)
	var choice int
	var err error
OuterLoop:
	for {
		fmt.Println("\nAdmin Menu:")
		fmt.Println("1. Rooms")
		fmt.Println("2. Doctor function")
		fmt.Println("3. Assign Doctor")
		fmt.Println("4. View Appointments")
		fmt.Println("3. Go back to Main Menu")
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

					err := deleteRoom(roomNumber)
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
				err = printDoctors()
				if err != nil {
					fmt.Println("Error reading doctor data:", err)
				}

				fmt.Println("\nDoctor Menu:")
				fmt.Println("1. Add Doctor")
				fmt.Println("2. Edit Doctor")
				fmt.Println("3. Delete room")
				fmt.Println("4. Go back to Admin Menu")
				fmt.Print("Enter your choice: ")
				fmt.Scanln(&choice)

				switch choice {
				case 1:

					// To read the whole line, use standard input scanner
					var lastName, firstName, middleName, specialization string

					fmt.Print("Enter Dr. Last Name: ")
					scanner.Scan()
					lastName = scanner.Text()

					fmt.Print("Enter Dr. First Name: ")
					scanner.Scan()
					firstName = scanner.Text()

					fmt.Print("Enter Dr. Middle Name: ")
					scanner.Scan()
					middleName = scanner.Text()

					fmt.Print("Enter Dr. Middle Name: ")
					scanner.Scan()
					specialization = scanner.Text()

					err := addDoctor(lastName, firstName, middleName, specialization)
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
					var lastName, firstName, middleName string

					fmt.Print("Enter Dr. Last Name to be deleted: ")
					scanner.Scan()
					lastName = scanner.Text()

					fmt.Print("Enter Dr. First Name to be deleted: ")
					scanner.Scan()
					firstName = scanner.Text()

					fmt.Print("Enter Dr. Middle Name to be deleted: ")
					scanner.Scan()
					middleName = scanner.Text()

					err := deleteDoctor(lastName, firstName, middleName)
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

					err := deleteRoom(roomNumber)
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

func deleteRoom(roomNumber string) error {
	db, err := connectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	// Execute the delete operation
	result, err := db.Exec("DELETE FROM tbl_rooms WHERE room_number=?", roomNumber)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("Room doesn't exist")
	}

	return nil
}

func printDoctors() error {
	db, err := connectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	rows, err := db.Query("SELECT hp_id, last_name, first_name, middle_name, specialization from tbl_doctors")
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

func addDoctor(lastName string, firstName string, middleName string, specialization string) error {
	db, err := connectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	hpid := generateMiliSec()
	uuid := uuid.New().String()

	_, err = db.Exec("INSERT INTO tbl_doctors (doctor_id, hp_id, last_name, first_name, middle_name, specialization) VALUES (?, ?, ?, ?, ?, ?)", uuid, hpid, lastName, firstName, middleName, specialization)
	if err != nil {
		return err
	}

	return nil
}

func deleteDoctor(lastName string, firstName string, middleName string) error {
	db, err := connectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	result, err := db.Exec("DELETE FROM tbl_doctors WHERE last_name=? AND first_name=? AND middle_name=?", lastName, firstName, middleName)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("Doctor doesn't exist")
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

	fmt.Printf(roomID)

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
		query = "SELECT doctor_id FROM tbl_doctors WHERE hp_id = ?"

		err = db.QueryRow(query, doctorID).Scan(&id)
	} else if roomNumber != "" {
		query = "SELECT room_id FROM tbl_rooms WHERE room_number = ?"

		err = db.QueryRow(query, roomNumber).Scan(&id)
	} else {
		return "", errors.New("invalid arguments")
	}

	return id, nil
}
