package main

import (
	"fmt"

	"github.com/jacobd39/edteam/go_db_gorm/model"
	"github.com/jacobd39/edteam/go_db_gorm/storage"
)

func main() {
	driver := storage.MySQL
	storage.New(driver)

	storage.DB().AutoMigrate(&model.Product{}, &model.InvoiceHeader{}, &model.InvoiceItem{})

	//product1 := model.Product{
	//	Name:  "Curso de Go",
	//	Price: 120,
	//}
	//
	//obs := "Este curso está disponible"
	//
	//product2 := model.Product{
	//	Name:         "Curso de Testing",
	//	Price:        150,
	//	Observations: &obs,
	//}
	//
	//product3 := model.Product{
	//	Name:  "Curso de Python",
	//	Price: 200,
	//}
	//
	//storage.DB().Create(&product1)
	//storage.DB().Create(&product2)
	//storage.DB().Create(&product3)
	//
	//fmt.Println("Producto1 ID: ", product1.ID)
	//fmt.Println("Producto2 ID: ", product2.ID)
	//fmt.Println("Producto3 ID: ", product3.ID)

	products := make([]model.Product, 0)
	storage.DB().Find(&products)

	for _, product := range products {
		fmt.Printf("%d - %s\n", product.ID, product.Name)
	}

	product := model.Product{}
	storage.DB().First(&product, 1)
	fmt.Printf("%d - %s\n", product.ID, product.Name)

	product.Name = "Curso de BD con Go"
	product.Price = 50
	storage.DB().Save(&product)
}
