package main

import "fmt"

type person struct {
	name    string
	surname string
	age     int
}

func (p *person) print() {
	fmt.Printf(p.str())
}

func (p *person) str() string {
	return fmt.Sprintf("%s %s (%d лет)", p.name, p.surname, p.age)
}

type node struct {
	nodeId       int
	parentNodeId *int
	person       *person
}

type familyTree struct {
	nodes []*node
}

func (t *familyTree) addNode(n *node) (*node, error) {

	if n.parentNodeId != nil {
		found := false
		for _, node := range t.nodes {
			if node.nodeId == *n.parentNodeId {
				found = true
				break
			}
		}
		if !found {
			return nil, fmt.Errorf("parentNode not found")
		}
	}

	t.nodes = append(t.nodes, n)

	return n, nil
}

func (t *familyTree) print() {
	for _, n := range t.nodes {
		if n.parentNodeId == nil {
			fmt.Printf("nodeId: %d, parentNodeId: %v, val: %s\n", n.nodeId, nil, n.person.str())
		} else {
			fmt.Printf("nodeId: %d, parentNodeId: %v, val: %s\n", n.nodeId, *n.parentNodeId, n.person.str())
		}
	}
}

func (t *familyTree) printPretty(level int, parentNodeId *int) {
	for _, n := range t.nodes {
		if (n.parentNodeId == nil && parentNodeId == nil) || (n.parentNodeId != nil && parentNodeId != nil && *n.parentNodeId == *parentNodeId) {
			for i := 0; i < level; i++ {
				fmt.Print("   ")
			}
			fmt.Printf("%s\n", n.person.str())
			t.printPretty(level+1, &n.nodeId)
		}
	}
}

func newFamilyTree() *familyTree {
	return &familyTree{}
}

func main() {
	tree := newFamilyTree()
	// granny
	grandFather, err := tree.addNode(&node{
		nodeId: 1,
		person: &person{
			name:    "Андрей",
			surname: "Володин",
			age:     67,
		},
	})
	if err != nil {
		panic(err)
	}
	//
	father, err := tree.addNode(&node{
		nodeId:       2,
		parentNodeId: &grandFather.nodeId,
		person: &person{
			name:    "Алексей",
			surname: "Володин",
			age:     39,
		},
	})
	if err != nil {
		panic(err)
	}

	_, err = tree.addNode(&node{
		nodeId:       3,
		parentNodeId: &father.nodeId,
		person: &person{
			name:    "Макс",
			surname: "Володин",
			age:     15,
		},
	})
	if err != nil {
		panic(err)
	}

	tree.printPretty(0, nil)
}
