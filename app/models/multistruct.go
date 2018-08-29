package models

import (
	"app/db"
	"fmt"
	"log"
)

/*
JobDtl struct
*/
type JobDtl struct {
	Job     Job                `json:"job"`
	Jobflow []JobflowWithParam `json:"jobflow"`
}

/*
JobflowWithParam struct
Jobflow with parameters
*/
type JobflowWithParam struct {
	Jobflow
	JobflowParam []JobflowParam
}

/*
GetJobFlowWithParamsByJobFlowID func(jobFlowID int) ([]JobflowWithParam, error)
*/
func GetJobFlowsWithParamsByJobFlowID(jobID int) ([]JobflowWithParam, error) {
	jobFlowWithParams := make([]JobflowWithParam, 0)

	jobFlow := new(Jobflow)
	rows, err := db.Engine.Where("jf_job_id = ?", jobID).Rows(jobFlow)
	if err != nil {
		return nil, fmt.Errorf("models.multistruct.GetJobFlowsWithParamsByJobFlowID Error : %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(jobFlow); err != nil {
			return nil, fmt.Errorf("models.multistruct.GetJobFlowsWithParamsByJobFlowID rows.Scan Error : %v", err)
		}
		log.Println(jobFlow)
		jobflowWithParam := new(JobflowWithParam)
		jobflowWithParam.Jobflow = *jobFlow

		jobFlowParams, err := GetJobFlowParamsByJobFlowID(jobFlow.JfId)
		if err != nil {
			return nil, fmt.Errorf("models.multistruct.GetJobFlowsWithParamsByJobFlowID GetJobFlowParamsByJobFlowID Error : %v", err)
		}

		jobflowWithParam.JobflowParam = jobFlowParams

		jobFlowWithParams = append(jobFlowWithParams, *jobflowWithParam)
	}

	return jobFlowWithParams, nil
}

/*
GetJobDtlByID func(jobid int) (JobDtl, error)
*/
func GetJobDtlByID(jobid int) (JobDtl, error) {

	jobdtl := new(JobDtl)

	job, err := GetJobByID(jobid)
	if err != nil {
		log.Println(err)
		return JobDtl{}, err
	}

	jobFlows, err := GetJobFlowsWithParamsByJobFlowID(jobid)
	if err != nil {
		log.Println(err)
		return JobDtl{}, err
	}

	jobdtl.Job = job
	jobdtl.Jobflow = jobFlows

	return *jobdtl, nil
}
