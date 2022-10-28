package graph

type Graph struct {
	Value        interface{}
	AdjacentList []*Graph
	Visited      bool
}

func BFSRetrieveTraversal(graph *Graph) (result []string) {
	if graph == nil {
		return
	}

	q := []*Graph{graph}
	graph.Visited = true

	for len(q) > 0 {
		currentNode := q[0]
		q = q[1:]

		result = append(result, currentNode.Value.(string))

		for _, g := range currentNode.AdjacentList {
			if !g.Visited {
				q = append(q, g)
				g.Visited = true
			}
		}
	}

	return
}

func DFSRetrieveTraversal(graph *Graph) (result []string) {
	if graph == nil {
		return
	}

	graph.Visited = true
	result = append(result, graph.Value.(string))

	for _, current := range graph.AdjacentList {
		if !current.Visited {
			result = append(result, DFSRetrieveTraversal(current)...)
		}
	}

	return

}
