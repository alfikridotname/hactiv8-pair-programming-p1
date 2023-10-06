package handler

import (
	"context"
	"fmt"
	"hactiv8-pair-programming-p1/config"
)

type ResponseListComics struct {
	ID          int
	Title       string
	Author      string
	Description string
}

func Read() {
	db, _ := config.InitDB()
	ctx := context.Background()

	rows, err := db.QueryContext(ctx, "SELECT id, title, author, description FROM comic")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	results := []ResponseListComics{}
	for rows.Next() {
		response := ResponseListComics{}
		err := rows.Scan(&response.ID, &response.Title, &response.Author, &response.Description)
		if err != nil {
			panic(err)
		}

		results = append(results, response)
	}

	if err = rows.Err(); err != nil {
		panic(err)
	}

	fmt.Println("List Comics")
	fmt.Println("--------------------------------------------------------------")
	fmt.Println("ID    | Title                     | Author          | Description    |")
	fmt.Println("--------------------------------------------------------------")
	for _, field := range results {
		fmt.Printf("%-5d | %-25s | %-15s | %-15s\n", field.ID, field.Title, field.Author, field.Description)
	}
	fmt.Println("--------------------------------------------------------------")
	fmt.Print("\nPress ENTER to return to the main menu...")
	fmt.Scanln()

}
