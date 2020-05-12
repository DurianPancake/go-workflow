package service

import (
	"github.com/jinzhu/gorm"
	"github.com/mumushuiding/util"
	"go-workflow/workflow-engine/model"
)

// SaveIdentitylinkTx SaveIdentitylinkTx
func SaveIdentitylinkTx(i *model.Identitylink, tx *gorm.DB) error {
	return i.SaveTx(tx)
}

// AddNotifierTx 添加抄送人候选用户组
func AddNotifierTx(group, tenant string, step int, procInstID string, tx *gorm.DB) error {
	yes, err := ExistsNotifierByProcInstIDAndGroup(procInstID, group)
	if err != nil {
		return err
	}
	if yes {
		return nil
	}
	i := &model.Identitylink{
		Group:      group,
		Type:       model.IdentityTypes[model.NOTIFIER],
		Step:       step,
		ProcInstID: procInstID,
		Tenant:     tenant,
	}
	return SaveIdentitylinkTx(i, tx)
}

// AddCandidateGroupTx AddCandidateGroupTx
// 添加候选用户组
func AddCandidateGroupTx(group, tenant string, step int, taskID, procInstID string, tx *gorm.DB) error {
	err := DelCandidateByProcInstID(procInstID, tx)
	if err != nil {
		return err
	}
	i := &model.Identitylink{
		Group:      group,
		Type:       model.IdentityTypes[model.CANDIDATE],
		TaskID:     taskID,
		Step:       step,
		ProcInstID: procInstID,
		Tenant:     tenant,
	}
	return SaveIdentitylinkTx(i, tx)
}

// AddCandidateUserTx AddCandidateUserTx
// 添加候选用户
func AddCandidateUserTx(userID, tenant string, step int, taskID, procInstID string, tx *gorm.DB) error {
	err := DelCandidateByProcInstID(procInstID, tx)
	if err != nil {
		return err
	}
	i := &model.Identitylink{
		UserID:     userID,
		Type:       model.IdentityTypes[model.CANDIDATE],
		TaskID:     taskID,
		Step:       step,
		ProcInstID: procInstID,
		Tenant:     tenant,
	}
	return SaveIdentitylinkTx(i, tx)
	// var wg sync.WaitGroup
	// var err1, err2 error
	// numberOfRoutine := 2
	// wg.Add(numberOfRoutine)
	// go func() {
	// 	defer wg.Done()
	// 	err1 = DelCandidateByProcInstID(procInstID, tx)
	// }()
	// go func() {
	// 	defer wg.Done()
	// 	i := &model.Identitylink{
	// 		UserID:     userID,
	// 		Type:       model.IdentityTypes[model.CANDIDATE],
	// 		TaskID:     taskID,
	// 		Step:       step,
	// 		ProcInstID: procInstID,
	// 		Tenant:    tenant,
	// 	}
	// 	err2 = SaveIdentitylinkTx(i, tx)
	// }()
	// wg.Wait()
	// fmt.Println("保存identyilink结束")
	// if err1 != nil {
	// 	return err1
	// }
	// return err2
}

//AddParticipantTx AddParticipantTx
// 添加任务参与人
func AddParticipantTx(userID, username, tenant, comment, taskID, procInstID string, step int, tx *gorm.DB) error {
	i := &model.Identitylink{
		Type:       model.IdentityTypes[model.PARTICIPANT],
		UserID:     userID,
		UserName:   username,
		ProcInstID: procInstID,
		Step:       step,
		Tenant:     tenant,
		TaskID:     taskID,
		Comment:    comment,
	}
	return SaveIdentitylinkTx(i, tx)
}

// IfParticipantByTaskID IfParticipantByTaskID
// 针对指定任务判断用户是否已经审批过了
func IfParticipantByTaskID(userID, tenant string, taskID string) (bool, error) {
	return model.IfParticipantByTaskID(userID, tenant, taskID)
}

// DelCandidateByProcInstID DelCandidateByProcInstID
// 删除历史候选人
func DelCandidateByProcInstID(procInstID string, tx *gorm.DB) error {
	return model.DelCandidateByProcInstID(procInstID, tx)
}

// ExistsNotifierByProcInstIDAndGroup 抄送人是否已经存在
func ExistsNotifierByProcInstIDAndGroup(procInstID string, group string) (bool, error) {
	return model.ExistsNotifierByProcInstIDAndGroup(procInstID, group)
}

// FindParticipantByProcInstID 查询参与审批的人
func FindParticipantByProcInstID(procInstID string) (string, error) {
	datas, err := model.FindParticipantByProcInstID(procInstID)
	if err != nil {
		return "", err
	}
	str, err := util.ToJSONStr(datas)
	if err != nil {
		return "", err
	}
	return str, nil
}
