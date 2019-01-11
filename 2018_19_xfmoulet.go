package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Instr struct {
	op      string
	a, b, c int
}
type Regs [6]int

type Machine struct {
	regs   [6]int
	ip_map int
	prog   []Instr
	ip     int
}

func opcode(instr Instr, regs Regs) Regs {
	A, B, C := instr.a, instr.b, instr.c

	// load operands
	arg1, arg2 := 0, 0
	switch instr.op {
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
	switch instr.op {
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

func (m *Machine) readfile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error reading file")
		return
	}
	// defer file.Close() -> not now !
	scanner := bufio.NewScanner(file)

	// read ip mapping
	if !scanner.Scan() {
		return
	}
	ip_map_str := scanner.Text()
	if !strings.HasPrefix(ip_map_str, "#ip ") {
		fmt.Println("Error : not an assignment ?")
	}
	ip_map, err := strconv.Atoi(ip_map_str[4:5])
	if err != nil {
		fmt.Println("cannot convert arguments in", ip_map_str)
		return
	}
	m.ip_map = ip_map

	//fmt.Println("IP mapped to register", ip_map)

	// read instructions
	var instr Instr
	for scanner.Scan() {
		instr_s := strings.Split(scanner.Text(), " ")
		instr.op = instr_s[0]
		instr.a, _ = strconv.Atoi(instr_s[1])
		instr.b, _ = strconv.Atoi(instr_s[2])
		instr.c, _ = strconv.Atoi(instr_s[3])

		//fmt.Println(instr)
		m.prog = append(m.prog, instr)
	}
}

func (m *Machine) execute() {
	n_steps := 0
	for m.regs[m.ip_map] < len(m.prog) {
		n_steps += 1

		instr := m.prog[m.regs[m.ip_map]]


		fmt.Println("ip=",m.regs[m.ip_map]," ",m.regs," ",instr)
		/* patch/fast forward [0 a=2 IP=3 b=2 max=10551275 scratch=0] -> [0 b=max-1 3 nb=max-1 10551275 scratch] so that a*b == max
		that will happen first when a=5 and b=10551275/5, just remove a bit to catch it
		*/
		
		if m.regs[2] == 3 {
			fmt.Println("was: ip=",m.regs[m.ip_map]," ",m.regs," ",instr)
			m.regs[3]=10551275/m.regs[1]
			fmt.Println("Skip ! now ip=",m.regs[m.ip_map]," ",m.regs," ",instr)
		}
		
		m.regs = opcode(instr, m.regs)

		//fmt.Println(m.regs)
		m.regs[m.ip_map] += 1
		if n_steps%1000000==0 { fmt.Println(n_steps)}
	}
	fmt.Println("final state:", m.regs, "after", n_steps, "steps")
}

func main() {
	var m Machine
	m.readfile(os.Args[1])
	m.regs[0] = 1 // part B
	m.execute()
}
/* real solution was
-  looking how the program behaves 
- decompiling the source and looking at it as

a=0
for (int i=1;i<BIGNUMBER;i++)
	for (int j=1;j<BIGNUMBER;i++)
	   if BIGNUMBER==i*j a+=i

in other words, take the sum of divisors of BIGNUMBER

so a quick python file did the trick ... see py file
*/