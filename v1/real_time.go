/**
    @author: dongjs
    @date: 2025/2/8
    @description: 实时数据
**/

package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// 根据appkey获取对应的Mqtt信息
// 该接口根据入参中的appkey获取该appkey对应的Mqtt公钥、Mqtt地址、用户名和密码等信息。
func (c *Client) DataSubscribeGetConfig(request DataSubscribeGetConfigRequest) (*DataSubscribeGetConfigResponse, error) {
	if request.Appkey == "" {
		request.Appkey = c.appkey
	}
	requestBytes, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}
	fmt.Println("requestBytes:", string(requestBytes))
	buf := bytes.NewBuffer(requestBytes)
	body, err := c.doRequest("openapi/datasubscribe/getConfig", buf)
	if err != nil {
		return nil, err
	}
	fmt.Printf("body:%+v", string(body))
	response := DataSubscribeGetConfigResponse{}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 开启实时数据上传功能
func (c *Client) DataSubscribeStart(request DataSubscribeStartRequest) (*DataSubscribeStartResponse, error) {
	if request.Appkey == "" {
		request.Appkey = c.appkey
	}
	requestBytes, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}
	fmt.Println("requestBytes:", string(requestBytes))
	buf := bytes.NewBuffer(requestBytes)
	body, err := c.doRequest("openapi/datasubscribe/start", buf)
	if err != nil {
		return nil, err
	}
	fmt.Printf("body:%+v", string(body))
	response := DataSubscribeStartResponse{}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 关闭实时数据上传功能
// 给逆变器下发指令，关闭实时数据上传功能
func (c *Client) DataSubscribeStop(request DataSubscribeStopRequest) (*DataSubscribeStopResponse, error) {
	if request.Appkey == "" {
		request.Appkey = c.appkey
	}
	requestBytes, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}
	fmt.Println("requestBytes:", string(requestBytes))
	buf := bytes.NewBuffer(requestBytes)
	body, err := c.doRequest("openapi/datasubscribe/stop", buf)
	if err != nil {
		return nil, err
	}
	fmt.Printf("body:%+v", string(body))
	response := DataSubscribeStopResponse{}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 获取历史数据
// 获取设备历史数据，目前我们仅提供最近一小时的数据。只有启用了实时数据上传功能的设备支持历史数据查询。
func (c *Client) DataSubscribeGetHisData(request DataSubscribeGetHisDataRequest) (*DataSubscribeGetHisDataResponse, error) {
	if request.Appkey == "" {
		request.Appkey = c.appkey
	}
	requestBytes, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}
	fmt.Println("requestBytes:", string(requestBytes))
	buf := bytes.NewBuffer(requestBytes)
	body, err := c.doRequest("openapi/datasubscribe/getHisData", buf)
	if err != nil {
		return nil, err
	}
	fmt.Printf("body:%+v", string(body))
	response := DataSubscribeGetHisDataResponse{}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
