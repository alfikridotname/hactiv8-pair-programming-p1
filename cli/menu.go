package cli

import (
	"fmt"
	"hactiv8-pair-programming-p1/handler"
)

func Run() {
	for {
		fmt.Println("=====================================")
		fmt.Println("Welcome to Comic Book Store")
		fmt.Println("=====================================")
		fmt.Println("1. Add Comic")
		fmt.Println("2. Update Comic")
		fmt.Println("3. Delete Comic")
		fmt.Println("4. Delete All Data")
		fmt.Println("5. Show Comic")
		fmt.Println("6. Exit")
		fmt.Println("7. Back to Main Menu")
		fmt.Println("=====================================")

		var input int
		fmt.Print("Choose menu: ")
		fmt.Scanln(&input)

		switch input {
		case 1:
			handler.Create()
		case 2:
			handler.Update()
		case 3:
			handler.Delete()
		case 4:
			handler.DeleteAll()
		case 5:
			handler.Read()
		case 6:
			fmt.Println("Exit")
			return
		case 7:
			fmt.Println("Back to Main Menu")
		default:
			fmt.Println("Wrong input")
		}
	}
}
