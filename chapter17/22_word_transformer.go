package chapter17

func newWordGraphNode(word string) *graphNode {
	return &graphNode{
		isVisited:  false,
		name:       word,
		neighbours: []*graphNode{},
	}
}

type namedPath struct {
	name  string
	path  []string
	depth int
}

func newNamedPath(name string, path []string, depth int) namedPath {
	return namedPath{
		name:  name,
		path:  path,
		depth: depth,
	}
}

func (p *namedPath) append(word string) {
	p.path = append(p.path, word)
}

// TransformWord shows the path for transforming a word from source to target.
func TransformWord(dictionary []string, source string, target string) []string {
	dict := map[string]*graphNode{}

	for _, word := range dictionary {
		dict[word] = nil
	}

	dict[source] = nil
	dict[target] = nil

	for key := range dict {
		dict[key] = newWordGraphNode(key)
	}

	for key := range dict {
		for otherKey := range dict {
			if hasOneWordDifference(key, otherKey) {
				dict[key].addEdge(dict[otherKey])
			}
		}
	}

	depth := 1
	visited := map[string]bool{dict[source].name: true}
	paths := []namedPath{newNamedPath(dict[source].name, []string{dict[source].name}, depth)}
	queue := []*graphNode{dict[source]}

	for len(queue) > 0 {
		var newQueue []*graphNode
		for _, node := range queue {
			for _, childNode := range node.neighbours {
				if childNode.name == target {
					return append(lastNodeEndingIn(paths, depth, node.name), target)
				}

				if _, ok := visited[childNode.name]; !ok {
					newQueue = append(newQueue, childNode)
					paths = addToPath(paths, depth, node.name, childNode.name)
					visited[childNode.name] = true
				}
			}
		}

		depth++
		queue = newQueue
		paths = cleanOldNodes(paths, depth)
	}

	return []string{}
}

func cleanOldNodes(paths []namedPath, depth int) []namedPath {
	var newPaths []namedPath
	for _, path := range paths {
		if path.depth == depth {
			newPaths = append(newPaths, path)
		}
	}
	return newPaths
}

func lastNodeEndingIn(paths []namedPath, depth int, word string) []string {
	for _, path := range paths {
		if path.depth == depth && path.path[len(path.path)-1] == word {
			return path.path
		}
	}

	return []string{}
}

func addToPath(paths []namedPath, depth int, parentNode string, value string) []namedPath {
	var pathsToExtend []int

	for index, path := range paths {
		if path.depth == depth && path.path[len(path.path)-1] == parentNode {
			pathsToExtend = append(pathsToExtend, index)
		}
	}

	for _, index := range pathsToExtend {
		newDepth := depth + 1
		paths = append(paths, newNamedPath(value, append(paths[index].path, value), newDepth))
	}

	return paths
}

func hasOneWordDifference(key, otherKey string) bool {
	count := 0

	for i := 0; i < len(key); i++ {
		if key[i] != otherKey[i] {
			count++
			if count > 1 {
				return false
			}
		}
	}

	return count == 1
}
