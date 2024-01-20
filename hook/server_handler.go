package hook

import (
  "context"
  "encoding/json"
  "github.com/taosdata/driver-go/v3/errors"
  "go-emqx-ex-hook-to-tdengine/container"
  "go-emqx-ex-hook-to-tdengine/db"
  pb "go-emqx-ex-hook-to-tdengine/protobuf"
  "go-emqx-ex-hook-to-tdengine/tdengine"
  utils "go-emqx-ex-hook-to-tdengine/utils"
  "log"
)

var counter *utils.Counter = utils.NewCounter(0, 100)

// Server is used to implement emqx_exhook_v1.s *Server
type Server struct {
  pb.UnimplementedHookProviderServer
}

// HookProviderServer callbacks

func (s *Server) OnProviderLoaded(ctx context.Context, in *pb.ProviderLoadedRequest) (*pb.LoadedResponse, error) {
  counter.Count(1)
  hooks := []*pb.HookSpec{
    &pb.HookSpec{Name: "client.connect"},
    &pb.HookSpec{Name: "client.connack"},
    &pb.HookSpec{Name: "client.connected"},
    &pb.HookSpec{Name: "client.disconnected"},
    &pb.HookSpec{Name: "client.authenticate"},
    &pb.HookSpec{Name: "client.check_acl"},
    &pb.HookSpec{Name: "client.subscribe"},
    &pb.HookSpec{Name: "client.unsubscribe"},
    &pb.HookSpec{Name: "session.created"},
    &pb.HookSpec{Name: "session.subscribed"},
    &pb.HookSpec{Name: "session.unsubscribed"},
    &pb.HookSpec{Name: "session.resumed"},
    &pb.HookSpec{Name: "session.discarded"},
    &pb.HookSpec{Name: "session.takeovered"},
    &pb.HookSpec{Name: "session.terminated"},
    &pb.HookSpec{Name: "message.publish"},
    &pb.HookSpec{Name: "message.delivered"},
    &pb.HookSpec{Name: "message.acked"},
    &pb.HookSpec{Name: "message.dropped"},
  }
  return &pb.LoadedResponse{Hooks: hooks}, nil
}

func (s *Server) OnProviderUnloaded(ctx context.Context, in *pb.ProviderUnloadedRequest) (*pb.EmptySuccess, error) {
  counter.Count(1)
  log.Println("OnProviderUnloaded")
  return &pb.EmptySuccess{}, nil
}

func (s *Server) OnClientConnect(ctx context.Context, in *pb.ClientConnectRequest) (*pb.EmptySuccess, error) {
  counter.Count(1)
  log.Println("OnClientConnect")
  return &pb.EmptySuccess{}, nil
}

func (s *Server) OnClientConnack(ctx context.Context, in *pb.ClientConnackRequest) (*pb.EmptySuccess, error) {
  counter.Count(1)
  log.Println("OnClientConnack")
  return &pb.EmptySuccess{}, nil
}

func (s *Server) OnClientConnected(ctx context.Context, in *pb.ClientConnectedRequest) (*pb.EmptySuccess, error) {
  log.Println("OnClientConnected")
  counter.Count(1)
  return &pb.EmptySuccess{}, nil
}

func (s *Server) OnClientDisconnected(ctx context.Context, in *pb.ClientDisconnectedRequest) (*pb.EmptySuccess, error) {
  log.Println("OnClientDisconnected")
  counter.Count(1)
  return &pb.EmptySuccess{}, nil
}

func (s *Server) OnClientAuthenticate(ctx context.Context, in *pb.ClientAuthenticateRequest) (*pb.ValuedResponse, error) {
  log.Println("OnClientAuthenticate")
  counter.Count(1)
  reply := &pb.ValuedResponse{}
  reply.Type = pb.ValuedResponse_STOP_AND_RETURN
  reply.Value = &pb.ValuedResponse_BoolResult{BoolResult: true}
  return reply, nil
}

func (s *Server) OnClientCheckAcl(ctx context.Context, in *pb.ClientCheckAclRequest) (*pb.ValuedResponse, error) {
  log.Println("OnClientCheckAcl")
  counter.Count(1)
  reply := &pb.ValuedResponse{}
  reply.Type = pb.ValuedResponse_STOP_AND_RETURN
  reply.Value = &pb.ValuedResponse_BoolResult{BoolResult: true}
  return &pb.ValuedResponse{}, nil
}

func (s *Server) OnClientSubscribe(ctx context.Context, in *pb.ClientSubscribeRequest) (*pb.EmptySuccess, error) {
  log.Println("OnClientSubscribe")
  counter.Count(1)
  return &pb.EmptySuccess{}, nil
}

func (s *Server) OnClientUnsubscribe(ctx context.Context, in *pb.ClientUnsubscribeRequest) (*pb.EmptySuccess, error) {
  log.Println("OnClientUnsubscribe")
  counter.Count(1)
  return &pb.EmptySuccess{}, nil
}

func (s *Server) OnSessionCreated(ctx context.Context, in *pb.SessionCreatedRequest) (*pb.EmptySuccess, error) {
  log.Println("OnSessionCreated")
  counter.Count(1)
  return &pb.EmptySuccess{}, nil
}
func (s *Server) OnSessionSubscribed(ctx context.Context, in *pb.SessionSubscribedRequest) (*pb.EmptySuccess, error) {
  log.Println("OnSessionSubscribed")
  counter.Count(1)
  return &pb.EmptySuccess{}, nil
}

func (s *Server) OnSessionUnsubscribed(ctx context.Context, in *pb.SessionUnsubscribedRequest) (*pb.EmptySuccess, error) {
  log.Println("OnSessionUnsubscribed")
  counter.Count(1)
  return &pb.EmptySuccess{}, nil
}

func (s *Server) OnSessionResumed(ctx context.Context, in *pb.SessionResumedRequest) (*pb.EmptySuccess, error) {
  log.Println("OnSessionResumed")
  counter.Count(1)
  return &pb.EmptySuccess{}, nil
}

func (s *Server) OnSessionDiscarded(ctx context.Context, in *pb.SessionDiscardedRequest) (*pb.EmptySuccess, error) {
  log.Println("OnSessionDiscarded")
  counter.Count(1)
  return &pb.EmptySuccess{}, nil
}

func (s *Server) OnSessionTakeovered(ctx context.Context, in *pb.SessionTakeoveredRequest) (*pb.EmptySuccess, error) {
  log.Println("OnSessionTakeovered")
  counter.Count(1)
  return &pb.EmptySuccess{}, nil
}

func (s *Server) OnSessionTerminated(ctx context.Context, in *pb.SessionTerminatedRequest) (*pb.EmptySuccess, error) {
  log.Println("OnSessionTerminated")
  counter.Count(1)
  return &pb.EmptySuccess{}, nil
}

func (s *Server) OnMessagePublish(ctx context.Context, in *pb.MessagePublishRequest) (*pb.ValuedResponse, error) {
  log.Println("OnMessagePublish")
  counter.Count(1)
  message := in.GetMessage()
  log.Println("message:", message)
  if in.Message != nil {
    //in.Message.Payload = []byte("hardcode payload by exhook-svr-go :)")
    //log.Printf("Received message: %s from topic: %s\n", in.Message.Payload, in.Message.Topic)
    var payload map[string]interface{}
    err := json.Unmarshal(in.Message.Payload, &payload)
    if err != nil {
      log.Println("Error in parsing JSON:", err)
    }

    // Replace placeholders in the SQL template
    if container.InsertSqlTemplate != "" {
      sqlString := container.InsertSqlTemplate
      sqlString = db.GetSql(payload, sqlString)
      // 使用示例
      _, err = container.SqlDb.Exec(sqlString)
      if err != nil {
        log.Println("error sql", sqlString)
        tError, is := err.(*errors.TaosError)
        if is {
          errorCode := int(tError.Code)
          log.Println("errorCode:", errorCode)
          log.Println("errorMessage:", tError.ErrStr)
          if errorCode == 11 {
            maxRetries := 10
            for i := 0; i < maxRetries; i++ {
              newDb, err := tdengine.Reconnect(container.AppConfig.TdDriverName, container.AppConfig.TdDataSourceName)
              if err != nil {
                log.Println("error in reconnect")
              } else {
                container.SqlDb = newDb // 更新外部的 sqlDb 变量
                _, err = container.SqlDb.Exec(sqlString)
                if err == nil {
                  break // 如果执行成功则跳出循环
                }
                log.Println("error in reconnect success and exec sql")
              }
            }
          }
        } else {
          log.Println(err.Error())
        }
      }
    }
  }

  reply := &pb.ValuedResponse{}
  reply.Type = pb.ValuedResponse_STOP_AND_RETURN
  reply.Value = &pb.ValuedResponse_Message{Message: in.Message}
  return reply, nil
}

func (s *Server) OnMessageDelivered(ctx context.Context, in *pb.MessageDeliveredRequest) (*pb.EmptySuccess, error) {
  log.Println("OnMessageDelivered")
  counter.Count(1)
  return &pb.EmptySuccess{}, nil
}

func (s *Server) OnMessageDropped(ctx context.Context, in *pb.MessageDroppedRequest) (*pb.EmptySuccess, error) {
  log.Println("OnMessageDropped")
  counter.Count(1)
  return &pb.EmptySuccess{}, nil
}

func (s *Server) OnMessageAcked(ctx context.Context, in *pb.MessageAckedRequest) (*pb.EmptySuccess, error) {
  log.Println("OnMessageAcked")
  counter.Count(1)
  return &pb.EmptySuccess{}, nil
}
