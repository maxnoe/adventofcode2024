package day21


type Direction int8
const (
	dirUp Direction = iota
	dirDown
	dirLeft
	dirRight
)

type Button int8
const (
	b0 Button = iota
	b1
	b2
	b3
	b4
	b5
	b6
	b7
	b8
	b9
	bA
	bUp
	bDown
	bLeft
	bRight
)


func inverse(dir Direction) Direction {
	switch (dir) {
	case dirUp: return dirDown
	case dirDown: return dirUp
	case dirLeft: return dirRight
	case dirRight: return dirLeft
	}
	panic("Invalid Direction")
}

type PadType int8
const (
	NumPad PadType = iota
	DirPad 
)

type ButtonPair struct {
	From Button
	To Button
}

type Connections map[ButtonPair][]Direction
type Neighbor struct {
	B Button
	D Direction
}
type Neighbors map[Button][]Neighbor

type Pad struct {
	Buttons Neighbors
	Connections Connections
}

var Pads = map[PadType]Pad {
	NumPad: {
		Neighbors{
			bA: {{b0, dirLeft}, {b3, dirUp}},
			b0: {{b2, dirUp}},

			b3: {{b2, dirLeft}, {b6, dirUp}},
			b2: {{b1, dirLeft}, {b5, dirUp}},
			b1: {{b4, dirUp}},

			b6: {{b5, dirLeft}, {b9, dirUp}},
			b5: {{b4, dirLeft}, {b8, dirUp}},
			b4: {{b7, dirUp}},
		},
		Connections{},
	},
	DirPad: {
		Neighbors{
			bA: {{bUp, dirLeft}, {bRight, dirDown}},
			bUp: {{bDown, dirDown}},
			bRight: {{bDown, dirLeft}},
			bDown: {{bLeft, dirLeft}},
		},
		Connections{},
	},
}

func fillInverseNeighbors(pad Pad) {
	for button, neighbors := range pad.Buttons {
		for _, neighbor := range neighbors {
			inv := Neighbor{button, inverse(neighbor.D)}
			pad.Buttons[neighbor.B] = append(pad.Buttons[neighbor.B], inv)
		}
	}
}

func invertSequence(sequence []Direction) []Direction {
	inv := make([]Direction, len(sequence))
	return inv
}


func ShortestConnection(pad Pad, from Button, to Button) []Direction {
	key := ButtonPair{from, to}
	sequence, found := pad.Connections[key]
	if found {
		return sequence
	}

	inverse_sequence, found := pad.Connections[ButtonPair{to, from}]
	if found {
		pad.Connections[key] = invertSequence(inverse_sequence)
		return pad.Connections[key]
	}

	to_check := [][]Direction{{}}
	path := make([]Direction, 0)
	for len(to_check) > 0 {
	}

	pad.Connections[key] = path
	return path
}

func init () {
	fillInverseNeighbors(Pads[NumPad])
}
