package mcsmapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

// 获取实例状态
func (server *MCSMServer) GetInstanceStatus(instance_uuid string, node_uuid string) (*InstanceStatusResponse, error) {
	apiUrl := server.ServerEndpoint + "/api/instance"
	queryParams := url.Values{}
	queryParams.Set("apikey", server.APIKey)
	queryParams.Set("uuid", instance_uuid)
	queryParams.Set("remote_uuid", node_uuid)
	u, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = queryParams.Encode()
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	instanceStatus := &InstanceStatusResponse{}
	err = json.Unmarshal(data, instanceStatus)
	return instanceStatus, err
}

// 停止实例
func (server *MCSMServer) StopInstance(instance_uuid string, node_uuid string) (*CommonResponse, error) {
	apiUrl := server.ServerEndpoint + "/api/protected_instance/stop"
	queryParams := url.Values{}
	queryParams.Set("apikey", server.APIKey)
	queryParams.Set("uuid", instance_uuid)
	queryParams.Set("remote_uuid", node_uuid)
	u, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = queryParams.Encode()
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	response := &CommonResponse{}
	err = json.Unmarshal(data, response)
	return response, err
}

// 强制停止实例
func (server *MCSMServer) KillInstance(instance_uuid string, node_uuid string) (*CommonResponse, error) {
	apiUrl := server.ServerEndpoint + "/api/protected_instance/kill"
	queryParams := url.Values{}
	queryParams.Set("apikey", server.APIKey)
	queryParams.Set("uuid", instance_uuid)
	queryParams.Set("remote_uuid", node_uuid)
	u, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = queryParams.Encode()
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	response := &CommonResponse{}
	err = json.Unmarshal(data, response)
	return response, err
}

// 开启实例
func (server *MCSMServer) StartInstance(instance_uuid string, node_uuid string) (*CommonResponse, error) {
	apiUrl := server.ServerEndpoint + "/api/protected_instance/open"
	queryParams := url.Values{}
	queryParams.Set("apikey", server.APIKey)
	queryParams.Set("uuid", instance_uuid)
	queryParams.Set("remote_uuid", node_uuid)
	u, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = queryParams.Encode()
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	response := &CommonResponse{}
	err = json.Unmarshal(data, response)
	return response, err
}

// 重启实例
func (server *MCSMServer) RestartInstance(instance_uuid string, node_uuid string) (*CommonResponse, error) {
	apiUrl := server.ServerEndpoint + "/api/protected_instance/restart"
	queryParams := url.Values{}
	queryParams.Set("apikey", server.APIKey)
	queryParams.Set("uuid", instance_uuid)
	queryParams.Set("remote_uuid", node_uuid)
	u, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = queryParams.Encode()
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	response := &CommonResponse{}
	err = json.Unmarshal(data, response)
	return response, err
}
