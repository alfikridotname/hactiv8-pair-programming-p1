package handler

import (
	"context"
	"fmt"
	"hactiv8-pair-programming-p1/config"
)

func Delete() {
	var id int
	fmt.Print("Input id yang akan didelete : ")
	fmt.Scanln(&id)

	db, err := config.InitDB()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	result, err := db.ExecContext(ctx, "DELETE FROM comic WHERE id = ? ", id)
	if err != nil {
		panic(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}

	if rowsAffected == 0 {
		fmt.Printf("Data dengan id %d tidak ditemukan\n", id)
		return
	}

	fmt.Println("Berhasil hapus data comic")
}

func DeleteAll()  {
	var input string
	fmt.Print("Apakah anda ingin menghapus semua data (y/n)? ")
	fmt.Scanln(&input)
	
	db, err := config.InitDB()
    if err != nil {
		panic(err)
	}
    defer db.Close()
	
	if input == "y" || input == "Y" {
		_, err := db.Exec("DELETE FROM comic")
		if err != nil {
			panic(err)
		}
		fmt.Println("Berhasil menghapus semua data comic")
	} else {
		fmt.Println("Tidak ada data yang dihapus.")
	}

	fmt.Print("\nPress ENTER to return to the main menu...")
	fmt.Scanln()
}