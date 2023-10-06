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
	_, err = db.ExecContext(ctx, "DELETE FROM comic WHERE id = ? ", id)
	if err != nil {
		panic(err)
	}

	fmt.Println("Berhasil hapus data comic")
}
