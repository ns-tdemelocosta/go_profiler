package models

//Process struct
type Process struct {
	Name      string  `json:"name"`
	CPUUsage  float64 `json:"cpu_usage"`
	Memory    float32 `json:"memory"`
	ProcessId int32   `json:"process_id"`
}
