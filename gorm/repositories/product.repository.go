package repositories

import (
	"orm/models"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	pr := &ProductRepository{}
	pr.db = db
	return pr
}
func (pr *ProductRepository) Get(id int) *models.Product {
	product := &models.Product{}
	pr.db.First(product, id)
	return product
}
func (pr *ProductRepository) GetAll() *[]models.Product {
	list := &[]models.Product{}
	pr.db.Find(list)
	return list

}
func (pr *ProductRepository) Save(product *models.Product) *models.Product {
	result := pr.db.Create(product)
	if result.Error != nil {
		panic("Ocurrio un error al insertar el Producto")
	}
	return product
}
func (pr *ProductRepository) Update(product *models.Product) bool {
	result := pr.db.Model(&product).Updates(product)
	return result.Error == nil
}
func (pr *ProductRepository) Delete(id uint) bool {
	result := pr.db.Delete(&models.Product{}, id)
	return result.Error == nil
}
