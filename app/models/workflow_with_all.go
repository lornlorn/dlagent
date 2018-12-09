package models

/*
WorkflowWithAll struct
*/
type WorkflowWithAll struct {
	Wfi           SysWorkflowInf          `json:"wfi"`
	WfdWithParams []WorkflowDtlWithParams `json:"wfd"`
}

/*
WorkflowDtlWithParams struct
workflow_dtl with parameters
*/
type WorkflowDtlWithParams struct {
	WFD SysWorkflowDtl
	WFP []SysWorkflowParam
}
