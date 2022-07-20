package Models

import (
	"awesomeProject/go-exercises/day4-5/Config"
	_ "github.com/go-sql-driver/mysql"
)
//add prod -
func CreateProduct(prod *Product) (err error) {
	if err = Config.DB.Create(prod).Error; err != nil {
		return err
	}
	return nil
}


//update quantity and price
func UpdateProd(prod *Product, id string) (err error) {
	//fmt.Println(prod)
	if err = Config.DB.Where("prod_id = ?", id).Save(prod).Error; err != nil {
		return err
	}
	return nil
}

//get prod, get prod by id -
func GetProdByID(prod *Product, id string) (err error) {
	if err = Config.DB.Where("prod_id = ?", id).Find(prod).Error; err != nil {
		return err
	}
	return nil
}
//GetAllUsers Fetch all user data
func GetAllProds(prod *[]Product) (err error) {
	if err = Config.DB.Find(prod).Error; err != nil {
		return err
	}
	return nil
}
//DeleteProd
func DeleteProd(prod *Product, id string) (err error) {
	Config.DB.Where("prod_id = ?", id).Delete(prod)
	return nil
}

