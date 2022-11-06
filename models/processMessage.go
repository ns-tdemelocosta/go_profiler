package models

type ProcessMessage struct {
	Pid       int32   `json:"pid"`
	Cpu       float64 `json:"cpu"`
	Mem       float32 `json:"mem"`
	Name      string  `json:"name"`
	TimeStamp int64   `json:"time"`
	Ctime     int64   `json:"ctime"`
}
