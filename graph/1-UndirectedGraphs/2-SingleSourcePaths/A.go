package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve() {
	n := readInt()
	m := readInt()
	g := make([][]int, n)
	for i := 0; i < n; i++ {
		g[i] = make([]int, 0)
	}
	for i := 0; i < m; i++ {
		u := readInt() - 1
		v := readInt() - 1
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}
	s := readInt() - 1
	t := readInt()
	a := readIntArray(t)
	edgeTo := make([]int, n)
	vis := make([]bool, n)
	dfs(g, edgeTo, vis, s)
	for _, v := range a {
		v--
		if vis[v] {
			fmt.Println("YES")
			items := getPath(edgeTo, s, v)
			fmt.Println(len(items))
			for i, item := range items {
				if i > 0 {
					fmt.Print(" ")
				}
				fmt.Print(item)
			}
			fmt.Println()
		} else {
			fmt.Println("NO")
		}
	}
}

func getPath(edgeTo []int, s int, v int) []int {
	res := make([]int, 0)
	for s != v {
		res = append(res, v+1)
		v = edgeTo[v]
	}
	res = append(res, s+1)
	n := len(res)
	for i := 0; i < n/2; i++ {
		res[i], res[n-i-1] = res[n-i-1], res[i]
	}
	return res
}

func dfs(g [][]int, edgeTo []int, vis []bool, w int) {
	vis[w] = true
	for _, next := range g[w] {
		if !vis[next] {
			edgeTo[next] = w
			dfs(g, edgeTo, vis, next)
		}
	}
}

var scanner *bufio.Scanner

func main() {
	scanner = bufio.NewScanner(os.Stdin)
	const MaxTokenLength = 1024 * 1024
	scanner.Buffer(make([]byte, 0, MaxTokenLength), MaxTokenLength)
	scanner.Split(bufio.ScanWords)
	solve()
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a int64, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

// IO

func readString() string {
	scanner.Scan()
	return scanner.Text()
}

func readInt() int {
	val, _ := strconv.Atoi(readString())
	return val
}

func readInt64() int64 {
	v, _ := strconv.ParseInt(readString(), 10, 64)
	return v
}

func readIntArray(sz int) []int {
	a := make([]int, sz)
	for i := 0; i < sz; i++ {
		a[i] = readInt()
	}
	return a
}

func readInt64Array(n int) []int64 {
	res := make([]int64, n)
	for i := 0; i < n; i++ {
		res[i] = readInt64()
	}
	return res
}
