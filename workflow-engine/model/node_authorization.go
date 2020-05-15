package model

// 在流程实例创建过程中，流程中每一个节点对应的参与人和其拥有的权限都记录在此表中
type NodeAuth struct {
	Model
	ProcdefID  string `json:"procdef_id"`
	ProcInstID string `json:"proc_inst_id"`
	NodeID     string `json:"node_id"`
	TaskID     string `json:"task_id"`
	Step       int    `json:"step"`

	Candidate   []string `json:"candidate"`
	MultiEdit   bool     `json:"multi_edit" gorm:"default:0"`
	RefuseAllow bool     `json:"refuse_allow"`
}
