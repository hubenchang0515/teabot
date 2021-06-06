package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path"

	"github.com/hubenchang0515/botsend"
)

func drink() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Println(err)
		return
	}

	dir := path.Join(home, "Pictures/tea")

	pic := botsend.NewPictureMessage()
	pictures, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Println(err)
		return
	}

	choice := pictures[rand.Intn(len(pictures))].Name()
	pic.SetPicture(path.Join(dir, choice))
	botsend.Request(CONFIG["BOT_KEY"], pic.Json())

	txt := botsend.NewTextMessgae()
	txt.SetMentionString("@all")
	txt.SetContent("摸会儿鱼，喝杯茶吧。虽然被HR抓到会很惨，但是鱼很好吃，所以值得。")
	botsend.Request(CONFIG["BOT_KEY"], txt.Json())
}
