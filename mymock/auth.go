package mymock

//假设我们有一个依赖http请求的鉴权接口
type AuthService interface{
	Login(username string,password string) (token string,e error)
	Logout(token string) error
}
