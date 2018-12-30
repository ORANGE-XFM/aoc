package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Regs [6]int

type Instr struct {
	op string
	a,b,c int
}

func opcode(op string, A int, B int, C int, regs Regs) Regs {

	// load operands
	arg1, arg2 := 0, 0
	switch op {
	// reg,reg
	case "addr", "mulr", "banr", "borr", "gtrr", "eqrr":
		arg1, arg2 = regs[A], regs[B]
	// reg,imm (or only reg, ignore B)
	case "addi", "muli", "bani", "bori", "gtri", "eqri":
		arg1, arg2 = regs[A], B
	// imm,reg
	case "gtir", "eqir":
		arg1, arg2 = A, regs[B]
	// imm
	case "seti":
		arg1 = A
	// reg
	case "setr":
		arg1 = regs[A]
	}
	// execute
	v := 0
	switch op {
	case "addr", "addi":
		v = arg1 + arg2
	case "muli", "mulr":
		v = arg1 * arg2
	case "banr", "bani":
		v = arg1 & arg2
	case "borr", "bori":
		v = arg1 | arg2
	case "setr", "seti":
		v = arg1
	case "gtir", "gtri", "gtrr":
		if arg1 > arg2 {
			v = 1
		}
	case "eqir", "eqri", "eqrr":
		if arg1 == arg2 {
			v = 1
		}
	}

	outregs := regs
	outregs[C] = v
	return outregs
}

func readfile(filename string) (int,[]Instr) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error reading file")
		return 0,[]Instr{}
	}
	defer file.Close() 
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	ipmap, _ := strconv.Atoi(scanner.Text()[4:])

	prog := make ([]Instr,0)
	for scanner.Scan() {
		ins := strings.Split(scanner.Text()," ")
		a,_ := strconv.Atoi(ins[1])
		b,_ := strconv.Atoi(ins[2])
		c,_ := strconv.Atoi(ins[3])

		prog = append(prog, Instr{ins[0],a,b,c})
	}

	return ipmap, prog
}

func exec_prog(r0 int, ipmap int, program []Instr, max_instr int, breakpoint int, break_times int) int {
	var regs Regs
	regs[0]=r0
	instrs := 0
	seen_regs := []int{}
	for {
		ip := regs[ipmap]
		if ip == breakpoint {
			break_times -= 1
			fmt.Printf("r5:%24b   ",regs[5])
			fmt.Println("bktpoint hit at ",regs,"seen ",len(seen_regs),"instrs ",instrs)
			for _,reg := range seen_regs {
				if reg == regs[5] {
					fmt.Println("already seen : ",reg)
					return instrs
				}
			}
			seen_regs = append(seen_regs,regs[5])
		}
		if ip >= len(program) || instrs >= max_instr || break_times==0 {
			//fmt.Println("break at ", instrs, "ip=", ip)
			break
		}
		instr := program[ip]

		show := false
		if show {
			fmt.Print("ip=",ip,"   ")
			show_instr(instr)
			fmt.Print(" ",regs)
		}
		regs = opcode(instr.op,instr.a, instr.b, instr.c, regs)
		if show { 
			fmt.Println(" --->", regs)
		}
		regs[ipmap] += 1
		instrs += 1
	}
	return instrs
}

func show_instr(instr Instr) {
	fmt.Print("r",instr.c," = ",)
	switch instr.op {
	// reg,reg
	case "addr", "mulr", "banr", "borr", "gtrr", "eqrr":
		fmt.Print("r",instr.a," ",instr.op[:3]," r", instr.b)
	// reg,imm (or only reg, ignore B)
	case "addi", "muli", "bani", "bori", "gtri", "eqri":
		fmt.Print("r",instr.a," ",instr.op[:3]," ",instr.b)
	// imm,reg
	case "gtir", "eqir":
		fmt.Print(instr.op[:2]," ",instr.a," r",instr.b)
	// imm
	case "seti":
		fmt.Print(instr.a)
	// reg
	case "setr":
		fmt.Print("r",instr.a)
	}
}

func show_prog(program []Instr) {
	for ip,instr := range program {
		//fmt.Println(instr)
		
		fmt.Print(ip,"  - ")
		show_instr(instr)
		fmt.Println()
	}
}

func main() {
	ipmap, prog := readfile("2018_21.txt")
	show_prog(prog)
	r0 := 1
	maxi := 10000000000
	instrs := exec_prog(r0,ipmap,prog,maxi,30,99999999999) // bkptpoint
	fmt.Println("done in",instrs)
}
