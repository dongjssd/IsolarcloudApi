/**
    @author: dongjs
    @date: 2025/2/7
    @description:
**/

package common

var (
	ApiDomainCn = "https://gateway.isolarcloud.com/"    //中国站
	ApiDomainHk = "https://gateway.isolarcloud.com.hk/" //国际站
	ApiDomainEu = "https://gateway.isolarcloud.eu/"     //欧洲站
	ApiDomainAu = "https://augateway.isolarcloud.com/"  //澳洲站
)

type Station string

var (
	StationCn Station = "cn"
	StationHk Station = "hk"
	StationAu Station = "au"
	StationEu Station = "eu"
)
