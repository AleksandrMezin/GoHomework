package main

import "fmt"

type Graph struct {
	adjacencyList map[int][]int
}

// NewGraph Инициализация графа
func NewGraph() *Graph {
	return &Graph{
		adjacencyList: make(map[int][]int),
	}
}

// AddEdge Добавление ребра в граф
func (g *Graph) AddEdge(left, right int) {
	g.adjacencyList[left] = append(g.adjacencyList[left], right)
	g.adjacencyList[right] = append(g.adjacencyList[right], left) // добавляем обратное ребро
}

// BFS Поиск в ширину
func (g *Graph) BFS(start int) {
	visited := make(map[int]bool)
	queue := []int{start}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if visited[node] {
			continue
		}

		visited[node] = true
		fmt.Printf("%d ", node)

		neighbors := g.adjacencyList[node]
		for _, neighbor := range neighbors {
			if !visited[neighbor] {
				queue = append(queue, neighbor)
			}
		}
	}
}

func main() {
	graph := NewGraph()

	// Добавление ребер в граф
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(1, 2)
	graph.AddEdge(2, 3)
	graph.AddEdge(3, 4)
	graph.AddEdge(4, 2)

	startNode := 0
	fmt.Printf("Поиск в ширину (BFS) начиная с вершины %d: ", startNode)
	graph.BFS(startNode)
}
