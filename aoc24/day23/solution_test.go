package day23

import (
	"testing"

	"github.com/maxnoe/adventofcode2024/aoc24"
)


var test_input = `
kh-tc
qp-kh
de-cg
ka-co
yn-aq
qp-ub
cg-tb
vc-aq
tb-ka
wh-tc
yn-cg
kh-ub
ta-co
de-co
tc-td
tb-wq
wh-td
ta-ka
td-qp
aq-cg
wq-ub
ub-vc
de-ta
wq-aq
wq-vc
wh-yn
ka-de
kh-ta
co-tc
wh-qp
tb-vc
td-yn
`

func TestPart1(t *testing.T) {
	conns, err := Parse(test_input)
	aoc24.AssertEqual(t, err, nil)

	answer, err := Part1(conns)
	aoc24.AssertEqual(t, err, nil)
	aoc24.AssertEqual(t, answer, 7)
}
