package main

import (
	"errors"
	"github.com/golang/mock/gomock"
	"testing"
)

func Test_FromDataBase(t *testing.T) {
	ct := gomock.NewController(t)
	defer ct.Finish()

	mock := NewMockDeferDataBase(ct)
	mock.EXPECT().Get(gomock.Eq("tom")).Return(100, errors.New("not exists"))

	if v := GetFromDataBase(mock, "tom"); v != -1 {
		t.Fatal("except -1 , but got ", v)
	}
}

func Test_FormDataBase2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	database := NewMockDeferDataBase(ctrl)

	database.EXPECT().Get(gomock.Eq("a")).Return(100, nil)
	t.Log("A:", GetFromDataBase(database, "a"))

	database.EXPECT().Get(gomock.Any()).Return(1000, nil)
	t.Log("B:", GetFromDataBase(database, "bcd"))

	database.EXPECT().Get(gomock.Not("c")).Return(222, nil)
	t.Log("C2:", GetFromDataBase(database, "cd"))

	database.EXPECT().Get(gomock.Eq("")).Return(333, nil)
	t.Log("D2:", GetFromDataBase(database, ""))

}

func Test_FromDataBase3(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	database := NewMockDeferDataBase(ctrl)

	database.EXPECT().Get(gomock.Any()).Do(func(key string) {
		t.Log("key -> ", key)
	}).DoAndReturn(func(key string) (int, error) {
		if key == "succ" {
			return 100, nil
		}
		return -100, errors.New("arguments invalid")
	}).AnyTimes()

	t.Log("A:", GetFromDataBase(database, "failed"))
	t.Log("B:", GetFromDataBase(database, "succ"))

}

func Test_FromDataBase4(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	database := NewMockDeferDataBase(ctrl)
	//MaxTimes最大调用次数  MinTimes最小调用次数 AnyTimes任意调用次数 Times规定的调用次数
	database.EXPECT().Get(gomock.Not("c")).Return(100, nil).Times(3)

	t.Log("A:", GetFromDataBase(database, "a"))
	t.Log("B:", GetFromDataBase(database, "b"))
	t.Log("C:", GetFromDataBase(database, "d"))
	//t.Log("C:", GetFromDataBase(database, "e"))

}
