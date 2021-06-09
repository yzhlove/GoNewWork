package main

func main() {

	Init("127.0.0.1:6379", 1)

	user := &User{Name: "yzh", Age: 0, Birthday: ""}
	data, err := user.MarshalMsg(nil)
	if err != nil {
		panic(err)
	}

	redis := Get()
	if _, err := redis.Get().Do("SET", "Test_Msgp_Key", data); err != nil {
		panic(err)
	}

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
