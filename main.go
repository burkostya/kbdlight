package main

import (
	"log"

	"github.com/godbus/dbus"
)

func main() {
	conn, err := dbus.SystemBus()
	if err != nil {
		log.Fatalf("getting system bus failed: %s", err)
	}

	obj := conn.Object("org.freedesktop.UPower", "/org/freedesktop/UPower/KbdBacklight")
	getCall := obj.Call("org.freedesktop.UPower.KbdBacklight.GetBrightness", 0)

	var current int
	err = getCall.Store(&current)
	if err != nil {
		log.Fatalf("getting keyboard highlight failed failed: %s", err)
	}

	next := toogle(current)
	setCall := obj.Call("org.freedesktop.UPower.KbdBacklight.SetBrightness", 0, next)
	if setCall.Err != nil {
		log.Fatalf("setting keyboard highlight to %d failed failed: %s", next, err)
	}
}

func toogle(val int) int {
	if val >= 2 {
		return 0
	}

	return val + 1
}
