package registry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func RegistryService(reg Registration) error {

	data, err := json.Marshal(reg)
	if err != nil {
		return err
	}

	res, err := http.Post(ServicesURL, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("注册服务{%s}失败,response code:%d \n", string(data), res.StatusCode)
	}
	return nil
}
