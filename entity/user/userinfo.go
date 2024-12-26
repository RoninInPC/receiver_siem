package user

import (
	"receiver_siem/other"
	"strconv"
)

type UserInfo struct {
	UserName   string
	Name       string
	SecondName string
	ChatId     int64
	Rights     int
}

func InitUserInfo(
	userName string,
	name string,
	second string,
	chatid int64,
	rights int) *UserInfo {
	return &UserInfo{
		userName,
		name,
		second,
		chatid,
		rights}
}

func (u UserInfo) ToString(isTelegramFormat bool) string {
	if isTelegramFormat {
		if len(u.UserName) > 0 {
			return other.ToBold("UserName: ") + other.FixString(u.UserName) +
				", " + other.ToBold("Rights: ") + strconv.Itoa(u.Rights)
		}
		return other.ToBold("FirstName: ") + other.FixString(u.Name) +
			", " + other.ToBold("SecondName: ") + other.FixString(u.SecondName) +
			", " + other.ToBold("Rights: ") + strconv.Itoa(u.Rights)
	}

	if len(u.UserName) > 0 {
		return "UserName: " + u.UserName +
			", Rights: " + strconv.Itoa(u.Rights)
	}
	return "FirstName: " + u.Name +
		", SecondName: " + u.SecondName +
		", Rights: " + strconv.Itoa(u.Rights)
}

func (u UserInfo) ToUsernameString() string {
	if len(u.UserName) > 0 {
		return "UserName: " + u.UserName
	}
	return "FirstName: " + u.Name +
		", SecondName: " + u.SecondName
}
