package sto

type Point struct {
	x int
	y int
}

func movingCount(m int, n int, k int) int {
	startPoint := Point{0, 0}
	l := []Point{startPoint}
	pointMap := make(map[Point]int)
	for len(l) > 0 {
		p := l[0]
		l = l[1:]
		if  pointMap[p] != 1 && p.x < m && p.y < n && (getSum(p.x) + getSum(p.y)) <= k {
			pointMap[p] = 1
			l = append(l, Point{p.x+1, p.y})
			l = append(l, Point{p.x, p.y+1})
		}
	}
	return len(pointMap)
}

func getSum(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n = n / 10
	}
	return sum
}
