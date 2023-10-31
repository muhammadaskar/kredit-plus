package main

import (
	mysqldriver "github.com/muhammadaskar/kredit-plus/infrastructures/mysql_driver"
	"github.com/muhammadaskar/kredit-plus/seeders"
)

func main() {
	db := mysqldriver.InitDatabase()

	seeders.SeedConsumers(db)
	seeders.SeedTenors(db)
}
