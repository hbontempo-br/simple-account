package controllers

import (
	"github.com/jinzhu/gorm"
	"simple-account/api/models"
)

func CreateAccount(tx *gorm.DB, documentNumber string) (models.Account, error){
	account := models.Account{
		DocumentNumber: documentNumber,
	}
	err := tx.Create(&account).Error

	return account, err
}

func GetAccountList(tx *gorm.DB) []models.Account {
	var accounts []models.Account
	tx.Find(&accounts)
	return accounts
}

func GetAccount(tx *gorm.DB, id int, documentNumber string) models.Account {
	filters := make(map[string]interface{})
	if id != 0 {
		filters["id"] = id
	}
	if documentNumber != "" {
		filters["document_number"] = documentNumber
	}

	var account models.Account
	tx.First(&account, filters)

	return account
}

