package controllers

import (
	"github.com/jinzhu/gorm"
	"simple-account/api/models"
)

func GetOperationTypeList(tx *gorm.DB, ids []int, enumerators []string) []models.OperationType {

	var query *gorm.DB

	query = tx.Where("1=1")
	if len(ids) > 0 {
		query = query.Where("id IN (?)", ids)
	}

	if len(enumerators) > 0 {
		query = query.Where("enumerator IN (?)", enumerators)
	}

	var operationType []models.OperationType
	query.Find(&operationType)

	return operationType
}

func GetOperationType(tx *gorm.DB, id int) models.OperationType {

	var operationType models.OperationType
	tx.First(&operationType, id)

	return operationType

}
