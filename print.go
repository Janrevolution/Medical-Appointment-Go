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

	rows, err := db.Query("SELECT room_id, room_number, room_type FROM tbl_rooms")
	if err != nil {
		return err
	}
	defer rows.Close()

	fmt.Println("\nExisting Rooms: ")
	for rows.Next() {
		var roomId, roomNumber, capacity string

		err := rows.Scan(&roomId, &roomNumber, &capacity)
		if err != nil {
			return err
		}
		roomId = strings.Split(roomId, "-")[0]
		fmt.Printf("Room ID: %s | Room Number: %s | Room Type: %s\n", roomId, roomNumber, capacity)
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

	rows, err := db.Query("SELECT emp_id, last_name, first_name, middle_name, profession, specialization from tbl_employees")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var empId, lastName, firstName, middleName, profession, specialization string

		err := rows.Scan(&empId, &lastName, &firstName, &middleName, &profession, &specialization)
		if err != nil {
			return err
		}
		empId = strings.Split(empId, "-")[0]

		fmt.Printf("Employee ID: %s | Employees: %s, %s %s | Profession: %s | Specialization: %s\n", empId, lastName, firstName, middleName, profession, specialization)
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

	rows, err := db.Query("SELECT emp_id, last_name, first_name, middle_name, specialization from tbl_employees where profession='Doctor'")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var empId, lastName, firstName, middleName, specialization string

		err := rows.Scan(&empId, &lastName, &firstName, &middleName, &specialization)
		if err != nil {
			return err
		}
		empId = strings.Split(empId, "-")[0]
		fmt.Printf("ID: %s | Doctor: %s, %s %s Specialization: %s\n", empId, lastName, firstName, middleName, specialization)
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

	rows, err := db.Query(`SELECT rd.rd_id, CONCAT(e.last_name, ', ', e.first_name, ' ', e.middle_name) AS doctor_full_name, 
			e.specialization, r.room_number FROM tbl_room_doctor rd 
			JOIN tbl_employees e ON rd.doctor_id_fk = e.emp_id 
			JOIN tbl_rooms r ON rd.room_id_fk = r.room_id;`)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var roomId, drName, roomNumber, specialization string

		err := rows.Scan(&roomId, &drName, &specialization, &roomNumber)
		if err != nil {
			return err
		}
		roomId = strings.Split(roomId, "-")[0]
		fmt.Printf("Doctor Room ID: %s | Doctor Name: %s| Specialization: %s| Room Number: %s\n", roomId, drName, specialization, roomNumber)
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

	rows, err := db.Query("SELECT e.emp_id, e.first_name, e.middle_name, e.last_name, a.username FROM tbl_accounts a JOIN tbl_employees e ON a.emp_id = e.emp_id;")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var empId, firstName, middleName, lastName, username string

		err := rows.Scan(&empId, &firstName, &middleName, &lastName, &username)
		if err != nil {
			return err
		}
		empId = strings.Split(empId, "-")[0]
		fmt.Printf("Employee ID: %s | Doctor Name: %s %s %s | Username: %s\n", empId, firstName, middleName, lastName, username)
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

		fmt.Printf("Patient ID: %s | Employee: %s, %s %s | Age: %d | Gender: %s\n", patid, lastName, firstName, middleName, age, gender)
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

	rows, err := db.Query(`
    SELECT 
        time_id,
        DATE_FORMAT(start_time, '%h:%i:%s %p') AS start_time,
        DATE_FORMAT(end_time, '%h:%i:%s %p') AS end_time
    FROM 
        tbl_time;
	`)
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
        SELECT rd.rd_id, e.last_name, e.first_name, e.middle_name,
        CONCAT(DATE_FORMAT(t.start_time, '%h:%i:%s %p'), ' - ', DATE_FORMAT(t.end_time, '%h:%i:%s %p')) AS time_slot,
        td.time_id, td.ad_id
        FROM tbl_employees e
        JOIN tbl_room_doctor rd ON e.emp_id = rd.doctor_id_fk
        JOIN tbl_time_doctor td ON rd.rd_id = td.rd_id
        JOIN tbl_time t ON td.time_id = t.time_id
        ORDER BY e.last_name, e.first_name, e.middle_name, t.start_time;
    `)
	if err != nil {
		return err
	}
	defer rows.Close()

	var prevDoctorID string
	var isNewDoctor bool = true

	for rows.Next() {
		var doctorID, lastName, firstName, middleName, timeID, timeSlot, adID string

		err := rows.Scan(&doctorID, &lastName, &firstName, &middleName, &timeSlot, &timeID, &adID)
		if err != nil {
			return err
		}

		// Extracting only the first part of the UUIDs
		doctorIDParts := strings.Split(doctorID, "-")
		timeIDParts := strings.Split(timeID, "-")
		adIDParts := strings.Split(adID, "-")

		// Checker for new Doctor
		isNewDoctor = !(doctorIDParts[0] == prevDoctorID)

		// Output new Doctor
		if isNewDoctor {
			fmt.Printf("Doctor: %s %s %s (%s)\n", firstName, middleName, lastName, doctorIDParts[0])
		}

		fmt.Printf("\t- Time Slot: %s (Time ID: %s) (Doctor Time ID: %s)\n", timeSlot, timeIDParts[0], adIDParts[0])

		// Update Doctor ID
		prevDoctorID = doctorIDParts[0]
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}

func printUnavDoctors() error {
	db, err := connectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	rows, err := db.Query(`
			SELECT e.last_name, e.first_name, e.middle_name, ad.ad_id, 
			DATE_FORMAT(ad.date, '%M %e, %Y') AS formatted_date,
			DATE_FORMAT(t.start_time, '%h:%i:%s %p') AS start_time,
			DATE_FORMAT(t.end_time, '%h:%i:%s %p') AS end_time
			FROM tbl_employees e
			JOIN tbl_room_doctor rd ON e.emp_id = rd.doctor_id_fk
			JOIN tbl_time_doctor td ON rd.rd_id = td.rd_id
			JOIN tbl_avail_doctor ad ON td.ad_id = ad.ad_id
			JOIN tbl_time t ON td.time_id = t.time_id;
	`)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var lastName, firstName, middleName, adID, formattedDate, startTime, endTime string

		err := rows.Scan(&lastName, &firstName, &middleName, &adID, &formattedDate, &startTime, &endTime)
		if err != nil {
			return err
		}

		adIDParts := strings.Split(adID, "-")

		fmt.Printf("Doctor Time ID: %s | Doctor: %s, %s, %s | %s | Start Time: %s | End Time: %s\n", adIDParts[0], lastName, firstName, middleName, formattedDate, startTime, endTime)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}

func freeTime() error {
	db, err := connectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	rows, err := db.Query(`
			SELECT CONCAT(e.first_name, ' ', e.middle_name, ' ', e.last_name) AS doctor_name, e.specialization, td.rd_id, td.time_id, DATE_FORMAT(t.start_time, '%h:%i:%s %p') AS formatted_start_time, DATE_FORMAT(t.end_time, '%h:%i:%s %p') AS formatted_end_time
			FROM tbl_employees e
			JOIN tbl_room_doctor rd ON e.emp_id = rd.doctor_id_fk
			JOIN tbl_time_doctor td ON rd.rd_id = td.rd_id
			JOIN tbl_time t ON td.time_id = t.time_id
			LEFT JOIN tbl_avail_doctor ad ON td.ad_id = ad.ad_id AND ad.date = CURDATE()
			LEFT JOIN tbl_appointment_details ap ON td.rd_id = ap.rd_id AND t.time_id = ap.time AND ap.date = CURDATE()
			WHERE ad.date IS NULL AND ap.reserve_id IS NULL;
    `)
	if err != nil {
		return err
	}
	defer rows.Close()

	var prevDoctorID string
	var isNewDoctor bool = true

	for rows.Next() {
		var doctorName, specialization, rdID, timeID, formattedStartTime, formattedEndTime string

		err := rows.Scan(&doctorName, &specialization, &rdID, &timeID, &formattedStartTime, &formattedEndTime)
		if err != nil {
			return err
		}

		// Checker for new Doctor
		isNewDoctor = !(doctorName == prevDoctorID)

		// Output new Doctor
		if isNewDoctor {
			// Split rdID and take the first part
			rdID = strings.Split(rdID, "-")[0]
			fmt.Printf("Doctor: %s (rd_id: %s) | Specialization: %s \n", doctorName, rdID, specialization)
		}

		// Split timeID and take the first part
		timeID = strings.Split(timeID, "-")[0]

		fmt.Printf("\t- Time Slot: %s - %s (Time ID: %s)\n", formattedStartTime, formattedEndTime, timeID)

		// Update Doctor ID
		prevDoctorID = doctorName
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}

func printReservedPatients() error {
	db, err := connectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	rows, err := db.Query(`
        SELECT ad.reserve_id, ad.date, 
        DATE_FORMAT(t.start_time, '%h:%i:%s %p') AS start_time, 
        DATE_FORMAT(t.end_time, '%h:%i:%s %p') AS end_time,
        t.time_id,
        p.first_name AS patient_first_name,
        p.middle_name AS patient_middle_name,
        p.last_name AS patient_last_name,
        e.first_name AS doctor_first_name,
        e.middle_name AS doctor_middle_name,
        e.last_name AS doctor_last_name
        FROM tbl_appointment_details ad
        JOIN tbl_time t ON ad.time = t.time_id
        JOIN tbl_patients p ON ad.patient_id_fk = p.patient_id
        JOIN tbl_room_doctor rd ON ad.rd_id = rd.rd_id
        JOIN tbl_employees e ON rd.doctor_id_fk = e.emp_id;
    `)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			reserveID, date, startTime, endTime,
			timeID, patientFirstName, patientMiddleName, patientLastName,
			doctorFirstName, doctorMiddleName, doctorLastName string
		)
		if err := rows.Scan(&reserveID, &date, &startTime, &endTime, &timeID,
			&patientFirstName, &patientMiddleName, &patientLastName,
			&doctorFirstName, &doctorMiddleName, &doctorLastName); err != nil {
			return err
		}
		reserveID = strings.Split(reserveID, "-")[0]
		timeID = strings.Split(timeID, "-")[0]
		fmt.Printf("Reserve ID: %s | Date: %s | Start Time: %s | End Time: %s | Time ID: %s | Patient Name: %s  %s %s | Doctor Name: %s  %s  %s\n",
			reserveID, date, startTime, endTime, timeID,
			patientFirstName, patientMiddleName, patientLastName,
			doctorFirstName, doctorMiddleName, doctorLastName)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}

// Display the patient booked for the day and don't print if it already has a diagnosis
func printPatientDay() error {
	db, err := connectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	rows, err := db.Query(`
	SELECT tad.reserve_id,
    tp.last_name,tp.first_name,tp.middle_name,
    tt.start_time, tt.end_time
	FROM tbl_appointment_details AS tad
	JOIN tbl_patients AS tp ON tad.patient_id_fk = tp.patient_id
	JOIN tbl_time AS tt ON tad.time = tt.time_id
	WHERE tad.date = CURDATE()	
	AND tad.reserve_id NOT IN (SELECT reserve_id FROM tbl_patient_diagnosis);
    `)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var reserveId, lastName, firstName, middleName,
			startTime, endTime string

		if err := rows.Scan(&reserveId, &lastName, &firstName, &middleName, &startTime, &endTime); err != nil {
			return err
		}
		reserveId = strings.Split(reserveId, "-")[0]

		fmt.Printf("Reserve ID: %s | Name: %s, %s %s | Appointment Time: %s - %s\n", reserveId, lastName, firstName, middleName, startTime, endTime)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}

func printPatientDiagnosis(doctorId string) error {
	db, err := connectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	rows, err := db.Query(`
        SELECT pd.reserve_id, tp.last_name, tp.first_name, tp.middle_name, pd.diagnosis
        FROM tbl_patient_diagnosis AS pd
        JOIN tbl_appointment_details AS tad ON pd.reserve_id = tad.reserve_id
        JOIN tbl_patients AS tp ON tad.patient_id_fk = tp.patient_id
        WHERE pd.doctor_id = ?;
    `, doctorId)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var reserveID, lastName, firstName, middleName, diagnosis string
		if err := rows.Scan(&reserveID, &lastName, &firstName, &middleName, &diagnosis); err != nil {
			return err
		}
		reserveID = strings.Split(reserveID, "-")[0]
		fmt.Printf("Reserve ID: %s | Patient Name: %s %s %s | Diagnosis: %s\n", reserveID, firstName, middleName, lastName, diagnosis)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}
