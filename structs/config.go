package structs

import "time"

type Increment struct {
	ID uint `gorm:"primary_key" json:"id"`
}

type Timestamps struct {
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

type TableConfig interface {
	TableName() string
}

func (Person) TableName() string {
	return "persons"
}
