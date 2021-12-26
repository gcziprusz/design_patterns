package main

import (
	"fmt"
	"sort"
)

type collection interface {
	ascIterator() *iterator
	descIterator() *iterator
	ageIterator() *iterator
}

type myCollection struct {
	users        []user
	asciterator  iterator
	desciterator iterator
	ageiterator  iterator
}

func (m *myCollection) ascIterator() iterator {
	return &myIterator{index: 0, reverse: false, users: m.users}
}

func (m *myCollection) descIterator() iterator {
	return &myIterator{index: len(m.users) - 1, reverse: true, users: m.users}
}

func (m *myCollection) ageIterator(desc bool) iterator {
	return newAgeIterator(desc, m.users)
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

type ageIterator struct {
	index int
	desc  bool
	users []user
}

func newAgeIterator(desc bool, users []user) *ageIterator {
	index := 0
	if desc == true {
		index = len(users) - 1
	}
	sort.SliceStable(users, func(i, j int) bool {
		return users[i].age < users[j].age
	})
	return &ageIterator{
		desc:  desc,
		users: users,
		index: index,
	}
}

func (a *ageIterator) has() bool {
	if a.desc == true {
		if a.index >= 0 {
			return true
		}
	} else {
		if a.index < len(a.users) {
			return true
		}
	}
	return false
}

func (a *ageIterator) next() *user {
	if a.has() {
		u := a.users[a.index]
		if a.desc {
			a.index--

		} else {
			a.index++
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
	uc := myCollection{users: users}

	fmt.Println("ASC INDEX ORDER:")
	it := uc.ascIterator()
	for it.has() {
		u := it.next()
		fmt.Println("User: ", u.name, u.age)
	}

	fmt.Println("DESC INDEX ORDER:")
	it = uc.descIterator()
	for it.has() {
		u := it.next()
		fmt.Println("User: ", u.name, u.age)
	}
	fmt.Println("ASC AGE ORDER:")
	it = uc.ageIterator(false)
	for it.has() {
		u := it.next()
		fmt.Println("User: ", u.name, u.age)
	}
	fmt.Println("DESC AGE ORDER:")
	it = uc.ageIterator(true)
	for it.has() {
		u := it.next()
		fmt.Println("User: ", u.name, u.age)
	}
}
