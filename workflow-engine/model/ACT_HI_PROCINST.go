package model

import (
	"strings"
	"sync"

	"github.com/jinzhu/gorm"
)

// ProcInst 流程实例
type ProcInst struct {
	Model
	// 流程定义ID
	ProcDefID string `json:"procDefId"`
	// 流程定义名
	ProcDefName string `json:"procDefName"`
	// title 标题
	Title string `json:"title"`
	// 用户部门 todo 获取部门直属领导或主管 、或者去除该字段
	Department string `json:"department"`
	// 公司已改为tenant
	Tenant string `json:"tenant"`
	// 当前节点
	NodeID string `json:"nodeID"`
	// 审批人
	Candidate string `json:"candidate"`
	// 当前任务
	TaskID string `json:"taskID"`
	// 实例创建时间
	StartTime string `json:"startTime"`
	// 实例结束时间
	EndTime string `json:"endTime"`
	// 实例
	Duration      int64  `json:"duration"`
	StartUserID   string `json:"startUserId"`
	StartUserName string `json:"startUserName"`
	IsFinished    bool   `gorm:"default:false" json:"isFinished"`
}

// GroupsNotNull 候选组
func GroupsNotNull(groups []string, sql string) func(db *gorm.DB) *gorm.DB {
	if len(groups) > 0 {
		return func(db *gorm.DB) *gorm.DB {
			return db.Or("candidate in (?) and "+sql, groups)
		}
	}
	return func(db *gorm.DB) *gorm.DB {
		return db
	}
}

// DepartmentsNotNull 分管部门
func DepartmentsNotNull(departments []string, sql string) func(db *gorm.DB) *gorm.DB {
	if len(departments) > 0 {
		return func(db *gorm.DB) *gorm.DB {
			return db.Or("department in (?) and candidate=? and "+sql, departments, IdentityTypes[MANAGER])
		}
	}
	return func(db *gorm.DB) *gorm.DB {
		return db
	}
}

// StartByMyself 我发起的流程
func StartByMyself(userID, tenant string, pageIndex, pageSize int) ([]*ProcInst, int, error) {
	maps := map[string]interface{}{
		"start_user_id": userID,
		"tenant":        tenant,
	}
	return findProcInsts(maps, pageIndex, pageSize)
}

// FindProcInstByID FindProcInstByID
func FindProcInstByID(id int) (*ProcInst, error) {
	var data = ProcInst{}
	err := db.Where("id=?", id).Find(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// FindProcNotify 查询抄送我的流程
func FindProcNotify(userID, tenant string, groups []string, pageIndex, pageSize int) ([]*ProcInst, int, error) {
	var datas []*ProcInst
	var count int
	var sql string
	if len(groups) != 0 {
		var s []string
		for _, val := range groups {
			s = append(s, "\""+val+"\"")
		}
		sql = "select proc_inst_id from identitylink i where i.type='notifier' and i.tenant='" + tenant + "' and (i.user_id='" + userID + "' or i.group in (" + strings.Join(s, ",") + "))"
	} else {
		sql = "select proc_inst_id from identitylink i where i.type='notifier' and i.tenant='" + tenant + "' and i.user_id='" + userID + "'"
	}
	err := db.Where("id in (" + sql + ")").Offset((pageIndex - 1) * pageSize).Limit(pageSize).Order("start_time desc").Find(&datas).Error
	if err != nil {
		return datas, count, err
	}
	err = db.Model(&ProcInst{}).Where("id in (" + sql + ")").Count(&count).Error
	if err != nil {
		return nil, count, err
	}
	return datas, count, err
}
func findProcInsts(maps map[string]interface{}, pageIndex, pageSize int) ([]*ProcInst, int, error) {
	var datas []*ProcInst
	var count int
	selectDatas := func(in chan<- error, wg *sync.WaitGroup) {
		go func() {
			err := db.Where(maps).Offset((pageIndex - 1) * pageSize).Limit(pageSize).Order("start_time desc").Find(&datas).Error
			in <- err
			wg.Done()
		}()
	}
	selectCount := func(in chan<- error, wg *sync.WaitGroup) {
		err := db.Model(&ProcInst{}).Where(maps).Count(&count).Error
		in <- err
		wg.Done()
	}
	var err1 error
	var wg sync.WaitGroup
	numberOfRoutine := 2
	wg.Add(numberOfRoutine)
	errStream := make(chan error, numberOfRoutine)
	// defer fmt.Println("close channel")
	selectDatas(errStream, &wg)
	selectCount(errStream, &wg)
	wg.Wait()
	defer close(errStream) // 关闭通道
	for i := 0; i < numberOfRoutine; i++ {
		// log.Printf("send: %v", <-errStream)
		if err := <-errStream; err != nil {
			err1 = err
		}
	}
	// fmt.Println("结束")
	return datas, count, err1
}

// FindProcInsts FindProcInsts
// 分页查询
func FindProcInsts(userID, procName, tenant string, groups, departments []string, pageIndex, pageSize int) ([]*ProcInst, int, error) {
	var datas []*ProcInst
	var count int
	var sql = " tenant='" + tenant + "' and is_finished=0 "
	if len(procName) > 0 {
		sql += "and proc_def_name='" + procName + "'"
	}
	// fmt.Println(sql)
	selectDatas := func(in chan<- error, wg *sync.WaitGroup) {
		go func() {
			err := db.Scopes(GroupsNotNull(groups, sql), DepartmentsNotNull(departments, sql)).
				Or("candidate=? and "+sql, userID).
				Offset((pageIndex - 1) * pageSize).Limit(pageSize).
				Order("start_time desc").
				Find(&datas).Error
			in <- err
			wg.Done()
		}()
	}
	selectCount := func(in chan<- error, wg *sync.WaitGroup) {
		go func() {
			err := db.Scopes(GroupsNotNull(groups, sql), DepartmentsNotNull(departments, sql)).Model(&ProcInst{}).Or("candidate=? and "+sql, userID).Count(&count).Error
			in <- err
			wg.Done()
		}()
	}
	var err1 error
	var wg sync.WaitGroup
	numberOfRoutine := 2
	wg.Add(numberOfRoutine)
	errStream := make(chan error, numberOfRoutine)
	// defer fmt.Println("close channel")
	selectDatas(errStream, &wg)
	selectCount(errStream, &wg)
	wg.Wait()
	defer close(errStream) // 关闭通道

	for i := 0; i < numberOfRoutine; i++ {
		// log.Printf("send: %v", <-errStream)
		if err := <-errStream; err != nil {
			err1 = err
		}
	}
	// fmt.Println("结束")
	return datas, count, err1
}

// Save save
func (p *ProcInst) Save() (string, error) {
	err := db.Create(p).Error
	if err != nil {
		return "", err
	}
	return p.ID, nil
}

//SaveTx SaveTx
func (p *ProcInst) SaveTx(tx *gorm.DB) (string, error) {
	if err := tx.Create(p).Error; err != nil {
		tx.Rollback()
		return "", err
	}
	return p.ID, nil
}

// DelProcInstByID DelProcInstByID
func DelProcInstByID(id string) error {
	return db.Where("id=?", id).Delete(&ProcInst{}).Error
}

// DelProcInstByIDTx DelProcInstByIDTx
// 事务
func DelProcInstByIDTx(id string, tx *gorm.DB) error {
	return tx.Where("id=?", id).Delete(&ProcInst{}).Error
}

// UpdateTx UpdateTx
func (p *ProcInst) UpdateTx(tx *gorm.DB) error {
	return tx.Model(&ProcInst{}).Updates(p).Error
}

// FindFinishedProc FindFinishedProc
func FindFinishedProc() ([]*ProcInst, error) {
	var datas []*ProcInst
	err := db.Where("is_finished=1").Find(&datas).Error
	return datas, err
}
