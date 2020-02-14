package chapter4

import (
	"container/list"
)

func (graph *Graph) RouteExists(start string, end string) bool {
	root := graph.GetNodeFromName(start)

	visitedNodes := map[interface{}]bool{}

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() != 0 {
		root := queue.Remove(queue.Back()).(GraphNode)
		visitedNodes[root.Value] = true

		for _, nodeName := range root.Children {
			node := graph.GetNodeFromName(nodeName.(string))
			if _, ok := visitedNodes[node.Value]; !ok {
				visitedNodes[node.Value] = true
				queue.PushBack(node)
			}

			if node.Value == end {
				return true
			}
		}
	}

	return false
}

func (graph *Graph) GetNodeFromName(name string) GraphNode {
	for _, node := range graph.Nodes {
		if node.Value == name {
			return node
		}
	}
	return GraphNode{name, []interface{}{}}
}
