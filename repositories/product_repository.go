package repositories

import (
	"api_homework_structure/config"
	"api_homework_structure/models"
	listProduct "api_homework_structure/models"
)

func GetAllProducts() ([]listProduct.Product,error){

	var products []listProduct.Product

	result := config.DB.Find(&products)

	return products, result.Error;
}

func GetProductByID(id uint) (*listProduct.Product,error){

	var product models.Product

	result := config.DB.First(&product,id)

	if result.Error != nil {
		return nil,result.Error
	}

	return &product,nil;
}

// create product 
 
func CreateProduct(product *listProduct.Product) error{

	result := config.DB.Create(product)
	
	return result.Error; 
}

// update 
func UpdateProduct(product  *listProduct.Product) error{

	result := config.DB.Save(product)

	return result.Error;
}

// delete product with id 

func DeleteProduct(id uint) error{
	result := config.DB.Delete(&listProduct.Product{},id)
	return result.Error;
}




