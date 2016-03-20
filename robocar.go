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

	motor1A := gpio.NewDirectPinDriver(raspberryAdaptor, "pin", "16")
	motor1B := gpio.NewDirectPinDriver(raspberryAdaptor, "pin", "18")
	motor1E := gpio.NewDirectPinDriver(raspberryAdaptor, "pin", "22")

	motor2A := gpio.NewDirectPinDriver(raspberryAdaptor, "pin", "19")
	motor2B := gpio.NewDirectPinDriver(raspberryAdaptor, "pin", "21")
	motor2E := gpio.NewDirectPinDriver(raspberryAdaptor, "pin", "23")

	HIGH := byte(1)
	LOW := byte(0)

	robot := gobot.NewRobot("dc_motor_bot",
		[]gobot.Connection{raspberryAdaptor},
		[]gobot.Device{motor1A, motor1B, motor1E, motor2A, motor2B, motor2E})

	addedBot := gbot.AddRobot(robot)

	addedBot.AddCommand("right",
		func(params map[string]interface{}) interface{} {
			motor1A.DigitalWrite(HIGH)
			motor1B.DigitalWrite(LOW)
			motor1E.DigitalWrite(HIGH)
			return "Going right!"
		})

	addedBot.AddCommand("left",
		func(params map[string]interface{}) interface{} {
			motor1A.DigitalWrite(LOW)
			motor1B.DigitalWrite(HIGH)
			motor1E.DigitalWrite(HIGH)
			return "Going left!"
		})

	addedBot.AddCommand("forward",
		func(params map[string]interface{}) interface{} {
			motor2A.DigitalWrite(HIGH)
			motor2B.DigitalWrite(LOW)
			motor2E.DigitalWrite(HIGH)
			return "Going forward!"
		})

	addedBot.AddCommand("backward",
		func(params map[string]interface{}) interface{} {
			motor2A.DigitalWrite(LOW)
			motor2B.DigitalWrite(HIGH)
			motor2E.DigitalWrite(HIGH)
			return "Going backward!"
		})

	addedBot.AddCommand("stop_acceleration",
		func(params map[string]interface{}) interface{} {
			motor2E.DigitalWrite(LOW)
			return "Stop acceleration! "
		})

	addedBot.AddCommand("stop_turning",
		func(params map[string]interface{}) interface{} {
			motor1E.DigitalWrite(LOW)
			return "Stop turning! "
		})

	addedBot.AddCommand("stop",
		func(params map[string]interface{}) interface{} {
			motor1E.DigitalWrite(LOW)
			motor2E.DigitalWrite(LOW)
			return "Stop!"
		})

	gbot.Start()
}
