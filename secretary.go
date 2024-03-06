package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

// NOT IMPERATIVE
func isAlphaOrSpace(s string) bool {
	for _, r := range s {
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
			err = freeTime()
			if err != nil {
				fmt.Println("Error reading doctor free time data:", err)
			}
			fmt.Println()
			err = printPatients()
			if err != nil {
				fmt.Println("Error reading patientsdata:", err)
			}
			fmt.Println("1. Add Reservation")
			fmt.Println("2. Edit Reservation")
			fmt.Println("2. Delete Reservation")
			fmt.Print("Enter your choice: ")
			fmt.Scanln(&choice)
			switch choice {
			case 1:
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

				fmt.Print("Enter the Patient's ID: ")
				fmt.Scanln(&patientId)
				patientId, err := getIdTemp(patientId, "patient")
				if err != nil {
					fmt.Println("Error getting patient ID:", err)
				}

				fmt.Print("Enter Room Doctor's ID: ")
				fmt.Scanln(&rdId)
				rdId, err = getIdTemp(rdId, "time_doctorRD")
				if err != nil {
					fmt.Println("Error getting time ID:", err)
				}

				fmt.Print("Enter Date for the patient(YYYY-MM-DD): ")
				fmt.Scanln(&date)

				fmt.Print("Enter Time ID: ")
				fmt.Scanln(&timeId)
				timeId, err = getIdTemp(timeId, "time_doctorT")
				if err != nil {
					fmt.Println("Error getting time ID:", err)
				}

				fmt.Print("Enter brief symptoms: ")
				fmt.Scanln(&description)

				uuid := uuid.New().String()

				query := "INSERT INTO tbl_appointment_details (reserve_id, patient_id_fk, rd_id, date, time, secretary_id, description) VALUES (?, ?, ?, ?, ?, ?, ?)"
				err = SQLManager(query, uuid, patientId, rdId, date, timeId, empId, description)
				if err != nil {
					fmt.Println("Error executing SQL query: ", err)
				}
				fmt.Println("Added time to doctor")
			}
		}
	}
}
