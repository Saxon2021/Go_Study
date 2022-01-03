package db_common

var DNS string = "Anireel:Wondershare2021@tcp(120.79.76.181:3306)/Anireel?charset=utf8mb4&parseTime=True"
var DNS_test string = "Anireel_test:Wondershare2021@tcp(120.79.76.181:3306)/Anireel_test?charset=utf8mb4&parseTime=True"

const (
	// 正式数据库
	// SQL_Asset_Add_NameInfo = "INSERT INTO nameInfo(pmName,designName,finalName,sort,recording) VALUES(?, ?, ?, ?, ?)"
	SQL_Asset_Add_NameInfo = "INSERT INTO nameInfo_test(pmName,designName,finalName,sort,recording) VALUES(?, ?, ?, ?, ?)"
)

// func Find(data interface{}) (result interface{}) {

// }
