package controller

import (
	"net/http"
)

// TODO 保存/更新流程与表单与操作之间的关系
//
func SaveProcDefManualTypeRelations(writer http.ResponseWriter, request *http.Request) {

}

// TODO 激活/上架/下架流程定义
// 如果是下架流程，则关闭其所有操作类型，并置Enable为false
// 如果是激活/上架流程
// 除置Enable为true外，首先根据流程ID查询到某表单下的操作类型
// 用需要激活的操作类型和表单ID反查可能存在冲突的已激活的流程
// 将冲突流程的操作类型关掉，若冲突流程失去所有操作类型，则将该流程作废
// 更新新流程的关系中的操作类型
// /*例如，启用某个新流程后，需要负责新增、编辑、删除审批，则所属表单以前的流程都不再生效*/
func EnableProcDefOfForm(writer http.ResponseWriter, request *http.Request) {

}

// TODO 分页查询可用流程列表，包括最高版本的未激活流程和激活流程
// 查询条件：最基本的tenant，可选：formId，创建时间，审核状态查询
func FindProcDefPageByForm(writer http.ResponseWriter, request *http.Request) {

}
