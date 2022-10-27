package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testcicd/forms"
	"testcicd/logger"
	"testcicd/tasks"
	"time"
)

type HookController struct {
}

func (h *HookController) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	//判断请求头是否含有正确的token(先忽略)
	var form forms.GitlabHookForm
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.Error(err)
	}

	if err := json.Unmarshal(body, &form); err == nil {
		logger.Debug("开始")
		if form.Eventname == "tag_push" && form.Commit != "" {
			//删除也会触发tag_push事件，当删除时checkout_sha为nill，这里映射为Commit
			//开始构建
			err := tasks.BuildTask.Build(&form)
			// fmt.Fprintf(w, "ret: %s", "x")
			if err == nil {
				//当构建没问题后，开始部署
				tasks.DeployTask.Deploy(&form)
			}
		}
	} else {
		logger.Error(err)
	}

	//返回
	fmt.Fprintf(resp, "ret: %d", time.Now().Unix())
}
