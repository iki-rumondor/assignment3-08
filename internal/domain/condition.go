package domain

type Condition struct{
	Id int `gorm:"primaryKey"`
	Wind int `gorm:"not_null"`
	Water int `gorm:"not_null"`
}