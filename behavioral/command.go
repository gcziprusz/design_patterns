package main

import "fmt"

type device interface {
	on()
	off()
}

type tv struct {
	isRunning bool
}

func (t *tv) on() {
	t.isRunning = true
	fmt.Println("Turning TV on")
}
func (t *tv) off() {
	t.isRunning = false
	fmt.Println("Turning TV off")
}

type button struct {
	command command
}

func (b *button) press() {
	b.command.execute()
}

type command interface {
	execute()
}
type onCommand struct {
	device device
}

func (c *onCommand) execute() {
	c.device.on()
}

type offCommand struct {
	device device
}

func (c *offCommand) execute() {
	c.device.off()
}

func main() {
	tv := &tv{}

	onCommand := &onCommand{tv}
	offCommand := &offCommand{tv}

	onButton := &button{onCommand}
	offButton := &button{offCommand}

	onButton.press()
	offButton.press()
}
