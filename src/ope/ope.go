package ope

// 命令コードニモニック

const (
	MOV = iota
	ADD
	SUB
	AND
	OR
	SL
	SR
	SRA
	LDL
	LDH
	CMP
	JE
	JMP
	LD
	ST
	HLT
)

const (
	REG0 = iota
	REG1
	REG2
	REG3
	REG4
	REG5
	REG6
	REG7
)

// 命令コード生成関数

func makeOp(opCode, ra, rb uint16) uint16 {
	return opCode<<11 | (ra << 8) | (rb << 5)
}

func makeData(val uint16) uint16 {
	return val & 0x00ff
}

func makeAddr(addr uint16) uint16 {
	return addr & 0x00ff
}

func MakeMov(ra, rb uint16) uint16 {
	return makeOp(MOV, ra, rb)
}

func MakeAdd(ra, rb uint16) uint16 {
	return makeOp(ADD, ra, rb)
}

func MakeSub(ra, rb uint16) uint16 {
	return makeOp(SUB, ra, rb)
}

func MakeAnd(ra, rb uint16) uint16 {
	return makeOp(AND, ra, rb)
}

func MakeOr(ra, rb uint16) uint16 {
	return makeOp(OR, ra, rb)
}

func MakeSl(ra uint16) uint16 {
	return makeOp(SL, ra, 0)
}

func MakeSr(ra uint16) uint16 {
	return makeOp(SR, ra, 0)
}

func MakeSra(ra uint16) uint16 {
	return makeOp(SRA, ra, 0)
}

func MakeLdl(ra, ival uint16) uint16 {
	return makeOp(LDL, ra, 0) | makeData(ival)
}

func MakeLdh(ra, ival uint16) uint16 {
	return makeOp(LDH, ra, 0) | makeData(ival)
}

func MakeCmp(ra, rb uint16) uint16 {
	return makeOp(CMP, ra, rb)
}

func MakeJe(addr uint16) uint16 {
	return makeOp(JE, 0, 0) | makeAddr(addr)
}

func MakeJmp(addr uint16) uint16 {
	return makeOp(JMP, 0, 0) | makeAddr(addr)
}

func MakeLd(ra, addr uint16) uint16 {
	return makeOp(LD, ra, 0) | makeAddr(addr)
}

func MakeSt(ra, addr uint16) uint16 {
	return makeOp(ST, ra, 0) | makeAddr(addr)
}

func MakeHlt() uint16 {
	return makeOp(HLT, 0, 0)
}

// 命令コード分離関数

func GetOpCode(ir uint16) uint16 {
	return (ir >> 11)
}

func GetRegA(ir uint16) uint16 {
	return (ir >> 8) & 0x007
}

func GetRegB(ir uint16) uint16 {
	return (ir >> 5) & 0x007
}

func GetData(ir uint16) uint16 {
	return ir & 0x00ff
}

func GetAddr(ir uint16) uint16 {
	return ir & 0x00ff
}
