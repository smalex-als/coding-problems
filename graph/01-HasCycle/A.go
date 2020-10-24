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
	if hasCycle(n, g) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func hasCycle(n int, g [][]int) bool {
	vis := make([]bool, n)
	for i := 0; i < n; i++ {
		if !vis[i] && dfs(g, vis, i, i) {
			return true
		}
	}
	return false
}

func dfs(g [][]int, vis []bool, u int, parent int) bool {
	vis[u] = true
	for _, next := range g[u] {
		if !vis[next] {
			if dfs(g, vis, next, u) {
				return true
			}
		} else if next != parent {
			return true
		}
	}
	return false
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
