package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const dbName = "go_demo"

type Product struct {
	Id    int
	Name  string
	Price int
}

func main() {
	myDB := connectDB()
	defer myDB.Close()

	insertHandler(myDB, 5, "A4 Paper", 250)
	insertHandler(myDB, 6, "A4 gsm Paper", 250)
	// updateHandler(myDB, "Update A3 Paper", 3)
	// deleteHandler(myDB, 3)
	queryAllHandler(myDB)
}

func insertHandler(myDB *sql.DB, itemId int, itemName string, itemPrice int) {
	insertQuery, err := myDB.Prepare("INSERT INTO products(Id, Name, Price) VALUES(?, ?, ?)")
	handleError(err)
	defer insertQuery.Close()

	// execute
	insertRes, err := insertQuery.Exec(itemId, itemName, itemPrice)
	handleError(err)

	id, err := insertRes.LastInsertId()
	handleError(err)
	fmt.Println("[INSERT] Item with id: ", id)
}

func updateHandler(myDB *sql.DB, itemUpdateName string, itemId int) {
	// Update into DB
	//prepare
	updateQuery, err := myDB.Prepare("UPDATE products SET Name=? where Id=?")
	handleError(err)

	// execute
	updateRes, err := updateQuery.Exec(itemUpdateName, itemId)
	handleError(err)

	// affected rows
	row, err := updateRes.RowsAffected()
	handleError(err)
	fmt.Println("[UPDATE] Rows Affected: ", row)
}

func queryAllHandler(myDB *sql.DB) {
	// query all data
	rows, err := myDB.Query("SELECT * FROM products")
	handleError(err)

	// declare empty products variable
	var product = Product{}

	// iterate over rows
	fmt.Println("All Products: ")
	for rows.Next() {
		err = rows.Scan(&product.Id, &product.Name, &product.Price)
		handleError(err)
		fmt.Println(product)
	}
}

func deleteHandler(myDB *sql.DB, itemId int) {
	deleteQuery, err := myDB.Prepare("DELETE FROM products WHERE Id=?")
	handleError(err)

	// delete nth product
	deleteRes, err := deleteQuery.Exec(itemId)
	handleError(err)

	// affected rows
	row, err := deleteRes.RowsAffected()
	handleError(err)
	fmt.Println("[DELETE] Rows Affected: ", row)
}

func connectDB() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "my-secret-pw"
	dbName := "go_demo"
	myDB, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	handleError(err)
	fmt.Println("MYSQL connection success")
	// defer myDB.Close()
	return myDB
}

func handleError(err error) {
	if err != nil {
		panic((err.Error()))
	}
}
