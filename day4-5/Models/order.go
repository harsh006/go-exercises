package Models

import "awesomeProject/go-exercises/day4-5/Config"

//create transactions
func CreateTransaction(trans *Transaction) (err error) {
	if err = Config.DB.Create(trans).Error; err != nil {
		return err
	}
	return nil
}
//get all transactions
func GetAllTrans(trans *[]Transaction) (err error) {
	if err = Config.DB.Find(trans).Error; err != nil {
		return err
	}
	return nil
}
//get transactions of particular customer

func GetTransByID(trans *Transaction, name string) (err error) {
	if err = Config.DB.Where("name = ?", name).Find(trans).Error; err != nil {
		return err
	}
	return nil
}