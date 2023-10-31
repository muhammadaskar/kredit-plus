package main

import (
	"github.com/muhammadaskar/kredit-plus/domains"
	mysqldriver "github.com/muhammadaskar/kredit-plus/infrastructures/mysql_driver"
)

func main() {
	db := mysqldriver.InitDatabase()
	db.AutoMigrate(&domains.Consumer{}, &domains.Tenor{}, &domains.Transaction{})
}
