/**
    @author: dongjs
    @date: 2025/2/7
    @description:
**/

package main

import (
	"fmt"
	"github.com/dongjssd/IsolarcloudApi/common"
	v1 "github.com/dongjssd/IsolarcloudApi/v1"
)

// 983288_b32158fb0d4b4a578ea70c6e45a72dce
func main() {
	v1Client, _ := v1.InitClient("iitgiu641e94hfrqviv1hw9jxj9k27n9", "3830501ECFA828A10FA2C6B0DC24A062", common.StationCn)
	//response, err := v1Client.Login("dongjs", "dongjs86!")
	response, err := v1Client.PowerStationList(v1.PowerStationListRequest{
		CurPage: 1,
		Size:    10,
		Token:   "983288_b32158fb0d4b4a578ea70c6e45a72dce",
	})
	if err != nil {
		fmt.Println(err)
	}
	if response != nil {
		fmt.Printf("response:%+v", response)
	}

}
