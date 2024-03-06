package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/MasterDimmy/go-cls"
	_ "github.com/go-sql-driver/mysql"
)

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
				fmt.Println("1. Assign Doctor Room")
				fmt.Println("2. Edit Doctor Room")
				fmt.Println("3. Remove Doctor Room")
				fmt.Println("4. Assign Doctor Time")
				fmt.Println("5. Edit Doctor Time")
				fmt.Println("6. Remove Doctor Time")
				fmt.Println("7. Go back to Admin Menu")
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
					fmt.Println("\nTime Slots: ")
					err = printTimeSlot()
					if err != nil {
						fmt.Println("Error Printing Time Slots!:", err)
					}

					fmt.Println("\nAssigned Doctor Room List: ")
					err = printDoctorsTemp()
					if err != nil {
						fmt.Println("Error Printing Assigned Doctors!:", err)
					}

					fmt.Println("\nAssigned Doctor Room with Time List: ")
					err = printAssignedDoctorTime()
					if err != nil {
						fmt.Println("Error Printing Assigned Doctors!:", err)
					}
					var timeId, doctorId string

					fmt.Print("Enter the Doctor's ID: ")
					fmt.Scanln(&doctorId)
					doctorId, err := getIdTemp(doctorId, "room_doctor")
					if err != nil {
						fmt.Println("Error getting doctor ID:", err)
					}

					fmt.Print("Enter the time ID: ")
					fmt.Scanln(&timeId)
					timeId, err = getIdTemp(timeId, "tbl_time")
					if err != nil {
						fmt.Println("Error getting time ID:", err)
					}
					fmt.Println(timeId)

					query := "INSERT INTO tbl_time_doctor (rd_id, time_id, status_id_fk) VALUES (?, ?, ?)"
					err = SQLManager(query, doctorId, timeId, "4b8b8801-db0e-11ee-9efc-902e16b789a2")
					if err != nil {
						fmt.Println("Error executing SQL query: ", err)
					}
					fmt.Println("Added time to doctor")
				case 7:
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