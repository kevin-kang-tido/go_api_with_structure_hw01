package routes

import (
	"api_homework_structure/controllers"
	"github.com/gin-gonic/gin"
)

func SetUpAllRouter() *gin.Engine{

	router := gin.Default();
    
	// grup router under prouduct 

	productGroup := router.Group("/product")
	{
		productGroup.GET("/",controllers.GetAllProducts)
		productGroup.GET("/:id",controllers.GetProductByID)
		productGroup.POST("/",controllers.CreateProduct)
		productGroup.PATCH("/:id",controllers.UpdateProduct)
		productGroup.DELETE("/:id",controllers.DeleteProduct)
	}

	// Route to trigger CouchDB replication
	router.POST("/replicate",controllers.ReplicateToCouchDB)

	return router;

}