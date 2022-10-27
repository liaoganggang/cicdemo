package tasks

import (
	"fmt"
	"sync"
)

type checkRunTask struct {
	mutex    *sync.RWMutex
	projects map[int64]bool
}

func NewCheckRunTask() *checkRunTask {
	return &checkRunTask{
		mutex:    &sync.RWMutex{},
		projects: map[int64]bool{},
	}
}

var CheckRunTask = NewCheckRunTask()

func (c *checkRunTask) CheckRun(projectid int64) error {
	c.mutex.RLock()
	if _, ok := c.projects[projectid]; ok {
		c.mutex.RUnlock()
		return fmt.Errorf("task running,please check")
	}
	c.mutex.RUnlock()

	//true表示项目正在运行
	c.projects[projectid] = true

	return nil
}
