package main

import "fmt"

var nextId int

func main() {
	level := 20
	n := 0
	cur := 1
	for i := 0; i < level; i++ {
		n += cur
		cur *= 2
	}
	g := make([][]int, n)
	for i := 0; i < n; i++ {
		g[i] = make([]int, 0)
	}
	dfs(g, 0, level)
	fmt.Println(n, nextId)
	for i := 0; i < n; i++ {
		for _, u := range g[i] {
			fmt.Println(i+1, u+1)
		}
	}
}

func dfs(g [][]int, from int, level int) {
	if level == 1 {
		return
	}
	for i := 0; i < 2; i++ {
		nextId++
		g[from] = append(g[from], nextId)
		dfs(g, nextId, level-1)
	}
}
