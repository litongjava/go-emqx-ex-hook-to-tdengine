module go-emqx-ex-hook-to-tdengine

go 1.11

replace emqx.io/grpc/exhook => ./

require (
	google.golang.org/grpc v1.36.0
	google.golang.org/protobuf v1.27.1 // indirect
)
