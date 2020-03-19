package controllers

import (
	"errors"
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

	var accounts []models.Account
	query.Find(&accounts)

	return accounts
}

func GetAccount(tx *gorm.DB, id int) models.Account {

	var account models.Account
	tx.First(&account, id)

	return account
}

func UpdateBalance(tx *gorm.DB, accountId int, transactionAmount float64) error{

	account := GetAccount(tx, accountId)

	remainingBalance := account.Balance + transactionAmount

	if remainingBalance < 0 {
		return errors.New("Insufficient funds")
	}

	account.Balance = remainingBalance

	tx.Save(&account)

	return nil
}


type Gormw interface {
	First(out interface{}, where ...interface{}) Gormw
}

