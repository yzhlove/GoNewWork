package main

import (
	"log"
	"think-library/ProtoBuffers/chat07/proto"
)

func main() {

	w1 := &proto.Work{Company: "tenent", Address: "shenzhen futian", Email: "12345@email.com", Code: "11112222233334444", IpaTime: "2026-13-32"}
	w2 := &proto.Work{Company: "alibaba ant design", Address: "zhejiang hangzhou", Email: "121382149@email.com", Code: "aaabbbcccbbdabfbasfdbsa", IpaTime: "2026-13-32"}

	u1 := &proto.User{Name: "aa", Age: 111, Birthday: "2021-01-02", Address: "xin long men ke zhan", Email: "ccc.aaa.bbb@email.com"}
	u2 := &proto.User{Name: "bb", Age: 111, Birthday: "2021-01-02", Address: "xin long men ke zhan", Email: "ccc.aaa.bbb@email.com"}
	u3 := &proto.User{Name: "cc", Age: 111, Birthday: "2021-01-02", Address: "xin long men ke zhan", Email: "ccc.aaa.bbb@email.com"}

	uw1 := &proto.User_Work{Name: "aa", Age: 111, Birthday: "2021-01-02", Address: "xin long men ke zhan", Email: "ccc.aaa.bbb@email.com", Work: w1}
	uw2 := &proto.User_Work{Name: "bb", Age: 111, Birthday: "2021-01-02", Address: "xin long men ke zhan", Email: "ccc.aaa.bbb@email.com", Work: w1}
	uw3 := &proto.User_Work{Name: "cc", Age: 111, Birthday: "2021-01-02", Address: "xin long men ke zhan", Email: "ccc.aaa.bbb@email.com", Work: w2}

	w, _ := w1.Marshal()
	log.Print("w ms => ", len(w))

	po := &proto.PutOne{Users: []*proto.User{u1, u2, u3}, Works: []*proto.Work{w1, w2}}

	pt := &proto.PutTwo{[]*proto.User_Work{uw1, uw2, uw3}}

	a, _ := po.Marshal()
	b, _ := pt.Marshal()

	log.Printf("a len = %d b len = %d ", len(a), len(b))

}
