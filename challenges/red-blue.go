package challenges

func makeGraph(n int, red_edges [][]int, blue_edges [][]int) ([][]int, [][]int) {
	graphsRed := make([][]int, n)
	graphsBlue := make([][]int, n)
	for _, e := range red_edges {
		graphsRed[e[0]] = append(graphsRed[e[0]], e[1])
	}
	for _, e := range blue_edges {
		graphsBlue[e[0]] = append(graphsBlue[e[0]], e[1])
	}
	return graphsRed, graphsBlue
}

func shortestAlternatingPaths(n int, red_edges [][]int, blue_edges [][]int) []int {
	graphsRed, graphsBlue := makeGraph(n, red_edges, blue_edges)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = -1
	}

	visited := make([][]bool, n)
	for j := 0; j < n; j++ {
		visited[j] = []bool{false, false}
	}

	curr := make([][]int, 1)
	curr[0] = []int{0, -1}
	level := 0
	for len(curr) != 0 {
		next := [][]int{}
		for _, c := range curr {
			color := c[1]
			node := c[0]
			if color != -1 {
				visited[node][color] = true
			}
			if res[node] == -1 || res[node] > level {
				res[node] = level
			}
			if color == 0 || color == -1 {
				for _, neighbor := range graphsRed[node] {
					if !visited[neighbor][1] {
						next = append(next, []int{neighbor, 1})
					}
				}
			}
			if color == 1 || color == -1 {
				for _, neighbor := range graphsBlue[node] {
					if !visited[neighbor][0] {
						next = append(next, []int{neighbor, 0})
					}
				}
			}
		}
		curr = next
		level++
	}
	return res
}
