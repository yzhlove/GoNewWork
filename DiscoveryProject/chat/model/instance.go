package model

import "time"

type Instance struct {
	Env      string   `json:"env"`
	AppId    string   `json:"app_id"`
	Hostname string   `json:"hostname"`
	Address  []string `json:"address"`
	Version  string   `json:"version"`
	Status   uint32   `json:"status"`

	RegTimestamp   int64 `json:"register_time"`
	UpTimestamp    int64 `json:"up_time"`
	RenewTimestamp int64 `json:"renew_time"`
	DirtyTimestamp int64 `json:"dirty_time"`
	LastTimestamp  int64 `json:"last_time"`
}

func NewInstance(req *RequestRegister) *Instance {
	now := time.Now().UnixNano()
	return &Instance{
		Env:            req.Env,
		AppId:          req.AppID,
		Hostname:       req.Hostname,
		Address:        req.Address,
		Version:        req.Version,
		Status:         req.Status,
		RegTimestamp:   now,
		UpTimestamp:    now,
		RenewTimestamp: now,
		DirtyTimestamp: now,
		LastTimestamp:  now,
	}
}
