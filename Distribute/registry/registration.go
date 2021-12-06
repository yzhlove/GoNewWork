package registry

type Registration struct {
	ServiceName ServiceName `json:"name"` // 服务名称
	ServiceURL  string      `json:"url"`  // 服务地址
}

type ServiceName string

const (
	LogService = ServiceName("LogService")
)
