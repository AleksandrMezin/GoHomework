package main

import (
	"fmt"
	"math"
)

type Graph struct {
	adjacencyList map[int]map[int]float64
}

// Инициализация графа
func NewGraph() *Graph {
	return &Graph{
		adjacencyList: make(map[int]map[int]float64),
	}
}

// Добавление ребра в граф
func (g *Graph) AddEdge(src, dest int, weight float64) {
	if _, ok := g.adjacencyList[src]; !ok {
		g.adjacencyList[src] = make(map[int]float64)
	}
	g.adjacencyList[src][dest] = weight
}

// Алгоритм Дейкстры
func (g *Graph) Dijkstra(source int) map[int]float64 {
	distance := make(map[int]float64)
	visited := make(map[int]bool)

	for v := range g.adjacencyList {
		distance[v] = math.Inf(1)
		visited[v] = false
	}

	distance[source] = 0

	for count := 0; count < len(g.adjacencyList)-1; count++ {
		u := g.minDistance(distance, visited)
		visited[u] = true

		for v, weight := range g.adjacencyList[u] {
			if !visited[v] && distance[u]+weight < distance[v] {
				distance[v] = distance[u] + weight
			}
		}
	}

	return distance
}

// Нахождение вершины с минимальным расстоянием
func (g *Graph) minDistance(distance map[int]float64, visited map[int]bool) int {
	min := math.Inf(1)
	minIndex := -1

	for v, d := range distance {
		if !visited[v] && d < min {
			min = d
			minIndex = v
		}
	}

	return minIndex
}

func main() {
	graph := NewGraph()

	// Добавление ребер в граф
	graph.AddEdge(0, 1, 2)
	graph.AddEdge(0, 2, 4)
	graph.AddEdge(1, 2, 1)
	graph.AddEdge(1, 3, 7)
	graph.AddEdge(2, 4, 3)
	graph.AddEdge(3, 4, 2)
	graph.AddEdge(3, 5, 1)
	graph.AddEdge(4, 5, 5)

	source := 0
	distances := graph.Dijkstra(source)

	for v, d := range distances {
		fmt.Printf("Кратчайшее расстояние от вершины %d до вершины %d: %.2f\n", source, v, d)
	}
}
