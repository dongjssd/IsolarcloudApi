/**
    @author: dongjs
    @date: 2025/2/7
    @description: 监控接口
**/

package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// 电站列表信息查询
func (c *Client) PowerStationList(request PowerStationListRequest) (*PowerStationListResponse, error) {
	if request.Appkey == "" {
		request.Appkey = c.appkey
	}
	requestBytes, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}
	fmt.Println("requestBytes:", string(requestBytes))
	buf := bytes.NewBuffer(requestBytes)
	body, err := c.doRequest("openapi/getPowerStationList", buf)
	if err != nil {
		return nil, err
	}
	fmt.Printf("body:%+v", string(body))
	response := PowerStationListResponse{}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 查询电站下设备列表
func (c *Client) GetDeviceList(request GetDeviceListRequest) (*GetDeviceListResponse, error) {
	if request.Appkey == "" {
		request.Appkey = c.appkey
	}
	requestBytes, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}
	fmt.Println("requestBytes:", string(requestBytes))
	buf := bytes.NewBuffer(requestBytes)
	body, err := c.doRequest("openapi/getDeviceList", buf)
	if err != nil {
		return nil, err
	}
	fmt.Printf("body:%+v", string(body))
	response := GetDeviceListResponse{}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 查询电站的基本信息
func (c *Client) GetPowerStationDetail(request GetPowerStationDetailRequest) (*GetPowerStationDetailResponse, error) {
	if request.Appkey == "" {
		request.Appkey = c.appkey
	}
	requestBytes, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}
	fmt.Println("requestBytes:", string(requestBytes))
	buf := bytes.NewBuffer(requestBytes)
	body, err := c.doRequest("openapi/getPowerStationDetail", buf)
	if err != nil {
		return nil, err
	}
	fmt.Printf("body:%+v", string(body))
	response := GetPowerStationDetailResponse{}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 查询设备的实时测点数据
func (c *Client) getDeviceRealTimeData(request GetDeviceRealTimeDataRequest) (*GetDeviceRealTimeDataResponse, error) {
	if request.Appkey == "" {
		request.Appkey = c.appkey
	}
	requestBytes, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}
	fmt.Println("requestBytes:", string(requestBytes))
	buf := bytes.NewBuffer(requestBytes)
	body, err := c.doRequest("openapi/getDeviceRealTimeData", buf)
	if err != nil {
		return nil, err
	}
	fmt.Printf("body:%+v", string(body))
	response := GetDeviceRealTimeDataResponse{}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 查询同类型设备的历史测点分钟级数据
func (c *Client) GetDevicePointMinuteDataList(request GetDevicePointMinuteDataListRequest) (*GetDevicePointMinuteDataListResponse, error) {
	if request.Appkey == "" {
		request.Appkey = c.appkey
	}
	requestBytes, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}
	fmt.Println("requestBytes:", string(requestBytes))
	buf := bytes.NewBuffer(requestBytes)
	body, err := c.doRequest("openapi/getDevicePointMinuteDataList", buf)
	if err != nil {
		return nil, err
	}
	fmt.Printf("body:%+v", string(body))
	response := GetDevicePointMinuteDataListResponse{}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 查询同类型设备的历史测点日月年数据
func (c *Client) GetDevicePointsDayMonthYearDataList(request GetDevicePointsDayMonthYearDataListRequest) (*GetDevicePointsDayMonthYearDataListResponse, error) {
	if request.Appkey == "" {
		request.Appkey = c.appkey
	}
	requestBytes, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}
	fmt.Println("requestBytes:", string(requestBytes))
	buf := bytes.NewBuffer(requestBytes)
	body, err := c.doRequest("openapi/getDevicePointsDayMonthYearDataList", buf)
	if err != nil {
		return nil, err
	}
	fmt.Printf("body:%+v", string(body))
	response := GetDevicePointsDayMonthYearDataListResponse{}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 查询设备的属性测点数据
func (c *Client) GetDevPropertyPointValue(request GetDevPropertyPointValueRequest) (*GetDevPropertyPointValueResponse, error) {
	if request.Appkey == "" {
		request.Appkey = c.appkey
	}
	requestBytes, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}
	fmt.Println("requestBytes:", string(requestBytes))
	buf := bytes.NewBuffer(requestBytes)
	body, err := c.doRequest("openapi/getDevPropertyPointValue", buf)
	if err != nil {
		return nil, err
	}
	fmt.Printf("body:%+v", string(body))
	response := GetDevPropertyPointValueResponse{}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 查询设备的故障告警信息。
func (c *Client) GetFaultAlarmInfo(request GetFaultAlarmInfoRequest) (*GetFaultAlarmInfoResponse, error) {
	if request.Appkey == "" {
		request.Appkey = c.appkey
	}
	requestBytes, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}
	fmt.Println("requestBytes:", string(requestBytes))
	buf := bytes.NewBuffer(requestBytes)
	body, err := c.doRequest("openapi/getFaultAlarmInfo", buf)
	if err != nil {
		return nil, err
	}
	fmt.Printf("body:%+v", string(body))
	response := GetFaultAlarmInfoResponse{}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 查询阳光云开放的测点定义信息
// 根据设备类型查询该类型设备开放的遥测、遥信、属性测点信息,同时支持根据设备类型和设备型号ID查询该型号设备支持的开放遥测测点信息。
func (c *Client) GetOpenPointInfo(request GetOpenPointInfoRequest) (*GetOpenPointInfoResponse, error) {
	if request.Appkey == "" {
		request.Appkey = c.appkey
	}
	requestBytes, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}
	fmt.Println("requestBytes:", string(requestBytes))
	buf := bytes.NewBuffer(requestBytes)
	body, err := c.doRequest("openapi/getOpenPointInfo", buf)
	if err != nil {
		return nil, err
	}
	fmt.Printf("body:%+v", string(body))
	response := GetOpenPointInfoResponse{}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 查询用户下设备列表
func (c *Client) GetDeviceListByUser(request GetDeviceListByUserRequest) (*GetDeviceListByUserResponse, error) {
	if request.Appkey == "" {
		request.Appkey = c.appkey
	}
	requestBytes, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}
	fmt.Println("requestBytes:", string(requestBytes))
	buf := bytes.NewBuffer(requestBytes)
	body, err := c.doRequest("openapi/getDeviceListByUser", buf)
	if err != nil {
		return nil, err
	}
	fmt.Printf("body:%+v", string(body))
	response := GetDeviceListByUserResponse{}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 根据设备SN查询对应的通信设备信息
func (c *Client) GetCommunicationDevInfoByDevSn(request GetCommunicationDevInfoByDevSnRequest) (*GetCommunicationDevInfoByDevSnResponse, error) {
	if request.Appkey == "" {
		request.Appkey = c.appkey
	}
	requestBytes, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}
	fmt.Println("requestBytes:", string(requestBytes))
	buf := bytes.NewBuffer(requestBytes)
	body, err := c.doRequest("openapi/getCommunicationDevInfoByDevSn", buf)
	if err != nil {
		return nil, err
	}
	fmt.Printf("body:%+v", string(body))
	response := GetCommunicationDevInfoByDevSnResponse{}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 查询光伏逆变器的实时测点数据
// 根据光伏逆变器的ps_key或SN查询实时测点数据(数据的单位均为最小基本单位)。
// 注：该接口返回的数据的单位都是最小单位，如发电量数据对应单位是wh，功率数据对应的单位是w，电流对应单位是A，电压对应单位是V。
// 该接口只针对光伏逆变器，其他类型的设备可以通过2.7查询同类型设备的历史测点分钟级数据接口传入最近的时间获取。
func (c *Client) GetPVInverterRealTimeData(request GetPVInverterRealTimeDataRequest) (*GetPVInverterRealTimeDataResponse, error) {
	if request.Appkey == "" {
		request.Appkey = c.appkey
	}
	requestBytes, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}
	fmt.Println("requestBytes:", string(requestBytes))
	buf := bytes.NewBuffer(requestBytes)
	body, err := c.doRequest("openapi/getCommunicationDevInfoByDevSn", buf)
	if err != nil {
		return nil, err
	}
	fmt.Printf("body:%+v", string(body))
	response := GetPVInverterRealTimeDataResponse{}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 查询API调用情况
func (c *Client) GetOpenApiCallInfo(request GetOpenApiCallInfoRequest) (*GetOpenApiCallInfoResponse, error) {
	if request.Appkey == "" {
		request.Appkey = c.appkey
	}
	requestBytes, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}
	fmt.Println("requestBytes:", string(requestBytes))
	buf := bytes.NewBuffer(requestBytes)
	body, err := c.doRequest("openapi/getOpenApiCallInfo", buf)
	if err != nil {
		return nil, err
	}
	fmt.Printf("body:%+v", string(body))
	response := GetOpenApiCallInfoResponse{}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 批量查询电站的基本信息
func (c *Client) GetBatchPsDetail(request GetBatchPsDetailRequest) (*GetBatchPsDetailResponse, error) {
	if request.Appkey == "" {
		request.Appkey = c.appkey
	}
	requestBytes, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}
	fmt.Println("requestBytes:", string(requestBytes))
	buf := bytes.NewBuffer(requestBytes)
	body, err := c.doRequest("openapi/getBatchPsDetail", buf)
	if err != nil {
		return nil, err
	}
	fmt.Printf("body:%+v", string(body))
	response := GetBatchPsDetailResponse{}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 查询设备组串配置信息
func (c *Client) GetDeviceStringInfo(request GetDeviceStringInfoRequest) (*GetDeviceStringInfoResponse, error) {
	if request.Appkey == "" {
		request.Appkey = c.appkey
	}
	requestBytes, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}
	fmt.Println("requestBytes:", string(requestBytes))
	buf := bytes.NewBuffer(requestBytes)
	body, err := c.doRequest("openapi/getDeviceStringInfo", buf)
	if err != nil {
		return nil, err
	}
	fmt.Printf("body:%+v", string(body))
	response := GetDeviceStringInfoResponse{}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 获取电站下优化器设备列表
func (c *Client) GetMlpeDeviceList(request GetMlpeDeviceListRequest) (*GetMlpeDeviceListResponse, error) {
	if request.Appkey == "" {
		request.Appkey = c.appkey
	}
	requestBytes, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}
	fmt.Println("requestBytes:", string(requestBytes))
	buf := bytes.NewBuffer(requestBytes)
	body, err := c.doRequest("openapi/getMlpeDeviceList", buf)
	if err != nil {
		return nil, err
	}
	fmt.Printf("body:%+v", string(body))
	response := GetMlpeDeviceListResponse{}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 查询优化器的实时数据API
func (c *Client) GetMlpeRealTimeData(request GetMlpeRealTimeDataRequest) (*GetMlpeRealTimeDataResponse, error) {
	if request.Appkey == "" {
		request.Appkey = c.appkey
	}
	requestBytes, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}
	fmt.Println("requestBytes:", string(requestBytes))
	buf := bytes.NewBuffer(requestBytes)
	body, err := c.doRequest("openapi/getMlpeRealTimeData", buf)
	if err != nil {
		return nil, err
	}
	fmt.Printf("body:%+v", string(body))
	response := GetMlpeRealTimeDataResponse{}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 查询优化器的历史测点分钟级数据 取设备分钟数据，主要用于提供给第三方平台使用。
// 注意：时间跨度不能超过30天。注意：该接口返回的数据的单位都是最小单位，
// 如发电量数据对应的单位都是最小单位wh，功率数据对应的单位都是最小单位w。
func (c *Client) GetMlpeMinuteDataList(request GetMlpeMinuteDataListRequest) (*GetMlpeMinuteDataListResponse, error) {
	if request.Appkey == "" {
		request.Appkey = c.appkey
	}
	requestBytes, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}
	fmt.Println("requestBytes:", string(requestBytes))
	buf := bytes.NewBuffer(requestBytes)
	body, err := c.doRequest("openapi/getMlpeMinuteDataList", buf)
	if err != nil {
		return nil, err
	}
	fmt.Printf("body:%+v", string(body))
	response := GetMlpeMinuteDataListResponse{}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// 查询优化器的的历史测点日月年数据
// 查询优化器的的历史测点日月年数据 根据开始时间和结束时间以及设备ps_key，和数据类型来获取设备的日月年数据相关数据，支持多测点查询。
// 注意：该接口返回的数据的单位都是最小单位，如发电量数据对应的单位都是最小单位wh，功率数据对应的单位都是最小单位w。
func (c *Client) GetMlpeDayMonthYearDataList(request GetMlpeDayMonthYearDataListRequest) (*GetMlpeDayMonthYearDataListResponse, error) {
	if request.Appkey == "" {
		request.Appkey = c.appkey
	}
	requestBytes, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}
	fmt.Println("requestBytes:", string(requestBytes))
	buf := bytes.NewBuffer(requestBytes)
	body, err := c.doRequest("openapi/getMlpeDayMonthYearDataList", buf)
	if err != nil {
		return nil, err
	}
	fmt.Printf("body:%+v", string(body))
	response := GetMlpeDayMonthYearDataListResponse{}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
