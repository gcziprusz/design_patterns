package main

import "fmt"

type department interface {
	execute(*patient)
	setNext(department)
}

type patient struct {
	name              string
	registrationDone  bool
	doctorCheckUpDone bool
	medicineDone      bool
	paymentDone       bool
}

/*********** CONCRETE DEPARTMENTS *****************/
type cashier struct {
	next department
}

func (d *cashier) execute(p *patient) {
	if p.paymentDone {
		fmt.Println("Payment was already Done")
		if d.next != nil {
			d.next.execute(p)
		}
		return
	}

	fmt.Println("Cashier getting money from patient!")
	p.paymentDone = true
	if d.next != nil {
		d.next.execute(p)
	}
}
func (d *cashier) setNext(next department) {
	d.next = next
}

type medical struct {
	next department
}

func (d *medical) execute(p *patient) {
	if p.medicineDone {
		fmt.Println("Medicine already given to patient")
		if d.next != nil {
			d.next.execute(p)
		}
		return
	}
	fmt.Println("Medical giving medicine to patient")
	p.medicineDone = true
	if d.next != nil {
		d.next.execute(p)
	}
}
func (d *medical) setNext(next department) {
	d.next = next
}

type doctor struct {
	next department
}

func (d *doctor) execute(p *patient) {
	if p.doctorCheckUpDone {
		fmt.Println("Doctor checkup already done")
		if d.next != nil {
			d.next.execute(p)
		}
		return
	}
	fmt.Println("Doctor checking patient")
	p.doctorCheckUpDone = true
	if d.next != nil {
		d.next.execute(p)
	}
}
func (d *doctor) setNext(next department) {
	d.next = next
}

type reception struct {
	next department
}

func (d *reception) execute(p *patient) {
	if p.registrationDone {
		fmt.Println("Patient registration already done")
		if d.next != nil {
			d.next.execute(p)
		}
		return
	}
	fmt.Println("Reception registering patient")
	p.registrationDone = true
	if d.next != nil {
		d.next.execute(p)
	}
}
func (d *reception) setNext(next department) {
	d.next = next
}

/*********** CLIENT CODE *****************/
func main() {
	patients := []*patient{{name: "Bob Simple"},
		{name: "Joe Prepaid", paymentDone: true},
		{name: "Maria Preregistered", registrationDone: true, paymentDone: true, medicineDone: true}}
	for _, p := range patients {
		fmt.Printf("\nPatient '%s' visits the Hospital\n\n", p.name)
		clientCode(p)
		fmt.Printf("\nTime to go home '%s' \n", p.name)
	}
}
func clientCode(p *patient) {
	cashier := &cashier{}
	//Set next for medical department
	medical := &medical{}
	medical.setNext(cashier)

	//Set next for doctor department
	doctor := &doctor{}
	doctor.setNext(medical)

	//Set next for reception department
	reception := &reception{}
	reception.setNext(doctor)

	//Patient visiting
	fmt.Println("CHAIN: reception > doctor > medical > cashier")
	reception.execute(p)
}
