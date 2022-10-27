package tasks

//定义build参数

import (
	"fmt"
	"path/filepath"
	"testcicd/forms"
	"testcicd/logger"
	"testcicd/utils"
	"time"
)

type buildTask struct {
}

var BuildTask = new(buildTask)

func (c *buildTask) Build(gitlabhook *forms.GitlabHookForm) error {
	//正在build，返回错误
	if err := CheckRunTask.CheckRun(gitlabhook.ProjectID); err != nil {
		logger.Error(err)
		return err
	}

	//build结束删除projectID
	defer func() {
		CheckRunTask.mutex.Lock() //读写锁
		delete(CheckRunTask.projects, gitlabhook.ProjectID)
		CheckRunTask.mutex.Unlock()
	}()

	//源码build路径golang/golangproject/time/
	dir, err := filepath.Abs(filepath.Join(
		"build/",
		gitlabhook.Project.Namespace,
		gitlabhook.Project.Name,
		time.Now().Format("20060102_150405"),
	))

	if err != nil {
		logger.Error(err)
		return err
	}

	//创建build目录
	if err = utils.Mkdir(dir); err != nil {
		logger.Error(err)
		return err
	}

	//从gitlab仓库下载源码并构建部署包(需要提前配置好ssh秘钥)
	file, err := filepath.Abs("scripts/build.sh")
	if err != nil {
		logger.Error(err)
		return err
	}
	runfile := fmt.Sprintf("%s %s %s %s %s",
		file,
		dir,
		gitlabhook.Project.GitSSHURL,
		gitlabhook.TagName(),
		gitlabhook.Commit)
	err = Exec(runfile, dir)
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}
