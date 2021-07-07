package model

type UserInfo struct {
	Name   string
	UserId string
	Age    int
}

func GetUserFromDb(userId string) *UserInfo {
	switch userId {
	case "1":
		return &UserInfo{"张三", "1", 20}
	case "2":
		return &UserInfo{"李四", "2", 21}
	default:
		return &UserInfo{"default", "0", 1}
	}
}
