package main

import (
	"fmt"

	"github.com/jacobd39/edteam/go_db_gorm/model"
	"github.com/jacobd39/edteam/go_db_gorm/storage"
)

func main() {
	driver := storage.MySQL
	storage.New(driver)

	//storage.DB().AutoMigrate(&model.Product{}, &model.InvoiceHeader{}, &model.InvoiceItem{})

	//Esto ya no es necesario - Queda para consulta si se usa una versión anterior de Gorm
	//storage.DB().Model(&model.InvoiceItem{}).AddForeignKey("product_id", "products(id)", "RESTRICT", "RESTRICT")
	//storage.DB().Model(&model.InvoiceItem{}).AddForeignKey("invoice_header_id", "invoice_headers(id)", "RESTRICT", "RESTRICT")

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

	//products := make([]model.Product, 0)
	//storage.DB().Find(&products)
	//
	//for _, product := range products {
	//	fmt.Printf("%d - %s\n", product.ID, product.Name)
	//}

	//product := model.Product{}
	//if result := storage.DB().First(&product, 1); result.Error != nil {
	//	fmt.Println("Registro no encontrado")
	//	return
	//}

	//fmt.Printf("%d - %s\n", product.ID, product.Name)

	//product.Name = "Curso de BD con Go"
	//product.Price = 50
	//storage.DB().Save(&product)
	//storage.DB().Model(&product).Updates(model.Product{
	//	Name:  "Curso de POO en Go",
	//	Price: 70,
	//})

	//Borrado suave
	//storage.DB().Delete(&product)

	//Borrado permanente
	//product := model.Product{}
	//product.ID = 3
	//storage.DB().Unscoped().Delete(&product)

	invoice := model.InvoiceHeader{
		Client: "Juan Pérez",
		InvoiceItems: []model.InvoiceItem{
			{ProductID: 1},
			{ProductID: 2},
		},
	}

	if result := storage.DB().Create(&invoice); result.Error != nil {
		fmt.Println("Error al crear la factura")
		return
	}
}
