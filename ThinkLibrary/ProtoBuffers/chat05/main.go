package main

import "log"

func main() {

	Init("127.0.0.1:6379", 1)

	//user := &User{Name: "yzh", Age: 0, Birthday: ""}
	user := &User{Name: "yzh", Age: 22, Birthday: "2012-12-28"}
	data, err := user.MarshalMsg(nil)
	log.Print("data length => ", len(data))
	if err != nil {
		panic(err)
	}

	redis := Get()
	if _, err := redis.Get().Do("SET", "Test_Msgp_Key", data); err != nil {
		panic(err)
	}

	redis.Get().Do("HMSET", "Test_MSGP_HASH", "KEY_1", data, "KEY_2", data)

	user = &User{Name: "yzh", Age: 0, Birthday: ""}
	data2, _ := user.MarshalMsg(nil)
	log.Print("data2 length => ", len(data2))
	redis.Get().Do("HMSET", "Test_MSGP_HASH", "KEY_1", data, "KEY_2", data2)

	/*
		user := &User{Name: "yzh", Age: 22, Birthday: "2012-12-28"}
		{
			"n": "yzh",
			"a": 22,
			"b": "2012-12-28"
		}

		user := &User{Name: "yzh", Age: 0, Birthday: ""}
		{
			"n": "yzh"
		}
	*/

}
