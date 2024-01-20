package app

import (
	"flag"
	"github.com/spf13/viper"
	"go-emqx-ex-hook-to-tdengine/container"
	"go-emqx-ex-hook-to-tdengine/db"
	"go-emqx-ex-hook-to-tdengine/hook"
	"go-emqx-ex-hook-to-tdengine/model"
	pb "go-emqx-ex-hook-to-tdengine/protobuf"
	"go-emqx-ex-hook-to-tdengine/tdengine"
	"google.golang.org/grpc"
	"log"
	"net"
)

func RunGoEmqxExHookToTdengine() {
	configFile := flag.String("c", "config.toml", "Configuration file")
	flag.Parse()

	if err := loadConfig(*configFile); err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// 将配置值赋给全局变量 AppConfig
	container.AppConfig = model.Config{
		ServerAddress:    viper.GetString("serverAddress"),
		TdDriverName:     viper.GetString("tdDriverName"),
		TdDataSourceName: viper.GetString("tdDataSourceName"),
		CreateSql:        viper.GetString("createSql"),
		InsertSql:        viper.GetString("insertSql"),
	}

	container.InsertSqlTemplate = db.GetInsertSqlTemplate(container.AppConfig.InsertSql)
	container.SqlDb = tdengine.ConnectTdEngine(container.AppConfig.TdDriverName, container.AppConfig.TdDataSourceName)
	defer container.SqlDb.Close()
	tdengine.CreateTable(container.SqlDb, container.AppConfig.CreateSql)

	lis, err := net.Listen("tcp", container.AppConfig.ServerAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHookProviderServer(s, &hook.Server{})

	log.Printf("Started gRPC server on %s \n", container.AppConfig.ServerAddress)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func loadConfig(path string) error {
	viper.SetConfigFile(path)
	viper.SetConfigType("toml")
	return viper.ReadInConfig()
}
