package main

import (
	"fmt"
	"orm/connection"
	"orm/migrations"
	"orm/models"
	"orm/repositories"
)

func main() {
	fmt.Println("GORM")
	db := connection.GetConnection()
	migrations.Migrate(db)
	pr := repositories.NewProductRepository(db)

	p1 := pr.Save(&models.Product{
		Code:   "D42",
		Price:  12000,
		Detail: "Plastilina",
	})
	fmt.Println(p1.ID, p1.Code)
	p2 := pr.Save(&models.Product{
		Code:   "D43",
		Price:  2000,
		Detail: "Lapiz",
	})
	fmt.Println(p2.ID, p2.Code)

	p2.Detail = "Lapiz de Color Rojo"
	pr.Update(p2)
	all := pr.GetAll()
	for _, producto := range *all {
		fmt.Println(producto.ID, producto.Detail)
		pr.Delete(producto.ID)
	}
}
