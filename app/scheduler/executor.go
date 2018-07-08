package scheduler

import (
	"os/exec"
	"strings"
)

func RunCmd(cronsh string, croncmd string) ([]byte, error) {
	var command string
	var args []string

	if cronsh == "" {
		cmdargs := strings.Split(croncmd, " ")
		command = cmdargs[0]
		args = cmdargs[1:]
	} else {
		command = cronsh
		// cmdargs := "-c " + croncmd
		// 转换为可变长数组
		// args = strings.Split(cmdargs, " ")
		args = []string{"-c", croncmd}
	}
	cmd := exec.Command(command, args...)
	/*
	   测试start
	*/
	// t, _ := exec.LookPath(command)
	// log.Println(t)
	// log.Println(filepath.Base(command))
	/*
	   测试end
	*/

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	if err = cmd.Start(); err != nil {
		return nil, err
	}

	var out = make([]byte, 0, 1024)
	for {
		tmp := make([]byte, 128)
		n, err := stdout.Read(tmp)
		out = append(out, tmp[:n]...)
		if err != nil {
			break
		}
	}

	if err = cmd.Wait(); err != nil {
		return nil, err
	}

	return out, nil
}
