package main

import (
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

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

	fmt.Println("\nExisting Rooms: ")
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

func printTimeSlot() error {
	db, err := connectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM tbl_time")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var timeID, startTime, endTime string

		err := rows.Scan(&timeID, &startTime, &endTime)
		if err != nil {
			return err
		}
		timeID = strings.Split(timeID, "-")[0]

		fmt.Printf("Time ID: %s | Start Time: %s | End Time: %s\n", timeID, startTime, endTime)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}

func printDoctorsTemp() error {
	db, err := connectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	rows, err := db.Query("SELECT rd.rd_id, e.last_name, e.first_name, e.middle_name, e.specialization, r.room_number FROM tbl_room_doctor rd INNER JOIN tbl_employees e ON rd.doctor_id_fk = e.emp_id INNER JOIN tbl_rooms r ON rd.room_id_fk = r.room_id")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var drID, lastName, firstName, middleName, specialization, roomNumber string

		err := rows.Scan(&drID, &lastName, &firstName, &middleName, &specialization, &roomNumber)
		if err != nil {
			return err
		}

		drID = strings.Split(drID, "-")[0]

		fmt.Printf("ID: %s | Doctor: %s, %s %s | Specialization: %s | Room Number: %s\n", drID, lastName, firstName, middleName, specialization, roomNumber)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}

func printAssignedDoctorTime() error {
	db, err := connectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	rows, err := db.Query(`
		SELECT e.last_name, e.first_name, e.middle_name,
			CONCAT(DATE_FORMAT(t.start_time, '%h:%i:%s %p'), ' - ', DATE_FORMAT(t.end_time, '%h:%i:%s %p')) AS time_slot,
			s.status_name
		FROM tbl_employees e
		JOIN tbl_room_doctor rd ON e.emp_id = rd.doctor_id_fk
		JOIN tbl_time_doctor td ON rd.rd_id = td.rd_id
		JOIN tbl_time t ON td.time_id = t.time_id
		JOIN tbl_status s ON td.status_id_fk = s.status_id
		ORDER BY e.last_name, e.first_name, e.middle_name, t.start_time;
	`)
	if err != nil {
		return err
	}
	defer rows.Close()

	var prevLastName, prevFirstName, prevMiddleName string
	var isNewDoctor bool = true

	for rows.Next() {
		var lastName, firstName, middleName, timeSlot, statusName string

		err := rows.Scan(&lastName, &firstName, &middleName, &timeSlot, &statusName)
		if err != nil {
			return err
		}

		// Checker for new Doctor
		isNewDoctor = !(lastName == prevLastName && firstName == prevFirstName && middleName == prevMiddleName)

		// Output new Doctor
		if isNewDoctor {
			fmt.Printf("Doctor: %s, %s %s\n", lastName, firstName, middleName)
		}

		fmt.Printf("\t- %s (%s)\n", timeSlot, statusName)

		// Upate Doctor Name
		prevLastName = lastName
		prevFirstName = firstName
		prevMiddleName = middleName
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}
