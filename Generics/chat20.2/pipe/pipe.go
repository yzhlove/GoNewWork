package pipe

import (
	"context"
	"encoding/binary"
	"generics-chat/chat20.2/log"
	"go.uber.org/zap"
	"io"
)

type (
	RDI64 = <-chan int64
	WRI64 = chan int64
)

func Read(ctx context.Context, reader io.Reader) RDI64 {
	out := make(WRI64, 128)
	go func() {
		buf := make([]byte, 8)
		for {
			select {
			case <-ctx.Done():
				return
			default:
				n, err := reader.Read(buf)
				if err != nil {
					log.Error("pipe read error", zap.Error(err))
					return
				}
				if n > 0 {
					out <- int64(binary.BigEndian.Uint64(buf))
				}
			}
		}
	}()
	return out
}
