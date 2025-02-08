/**
    @author: dongjs
    @date: 2025/2/6
    @description:
**/

package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// 用户登陆
func (c *Client) Login(userAccount, userPassword string) (*LoginResponse, error) {
	loginRequest := LoginRequest{
		AppKey:       c.appkey,
		UserPassword: userPassword,
		UserAccount:  userAccount,
	}
	requestBytes, err := json.Marshal(&loginRequest)
	if err != nil {
		return nil, err
	}
	fmt.Println("requestBytes:", string(requestBytes))
	buf := bytes.NewBuffer(requestBytes)
	body, err := c.doRequest("openapi/login", buf)
	if err != nil {
		return nil, err
	}
	fmt.Printf("body:%+v", string(body))
	response := LoginResponse{}
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

type LoginRequest struct {
	AppKey       string `json:"appkey"`
	UserAccount  string `json:"user_account"`  // dongjs
	UserPassword string `json:"user_password"` // dongjs86
}

type LoginResponse struct {
	ReqSerialNum string `json:"req_serial_num"`
	ResultCode   string `json:"result_code"`
	ResultMsg    string `json:"result_msg"`
	ResultData   struct {
		UserMasterOrgId   string `json:"user_master_org_id"`
		MobileTel         string `json:"mobile_tel"`
		UserName          string `json:"user_name"`
		Language          string `json:"language"`
		Token             string `json:"token"`
		ErrTimes          string `json:"err_times"`
		UserId            string `json:"user_id"`
		LoginState        string `json:"login_state"`
		DisableTime       string `json:"disable_time"`
		CountryName       string `json:"country_name"`
		UserAccount       string `json:"user_account"`
		UserMasterOrgName string `json:"user_master_org_name"`
		Email             string `json:"email"`
		CountryId         string `json:"country_id"`
	} `json:"result_data"`
}
