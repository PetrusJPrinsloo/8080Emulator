package main

import (
	"fmt"
	"log"
)

type ConditionCodes struct {
	Z   bool
	S   bool
	P   bool
	CY  bool
	AC  bool
	Pad uint8
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
	case 0x40:
		UnimplementedInstruction(state)
		break
	case 0x41:
		UnimplementedInstruction(state)
		break
	case 0x42:
		UnimplementedInstruction(state)
		break
	case 0x43:
		UnimplementedInstruction(state)
		break
	case 0x44:
		UnimplementedInstruction(state)
		break
	case 0x45:
		UnimplementedInstruction(state)
		break
	case 0x46:
		UnimplementedInstruction(state)
		break
	case 0x47:
		UnimplementedInstruction(state)
		break
	case 0x48:
		UnimplementedInstruction(state)
		break
	case 0x49:
		UnimplementedInstruction(state)
		break
	case 0x4a:
		UnimplementedInstruction(state)
		break
	case 0x4b:
		UnimplementedInstruction(state)
		break
	case 0x4c:
		UnimplementedInstruction(state)
		break
	case 0x4d:
		UnimplementedInstruction(state)
		break
	case 0x4e:
		UnimplementedInstruction(state)
		break
	case 0x4f:
		UnimplementedInstruction(state)
		break
	case 0x50:
		UnimplementedInstruction(state)
		break
	case 0x51:
		UnimplementedInstruction(state)
		break
	case 0x52:
		UnimplementedInstruction(state)
		break
	case 0x53:
		UnimplementedInstruction(state)
		break
	case 0x54:
		UnimplementedInstruction(state)
		break
	case 0x55:
		UnimplementedInstruction(state)
		break
	case 0x56:
		UnimplementedInstruction(state)
		break
	case 0x57:
		UnimplementedInstruction(state)
		break
	case 0x58:
		UnimplementedInstruction(state)
		break
	case 0x59:
		UnimplementedInstruction(state)
		break
	case 0x5a:
		UnimplementedInstruction(state)
		break
	case 0x5b:
		UnimplementedInstruction(state)
		break
	case 0x5c:
		UnimplementedInstruction(state)
		break
	case 0x5d:
		UnimplementedInstruction(state)
		break
	case 0x5e:
		UnimplementedInstruction(state)
		break
	case 0x5f:
		UnimplementedInstruction(state)
		break
	case 0x60:
		UnimplementedInstruction(state)
		break
	case 0x61:
		UnimplementedInstruction(state)
		break
	case 0x62:
		UnimplementedInstruction(state)
		break
	case 0x63:
		UnimplementedInstruction(state)
		break
	case 0x64:
		UnimplementedInstruction(state)
		break
	case 0x65:
		UnimplementedInstruction(state)
		break
	case 0x66:
		UnimplementedInstruction(state)
		break
	case 0x67:
		UnimplementedInstruction(state)
		break
	case 0x68:
		UnimplementedInstruction(state)
		break
	case 0x69:
		UnimplementedInstruction(state)
		break
	case 0x6a:
		UnimplementedInstruction(state)
		break
	case 0x6b:
		UnimplementedInstruction(state)
		break
	case 0x6c:
		UnimplementedInstruction(state)
		break
	case 0x6d:
		UnimplementedInstruction(state)
		break
	case 0x6e:
		UnimplementedInstruction(state)
		break
	case 0x6f:
		UnimplementedInstruction(state)
		break
	case 0x70:
		UnimplementedInstruction(state)
		break
	case 0x71:
		UnimplementedInstruction(state)
		break
	case 0x72:
		UnimplementedInstruction(state)
		break
	case 0x73:
		UnimplementedInstruction(state)
		break
	case 0x74:
		UnimplementedInstruction(state)
		break
	case 0x75:
		UnimplementedInstruction(state)
		break
	case 0x76:
		UnimplementedInstruction(state)
		break
	case 0x77:
		UnimplementedInstruction(state)
		break
	case 0x78:
		UnimplementedInstruction(state)
		break
	case 0x79:
		UnimplementedInstruction(state)
		break
	case 0x7a:
		UnimplementedInstruction(state)
		break
	case 0x7b:
		UnimplementedInstruction(state)
		break
	case 0x7c:
		UnimplementedInstruction(state)
		break
	case 0x7d:
		UnimplementedInstruction(state)
		break
	case 0x7e:
		UnimplementedInstruction(state)
		break
	case 0x7f:
		UnimplementedInstruction(state)
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
	case 0x82:
		UnimplementedInstruction(state)
		break
	case 0x83:
		UnimplementedInstruction(state)
		break
	case 0x84:
		UnimplementedInstruction(state)
		break
	case 0x85:
		UnimplementedInstruction(state)
		break
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
	case 0x87:
		UnimplementedInstruction(state)
		break
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
	case 0x90:
		UnimplementedInstruction(state)
		break
	case 0x91:
		UnimplementedInstruction(state)
		break
	case 0x92:
		UnimplementedInstruction(state)
		break
	case 0x93:
		UnimplementedInstruction(state)
		break
	case 0x94:
		UnimplementedInstruction(state)
		break
	case 0x95:
		UnimplementedInstruction(state)
		break
	case 0x96:
		UnimplementedInstruction(state)
		break
	case 0x97:
		UnimplementedInstruction(state)
		break
	case 0x98:
		UnimplementedInstruction(state)
		break
	case 0x99:
		UnimplementedInstruction(state)
		break
	case 0x9a:
		UnimplementedInstruction(state)
		break
	case 0x9b:
		UnimplementedInstruction(state)
		break
	case 0x9c:
		UnimplementedInstruction(state)
		break
	case 0x9d:
		UnimplementedInstruction(state)
		break
	case 0x9e:
		UnimplementedInstruction(state)
		break
	case 0x9f:
		UnimplementedInstruction(state)
		break
	case 0xa0:
		UnimplementedInstruction(state)
		break
	case 0xa1:
		UnimplementedInstruction(state)
		break
	case 0xa2:
		UnimplementedInstruction(state)
		break
	case 0xa3:
		UnimplementedInstruction(state)
		break
	case 0xa4:
		UnimplementedInstruction(state)
		break
	case 0xa5:
		UnimplementedInstruction(state)
		break
	case 0xa6:
		UnimplementedInstruction(state)
		break
	case 0xa7:
		UnimplementedInstruction(state)
		break
	case 0xa8:
		UnimplementedInstruction(state)
		break
	case 0xa9:
		UnimplementedInstruction(state)
		break
	case 0xaa:
		UnimplementedInstruction(state)
		break
	case 0xab:
		UnimplementedInstruction(state)
		break
	case 0xac:
		UnimplementedInstruction(state)
		break
	case 0xad:
		UnimplementedInstruction(state)
		break
	case 0xae:
		UnimplementedInstruction(state)
		break
	case 0xaf:
		UnimplementedInstruction(state)
		break
	case 0xb0:
		UnimplementedInstruction(state)
		break
	case 0xb1:
		UnimplementedInstruction(state)
		break
	case 0xb2:
		UnimplementedInstruction(state)
		break
	case 0xb3:
		UnimplementedInstruction(state)
		break
	case 0xb4:
		UnimplementedInstruction(state)
		break
	case 0xb5:
		UnimplementedInstruction(state)
		break
	case 0xb6:
		UnimplementedInstruction(state)
		break
	case 0xb7:
		UnimplementedInstruction(state)
		break
	case 0xb8:
		UnimplementedInstruction(state)
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
	case 0xc0:
		UnimplementedInstruction(state)
		break
	case 0xc1:
		UnimplementedInstruction(state)
		break
	case 0xc2:
		UnimplementedInstruction(state)
		break
	case 0xc3:
		UnimplementedInstruction(state)
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
