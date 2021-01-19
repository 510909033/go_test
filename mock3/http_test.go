package mock3

import (
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestLogin(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish() // 断言 DB.Get() 方法是否被调用

	m := NewMockUser(ctrl)

	a1 := m.EXPECT().Login(gomock.Eq(1)).Return(nil, errors.New("not exist")).Times(1)
	a2 := m.EXPECT().Login(gomock.Eq(2)).Return(nil, nil).Times(1)

	gomock.InOrder(a1, a2)

	a, b := UserLogin(m, 1)
	fmt.Println(1, a, b)

	a, b = UserLogin(m, 2)
	fmt.Println(2, a, b)

	//	a, b = UserLogin(m, 3)
	//	fmt.Println(3, a, b)

}
