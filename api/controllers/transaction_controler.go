package controllers

import (
	"github.com/jinzhu/gorm"
	"simple-account/api/models"
)

func CreateTransaction(tx *gorm.DB, accountId int, operationId int, amount float64) (models.Transaction, error) {

	//Validate if Operation Type exists
	//Validate if Account Type exists

	transaction := models.Transaction{
		AccountId:       accountId,
		OperationTypeId: operationId,
		Amount:          amount,
	}
	err := tx.Create(&transaction).Error

	return transaction, err
}

//func GetTransactionList(tx *gorm.DB, ids []int, documentNumbers []string) []models.Account {
//
//	var query *gorm.DB
//
//	query = tx.Where("document_number IN (?)", []string{"12345678901", "00000000001"})
//	query = query.Where("document_number <> ?", "00000000001")
//
//	var accounts []models.Account
//	query.Find(&accounts)
//
//	query = tx.Where("document_number IN (?)", []string{"12345678901", "00000000001"})
//	query.Find(&accounts)
//
//	query = tx.Where("document_number IN (?)", []string{"12345678901", "00000000001"})
//	query.Find(&accounts)
//
//	query = tx.Where("document_number <> ?", "00000000001")
//	query.Find(&accounts)
//
//	return accounts
//}

func GetTransaction(tx *gorm.DB, id int) models.Transaction {

	var transaction models.Transaction
	tx.First(&transaction, id)

	return transaction

}
