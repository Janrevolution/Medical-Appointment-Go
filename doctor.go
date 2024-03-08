package main

import (
	"bufio"
	"fmt"
	"os"
)

func doctor(empId string) {
	var choice int
	var err error
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("\nPatients for today: ")
	err = printPatientDay()
	if err != nil {
		fmt.Println("Error reading patient data:", err)
	}
	for {
		fmt.Println("---------------------------------")
		fmt.Println("| [1] Attend to a patient       |")
		fmt.Println("| [2] Edit a patient diagnosis  |")
		fmt.Println("| [3] Go back to Main Menu      |")
		fmt.Println("---------------------------------")
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var reserveId, diagnosis string
			fmt.Print("Enter the Reservation ID: ")
			fmt.Scanln(&reserveId)
			reserveId, err := getIdTemp(reserveId, "reserve")
			if err != nil {
				fmt.Println("Error getting reservation ID:", err)
			}

			fmt.Print("Enter Diagnosis: ")
			scanner.Scan()
			diagnosis = scanner.Text()

			query := "INSERT INTO tbl_patient_diagnosis (reserve_id, diagnosis, doctor_id) VALUES (?, ?, ?)"
			err = SQLManager(query, reserveId, diagnosis, empId)
			if err != nil {
				fmt.Println("Error executing SQL query: ", err)
				continue
			}
			fmt.Println("Diagnosed Patient successfully.")
		case 2:
			err = printPatientDiagnosis(empId)
			if err != nil {
				fmt.Println("Error reading patient diagnosis data:", err)
			}
			var reserveId, newDiagnosis string
			fmt.Print("Enter Reservation ID: ")
			fmt.Scanln(&reserveId)
			reserveId, err := getIdTemp(reserveId, "reserve")
			if err != nil {
				fmt.Println("Error getting reservation ID:", err)
			}

			fmt.Print("Enter Updated Diagnosis ID: ")
			scanner.Scan()
			newDiagnosis = scanner.Text()

			query := "UPDATE tbl_patient_diagnosis SET diagnosis = ? WHERE reserve_id = ?"
			err = SQLManager(query, newDiagnosis, reserveId)
			if err != nil {
				fmt.Println("Error executing SQL query: ", err)
				continue
			}
			fmt.Println("Diagnosed Patient successfully.")
		case 3:
			main()
		}
	}
}
