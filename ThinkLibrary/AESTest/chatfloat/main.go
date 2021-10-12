package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"think-library/AESTest/chatfloat/proto"
)

func main() {

	str := make([]string, 0, 64)

	for i := 1; i < 20; i++ {
		var number = "0."
		for j := 0; j < i; j++ {
			number += strconv.Itoa(j + 1)
		}
		str = append(str, number)
	}

	arrJSON := make([]JSON_A, 0, len(str))
	arrPROTO := make([]proto.PROTO_A, 0, len(str))

	for _, t := range str {

		f64, err := strconv.ParseFloat(t, 10)
		if err != nil {
			panic(fmt.Sprintf("transfrom error: %s:%v", t, err))
		}

		arrJSON = append(arrJSON, JSON_A{Name1: t, Name2: float32(f64), Name3: f64})
		arrPROTO = append(arrPROTO, proto.PROTO_A{Number1: t, Number2: float32(f64), Number3: f64})
	}

	w := func(size int) string {
		return strings.Repeat("=", size)
	}
	fmt.Println(w(30), "PROTO", w(30))
	for _, pto := range arrPROTO {
		number2 := strconv.FormatFloat(float64(pto.Number2), 'f', -1, 32)
		number3 := strconv.FormatFloat(pto.Number3, 'f', -1, 64)
		if number2 != pto.Number1 {
			fmt.Printf("↓ <f32>失败:[字符串值:（%s）][字符串转浮点数:（%v）][浮点数转回字符串:（%s）]\n", pto.Number1, pto.Number2, number2)
		} else {
			fmt.Printf("↑ <f32>成功:[字符串值:（%s）][字符串转浮点数:（%v）][浮点数转回字符串:（%s）]\n\n", pto.Number1, pto.Number2, number2)
		}
		if number3 != pto.Number1 {
			fmt.Printf("↓ <f64>失败:[字符串值:（%s）][字符串转浮点数:（%v）][浮点数转回字符串:（%s）]\n", pto.Number1, pto.Number3, number3)
		} else {
			fmt.Printf("↑ <f64>成功:[字符串值:（%s）][字符串转浮点数:（%v）][浮点数转回字符串:（%s）]\n\n", pto.Number1, pto.Number3, number3)
		}
	}

	fmt.Println(w(30), "JSON", w(30))
	for _, obj := range arrJSON {
		data, err := json.Marshal(obj)
		if err != nil {
			panic(fmt.Sprintf("json error: %v:%v", obj, err))
		}

		var tObj = &JSON_A{}
		if err := json.Unmarshal(data, tObj); err != nil {
			panic(fmt.Sprintf("json un error:%s:%v", string(data), err))
		}

		number2 := strconv.FormatFloat(float64(tObj.Name2), 'f', -1, 32)
		number3 := strconv.FormatFloat(tObj.Name3, 'f', -1, 64)
		if number2 != tObj.Name1 {
			fmt.Printf("↓ <f32>失败:[字符串值:（%s）][字符串转浮点数:（%v）][浮点数转回字符串:（%s）][序列化字符串:（%s）]\n", tObj.Name1, tObj.Name2, number2, string(data))
		} else {
			fmt.Printf("↑ <f32>成功:[字符串值:（%s）][字符串转浮点数:（%v）][浮点数转回字符串:（%s）][序列化字符串:（%s）]\n\n", tObj.Name1, tObj.Name2, number2, string(data))
		}
		if number3 != tObj.Name1 {
			fmt.Printf("↓ <f64>失败:[字符串值:（%s）][字符串转浮点数:（%v）][浮点数转回字符串:（%s）][序列化字符串:（%s）]\n", tObj.Name1, tObj.Name3, number3, string(data))
		} else {
			fmt.Printf("↑ <f64>成功:[字符串值:（%s）][字符串转浮点数:（%v）][浮点数转回字符串:（%s）][序列化字符串:（%s）]\n\n", tObj.Name1, tObj.Name3, number3, string(data))
		}

	}

}

type JSON_A struct {
	Name1 string
	Name2 float32
	Name3 float64
}
