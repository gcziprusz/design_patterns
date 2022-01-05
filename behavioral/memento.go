package main

import "fmt"

type originator interface {
	createSnapshot(state string)
	restoreSnapshot(m snapshotMaker)
}
type origin struct {
	state string
}

func newOriginator(state string) *origin {
	return &origin{state}
}

func (o *origin) createSnapshot() *snapshot {
	return &snapshot{o.state}
}
func (o *origin) restoreSnapshot(s snapshot) {
	o.state = s.get()
}

func (o *origin) setState(state string) {
	o.state = state
}
func (o *origin) getState() string {
	return o.state
}

type historyKeeper interface {
	save(snapshot)
	undo()
}

type history struct {
	snapshots []*snapshot
	origin    *origin
}

func newHistory(o *origin) *history {
	return &history{origin: o}
}
func (h *history) save(s *snapshot) {
	h.snapshots = append(h.snapshots, s)
}
func (h *history) undo() {
	if len(h.snapshots) > 0 {
		h.snapshots = (h.snapshots)[:len(h.snapshots)-1]
		latestSnapshot := (h.snapshots)[len(h.snapshots)-1]
		h.origin.restoreSnapshot(*latestSnapshot)
	}
}

type snapshotMaker interface {
	get() string
}
type snapshot struct {
	state string
}

func (s *snapshot) get() string {
	return s.state
}

func main() {
	origin := newOriginator("STATE_A")
	history := newHistory(origin)

	history.save(origin.createSnapshot())
	fmt.Println("get state", origin.getState())

	origin.setState("STATE_B")
	history.save(origin.createSnapshot())
	fmt.Println("get state", origin.getState())

	origin.setState("STATE_C")
	history.save(origin.createSnapshot())
	fmt.Println("get state", origin.getState())

	history.undo()
	fmt.Println("restored state", origin.getState())
	history.undo()
	fmt.Println("restored state", origin.getState())

}
