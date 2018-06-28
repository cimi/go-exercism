package tree

import (
	"fmt"
	"sort"
)

type Record struct {
	ID, Parent int
}

// ByID implements sort.Interface for []Record based on
// the ID field.
type ByID []Record

func (a ByID) Len() int           { return len(a) }
func (a ByID) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByID) Less(i, j int) bool { return a[i].ID < a[j].ID }

type Node struct {
	ID       int
	Children []*Node
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}
	sort.Sort(ByID(records))

	nodes := make(map[int]*Node)
	for _, r := range records {
		if r.ID >= len(records) {
			return nil, fmt.Errorf("%d out of bounds", r.ID)
		}
		if r.Parent > r.ID || r.Parent != 0 && r.ID == r.Parent {
			return nil, fmt.Errorf("Invalid node: %d %d", r.Parent, r.ID)
		}
		if _, ok := nodes[r.ID]; ok {
			return nil, fmt.Errorf("Node %d already declared", r.ID)
		}
		nodes[r.ID] = &Node{ID: r.ID}
		if r.ID != r.Parent {
			nodes[r.Parent].Children = append(nodes[r.Parent].Children, nodes[r.ID])
		}
	}
	return nodes[0], nil
}
