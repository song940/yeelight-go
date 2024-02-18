package main

import (
	"log"

	"github.com/song940/yeelight-go/yeelight"
)

func main() {
	// y := yeelight.Find()
	y := yeelight.New(&yeelight.Config{
		IP:   "192.168.2.182",
		Port: 55443,
	})

	result, err := y.GetProp("name")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("name:", result.Result[0])

	result, _ = y.SetName("Yeelight")
	log.Println(result)

	result, err = y.GetProp("power")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(result.Result[0])

	result, err = y.SetPower("on", &yeelight.Effect{Effect: "smooth", Duration: 1500}, 2)
	log.Println(result, err)

	result, err = y.SetRGB(
		255,
		&yeelight.Effect{Effect: "smooth", Duration: 1500},
	)
	log.Println(result, err)

	result, err = y.SetBright(10, &yeelight.Effect{Effect: "smooth", Duration: 1500})
	log.Println(result, err)

	result, err = y.GetProp("bright")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("bright:", result.Result[0])

	// result, err = y.Toggle()
	// log.Println(result, err)
}
