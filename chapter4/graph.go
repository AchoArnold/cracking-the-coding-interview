package chapter4

type Graph struct {
	Nodes []GraphNode
}

type GraphNode struct {
	Value    interface{}
	Children []interface{}
}

func StringToInterfaceSlice(array []string) []interface{} {
	result := make([]interface{}, len(array))
	for index, value := range array {
		result[index] = value
	}

	return result
}

func IntToInterfaceSlice(array []int) []interface{} {
	result := make([]interface{}, len(array))
	for index, value := range array {
		result[index] = value
	}
	return result
}
