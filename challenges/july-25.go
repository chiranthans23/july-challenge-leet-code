package challenges

func allPathsSourceTarget(graph [][]int) [][]int {

	var result [][]int
	visited := make([]bool, len(graph))
	dfsToTarget(graph, 0, []int{}, visited, &result)

	return result
}

func dfsToTarget(graph [][]int, source int, path []int, visited []bool, result *[][]int) (bool, []int) {

	// When I reach the destination I return the path
	if source == len(graph)-1 {
		path = append(path, source)
		return true, path
	}
	if visited[source] {
		return false, []int{}
	}
	path = append(path, source)
	visited[source] = true

	for _, node := range graph[source] {
		path = append(path, node)
		reachable, path := dfsToTarget(graph, node, path, visited, result)

		// Add paths list only if reachable
		if reachable {
			*result = append(*result, path)
		}

	}
	path = path[:len(path)-1]
	visited[source] = false

	return false, path
}
