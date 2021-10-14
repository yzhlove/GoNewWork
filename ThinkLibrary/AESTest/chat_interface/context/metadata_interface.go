package context

import "go.uber.org/zap/zapcore"

type LogDataInterface interface {
	Encode() ([]byte, error)
	Decode([]byte) error

	zapcore.ObjectMarshaler
}

type LogSourceInterface interface {
	Get() int //返回sourceType

	LogDataInterface
}
