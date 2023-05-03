package app

type Config struct {
	MCSMEndpoint  string                `json:"mcsm_endpoint"`
	APIKey        string                `json:"mcsm_api_key"`
	InstanceList  []MCSMInstanceConfig  `json:"instance_list"`
	CQEndpoint    string                `json:"cq_endpoint"`
	TargetQQGroup int64                 `json:"target_qq_group"`
	AdminQQList   []adminQQConfig       `json:"admin_qq_list"`
	CommandPrefix string                `json:"prefix"`
	FunctionLevel []functionLevelConfig `json:"function_level"`
}

type functionLevelConfig struct {
	FunctionCommand string `json:"function"`
	Level           int    `json:"level"`
}

type adminQQConfig struct {
	UserID int64 `json:"id"`
	Level  int   `json:"level"`
}

type MCSMInstanceConfig struct {
	InstanceName string `json:"instance_name"`
	InstanceUUID string `json:"instance_uuid"`
	NodeUUID     string `json:"node_uuid"`
}
