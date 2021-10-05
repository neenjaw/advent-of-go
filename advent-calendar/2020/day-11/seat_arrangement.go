package main

type Space int

const (
	EmptySeat Space = iota + 1
	OccupiedSeat
	Floor
)

type Arrangement [][]Space

func ConvInputToSeatingArrangement(input []string) (arr Arrangement) {
	arr = make(Arrangement, len(input))

	for i, line := range input {
		arr[i] = make([]Space, 0, len(line))

		for _, r := range line {
			switch r {
			case '.':
				arr[i] = append(arr[i], Floor)
			case '#':
				arr[i] = append(arr[i], OccupiedSeat)
			case 'L':
				arr[i] = append(arr[i], EmptySeat)
			default:
				panic("unhandled space type")
			}
		}
	}

	return
}

type Coordinate struct {
	x int
	y int
}

func slopes() []Coordinate {
	return []Coordinate{{1, 1}, {1, 0}, {1, -1}, {0, 1}, {0, -1}, {-1, 1}, {-1, 0}, {-1, -1}}
}

func lookaroundSum(arr Arrangement, c Coordinate, limit bool) (sum int) {
	var fx func(Arrangement, Coordinate, Coordinate) bool
	if limit {
		fx = lookImmediate
	} else {
		fx = lookUnlimited
	}

	for _, slope := range slopes() {
		if fx(arr, slope, c) {
			sum += 1
		}
	}

	return
}

func lookImmediate(arr Arrangement, slope Coordinate, c Coordinate) bool {
	minY := 0
	minX := 0
	maxY := len(arr) - 1
	maxX := len(arr[0]) - 1
	c.x += slope.x
	c.y += slope.y

	if c.x < minX || c.x > maxX || c.y < minY || c.y > maxY {
		return false
	}

	if arr[c.y][c.x] == OccupiedSeat {
		return true
	}

	return false
}

func lookUnlimited(arr Arrangement, slope Coordinate, c Coordinate) bool {
	minY := 0
	minX := 0
	maxY := len(arr) - 1
	maxX := len(arr[0]) - 1

	for {
		c.x += slope.x
		c.y += slope.y

		if c.x < minX || c.x > maxX || c.y < minY || c.y > maxY {
			return false
		}

		if arr[c.y][c.x] == OccupiedSeat {
			return true
		}

		if arr[c.y][c.x] == EmptySeat {
			return false
		}
	}
}

type StepConfig struct {
	fillThreshold  int
	emptyThreshold int
	limit          bool
}

func Step(arr Arrangement, config StepConfig) (next Arrangement, changed bool) {
	next = make(Arrangement, len(arr))
	for y, row := range arr {
		next[y] = make([]Space, len(row))

		for x, spot := range row {
			if spot == Floor {
				next[y][x] = Floor
				continue
			}

			sum := lookaroundSum(arr, Coordinate{x, y}, config.limit)

			if spot == OccupiedSeat && sum >= config.emptyThreshold {
				next[y][x] = EmptySeat
				changed = true
				continue
			}

			if spot == EmptySeat && sum <= config.fillThreshold {
				next[y][x] = OccupiedSeat
				changed = true
				continue
			}

			next[y][x] = spot
		}
	}
	return
}

func CountOccupied(arr Arrangement) (count int) {
	for _, row := range arr {
		for _, spot := range row {
			if spot == OccupiedSeat {
				count += 1
			}
		}
	}
	return
}
