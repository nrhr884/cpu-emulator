package main

import (
	"fmt"
	"ope"
	. "ope"
)

func execute(rom *[256]uint16) {
	var reg [16]uint16
	var ram [256]uint16

	var pc uint16 = 0
	var ir uint16 = rom[pc]
	var flagEq = false

	ram[0] = 1

	for GetOpCode(ir) != ope.HLT && pc < 255 {
		fmt.Printf("pc=%-5d ir=%-5d reg0=%-5d reg1=%-5d reg2=%-5d reg3=%-5d\n",
			pc, ir, reg[REG0], reg[REG1], reg[REG2], reg[REG3])
		switch GetOpCode(ir) {
		case MOV:
			reg[GetRegA(ir)] = reg[GetRegB(ir)]
		case ADD:
			reg[GetRegA(ir)] += reg[GetRegB(ir)]
		case SUB:
			reg[GetRegA(ir)] -= reg[GetRegB(ir)]
		case AND:
			reg[GetRegA(ir)] &= reg[GetRegB(ir)]
		case OR:
			reg[GetRegA(ir)] |= reg[GetRegB(ir)]
		case SL:
			reg[GetRegA(ir)] <<= 1
		case SR:
			reg[GetRegA(ir)] >>= 1
		case SRA:
			reg[GetRegA(ir)] = (reg[GetRegA(ir)] & 0x8000) | (reg[GetRegA(ir)] >> 1)
		case LDL:
			reg[GetRegA(ir)] = (reg[GetRegA(ir)] & 0xff00) | (GetData(ir) & 0x00ff)
		case LDH:
			reg[GetRegA(ir)] = (reg[GetRegA(ir)] & 0x00ff) | (GetData(ir) << 8)
		case CMP:
			flagEq = (reg[GetRegA(ir)] == reg[GetRegB(ir)])
		case JE:
			if flagEq {
				pc = GetAddr(ir)
				ir = rom[pc]
				continue
			}
		case JMP:
			pc = GetAddr(ir)
			ir = rom[pc]
			continue
		case LD:
			reg[GetRegA(ir)] = ram[GetAddr(ir)]
		case ST:
			ram[GetAddr(ir)] = reg[GetRegA(ir)]
		}

		pc = pc + 1
		ir = rom[pc]
	}
}

func makeAssembler(rom *[256]uint16) {
	irs := []uint16{
		MakeLdh(REG0, 0),
		MakeLdl(REG0, 0),
		MakeLdh(REG1, 0),
		MakeLdl(REG1, 1),
		MakeLdh(REG2, 0),
		MakeLdl(REG2, 0),
		MakeLdh(REG3, 0),
		MakeLdl(REG3, 10),
		MakeAdd(REG2, REG1),
		MakeAdd(REG0, REG2),
		MakeSt(REG0, 64),
		MakeCmp(REG2, REG3),
		MakeJe(14),
		MakeJmp(8),
		MakeHlt(),
	}

	for i, ir := range irs {
		rom[i] = ir
	}
}

func main() {
	var rom [256]uint16
	makeAssembler(&rom)
	execute(&rom)
}
