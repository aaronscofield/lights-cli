package govee

type stateProperties map[string]interface{}

type stateData struct {
	Device     string            `json:"device"`
	Model      string            `json:"model"`
	Properties []stateProperties `json:"properties"`
}

type stateResponse struct {
	Data    stateData `json:"data"`
	Message string    `json:"message"`
	Code    int       `json:"code"`
}

type turnResponse struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data,omitempty"`
}

type cmd struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type body struct {
	Device string `json:"device"`
	Model  string `json:"model"`
	Cmd    cmd    `json:"cmd"`
}

type Device struct {
	Device       string   `json:"device"`
	Model        string   `json:"model"`
	DeviceName   string   `json:"deviceName"`
	Controllable bool     `json:"controllable"`
	Retrievable  bool     `json:"retrievable"`
	SupportCmds  []string `json:"supportCmds"`
}

type devicesData struct {
	Devices []Device `json:"devices"`
}

type devicesResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    devicesData `json:"data,omitempty"`
}
