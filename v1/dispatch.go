/**
    @author: dongjs
    @date: 2025/2/8
    @description: 调度
**/

package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// 参数设置校验
// 参数设置之前校验当前设备是否可以进行参数设置。
func (c *Client) ParamSettingCheck(request ParamSettingCheckRequest) (*ParamSettingCheckResponse, error) {
	if request.Appkey == "" {
		request.Appkey = c.appkey
	}
	requestBytes, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}
	fmt.Println("requestBytes:", string(requestBytes))
	buf := bytes.NewBuffer(requestBytes)
	body, err := c.doRequest("openapi/paramSettingCheck", buf)
	if err != nil {
		return nil, err
	}
	fmt.Printf("body:%+v", string(body))
	response := ParamSettingCheckResponse{}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 参数设置任务下发
// 通过接口下发对应的设备的参数设置信息。
func (c *Client) ParamSetting(request ParamSettingRequest) (*ParamSettingResponse, error) {
	if request.Appkey == "" {
		request.Appkey = c.appkey
	}
	requestBytes, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}
	fmt.Println("requestBytes:", string(requestBytes))
	buf := bytes.NewBuffer(requestBytes)
	body, err := c.doRequest("openapi/paramSetting", buf)
	if err != nil {
		return nil, err
	}
	fmt.Printf("body:%+v", string(body))
	response := ParamSettingResponse{}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 参数设置结果查询
// 通过参数设置任务id查询参数设置结果信息。
func (c *Client) GetParamSettingTask(request GetParamSettingTaskRequest) (*GetParamSettingTaskResponse, error) {
	if request.Appkey == "" {
		request.Appkey = c.appkey
	}
	requestBytes, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}
	fmt.Println("requestBytes:", string(requestBytes))
	buf := bytes.NewBuffer(requestBytes)
	body, err := c.doRequest("openapi/getParamSettingTask", buf)
	if err != nil {
		return nil, err
	}
	fmt.Printf("body:%+v", string(body))
	response := GetParamSettingTaskResponse{}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 下发只读参数读取任务
func (c *Client) ReadOnlyParamSet(request ReadOnlyParamSetRequest) (*ReadOnlyParamSetResponse, error) {
	if request.Appkey == "" {
		request.Appkey = c.appkey
	}
	requestBytes, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}
	fmt.Println("requestBytes:", string(requestBytes))
	buf := bytes.NewBuffer(requestBytes)
	body, err := c.doRequest("openapi/readOnlyParamSet", buf)
	if err != nil {
		return nil, err
	}
	fmt.Printf("body:%+v", string(body))
	response := ReadOnlyParamSetResponse{}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 只读参数读取结果查询
func (c *Client) GetReadOnlyResult(request GetReadOnlyResultRequest) (*GetReadOnlyResultResponse, error) {
	if request.Appkey == "" {
		request.Appkey = c.appkey
	}
	requestBytes, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}
	fmt.Println("requestBytes:", string(requestBytes))
	buf := bytes.NewBuffer(requestBytes)
	body, err := c.doRequest("openapi/getReadOnlyResult", buf)
	if err != nil {
		return nil, err
	}
	fmt.Printf("body:%+v", string(body))
	response := GetReadOnlyResultResponse{}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 开启实时数据上传功能
// 给逆变器下发指令，开启实时数据上传功能 ,可通过mqtt订阅实时数据。
func (c *Client) GetReadOnlyParamDefinition(request GetReadOnlyParamDefinitionRequest) (*GetReadOnlyParamDefinitionResponse, error) {
	if request.Appkey == "" {
		request.Appkey = c.appkey
	}
	requestBytes, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}
	fmt.Println("requestBytes:", string(requestBytes))
	buf := bytes.NewBuffer(requestBytes)
	body, err := c.doRequest("openapi/getReadOnlyParamDefinition", buf)
	if err != nil {
		return nil, err
	}
	fmt.Printf("body:%+v", string(body))
	response := GetReadOnlyParamDefinitionResponse{}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
