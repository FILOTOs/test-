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

	t.nodes = append(t.nodes, n)

	return nil
}

func (t *tree) print() {
	for _, n := range t.nodes {
		if n.parentNodeId == nil {
			fmt.Printf("nodeId: %d, parentNodeId: %v, val: %s\n", n.nodeId, nil, n.val)
		} else {
			fmt.Printf("nodeId: %d, parentNodeId: %v, val: %s\n", n.nodeId, *n.parentNodeId, n.val)
		}
	}
}

func (t *tree) printPretty(level int, parentNodeId *int) {
	for _, n := range t.nodes {
		if (n.parentNodeId == nil && parentNodeId == nil) || (n.parentNodeId != nil && parentNodeId != nil && *n.parentNodeId == *parentNodeId) {
			v := fmt.Sprintf("%%%ds\n", level*10)
			fmt.Printf(v, n.val)
			t.printPretty(level+1, &n.nodeId)
		}
	}
}

func newTree() *tree {
	return &tree{}
}

func main() {
	tree := newTree()
	rootNodeId := 1
	tree.addNode(&node{
		nodeId: rootNodeId,
		val:    "root",
	})

	tree.addNode(&node{
		nodeId:       2,
		parentNodeId: &rootNodeId,
		val:          "node2",
	})
	tree.addNode(&node{
		nodeId:       3,
		parentNodeId: &rootNodeId,
		val:          "node3",
	})
	tree.addNode(&node{
		nodeId:       4,
		parentNodeId: &rootNodeId,
		val:          "node4",
	})
	tree.addNode(&node{
		nodeId:       5,
		parentNodeId: &rootNodeId,
		val:          "node3",
	})

	tree.printPretty(0, nil)
}
