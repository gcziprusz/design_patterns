package main

import "fmt"

type computer interface {
	print()
	setPrinter(printer)
}
type printer interface {
	printFile()
}

// computers
type mac struct{ printer printer }

func (m *mac) print() {
	fmt.Println("Print request for mac")
	m.printer.printFile()
}
func (m *mac) setPrinter(printer printer) {
	m.printer = printer
}

type dell struct{ printer printer }

func (d *dell) print() {
	fmt.Println("Print request for dell")
	d.printer.printFile()
}
func (d *dell) setPrinter(printer printer) {
	d.printer = printer
}

type hp struct{ printer printer }

func (h *hp) print() {
	fmt.Println("Print request for hp")
	h.printer.printFile()
}
func (h *hp) setPrinter(printer printer) {
	h.printer = printer
}

// printers
type epson struct{}

func (p *epson) printFile() {
	fmt.Println("Printing by EPSON printer.")
}

type canon struct{}

func (p *canon) printFile() {
	fmt.Println("Printing by CANON printer.")
}

type xerox struct{}

func (p *xerox) printFile() {
	fmt.Println("Printing by XEROX printer.")
}

// client code
func main() {

	canonPrinter := &canon{}
	epsonPrinter := &epson{}
	xeroxPrinter := &xerox{}

	macComputer := &mac{}
	hpComputer := &hp{}
	dellComputer := &dell{}

	macComputer.setPrinter(canonPrinter)
	macComputer.print()

	macComputer.setPrinter(epsonPrinter)
	macComputer.print()

	macComputer.setPrinter(xeroxPrinter)
	macComputer.print()
	fmt.Println()

	hpComputer.setPrinter(canonPrinter)
	hpComputer.print()

	hpComputer.setPrinter(epsonPrinter)
	hpComputer.print()

	hpComputer.setPrinter(xeroxPrinter)
	hpComputer.print()
	fmt.Println()

	dellComputer.setPrinter(canonPrinter)
	dellComputer.print()

	dellComputer.setPrinter(epsonPrinter)
	dellComputer.print()

	dellComputer.setPrinter(xeroxPrinter)
	dellComputer.print()
	fmt.Println()

}
