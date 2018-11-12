package utils

import (
	"bufio"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"os"
	"strings"
)

/*
任意数据类型转JSON
*/

// Convert2JSON (data interface{}) []byte
func Convert2JSON(data interface{}) []byte {
	switch data.(type) {
	case []byte:
		// log.Println("Convert To JSON args []byte")
		retdata := data.([]byte)
		return retdata
	default:
		// log.Println("Convert To JSON args not []byte")
		retdata, err := json.Marshal(data)
		if err != nil {
			log.Printf("utils.tools.Convert2JSON -> json.Marshal Error : %v\n", err)
			return []byte("")
		}
		return retdata
	}
}

/*
读文件 并 设置偏移量和行数
*/

// ReadLines reads contents from file and splits them by new line.
// A convenience wrapper to ReadLinesOffsetN(filename, 0, -1).
func ReadLines(filename string) ([]string, error) {
	return ReadLinesOffsetN(filename, 0, -1)
}

// ReadLinesOffsetN reads contents from file and splits them by new line.
// The offset tells at which line number to start.
// The count determines the number of lines to read (starting from offset):
//   n >= 0: at most n lines
//   n < 0: whole file
func ReadLinesOffsetN(filename string, offset uint, n int) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return []string{""}, err
	}
	defer f.Close()

	var ret []string

	r := bufio.NewReader(f)
	for i := 0; i < n+int(offset) || n < 0; i++ {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}
		if i < int(offset) {
			continue
		}
		ret = append(ret, strings.Trim(line, "\n"))
	}

	return ret, nil
}

/*
生成随机UID GetUniqueID()
*/

// GetMd5String 生成32位MD5字符串
func GetMd5String(s string) string {
	newmd5 := md5.New()
	newmd5.Write([]byte(s))
	return hex.EncodeToString(newmd5.Sum(nil))
}

// GetUniqueID 生成UID唯一标识
func GetUniqueID() string {
	newbyte := make([]byte, 48)

	_, err := io.ReadFull(rand.Reader, newbyte)
	if err != nil {
		log.Printf("utils.tools.GetUniqueID -> io.ReadFull Error : %v\n", err)
		return "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	}
	return GetMd5String(base64.URLEncoding.EncodeToString(newbyte))
}
