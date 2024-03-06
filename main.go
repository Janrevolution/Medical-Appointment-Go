package main

import (
	"bufio"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/MasterDimmy/go-cls"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

func isAlpha(s string) bool {
	reg := regexp.MustCompile(`^[a-zA-Z]+$`)
	return reg.MatchString(s)
}

func secretary() {
	scanner := bufio.NewScanner(os.Stdin)
	var choice int
	var err error
	for {
		fmt.Print(`
Secretary Menu:
1. Patients
2. Reservation
3. Go Back
Enter your choice: `)
		_, err = fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Error reading input: ", err)
			continue
		}

		switch choice {
		case 1:
			fmt.Print(`
Patient Menu:
1. Add Patients
2. Edit Patient
3. Go Back to Admin Menu
Enter your choice: `)

			_, err = fmt.Scanln(&choice)
			if err != nil {
				fmt.Println("Error reading input: ", err)
				continue
			}

			switch choice {
			case 1:
				var lastName, firstName, middleName, gender string
				var age int

				for {
					fmt.Print("Enter Last Name: ")
					scanner.Scan()
					lastName = scanner.Text()
					if !isAlpha(lastName) {
						fmt.Println("Invalid input!")
					} else {
						break
					}
				}

				for {
					fmt.Print("Enter First Name: ")
					scanner.Scan()
					firstName = scanner.Text()
					if !isAlpha(firstName) {
						fmt.Println("Invalid input!")
					} else {
						break
					}
				}

				for {
					fmt.Print("Enter Middle Name: ")
					scanner.Scan()
					middleName = scanner.Text()
					if !isAlpha(middleName) {
						fmt.Println("Invalid input!")
					} else {
						break
					}
				}

				for {
					fmt.Print("Enter Age: ")
					_, err = fmt.Scanf("%d\n", &age)
					if err != nil {
						fmt.Println("Invalid input! Age should only be a number.")
						// Clear the input buffer
						reader := bufio.NewReader(os.Stdin)
						_, _ = reader.ReadString('\n')
					} else {
						break
					}
				}
				
				for {
					fmt.Print("Enter Gender: ")
					scanner.Scan()
					gender = strings.ToLower(scanner.Text())
					if gender != "male" && gender != "female" {
						fmt.Println("Invalid input! Gender should be either 'male' or 'female'.")
					} else {
						break
					}
				}

				uuid := uuid.New().String()

				query := "INSERT INTO tbl_patients (patient_id, last_name, first_name, middle_name, age, gender) VALUES (?, ?, ?, ?, ?, ?)"
				err := SQLManager(query, uuid, lastName, firstName, middleName, age, gender)
				if err != nil {
					fmt.Println("Error executing SQL query: ", err)
					continue
				}

				fmt.Println("Patient added successfully.")
			case 2:
				err = printPatients()
				if err != nil {
					fmt.Println("Error reading employee data:", err)
					continue
				}
			case 3:
				secretary()
			}
		}
	}
}

// Insert, Update, Delete
func SQLManager(query string, args ...interface{}) error {
	db, err := connectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(args...)
	if err != nil {
		return err
	}

	return nil
}

func doctor() {
	fmt.Println("Welcome Doc!")
}
func login() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\nLog-in Menu:")
		var username string
		fmt.Print("Enter Username (Hit enter to go back): ")
		scanner.Scan()
		username = scanner.Text()
		if username == "" {
			main()
		}

		var password string
		fmt.Print("Enter Password: ")
		scanner.Scan()
		password = scanner.Text()

		err, empID := loginUtil(username, password)
		if err != nil {
			fmt.Println("Login failed:", err)
		} else {
			fmt.Println("Login successful!")
			fmt.Println("Employee ID:", empID)
			profession, profErr := GetProfession(empID)
			if profErr != nil {
				fmt.Println("Error getting profession:", profErr)
			} else {
				if profession == "Doctor" {
					doctor()
				} else {
					secretary()
				}
			}
		}
	}
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

				fmt.Println("\nAssigned Doctors: ")
				err = printAssignedDoctor()
				if err != nil {
					fmt.Println("Error deleting doctor & room data:", err)
				}
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
					fmt.Println("\nAssigned Doctors: ")
					err = printAssignedDoctor()
					if err != nil {
						fmt.Println("Error deleting doctor & room data:", err)
					}

					var roomNumber string
					fmt.Print("Enter the room number to be deleted: ")
					fmt.Scanln(&roomNumber)

					var drID string
					fmt.Print("Enter Doctors ID to be deleted: ")
					fmt.Scanln(&drID)

					rd_id, err := getId(roomNumber, drID)
					if err != nil {
						fmt.Println("Error getting ID:", err)
						return
					}

					err = deleteRecord(rd_id, "assignment")
					if err != nil {
						cls.CLS()
						fmt.Println("Error deleting assignment:", err)
					} else {
						cls.CLS()
						fmt.Println("Room deleted assignment")
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
				fmt.Println("Accounts :")
				err = printAccounts()
				if err != nil {
					fmt.Println("Error reading room data:", err)
				}

				fmt.Println("\nAccount Creation")
				fmt.Println("1. Create Account")
				fmt.Println("2. Update Account")
				fmt.Println("3. Delete Account")
				fmt.Println("4. Go back to Admin Menu")
				fmt.Print("Enter your choice: ")
				fmt.Scanln(&choice)

				switch choice {
				case 1:
					fmt.Println("Accounts :")
					err = printAccounts()
					if err != nil {
						fmt.Println("Error reading room data:", err)
					}

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
		fmt.Println("1. Log-in")
		fmt.Println("2. Admin")
		fmt.Println("3. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			login()
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

// func SQLManager(String SQL){
// 	db, err := connectDB()
// 	db.Open()

// 	db.Exec(SQL, identifier)

// 	db.Close()
// 	return nil
// }

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
	case "assignment":
		query = "DELETE FROM tbl_room_doctor WHERE rd_id=?"
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

func loginUtil(username, password string) (error, string) {
	db, err := connectDB()
	if err != nil {
		return err, ""
	}
	defer db.Close()

	// No Username
	var storedPassword, empID string
	err = db.QueryRow("SELECT password, emp_id FROM tbl_accounts WHERE username = ?", username).Scan(&storedPassword, &empID)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("username not found"), ""
		}
		return err, ""
	}

	// Wrong Password
	if storedPassword != password {
		return errors.New("incorrect password"), ""
	}
	return nil, empID
}

func GetProfession(empId string) (string, error) {
	db, err := connectDB()
	if err != nil {
		return "", err
	}
	defer db.Close()

	var profession string
	err = db.QueryRow("SELECT profession FROM tbl_employees WHERE emp_id = ?", empId).Scan(&profession)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("profession not found")
		}
		return "", err
	}
	return profession, nil
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

func getId(roomNumber string, doctorID string) (string, error) {
	db, err := connectDB()
	if err != nil {
		return "", err
	}
	defer db.Close()

	var id string
	var query string
	if doctorID != "" && roomNumber != "" {
		// First, retrieve the emp_id from tbl_employees
		var empID string
		query = "SELECT emp_id FROM tbl_employees WHERE hp_id = ?"
		err = db.QueryRow(query, doctorID).Scan(&empID)
		if err != nil {
			return "", err
		}

		// Then, retrieve the room_id from tbl_rooms
		var roomID string
		query = "SELECT room_id FROM tbl_rooms WHERE room_number = ?"
		err = db.QueryRow(query, roomNumber).Scan(&roomID)
		if err != nil {
			return "", err
		}

		// Finally, retrieve the rd_id from tbl_room_doctor using both emp_id and room_id
		query = "SELECT rd_id FROM tbl_room_doctor WHERE doctor_id_fk = ? AND room_id_fk = ?"
		err = db.QueryRow(query, empID, roomID).Scan(&id)
		if err != nil {
			return "", err
		}
	} else if doctorID != "" {
		query = "SELECT emp_id FROM tbl_employees WHERE hp_id = ?"
		err = db.QueryRow(query, doctorID).Scan(&id)
	} else if roomNumber != "" {
		query = "SELECT room_id FROM tbl_rooms WHERE room_number = ?"
		err = db.QueryRow(query, roomNumber).Scan(&id)
	} else {
		return "", errors.New("invalid arguments")
	}

	if err != nil {
		return "", err
	}

	return id, nil
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

func printAccounts() error {
	db, err := connectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	rows, err := db.Query("SELECT e.hp_id, a.username FROM tbl_accounts a JOIN tbl_employees e ON a.emp_id = e.emp_id;")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var hpid, username string

		err := rows.Scan(&hpid, &username)
		if err != nil {
			return err
		}

		fmt.Printf("ID: %s | Username: %s\n", hpid, username)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}

func printPatients() error {
	db, err := connectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM tbl_patients")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var patid, lastName, firstName, middleName, gender string
		var age int

		err := rows.Scan(&patid, &lastName, &firstName, &middleName, &age, &gender)
		if err != nil {
			return err
		}

		patid = strings.Split(patid, "-")[0]

		fmt.Printf("ID: %s | Employee: %s, %s %s | Age: %d | Gender: %s\n", patid, lastName, firstName, middleName, age, gender)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}
