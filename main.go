package main

import (
	"bufio"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

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
			profession, profErr := GetProfession(empID)
			if profErr != nil {
				fmt.Println("Error getting profession:", profErr)
			} else {
				if profession == "Doctor" {
					doctor(empID)
				} else {
					secretary(empID)
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
----Main Menu---
| [1] Log-in   |
| [2] Admin    |
| [3] Exit     |
----------------
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
	} else if table == "time_doctor" {
		query := "SELECT ad_id FROM tbl_time_doctor WHERE LEFT(ad_id, 8) = ?"
		err = db.QueryRow(query, cutId).Scan(&id)
	} else if table == "avail_doctor" {
		query := "SELECT ad_id FROM tbl_avail_doctor WHERE LEFT(ad_id, 8) = ?"
		err = db.QueryRow(query, cutId).Scan(&id)
	} else if table == "patient" {
		query := "SELECT patient_id FROM tbl_patients WHERE LEFT(patient_id, 8) = ?"
		err = db.QueryRow(query, cutId).Scan(&id)
	} else if table == "time_doctorT" {
		query := "SELECT time_id FROM tbl_time_doctor WHERE LEFT(time_id, 8) = ?"
		err = db.QueryRow(query, cutId).Scan(&id)
	} else if table == "time_doctorRD" {
		query := "SELECT rd_id FROM tbl_time_doctor WHERE LEFT(rd_id, 8) = ?"
		err = db.QueryRow(query, cutId).Scan(&id)
	} else if table == "appoint" {
		query := "SELECT reserve_id FROM tbl_appointment_details WHERE LEFT(reserve_id, 8) = ?"
		err = db.QueryRow(query, cutId).Scan(&id)
	} else if table == "newTime" {
		query := "SELECT time_id FROM tbl_time_doctor WHERE LEFT(time_id, 8) = ?"
		err = db.QueryRow(query, cutId).Scan(&id)
	} else if table == "reserve" {
		query := "SELECT reserve_id FROM tbl_appointment_details WHERE LEFT(reserve_id, 8) = ?"
		err = db.QueryRow(query, cutId).Scan(&id)
	} else if table == "employee" {
		query := "SELECT emp_id FROM tbl_employees WHERE LEFT(emp_id, 8) = ?"
		err = db.QueryRow(query, cutId).Scan(&id)
	} else if table == "room" {
		query := "SELECT room_id FROM tbl_rooms WHERE LEFT(room_id, 8) = ?"
		err = db.QueryRow(query, cutId).Scan(&id)
	} else if table == "account" {
		query := "SELECT emp_id FROM tbl_accounts WHERE LEFT(emp_id, 8) = ?"
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

func duplicatePatient(patientID string, currentDate time.Time) bool {
	db, err := connectDB()
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return false
	}
	defer db.Close()

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM tbl_appointment_details WHERE patient_id_fk = ? AND date = ?", patientID, currentDate).Scan(&count)
	if err != nil {
		// Handle error
		fmt.Println("Error querying database:", err)
		return false
	}

	return count > 0
}
