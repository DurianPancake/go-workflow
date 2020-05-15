package model

/**
 * 操作记录日志记录表
 */
type EventLog struct {
	//
	ProcInstID     string `json:"proc_inst_id"`
	ProcDefID      string `json:"procDefId"`
	ProcDefName    string `json:"procDefName"`
	Title          string `json:"title"`
	FormID         string `json:"form_id"`
	FormDataID     string `json:"form_data_id"`
	TaskID         string `json:"taskID"`
	NodeID         string `json:"nodeID"`
	Step           int    `json:"step"`            // 记录执行的某一步
	ExecutionOrder int    `json:"execution_order"` // 实际执行的次序（在相同的一步中，多人审批共用同一个Order)
	Candidate      string `json:"candidate"`
	ActType        string `json:"act_type"`      // 表示是会签还是或签
	AuditOutcome   string `json:"audit_outcome"` // 审批结果
	AuditOpinion   string `json:"audit_opinion"` // 审批意见
}
