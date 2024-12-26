package user

import (
	"os"
	"receiver_siem/other"
	"strconv"
)

type User struct {
	UserId string
	Info   UserInfo
}

func InitUser(userName string, name string, second string, userid int64, chatid int64, rights int) *User {
	return &User{UserId: strconv.FormatInt(userid, 10),
		Info: *InitUserInfo(userName, name, second, chatid, rights)}
}

func (u User) ToString(isTelegramFormat bool) string {
	if isTelegramFormat {
		return other.ToBold("UserId: ") + u.UserId + ", " + u.Info.ToString(isTelegramFormat)
	}
	return "UserId: " + u.UserId + ", " + u.Info.ToString(isTelegramFormat)
}

type Users []User

func (u Users) ToFile(filename string) bool {
	_, err2 := os.Create(filename)
	if err2 != nil {
		return false
	}
	file, err := os.OpenFile(filename, os.O_RDWR, 0666)
	if err != nil {
		return false
	}
	for _, user := range u {
		_, err := file.WriteString(user.ToString(false) + "\n")
		if err != nil {
			return false
		}
	}
	err = file.Close()
	if err != nil {
		return false
	}
	return true
}

func (u Users) ToString(isTelegramFormat bool) string {
	var answer = ""
	for _, us := range u {
		answer += us.ToString(isTelegramFormat) + "\n\n"
	}
	return answer
}
