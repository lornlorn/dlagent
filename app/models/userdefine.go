package models

// AjaxReturn struct
type AjaxReturn struct {
	RetCode string
	RetMsg  string
}

// ReflectReturn struct
type ReflectReturn struct {
	AjaxReturn
	RetData []byte
}
