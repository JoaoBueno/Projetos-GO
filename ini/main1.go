package main

import (
	"gopkg.in/ini.v1"
)

func main() {
	cfg := ini.Empty()
	cfg.Section("").Key("app_mode").SetValue("production")
	cfg.Section("id").Key("Consumer Key").SetValue("rs_f7yOqYa9K5s45kkun6EBqwcMa")
	cfg.Section("id").Key("Consumer Secret").SetValue("xhVqUCZ69Pyjwo6aIp2PDWd8py8a")

	cfg.SaveTo("my.ini.loca1")

}
