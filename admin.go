package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

func isRoomTypeValid(roomType string) bool {
	match, _ := regexp.MatchString("^[a-zA-Z]+$", roomType)
	return match
}

// Implement checkRoomExists and addRoom functions here

func clearInputBuffer() {
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}

func adminFunction() {
	scanner := bufio.NewScanner(os.Stdin)
	var choice int
	var err error
OuterLoop:
	for {
		fmt.Print(`
----------Admin Menu---------
| [1] Rooms                 |
| [2] Employee function     |
| [3] Assign Doctor         |
| [4] Account Setting       |
| [5] Go back to Main Menu  |
-----------------------------
Enter your choice: `)
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			for {
				err = printRooms()
				if err != nil {
					fmt.Println("Error reading room data:", err)
				}
				fmt.Print(`
----------Rooms Menu----------
| [1] Add room               |
| [2] Delete room            |
| [3] Go back to Admin Menu  |
------------------------------
Enter your Choice: `)
				fmt.Scanln(&choice)

				switch choice {
				case 1:
					var roomType string
					var roomNumber int
					scanner := bufio.NewScanner(os.Stdin)

					for {
						fmt.Print("Enter the room type: ")
						scanner.Scan()
						roomType = scanner.Text()

						// Check if the room type is valid
						if !isRoomTypeValid(roomType) {
							fmt.Println("Invalid Input! Room type should only contain letters.")
						} else {
							// Convert room type to uppercase
							roomType = strings.ToUpper(roomType)
							break
						}
					}

					for {
						fmt.Print("Enter the room number: ")
						n, err := fmt.Scanln(&roomNumber)

						// Check if the room number is valid
						if err != nil || n != 1 {
							fmt.Println("Invalid Input! Room number should be a number.")
							clearInputBuffer()
							continue
						}

						// Check if the room number already exists
						exists, err := checkRoomExists(roomNumber)
						if err != nil {
							fmt.Println("Error checking room:", err)
							continue
						}
						if exists {
							fmt.Println("ERROR! Room is already Existing")
							continue
						}

						break
					}

					uuid := uuid.New().String()

					query := "INSERT INTO tbl_rooms (room_id, room_number, room_type) VALUES (?,?,?)"
					err = SQLManager(query, uuid, roomNumber, roomType)
					if err != nil {
						fmt.Println("Error executing SQL query: ", err)
						continue
					}
					fmt.Println("Room added successfully.")
				case 2:
					var roomNumber string
					fmt.Print("Enter the room number to be deleted: ")
					fmt.Scanln(&roomNumber)

					query := "DELETE FROM tbl_rooms WHERE room_number=?"
					err = SQLManager(query, roomNumber)
					if err != nil {
						fmt.Println("Error executing SQL query: ", err)
						continue
					}
					fmt.Println("Room Deleted Successfully")

				case 3:
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
				fmt.Print(`
---------Employee Menu--------
| [1] Add Employee           |
| [2] Delete Employee        |
| [3] Go back to Admin Menu  |
------------------------------
Enter your Choice: `)
				fmt.Scanln(&choice)

				switch choice {
				case 1:

					// To read the whole line, use standard input scanner
					var lastName, firstName, middleName, profession, specialization string

					for {
						fmt.Print("Enter Last Name: ")
						scanner.Scan()
						lastName = scanner.Text()
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
						fmt.Print("Enter Profession: ")
						scanner.Scan()
						profession = scanner.Text()
						professionLower := strings.ToLower(profession)
						if professionLower != "doctor" && professionLower != "secretary" {
							fmt.Println("Invalid input! Profession should be either 'Doctor' or 'Secretary'.")
						} else {
							if professionLower == "doctor" {
								profession = "Doctor"
							} else if professionLower == "secretary" {
								profession = "Secretary"
							}
							break
						}
					}									

					specializations := []string{"Cardiology", "Dermatology", "Endocrinology", "Gastroenterology", "Hematology", "Immunology", "Neurology", "Oncology", "Radiology", "Orthopedics", "Pediatrics", "Psychiatry", "Urology"}

					for {
						if profession == "Doctor" {
							fmt.Println("List of Specializations:")
							for _, s := range specializations {
								fmt.Println(s)
							}
							fmt.Print("Enter Specialization: ")
							scanner.Scan()
							specialization = scanner.Text()
							validSpecialization := false
							for _, s := range specializations {
								if strings.EqualFold(s, specialization) {
									validSpecialization = true
									break
								}
							}
							if !validSpecialization {
								fmt.Println("Invalid input! Please enter a valid specialization.")
							} else {
								specialization = strings.Title(strings.ToLower(specialization)) // Convert to title case for storage
								break
							}
						} else {
							specialization = "N/A"
							break
						}
					}

					
					uuid := uuid.New().String()

					query := "INSERT INTO tbl_employees (emp_id, last_name, first_name, middle_name, profession, specialization) VALUES (?,?,?,?,?,?)"
					err = SQLManager(query, uuid, lastName, firstName, middleName, profession, specialization)
					if err != nil {
						fmt.Println("Error executing SQL query: ", err)
						continue
					}
					fmt.Println("Employee added successfully.")

				case 2:
					err = printEmployees()
					if err != nil {
						fmt.Println("Error reading employee data:", err)
					}

					var empId string

					fmt.Print("Enter Employee ID to be deleted: ")
					fmt.Scanln(&empId)
					empId, err = getIdTemp(empId, "employee")
					if err != nil {
						fmt.Println("Error reservation ID:", err)
					}

					query := "DELETE FROM tbl_employees WHERE emp_id=?"
					err = SQLManager(query, empId)
					if err != nil {
						fmt.Println("Error executing SQL query: ", err)
						continue
					}
					fmt.Println("Employee Deleted Successfully")

				case 3:
					fmt.Println("Going back to Admin Menu...")
					continue OuterLoop
				default:
					fmt.Println("Invalid choice. Please try again.")
				}
			}
		case 3:
		assignMenu:
			for {
				fmt.Println("\nAssigned Doctors: ")
				err = printAssignedDoctor()
				if err != nil {
					fmt.Println("Error deleting doctor & room data:", err)
				}

				fmt.Println("--------------Assign Menu--------------")
				fmt.Println("| [1] Assign Doctor Room              |")
				fmt.Println("| [2] Remove Doctor Room              |")
				fmt.Println("| [3] Assign Doctor Time              |")
				fmt.Println("| [4] Add Unavailable Doctor Time     |")
				fmt.Println("| [5] Remove Unavailable Doctor Time  |")
				fmt.Println("| [6] Go back to Admin Menu           |")
				fmt.Println("---------------------------------------")
				fmt.Print("Enter your choice: ")

				fmt.Scanln(&choice)

				switch choice {
				case 1:
					fmt.Println("\nRoom Data:")
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
					var roomId, doctorId string
					fmt.Print("\nEnter room ID: ")
					fmt.Scanln(&roomId)
					roomId, err = getIdTemp(roomId, "room")
					if err != nil {
						fmt.Println("Error getting doctor ID:", err)
						continue
					}

					fmt.Print("Enter doctor ID: ")
					fmt.Scanln(&doctorId)
					doctorId, err = getIdTemp(doctorId, "employee")
					if err != nil {
						fmt.Println("Error getting doctor ID:", err)
						continue
					}

					uuid := uuid.New().String()
					query := "INSERT INTO tbl_room_doctor (rd_id, doctor_id_fk, room_id_fk) VALUES (?, ?, ?)"
					err = SQLManager(query, uuid, doctorId, roomId)
					if err != nil {
						fmt.Println("Error executing SQL query: ", err)
					} else {
						fmt.Println("Assigned doctor to a room")
					}
				case 2:
					fmt.Println("\nAssigned Doctors: ")
					err = printAssignedDoctor()
					if err != nil {
						fmt.Println("Error printing doctor & room data:", err)
					}

					var doctorRoomId string

					for {
						fmt.Print("Enter the Doctor Room ID to be deleted: ")
						fmt.Scanln(&doctorRoomId)
						doctorRoomId = strings.TrimSpace(doctorRoomId) // Remove the newline character
						if doctorRoomId == "" {
							continue assignMenu
						}
						doctorRoomId, err = getIdTemp(doctorRoomId, "room_doctor")
						if err != nil {
							fmt.Println("Error getting doctor ID:", err)
							continue
						}
						break
					}

					query := "DELETE FROM tbl_room_doctor WHERE rd_id = ?"
					err = SQLManager(query, doctorRoomId)
					if err != nil {
						fmt.Println("Error executing SQL query: ", err)
					} else {
						fmt.Println("Deleted the doctor from the room!")
					}
				case 3:
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

					for {
						fmt.Print("\nEnter the Doctor's Room ID  [enter if you want to go back to assign menu]: ")
						fmt.Scanln(&doctorId)
						doctorId = strings.TrimSpace(doctorId) // Remove the newline character
						if doctorId == "" {
							continue assignMenu
						}
						doctorId, err = getIdTemp(doctorId, "room_doctor")
						if err != nil {
							fmt.Println("Error getting doctor ID:", err)
							continue
						}
						break
					}

					for {
						fmt.Print("Enter the time ID  [enter if you want to go back to assign menu]: ")
						fmt.Scanln(&timeId)
						timeId = strings.TrimSpace(timeId) // Remove the newline character
						if timeId == "" {
							continue assignMenu
						}
						timeId, err = getIdTemp(timeId, "tbl_time")
						if err != nil {
							fmt.Println("Error getting time ID:", err)
							continue
						}
						break
					}

					uuid := uuid.New().String()
					query := "INSERT INTO tbl_time_doctor (rd_id, time_id, ad_id) VALUES (?, ?, ?)"
					err = SQLManager(query, doctorId, timeId, uuid)
					if err != nil {
						fmt.Println("Error executing SQL query: ", err)
					} else {
						fmt.Println("Added time to doctor")
					}
				case 4:
					fmt.Println("\nAssigned Doctors: ")
					err = printAssignedDoctor()
					if err != nil {
						fmt.Println("Error printing doctor & room data:", err)
					}

					fmt.Println("\nAssigned Doctor Room with Time List: ")
					err = printAssignedDoctorTime()
					if err != nil {
						fmt.Println("Error Printing Assigned Doctors!:", err)
					}

					var adId, date string

					for {
						fmt.Print("Enter the Doctor's Time ID whose time to be unavailable [enter if you want to go back to assign menu]: ")
						fmt.Scanln(&adId)
						adId = strings.TrimSpace(adId) // Remove the newline character
						if adId == "" {
							continue assignMenu
						}
						adId, err = getIdTemp(adId, "time_doctor")
						if err != nil {
							fmt.Println("Invalid Input! Error getting time doctor ID:", err)
							continue
						}
						break
					}

					reader := bufio.NewReader(os.Stdin)

					for {
						fmt.Print("Enter the date to be unavailable (YYYY-MM-DD) [enter if you want to go back to assign menu]: ")
						date, _ = reader.ReadString('\n')
						date = strings.TrimSpace(date) // Remove the newline character

						if date == "" {
							continue assignMenu
						}
						// Check if date is in the correct format
						_, err := time.Parse("2006-01-02", date)
						if err != nil {
							fmt.Println("Invalid date format. Please enter a date in the format YYYY-MM-DD.")
							continue
						}
						break
					}
					query := "INSERT INTO tbl_avail_doctor (ad_id, date) VALUES (?, ?)"
					err = SQLManager(query, adId, date)
					if err != nil {
						fmt.Println("Error executing SQL query: ", err)
					} else {
						fmt.Println("Successfully made the Time to be unavailable!")
					}
				case 5:
					fmt.Println("\nDoctor Unavailable List: ")
					err = printUnavDoctors()
					if err != nil {
						fmt.Println("Error Printing Unavailable Doctors!:", err)
					}

					var adId string

					for {
						fmt.Print("Enter the Doctor's Time ID whose date and time is to be removed: ")
						fmt.Scanln(&adId)
						adId = strings.TrimSpace(adId) // Remove the newline character
						if adId == "" {
							continue assignMenu
						}
						adId, err = getIdTemp(adId, "avail_doctor")
						if err != nil {
							fmt.Println("Error getting time doctor ID:", err)
							continue
						}
						break
					}

					query := "DELETE FROM tbl_avail_doctor WHERE ad_id = ?"
					err = SQLManager(query, adId)
					if err != nil {
						fmt.Println("Error executing SQL query: ", err)
					} else {
						fmt.Println("Made the slot available again!")
					}
				case 6:
					fmt.Println("Going back to Admin Menu...")
					continue OuterLoop
				default:
					fmt.Println("Invalid choice. Please try again.")
				}
			}
		case 4:
		accountCreation:
			for {
				fmt.Println("Accounts :")
				err = printAccounts()
				if err != nil {
					fmt.Println("Error reading room data:", err)
				}
				fmt.Print(`
-------Account Creation-------
| [1] Create Account         |
| [2] Update Account         |
| [3] Delete Account         |
| [4] Go back to Admin Menu  |
------------------------------
Enter your Choice: `)

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
					var empId, username, password string
					fmt.Print("\nEnter Employee ID number [enter if you want to go back to assign menu]: ")
					fmt.Scanln(&empId)
					empId, err = getIdTemp(empId, "employee")
					if err != nil {
						fmt.Println("Error getting doctor ID:", err)
						continue accountCreation
					}

					if empId == "" {
						continue accountCreation
					}

					fmt.Print("Enter username [enter if you want to go back to assign menu]: ")
					fmt.Scanln(&username)

					if username == "" {
						continue accountCreation
					}

					fmt.Print("Enter password [enter if you want to go back to assign menu]: ")
					fmt.Scanln(&password)
					if password == "" {
						continue accountCreation
					}

					query := "INSERT INTO tbl_accounts (emp_id, username, password) VALUES (?, ?, ?)"
					err = SQLManager(query, empId, username, password)
					if err != nil {
						fmt.Println("Error executing SQL query: ", err)
					} else {
						fmt.Println("Created an account for the employee!")
					}
				case 2:
					fmt.Println("Accounts :")
					err = printAccounts()
					if err != nil {
						fmt.Println("Error reading room data:", err)
					}

					var empId, username, password string
					fmt.Print("\nEnter Employee ID number you want to edit: ")
					fmt.Scanln(&empId)
					empId, err = getIdTemp(empId, "employee")
					if err != nil {
						fmt.Println("Error getting doctor ID:", err)
						continue accountCreation
					}

					fmt.Print("Enter a new username: ")
					fmt.Scanln(&username)
					if username == "" {
						continue accountCreation
					}

					fmt.Print("Enter a new password: ")
					fmt.Scanln(&password)
					if password == "" {
						continue accountCreation
					}

					query := "UPDATE tbl_accounts SET username = ?, password = ? WHERE emp_id = ?"
					err = SQLManager(query, username, password, empId)
					if err != nil {
						fmt.Println("Error executing SQL query: ", err)
					}
					fmt.Println("Updated account settings for the patient!")
				case 3:
					fmt.Println("Accounts :")
					err = printAccounts()
					if err != nil {
						fmt.Println("Error reading room data:", err)
					}
					var accountId string
					fmt.Print("Enter the account ID to be deleted: ")
					fmt.Scanln(&accountId)
					if accountId == "" {
						continue accountCreation
					}
					accountId, err = getIdTemp(accountId, "account")
					if err != nil {
						fmt.Println("Error getting doctor ID:", err)
						continue accountCreation
					}

					query := "DELETE FROM tbl_accounts WHERE emp_id = ?"
					err = SQLManager(query, accountId)
					if err != nil {
						fmt.Println("Error executing SQL query: ", err)
					}
					fmt.Println("Deleted the account!")
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
