package main

import "fmt"

type collection interface {
	getIterator() *iterator
	setIterator(*iterator)
}

type myCollection struct {
	users    []user
	iterator iterator
}

func (m *myCollection) getIterator() iterator {
	return m.iterator
}

func (m *myCollection) setIterator(iterator iterator) {
	m.iterator = iterator
}

type iterator interface {
	has() bool
	next() *user
}

type myIterator struct {
	index   int
	reverse bool
	users   []user
}

func (i *myIterator) has() bool {
	if i.reverse == true {
		if i.index >= 0 {
			return true
		}
	} else {
		if i.index < len(i.users) {
			return true
		}
	}
	return false
}

func (i *myIterator) next() *user {
	if i.has() {
		u := i.users[i.index]
		if i.reverse {
			i.index--

		} else {
			i.index++
		}
		return &u
	}
	return nil
}

type user struct {
	name string
	age  int
}

func main() {
	u1 := user{
		"Bob", 33,
	}
	u2 := user{
		"Jan", 25,
	}
	u3 := user{
		"Joe", 89,
	}
	u4 := user{
		"Pista", 11,
	}
	users := []user{u1, u2, u3, u4}
	ascIterator := &myIterator{index: 0, reverse: false, users: users}
	descIterator := &myIterator{index: len(users) - 1, reverse: true, users: users}
	uc := myCollection{users: users, iterator: ascIterator}

	fmt.Println("ASC ORDER")
	it := uc.getIterator()
	for it.has() {
		u := it.next()
		fmt.Println("User: ", u.name, u.age)
	}

	fmt.Println("DESC ORDER")
	uc.setIterator(descIterator)
	it = uc.getIterator()
	for it.has() {
		u := it.next()
		fmt.Println("User: ", u.name, u.age)
	}
}
