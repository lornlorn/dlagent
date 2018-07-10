package utils

import (
	"app/models"
	"fmt"
)

// retCodeDict return code directory
var retCodeDict = map[string]string{
	"0000": "成功",
	"9999": "未知错误",
}

// GetRetMsg func(retcode string) string
// Get Return Message By Return Code
func GetRetMsg(retcode string) string {
	return retCodeDict[retcode]
}

/*
GetAjaxRetObj func(retcode string, errmsg error) []byte
Get return code and errors
Convert to models.AjaxReturn struct object
*/
func GetAjaxRetObj(retcode string, errmsg error) models.AjaxReturn {
	var retobj models.AjaxReturn

	if errmsg != nil {
		retobj = models.AjaxReturn{
			RetCode: retcode,
			RetMsg:  fmt.Sprintf("%v, %v", GetRetMsg(retcode), errmsg),
		}
	} else {
		retobj = models.AjaxReturn{
			RetCode: retcode,
			RetMsg:  GetRetMsg(retcode),
		}
	}
	return retobj
}

/*
GetAjaxRetJSON func(retcode string, errmsg error) []byte
Get return code and errors
Convert to models.AjaxReturn struct and Marshal to JSON
*/
func GetAjaxRetJSON(retcode string, errmsg error) []byte {
	var retobj models.AjaxReturn

	if errmsg != nil {
		retobj = models.AjaxReturn{
			RetCode: retcode,
			RetMsg:  fmt.Sprintf("%v, %v", GetRetMsg(retcode), errmsg),
		}
	} else {
		retobj = models.AjaxReturn{
			RetCode: retcode,
			RetMsg:  GetRetMsg(retcode),
		}
	}

	ret, _ := Convert2JSON(retobj)

	return ret
}
