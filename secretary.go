package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

// NOT IMPERATIVE
func isAlphaOrSpace(s string) bool {
	trimmed := strings.TrimSpace(s)
	if len(trimmed) == 0 {
		return false
	}
	for _, r := range trimmed {
		if !unicode.IsLetter(r) && !unicode.IsSpace(r) {
			return false
		}
	}
	return true
}

func secretary(empId string) {
	
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
					fmt.Print("Enter Last Name(Press enter to go back): ")
					scanner.Scan()
					lastName = scanner.Text()
					if lastName == "" {
						secretary(empId)
					}
					if !isAlphaOrSpace(lastName) {
						fmt.Println("Invalid input!")
					} else {
						break
					}
				}

				for {
					fmt.Print("Enter First Name: ")
					scanner.Scan()
					firstName = scanner.Text()
					if !isAlphaOrSpace(firstName) {
						fmt.Println("Invalid input!")
					} else {
						break
					}
				}

				for {
					fmt.Print("Enter Middle Name: ")
					scanner.Scan()
					middleName = scanner.Text()
					if !isAlphaOrSpace(middleName) {
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
				secretary(empId)
			}
		case 2:
			reservationMenu:
			for{
				fmt.Println("Free slots for the day are: ")
			err = freeTime()
			if err != nil {
				fmt.Println("Error reading doctor free time data:", err)
			}
			fmt.Println()
			err = printPatients()
			if err != nil {
				fmt.Println("Error reading patientsdata:", err)
			}
			fmt.Println("---------Reservation Menu---------")
			fmt.Println("| [1] Add Reservation            |")
			fmt.Println("| [2] Edit Reservation           |")
			fmt.Println("| [3] Delete Reservation         |")
			fmt.Println("| [4] Go back to Secretary Menu  |")
			fmt.Println("----------------------------------")
			fmt.Print(`Enter your choice: `)
			fmt.Scanln(&choice)

			switch choice {
			case 1:
				reader := bufio.NewReader(os.Stdin)
				err = freeTime()
				if err != nil {
					fmt.Println("Error reading doctor free time data:", err)
				}
				fmt.Println()
				err = printPatients()
				if err != nil {
					fmt.Println("Error reading patientsdata:", err)
				}

				var patientId, rdId, timeId, date, description string

				for{
					fmt.Print("Enter the Patient's ID: ")
					fmt.Scanln(&patientId)
					patientId, err := getIdTemp(patientId, "patient")
					if patientId == ""{
						continue reservationMenu
					}
					if err != nil {
					fmt.Println("Error getting patient ID:", err)
					}else{
						break
					}
				}

				for{
					fmt.Print("Enter Room Doctor's ID: ")
					fmt.Scanln(&rdId)
					rdId, err = getIdTemp(rdId, "time_doctorRD")
					if rdId==""{
						continue reservationMenu
					}
					if err != nil {
						fmt.Println("Error getting time ID:", err)
					}else{
						break
					}
				}
				
				for {
					fmt.Print("Enter the date to be unavailable (YYYY-MM-DD) [enter if you want to go back to assign menu]: ")
					date, _ = reader.ReadString('\n')
					date = strings.TrimSpace(date)  // Remove the newline character

					if date == "" {
						continue reservationMenu
					}
					// Check if date is in the correct format
					_, err := time.Parse("2006-01-02", date)
					if err != nil {
						fmt.Println("Invalid date format. Please enter a date in the format YYYY-MM-DD.")
						continue
					}
					break
				}

				for{
					fmt.Print("Enter Time ID: ")
					fmt.Scanln(&timeId)
					timeId, err = getIdTemp(timeId, "time_doctorT")
					if timeId == ""{
						continue reservationMenu
					}
					if err != nil {
						fmt.Println("Error getting time ID:", err)
					}else{
						break
					}
				}

				for{
					fmt.Print("Enter brief symptoms: ")
					fmt.Scanln(&description)
					if description==""{
						continue reservationMenu
					}else{
						break
					}
				}

				uuid := uuid.New().String()

				query := "INSERT INTO tbl_appointment_details (reserve_id, patient_id_fk, rd_id, date, time, secretary_id, description) VALUES (?, ?, ?, ?, ?, ?, ?)"
				err = SQLManager(query, uuid, patientId, rdId, date, timeId, empId, description)
				if err != nil {
					fmt.Println("Error executing SQL query: ", err)
				} else{
					fmt.Println("Added time to doctor")
				}
			case 2:
				err = printReservedPatients()
				if err != nil {
					fmt.Println("Error reading reserved patient data:", err)
				}

				err = freeTime()
				if err != nil {
					fmt.Println("Error reading doctor free time data:", err)
				}

				var reserveId, timeId string
				fmt.Print("Enter reserve ID to be edited: ")
				fmt.Scanln(&reserveId)
				reserveId, err = getIdTemp(reserveId, "appoint")
				if err != nil {
					fmt.Println("Error reservation ID:", err)
				}

				fmt.Print("Enter time ID to be changed to: ")
				fmt.Scanln(&timeId)
				timeId, err = getIdTemp(timeId, "newTime")
				if err != nil {
					fmt.Println("Error getting time ID:", err)
				}
				query := "UPDATE tbl_appointment_details SET time = ? WHERE reserve_id = ?"
				err = SQLManager(query, timeId, reserveId)
				if err != nil {
					fmt.Println("Error executing SQL query: ", err)
				}
				fmt.Println("Updated time for the patient!")
			case 3:
				err = printReservedPatients()
				if err != nil {
					fmt.Println("Error reading reserved patient data:", err)
				}

				var reserveId string
				fmt.Print("Enter reserve ID to be deleted: ")
				fmt.Scanln(&reserveId)
				reserveId, err = getIdTemp(reserveId, "appoint")
				if err != nil {
					fmt.Println("Error reservation ID:", err)
				}

				query := "DELETE FROM tbl_appointment_details WHERE reserve_id = ?"
				err = SQLManager(query, reserveId)
				if err != nil {
					fmt.Println("Error executing SQL query: ", err)
				}
				fmt.Println("Deleted Reservation!")
			case 4:
				secretary(empId)
			}
			}
		case 3:
			main()
		}
	}
}
