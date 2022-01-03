package asset

import (
	Common "Go_Study/Asset_Admin/common"
	Database "Go_Study/Asset_Admin/database"
	"encoding/json"
	"fmt"
	"net/http"
)

// 向数据库中插入一条资源名称信息
func Add_assets_name(writer http.ResponseWriter, reader *http.Request) {

	// 注册一个关闭连接的事件
	defer reader.Body.Close()

	// 解析客户端传来的参数
	reader.ParseForm()
	fmt.Println(reader.PostForm)
	pmName := reader.PostForm.Get("pmName")
	designName := reader.PostForm.Get("designName")
	finalName := reader.PostForm.Get("finalName")
	sort := reader.PostForm.Get("sort")
	recording := reader.PostForm.Get("recording")
	skeleton := reader.PostForm.Get("skeleton")
	designer := reader.PostForm.Get("designer")

	// 实例化一个返回参数的结构体
	var result = Common.NewResult()

	// 示例化一个接收请求参数的结构体
	var requestData = Common.NewRequestInfo()

	// 赋值请求参数到结构体
	requestData.PmName = pmName
	requestData.DesignName = designName
	requestData.FinalName = finalName
	requestData.Sort = sort
	requestData.Recording = recording
	requestData.Skeleton = skeleton
	requestData.Designer = designer

	// 请求参数为空
	if pmName == "" && sort == "" && recording == "" {
		result.Code = 201
		result.Message = "请勿非法访问"
	} else {    // 正常请求

		// 进行数据库操作
		if Database.DB_Asset_Add_NameInfo(*requestData) {
			result.Code = 200
			result.Message = "请求成功"
		} else {
			result.Code = 202
			result.Message = "写入数据库失败"
		}
	}

	// 返回请求参数
	bytes, _ := json.Marshal(result)
	fmt.Fprint(writer, string(bytes))
}

// 查询接口，按照记录查询，按照分类查询，查询所有值，返回一个结构体切片
func Find_assets_info(w http.ResponseWriter, reader *http.Request)(data []Common.AssetInfo){
	
}