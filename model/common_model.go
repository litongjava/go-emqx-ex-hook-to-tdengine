package model

// AppConfig 结构体保存 MQTT 配置信息
type Config struct {
	ServerAddress    string
	TdDriverName     string
	TdDataSourceName string
	CreateSql        string
	InsertSql        string
}
