package day17

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/maxnoe/adventofcode2024/aoc24"
)

type Instruction int

const (
	ADV Instruction = 0
	BXL             = 1
	BST             = 2
	JNZ             = 3
	BXC             = 4
	OUT             = 5
	BDV             = 6
	CDV             = 7
)

type Machine struct {
	A       int
	B       int
	C       int
	Program []int
	Ptr     int
	Output  []int
}

func (m *Machine) arg() int {
	return m.Program[m.Ptr+1]
}

func (m *Machine) combo() int {
	arg := m.arg()
	switch arg {
	case 0, 1, 2, 3:
		return arg
	case 4:
		return m.A
	case 5:
		return m.B
	case 6:
		return m.C
	default:
		log.Panicf("Invalid argument value: %v", arg)
	}
	return 0
}

func (m *Machine) dv_impl() int {
	result := m.A / (1 << m.combo())
	m.Ptr += 2
	return result
}

func (m *Machine) adv() {
	m.A = m.dv_impl()
}

func (m *Machine) bxl() {
	m.B = m.B ^ m.arg()
	m.Ptr += 2
}

func (m *Machine) bst() {
	m.B = m.combo() % 8
	m.Ptr += 2
}

func (m *Machine) jnz() {
	if m.A != 0 {
		m.Ptr = m.arg()
	} else {
		m.Ptr += 2
	}
}

func (m *Machine) bxc() {
	m.B = m.B ^ m.C
	m.Ptr += 2
}

func (m *Machine) out() {
	m.Output = append(m.Output, m.combo()%8)
	m.Ptr += 2
}

func (m *Machine) bdv() {
	m.B = m.dv_impl()
}

func (m *Machine) cdv() {
	m.C = m.dv_impl()
}

func (m *Machine) PrintInstruction(ptr int) {
	instruction := Instruction(m.Program[ptr])
	arg := m.Program[ptr+1]

	var combo string
	if arg <= 3 {
		combo = strconv.Itoa(arg)
	} else if arg >= 4 && arg <= 6 {
		combo = []string{"A", "B", "C"}[arg-4]
	} else {
		combo = "invalid"
	}

	switch instruction {
	case ADV:
		log.Printf("adv %d: A = A / (2 ** %s) ", arg, combo)
	case BXL:
		log.Printf("bxl %d: B = B ^ %d", arg, arg)
	case BST:
		log.Printf("bst %d: B = %s %% 8", arg, combo)
	case JNZ:
		log.Printf("jnz %d", arg)
	case BXC:
		log.Printf("bxc  : B = B ^ C")
	case OUT:
		log.Printf("out %d: %s %% 8", arg, combo)
	case BDV:
		log.Printf("bdv %d: B = A / (2 ** %s) ", arg, combo)
	case CDV:
		log.Printf("cdv %d: C = A / (2 ** %s) ", arg, combo)
	}
}

func (m *Machine) PrintProgram() {
	log.Printf("%v", m)
	for i := range len(m.Program) / 2 {
		ptr := 2 * i
		m.PrintInstruction(ptr)
	}
}

func (m *Machine) RunProgram() {
	for m.Ptr < len(m.Program) {
		// log.Printf("Ptr=%2d, A=%8d B=%8d C=%8d, output=%v", m.Ptr, m.A, m.B, m.C, m.Output)
		// m.PrintInstruction(m.Ptr)
		instruction := Instruction(m.Program[m.Ptr])
		switch instruction {
		case ADV:
			m.adv()
		case BXL:
			m.bxl()
		case BST:
			m.bst()
		case JNZ:
			m.jnz()
		case BXC:
			m.bxc()
		case OUT:
			m.out()
		case BDV:
			m.bdv()
		case CDV:
			m.cdv()
		}
	}
}

func Parse(input string) (Machine, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) != 5 {
		return Machine{}, fmt.Errorf("Invalid format of input, expected 5 lines got %d", len(lines))
	}

	A, err := strconv.Atoi(strings.Split(lines[0], ": ")[1])
	if err != nil {
		return Machine{}, err
	}
	B, err := strconv.Atoi(strings.Split(lines[1], ": ")[1])
	if err != nil {
		return Machine{}, err
	}
	C, err := strconv.Atoi(strings.Split(lines[2], ": ")[1])
	if err != nil {
		return Machine{}, err
	}

	values := strings.Split(strings.Split(lines[4], ": ")[1], ",")
	program := make([]int, len(values))
	for i, val := range values {
		val, err := strconv.Atoi(val)
		if err != nil {
			return Machine{}, err
		}
		program[i] = val
	}

	m := Machine{A, B, C, program, 0, nil}
	m.PrintProgram()
	return m, nil
}

func JoinInts(values []int) string {
	if len(values) == 0 {
		return ""
	}
	s := strconv.Itoa(values[0])
	if len(values) == 1 {
		return s
	}

	for _, v := range values[1:] {
		s += ","
		s += strconv.Itoa(v)
	}
	return s
}

func Part1(m Machine) (string, error) {
	m.RunProgram()
	return JoinInts(m.Output), nil
}

func Part2(m Machine) (int, error) {
	A := 0
	previous := 0
	for i := range m.Program {

		// we go from the last output to the first
		// looking for the first number that outputs the partial program
		// starting from that index to the end
		pos := len(m.Program) - i - 1

		// The program always seems to divide A by 8, so we multiply by 8 to start the next scan
		A = previous * 8
		for {
			t := Machine{A, m.B, m.C, m.Program, 0, nil}
			t.RunProgram()

			if slices.Equal(t.Output, m.Program[pos:]) {
				previous = A
				break
			}
			A += 1
		}
	}

	return A, nil
}

func init() {
	aoc24.AddSolution(17, Parse, Part1, Part2)
}
