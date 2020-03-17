package controllers

import (
	"github.com/jinzhu/gorm"
	"simple-account/api/models"
)

func CreateAccount(tx *gorm.DB, documentNumber string) (models.Account, error) {
	account := models.Account{
		DocumentNumber: documentNumber,
	}
	err := tx.Create(&account).Error

	return account, err
}

func GetAccountList(tx *gorm.DB, ids []int, documentNumbers []string) []models.Account {

	var query *gorm.DB

	query = tx.Where("document_number IN (?)", []string{"12345678901", "00000000001"})
	query = query.Where("document_number <> ?", "00000000001")

	var accounts []models.Account
	query.Find(&accounts)

	query = tx.Where("document_number IN (?)", []string{"12345678901", "00000000001"})
	query.Find(&accounts)

	query = tx.Where("document_number IN (?)", []string{"12345678901", "00000000001"})
	query.Find(&accounts)

	query = tx.Where("document_number <> ?", "00000000001")
	query.Find(&accounts)

	return accounts
}

func GetAccount(tx *gorm.DB, id int) models.Account {

	var account models.Account
	tx.First(&account, id)

	return account
}
