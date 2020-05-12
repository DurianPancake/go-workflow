package model

import "time"

type FormRelation struct {
	Model
	Tenant    string     `json:"tenant"`
	FormId    string     `json:"form_id"`
	ProcdefId string     `json:"procdef_id"`
	Insert    bool       `json:"insert,omitempty" gorm:"default:'false'"`
	Edit      bool       `json:"insert,omitempty" gorm:"default:'false'"`
	Delete    bool       `json:"insert,omitempty" gorm:"default:'false'"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}
