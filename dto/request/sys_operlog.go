package request

import (
	"time"
)

// 获取操作日志列表结构体
type OperationLogListReq struct {
	Method            string `json:"method" form:"method"`
	Path              string `json:"path" form:"path"`
	Username          string `json:"username" form:"username"`
	Ip                string `json:"ip" form:"ip"`
	Status            string `json:"status" form:"status"`
	PageInfo        // 分页参数
}

// 翻译需要校验的字段名称
func (s OperationLogListReq) FieldTrans() map[string]string {
	m := make(map[string]string, 0)
	m["Status"] = "响应状态码"
	return m
}

// 创建操作日志结构体
type CreateOperationLogRequestStruct struct {
	ApiDesc    string        `json:"apiDesc"`
	Path       string        `json:"path"`
	Method     string        `json:"method"`
	Params     string        `json:"params"`
	Body       string        `json:"body"`
	Data       string        `json:"data"`
	Status     int           `json:"status"`
	Username   string        `json:"username"`
	RoleName   string        `json:"roleName"`
	Ip         string        `json:"ip"`
	IpLocation string        `json:"ipLocation"`
	Latency    time.Duration `json:"latency"`
	UserAgent  string        `json:"userAgent"`
}
