package Controllers

import (
	"awesomeProject/go-exercises/day4-5/Config"
	"awesomeProject/go-exercises/day4-5/Models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//CreateUser ... Create User
func CreateTransaction(c *gin.Context) {
	var user Models.Transaction
	c.BindJSON(&user)
	if (time.Now().Unix() - user.TransTime.Unix())<300{
		//user.Status = "Failed! Cooldown period of 5 minutes required"
		c.AbortWithStatusJSON(http.StatusBadRequest,gin.H{"message":"Cooldown period required, before placing new order"})

	}
	var prod Models.Product
	Config.DB.Where("prod_id = ?", user.ProductId).Find(prod)
	if user.ProductQuantity>prod.Quantity {
		//user.Status = "Failed! Cooldown period of 5 minutes required"
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message":"OUT OF STOCK!"})

	}
	user.Status = "Order placed"
	err := Models.CreateTransaction(&user)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}

}

//get all transactions
func GetAllTrans(c *gin.Context) {
	var trans []Models.Transaction
	err := Models.GetAllTrans(&trans)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, trans)
	}
}

func GetTransByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var trans Models.Transaction
	err := Models.GetTransByID(&trans, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, trans)
	}
}