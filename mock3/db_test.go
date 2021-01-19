package mock3

import (
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestGetFromDB(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish() // 断言 DB.Get() 方法是否被调用

	m := NewMockDB(ctrl)
	m.EXPECT().Get(gomock.Eq("Tom")).Return(100, errors.New("not exist"))
	m.EXPECT().Get(gomock.Eq("ok100")).Return(100, nil)

	fmt.Println("haha")

	v := GetFromDB(m, "Tom")
	fmt.Println("TestGetFromDB.GetFromDB, v=", v)
	if v != -1 {
		t.Fatal("expected -1, but got", v)
	}

	v = GetFromDB(m, "ok100")
	fmt.Println("TestGetFromDB.GetFromDB, v=", v)
	if v != 100 {
		t.Fatal("expected 100, but got", v)
	}
}
