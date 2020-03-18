package controllers

import (
	"github.com/jinzhu/gorm"
	"simple-account/api/models"
)

func CreateTransaction(tx *gorm.DB, accountId int, operationId int, amount float64) (models.Transaction, error) {

	transaction := models.Transaction{
		AccountId:       accountId,
		OperationTypeId: operationId,
		Amount:          amount,
	}
	err := tx.Create(&transaction).Error

	return transaction, err
}

//func GetTransactionList(tx *gorm.DB, ids []int, documentNumbers []string) []models.Account {


func GetTransaction(tx *gorm.DB, id int) models.Transaction {

	var transaction models.Transaction
	tx.First(&transaction, id)

	return transaction

}
