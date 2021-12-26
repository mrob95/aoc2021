package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

const INP = 0
const ADD = 1
const MUL = 2
const DIV = 3
const MOD = 4
const EQL = 5

type Instruction struct {
	op, a, b int
	literal  bool
}

type ALU struct {
	regs [4]int
	ic   int
}

func runProgram(program *[]Instruction, in *[]int, z int) [4]int {
	var alu ALU
	alu.regs[3] = z
	for _, inst := range *program {
		b := 0
		if inst.literal {
			b = inst.b
		} else {
			b = alu.regs[inst.b]
		}

		switch inst.op {
		case INP:
			alu.regs[inst.a] = (*in)[alu.ic]
			alu.ic++
		case ADD:
			alu.regs[inst.a] = alu.regs[inst.a] + b
		case MUL:
			alu.regs[inst.a] = alu.regs[inst.a] * b
		case DIV:
			alu.regs[inst.a] = alu.regs[inst.a] / b
		case MOD:
			alu.regs[inst.a] = alu.regs[inst.a] % b
		case EQL:
			if alu.regs[inst.a] == b {
				alu.regs[inst.a] = 1
			} else {
				alu.regs[inst.a] = 0
			}
		}
	}
	return alu.regs
}

func parseB(l []byte) (int, bool) {
	bVal := l[6]
	if bVal >= 'w' && bVal <= 'z' {
		return int(bVal - 'w'), false
	}
	i, _ := strconv.Atoi(string(l[6:]))
	return i, true
}

func parseInput(r io.Reader) [][]Instruction {
	var subprograms [][]Instruction
	scanner := bufio.NewScanner(r)
	i := -1
	for scanner.Scan() {
		line := scanner.Bytes()
		switch line[0] {
		case 'i':
			subprograms = append(subprograms, []Instruction{})
			i++
			subprograms[i] = append(subprograms[i], Instruction{INP, int(line[4] - 'w'), 0, false})
		case 'a':
			b, literal := parseB(line)
			subprograms[i] = append(subprograms[i], Instruction{ADD, int(line[4] - 'w'), b, literal})
		case 'm':
			op := 0
			if line[1] == 'u' {
				op = MUL
			} else {
				op = MOD
			}
			b, literal := parseB(line)
			subprograms[i] = append(subprograms[i], Instruction{op, int(line[4] - 'w'), b, literal})
		case 'd':
			b, literal := parseB(line)
			subprograms[i] = append(subprograms[i], Instruction{DIV, int(line[4] - 'w'), b, literal})
		case 'e':
			b, literal := parseB(line)
			subprograms[i] = append(subprograms[i], Instruction{EQL, int(line[4] - 'w'), b, literal})
		default:
			panic("Unreachable")
		}
	}
	return subprograms
}

func gt(a, b []int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			return true
		} else if a[i] < b[i] {
			return false
		}
	}
	return false
}
func lt(a, b []int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] < b[i] {
			return true
		} else if a[i] < b[i] {
			return false
		}
	}
	return false
}

func main() {
	subprograms := parseInput(os.Stdin)

	var Zpaths [15]map[int][]int
	Zpaths[0] = map[int][]int{0: {}}
	for i, subprogram := range subprograms {
		Zpaths[i+1] = make(map[int][]int)
		for w := 1; w < 10; w++ {
			in := []int{w}
			for z, path := range Zpaths[i] {
				regs := runProgram(&subprogram, &in, z)
				newZ := regs[3]
				if newZ < 0 {
					continue // If we start a subprogram with a negative Z, we immediately do a negative % op, which is not allowed!
				}
				newPath := append([]int{}, path...)
				newPath = append(newPath, w)
				if p, ok := Zpaths[i+1][newZ]; ok {
					if lt(newPath, p) {
						Zpaths[i+1][newZ] = newPath
					}
				} else {
					Zpaths[i+1][newZ] = newPath
				}

			}
		}
		fmt.Println(len(Zpaths[i+1]))
	}
	p := Zpaths[14][0]
	fmt.Println(p)

	z := 0
	for i, subprogram := range subprograms {
		in := []int{p[i]}
		r := runProgram(&subprogram, &in, z)
		z = r[3]
		fmt.Println(p[:i+1], z)
	}
}
