package models

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var users = make(map[string]string)

func CreateUser(creds Credentials) error {
	// 创建用户并保存到内存或数据库
	users[creds.Username] = creds.Password
	return nil
}
