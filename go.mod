module go-emqx-ex-hook-to-tdengine

go 1.11

replace emqx.io/grpc/exhook => ./

require (
	github.com/spf13/viper v1.18.2
	github.com/taosdata/driver-go/v3 v3.5.1
	google.golang.org/grpc v1.59.0
	google.golang.org/protobuf v1.31.0
)
