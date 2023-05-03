package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

func stopInstance(cmdline []string, userID int64, appInstance *appStatus) string {
	if len(cmdline) < 3 {
		return "请提供服务器名称"
	}
	instanceIndex := getInstanceByInstanceName(appInstance.config.InstanceList, cmdline[2])
	response, err := appInstance.mcsm.StopInstance(appInstance.config.InstanceList[instanceIndex].InstanceUUID,
		appInstance.config.InstanceList[instanceIndex].NodeUUID,
	)
	if err != nil {
		return fmt.Sprintf("服务器 %s 关闭失败: %v", appInstance.config.InstanceList[instanceIndex].InstanceName, err.Error())
	}
	if response.Status != 200 {
		return fmt.Sprintf("服务器 %s 关闭失败: mcsm返回%d", appInstance.config.InstanceList[instanceIndex].InstanceName, response.Status)
	}
	return fmt.Sprintf("服务器 %s 开始关闭", appInstance.config.InstanceList[instanceIndex].InstanceName)
}

func startInstance(cmdline []string, userID int64, appInstance *appStatus) string {
	if len(cmdline) < 3 {
		return "请提供服务器名称"
	}
	instanceIndex := getInstanceByInstanceName(appInstance.config.InstanceList, cmdline[2])
	response, err := appInstance.mcsm.StartInstance(appInstance.config.InstanceList[instanceIndex].InstanceUUID,
		appInstance.config.InstanceList[instanceIndex].NodeUUID,
	)
	if err != nil {
		return fmt.Sprintf("服务器 %s 开启失败: %v", appInstance.config.InstanceList[instanceIndex].InstanceName, err.Error())
	}
	if response.Status != 200 {
		return fmt.Sprintf("服务器 %s 开启失败: mcsm返回%d", appInstance.config.InstanceList[instanceIndex].InstanceName, response.Status)
	}
	return fmt.Sprintf("服务器 %s 开始启动", appInstance.config.InstanceList[instanceIndex].InstanceName)
}

func restartInstance(cmdline []string, userID int64, appInstance *appStatus) string {
	if len(cmdline) < 3 {
		return "请提供服务器名称"
	}
	instanceIndex := getInstanceByInstanceName(appInstance.config.InstanceList, cmdline[2])
	response, err := appInstance.mcsm.RestartInstance(appInstance.config.InstanceList[instanceIndex].InstanceUUID,
		appInstance.config.InstanceList[instanceIndex].NodeUUID,
	)
	if err != nil {
		return fmt.Sprintf("服务器 %s 重启失败: %v", appInstance.config.InstanceList[instanceIndex].InstanceName, err.Error())
	}
	if response.Status != 200 {
		return fmt.Sprintf("服务器 %s 重启失败: mcsm返回%d", appInstance.config.InstanceList[instanceIndex].InstanceName, response.Status)
	}
	return fmt.Sprintf("服务器 %s 开始重启", appInstance.config.InstanceList[instanceIndex].InstanceName)
}

func killInstance(cmdline []string, userID int64, appInstance *appStatus) string {
	if len(cmdline) < 3 {
		return "请提供服务器名称"
	}
	instanceIndex := getInstanceByInstanceName(appInstance.config.InstanceList, cmdline[2])
	response, err := appInstance.mcsm.KillInstance(appInstance.config.InstanceList[instanceIndex].InstanceUUID,
		appInstance.config.InstanceList[instanceIndex].NodeUUID,
	)
	if err != nil {
		return fmt.Sprintf("服务器 %s 强行停止失败: %v", appInstance.config.InstanceList[instanceIndex].InstanceName, err.Error())
	}
	if response.Status != 200 {
		return fmt.Sprintf("服务器 %s 强行停止失败: mcsm返回%d", appInstance.config.InstanceList[instanceIndex].InstanceName, response.Status)
	}
	return fmt.Sprintf("服务器 %s 强行停止", appInstance.config.InstanceList[instanceIndex].InstanceName)
}

func listInstanceStatus(cmdline []string, userID int64, appInstance *appStatus) string {
	result := "已配置的服务器:"
	for i := 0; i < len(appInstance.config.InstanceList); i++ {
		instanceConfig := appInstance.config.InstanceList[i]
		instanceStatus, err := appInstance.mcsm.GetInstanceStatus(instanceConfig.InstanceUUID, instanceConfig.NodeUUID)
		if err != nil {
			result = result + fmt.Sprintf("\n无法获取服务器 %s 的信息: %s", instanceConfig.InstanceName, err.Error())
		} else {
			var status string
			// ↓ 会返回的值及其解释：-1（状态未知）；0（已停止）；1（正在停止）；2（正在启动）；3（正在运行）
			switch instanceStatus.Data.Status {
			case -1:
				status = "未知"
			case 0:
				status = "已停止"
			case 1:
				status = "正在停止"
				result = result + fmt.Sprintf("\n服务器 %s 状态: 正在停止 占用内存: %.2f GB",
					instanceConfig.InstanceName, float32(instanceStatus.Data.ProcessInfo.Memory)/1024/1024)
				continue
			case 2:
				status = "正在启动"
			case 3:
				status = "正在运行"
				result = result + fmt.Sprintf("\n服务器 %s 状态: 正在运行 占用内存: %.2f GB",
					instanceConfig.InstanceName, float32(instanceStatus.Data.ProcessInfo.Memory)/1024/1024)
				continue
			}
			result = result + fmt.Sprintf("\n服务器 %s 状态: %s", instanceConfig.InstanceName, status)
		}
	}
	return result
}

func listUser(cmdline []string, userID int64, appInstance *appStatus) string {
	result := "服务器控制器用户以及权限(未记录的用户权限为0)"
	for i := 0; i < len(appInstance.config.AdminQQList); i++ {
		result = fmt.Sprintf("%s\nQQ号:%d 权限:%d", result, appInstance.config.AdminQQList[i].UserID, appInstance.config.AdminQQList[i].Level)
	}
	return result
}

func setUserLevel(cmdline []string, userID int64, appInstance *appStatus) string {
	// /mcsm setuser 20 1020100000
	if len(cmdline) < 4 {
		return "缺少参数"
	}
	var level int64
	var id int64
	var err error
	level, err = strconv.ParseInt(cmdline[2], 10, 32)
	if err != nil {
		return "参数错误。level和qq号必须为数字"
	}
	id, err = strconv.ParseInt(cmdline[3], 10, 64)
	if err != nil {
		return "参数错误。level和qq号必须为数字"
	}
	isSet := false
	for i := 0; i < len(appInstance.config.AdminQQList); i++ {
		if appInstance.config.AdminQQList[i].UserID == id {
			appInstance.config.AdminQQList[i].Level = int(level)
			isSet = true
			break
		}
	}
	if !isSet {
		appInstance.config.AdminQQList = append(appInstance.config.AdminQQList, adminQQConfig{
			UserID: id,
			Level:  int(level),
		})
	}

	configFile, _ := json.Marshal(appInstance.config)
	err = ioutil.WriteFile(appInstance.configPath, configFile, 0644)
	if err != nil {
		return fmt.Sprintf("成功将qq %d 的权限设置为 %d, 保存配置文件失败:%v", id, level, err.Error())
	}
	return fmt.Sprintf("成功将qq %d 的权限设置为 %d, 并保存到配置文件", id, level)
}
