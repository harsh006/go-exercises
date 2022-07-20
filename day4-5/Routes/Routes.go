package Routes

import (
	"awesomeProject/go-exercises/day4-5/Controllers"

	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/prod-api")
	{
		grp1.GET("prod", Controllers.GetAllProds)
		grp1.POST("prod", Controllers.CreateProduct)
		grp1.GET("prod/:id", Controllers.GetProdByID)
		grp1.PUT("prod/:id", Controllers.UpdateProd)
		grp1.DELETE("prod/:id", Controllers.DeleteProd)
	}
	grp2 := r.Group("/trans-api")
	{
		grp2.GET("trans", Controllers.GetAllTrans)
		grp2.POST("trans", Controllers.CreateTransaction)
		grp2.GET("trans/:id", Controllers.GetTransByID)
		//grp2.PUT("prod/:id", Controllers.Update)
		//grp2.DELETE("prod/:id", Controllers.DeleteProd)
	}
	return r
}
