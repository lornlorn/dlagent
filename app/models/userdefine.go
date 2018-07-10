package models

// AjaxReturn struct
type AjaxReturn struct {
	RetCode string
	RetMsg  string
	RetData []byte
}

// ReflectReturn struct
type ReflectReturn struct {
	AjaxReturn
}
