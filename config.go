package main

var CONFIG map[string]string

func init() {
	CONFIG = make(map[string]string)
	CONFIG["BOT_KEY"] = ""
	CONFIG["WEATHER_KEY"] = ""
}
