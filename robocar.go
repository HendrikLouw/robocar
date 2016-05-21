package main

import (
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

	ch1DIR := gpio.NewDirectPinDriver(raspberryAdaptor, "pin", "24")
	ch1PWM := gpio.NewDirectPinDriver(raspberryAdaptor, "pin", "22")
	ch1GND := gpio.NewDirectPinDriver(raspberryAdaptor, "pin", "20")

	ch2DIR := gpio.NewDirectPinDriver(raspberryAdaptor, "pin", "18")
	ch2PWM := gpio.NewDirectPinDriver(raspberryAdaptor, "pin", "16")
	ch2GND := gpio.NewDirectPinDriver(raspberryAdaptor, "pin", "14")

	HIGH := byte(1)
	LOW := byte(0)

	robot := gobot.NewRobot("robocar",
		[]gobot.Connection{raspberryAdaptor},
		[]gobot.Device{ch1DIR, ch1PWM, ch1GND, ch2DIR, ch2PWM, ch2GND})

	roboCar := gbot.AddRobot(robot)

	roboCar.AddCommand("left",
		func(params map[string]interface{}) interface{} {
			ch1DIR.DigitalWrite(HIGH)
			ch1PWM.DigitalWrite(HIGH)
			return "Turning left"
		})
		roboCar.AddCommand("right",
			func(params map[string]interface{}) interface{} {
				ch1DIR.DigitalWrite(LOW)
				ch1PWM.DigitalWrite(HIGH)
				return "Turning right"
			})

	roboCar.AddCommand("backward",
		func(params map[string]interface{}) interface{} {
			ch2DIR.DigitalWrite(LOW)
			ch2PWM.DigitalWrite(HIGH)
		return "Going backward!"
		})

	roboCar.AddCommand("forward",
		func(params map[string]interface{}) interface{} {
			ch2DIR.DigitalWrite(HIGH)
			ch2PWM.DigitalWrite(HIGH)
			return "Going forward!"
		})

	roboCar.AddCommand("stop_acceleration",
		func(params map[string]interface{}) interface{} {
			ch2PWM.DigitalWrite(LOW)
			return "Stop acceleration! "
		})

	roboCar.AddCommand("stop_turning",
		func(params map[string]interface{}) interface{} {
			ch1PWM.DigitalWrite(LOW)
			return "Stop turning! "
		})

	roboCar.AddCommand("stop",
		func(params map[string]interface{}) interface{} {
			ch1PWM.DigitalWrite(LOW)
			ch2PWM.DigitalWrite(LOW)
			return "Stop!"
		})

	gbot.Start()
}
