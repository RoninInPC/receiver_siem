package command

import (
	"receiver_siem/storageuser"
)

type Command string

var (
	Rights     Command = "rights"
	Help       Command = "help"
	Start      Command = "start"
	GetUsers   Command = "get_users"
	SetRights  Command = "set_rights"
	TimeNow    Command = "time_now"
	Mute       Command = "mute"
	UnMute     Command = "unmute"
	MuteStatus Command = "mute_status"
	WhoAmI     Command = "whoami"

	accessRights = map[string]int{
		GetNormalCommand(Rights):     5,
		GetNormalCommand(Help):       0,
		GetNormalCommand(Start):      0,
		GetNormalCommand(GetUsers):   10,
		GetNormalCommand(SetRights):  10,
		GetNormalCommand(TimeNow):    5,
		GetNormalCommand(Mute):       5,
		GetNormalCommand(UnMute):     5,
		GetNormalCommand(MuteStatus): 10,
		GetNormalCommand(WhoAmI):     0,
	}
)

func GetNormalCommand(command Command) string {
	return "/" + string(command)
}

type AccessRights interface {
	IsRights() bool
	SetUserName(string)
	SetCommand(string)
}

/*
Пример реализации интерфейса AccessRights для бота, подключенного к базе с пользователями
*/
type AccessRightsTableFunction struct {
	Users    storageuser.StorageUser
	Username string
	Command  string
}

func (access *AccessRightsTableFunction) SetUserName(username string) {
	access.Username = username
}
func (access *AccessRightsTableFunction) SetCommand(command string) {
	access.Command = command
}

func (access *AccessRightsTableFunction) IsRights() bool {
	rights := accessRights[access.Command]
	if rights == 0 {
		return true
	}
	return access.Users.IsHaveCorrectRights(access.Username, rights)
}
