package Controllers

import (
	"awesomeProject/go-exercises/day4-5/Models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//GetUsers ... Get all users
func GetAllProds(c *gin.Context) {
	var prod []Models.Product
	err := Models.GetAllProds(&prod)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, prod)
	}
}
//CreateUser ... Create User
func CreateProduct(c *gin.Context) {
	var user Models.Product
	c.BindJSON(&user)
	err := Models.CreateProduct(&user)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func GetProdByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var user Models.Product
	err := Models.GetProdByID(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}


//UpdateUser ... Update the user information
func UpdateProd(c *gin.Context) {
	var user Models.Product
	id := c.Params.ByName("id")
	err := Models.GetProdByID(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, user)
	}
	c.BindJSON(&user)
	err = Models.UpdateProd(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//DeleteUser ... Delete the user
func DeleteProd(c *gin.Context) {
	var user Models.Product
	id := c.Params.ByName("id")
	err := Models.DeleteProd(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
