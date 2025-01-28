package main

import "receiver_siem/service/program"

func main() {
	program := program.InitProgram("config/config.ini")
	program.Work()
}
