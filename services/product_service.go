package services

import (
	listProduct "api_homework_structure/models"
	"api_homework_structure/repositories"
)
func GetAllProducts()([]listProduct.Product, error){
	return repositories.GetAllProducts();
}

func GetProductByID(id uint)(*listProduct.Product,error){
	return repositories.GetProductByID(id);
}

func CreatePruct(product *listProduct.Product) error{

	return repositories.CreateProduct(product);        
	
}   

func UpdateProduct(product *listProduct.Product) error {
	return repositories.UpdateProduct(product);
}

// delete pruduct 
func DeleteProudct(id uint) error{
	return repositories.DeleteProduct(id);
}                        