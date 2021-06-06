package main

import (
	"flag"
	"log"
)

func main() {
	t := flag.String("type", "now", "now today tomorrow drink")
	city := flag.String("city", "武汉", "City")
	flag.Parse()

	switch *t {
	case "now":
		now(*city)
	case "today":
		today(*city)
	case "tomorrow":
		tomorrow(*city)
	case "drink":
		drink()
	default:
		log.Fatalf("type error %s", *t)
	}
}
