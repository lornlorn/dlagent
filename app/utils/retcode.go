package utils

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
