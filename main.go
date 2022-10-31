package main

import "github.com/jacobd39/edteam/go_db_gorm/storage"

func main() {
	driver := storage.MySQL
	storage.New(driver)
}
