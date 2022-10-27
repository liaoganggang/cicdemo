package tasks

import (
	"testcicd/forms"
	"testcicd/logger"
)

type deployTask struct {
}

var DeployTask = new(deployTask)

func (c *deployTask) Deploy(gitlabhook *forms.GitlabHookForm) error {
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

	//开始部署
	return nil
}
