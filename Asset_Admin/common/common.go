package common

// 网络请求返回数据结构体体
type Result struct {
	Code    int         `json:code`
	Message string      `json:message`
	Data    interface{} `json:data`
}

// 添加素材前台请求信息
type RequestInfo struct {
	PmName     string
	DesignName string
	FinalName  string
	Sort       string
	Recording  string
	Skeleton   string
	Designer   string
}

// 数据库通用资源信息存储结构体
type AssetInfo struct {
	AssetID    int
	PmName     string
	DesignName string
	FinalName  string
	Sort       string
	Recording  string
	Skeleton   string
	Designer   string
}

// 实例化
func NewResult() *Result {
	return &Result{}
}

func NewRequestInfo() *RequestInfo {
	return &RequestInfo{}
}

func NewAssetInfo() *AssetInfo {
	return &AssetInfo{}
}
