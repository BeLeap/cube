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
	var bestNode *node.Node
	var lowestScore float64
	for idx, node := range candidates {
		if idx == 0 {
			bestNode = node
			lowestScore = scores[node.Name]
			continue
		}

		if scores[node.Name] < lowestScore {
			bestNode = node
			lowestScore = scores[node.Name]
		}
	}

	return bestNode
}

// Score implements Scheduler.
func (r *RoundRobin) Score(t task.Task, nodes []*node.Node) map[string]float64 {
	nodeScores := make(map[string]float64)
	var newWorker int
	if r.LastWorker+1 < len(nodes) {
		newWorker = r.LastWorker + 1
		r.LastWorker++
	} else {
		newWorker = 0
		r.LastWorker = 0
	}

	for idx, node := range nodes {
		if idx == newWorker {
			nodeScores[node.Name] = 0.1
		} else {
			nodeScores[node.Name] = 1.0
		}
	}

	return nodeScores
}

// SelectCandidateNodes implements Scheduler.
func (r *RoundRobin) SelectCandidateNodes(t task.Task, nodes []*node.Node) []*node.Node {
	return nodes
}
