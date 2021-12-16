package main

import "fmt"

// DRESS 'CLASSES' and INTERFACE
type dress interface {
	getColor() string
}

type terroristDress struct {
	color string
}

func (t *terroristDress) getColor() string {
	return t.color
}
func newTerroristDress() *terroristDress {
	return &terroristDress{"red"}
}

type counterTerroristDress struct {
	color string
}

func (c *counterTerroristDress) getColor() string {
	return c.color
}
func newCounterTerroristDress() *counterTerroristDress {
	return &counterTerroristDress{"green"}
}

// DRESS FACTORY CLASS
const (
	terroristDressType        = "tDress"
	counterTerroristDressType = "ctDress"
)

var dressFactorySingleton = &dressFactory{dressMap: make(map[string]dress)}

type dressFactory struct {
	dressMap map[string]dress
}

func (d *dressFactory) getDressByType(dressType string) (dress, error) {
	if d.dressMap[dressType] != nil {
		return d.dressMap[dressType], nil
	}

	if dressType == terroristDressType {
		d.dressMap[dressType] = newTerroristDress()
		return d.dressMap[dressType], nil
	}
	if dressType == counterTerroristDressType {
		d.dressMap[dressType] = newCounterTerroristDress()
		return d.dressMap[dressType], nil
	}
	return nil, fmt.Errorf("wrong dress type passed")
}

func getDressFactorySingleton() *dressFactory {
	return dressFactorySingleton
}

// PLAYER
type player struct {
	dress      dress
	playerType string
	lat, long  int
}

func newPLayer(playerType, dressType string) *player {
	dress, _ := getDressFactorySingleton().getDressByType(dressType)
	return &player{playerType: playerType, dress: dress}
}

func (p *player) newLocation(lat, long int) {
	p.lat = lat
	p.long = long
}

// GAME
type game struct {
	terrorists        []*player
	counterTerrorists []*player
}

func newGame() *game {
	return &game{
		terrorists:        make([]*player, 1),
		counterTerrorists: make([]*player, 1),
	}
}

func (g *game) addTerrorist(dressType string) {
	player := newPLayer("T", dressType)
	g.terrorists = append(g.terrorists, player)
}
func (g *game) addCounterTerrorist(dressType string) {
	player := newPLayer("CT", dressType)
	g.counterTerrorists = append(g.counterTerrorists, player)
}

// CLIENT CODE

func main() {
	game := newGame()
	game.addTerrorist(terroristDressType)
	game.addTerrorist(terroristDressType)
	game.addTerrorist(terroristDressType)
	game.addTerrorist(terroristDressType)

	game.addCounterTerrorist(counterTerroristDressType)
	game.addCounterTerrorist(counterTerroristDressType)
	game.addCounterTerrorist(counterTerroristDressType)

	for dressType, dress := range getDressFactorySingleton().dressMap {
		fmt.Printf("DressColorType: %s\nDressColor: %s\n", dressType, dress.getColor())
	}
}
