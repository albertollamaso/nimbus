package main

import (
	"fmt"

	"github.com/albertollamaso/nimbus"
)

func main() {

	// Init the database
	db := nimbus.InitDB("db")

	defer db.Close()

	// Add a key value pair
	nimbus.AddKV(db, "me@gmail.com", "myWeakPassword")

	// Read the value for a key
	value := nimbus.ReadKV(db, "me@gmail.com")
	fmt.Println("password is: ", value)
}
