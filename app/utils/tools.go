package utils

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"strings"
)

// Convert2JSON (data interface{}) []byte
func Convert2JSON(data interface{}) ([]byte, error) {
	switch data.(type) {
	case []byte:
		// log.Println("Convert To JSON args []byte")
		retdata := data.([]byte)
		return retdata, nil
	default:
		// log.Println("Convert To JSON args not []byte")
		retdata, err := json.Marshal(data)
		if err != nil {
			log.Printf("Marshal Json Error : %v\n", err)
		}
		return retdata, err
	}
}

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
