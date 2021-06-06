package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hubenchang0515/botsend"
	"github.com/hubenchang0515/weather"
)

func now(city string) {
	now := weather.Now(CONFIG["WEATHER_KEY"], city)
	msg := botsend.NewMarkdownMessage()
	msg.SetString(fmt.Sprintf("># 实时气象\n  > 天气: `%s`\n  > 温度: `%s℃`\n  ", now.Text, now.Temperature))
	log.Printf("%s", msg.Json())
	botsend.Request(CONFIG["BOT_KEY"], msg.Json())
}

func today(city string) {
	daily := weather.Forecast(CONFIG["WEATHER_KEY"], city, 1)
	if len(daily) < 1 {
		log.Fatal("failed to get the weather")
		return
	}

	var text string
	if daily[0].DayText == daily[0].NightText {
		text = fmt.Sprintf("># 今日气象\n  > 天气: `%s`\n  > 温度: `%s~%s`℃\n  ", daily[0].DayText, daily[0].LowTemperature, daily[0].HighTemperature)
	} else {
		text = fmt.Sprintf("># 今日气象\n  > 天气: `%s`转`%s`\n  > 温度: `%s~%s`℃\n  ", daily[0].DayText, daily[0].NightText, daily[0].LowTemperature, daily[0].HighTemperature)
	}

	dayCode, _ := strconv.Atoi(daily[0].DayCode)
	nightCode, _ := strconv.Atoi(daily[0].NightCode)

	if dayCode > 9 || nightCode > 9 {
		text += "> 记得带伞哦\n  "
	}

	msg := botsend.NewMarkdownMessage()
	msg.SetString(text)
	log.Printf("%s", msg.Json())
	botsend.Request(CONFIG["BOT_KEY"], msg.Json())
}

func tomorrow(city string) {
	daily := weather.Forecast(CONFIG["WEATHER_KEY"], city, 2)
	if len(daily) < 2 {
		log.Fatal("failed to get the weather")
		return
	}

	var text string
	if daily[1].DayText == daily[1].NightText {
		text = fmt.Sprintf("># 明日气象\n  > 天气: `%s`\n  > 温度: `%s~%s`℃\n  ", daily[1].DayText, daily[1].LowTemperature, daily[1].HighTemperature)
	} else {
		text = fmt.Sprintf("># 明日气象\n  > 天气: `%s`转`%s`\n  > 温度: `%s~%s`℃\n  ", daily[1].DayText, daily[1].NightText, daily[1].LowTemperature, daily[1].HighTemperature)
	}

	dayCode, _ := strconv.Atoi(daily[1].DayCode)
	nightCode, _ := strconv.Atoi(daily[1].NightCode)

	if dayCode > 9 || nightCode > 9 {
		text += "> 最好把伞带回家哦\n  "
	}

	msg := botsend.NewMarkdownMessage()
	msg.SetString(text)
	log.Printf("%s", msg.Json())
	botsend.Request(CONFIG["BOT_KEY"], msg.Json())
}
