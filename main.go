package main

import (
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	adapter := raspi.NewAdaptor()
	blue_led := gpio.NewLedDriver(adapter, "18")
	red_led := gpio.NewLedDriver(adapter, "19")
	green_led := gpio.NewLedDriver(adapter, "21")

	work := func() {
		gobot.Every(1*time.Second, func() {
			blue_led.Toggle()
			red_led.Toggle()
			green_led.Toggle()
		})
	}

	robot := gobot.NewRobot("snapbot",
		[]gobot.Connection{adapter},
		[]gobot.Device{blue_led, red_led, green_led},
		work,
	)

	robot.Start()
}
