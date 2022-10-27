package tasks

//执行build脚本，将gitlab代码clone到本地
//将build脚本执行过程输出到txt文件

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"testcicd/logger"
	"time"
)

type ConfigFile struct {
	Stdout    string
	Stderr    string
	ResCode   int64
	StartTime string
	EndTime   *time.Time
	RunTime   string
}

var (
	stdoutstring string
	stderrstring string
	buffer       bytes.Buffer
	configfile   *ConfigFile
)

func Exec(runfile, resultPath string) error {
	startTime := time.Now().Format("2006-00-02 15:04:05")
	fmt.Println(runfile)
	cmd := exec.Command("/bin/bash", "-c", runfile)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		logger.Error(err)
		return err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		logger.Error(err)
		return err
	}
	cmd.Start()

	//正确和错误输出
	stdoutReader := bufio.NewReader(stdout)
	stderrReader := bufio.NewReader(stderr)

	//定义结果写入的文件
	resPath := filepath.Join(resultPath, "build.txt")
	file, _ := os.OpenFile(resPath, os.O_CREATE|os.O_RDWR, 0644)
	defer file.Close()

	for {
		stdoutctx, _, err := stdoutReader.ReadLine()
		if err == io.EOF {
			break
		}
		stdoutstring = stdoutstring + string(stdoutctx) + "\n"

		stderrctx, _, _ := stderrReader.ReadLine()

		stderrstring = stderrstring + string(stderrctx) + "\n"

		configfile = &ConfigFile{
			StartTime: startTime,
			EndTime:   nil,
			RunTime:   "",
			Stdout:    stdoutstring,
			Stderr:    stderrstring,
			ResCode:   0,
		}
		st, _ := json.Marshal(configfile)
		logger.Debug(string(st))
	}
	res, err := json.Marshal(configfile)
	if err != nil {
		logger.Error(err)
		return err
	}
	json.Indent(&buffer, res, "", "\t")
	buffer.WriteTo(file) //将结果写入file

	return nil
}
