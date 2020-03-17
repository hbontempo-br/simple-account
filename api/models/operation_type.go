package models

type OperationType struct {
	ID          uint   `gorm:"primary_key"`
	Enumerator  string `gorm:"unique;not null"`
	Description string `gorm:"unique;not null"`
	Behavior    string `gorm:"not null"`
	//CreateAt     *time.Time
}

func (OperationType) TableName() string {
	return "OperationType"
}
