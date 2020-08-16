package ope

import "testing"

func TestMakeFunctions(t *testing.T) {
	tests := []struct {
		ir      uint16
		opeCode uint16
	}{
		{MakeMov(0, 0), MOV},
		{MakeAdd(0, 0), ADD},
		{MakeSub(0, 0), SUB},
		{MakeAnd(0, 0), AND},
		{MakeOr(0, 0), OR},
		{MakeSl(0), SL},
		{MakeSr(0), SR},
		{MakeSra(0), SRA},
		{MakeLdl(0, 0), LDL},
		{MakeLdh(0, 0), LDH},
		{MakeCmp(0, 0), CMP},
		{MakeJe(0), JE},
		{MakeJmp(0), JMP},
		{MakeLd(0, 0), LD},
		{MakeSt(0, 0), ST},
		{MakeHlt(), HLT},
	}

	for _, tt := range tests {
		if (tt.ir >> 11) != tt.opeCode {
			t.Errorf("OpCode error actual=%d expected=%d", tt.ir>>11, tt.opeCode)
		}
	}
}
