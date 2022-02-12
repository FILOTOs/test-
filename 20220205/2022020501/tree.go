package main

import "fmt"

type node struct {
	nodeId       int
	parentNodeId *int
	val          string
}

type tree struct {
	nodes []*node
}

func (t *tree) addNode(n *node) error {

	if n.parentNodeId != nil {
		found := false
		for _, node := range t.nodes {
			if node.nodeId == *n.parentNodeId {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("parentNode not found")
		}
	}

	return nil
}

func (t *tree) print() {

}
