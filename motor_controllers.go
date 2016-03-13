package main

import (
	"time"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/api"
	"github.com/hybridgroup/gobot/platforms/gpio"
	"github.com/hybridgroup/gobot/platforms/raspi"
)

func main() {
	gbot := gobot.NewGobot()

	raspberryAdaptor := raspi.NewRaspiAdaptor("raspi")
	gbotApi := api.NewAPI(gbot)
	gbotApi.Debug()
	gbotApi.Start()

	pin := gpio.NewDirectPinDriver(raspberryAdaptor, "pin", "13")
	work := func() {
		level := byte(1)

		gobot.Every(1*time.Second, func() {
			pin.DigitalWrite(level)
			if level == 1 {
				level = 0
			} else {
				level = 1
			}
		})
	}
	robot := gobot.NewRobot("dc_motor_bot",
		[]gobot.Connection{raspberryAdaptor},
		[]gobot.Device{pin},
		work,
	)

	addedBot := gbot.AddRobot(robot)

	addedBot.AddCommand("forward",
		func(params map[string]interface{}) interface{} {
			return "Must implement forward!"
		})

	addedBot.AddCommand("backward",
		func(params map[string]interface{}) interface{} {
			return "Must implement backward!"
		})

	addedBot.AddCommand("left",
		func(params map[string]interface{}) interface{} {
			return "Must implement left!"
		})

	addedBot.AddCommand("right",
		func(params map[string]interface{}) interface{} {
			return "Must implement right!"
		})

	gbot.Start()
}
