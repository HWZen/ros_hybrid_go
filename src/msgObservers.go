package ros_hybrid_go

import (
	"google.golang.org/protobuf/proto"
)

type msgObserver interface {
	Update(proto.Message)
	Shutdown()
}
