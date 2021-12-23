package main

import (
	"advent-of-go/util/file"
	"fmt"
)

type Limit struct {
	min, max int
}

type Run struct {
	file     string
	expected int64
	hasLimit bool
	limit    Limit
}

func main() {
	runs := []Run{
		{"example.txt", 39, true, Limit{-50, 50}},
		{"example2.txt", 590784, true, Limit{-50, 50}},
		{"input.txt", -1, true, Limit{-50, 50}},
	}

	runs2 := []Run{
		// {"example.txt", 39, false, Limit{}},
		// {"example2.txt", 590784, false, Limit{}},
		// {"input.txt", -1, false, Limit{}},
	}

	for i, run := range runs {
		do_main(i, run)
	}

	for i, run := range runs2 {
		do_main(i, run)
	}
}

func do_main(i int, run Run) {
	fileContent, _ := file.ReadFile(run.file)
	cubeSteps := ParseCuboidSteps(fileContent)

	// fmt.Println(cubeSteps)
	ans := simulate(cubeSteps, run.hasLimit, run.limit)
	if ans != run.expected {
		fmt.Printf("[FAIL] %d: expected: %v go: %v\n", i, run.expected, ans)
		return
	} else {
		fmt.Printf("[PASS] %d: expected: %v got: %v\n", i, run.expected, ans)
	}
}

type Position struct {
	z, y, x int
}

func simulate(steps []Rule, hasLimit bool, limit Limit) int64 {
	onCuboids := []Cuboid{}
	for i, step := range steps {
		fmt.Println("STEP >> ", i)
		nextOnCuboids := []Cuboid{}
		for _, onCuboid := range onCuboids {
			if step.cuboid.encompasses(onCuboid) {
				continue
			} else if step.cuboid.intersects(onCuboid) {
				resultingCuboids := onCuboid.subtract(step.cuboid)
				nextOnCuboids = append(nextOnCuboids, resultingCuboids...)
			} else {
				nextOnCuboids = append(nextOnCuboids, onCuboid)
			}
		}

		if step.state == On {
			nextOnCuboids = append(nextOnCuboids, step.cuboid)
		}
		onCuboids = nextOnCuboids

		// sum := int64(0)
		// for _, c := range onCuboids {
		// 	sum += c.volume()
		// }
		// fmt.Println("step sum ==============", sum)
	}

	var count int64 = 0
	for _, onCuboid := range onCuboids {
		if !hasLimit || !onCuboid.withinLimit(limit) {
			continue
		}
		count += onCuboid.volume()
	}

	return count
}

func (c1 Cuboid) encompasses(c2 Cuboid) bool {
	return c1.z1 <= c2.z1 && c1.z2 >= c2.z2 &&
		c1.y1 <= c2.y1 && c1.y2 >= c2.y2 &&
		c1.x1 <= c2.x1 && c1.x2 >= c2.x2
}

func (c1 Cuboid) intersects(c2 Cuboid) bool {
	return !(c1.x1 > c2.x2 || c1.x2 < c2.x1 ||
		c1.y1 > c2.y2 || c1.y2 < c2.y1 ||
		c1.z1 > c2.z2 || c1.z2 < c1.z2)
}

func (c1 Cuboid) subtract(c2 Cuboid) []Cuboid {
	if !c1.intersects(c2) {
		return []Cuboid{c1}
	}

	var xmin, xmax, ymin, ymax, zmin, zmax int

	if c1.x1 < c2.x1 {
		xmin = c1.x1
	} else {
		xmin = c2.x1
	}

	if c1.x2 > c2.x2 {
		xmax = c1.x2
	} else {
		xmax = c2.x2
	}

	if c1.y1 < c2.y1 {
		ymin = c1.y1
	} else {
		ymin = c2.y1
	}

	if c1.y2 > c2.y2 {
		ymax = c1.y2
	} else {
		ymax = c2.y2
	}

	if c1.z1 < c2.z1 {
		zmin = c1.z1
	} else {
		zmin = c2.z1
	}

	if c1.z2 > c2.z2 {
		zmax = c1.z2
	} else {
		zmax = c2.z2
	}

	for z := zmin; z <= zmax; z++ {
		for y := ymin; y <= ymax; y++ {
			for x := xmin; x <= xmax; x++ {
				if z < c1.z1 || z > c1.z2 || y < c1.y1 || y > c1.y2 || x < c1.x1 || x > c1.x2 {
					continue
				}

				c1.subtractions[Position{z: z, y: y, x: x}] = true
			}
		}
	}

	return []Cuboid{c1}
}

func (c Cuboid) volume() int64 {
	volume := int64((c.z2 - c.z1 + 1) * (c.y2 - c.y1 + 1) * (c.x2 - c.x1 + 1))
	return volume
}

func (c Cuboid) withinLimit(limit Limit) bool {
	return !(c.x1 > limit.max || c.x2 < limit.min ||
		c.y1 > limit.max || c.y2 < limit.min ||
		c.z1 > limit.max || c.z2 < limit.min)
}
