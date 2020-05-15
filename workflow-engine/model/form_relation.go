package model

import "time"

// Form与工作流的关系表
// Form与工作流为一对多的关系，这里的ProcdefId与流程定义表里一一对应
// 一个流程定义可以控制新增、编辑、删除功能，当bool值为true时，表示当前表单的某些操作被某流程激活
// 联立流程定义表的Enable字段，Enable下至多有一条流程的同一类型的开关为true
type FormRelation struct {
	Model
	Tenant    string     `json:"tenant"`
	AppId     string     `json:"app_id"`
	FormId    string     `json:"form_id"`
	ProcdefID string     `json:"procdef_id"`
	Insert    bool       `json:"insert,omitempty" gorm:"default:0"`
	Edit      bool       `json:"edit,omitempty" gorm:"default:0"`
	Delete    bool       `json:"delete,omitempty" gorm:"default:0"`
	Version   int        `json:"version"`
	CreatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}
