/**
    @author: dongjs
    @date: 2025/2/6
    @description:
**/

package v1

import "github.com/dongjssd/IsolarcloudApi/common"

// 客户端
type Client struct {
	accessKey string         `json:"access_key"` //密钥
	sysCode   string         `json:"sys_code"`   //系统编码：第三方调用传入901
	appkey    string         `json:"app_key"`    //Appkey
	station   common.Station `json:"station"`    //站点 // cn:中国站  hk:国际站 eu:欧洲站 au:澳洲站
}

// 初始化请求客户端
func InitClient(accessKey, appKey string, station common.Station) (*Client, error) {
	return &Client{
		accessKey: accessKey,
		appkey:    appKey,
		station:   station,
	}, nil
}

// 获取请求地址
func (c *Client) GetRequestUrl() string {
	if c.station == common.StationCn {
		return common.ApiDomainCn
	} else if c.station == common.StationHk {
		return common.ApiDomainHk
	} else if c.station == common.StationAu {
		return common.ApiDomainAu
	} else if c.station == common.StationEu {
		return common.ApiDomainEu
	}
	return common.ApiDomainCn
}
