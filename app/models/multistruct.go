package models

/*
JobDtl struct
*/
type JobDtl struct {
	Job          Job
	JobFlow      []Jobflow
	JobflowParam []JobflowParam
}
