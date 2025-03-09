package scheduler

import (
	"beleap.dev/cube/node"
	"beleap.dev/cube/task"
)

var (
	_ Scheduler = &RoundRobin{}
)

type RoundRobin struct {
	Name       string
	LastWorker int
}

// Pick implements Scheduler.
func (r *RoundRobin) Pick(scores map[string]float64, candidates []*node.Node) *node.Node {
	panic("unimplemented")
}

// Score implements Scheduler.
func (r *RoundRobin) Score(t task.Task, nodes []*node.Node) map[string]float64 {
	panic("unimplemented")
}

// SelectCandidateNodes implements Scheduler.
func (r *RoundRobin) SelectCandidateNodes(t task.Task, nodes []*node.Node) []*node.Node {
	panic("unimplemented")
}
