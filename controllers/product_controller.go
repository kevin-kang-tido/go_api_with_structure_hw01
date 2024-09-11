package controllers

import (
	"api_homework_structure/base"
	"api_homework_structure/models"
	"api_homework_structure/services"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	// "github.com/go-playground/locales/id"
)

func GetAllProducts(handleRequest*gin.Context){

	products, error := services.GetAllProducts();

	if error != nil {
	
		base.ErrorMessage(handleRequest,http.StatusInternalServerError,"Something want wrong!")
		return
	}

	base.SendSuccess(handleRequest,http.StatusOK,products)

}

// find by id 
func GetProductByID(handleRequest *gin.Context){
    id, _:= strconv.Atoi(handleRequest.Param("id"))  // change string to ingter 
	product, errr := services.GetProductByID(uint(id))
	
	if errr !=  nil {

		base.ErrorMessage(handleRequest,http.StatusNotFound,"Product has been not found!")
		
		return
	}

	base.SendSuccess(handleRequest,http.StatusOK,product)
}

// create 
func CreateProduct(handleRequest *gin.Context){
	var  product models.Product

	if err:= handleRequest.ShouldBindJSON(&product);
	err != nil {
		base.ErrorMessage(handleRequest,http.StatusBadRequest,"Something went wrong!, Please check your JSON format again!")
		return
	}

	// handle name and price is blank 
	if product.Name ==""{
		base.ErrorMessage(handleRequest,http.StatusBadRequest,"Product Nameis Requrid!")
		return
	}

	if  product.Price <=0{
		base.ErrorMessage(handleRequest,http.StatusBadRequest,"Price is Required!")
		return
	}
	// handle when internal server error 
	if  err:= services.CreatePruct(&product); err != nil {
		base.ErrorMessage(handleRequest,http.StatusInternalServerError,"Something went wrong!")
		return
		
	}

	base.SendSuccess(handleRequest,http.StatusCreated,product)

}

// update proudct by id 
func UpdateProduct(handleRequest *gin.Context){
	
    id, _ := strconv.Atoi(handleRequest.Param("id"))
	product,err := services.GetProductByID(uint(id));

	if err != nil {
		base.ErrorMessage(handleRequest,http.StatusNotFound,"Product has been not found!")
		return
	}
	// handle error 
	if err:= handleRequest.ShouldBindJSON(product);
	err != nil {
		base.ErrorMessage(handleRequest,http.StatusBadRequest,"Something went wrong!, Please check your JSON format again!")
		return
	}

	// handle name and price is blank 
	if product.Name ==" "{
		base.ErrorMessage(handleRequest,http.StatusBadRequest,"Product Nameis Requrid!")
		return
	}

	if  product.Price <= 0{
		base.ErrorMessage(handleRequest,http.StatusBadRequest,"Price is Required!")
		return
	}
	// handle when internal server error 
	if  err:= services.UpdateProduct(product); err != nil {
		base.SendError(handleRequest,http.StatusInternalServerError,"Something went wrong!",err)
		return
	}
   
	base.UpdateMessage(handleRequest,http.StatusOK,"Product has been update sueccssfully!",product)

}

// handle detele product 
func DeleteProduct(handleDeleteProduct *gin.Context){
	
	id, _ := strconv.Atoi(handleDeleteProduct.Param("id"))
    
	if err := services.DeleteProudct(uint(id)); err != nil {
		base.ErrorMessage(handleDeleteProduct,http.StatusNotFound,"Product has been not found!")
	    return
	}

	base.SendDeleteSuccess(handleDeleteProduct,http.StatusNoContent,"Product has been Delete Successfully!")
}
