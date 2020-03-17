package models

import (
	"fmt"
	"time"
)

type Transaction struct {
	ID        uint `gorm:"primary_key"`
	AccountId int
	//Account	   Account
	OperationTypeId int
	//OperationType	   OperationType
	Amount    float64
	EventDate *time.Time `sql:"DEFAULT:current_timestamp"`
	//CreateAt     time.Time

}

func (Transaction) TableName() string {
	return "Transaction"
}

func (transaction Transaction) ToString() string {
	return fmt.Sprintf("Transaction: id: %d / account_id: %s / operation_type_id: %s / amount: %s / event_date: %s", transaction.ID, transaction.AccountId, transaction.OperationTypeId, transaction.Amount, transaction.EventDate)
}
