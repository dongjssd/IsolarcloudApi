/**
    @author: dongjs
    @date: 2025/2/7
    @description:
**/

package v1

import "time"

type PowerStationListRequest struct {
	PsName      string `json:"ps_name"`       //电站名称(模糊查询)
	PsType      string `json:"ps_type"`       //电站类型：（多个使用英文逗号隔开传入，默认查询所有） 1：地面电站 3：分布式光伏 4：户用光伏 5：户用储能 6：村级电站 7：分布式储能 8：扶贫电站 9：风能电站 10：地面储能电站 12：工商业储能电站
	ShareType   string `json:"share_type"`    //电站的分享类型： 1：查询具有浏览权限的分享电站 2：查询具有管理权限的分享电站 0：查询本人电站 多个类型使用逗号隔开传入，默认查询所有
	ValidFlag   string `json:"valid_flag"`    //需要查询的电站状态： （1：正常，2：停用, 3：接入中）如果不传该节点，则查询1状态的电站。 如果需要查询多个状态，使用逗号隔开。
	OrgId       string `json:"org_id"`        //如果是根据组织id进行查询，则传入，对应login接口返回的 user_master_org_id出参
	CurPage     int    `json:"curPage"`       //页码 必填
	Size        int    `json:"size"`          //每页大小  必填
	UserId      string `json:"user_id"`       //用户id （用户必须是登录账号的下级）
	IsSelfBuilt int    `json:"is_self_built"` //是否自建电站0-全部 1-自建 默认0全部
	Appkey      string `json:"appkey"`        //32	授权码，必传 (接口给客户端系统分配的appkey)。	是
	Token       string `json:"token"`         //40	Token (登陆成功后API返回的token)	是
}

type GetPowerStationDetailRequest struct {
	Sn             string `json:"sn"`                //设备S/N
	IsGetPsRemarks string `json:"is_get_ps_remarks"` //是否获取电站的备注信息： 1：获取， 不传默认不获取
	Appkey         string `json:"appkey"`            //32	授权码，必传 (接口给客户端系统分配的appkey)。	是
	Token          string `json:"token"`             //40	Token (登陆成功后API返回的token)	是
}

type GetDeviceListRequest struct {
	PsId                 string `json:"ps_id"`                   //电站ID 必传
	CurPage              int    `json:"curPage"`                 //页码 必传
	Size                 int    `json:"size"`                    //每页大小 必传
	IsVirtualUnit        string `json:"is_virtual_unit"`         //是否是查询虚拟设备，1：查询的是虚拟设备，0：查询的是物理设备，默认为0，即表示查询物理设备，电站、并网点、单元都属于虚拟设备
	DeviceTypeList       []int  `json:"device_type_list"`        //设备类型列表，限定条件，为空，查看全部 如： 11：电站、 1：逆变器、 3：并网点、 17：单元。具体参考附录中的设备类型定义
	RelState             string `json:"rel_state"`               //设备认领状态：0：未认领，1：认领
	IsGetFirmwareVersion string `json:"is_get_firmware_version"` //是否需要获取设备的固件版本信息 1：需要、0：不需要，默认不需要
	Appkey               string `json:"appkey"`                  //32	授权码，必传 (接口给客户端系统分配的appkey)。	是
	Token                string `json:"token"`                   //40	Token (登陆成功后API返回的token)	是
}

type GetDeviceRealTimeDataRequest struct {
	PointIdList []string `json:"point_id_list"`
	PsKeyList   []string `json:"ps_key_list"`
	DeviceType  int      `json:"device_type"`
	SnList      []string `json:"sn_list"`
	Appkey      string   `json:"appkey"` //32	授权码，必传 (接口给客户端系统分配的appkey)。	是
	Token       string   `json:"token"`  //40	Token (登陆成功后API返回的token)	是
}

type GetDevicePointMinuteDataListRequest struct {
	StartTimeStamp           string   `json:"start_time_stamp"`             //开始时间 (电站所在时区时间) ，格式为：yyyyMMddHHmmss
	Points                   string   `json:"points"`                       //p+测点ID，如p3022,p3024，多个测点使用英文逗号隔开，具体开放测点定义信息通过常用遥测测点获取。
	EndTimeStamp             string   `json:"end_time_stamp"`               //结束时间 （电站所在时区时间），格式为：yyyyMMddHHmmss
	MinuteInterval           int      `json:"minute_interval"`              //分钟的时间间隔 如：15 代表每15分钟的测点数据 如果不传或者为空，则默认查询每5分钟的测点数据
	PsKeyList                []string `json:"ps_key_list"`                  //设备ps_key集合（相同类型的设备）
	IsGetDataAcquisitionTime string   `json:"is_get_data_acquisition_time"` //是否需要返回测点字典，1：需要，0：不需要，不传默认为不需要
	Appkey                   string   `json:"appkey"`                       //32	授权码，必传 (接口给客户端系统分配的appkey)。	是
	Token                    string   `json:"token"`                        //40	Token (登陆成功后API返回的token)	是
}

type GetDevicePointsDayMonthYearDataListRequest struct {
	DataPoint string   `json:"data_point"`  //格式p+测点，多个测点使用逗号隔开。例如：如果查询的设备的ps_key是逆变器设备，则 p1表示逆变器发电量（返回数值对应单位为Wh），p2表示逆变器累计发电量（返回数值对应单位为Wh），p24表示逆变器总有功功率（返回数值对应单位为W），具体开放测点定义信息通过常用遥测测点获取。
	EndTime   string   `json:"end_time"`    //结束时间 说明： query_type为day的时候，日期格式为yyyyMMdd, query_type为month的时候，日期格式为yyyyMM, query_type为year的时候，日期格式为yyyy
	QueryType string   `json:"query_type"`  //查询类型： 查询日数据：1 查询月数据：2 查询年数据：3
	StartTime string   `json:"start_time"`  //开始时间 说明： query_type为day的时候，日期格式为yyyyMMdd, query_type为month的时候，日期格式为yyyyMM, query_type为year的时候，日期格式为yyyy
	PsKeyList []string `json:"ps_key_list"` //设备ps_key集合（相同类型的设备）
	DataType  string   `json:"data_type"`   //1：均值，2：峰值，3：谷值，4：合计值（日维度的数据无合计值，月年维度的数据才有合计值） 多个使用逗号隔开。 如果是可以进行合计的数据，则： 查询日数据的时候，data_type为2， 查询月数据的时候，data_type为4， 查询年数据的时候，data_type为4
	Order     string   `json:"order"`       //排序方式： 1 倒序， 0 正序（按照时间的顺序）
	Appkey    string   `json:"appkey"`      //32	授权码，必传 (接口给客户端系统分配的appkey)。	是
	Token     string   `json:"token"`       //40	Token (登陆成功后API返回的token)	是
}

type GetDevPropertyPointValueRequest struct {
	DeviceType  int      `json:"device_type"`   //设备类型
	PointIdList []string `json:"point_id_list"` //测点id集合，具体开放测点定义信息通过常用遥测测点获取。
	PsKeyList   []string `json:"ps_key_list"`   //设备ps_key集合（相同类型的设备）
	PsId        string   `json:"ps_id"`         //电站ID
	Appkey      string   `json:"appkey"`        //32	授权码，必传 (接口给客户端系统分配的appkey)。	是
	Token       string   `json:"token"`         //40	Token (登陆成功后API返回的token)	是
}

type GetFaultAlarmInfoRequest struct {
	StartTime     string `json:"startTime"`      //开始时间（格式：yyyyMMddHHmm）可为空。注：查询process_status=9或999的故障列表的时候，时间区间只能查询某个月的数据，不能跨月查询。不传入时间，默认查询当月的故障。
	FaultType     string `json:"fault_type"`     //故障类型： 1：故障 2：告警 3：提示 4：建议 （传多个用逗号分割）
	FaultLevel    string `json:"fault_level"`    //故障级别： 1：重要 2：次要 3：一般 4：轻微 (传多个用逗号分割）
	CurPage       int    `json:"curPage"`        //页码
	Size          int    `json:"size"`           //每页大小
	PsId          int    `json:"ps_id"`          //电站id (ps_id不传时默认查询所有已授权电站的故障信息)
	PsKey         string `json:"ps_key"`         //设备ps_key
	FaultCode     string `json:"fault_code"`     //故障唯一标识，如果传入，则查询该故障告警的信息
	FaultName     string `json:"fault_name"`     //故障名称（不为空时进行模糊查询）
	ShareType     string `json:"share_type"`     //电站的分享类型： 1：查询具有浏览权限的分享电站 2：查询具有管理权限的分享电站 0：查询本人电站 多个类型使用逗号隔开传入，默认查询所有
	ProcessStatus string `json:"process_status"` //故障状态： 8：未关闭 9：已关闭 999: 未关闭和已关闭（不能跨月查询） 如果该字段为空，则查询未关闭的故障
	EndTime       string `json:"endTime"`        //结束时间（格式：yyyyMMddHHmm）可为空。 注：查询process_status=9或999的故障列表的时候，时间区间只能查询某个月的数据，不能跨月查询。不传入时间，默认查询当月的故障。
	Appkey        string `json:"appkey"`         //32	授权码，必传 (接口给客户端系统分配的appkey)。	是
	Token         string `json:"token"`          //40	Token (登陆成功后API返回的token)	是
}

type GetOpenPointInfoRequest struct {
	DeviceType    int    `json:"device_type"`     //设备类型
	Type          int    `json:"type"`            //测点类型： 1：遥信测点 2：遥测测点 5：属性测点
	CurPage       int    `json:"curPage"`         //页码
	Size          int    `json:"size"`            //每页大小
	DeviceModelId string `json:"device_model_id"` //设备型号ID 查询遥测测点信息时，如果传入该参数则表示根据型号ID查询该型号支持的遥测测点信息，如果不传入该参数则查询该类型设备的所有开放测点
	Appkey        string `json:"appkey"`          //32	授权码，必传 (接口给客户端系统分配的appkey)。	是
	Token         string `json:"token"`           //40	Token (登陆成功后API返回的token)	是
}

type GetDeviceListByUserRequest struct {
	CurPage              string `json:"curPage"`                 //页码
	Size                 string `json:"size"`                    //每页大小
	PsId                 string `json:"ps_id"`                   //电站ID
	IsVirtualUnit        string `json:"is_virtual_unit"`         //是否是查询虚拟设备，1：查询的是虚拟设备，0：查询的是物理设备，默认为0，即表示查询物理设备，电站、并网点、单元都属于虚拟设备
	DeviceTypeList       []int  `json:"device_type_list"`        //设备类型列表，限定条件，为空，查看全部 如： 11：电站、 1：逆变器、 3：并网点、 17：单元。具体参考附录中的设备类型定义
	RelState             string `json:"rel_state"`               //设备认领状态：0：未认领，1：认领
	IsGetFirmwareVersion string `json:"is_get_firmware_version"` //是否需要获取设备的固件版本信息 1：需要、0：不需要，默认不需要
	ShareType            string `json:"share_type"`              //电站的分享类型： 1：查询具有浏览权限的分享电站 2：查询具有管理权限的分享电站 0：查询本人电站 多个类型使用逗号隔开传入，默认查询所有
	Appkey               string `json:"appkey"`                  //32	授权码，必传 (接口给客户端系统分配的appkey)。	是
	Token                string `json:"token"`                   //40	Token (登陆成功后API返回的token)	是
}

type GetCommunicationDevInfoByDevSnRequest struct {
	DevSn     string `json:"dev_sn"`     //设备SN
	ShareType string `json:"share_type"` //设备sn对应电站的分享类型： 1：查询具有浏览权限的分享电站 2：查询具有管理权限的分享电站 0：查询本人电站 多个类型使用逗号隔开传入，默认查询本人电站
	Appkey    string `json:"appkey"`     //32	授权码，必传 (接口给客户端系统分配的appkey)。	是
	Token     string `json:"token"`      //40	Token (登陆成功后API返回的token)	是
}

type GetPVInverterRealTimeDataRequest struct {
	SnList    []string `json:"sn_list"`     //sn集合，ps_key_list与sn_list只需传入一个
	PsKeyList []string `json:"ps_key_list"` //设备ps_key集合（相同类型的设备）
	Appkey    string   `json:"appkey"`      //32	授权码，必传 (接口给客户端系统分配的appkey)。	是
	Token     string   `json:"token"`       //40	Token (登陆成功后API返回的token)	是
}

type GetOpenApiCallInfoRequest struct {
	Token  string `json:"token"`
	Appkey string `json:"appkey"`
}

type GetBatchPsDetailRequest struct {
	IsGetPsRemarks string `json:"is_get_ps_remarks"` //是否获取电站的备注信息： 1：获取， 不传默认不获取
	Lang           string `json:"lang"`              //国际化语言
	PsIds          string `json:"ps_ids"`            //电站id，多个之间用逗号隔开(最多支持查询20个，ps_ids和sns2个必须填写1个，同时传入时优先使用ps_ids查询，ps_ids为空时才会使用sns查询)
	Sns            string `json:"sns"`               //通讯设备sn，多个之间用逗号隔开(最多支持查询20个，ps_ids为空时才会使用sns查询)
	Appkey         string `json:"appkey"`            //32	授权码，必传 (接口给客户端系统分配的appkey)。	是
	Token          string `json:"token"`             //40	Token (登陆成功后API返回的token)	是
}

type GetDeviceStringInfoRequest struct {
	PsKeyList []string `json:"ps_key_list"`
	Lang      string   `json:"lang"`
	Appkey    string   `json:"appkey"` //32	授权码，必传 (接口给客户端系统分配的appkey)。	是
	Token     string   `json:"token"`  //40	Token (登陆成功后API返回的token)	是
}

type GetMlpeDeviceListRequest struct {
	PsId   string `json:"ps_id"`
	Appkey string `json:"appkey"` //32	授权码，必传 (接口给客户端系统分配的appkey)。	是
	Token  string `json:"token"`  //40	Token (登陆成功后API返回的token)	是
}

type GetMlpeRealTimeDataRequest struct {
	PsKeyList   []string `json:"ps_key_list"`   //优化器的ps_key列表，单次查询不超过100个
	PointIdList []string `json:"point_id_list"` //需要查询的测点列表，现支持测点列表：58101,58103,58104,58105,58106,58107
	Appkey      string   `json:"appkey"`        //32	授权码，必传 (接口给客户端系统分配的appkey)。	是
	Token       string   `json:"token"`         //40	Token (登陆成功后API返回的token)	是
}

type GetMlpeMinuteDataListRequest struct {
	MinuteInterval string   `json:"minute_interval"` //分钟的时间间隔 如：15 代表每15分钟的测点数据 如果不传或者为空，则默认查询每5分钟的测点数据
	IsGetPointDict string   `json:"is_get_point_dict"`
	Points         string   `json:"points"`           //测点id，多个测点使用逗号隔开；当前仅支持p58101,p58103,p58104,p58105,p58106,p58107
	StartTimeStamp string   `json:"start_time_stamp"` //开始时间格式：yyyyMMddHHmmss
	EndTimeStamp   string   `json:"end_time_stamp"`   //结束时间格式：yyyyMMddHHmmss
	PsKeyList      []string `json:"ps_key_list"`      //优化器的ps_key，支持单次最多100个
	Appkey         string   `json:"appkey"`           //32	授权码，必传 (接口给客户端系统分配的appkey)。	是
	Token          string   `json:"token"`            //40	Token (登陆成功后API返回的token)	是
}

type GetMlpeDayMonthYearDataListRequest struct {
	QueryType string   `json:"query_type"`  //查询类型： 查询日数据：1 查询月数据：2 查询年数据：3
	DataType  string   `json:"data_type"`   //1：均值，2：峰值，3：谷值，4：合计值（日维度的数据无合计值，月年维度的数据才有合计值） 多个使用逗号隔开。 如果是可以进行合计的数据，则： 查询日数据的时候，data_type为2， 查询月数据的时候，data_type为4， 查询年数据的时候，data_type为4
	DataPoint string   `json:"data_point"`  //格式p+测点，多个测点使用逗号隔开。例如：如果查询的设备的ps_key是逆变器设备，则 p1表示逆变器发电量（返回数值对应单位为Wh），p2表示逆变器累计发电量（返回数值对应单位为Wh），p24表示逆变器总有功功率（返回数值对应单位为W），具体开放测点定义信息通过常用遥测测点获取。
	StartTime string   `json:"start_time"`  //开始时间 说明： query_type为day的时候，日期格式为yyyyMMdd, query_type为month的时候，日期格式为yyyyMM, query_type为year的时候，日期格式为yyyy
	EndTime   string   `json:"end_time"`    //结束时间 说明： query_type为day的时候，日期格式为yyyyMMdd, query_type为month的时候，日期格式为yyyyMM, query_type为year的时候，日期格式为yyyy
	PsKeyList []string `json:"ps_key_list"` //优化器的ps_key，支持单次最多100个
	Order     string   `json:"order"`       //排序方式： 1 倒序， 0 正序（按照时间的顺序）
	Appkey    string   `json:"appkey"`      //32	授权码，必传 (接口给客户端系统分配的appkey)。	是
	Token     string   `json:"token"`       //40	Token (登陆成功后API返回的token)	是
}

// 调度参数
type ParamSettingCheckRequest struct {
	SetType string `json:"set_type"` //参数设置类型： 0：参数设置， 2：参数回读
	Uuid    string `json:"uuid"`     //设备的uuid，多设备使用英文逗号分隔传入。当set_type=2时，仅支持操作单个设备。
	Appkey  string `json:"appkey"`   //32	授权码，必传 (接口给客户端系统分配的appkey)。	是
	Token   string `json:"token"`    //40	Token (登陆成功后API返回的token)	是
}

type ParamSettingRequest struct {
	SetType      int    `json:"set_type"`      //参数设置类型： 0：参数设置， 2：参数回读， 默认0
	Uuid         string `json:"uuid"`          //设备的uuid，多设备使用英文逗号分隔传入。当set_type=2时，仅支持操作单个设备。
	TaskName     string `json:"task_name"`     //本次参数设置任务名称
	ExpireSecond int    `json:"expire_second"` //参数设置任务的超时时间，单位为秒,数据范围为[0,1800]
	ParamList    []struct {
		ParamCode int    `json:"param_code"` //参数代码，请参考控制参数代码定义
		SetValue  string `json:"set_value"`  //参数的设置值，如果是进行参数回读，该参数传空，否则该参数必传
	} `json:"param_list"` //设置的参数：每个元素的数据类型为Map，key为param_code和set_value, value为参数的测点id和参数设置的值
	Appkey string `json:"appkey"` //32	授权码，必传 (接口给客户端系统分配的appkey)。	是
	Token  string `json:"token"`  //40	Token (登陆成功后API返回的token)	是
}

type GetParamSettingTaskRequest struct {
	TaskId string `json:"task_id"`
	Uuid   string `json:"uuid"`
	Appkey string `json:"appkey"` //32	授权码，必传 (接口给客户端系统分配的appkey)。	是
	Token  string `json:"token"`  //40	Token (登陆成功后API返回的token)	是
}

type ReadOnlyParamSetRequest struct {
	ParamList []int  `json:"param_list"`
	UuidList  []int  `json:"uuid_list"`
	Appkey    string `json:"appkey"` //32	授权码，必传 (接口给客户端系统分配的appkey)。	是
	Token     string `json:"token"`  //40	Token (登陆成功后API返回的token)	是
}

type GetReadOnlyResultRequest struct {
	TaskList []int  `json:"task_list"` //任务id集合
	Appkey   string `json:"appkey"`    //32	授权码，必传 (接口给客户端系统分配的appkey)。	是
	Token    string `json:"token"`     //40	Token (登陆成功后API返回的token)	是
}
type GetReadOnlyParamDefinitionRequest struct {
	Appkey string `json:"appkey"` //32	授权码，必传 (接口给客户端系统分配的appkey)。	是
	Token  string `json:"token"`  //40	Token (登陆成功后API返回的token)	是
}

// 调度参数

// 实时数据
type DataSubscribeGetConfigRequest struct {
	Appkey  string `json:"appkey"`
	SysCode int    `json:"sys_code"`
	Token   string `json:"token"` //40	Token (登陆成功后API返回的token)	是
}

type DataSubscribeStartRequest struct {
	Second int      `json:"second"`  //上传频率（单位秒），默认60s，不能小于5s，如：10就是10s上传一次
	SnList []string `json:"sn_list"` //设备（逆变器）sn，最多10个
	Appkey string   `json:"appkey"`  //32	授权码，必传 (接口给客户端系统分配的appkey)。	是
	Token  string   `json:"token"`   //40	Token (登陆成功后API返回的token)	是
}

type DataSubscribeStopRequest struct {
	SnList []string `json:"sn_list"` //设备sn列表，最多10个
	Appkey string   `json:"appkey"`  //32	授权码，必传 (接口给客户端系统分配的appkey)。	是
	Token  string   `json:"token"`   //40	Token (登陆成功后API返回的token)	是
}

type DataSubscribeGetHisDataRequest struct {
	Appkey    string    `json:"appkey"`
	SysCode   int       `json:"sys_code"`
	SnList    []string  `json:"sn_list"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Token     string    `json:"token"` //40	Token (登陆成功后API返回的token)	是
}

// 实时数据
