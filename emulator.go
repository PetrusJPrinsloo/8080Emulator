package main

import (
	"fmt"
	"log"
)

type ConditionCodes struct {
	Z   bool  // Zero
	S   bool  // Sign
	P   bool  // Parity
	CY  bool  // Carry
	AC  bool  // Auxiliary Carry
	Pad uint8 // Padding
}

type State8080 struct {
	A         uint8
	B         uint8
	C         uint8
	D         uint8
	E         uint8
	H         uint8
	L         uint8
	SP        uint16
	PC        uint16
	Memory    []byte
	Cc        ConditionCodes
	IntEnable uint8
	Quit      chan struct{}
}

func NewState8080(rom []byte, quit chan struct{}) *State8080 {
	state := State8080{}
	state.Memory = make([]byte, 0x10000)
	copy(state.Memory[0x0000:], rom)
	state.SP = 0x0000
	state.PC = 0x0000
	state.IntEnable = 0
	state.Quit = quit
	return &state
}

func (state *State8080) Step() error {
	return Emulate8080Op(state)
}

func UnimplementedInstruction(state *State8080) {
	state.PC++
	log.Fatalln("Error: Unimplemented instruction")
}

func parity(b uint8, numBits uint8) bool {
	var i uint8
	var parity bool
	for i = 0; i < numBits; i++ {
		if (b & (1 << i)) != 0 {
			parity = !parity
		}
	}
	return parity
}

func Emulate8080Op(state *State8080) error {
	switch state.Memory[state.PC] {

	// NOP
	case 0x00:
		break

	// LXI B, D16
	case 0x01:
		state.B = state.Memory[state.PC+2]
		state.C = state.Memory[state.PC+1]
		state.PC += 3
		break

	// STAX B
	case 0x02:
		state.Memory[(uint16(state.B)<<8)|uint16(state.C)] = state.A
		state.PC++
		break

	// INX B
	case 0x03:
		state.B++
		state.C++
		state.PC++
		break

	// INR B
	case 0x04:
		state.B++
		state.PC++
		break

	// DCR B
	case 0x05:
		state.B--
		state.PC++
		break

	// MVI B, D8
	case 0x06:
		UnimplementedInstruction(state)
		break

	// RLC
	case 0x07:
		UnimplementedInstruction(state)
		break
	// NOP
	case 0x08:
		break
	// DAD B
	case 0x09:
		UnimplementedInstruction(state)
		break
	// LDAX B
	case 0x0a:
		UnimplementedInstruction(state)
		break
	// DCX B
	case 0x0b:
		UnimplementedInstruction(state)
		break
	// INR C
	case 0x0c:
		state.C++
		state.PC++
		break
	// DCR C
	case 0x0d:
		state.C--
		state.PC++
		break
	// MVI C, D8
	case 0x0e:
		state.Memory[state.PC+1] = state.C
		state.PC += 2
		break
	// RRC
	case 0x0f:
		UnimplementedInstruction(state)
		break
	// NOP
	case 0x10:
		break
	// LXI D, D16
	case 0x11:
		UnimplementedInstruction(state)
		break
	// STAX D
	case 0x12:
		UnimplementedInstruction(state)
		break
	// INX D
	case 0x13:
		UnimplementedInstruction(state)
		break
	// INR D
	case 0x14:
		state.D++
		state.PC++
		break
	// DCR D
	case 0x15:
		state.D--
		state.PC++
		break
	// MVI D, D8
	case 0x16:
		UnimplementedInstruction(state)
		break
	// RAL
	case 0x17:
		UnimplementedInstruction(state)
		break
	// NOP
	case 0x18:
		state.PC++
		break
	// DAD D
	case 0x19:
		UnimplementedInstruction(state)
		break
	// LDAX D
	case 0x1a:
		UnimplementedInstruction(state)
		break
	// DCX D
	case 0x1b:
		UnimplementedInstruction(state)
		break
	// INR E
	case 0x1c:
		state.E++
		state.PC++
		break
	// DCR E
	case 0x1d:
		state.E--
		state.PC++
		break
	// MVI E, D8
	case 0x1e:
		UnimplementedInstruction(state)
		break
	// RAR
	case 0x1f:
		UnimplementedInstruction(state)
		break
	// NOP
	case 0x20:
		state.PC++
		break
	// LXI H, D16
	case 0x21:
		UnimplementedInstruction(state)
		break
	// SHLD adr
	case 0x22:
		UnimplementedInstruction(state)
		break
	// INX H
	case 0x23:
		UnimplementedInstruction(state)
		break
	// INR H
	case 0x24:
		state.H++
		state.PC++
		break
	// DCR H
	case 0x25:
		state.H--
		state.PC++
		break
	// MVI H, D8
	case 0x26:
		UnimplementedInstruction(state)
		break
	// DAA
	case 0x27:
		UnimplementedInstruction(state)
		break
	// NOP
	case 0x28:
		state.PC++
		break
	// DAD H
	case 0x29:
		UnimplementedInstruction(state)
		break
	// LHLD adr
	case 0x2a:
		UnimplementedInstruction(state)
		break
	// DCX H
	case 0x2b:
		UnimplementedInstruction(state)
		break
	// INR L
	case 0x2c:
		state.L++
		state.PC++
		break
	// DCR L
	case 0x2d:
		state.L--
		state.PC++
		break
	// MVI L, D8
	case 0x2e:
		UnimplementedInstruction(state)
		break
	// CMA
	case 0x2f:
		UnimplementedInstruction(state)
		break
	// NOP
	case 0x30:
		state.PC++
		break
	// LXI SP, D16
	case 0x31:
		UnimplementedInstruction(state)
		break
	// STA adr
	case 0x32:
		UnimplementedInstruction(state)
		break
	// INX SP
	case 0x33:
		UnimplementedInstruction(state)
		break
	// INR M
	case 0x34:
		break
	// DCR M
	case 0x35:
		UnimplementedInstruction(state)
		break
	// MVI M, D8
	case 0x36:
		UnimplementedInstruction(state)
		break
	// STC
	case 0x37:
		UnimplementedInstruction(state)
		break
	// NOP
	case 0x38:
		break
	// DAD SP
	case 0x39:
		UnimplementedInstruction(state)
		break
	// LDA adr
	case 0x3a:
		UnimplementedInstruction(state)
		break
	// DCX SP
	case 0x3b:
		UnimplementedInstruction(state)
		break
	// INR A
	case 0x3c:
		state.A++
		state.PC++
		break
	case 0x3d:
		state.A--
		state.PC++
		break
	// MVI A, D8
	case 0x3e:
		UnimplementedInstruction(state)
		break
	// CMC
	case 0x3f:
		UnimplementedInstruction(state)
		break
	// MOV B, B
	case 0x40:
		state.B = state.B
		state.PC++
		break
	// MOV B, C
	case 0x41:
		state.B = state.C
		state.PC++
		break
	// MOV B, D
	case 0x42:
		state.B = state.D
		state.PC++
		break
	// MOV B, E
	case 0x43:
		state.B = state.E
		state.PC++
		break
	// MOV B, H
	case 0x44:
		state.B = state.H
		state.PC++
		break
	// MOV B, L
	case 0x45:
		state.B = state.L
		break
	// MOV B, M
	case 0x46:
		UnimplementedInstruction(state)
		break
	// MOV B, A
	case 0x47:
		state.B = state.A
		state.PC++
		break
	// MOV C, B
	case 0x48:
		state.C = state.B
		state.PC++
		break
	// MOV C, C
	case 0x49:
		state.C = state.C
		state.PC++
		break
	// MOV C, D
	case 0x4a:
		state.C = state.D
		state.PC++
		break
	// MOV C, E
	case 0x4b:
		state.C = state.E
		state.PC++
		break
	// MOV C, H
	case 0x4c:
		state.C = state.H
		state.PC++
		break
	// MOV C, L
	case 0x4d:
		state.C = state.L
		state.PC++
		break
	// MOV C, M
	case 0x4e:
		UnimplementedInstruction(state)
		break
	// MOV C, A
	case 0x4f:
		state.C = state.A
		state.PC++
		break
	// MOV D, B
	case 0x50:
		state.D = state.B
		state.PC++
		break
	// MOV D, C
	case 0x51:
		state.D = state.C
		state.PC++
		break
	// MOV D, D
	case 0x52:
		state.D = state.D
		state.PC++
		break
	// MOV D, E
	case 0x53:
		state.D = state.E
		state.PC++
		break
	// MOV D, H
	case 0x54:
		state.D = state.H
		state.PC++
		break
	// MOV D, L
	case 0x55:
		state.D = state.L
		state.PC++
		break
	// MOV D, M
	case 0x56:
		UnimplementedInstruction(state)
		break
	// MOV D, A
	case 0x57:
		state.D = state.A
		state.PC++
		break
	// MOV E, B
	case 0x58:
		state.E = state.B
		state.PC++
		break
	// MOV E, C
	case 0x59:
		state.E = state.C
		state.PC++
		break
	// MOV E, D
	case 0x5a:
		state.E = state.D
		state.PC++
		break
	// MOV E, E
	case 0x5b:
		state.E = state.E
		state.PC++
		break
	// MOV E, H
	case 0x5c:
		state.E = state.H
		state.PC++
		break
	// MOV E, L
	case 0x5d:
		state.E = state.L
		state.PC++
		break
	// MOV E, M
	case 0x5e:
		UnimplementedInstruction(state)
		break
	// MOV E, A
	case 0x5f:
		state.E = state.A
		state.PC++
		break
	// MOV H, B
	case 0x60:
		state.H = state.B
		state.PC++
		break
	// MOV H, C
	case 0x61:
		state.H = state.C
		state.PC++
		break
	// MOV H, D
	case 0x62:
		state.H = state.D
		state.PC++
		break
	// MOV H, E
	case 0x63:
		state.H = state.E
		state.PC++
		break
	// MOV H, H
	case 0x64:
		state.H = state.H
		state.PC++
		break
	// MOV H, L
	case 0x65:
		state.H = state.L
		state.PC++
		break
	// MOV H, M
	case 0x66:
		UnimplementedInstruction(state)
		break
	// MOV H, A
	case 0x67:
		state.H = state.A
		state.PC++
		break
	// MOV L, B
	case 0x68:
		state.L = state.B
		state.PC++
		break
	// MOV L, C
	case 0x69:
		state.L = state.C
		state.PC++
		break
	// MOV L, D
	case 0x6a:
		state.L = state.D
		state.PC++
		break
	// MOV L, E
	case 0x6b:
		state.L = state.E
		state.PC++
		break
	// MOV L, H
	case 0x6c:
		state.L = state.H
		state.PC++
		break
	// MOV L, L
	case 0x6d:
		state.L = state.L
		state.PC++
		break
	// MOV L, M
	case 0x6e:
		UnimplementedInstruction(state)
		break
	// MOV L, A
	case 0x6f:
		state.L = state.A
		state.PC++
		break
	// MOV M, B
	case 0x70:
		UnimplementedInstruction(state)
		break
	// MOV M, C
	case 0x71:
		UnimplementedInstruction(state)
		break
	// MOV M, D
	case 0x72:
		UnimplementedInstruction(state)
		break
	// MOV M, E
	case 0x73:
		UnimplementedInstruction(state)
		break
	// MOV M, H
	case 0x74:
		UnimplementedInstruction(state)
		break
	// MOV M, L
	case 0x75:
		UnimplementedInstruction(state)
		break
	// HLT
	case 0x76:
		state.Quit <- struct{}{}
		break
	// MOV M, A
	case 0x77:
		UnimplementedInstruction(state)
		break
	// MOV A, B
	case 0x78:
		state.A = state.B
		state.PC++
		break
	// MOV A, C
	case 0x79:
		state.A = state.C
		state.PC++
		break
	// MOV A, D
	case 0x7a:
		state.A = state.D
		state.PC++
		break
	// MOV A, E
	case 0x7b:
		state.A = state.E
		state.PC++
		break
	// MOV A, H
	case 0x7c:
		state.A = state.H
		state.PC++
		break
	// MOV A, L
	case 0x7d:
		state.A = state.L
		state.PC++
		break
	// MOV A, M
	case 0x7e:
		UnimplementedInstruction(state)
		break
	// MOV A, A
	case 0x7f:
		state.A = state.A
		state.PC++
		break

	// ADD A, B
	case 0x80:
		answer := uint16(state.A) + uint16(state.B)
		if answer > 255 {
			state.Cc.CY = true
		} else {
			state.Cc.CY = false
		}
		if answer == 0 {
			state.Cc.Z = true
		} else {
			state.Cc.Z = false
		}
		// check if sign bit is set in answer
		if answer&0x80 == 0x80 {
			state.Cc.S = true
		} else {
			state.Cc.S = false
		}
		state.Cc.P = parity(uint8(answer&0xff), 8)
		state.A = uint8(answer)
		state.PC++
		break

	// ADD A, C
	case 0x81:
		answer := uint16(state.A) + uint16(state.C)
		state.Cc.Z = (answer & 0xff) == 0
		state.Cc.S = (answer & 0x80) != 0
		state.Cc.CY = answer > 0xff
		state.Cc.P = parity(uint8(answer&0xff), 8)
		state.A = uint8(answer & 0xff)
		state.PC++
		break

	// ADD A, D
	case 0x82:
		answer := uint16(state.A) + uint16(state.D)
		state.Cc.Z = (answer & 0xff) == 0
		state.Cc.S = (answer & 0x80) != 0
		state.Cc.CY = answer > 0xff
		state.Cc.P = parity(uint8(answer&0xff), 8)
		state.A = uint8(answer & 0xff)
		state.PC++
		break

	// ADD A, E
	case 0x83:
		answer := uint16(state.A) + uint16(state.E)
		state.Cc.Z = (answer & 0xff) == 0
		state.Cc.S = (answer & 0x80) != 0
		state.Cc.CY = answer > 0xff
		state.Cc.P = parity(uint8(answer&0xff), 8)
		state.A = uint8(answer & 0xff)
		state.PC++
		break

	// ADD A, H
	case 0x84:
		answer := uint16(state.A) + uint16(state.H)
		state.Cc.Z = (answer & 0xff) == 0
		state.Cc.S = (answer & 0x80) != 0
		state.Cc.CY = answer > 0xff
		state.Cc.P = parity(uint8(answer&0xff), 8)
		state.A = uint8(answer & 0xff)
		state.PC++
		break

	// ADD A, L
	case 0x85:
		answer := uint16(state.A) + uint16(state.L)
		state.Cc.Z = (answer & 0xff) == 0
		state.Cc.S = (answer & 0x80) != 0
		state.Cc.CY = answer > 0xff
		state.Cc.P = parity(uint8(answer&0xff), 8)
		state.A = uint8(answer & 0xff)
		state.PC++
		break

	// ADD A, M
	case 0x86:
		offset := (uint16(state.H) << 8) | uint16(state.L)
		answer := uint16(state.A) + uint16(state.Memory[offset])
		state.Cc.Z = (answer & 0xff) == 0
		state.Cc.S = (answer & 0x80) != 0
		state.Cc.CY = answer > 0xff
		state.Cc.P = parity(uint8(answer&0xff), 8)
		state.A = uint8(answer & 0xff)
		state.PC++
		break

	// ADD A, A
	case 0x87:
		answer := uint16(state.A) + uint16(state.A)
		state.Cc.Z = (answer & 0xff) == 0
		state.Cc.S = (answer & 0x80) != 0
		state.Cc.CY = answer > 0xff
		state.Cc.P = parity(uint8(answer&0xff), 8)
		state.A = uint8(answer & 0xff)
		state.PC++
		break

	// ADC A, B
	case 0x88:
		UnimplementedInstruction(state)
		break
	case 0x89:
		UnimplementedInstruction(state)
		break
	case 0x8a:
		UnimplementedInstruction(state)
		break
	case 0x8b:
		UnimplementedInstruction(state)
		break
	case 0x8c:
		UnimplementedInstruction(state)
		break
	case 0x8d:
		UnimplementedInstruction(state)
		break
	case 0x8e:
		UnimplementedInstruction(state)
		break
	case 0x8f:
		UnimplementedInstruction(state)
		break

	// SUB B
	case 0x90:
		answer := uint16(state.A) - uint16(state.B)
		state.Cc.Z = (answer & 0xff) == 0
		state.Cc.S = (answer & 0x80) != 0
		state.Cc.CY = answer > 0xff
		state.Cc.P = parity(uint8(answer&0xff), 8)
		state.A = uint8(answer & 0xff)
		state.Cc.Z = state.A == 0
		state.PC++
		break

	// SUB C
	case 0x91:
		answer := uint16(state.A) - uint16(state.C)
		state.Cc.Z = (answer & 0xff) == 0
		state.Cc.S = (answer & 0x80) != 0
		state.Cc.CY = answer > 0xff
		state.Cc.P = parity(uint8(answer&0xff), 8)
		state.A = uint8(answer & 0xff)
		state.Cc.Z = state.A == 0
		state.PC++
		break

	// SUB D
	case 0x92:
		answer := uint16(state.A) - uint16(state.D)
		state.Cc.Z = (answer & 0xff) == 0
		state.Cc.S = (answer & 0x80) != 0
		state.Cc.CY = answer > 0xff
		state.Cc.P = parity(uint8(answer&0xff), 8)
		state.A = uint8(answer & 0xff)
		state.Cc.Z = state.A == 0
		state.PC++
		break

	// SUB E
	case 0x93:
		answer := uint16(state.A) - uint16(state.E)
		state.Cc.Z = (answer & 0xff) == 0
		state.Cc.S = (answer & 0x80) != 0
		state.Cc.CY = answer > 0xff
		state.Cc.P = parity(uint8(answer&0xff), 8)
		state.A = uint8(answer & 0xff)
		state.Cc.Z = state.A == 0
		state.PC++
		break

	// SUB H
	case 0x94:
		answer := uint16(state.A) - uint16(state.H)
		state.Cc.Z = (answer & 0xff) == 0
		state.Cc.S = (answer & 0x80) != 0
		state.Cc.CY = answer > 0xff
		state.Cc.P = parity(uint8(answer&0xff), 8)
		state.A = uint8(answer & 0xff)
		state.Cc.Z = state.A == 0
		state.PC++
		break

	// SUB L
	case 0x95:
		answer := uint16(state.A) - uint16(state.L)
		state.Cc.Z = (answer & 0xff) == 0
		state.Cc.S = (answer & 0x80) != 0
		state.Cc.CY = answer > 0xff
		state.Cc.P = parity(uint8(answer&0xff), 8)
		state.A = uint8(answer & 0xff)
		state.Cc.Z = state.A == 0
		state.PC++
		break

	// SUB M
	case 0x96:
		UnimplementedInstruction(state)
		break

	// SUB A
	case 0x97:
		answer := uint16(state.A) - uint16(state.A)
		state.Cc.Z = (answer & 0xff) == 0
		state.Cc.S = (answer & 0x80) != 0
		state.Cc.CY = answer > 0xff
		state.Cc.P = parity(uint8(answer&0xff), 8)
		state.A = uint8(answer & 0xff)
		state.Cc.Z = state.A == 0
		state.PC++
		break

	// SBB B
	case 0x98:
		UnimplementedInstruction(state)
		break

	// SBB C
	case 0x99:
		UnimplementedInstruction(state)
		break

	// SBB D
	case 0x9a:
		UnimplementedInstruction(state)
		break

	// SBB E
	case 0x9b:
		UnimplementedInstruction(state)
		break

	// SBB H
	case 0x9c:
		UnimplementedInstruction(state)
		break

	// SBB L
	case 0x9d:
		UnimplementedInstruction(state)
		break

	// SBB M
	case 0x9e:
		UnimplementedInstruction(state)
		break

	// SBB A
	case 0x9f:
		UnimplementedInstruction(state)
		break

	// ANA B
	case 0xa0:
		state.A = state.A & state.B
		state.Cc.Z = state.A == 0
		state.PC++
		break

	// ANA C
	case 0xa1:
		state.A = state.A & state.C
		state.Cc.Z = state.A == 0
		state.PC++
		break

	// ANA D
	case 0xa2:
		state.A = state.A & state.D
		state.Cc.Z = state.A == 0
		state.PC++
		break

	// ANA E
	case 0xa3:
		state.A = state.A & state.E
		state.Cc.Z = state.A == 0
		state.PC++
		break

	// ANA H
	case 0xa4:
		state.A = state.A & state.H
		state.Cc.Z = state.A == 0
		state.PC++
		break

	// ANA L
	case 0xa5:
		state.A = state.A & state.L
		state.Cc.Z = state.A == 0
		state.PC++
		break

	// ANA M
	case 0xa6:
		UnimplementedInstruction(state)
		break

	// ANA A
	case 0xa7:
		state.A = state.A & state.A
		state.Cc.Z = state.A == 0
		state.PC++
		break

	// XRA B
	case 0xa8:
		state.A = state.A ^ state.B
		state.Cc.Z = state.A == 0
		state.PC++
		break

	// XRA C
	case 0xa9:
		state.A = state.A ^ state.C
		state.Cc.Z = state.A == 0
		state.PC++
		break

	// XRA D
	case 0xaa:
		state.A = state.A ^ state.D
		state.Cc.Z = state.A == 0
		state.PC++
		break

	// XRA E
	case 0xab:
		state.A = state.A ^ state.E
		state.Cc.Z = state.A == 0
		state.PC++
		break

	// XRA H
	case 0xac:
		state.A = state.A ^ state.H
		state.Cc.Z = state.A == 0
		state.PC++
		break

	// XRA L
	case 0xad:
		state.A = state.A ^ state.L
		state.Cc.Z = state.A == 0
		state.PC++
		break

	// XRA M
	case 0xae:
		UnimplementedInstruction(state)
		break

	// XRA A
	case 0xaf:
		state.A = state.A ^ state.A
		state.Cc.Z = state.A == 0
		state.PC++
		break

	// ORA B
	case 0xb0:
		state.A = state.A | state.B
		state.Cc.Z = state.A == 0
		state.PC++
		break

	// ORA C
	case 0xb1:
		state.A = state.A | state.C
		state.Cc.Z = state.A == 0
		state.PC++
		break

	// ORA D
	case 0xb2:
		state.A = state.A | state.D
		state.Cc.Z = state.A == 0
		state.PC++
		break

	// ORA E
	case 0xb3:
		state.A = state.A | state.E
		state.Cc.Z = state.A == 0
		state.PC++
		break

	// ORA H
	case 0xb4:
		state.A = state.A | state.H
		state.Cc.Z = state.A == 0
		state.PC++
		break

	// ORA L
	case 0xb5:
		state.A = state.A | state.L
		state.Cc.Z = state.A == 0
		state.PC++
		break

	// ORA M
	case 0xb6:
		UnimplementedInstruction(state)
		break

	// ORA A
	case 0xb7:
		state.A = state.A | state.A
		state.Cc.Z = state.A == 0
		state.PC++
		break

	// CMP B
	case 0xb8:
		state.Cc.Z = state.A == state.B
		state.Cc.S = (state.A - state.B) > 0x7f
		state.Cc.P = parity(state.A-state.B, 8)
		state.Cc.CY = state.A < state.B
		state.PC++
		break
	case 0xb9:
		UnimplementedInstruction(state)
		break
	case 0xba:
		UnimplementedInstruction(state)
		break
	case 0xbb:
		UnimplementedInstruction(state)
		break
	case 0xbc:
		UnimplementedInstruction(state)
		break
	case 0xbd:
		UnimplementedInstruction(state)
		break
	case 0xbe:
		UnimplementedInstruction(state)
		break
	case 0xbf:
		UnimplementedInstruction(state)
		break

	// RNZ
	case 0xc0:
		UnimplementedInstruction(state)
		break

	// POP B
	case 0xc1:
		UnimplementedInstruction(state)
		break

	// JNZ adr
	case 0xc2:
		if !state.Cc.Z {
			state.PC = (uint16(state.Memory[state.PC+2]) << 8) | uint16(state.Memory[state.PC+1])
		} else {
			state.PC += 2
		}
		break

	// JMP adr
	case 0xc3:
		state.PC = (uint16(state.Memory[state.PC+2]) << 8) | uint16(state.Memory[state.PC+1])
		break
	case 0xc4:
		UnimplementedInstruction(state)
		break
	case 0xc5:
		UnimplementedInstruction(state)
		break
	case 0xc6:
		answer := uint16(state.A) + state.PC + 1
		state.Cc.Z = answer&0xff == 0
		state.Cc.S = (answer & 0x80) != 0
		state.Cc.CY = answer > 0xff
		state.Cc.P = parity(uint8(answer&0xff), 8)
		state.PC += 2
		break
	case 0xc7:
		UnimplementedInstruction(state)
		break
	case 0xc8:
		UnimplementedInstruction(state)
		break
	case 0xc9:
		UnimplementedInstruction(state)
		break
	case 0xca:
		UnimplementedInstruction(state)
		break
	case 0xcb:
		UnimplementedInstruction(state)
		break
	case 0xcc:
		UnimplementedInstruction(state)
		break
	case 0xcd:
		UnimplementedInstruction(state)
		break
	case 0xce:
		UnimplementedInstruction(state)
		break
	case 0xcf:
		UnimplementedInstruction(state)
		break
	case 0xd0:
		UnimplementedInstruction(state)
		break
	case 0xd1:
		UnimplementedInstruction(state)
		break
	case 0xd2:
		UnimplementedInstruction(state)
		break
	case 0xd3:
		UnimplementedInstruction(state)
		break
	case 0xd4:
		UnimplementedInstruction(state)
		break
	case 0xd5:
		UnimplementedInstruction(state)
		break
	case 0xd6:
		UnimplementedInstruction(state)
		break
	case 0xd7:
		UnimplementedInstruction(state)
		break
	case 0xd8:
		UnimplementedInstruction(state)
		break
	case 0xd9:
		UnimplementedInstruction(state)
		break
	case 0xda:
		UnimplementedInstruction(state)
		break
	case 0xdb:
		UnimplementedInstruction(state)
		break
	case 0xdc:
		UnimplementedInstruction(state)
		break
	case 0xdd:
		UnimplementedInstruction(state)
		break
	case 0xde:
		UnimplementedInstruction(state)
		break
	case 0xdf:
		UnimplementedInstruction(state)
		break
	case 0xe0:
		UnimplementedInstruction(state)
		break
	case 0xe1:
		UnimplementedInstruction(state)
		break
	case 0xe2:
		UnimplementedInstruction(state)
		break
	case 0xe3:
		UnimplementedInstruction(state)
		break
	case 0xe4:
		UnimplementedInstruction(state)
		break
	case 0xe5:
		UnimplementedInstruction(state)
		break
	case 0xe6:
		UnimplementedInstruction(state)
		break
	case 0xe7:
		UnimplementedInstruction(state)
		break
	case 0xe8:
		UnimplementedInstruction(state)
		break
	case 0xe9:
		UnimplementedInstruction(state)
		break
	case 0xea:
		UnimplementedInstruction(state)
		break
	case 0xeb:
		UnimplementedInstruction(state)
		break
	case 0xec:
		UnimplementedInstruction(state)
		break
	case 0xed:
		UnimplementedInstruction(state)
		break
	case 0xee:
		UnimplementedInstruction(state)
		break
	case 0xef:
		UnimplementedInstruction(state)
		break
	case 0xf0:
		UnimplementedInstruction(state)
		break
	case 0xf1:
		UnimplementedInstruction(state)
		break
	case 0xf2:
		UnimplementedInstruction(state)
		break
	case 0xf3:
		UnimplementedInstruction(state)
		break
	case 0xf4:
		UnimplementedInstruction(state)
		break
	case 0xf5:
		UnimplementedInstruction(state)
		break
	case 0xf6:
		UnimplementedInstruction(state)
		break
	case 0xf7:
		UnimplementedInstruction(state)
		break
	case 0xf8:
		UnimplementedInstruction(state)
		break
	case 0xf9:
		UnimplementedInstruction(state)
		break
	case 0xfa:
		UnimplementedInstruction(state)
		break
	case 0xfb:
		UnimplementedInstruction(state)
		break
	case 0xfc:
		UnimplementedInstruction(state)
		break
	case 0xfd:
		UnimplementedInstruction(state)
		break
	case 0xfe:
		UnimplementedInstruction(state)
		break
	case 0xff:
		UnimplementedInstruction(state)
		break
	default:
		fmt.Printf("Unknown opcode: %X", state.Memory[state.PC])
		state.Quit <- struct{}{}
	}
	return nil
}
