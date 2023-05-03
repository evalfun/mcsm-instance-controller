package app

import "fmt"

func getInstanceByInstanceName(instanceConfigList []MCSMInstanceConfig, instanceName string) int {
	for i := 0; i < len(instanceConfigList); i++ {
		if instanceConfigList[i].InstanceName == instanceName {
			return i
		}
	}
	return -1
}

func getUserLevel(userList []adminQQConfig, userID int64) int {
	for i := 0; i < len(userList); i++ {
		if userList[i].UserID == userID {
			return userList[i].Level
		}
	}
	return 0
}

type appFunction func(cmdline []string, userID int64, appInstance *appStatus) string

type appFunctionRecord struct {
	command     string
	level       int
	function    appFunction
	description string
}

func registerFunction(appInstance *appStatus) string {
	var helpMessage string = "服务器控制器指令:"

	appInstance.functionList = append(appInstance.functionList, appFunctionRecord{
		command:     "list",
		level:       getFunctionLevel(appInstance, "list"),
		function:    listInstanceStatus,
		description: "list 列出服务器和状态",
	})

	appInstance.functionList = append(appInstance.functionList, appFunctionRecord{
		command:     "start",
		level:       getFunctionLevel(appInstance, "start"),
		function:    startInstance,
		description: "start <服务器名称> 启动服务器",
	})

	appInstance.functionList = append(appInstance.functionList, appFunctionRecord{
		command:     "stop",
		level:       getFunctionLevel(appInstance, "stop"),
		function:    stopInstance,
		description: "stop <服务器名称> 停止服务器",
	})

	appInstance.functionList = append(appInstance.functionList, appFunctionRecord{
		command:     "kill",
		level:       getFunctionLevel(appInstance, "kill"),
		function:    killInstance,
		description: "kill <服务器名称> 强行停止服务器",
	})

	appInstance.functionList = append(appInstance.functionList, appFunctionRecord{
		command:     "restart",
		level:       getFunctionLevel(appInstance, "restart"),
		function:    restartInstance,
		description: "restart <服务器名称> 重启服务器",
	})

	appInstance.functionList = append(appInstance.functionList, appFunctionRecord{
		command:     "setuser",
		level:       getFunctionLevel(appInstance, "setuser"),
		function:    setUserLevel,
		description: "setuser <权限(数字)> <qq号(数字)> 设置用户的权限等级",
	})

	appInstance.functionList = append(appInstance.functionList, appFunctionRecord{
		command:     "listuser",
		level:       getFunctionLevel(appInstance, "listuser"),
		function:    listUser,
		description: "listuser 列出用户的权限等级",
	})

	for i := 0; i < len(appInstance.functionList); i++ {
		helpMessage = fmt.Sprintf("%s\n%s %s 权限:%d",
			helpMessage, appInstance.config.CommandPrefix,
			appInstance.functionList[i].description, appInstance.functionList[i].level)
	}
	return helpMessage
}

func getFunctionLevel(appInstance *appStatus, functionCommand string) int {
	for i := 0; i < len(appInstance.config.FunctionLevel); i++ {
		if appInstance.config.FunctionLevel[i].FunctionCommand == functionCommand {
			return appInstance.config.FunctionLevel[i].Level
		}
	}
	return 999
}

func getFunction(appInstance *appStatus, command string) int {
	for i := 0; i < len(appInstance.functionList); i++ {
		if appInstance.functionList[i].command == command {
			return i
		}
	}
	return -1
}
