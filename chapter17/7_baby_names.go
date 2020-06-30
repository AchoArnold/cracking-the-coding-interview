package chapter17

type graph struct {
	nodes     []*graphNode
	nodeNames map[string]*graphNode
	edges     [][2]*graphNode
}

func (g *graph) createNode(name string, frequency int) {
	node := createGraphNode(name, frequency)
	g.nodes = append(g.nodes, node)
	g.nodeNames[name] = node
}

func (g *graph) addEdge(name1 string, name2 string) {
	node1 := g.nodeNames[name1]
	node2 := g.nodeNames[name2]

	node1.addEdge(node2)
	node2.addEdge(node1)

	g.edges = append(g.edges, [2]*graphNode{node1, node2})
}

func (g *graph) getNodes() []*graphNode {
	return g.nodes
}

func newGraph() *graph {
	return &graph{
		nodes:     []*graphNode{},
		nodeNames: map[string]*graphNode{},
		edges:     [][2]*graphNode{},
	}
}

type graphNode struct {
	isVisited  bool
	name       string
	frequency  int
	neighbours []*graphNode
}

func (node *graphNode) addEdge(node2 *graphNode) {
	node.neighbours = append(node.neighbours, node2)
}

func (node *graphNode) setIsVisited(isVisited bool) {
	node.isVisited = isVisited
}

func createGraphNode(name string, frequency int) *graphNode {
	return &graphNode{
		isVisited:  false,
		name:       name,
		frequency:  frequency,
		neighbours: []*graphNode{},
	}
}

// TrulyMostPopular returns the most popular name taking into account synonyms
func TrulyMostPopular(names map[string]int, synonyms [][]string) map[string]int {
	graph := constructGraph(names)
	connectEdges(graph, synonyms)

	// find components
	rootNames := getTrueFrequencies(graph)
	return rootNames
}

// connect synonymous spellings
func connectEdges(graph *graph, synonyms [][]string) {
	for _, entry := range synonyms {
		graph.addEdge(entry[0], entry[1])
	}
}

// Do DFS of each component. If a node has been visited before, then it's component has already been computed
func getTrueFrequencies(graph *graph) map[string]int {
	rootNames := map[string]int{}

	for _, node := range graph.getNodes() {
		frequency := getComponentFrequency(node)
		name := node.name
		if frequency > 0 {
			rootNames[name] = frequency
		}
	}

	return rootNames
}

// do depth-first search to find the total frequency of this component and mark each node as visited
func getComponentFrequency(node *graphNode) int {
	if node.isVisited {
		return 0
	}

	node.setIsVisited(true)
	sum := node.frequency

	for _, child := range node.neighbours {
		sum += getComponentFrequency(child)
	}

	return sum
}

func constructGraph(names map[string]int) *graph {
	graph := newGraph()
	for name, frequency := range names {
		graph.createNode(name, frequency)
	}

	return graph
}
