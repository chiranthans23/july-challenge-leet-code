package challenges

func makePreRequisiteGraph(numCourses int, prerequisites [][]int) ([][]int, []int) {
	adjTable := make([][]int, numCourses)
	in := make([]int, numCourses)

	for _, adge := range prerequisites {
		adjTable[adge[1]] = append(adjTable[adge[1]], adge[0])
		in[adge[0]]++
	}
	return adjTable, in
}
func findOrder(numCourses int, prerequisites [][]int) []int {
	ret := []int{}
	if len(prerequisites) == 0 {
		for i := 0; i < numCourses; i++ {
			ret = append(ret, i)
		}
		return ret
	}
	adjTable, in := makePreRequisiteGraph(numCourses, prerequisites)

	que := []int{}
	for course, degree := range in {
		if degree == 0 {
			que = append(que, course)
		}
	}

	for len(que) > 0 {
		front := que[0]
		que = que[1:]
		ret = append(ret, front)
		for _, node := range adjTable[front] {
			in[node]--
			if in[node] == 0 {
				que = append(que, node)
			}
		}
	}

	for i := 0; i < numCourses; i++ {
		if in[i] != 0 {
			return []int{}
		}
	}
	return ret
}
