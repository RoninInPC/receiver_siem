package config

import "gopkg.in/ini.v1"

type Config struct {
	Host struct {
		HostAddressServer  string `ini:"host_address_server"`
		HostAddressCommand string `ini:"host_address_command"`
		HostCheck          string `ini:"host_check"`
	} `ini:"host"`
	TelegramBot struct {
		Token       string `ini:"token"`
		AccessStart string `ini:"access_start"`
		Address     string `ini:"address"`
		Port        string `ini:"code_name"`
	} `ini:"telegram_bot"`
	AdminUser struct {
		UserName   string `ini:"username"`
		FirstName  string `ini:"first_name"`
		SecondName string `ini:"second_name"`
		ID         int64  `ini:"id"`
		ChatID     int64  `ini:"chat_id"`
		Rights     int    `ini:"rights"`
	} `ini:"admin_user"`
	RedisUsers struct {
		Address  string `ini:"address"`
		Password string `ini:"password"`
		DB       int    `ini:"db"`
	} `ini:"redis_users"`
}

func ReadFromFile(fileName string) (Config, error) {
	cfg, err := ini.Load(fileName)
	config := Config{}
	if err != nil {
		return Config{}, err
	}
	err = cfg.MapTo(&config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}
