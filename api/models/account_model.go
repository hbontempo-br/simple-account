package models

import "fmt"

type Account struct {
	ID             uint `gorm:"primary_key"`
	DocumentNumber string
	Balance		   float64 `gorm:"default=0"`
	//CreateAt     time.Time

}

func (Account) TableName() string {
	return "Account"
}

func (account Account) ToString() string {
	return fmt.Sprintf("Account: id: %d / document_number: %s", account.ID, account.DocumentNumber)
}
