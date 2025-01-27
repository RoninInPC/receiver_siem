package config

import "gopkg.in/ini.v1"

type Config struct {
	TelegramBot struct {
		Token       string `ini:"token"`
		AccessStart string `ini:"access_start"`
		CodeName    string `ini:"code_name"`
	} `ini:"telegram_bot"`
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
