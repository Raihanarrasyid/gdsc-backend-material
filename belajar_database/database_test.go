package main

import (
	"belajar_database/app_database"
	"context"
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestOpenConnections(t *testing.T) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/golang_database")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("Berhasil")

}

func TestQueryInsert(t *testing.T) {
	db := app_database.GetConnection();
	defer db.Close()

	ctx := context.Background();
	result, err := db.ExecContext(ctx, "insert into customer(id, name) values('rai', 'rai')")
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	fmt.Println("success")
}

func TestQuerySelect(t *testing.T) {
	db := app_database.GetConnection();
	defer db.Close();

	ctx := context.Background();
	query := "select * from customer"
	result, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer result.Close()

	i := 1
	for result.Next() {
		var id, name, email sql.NullString;
		err := result.Scan(&id, &name, &email);
		if err != nil {
			panic(err)
		}
		fmt.Println(i)
		fmt.Println("Id : ", id.String)
		fmt.Println("Name : ", name.String)
		fmt.Println("email : ", email)
		i++;
	}
}