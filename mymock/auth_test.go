package mymock

import "errors"

//https://blog.csdn.net/qq_29648159/article/details/82145690

type authService struct {
}
func (auth *authService) Login (username string,password string) (string,error) {
	return "token", nil
}
func (auth *authService) Logout(token string) error{
	return nil
}

//在这里我们用authService实现了AuthService接口，这样测试Login,Logout就不再需需要依赖网络请求了。而且我们也可以模拟一些错误的情况进行测试：

//模拟登录失败
type authLoginErr struct {
	auth AuthService  //可以使用组合的特性，Logout方法我们不关心，只用“覆盖”Login方法即可
}
func (auth *authLoginErr) Login (username string,password string) (string,error) {
	return "", errors.New("用户名密码错误")
}

//模拟api服务器宕机
type authUnavailableErr struct {
	auth AuthService
}
func (auth *authUnavailableErr) Login (username string,password string) (string,error) {
	return "", errors.New("api服务不可用")
}
func (auth *authUnavailableErr) Logout(token string) error{
	return errors.New("api服务不可用")
}