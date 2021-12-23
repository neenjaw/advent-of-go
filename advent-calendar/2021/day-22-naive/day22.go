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
	if int64(ans) != run.expected {
		fmt.Printf("[FAIL] %d: expected: %v go: %v\n", i, run.expected, ans)
		return
	} else {
		fmt.Printf("[PASS] %d: expected: %v got: %v\n", i, run.expected, ans)
	}
}

type Position struct {
	z, y, x int
}

func simulate(steps []Rule, hasLimit bool, limit Limit) int {
	if !hasLimit {
		panic("expect limit")
	}

	s := map[Position]State{}
	for _, step := range steps {
		for z := step.cuboid.z1; z <= step.cuboid.z2; z++ {
			if z < limit.min || z > limit.max {
				continue
			}
			for y := step.cuboid.y1; y <= step.cuboid.y2; y++ {
				if y < limit.min || y > limit.max {
					continue
				}
				for x := step.cuboid.x1; x <= step.cuboid.x2; x++ {
					if x < limit.min || x > limit.max {
						continue
					}
					s[Position{z: z, y: y, x: x}] = step.state
				}
			}
		}
	}

	count := 0
	for _, v := range s {
		if v == On {
			count += 1
		}
	}
	return count
}
