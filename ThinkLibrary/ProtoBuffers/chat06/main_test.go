package main

import (
	"testing"
)

func Benchmark_Speed(b *testing.B) {

	stu := &Stu{Name: "abcdefgh", Age: 20000001, Birthday: "protoc-gen-gogofaster"}
	dta, _ := stu.MarshalMsg(nil)
	user := &User{Name: "abcdefgh", Age: 20000001, Birthday: "protoc-gen-gogofaster"}
	data, _ := user.Marshal(nil)

	b.Run("msgp-encode", func(b *testing.B) {
		b.ResetTimer()
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			_, _ = stu.MarshalMsg(nil)
		}
		b.StartTimer()
		b.ReportAllocs()
	})

	b.Run("msgp-decode", func(b *testing.B) {
		b.ResetTimer()
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			_, _ = stu.UnmarshalMsg(dta)
		}
		b.StartTimer()
		b.ReportAllocs()
	})

	b.Run("gencode-encode", func(b *testing.B) {
		b.ResetTimer()
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			_, _ = user.Marshal(nil)
		}
		b.StartTimer()
		b.ReportAllocs()
	})

	b.Run("gencode-decode", func(b *testing.B) {
		b.ResetTimer()
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			_, _ = user.Unmarshal(data)
		}
		b.StartTimer()
		b.ReportAllocs()
	})

}
