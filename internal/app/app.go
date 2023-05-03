package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/evalfun/mcsm-instance-controller/pkg/cqapi"
	"github.com/evalfun/mcsm-instance-controller/pkg/mcsmapi"
)

type appStatus struct {
	config       *Config
	cq           *cqapi.CQAPIClient
	mcsm         *mcsmapi.MCSMServer
	functionList []appFunctionRecord
	configPath   string
}

func StartMCSMController(configPath string) {
	configFileData, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Panicf("Can not open config file: %v", configPath)
	}
	config := &Config{}
	err = json.Unmarshal(configFileData, config)
	if err != nil {
		log.Panicf("Unmarshal config json failed: %v", err.Error())
	}
	//建立cqhttp连接
	appInstance := &appStatus{}
	appInstance.configPath = configPath
	appInstance.cq = cqapi.CreateCQAPIClient(config.CQEndpoint)
	appInstance.config = config
	appInstance.mcsm = &mcsmapi.MCSMServer{
		ServerEndpoint: config.MCSMEndpoint,
		APIKey:         config.APIKey,
	}
	helpMessage := registerFunction(appInstance)
	for {
		report := <-appInstance.cq.RecvChan
		messageReport, ok := report.(*cqapi.MessageReport)
		if ok {
			if messageReport.GroupID == config.TargetQQGroup {
				if messageReport.Message == "/help" {
					appInstance.cq.SendGroupMessage(config.TargetQQGroup, helpMessage)
				} else {
					if strings.Index(messageReport.Message, config.CommandPrefix+" ") == 0 {
						cmdline := strings.Split(messageReport.Message, " ")
						if len(cmdline) == 1 {
							appInstance.cq.SendGroupMessage(config.TargetQQGroup, helpMessage)
							continue
						}
						functionIndex := getFunction(appInstance, cmdline[1])
						if functionIndex == -1 {
							appInstance.cq.SendGroupMessage(config.TargetQQGroup, helpMessage)
						}
						userLevel := getUserLevel(config.AdminQQList, messageReport.UserID)
						if userLevel < appInstance.functionList[functionIndex].level {
							appInstance.cq.SendGroupMessage(config.TargetQQGroup, fmt.Sprintf("[CQ:at,qq=%d]执行命令 %s 至少需要 %d 权限, 你的权限为 %d", messageReport.UserID, cmdline[1], appInstance.functionList[functionIndex].level, userLevel))
							log.Printf("用户 %d 执行命令失败(权限不足): %v", messageReport.UserID, cmdline)
						} else {
							log.Printf("用户 %d 执行命令成功: %v", messageReport.UserID, cmdline)
							result := appInstance.functionList[functionIndex].function(cmdline, messageReport.UserID, appInstance)
							if result != "" {
								appInstance.cq.SendGroupMessage(config.TargetQQGroup, result)
							} else {
								appInstance.cq.SendGroupMessage(config.TargetQQGroup, fmt.Sprintf("[CQ:at,qq=%d]命令 %s 已经成功执行", messageReport.UserID, cmdline[1]))
							}
						}
					}
					continue
				}
			}
		}
	}
}
