package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	op_addr = iota
	op_addi
	op_mulr
	op_muli
	op_banr
	op_bani
	op_borr
	op_bori
	op_setr
	op_seti
	op_gtir
	op_gtri
	op_gtrr
	op_eqir
	op_eqri
	op_eqrr
	NB_OPS
)

var opnames = [...]string{
	"addr", "addi",
	"mulr", "muli",
	"banr", "bani",
	"borr", "bori",
	"setr", "seti",
	"gtir", "gtri", "gtrr",
	"eqir", "eqri", "eqrr",
}

type Regs [4]int

func opcode(op int, A int, B int, C int, regs Regs) Regs {

	// load operands
	arg1, arg2 := 0, 0
	switch op {
	// reg,reg
	case op_addr, op_mulr, op_banr, op_borr, op_gtrr, op_eqrr:
		arg1, arg2 = regs[A], regs[B]
	// reg,imm (or only reg, ignore B)
	case op_addi, op_muli, op_bani, op_bori, op_gtri, op_eqri:
		arg1, arg2 = regs[A], B
	// imm,reg
	case op_gtir, op_eqir:
		arg1, arg2 = A, regs[B]
	// imm
	case op_seti:
		arg1 = A
	// reg
	case op_setr:
		arg1 = regs[A]
	}
	// execute
	v := 0
	switch op {
	case op_addr, op_addi:
		v = arg1 + arg2
	case op_muli, op_mulr:
		v = arg1 * arg2
	case op_banr, op_bani:
		v = arg1 & arg2
	case op_borr, op_bori:
		v = arg1 | arg2
	case op_setr, op_seti:
		v = arg1
	case op_gtir, op_gtri, op_gtrr:
		if arg1 > arg2 {
			v = 1
		}
	case op_eqir, op_eqri, op_eqrr:
		if arg1 == arg2 {
			v = 1
		}
	}

	outregs := regs
	outregs[C] = v
	return outregs
}

func openfile(filename string) *bufio.Scanner {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error reading file")
		return nil
	}
	// defer file.Close() -> not now !
	scanner := bufio.NewScanner(file)
	return scanner
}

func readQuad(str string, sep string) Regs {
	s := strings.Split(str, sep)
	regs := Regs{0, 0, 0, 0}
	for i := 0; i < 4; i++ {
		regs[i], _ = strconv.Atoi(s[i])
	}
	return regs
}

func print_map(mat [16][16]bool) {
	// display
	fmt.Println("opcodes ===")
	for i := 0; i < 16; i++ {
		fmt.Print(i, ":")
		for j := 0; j < 16; j++ {
			s := " ."
			if mat[i][j] {
				s = " x"
			}
			fmt.Print(s)

		}
		fmt.Print("\n")
	}
}

func solve(mat [16][16]bool) [16]int {
	// while not all solved
	tosolve := 16
	known := [16]int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}
	k := 0
	// display
	fmt.Println("opcodes ===")

	for tosolve > 0 {
		k = (k + 1) % 16 // next, possibly loop
		if known[k] >= 0 {
			continue
		} // already solved

		// solved one ?
		nb_ok := 0
		op_ok := -1
		for i := 0; i < 16; i++ {
			if mat[k][i] {
				nb_ok += 1
				op_ok = i
			}
		}

		// yes, mark as known
		if nb_ok == 1 {
			tosolve -= 1
			known[k] = op_ok
			fmt.Println("solved", k, "->", opnames[op_ok], op_ok)
			for j := 0; j < 16; j++ {
				if j != k {
					mat[j][op_ok] = false // now others are now not possible
				}
			}
			print_map(mat)
		}
	}
	return known
}

func check_multiple(scanner *bufio.Scanner) [16][16]bool {

	var possible_ops [16][16]bool // opcode id -> possible opcodes.
	for i := 0; i < 16; i++ {
		for j := 0; j < 16; j++ {
			possible_ops[i][j] = true
		}
	}

	nb_multiples := 0
	for {
		if !scanner.Scan() {
			break
		}
		reg_in := readQuad(scanner.Text()[9:19], ", ")
		//fmt.Println("reg_in", reg_in)

		scanner.Scan()
		instr := readQuad(scanner.Text(), " ")
		//fmt.Println("instr", instr)

		scanner.Scan()
		reg_out := readQuad(scanner.Text()[9:19], ", ")
		//fmt.Println("reg_out", reg_out)

		scanner.Scan() // empty line

		// check if multiple possibles
		nbok := 0
		opid, a, b, c := instr[0], instr[1], instr[2], instr[3]
		for op := op_addr; op < NB_OPS; op++ {
			if opcode(op, a, b, c, reg_in) == reg_out {
				//fmt.Println("    op", opnames[op], "matches")
				nbok += 1
			} else {
				// make it not match
				possible_ops[opid][op] = false
			}
		}
		if nbok >= 3 {
			nb_multiples += 1
			//fmt.Println("multiples",nb_multiples)
		}
	}
	fmt.Println("nb multiples", nb_multiples)
	return possible_ops
}

func exec_prog(filename string, op_map [16]int) {
	var regs Regs
	scanner := openfile(filename)
	for scanner.Scan() {
		q := readQuad(scanner.Text(), " ")
		op, a, b, c := q[0], q[1], q[2], q[3]
		regs = opcode(op_map[op], a, b, c, regs)
		fmt.Println("op", op, "->", op_map[op], a, b, c, "regs", regs)
	}
}

func main() {
	//scan := openfile("2018_16_exinstrs.txt")
	scan := openfile("2018_16_instrs.data")
	mat := check_multiple(scan)
	print_map(mat)
	op_map := solve(mat)
	fmt.Println(op_map)
	exec_prog("2018_16_prog.data", op_map)
}
