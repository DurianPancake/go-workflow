package model

// ProcdefHistory 历史流程定义
type ProcdefHistory struct {
	Procdef
}

// Save Save
func (p *ProcdefHistory) Save() (ID string, err error) {
	err = db.Create(p).Error
	if err != nil {
		return "", err
	}
	return p.ID, nil
}
