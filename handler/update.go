package handler

import (
	"context"
	"fmt"
	"hactiv8-pair-programming-p1/config"
	"log"
)

func Update() {
	var id int
	var count int
	fmt.Print("Input id comic : ")
	fmt.Scanln(&id)

	db, _ := config.InitDB()
	ctx := context.Background()
	err := db.QueryRowContext(ctx, "SELECT COUNT(*) FROM comic WHERE id = ?", id).Scan(&count)

	if err != nil {
		log.Fatal(err)
	}

	if count == 0 {
		fmt.Printf("Data dengan id %d tidak ditemukan\n", id)
		return
	}

	var title, author, description string

	fmt.Println("Update Data Comic")
	fmt.Print("Title : ")
	fmt.Scanln(&title)
	fmt.Print("Author : ")
	fmt.Scanln(&author)
	fmt.Print("Description : ")
	fmt.Scanln(&description)

	_, err = db.ExecContext(ctx, "UPDATE comic SET title = ?, author= ? , description = ? WHERE id = ?", title, author, description, id)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success update comic")

}
