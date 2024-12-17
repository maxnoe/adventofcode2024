package day17

import (
	"fmt"
	"log"
	"math"
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
	Output []int
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
	m.Output = append(m.Output, m.combo() % 8)
	m.Ptr += 2
}

func (m *Machine) bdv() {
	m.B = m.dv_impl()
}

func (m *Machine) cdv() {
	m.C = m.dv_impl()
}

func (m *Machine) RunProgram() {
	for m.Ptr < len(m.Program) {
		instruction := Instruction(m.Program[m.Ptr])
		switch instruction {
		case ADV: m.adv()
		case BXL: m.bxl()
		case BST: m.bst()
		case JNZ: m.jnz()
		case BXC: m.bxc()
		case OUT: m.out()
		case BDV: m.bdv()
		case CDV: m.cdv()
		}
	}
}

func Parse(input string) (Machine, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) != 5 {
		return Machine{}, fmt.Errorf("Invalid format of input, expected 5 lines got %d", len(lines))
	}
	
	A, err := strconv.Atoi(strings.Split(lines[0], ": ")[1])
	if err != nil {return Machine{}, err}
	B, err := strconv.Atoi(strings.Split(lines[1], ": ")[1])
	if err != nil {return Machine{}, err}
	C, err := strconv.Atoi(strings.Split(lines[2], ": ")[1])
	if err != nil {return Machine{}, err}

	values := strings.Split(strings.Split(lines[4], ": ")[1], ",")
	program := make([]int, len(values))
	for i, val := range values {
		val, err := strconv.Atoi(val)
		if err != nil {return Machine{}, err}
		program[i] = val
	}

	return Machine{A, B, C, program, 0, nil}, nil
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
	for {
		if A % (1024 * 1024) == 0 {
			log.Printf("%d, log2(A) = %f \n", A, math.Log2(float64(A)))
		}
		new_m := Machine{A, m.B, m.C, m.Program, 0, nil}
		new_m.RunProgram()

		if slices.Equal(new_m.Output, new_m.Program) {
			return A, nil
		}
		A += 1
	}
}

func init() {
	aoc24.AddSolution(17, Parse, Part1, Part2)
}
