package chapter4

import (
	"errors"
)

type DependencyGraph struct {
	Nodes      []*Project
	ProjectMap map[string]*Project
}

func (graph *DependencyGraph) GetOrCreateNode(name string) *Project {
	if _, ok := graph.ProjectMap[name]; !ok {
		node := CreateProject(name)
		graph.Nodes = append(graph.Nodes, node)
		graph.ProjectMap[name] = node
	}

	return graph.ProjectMap[name]
}

func (graph *DependencyGraph) AddEdge(startName string, endName string) {
	start := graph.GetOrCreateNode(startName)
	end := graph.GetOrCreateNode(endName)

	start.AddNeighbour(end)
}

func (graph *DependencyGraph) GetNodes() []*Project {
	return graph.Nodes
}

type Project struct {
	Children     []*Project
	Name         string
	dependencies int
	ProjectMap   map[string]*Project
}

func CreateProject(name string) *Project {
	return &Project{
		Children:     nil,
		Name:         name,
		dependencies: 0,
		ProjectMap:   map[string]*Project{},
	}
}

func (project *Project) AddNeighbour(node *Project) {
	if _, ok := project.ProjectMap[node.Name]; !ok {
		project.Children = append(project.Children, node)
		project.ProjectMap[node.Name] = node
		node.IncrementDependencies()
	}
}

func (project *Project) GetName() string {
	return project.Name
}

func (project *Project) GetChildren() []*Project {
	return project.Children
}

func (project *Project) GetNumberDependencies() int {
	return project.dependencies
}

func (project *Project) IncrementDependencies() {
	project.dependencies++
}

func (project *Project) DecrementDependencies() {
	project.dependencies--
}

func BuildGraph(projects []string, dependencies [][]string) DependencyGraph {
	graph := DependencyGraph{
		Nodes:      []*Project{},
		ProjectMap: map[string]*Project{},
	}

	for _, project := range projects {
		graph.GetOrCreateNode(project)
	}

	for _, dependency := range dependencies {
		first := dependency[0]
		second := dependency[1]
		graph.AddEdge(first, second)
	}

	return graph
}

func FindBuildOrder(projects []string, dependencies [][]string) []*Project {
	graph := BuildGraph(projects, dependencies)

	return orderProjects(graph.Nodes)
}

func orderProjects(projects []*Project) []*Project {
	var order []*Project

	for _, project := range projects {
		if project.GetNumberDependencies() == 0 {
			order = append(order, project)
		}
	}

	toBeProcessed := 0

	for toBeProcessed < len(order) {
		// We have a circular dependency since there are  remaining projects with zero dependencies
		if toBeProcessed > len(order)-1 {
			panic(errors.New("circular dependency detected"))
		}

		// Remove project from list of dependencies
		current := order[toBeProcessed]
		children := current.GetChildren()

		for _, child := range children {
			child.DecrementDependencies()
		}

		// Add children that have to one depending on them
		for _, child := range children {
			if child.GetNumberDependencies() == 0 {
				order = append(order, child)
			}
		}

		toBeProcessed++
	}

	if len(order) != len(projects) {
		panic(errors.New("exiting due to circular dependency"))
	}

	return order
}
