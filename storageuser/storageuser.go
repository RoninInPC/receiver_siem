package storageuser

import "receiver_siem/entity/user"

type StorageUser interface {
	Append(user.User) bool
	Update(user.User) bool
	Delete(string) bool
	GetById(string) user.User
	GetByUserName(string) user.User
	GetAllIDs() []string
	GetUserByRights(int) user.Users
	SetUserNameRights(string, int) bool
	SetByIDRights(string, int) bool
	GetUsers() user.Users
	Contains(string) bool
	IsAdmin(string) bool
	IsHaveCorrectRights(string, int) bool
}
