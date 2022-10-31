package storage

import (
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

//Driver of storage
type Driver string

//Drivers
const (
	MySQL    Driver = "mysql"
	Postgres Driver = "postgres"
)

//New creates the conexion with the database
func New(d Driver) {
	switch d {
	case MySQL:
		newMySQLDB()
	case Postgres:
		newPostgresDB()
	}
}

//Inicia una nueva conexión a la Base de Datos Postgress
func newPostgresDB() {
	once.Do(func() {
		//Toda esta sección se ejecutará una sola vez - patrón Singleton
		var err error
		db, err = gorm.Open(postgres.Open("postgres://jaime:kbjnfqfsfy79@localhost:5432/godb?sslmode=disable"), &gorm.Config{})
		if err != nil {
			log.Fatalf("can't open db: %v", err)
		}

		fmt.Println("conectado a postgres")
	})
}

//Inicia una nueva conexión a la Base de Datos MySQL
func newMySQLDB() {
	once.Do(func() {
		//Toda esta sección se ejecutará una sola vez - patrón Singleton
		var err error
		db, err = gorm.Open(mysql.Open("bay:bayental2019@tcp(localhost:3306)/godb?parseTime=true"), &gorm.Config{})
		if err != nil {
			log.Fatalf("can't open db: %v", err)
		}

		fmt.Println("conectado a MySQL")
	})
}

//Pool return an unique instance of db
func DB() *gorm.DB {
	return db
}
