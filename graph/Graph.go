package graph

type Graph struct {
	edges map[interface{}][]interface{}
}

func NewGraph() *Graph {
	g := new(Graph)
	g.edges = make(map[interface{}][]interface{})
	return g
}

func (g Graph) AddEdge(from interface{}, to interface{}) {
	edges := g.edges[from]
	newEdges := append(edges, to)
	g.edges[from] = newEdges
}

func (g Graph) AddBothEdges(from interface{}, to interface{}) {
	g.AddEdge(from, to)
	g.AddEdge(to, from)
}

func (g Graph) ElementsDfs() []interface{} {
	visited := make(map[interface{}]struct{})
	result := make([]interface{}, 0)

	for k, _ := range g.edges {
		g.appendElementsDfs(k, &visited, &result)
	}
	return result
}

func (g Graph) IsConnected(from interface{}, to interface{}) bool {
	visited := make(map[interface{}]struct{})
	return g.isConnectedFrom(from, to, visited)
}

func (g Graph) isConnectedFrom(from interface{}, to interface{}, visited map[interface{}]struct{}) bool {
	_, isVisited := visited[from];
	if isVisited {
		return false
	}

	if from == to {
		return true
	}

	visited[from] = struct{}{}

	for _, conn := range g.edges[from] {
		if g.isConnectedFrom(conn, to, visited) {
			return true
		}
	}
	return false
}

func (g Graph) ElementsBfs() []interface{} {
	visited := make(map[interface{}]struct{})
	result := make([]interface{}, 0)

	for k, _ := range g.edges {
		queue := make([]interface{}, 0)
		queue = append(queue, k)

		for len(queue) > 0 {
			first := queue[0]
			queue = queue[1:]

			_, isVisited := visited[k];
			if isVisited {
				continue
			}
			result = append(result, first)
			visited[first] = struct{}{}

			for _, e := range g.edges[first] {
				queue = append(queue, e)
			}
		}
	}
	return result
}

func (g Graph) appendElementsDfs(from interface{}, visited *map[interface{}]struct{}, list *[]interface{}) {
	if _, isVisited := (*visited)[from]; isVisited {
		return
	}
	*list = append(*list, from)
	(*visited)[from] = struct{}{}

	for _, v := range g.edges[from] {
		g.appendElementsDfs(v, visited, list)
	}
}
