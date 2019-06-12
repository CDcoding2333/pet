package configs

//DbConfig db相关的配置
type DbConfig struct {
	// 数据库类型 mysql,postgres,sqlite,mssql
	Driver string
	// 连接字符串
	Source string
	//开启日志
	LogEnabled bool
}
