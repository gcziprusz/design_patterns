package main

import (
	"errors"
	"fmt"
)

type observer interface {
	update(i item)
	getID() string
}
type item struct {
	name      string
	inStock   bool
	observers []observer
}

type subject interface {
	register(o observer)
	deregister(o observer)
	notifyAll()
}

type customer struct {
	email string
}

func (c *customer) update(i item) {
	fmt.Printf("Send email to %s , %s is in stock\n", c.email, i.name)
}
func (c *customer) getID() string {
	return c.email
}

func (i *item) register(o observer) {
	i.observers = append(i.observers, o)
}

func (i *item) deregister(o observer) {
	ii, _ := i.getIndex(o)
	i.observers = append(i.observers[:ii], i.observers[ii+1:]...)
}

func (i *item) updateAvailability() {
	fmt.Printf("Item %s is now in stock\n", i.name)
	i.inStock = true
	i.notifyAll()
}
func (i *item) getIndex(o observer) (int, error) {
	for c, observer := range i.observers {
		if o.getID() == observer.getID() {
			return c, nil
		}
	}
	return -1, errors.New("not found")
}

func (i *item) notifyAll() {
	for _, o := range i.observers {
		o.update(*i)
	}
}

func main() {
	cust1 := &customer{email: "Joe@email.net"}
	cust2 := &customer{email: "Bob@email.net"}
	item1 := &item{name: "Adidas zero gravity shoe"}

	item1.register(cust1)
	item1.register(cust2)
	fmt.Printf("%s is interested in %s\n", cust1.email, item1.name)
	fmt.Printf("%s is interested in %s\n", cust2.email, item1.name)
	item1.updateAvailability()
}
