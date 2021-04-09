package model

type RequestRegister struct {
	Env            string   `json:"env"`
	AppID          string   `json:"app_id"`
	Hostname       string   `json:"hostname"`
	Address        []string `json:"address"`
	Version        string   `json:"version"`
	Status         uint32   `json:"status"`
	LastTimestamp  int64    `json:"last_time"`
	DirtyTimestamp int64    `json:"dirty_time"`
	Replication    bool     `json:"replication"`
}
