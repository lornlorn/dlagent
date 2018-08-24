package models

import "log"

/*
JobDtl struct
*/
type JobDtl struct {
	Job          Job            `json:"job"`
	Jobflow      []Jobflow      `json:"jobflow"`
	JobflowParam []JobflowParam `json:"jobflowparam"`
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
GetJobDtlByID func(jobid int) (JobDtl, error)
*/
func GetJobDtlByID(jobid int) (JobDtl, error) {

	jobdtl := new(JobDtl)

	job, err := GetJobByID(jobid)
	if err != nil {
		log.Println(err)
		return JobDtl{}, err
	}

	jobflows, err := GetJobFlowsByJobID(jobid)
	if err != nil {
		log.Println(err)
		return JobDtl{}, err
	}

	jobdtl.Job = job
	jobdtl.Jobflow = jobflows

	return *jobdtl, nil
}
