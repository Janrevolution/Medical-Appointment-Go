package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/MasterDimmy/go-cls"
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
| [4] Create Account        |
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

					err := addRoom(roomType, roomNumber)
					if err != nil {
						cls.CLS()
						fmt.Println("Error creating room:", err)
					} else {
						cls.CLS()
						fmt.Println("Room created successfully")
					}
				case 2:
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
						if !isAlphaOrSpace(firstName) {
							fmt.Println("Invalid input!")
						} else {
							break
						}
					}

					for {
						fmt.Print("Enter Specialization(N/A for non-doctors): ")
						scanner.Scan()
						specialization = scanner.Text()
						if !isAlphaOrSpace(firstName) {
							fmt.Println("Invalid input!")
						} else {
							break
						}
					}

					err := addEmployee(lastName, firstName, middleName, profession, specialization)

					if err != nil {
						cls.CLS()
						fmt.Println("Error creating user:", err)
					} else {
						cls.CLS()
						fmt.Println("User created successfully")
					}
				case 2:
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
				fmt.Print("--------------Assign Menu--------------")
				fmt.Print("| [1] Assign Doctor Room              |")
				fmt.Print("| [2] Edit Doctor Room                |")
				fmt.Print("| [3] Remove Doctor Room              |")
				fmt.Print("| [4] Assign Doctor Time              |")
				fmt.Print("| [5] Add Unavailable Doctor Time     |")
				fmt.Print("| [6] Remove Unavailable Doctor Time  |")
				fmt.Print("| [7] Go back to Admin Menu           |")
				fmt.Print("---------------------------------------")
				fmt.Print("Enter your choice")

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
					var roomNumber string
					fmt.Print("\nEnter room number: ")
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

					var roomNumber, drID string

					for {
						fmt.Print("Enter the room number to be deleted: ")
						fmt.Scanln(&roomNumber)
						roomNumber = strings.TrimSpace(roomNumber) // Remove the newline character
						if roomNumber == "" {
							continue assignMenu
						}
						break
					}

					for {
						fmt.Print("Enter Doctors ID to be deleted: ")
						fmt.Scanln(&drID)
						drID = strings.TrimSpace(drID) // Remove the newline character
						if drID == "" {
							continue assignMenu
						}
						break
					}

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

				case 5:
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

				case 6:

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
				case 7:
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
					var hp_id string
					fmt.Print("\nEnter ID number [enter if you want to go back to assign menu]: ")
					scanner.Scan()
					hp_id = scanner.Text()

					if hp_id == "" {
						continue accountCreation
					}

					var username string
					fmt.Print("Enter username [enter if you want to go back to assign menu]: ")
					scanner.Scan()
					username = scanner.Text()

					if username == "" {
						continue accountCreation
					}

					var password string
					fmt.Print("Enter password [enter if you want to go back to assign menu]: ")
					scanner.Scan()
					password = scanner.Text()
					if password == "" {
						continue accountCreation
					}

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
					if accountId == "" {
						continue accountCreation
					}
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
