/**
    @author: dongjs
    @date: 2025/2/7
    @description:
**/

package v1

import "time"

type PowerStationListResponse struct {
	ReqSerialNum string `json:"req_serial_num"`
	ResultCode   string `json:"result_code"`
	ResultMsg    string `json:"result_msg"`
	ResultData   struct {
		PageList []struct {
			TotalEnergy struct {
				Unit  string `json:"unit"`
				Value string `json:"value"`
			} `json:"total_energy"`
			AlarmCount            int         `json:"alarm_count"`
			Latitude              interface{} `json:"latitude"`
			Description           interface{} `json:"description"`
			TotalIncomeUpdateTime time.Time   `json:"total_income_update_time"`
			ValidFlag             int         `json:"valid_flag"`
			CurrPower             struct {
				Unit  string `json:"unit"`
				Value string `json:"value"`
			} `json:"curr_power"`
			PsFaultStatus            int         `json:"ps_fault_status"`
			DistrictName             string      `json:"district_name"`
			Co2ReduceUpdateTime      time.Time   `json:"co2_reduce_update_time"`
			CityName                 string      `json:"city_name"`
			InstallDate              string      `json:"install_date"`
			BuildStatus              int         `json:"build_status"`
			TodayEnergyUpdateTime    time.Time   `json:"today_energy_update_time"`
			TotalEnergyUpdateTime    time.Time   `json:"total_energy_update_time"`
			PsType                   int         `json:"ps_type"`
			Longitude                interface{} `json:"longitude"`
			TotalCapcityUpdateTime   interface{} `json:"total_capcity_update_time"`
			EquivalentHourUpdateTime time.Time   `json:"equivalent_hour_update_time"`
			PsName                   string      `json:"ps_name"`
			Co2ReduceTotal           struct {
				Unit  string `json:"unit"`
				Value string `json:"value"`
			} `json:"co2_reduce_total"`
			CurrPowerUpdateTime time.Time `json:"curr_power_update_time"`
			TodayIncome         struct {
				Unit  string `json:"unit"`
				Value string `json:"value"`
			} `json:"today_income"`
			GridConnectionStatus int `json:"grid_connection_status"`
			EquivalentHour       struct {
				Unit  string `json:"unit"`
				Value string `json:"value"`
			} `json:"equivalent_hour"`
			Co2ReduceTotalUpdateTime time.Time `json:"co2_reduce_total_update_time"`
			ProvinceName             string    `json:"province_name"`
			PsLocation               string    `json:"ps_location"`
			TotalIncome              struct {
				Unit  string `json:"unit"`
				Value string `json:"value"`
			} `json:"total_income"`
			TotalCapcity struct {
				Unit  string `json:"unit"`
				Value string `json:"value"`
			} `json:"total_capcity"`
			ShareType             string      `json:"share_type"`
			PsCurrentTimeZone     string      `json:"ps_current_time_zone"`
			TodayIncomeUpdateTime time.Time   `json:"today_income_update_time"`
			PsId                  int         `json:"ps_id"`
			GridConnectionTime    interface{} `json:"grid_connection_time"`
			ConnectType           int         `json:"connect_type"`
			TodayEnergy           struct {
				Unit  string `json:"unit"`
				Value string `json:"value"`
			} `json:"today_energy"`
			PsStatus  int `json:"ps_status"`
			Co2Reduce struct {
				Unit  string `json:"unit"`
				Value string `json:"value"`
			} `json:"co2_reduce"`
			FaultCount int `json:"fault_count"`
		} `json:"pageList"`
		RowCount int `json:"rowCount"`
	} `json:"result_data"`
}

type GetPowerStationDetailResponse struct {
	ReqSerialNum string `json:"req_serial_num"`
	ResultCode   string `json:"result_code"`
	ResultMsg    string `json:"result_msg"`
	ResultData   struct {
		DesignCapacity             float64 `json:"design_capacity"`
		AlarmCount                 int     `json:"alarm_count"`
		PsKey                      string  `json:"ps_key"`
		Latitude                   float64 `json:"latitude"`
		Description                string  `json:"description"`
		PsPriceKwh                 string  `json:"ps_price_kwh"`
		PsFaultStatus              int     `json:"ps_fault_status"`
		PsTypeName                 string  `json:"ps_type_name"`
		BuildStatus                int     `json:"build_status"`
		InstallDate                string  `json:"install_date"`
		PsType                     int     `json:"ps_type"`
		Email                      string  `json:"email"`
		Longitude                  float64 `json:"longitude"`
		ParamIncomeUnitName        string  `json:"param_income_unit_name"`
		PsPrice                    string  `json:"ps_price"`
		PsName                     string  `json:"ps_name"`
		ShareUserType              string  `json:"share_user_type"`
		PsLocation                 string  `json:"ps_location"`
		ShareType                  string  `json:"share_type"`
		PsCurrentTimeZone          string  `json:"ps_current_time_zone"`
		UserMobleTel               string  `json:"user_moble_tel"`
		PsId                       int     `json:"ps_id"`
		CommunicationDevDetailList []struct {
			IsEnable int    `json:"is_enable"`
			Sn       string `json:"sn"`
		} `json:"communication_dev_detail_list"`
		ConnectType int `json:"connect_type"`
		PsStatus    int `json:"ps_status"`
		FaultCount  int `json:"fault_count"`
	} `json:"result_data"`
}

type GetDeviceListResponse struct {
	ReqSerialNum string `json:"req_serial_num"`
	ResultCode   string `json:"result_code"`
	ResultMsg    string `json:"result_msg"`
	ResultData   struct {
		PageList []struct {
			ChnnlId            int    `json:"chnnl_id"`
			TypeName           string `json:"type_name"`
			PsKey              string `json:"ps_key"`
			DeviceSn           string `json:"device_sn"`
			DevStatus          string `json:"dev_status"`
			DeviceType         int    `json:"device_type"`
			FactoryName        string `json:"factory_name"`
			Uuid               int    `json:"uuid"`
			GridConnectionDate string `json:"grid_connection_date"`
			DeviceName         string `json:"device_name"`
			DevFaultStatus     int    `json:"dev_fault_status"`
			RelState           int    `json:"rel_state"`
			DeviceCode         int    `json:"device_code"`
			PsId               int    `json:"ps_id"`
			DeviceModelId      int    `json:"device_model_id"`
			CommunicationDevSn string `json:"communication_dev_sn"`
			DeviceModelCode    string `json:"device_model_code"`
		} `json:"pageList"`
		RowCount int `json:"rowCount"`
	} `json:"result_data"`
}
type GetDeviceRealTimeDataResponse struct {
	ReqSerialNum string `json:"req_serial_num"`
	ResultCode   string `json:"result_code"`
	ResultMsg    string `json:"result_msg"`
	ResultData   struct {
		FailPsKeyList   []interface{} `json:"fail_ps_key_list"`
		DevicePointList []struct {
			DevicePoint struct {
				DeviceName         string      `json:"device_name"`
				DevFaultStatus     int         `json:"dev_fault_status"`
				PsKey              string      `json:"ps_key"`
				DeviceSn           interface{} `json:"device_sn"`
				DevStatus          int         `json:"dev_status"`
				PsId               int         `json:"ps_id"`
				CommunicationDevSn interface{} `json:"communication_dev_sn"`
				Uuid               int         `json:"uuid"`
				P83022             string      `json:"p83022"`
				P83033             string      `json:"p83033"`
				DeviceTime         string      `json:"device_time"`
			} `json:"device_point"`
		} `json:"device_point_list"`
	} `json:"result_data"`
}
type GetDevicePointMinuteDataListResponse struct {
	ReqSerialNum string                 `json:"req_serial_num"`
	ResultCode   string                 `json:"result_code"`
	ResultMsg    string                 `json:"result_msg"`
	ResultData   map[string]interface{} `json:"result_data"`
}
type GetDevicePointsDayMonthYearDataListResponse struct {
	ReqSerialNum string                 `json:"req_serial_num"`
	ResultCode   string                 `json:"result_code"`
	ResultMsg    string                 `json:"result_msg"`
	ResultData   map[string]interface{} `json:"result_data"`
}
type GetDevPropertyPointValueResponse struct {
	ReqSerialNum string `json:"req_serial_num"`
	ResultCode   string `json:"result_code"`
	ResultMsg    string `json:"result_msg"`
	ResultData   struct {
		Code                   string `json:"code"`
		PropertyPointValueList []struct {
			PropertyCode  string `json:"property_code"`
			Unit          string `json:"unit"`
			PsKey         string `json:"ps_key"`
			DeviceType    int    `json:"device_type"`
			PropertyValue string `json:"property_value"`
			Uuid          int    `json:"uuid"`
		} `json:"property_point_value_list"`
	} `json:"result_data"`
}
type GetFaultAlarmInfoResponse struct {
	ReqSerialNum string `json:"req_serial_num"`
	ResultCode   string `json:"result_code"`
	ResultMsg    string `json:"result_msg"`
	ResultData   struct {
		PageList []struct {
			PsId                 int    `json:"ps_id"`
			PsKey                string `json:"ps_key"`
			FaultTypeCode        int    `json:"fault_type_code"`
			FaultCode            string `json:"fault_code"`
			FaultType            int    `json:"fault_type"`
			FaultLevel           int    `json:"fault_level"`
			FaultDesc            string `json:"fault_desc"`
			ProcessStatus        int    `json:"process_status"`
			CreateTime           string `json:"create_time"`
			ProcessTime          int64  `json:"process_time"`
			FaultReason          string `json:"fault_reason"`
			GridConnectionStatus int    `json:"grid_connection_status"`
			DeviceModelCode      string `json:"device_model_code"`
			DeviceModel          string `json:"device_model"`
			DeviceName           string `json:"device_name"`
			DeviceType           int    `json:"device_type"`
			PsName               string `json:"ps_name"`
			Uuid                 int    `json:"uuid"`
			TypeName             string `json:"type_name"`
			FaultName            string `json:"fault_name"`
			OverTime             string `json:"over_time"`
		} `json:"pageList"`
		RowCount int `json:"rowCount"`
	} `json:"result_data"`
}
type GetOpenPointInfoResponse struct {
	ReqSerialNum string `json:"req_serial_num"`
	ResultCode   string `json:"result_code"`
	ResultMsg    string `json:"result_msg"`
	ResultData   struct {
		RowCount int `json:"rowCount"`
		PageList []struct {
			PointId         int         `json:"point_id"`
			OpenPointRemark interface{} `json:"open_point_remark"`
			StorageUnit     *string     `json:"storage_unit"`
			ShowUnit        *string     `json:"show_unit"`
			DeviceType      int         `json:"device_type"`
			PointName       string      `json:"point_name"`
		} `json:"pageList"`
	} `json:"result_data"`
}

type GetDeviceListByUserResponse struct {
	ReqSerialNum string `json:"req_serial_num"`
	ResultCode   string `json:"result_code"`
	ResultMsg    string `json:"result_msg"`
	ResultData   struct {
		PageList []struct {
			ChnnlId            int    `json:"chnnl_id"`
			TypeName           string `json:"type_name"`
			PsKey              string `json:"ps_key"`
			DeviceSn           string `json:"device_sn"`
			DevStatus          string `json:"dev_status"`
			DeviceType         int    `json:"device_type"`
			FactoryName        string `json:"factory_name"`
			Uuid               int    `json:"uuid"`
			GridConnectionDate string `json:"grid_connection_date"`
			DeviceName         string `json:"device_name"`
			DevFaultStatus     int    `json:"dev_fault_status"`
			RelState           int    `json:"rel_state"`
			DeviceCode         int    `json:"device_code"`
			PsId               int    `json:"ps_id"`
			DeviceModelId      int    `json:"device_model_id"`
			CommunicationDevSn string `json:"communication_dev_sn"`
			DeviceModelCode    string `json:"device_model_code"`
		} `json:"pageList"`
		RowCount int `json:"rowCount"`
	} `json:"result_data"`
}

type GetCommunicationDevInfoByDevSnResponse struct {
	ReqSerialNum string `json:"req_serial_num"`
	ResultCode   string `json:"result_code"`
	ResultMsg    string `json:"result_msg"`
	ResultData   struct {
		DeviceModelCode string `json:"device_model_code"`
		PsId            int    `json:"ps_id"`
		PsKey           string `json:"ps_key"`
		PsName          string `json:"ps_name"`
		Sn              string `json:"sn"`
		Uuid            int    `json:"uuid"`
	} `json:"result_data"`
}

type GetPVInverterRealTimeDataResponse struct {
	ReqSerialNum string `json:"req_serial_num"`
	ResultCode   string `json:"result_code"`
	ResultMsg    string `json:"result_msg"`
	ResultData   struct {
		FailSnList      []string      `json:"fail_sn_list"`
		DevicePointList []interface{} `json:"device_point_list"`
	} `json:"result_data"`
}

type GetOpenApiCallInfoResponse struct {
	ReqSerialNum string `json:"req_serial_num"`
	ResultCode   string `json:"result_code"`
	ResultMsg    string `json:"result_msg"`
	ResultData   struct {
		PerDayAccessTimesConfig  string `json:"per_day_access_times_config"`
		CurrHourAccessedTimes    string `json:"curr_hour_accessed_times"`
		TodayAccessedTimes       string `json:"today_accessed_times"`
		PerHourAccessTimesConfig []struct {
			Hour         string `json:"hour"`
			MaxCallTimes int    `json:"max_call_times"`
		} `json:"per_hour_access_times_config"`
		PerHourResidueTimes []struct {
			Times string `json:"times"`
			Hour  string `json:"hour"`
		} `json:"per_hour_residue_times"`
	} `json:"result_data"`
}

type GetBatchPsDetailResponse struct {
	ReqSerialNum string `json:"req_serial_num"`
	ResultCode   string `json:"result_code"`
	ResultMsg    string `json:"result_msg"`
	ResultData   struct {
		Code              string   `json:"code"`
		IsNotBelongToUser []string `json:"is_not_belong_to_user"`
		DataList          []struct {
			DesignCapacity float64     `json:"design_capacity"`
			AlarmCount     int         `json:"alarm_count"`
			PsKey          string      `json:"ps_key"`
			Latitude       interface{} `json:"latitude"`
			Description    interface{} `json:"description"`
			PsRemarks      struct {
				Remark1 string `json:"remark1"`
				Remark2 string `json:"remark2"`
				Remark3 string `json:"remark3"`
				Remark4 string `json:"remark4"`
				Remark5 string `json:"remark5"`
			} `json:"ps_remarks"`
			PsPriceKwh                 string      `json:"ps_price_kwh"`
			PsFaultStatus              int         `json:"ps_fault_status"`
			PsTypeName                 string      `json:"ps_type_name"`
			BuildStatus                int         `json:"build_status"`
			InstallDate                string      `json:"install_date"`
			PsType                     int         `json:"ps_type"`
			Email                      string      `json:"email"`
			Longitude                  interface{} `json:"longitude"`
			ParamIncomeUnitName        string      `json:"param_income_unit_name"`
			PsPrice                    string      `json:"ps_price"`
			PsName                     string      `json:"ps_name"`
			ShareUserType              interface{} `json:"share_user_type"`
			PsLocation                 string      `json:"ps_location"`
			ShareType                  string      `json:"share_type"`
			PsCurrentTimeZone          string      `json:"ps_current_time_zone"`
			UserMobleTel               string      `json:"user_moble_tel"`
			PsId                       string      `json:"ps_id"`
			CommunicationDevDetailList []struct {
				IsEnable int    `json:"is_enable"`
				Sn       string `json:"sn"`
			} `json:"communication_dev_detail_list"`
			ConnectType int `json:"connect_type"`
			PsStatus    int `json:"ps_status"`
			FaultCount  int `json:"fault_count"`
		} `json:"dataList"`
	} `json:"result_data"`
}

type GetDeviceStringInfoResponse struct {
	ReqSerialNum string `json:"req_serial_num"`
	ResultCode   string `json:"result_code"`
	ResultMsg    string `json:"result_msg"`
	ResultData   struct {
		DeviceStringInfo []struct {
			PvCount    int    `json:"pv_count"`
			UpdateTime string `json:"update_time"`
			PsKey      string `json:"ps_key"`
			PvString   []struct {
				PvValue string `json:"pv_value"`
				PvName  string `json:"pv_name"`
			} `json:"pv_string"`
		} `json:"device_string_info"`
		Code      int      `json:"code"`
		FailPsKey []string `json:"fail_ps_key"`
	} `json:"result_data"`
	ExceptionStackTrace interface{} `json:"exception_stack_trace"`
}

type GetMlpeDeviceListResponse struct {
	ReqSerialNum string `json:"req_serial_num"`
	ResultCode   string `json:"result_code"`
	ResultMsg    string `json:"result_msg"`
	ResultData   struct {
		DeviceList []struct {
			PsKey string `json:"ps_key"`
			Sn    string `json:"sn"`
		} `json:"device_list"`
		Code string `json:"code"`
	} `json:"result_data"`
}

type GetMlpeRealTimeDataResponse struct {
	ReqSerialNum string `json:"req_serial_num"`
	ResultCode   string `json:"result_code"`
	ResultMsg    string `json:"result_msg"`
	ResultData   struct {
		DevicePointList []struct {
			DevicePoint struct {
				P58101     string `json:"p58101"`
				PsKey      string `json:"ps_key"`
				P58103     string `json:"p58103"`
				DeviceTime int64  `json:"device_time"`
			} `json:"device_point"`
		} `json:"device_point_list"`
	} `json:"result_data"`
}

type GetMlpeMinuteDataListResponse struct {
	ReqSerialNum string                 `json:"req_serial_num"`
	ResultCode   string                 `json:"result_code"`
	ResultMsg    string                 `json:"result_msg"`
	ResultData   map[string]interface{} `json:"result_data"`
}

type GetMlpeDayMonthYearDataListResponse struct {
	ReqSerialNum string                 `json:"req_serial_num"`
	ResultCode   string                 `json:"result_code"`
	ResultMsg    string                 `json:"result_msg"`
	ResultData   map[string]interface{} `json:"result_data"`
}

// 调度返回参数
type ParamSettingCheckResponse struct {
	ReqSerialNum string `json:"req_serial_num"`
	ResultCode   string `json:"result_code"`
	ResultMsg    string `json:"result_msg"`
	ResultData   struct {
		CheckResult   string `json:"check_result"`
		DevResultList []struct {
			CheckResult string `json:"check_result"`
			Uuid        string `json:"uuid"`
			CheckMsg    string `json:"check_msg,omitempty"`
		} `json:"dev_result_list"`
	} `json:"result_data"`
}
type ParamSettingResponse struct {
	ReqSerialNum string `json:"req_serial_num"`
	ResultCode   string `json:"result_code"`
	ResultMsg    string `json:"result_msg"`
	ResultData   struct {
		CheckResult   string `json:"check_result"`
		DevResultList []struct {
			Msg      string `json:"msg"`
			TaskName string `json:"task_name"`
			Code     string `json:"code"`
			TaskId   string `json:"task_id"`
			Uuid     string `json:"uuid"`
		} `json:"dev_result_list"`
	} `json:"result_data"`
}

type GetParamSettingTaskResponse struct {
	ReqSerialNum string `json:"req_serial_num"`
	ResultCode   string `json:"result_code"`
	ResultMsg    string `json:"result_msg"`
	ResultData   struct {
		TaskName       string      `json:"task_name"`
		CommandStatus  int         `json:"command_status"`
		OverTimeZone   interface{} `json:"over_time_zone"`
		CreateTime     string      `json:"create_time"`
		TaskId         int         `json:"task_id"`
		CreateTimeZone time.Time   `json:"create_time_zone"`
		OverTime       interface{} `json:"over_time"`
		ParamList      []struct {
			CommandStatus  int       `json:"command_status"`
			PointId        string    `json:"point_id"`
			ReturnValue    string    `json:"return_value"`
			SetValue       string    `json:"set_value"`
			CreateTime     string    `json:"create_time"`
			ParamCode      string    `json:"param_code"`
			CreateTimeZone time.Time `json:"create_time_zone"`
			SetPrecision   string    `json:"set_precision"`
			PointName      string    `json:"point_name"`
			Unit           string    `json:"unit"`
			UpdateTime     string    `json:"update_time"`
			SetValName     string    `json:"set_val_name"`
			SetValNameVal  string    `json:"set_val_name_val"`
			UpdateTimeZone time.Time `json:"update_time_zone"`
		} `json:"param_list"`
	} `json:"result_data"`
}

type ReadOnlyParamSetResponse struct {
	ReqSerialNum string `json:"req_serial_num"`
	ResultCode   string `json:"result_code"`
	ResultMsg    string `json:"result_msg"`
	ResultData   []struct {
		Code   string `json:"code"`
		TaskId string `json:"task_id"`
		Uuid   string `json:"uuid"`
	} `json:"result_data"`
}

type GetReadOnlyResultResponse struct {
	ReqSerialNum string `json:"req_serial_num"`
	ResultCode   string `json:"result_code"`
	ResultMsg    string `json:"result_msg"`
	ResultData   []struct {
		Code           string `json:"code"`
		TaskDetailList []struct {
			CommandStatus int         `json:"command_status"`
			PointId       string      `json:"point_id"`
			ReturnValue   interface{} `json:"return_value"`
			Unit          string      `json:"unit"`
			ModbusAddress string      `json:"modbus_address"`
			PsId          int         `json:"ps_id"`
			TaskId        int         `json:"task_id"`
			ParamCode     int         `json:"param_code"`
		} `json:"task_detail_list"`
		TaskId string `json:"task_id"`
		Uuid   string `json:"uuid"`
	} `json:"result_data"`
}

type GetReadOnlyParamDefinitionResponse struct {
	ReqSerialNum string `json:"req_serial_num"`
	ResultCode   string `json:"result_code"`
	ResultMsg    string `json:"result_msg"`
	ResultData   []struct {
		ParamCode int    `json:"param_code"`
		ParamName string `json:"param_name"`
	} `json:"result_data"`
}

//调度返回参数

// 实时数据
type DataSubscribeGetConfigResponse struct {
	ReqSerialNum string `json:"req_serial_num"`
	ResultCode   string `json:"result_code"`
	ResultMsg    string `json:"result_msg"`
	ResultData   struct {
		MqttType         string   `json:"mqtt_type"`
		MqttUrlList      []string `json:"mqtt_url_list"`
		MqttUsername     string   `json:"mqtt_username"`
		MqttPassword     string   `json:"mqtt_password"`
		MqttRsaPublicKey string   `json:"mqtt_rsa_public_key"`
		Code             string   `json:"code"`
	} `json:"result_data"`
}

type DataSubscribeStartResponse struct {
	ReqSerialNum string `json:"req_serial_num"`
	ResultCode   string `json:"result_code"`
	ResultData   struct {
		SuccessList []struct {
			Sn    string `json:"sn"`
			Topic string `json:"topic"`
		} `json:"success_list"`
		FailList []struct {
			Sn   string `json:"sn"`
			Code int    `json:"code"`
		} `json:"fail_list"`
	} `json:"result_data"`
	ResultMsg string `json:"result_msg"`
}

type DataSubscribeStopResponse struct {
	ReqSerialNum string `json:"req_serial_num"`
	ResultCode   string `json:"result_code"`
	ResultData   struct {
		SuccessList []struct {
			Sn string `json:"sn"`
		} `json:"success_list"`
		FailList []struct {
			Sn   string `json:"sn"`
			Code int    `json:"code"`
		} `json:"fail_list"`
	} `json:"result_data"`
	ResultMsg string `json:"result_msg"`
}
type DataSubscribeGetHisDataResponse struct {
	ReqSerialNum string `json:"req_serial_num"`
	ResultCode   string `json:"result_code"`
	ResultMsg    string `json:"result_msg"`
	ResultData   struct {
		FailList    []interface{} `json:"fail_list"`
		SuccessList []struct {
			DataList []struct {
				MsgTime  time.Time `json:"msgTime"`
				RecvTime time.Time `json:"recvTime"`
				Data     []struct {
					Id  string `json:"id"`
					Val string `json:"val"`
				} `json:"data"`
			} `json:"dataList"`
			Sn string `json:"sn"`
		} `json:"success_list"`
	} `json:"result_data"`
}

//实时数据
