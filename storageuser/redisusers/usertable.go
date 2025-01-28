package redisusers

import (
	"encoding/json"
	red "github.com/go-redis/redis"
	"log"
	"receiver_siem/entity/user"
)

type RedisUserTable struct {
	Client *red.Client
}

func InitRedisUserTable(addr string, password string, db int) *RedisUserTable {
	return &RedisUserTable{Client: red.NewClient(&red.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})}
}

func (r *RedisUserTable) Append(user user.User) bool {
	if r.Contains(user.UserId) {
		return false
	}
	p, err2 := json.Marshal(user.Info)
	if err2 != nil {
		return false
	}
	if err := r.Client.Set(user.UserId, p, 0).Err(); err != nil {
		return false
	}
	return true
}

func (r *RedisUserTable) Update(user user.User) bool {
	p, err2 := json.Marshal(user.Info)
	if err2 != nil {
		return false
	}
	if err := r.Client.Set(user.UserId, p, 0).Err(); err != nil {
		return false
	}
	return true
}

func (r *RedisUserTable) Delete(id string) bool {
	if err := r.Client.Del(id).Err(); err != nil {
		return false
	}
	return true
}

func (r *RedisUserTable) GetById(id string) user.User {
	bytes, err := r.Client.Get(id).Bytes()
	if err != nil {
		return user.User{}
	}
	userinfo := user.UserInfo{}
	err = json.Unmarshal(bytes, &userinfo)
	if err != nil {
		return user.User{}
	}
	return user.User{UserId: id, Info: userinfo}
}
func (r *RedisUserTable) GetByUserName(username string) user.User {
	users := r.GetUsers()
	for _, user := range users {
		if user.Info.UserName == username {
			return user
		}
	}
	return user.User{}
}

func (r *RedisUserTable) GetAllIDs() []string {
	keys, err := r.Client.Keys("*").Result()
	if err != nil {
		log.Panic(err)
	}
	return keys
}

func (r *RedisUserTable) GetUserByRights(rights int) user.Users {
	ids := r.GetAllIDs()
	users := make(user.Users, 0)
	for _, id := range ids {
		user := r.GetById(id)
		if user.Info.Rights >= rights {
			users = append(users, user)
		}
	}
	return users
}

func (r *RedisUserTable) SetUserNameRights(username string, rights int) bool {
	var user = r.GetByUserName(username)
	user.Info.Rights = rights
	r.Update(user)
	return true
}

func (r *RedisUserTable) SetByIDRights(id string, rights int) bool {
	var user = r.GetById(id)
	user.Info.Rights = rights
	r.Update(user)
	return true
}

func (r *RedisUserTable) GetUsers() user.Users {
	usernames := r.GetAllIDs()
	users := make(user.Users, 0)
	for _, username := range usernames {
		users = append(users, r.GetById(username))
	}
	return users
}

func (r *RedisUserTable) Contains(id string) bool {
	users := r.GetUsers()
	for _, user := range users {
		if user.UserId == id {
			return true
		}
	}
	return false
}

func (r *RedisUserTable) IsAdmin(username string) bool {

	return r.IsHaveCorrectRights(username, 10)
}

func (r *RedisUserTable) IsHaveCorrectRights(id string, rights int) bool {
	if !r.Contains(id) {
		return false
	}
	users := r.GetUserByRights(rights)

	for _, user := range users {
		if user.UserId == id {
			return true
		}
	}
	return false
}
