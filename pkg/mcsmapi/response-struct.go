package mcsmapi

type InstanceStatusResponse struct {
	Status int                        `json:"status"`
	Data   InstanceStatusResponseData `json:"data"`
	Time   int64                      `json:"time"`
}

type InstanceStatusResponseTerminalOption struct {
	HaveColor bool `json:"haveColor"`
}

type InstanceStatusResponseEventTask struct {
	AutoStart   bool `json:"autoStart"`
	AutoRestart bool `json:"autoRestart"`
	Ignore      bool `json:"ignore"`
}

type InstanceStatusResponseDocker struct {
	Image       string   `json:"image"`
	Ports       []string `json:"ports"`
	Memory      int      `json:"memory"`
	NetworkMode string   `json:"networkMode"`
	CpusetCpus  string   `json:"cpusetCpus"`
	CpuUsage    int      `json:"cpuUsage"`
	MaxSpace    string   `json:"maxSpace"`
	Io          string   `json:"io"`
	Network     string   `json:"network"`
}

type InstanceStatusResponsePingConfig struct {
	Ip   string `json:"ip"`
	Port int    `json:"port"`
	Type int    `json:"type"`
}

type InstanceStatusResponseInfo struct {
	CurrentPlayers int    `json:"currentPlayers"`
	MaxPlayers     int    `json:"maxPlayers"`
	Version        string `json:"version"`
}

type InstanceStatusResponseProcessInfo struct {
	Cpu       int `json:"cpu"`
	Memory    int `json:"memory"`
	Ppid      int `json:"ppid"`
	Pid       int `json:"pid"`
	Ctime     int `json:"ctime"`
	Elapsed   int `json:"elapsed"`
	Timestamp int `json:"timestamp"`
}

type InstanceStatusResponseExtraServiceConfig struct {
	OpenFrpTunnelId string `json:"openFrpTunnelId"`
	OpenFrpToken    string `json:"openFrpToken"`
}

type InstanceStatusResponseConfig struct {
	Nickname           string                                   `json:"nickname"`
	StartCommand       string                                   `json:"startCommand"`
	StopCommand        string                                   `json:"stopCommand"`
	Cwd                string                                   `json:"cwd"`
	Ie                 string                                   `json:"ie"`
	Oe                 string                                   `json:"oe"`
	CreateDatetime     string                                   `json:"createDatetime"`
	LastDatetime       string                                   `json:"lastDatetime"`
	Type               string                                   `json:"Type"`
	Tag                []string                                 `json:"tag"`
	EndTime            string                                   `json:"endTime"`
	FileCode           string                                   `json:"fileCode"`
	ProcessType        string                                   `json:"processType"`
	TerminalOption     InstanceStatusResponseTerminalOption     `json:"terminalOption"`
	EventTask          InstanceStatusResponseEventTask          `json:"eventTask"`
	Docker             InstanceStatusResponseDocker             `json:"docker"`
	PingConfig         InstanceStatusResponsePingConfig         `json:"pingConfig"`
	ExtraServiceConfig InstanceStatusResponseExtraServiceConfig `json:"extraServiceConfig"`
}

type InstanceStatusResponseData struct {
	InstanceUuid string                            `json:"instanceUuid"`
	Started      int                               `json:"started"`
	Status       int                               `json:"status"`
	Config       InstanceStatusResponseConfig      `json:"config"`
	Info         InstanceStatusResponseInfo        `json:"info"`
	Space        int                               `json:"space"`
	ProcessInfo  InstanceStatusResponseProcessInfo `json:"processInfo"`
}

type CommonResponse struct {
	Status int   `json:"status"`
	Time   int64 `json:"time"`
}
