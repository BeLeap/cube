package scheduler

import (
	"beleap.dev/cube/node"
	"beleap.dev/cube/task"
)

var (
	_ Scheduler = &Epvm{}
)

type Epvm struct {
	Name string
}

// Pick implements Scheduler.
func (e *Epvm) Pick(scores map[string]float64, candidates []*node.Node) *node.Node {
	panic("unimplemented")
}

// Score implements Scheduler.
func (e *Epvm) Score(t task.Task, nodes []*node.Node) map[string]float64 {
	panic("unimplemented")
}

// SelectCandidateNodes implements Scheduler.
func (e *Epvm) SelectCandidateNodes(t task.Task, nodes []*node.Node) []*node.Node {
	var candidates []*node.Node
	for node := range nodes {
		if checkDisk(t, nodes[node].Disk-nodes[node].DiskAllocated) {
			candidates = append(candidates, nodes[node])
		}
	}

	return candidates
}

func checkDisk(t task.Task, diskAvailable int64) bool {
	return t.Disk <= diskAvailable
}
