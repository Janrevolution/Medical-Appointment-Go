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
3. Go Back to Main Menu
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
2. Go Back to Secretary Menu
Enter your choice: `)

			_, err = fmt.Scanln(&choice)
			if err != nil {
				fmt.Println("Error reading input: ", err)
				continue
			}

			fmt.Println("\nPatient Menu:")
			err = printPatients()
			if err != nil {
				fmt.Println("Error reading patient data:", err)
			}
			fmt.Print(`
Patient Menu:
1. Add Patients
2. Go Back to Secretary Menu
Enter your choice: `)
			fmt.Scanln(&choice)
			switch choice {
			case 1:
				err = printPatients()
				if err != nil {
					fmt.Println("Error reading patient data:", err)
				}

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
				err = SQLManager(query, uuid, lastName, firstName, middleName, age, gender)
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
				main()
			}
		case 2:
			// fmt.Println("1. Add Reservation")
			// fmt.Println("2. Edit Reservation")
			// fmt.Println("2. Delete Reservation")
			// fmt.Print("Enter your choice: ")
			// fmt.Scanln(&choice)
			// switch choice {
			// case 1:

			// }
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

func main() {
	var choice int
	var username, password string

	for {
		fmt.Print(`
Main Menu:
1. Log-in
2. Admin
3. Exit
Enter your choice: `)
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

func getIdTemp(cutId, table string) (string, error) {
	db, err := connectDB()
	if err != nil {
		return "", err
	}
	defer db.Close()

	var id string

	if table == "room_doctor" {
		query := "SELECT rd_id FROM tbl_room_doctor WHERE LEFT(rd_id, 8) = ?"
		err = db.QueryRow(query, cutId).Scan(&id)
	} else if table == "tbl_time" {
		query := "SELECT time_id FROM tbl_time WHERE LEFT(time_id, 8) = ?"
		err = db.QueryRow(query, cutId).Scan(&id)
	}

	if err != nil {
		return "", err
	}

	return id, nil
}
func checkRoomExists(roomNumber int) (bool, error) {
	db, err := connectDB()
	if err != nil {
		return false, err
	}
	defer db.Close()

	var exists bool
	err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM tbl_rooms WHERE room_number=?)", roomNumber).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}