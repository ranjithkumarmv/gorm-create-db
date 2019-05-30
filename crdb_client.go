package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var address string
var dbname string

func init() {
	flag.StringVar(&address, "address", "", "string to get db connection address")
	flag.StringVar(&dbname, "dbname", "", "string to get db name")
	flag.Parse()
}

func main() {
	db, err := gorm.Open("postgres", address)
	if err != nil {
		fmt.Println(err)
		time.Sleep(10 * time.Second)
		main()
	}
	defer db.Close()
	db = db.Exec("CREATE DATABASE " + dbname + ";")
	if db.Error != nil {
		fmt.Println(db.Error)
	} else {
		fmt.Println(dbname + " Database Created Successfully")
	}
}
