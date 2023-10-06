package handler

import (
	"context"
	"fmt"
	"hactiv8-pair-programming-p1/config"
)

func Create() {
	fmt.Println("Add Comic")
	var title string
	var author string
	var description string

	fmt.Print("Title: ")
	fmt.Scanln(&title)
	fmt.Print("Author: ")
	fmt.Scanln(&author)
	fmt.Print("Description: ")
	fmt.Scanln(&description)

	db, _ := config.InitDB()

	ctx := context.Background()

	_, err := db.ExecContext(ctx, "INSERT INTO comic (title, author, description) VALUES (?, ?, ?)", title, author, description)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success add comic")

}
