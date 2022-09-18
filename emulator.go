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
	return &state
}

func (state *State8080) Step() error {
	return Emulate8080Op(state)
}

func UnimplementedInstruction() {
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
		state.PC += 2
		break

	// STAX B
	case 0x02:
		break

	// INX B
	case 0x03:
		state.B++
		state.C++
		break

	// INR B
	case 0x04:
		state.B++
		break

	// DCR B
	case 0x05:
		state.B--
		break

	// MVI B, D8
	case 0x06:
		UnimplementedInstruction()
		break

	// RLC
	case 0x07:
		UnimplementedInstruction()
		break
	case 0x08:
		UnimplementedInstruction()
		break
	case 0x09:
		UnimplementedInstruction()
		break
	case 0x0a:
		UnimplementedInstruction()
		break
	case 0x0b:
		UnimplementedInstruction()
		break
	case 0x0c:
		UnimplementedInstruction()
		break
	case 0x0d:
		UnimplementedInstruction()
		break
	case 0x0e:
		UnimplementedInstruction()
		break
	case 0x0f:
		UnimplementedInstruction()
		break
	case 0x10:
		UnimplementedInstruction()
		break
	case 0x11:
		UnimplementedInstruction()
		break
	case 0x12:
		UnimplementedInstruction()
		break
	case 0x13:
		UnimplementedInstruction()
		break
	case 0x14:
		UnimplementedInstruction()
		break
	case 0x15:
		UnimplementedInstruction()
		break
	case 0x16:
		UnimplementedInstruction()
		break
	case 0x17:
		UnimplementedInstruction()
		break
	case 0x18:
		UnimplementedInstruction()
		break
	case 0x19:
		UnimplementedInstruction()
		break
	case 0x1a:
		UnimplementedInstruction()
		break
	case 0x1b:
		UnimplementedInstruction()
		break
	case 0x1c:
		UnimplementedInstruction()
		break
	case 0x1d:
		UnimplementedInstruction()
		break
	case 0x1e:
		UnimplementedInstruction()
		break
	case 0x1f:
		UnimplementedInstruction()
		break
	case 0x20:
		UnimplementedInstruction()
		break
	case 0x21:
		UnimplementedInstruction()
		break
	case 0x22:
		UnimplementedInstruction()
		break
	case 0x23:
		UnimplementedInstruction()
		break
	case 0x24:
		UnimplementedInstruction()
		break
	case 0x25:
		UnimplementedInstruction()
		break
	case 0x26:
		UnimplementedInstruction()
		break
	case 0x27:
		UnimplementedInstruction()
		break
	case 0x28:
		UnimplementedInstruction()
		break
	case 0x29:
		UnimplementedInstruction()
		break
	case 0x2a:
		UnimplementedInstruction()
		break
	case 0x2b:
		UnimplementedInstruction()
		break
	case 0x2c:
		UnimplementedInstruction()
		break
	case 0x2d:
		UnimplementedInstruction()
		break
	case 0x2e:
		UnimplementedInstruction()
		break
	case 0x2f:
		UnimplementedInstruction()
		break
	case 0x30:
		UnimplementedInstruction()
		break
	case 0x31:
		UnimplementedInstruction()
		break
	case 0x32:
		UnimplementedInstruction()
		break
	case 0x33:
		UnimplementedInstruction()
		break
	case 0x34:
		UnimplementedInstruction()
		break
	case 0x35:
		UnimplementedInstruction()
		break
	case 0x36:
		UnimplementedInstruction()
		break
	case 0x37:
		UnimplementedInstruction()
		break
	case 0x38:
		UnimplementedInstruction()
		break
	case 0x39:
		UnimplementedInstruction()
		break
	case 0x3a:
		UnimplementedInstruction()
		break
	case 0x3b:
		UnimplementedInstruction()
		break
	case 0x3c:
		UnimplementedInstruction()
		break
	case 0x3d:
		UnimplementedInstruction()
		break
	case 0x3e:
		UnimplementedInstruction()
		break
	case 0x3f:
		UnimplementedInstruction()
		break
	case 0x40:
		UnimplementedInstruction()
		break
	case 0x41:
		UnimplementedInstruction()
		break
	case 0x42:
		UnimplementedInstruction()
		break
	case 0x43:
		UnimplementedInstruction()
		break
	case 0x44:
		UnimplementedInstruction()
		break
	case 0x45:
		UnimplementedInstruction()
		break
	case 0x46:
		UnimplementedInstruction()
		break
	case 0x47:
		UnimplementedInstruction()
		break
	case 0x48:
		UnimplementedInstruction()
		break
	case 0x49:
		UnimplementedInstruction()
		break
	case 0x4a:
		UnimplementedInstruction()
		break
	case 0x4b:
		UnimplementedInstruction()
		break
	case 0x4c:
		UnimplementedInstruction()
		break
	case 0x4d:
		UnimplementedInstruction()
		break
	case 0x4e:
		UnimplementedInstruction()
		break
	case 0x4f:
		UnimplementedInstruction()
		break
	case 0x50:
		UnimplementedInstruction()
		break
	case 0x51:
		UnimplementedInstruction()
		break
	case 0x52:
		UnimplementedInstruction()
		break
	case 0x53:
		UnimplementedInstruction()
		break
	case 0x54:
		UnimplementedInstruction()
		break
	case 0x55:
		UnimplementedInstruction()
		break
	case 0x56:
		UnimplementedInstruction()
		break
	case 0x57:
		UnimplementedInstruction()
		break
	case 0x58:
		UnimplementedInstruction()
		break
	case 0x59:
		UnimplementedInstruction()
		break
	case 0x5a:
		UnimplementedInstruction()
		break
	case 0x5b:
		UnimplementedInstruction()
		break
	case 0x5c:
		UnimplementedInstruction()
		break
	case 0x5d:
		UnimplementedInstruction()
		break
	case 0x5e:
		UnimplementedInstruction()
		break
	case 0x5f:
		UnimplementedInstruction()
		break
	case 0x60:
		UnimplementedInstruction()
		break
	case 0x61:
		UnimplementedInstruction()
		break
	case 0x62:
		UnimplementedInstruction()
		break
	case 0x63:
		UnimplementedInstruction()
		break
	case 0x64:
		UnimplementedInstruction()
		break
	case 0x65:
		UnimplementedInstruction()
		break
	case 0x66:
		UnimplementedInstruction()
		break
	case 0x67:
		UnimplementedInstruction()
		break
	case 0x68:
		UnimplementedInstruction()
		break
	case 0x69:
		UnimplementedInstruction()
		break
	case 0x6a:
		UnimplementedInstruction()
		break
	case 0x6b:
		UnimplementedInstruction()
		break
	case 0x6c:
		UnimplementedInstruction()
		break
	case 0x6d:
		UnimplementedInstruction()
		break
	case 0x6e:
		UnimplementedInstruction()
		break
	case 0x6f:
		UnimplementedInstruction()
		break
	case 0x70:
		UnimplementedInstruction()
		break
	case 0x71:
		UnimplementedInstruction()
		break
	case 0x72:
		UnimplementedInstruction()
		break
	case 0x73:
		UnimplementedInstruction()
		break
	case 0x74:
		UnimplementedInstruction()
		break
	case 0x75:
		UnimplementedInstruction()
		break
	case 0x76:
		UnimplementedInstruction()
		break
	case 0x77:
		UnimplementedInstruction()
		break
	case 0x78:
		UnimplementedInstruction()
		break
	case 0x79:
		UnimplementedInstruction()
		break
	case 0x7a:
		UnimplementedInstruction()
		break
	case 0x7b:
		UnimplementedInstruction()
		break
	case 0x7c:
		UnimplementedInstruction()
		break
	case 0x7d:
		UnimplementedInstruction()
		break
	case 0x7e:
		UnimplementedInstruction()
		break
	case 0x7f:
		UnimplementedInstruction()
		break
	case 0x80: // ADD A, B
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
		break
	case 0x81:
		answer := uint16(state.A) + uint16(state.C)
		state.Cc.Z = (answer & 0xff) == 0
		state.Cc.S = (answer & 0x80) != 0
		state.Cc.CY = answer > 0xff
		state.Cc.P = parity(uint8(answer&0xff), 8)
		state.A = uint8(answer & 0xff)
		break
	case 0x82:
		UnimplementedInstruction()
		break
	case 0x83:
		UnimplementedInstruction()
		break
	case 0x84:
		UnimplementedInstruction()
		break
	case 0x85:
		UnimplementedInstruction()
		break
	case 0x86:
		offset := (uint16(state.H) << 8) | uint16(state.L)
		answer := uint16(state.A) + uint16(state.Memory[offset])
		state.Cc.Z = (answer & 0xff) == 0
		state.Cc.S = (answer & 0x80) != 0
		state.Cc.CY = answer > 0xff
		state.Cc.P = parity(uint8(answer&0xff), 8)
		state.A = uint8(answer & 0xff)
		break
	case 0x87:
		UnimplementedInstruction()
		break
	case 0x88:
		UnimplementedInstruction()
		break
	case 0x89:
		UnimplementedInstruction()
		break
	case 0x8a:
		UnimplementedInstruction()
		break
	case 0x8b:
		UnimplementedInstruction()
		break
	case 0x8c:
		UnimplementedInstruction()
		break
	case 0x8d:
		UnimplementedInstruction()
		break
	case 0x8e:
		UnimplementedInstruction()
		break
	case 0x8f:
		UnimplementedInstruction()
		break
	case 0x90:
		UnimplementedInstruction()
		break
	case 0x91:
		UnimplementedInstruction()
		break
	case 0x92:
		UnimplementedInstruction()
		break
	case 0x93:
		UnimplementedInstruction()
		break
	case 0x94:
		UnimplementedInstruction()
		break
	case 0x95:
		UnimplementedInstruction()
		break
	case 0x96:
		UnimplementedInstruction()
		break
	case 0x97:
		UnimplementedInstruction()
		break
	case 0x98:
		UnimplementedInstruction()
		break
	case 0x99:
		UnimplementedInstruction()
		break
	case 0x9a:
		UnimplementedInstruction()
		break
	case 0x9b:
		UnimplementedInstruction()
		break
	case 0x9c:
		UnimplementedInstruction()
		break
	case 0x9d:
		UnimplementedInstruction()
		break
	case 0x9e:
		UnimplementedInstruction()
		break
	case 0x9f:
		UnimplementedInstruction()
		break
	case 0xa0:
		UnimplementedInstruction()
		break
	case 0xa1:
		UnimplementedInstruction()
		break
	case 0xa2:
		UnimplementedInstruction()
		break
	case 0xa3:
		UnimplementedInstruction()
		break
	case 0xa4:
		UnimplementedInstruction()
		break
	case 0xa5:
		UnimplementedInstruction()
		break
	case 0xa6:
		UnimplementedInstruction()
		break
	case 0xa7:
		UnimplementedInstruction()
		break
	case 0xa8:
		UnimplementedInstruction()
		break
	case 0xa9:
		UnimplementedInstruction()
		break
	case 0xaa:
		UnimplementedInstruction()
		break
	case 0xab:
		UnimplementedInstruction()
		break
	case 0xac:
		UnimplementedInstruction()
		break
	case 0xad:
		UnimplementedInstruction()
		break
	case 0xae:
		UnimplementedInstruction()
		break
	case 0xaf:
		UnimplementedInstruction()
		break
	case 0xb0:
		UnimplementedInstruction()
		break
	case 0xb1:
		UnimplementedInstruction()
		break
	case 0xb2:
		UnimplementedInstruction()
		break
	case 0xb3:
		UnimplementedInstruction()
		break
	case 0xb4:
		UnimplementedInstruction()
		break
	case 0xb5:
		UnimplementedInstruction()
		break
	case 0xb6:
		UnimplementedInstruction()
		break
	case 0xb7:
		UnimplementedInstruction()
		break
	case 0xb8:
		UnimplementedInstruction()
		break
	case 0xb9:
		UnimplementedInstruction()
		break
	case 0xba:
		UnimplementedInstruction()
		break
	case 0xbb:
		UnimplementedInstruction()
		break
	case 0xbc:
		UnimplementedInstruction()
		break
	case 0xbd:
		UnimplementedInstruction()
		break
	case 0xbe:
		UnimplementedInstruction()
		break
	case 0xbf:
		UnimplementedInstruction()
		break
	case 0xc0:
		UnimplementedInstruction()
		break
	case 0xc1:
		UnimplementedInstruction()
		break
	case 0xc2:
		UnimplementedInstruction()
		break
	case 0xc3:
		UnimplementedInstruction()
		break
	case 0xc4:
		UnimplementedInstruction()
		break
	case 0xc5:
		UnimplementedInstruction()
		break
	case 0xc6:
		answer := uint16(state.A) + state.PC + 1
		state.Cc.Z = answer&0xff == 0
		state.Cc.S = (answer & 0x80) != 0
		state.Cc.CY = answer > 0xff
		state.Cc.P = parity(uint8(answer&0xff), 8)
		break
	case 0xc7:
		UnimplementedInstruction()
		break
	case 0xc8:
		UnimplementedInstruction()
		break
	case 0xc9:
		UnimplementedInstruction()
		break
	case 0xca:
		UnimplementedInstruction()
		break
	case 0xcb:
		UnimplementedInstruction()
		break
	case 0xcc:
		UnimplementedInstruction()
		break
	case 0xcd:
		UnimplementedInstruction()
		break
	case 0xce:
		UnimplementedInstruction()
		break
	case 0xcf:
		UnimplementedInstruction()
		break
	case 0xd0:
		UnimplementedInstruction()
		break
	case 0xd1:
		UnimplementedInstruction()
		break
	case 0xd2:
		UnimplementedInstruction()
		break
	case 0xd3:
		UnimplementedInstruction()
		break
	case 0xd4:
		UnimplementedInstruction()
		break
	case 0xd5:
		UnimplementedInstruction()
		break
	case 0xd6:
		UnimplementedInstruction()
		break
	case 0xd7:
		UnimplementedInstruction()
		break
	case 0xd8:
		UnimplementedInstruction()
		break
	case 0xd9:
		UnimplementedInstruction()
		break
	case 0xda:
		UnimplementedInstruction()
		break
	case 0xdb:
		UnimplementedInstruction()
		break
	case 0xdc:
		UnimplementedInstruction()
		break
	case 0xdd:
		UnimplementedInstruction()
		break
	case 0xde:
		UnimplementedInstruction()
		break
	case 0xdf:
		UnimplementedInstruction()
		break
	case 0xe0:
		UnimplementedInstruction()
		break
	case 0xe1:
		UnimplementedInstruction()
		break
	case 0xe2:
		UnimplementedInstruction()
		break
	case 0xe3:
		UnimplementedInstruction()
		break
	case 0xe4:
		UnimplementedInstruction()
		break
	case 0xe5:
		UnimplementedInstruction()
		break
	case 0xe6:
		UnimplementedInstruction()
		break
	case 0xe7:
		UnimplementedInstruction()
		break
	case 0xe8:
		UnimplementedInstruction()
		break
	case 0xe9:
		UnimplementedInstruction()
		break
	case 0xea:
		UnimplementedInstruction()
		break
	case 0xeb:
		UnimplementedInstruction()
		break
	case 0xec:
		UnimplementedInstruction()
		break
	case 0xed:
		UnimplementedInstruction()
		break
	case 0xee:
		UnimplementedInstruction()
		break
	case 0xef:
		UnimplementedInstruction()
		break
	case 0xf0:
		UnimplementedInstruction()
		break
	case 0xf1:
		UnimplementedInstruction()
		break
	case 0xf2:
		UnimplementedInstruction()
		break
	case 0xf3:
		UnimplementedInstruction()
		break
	case 0xf4:
		UnimplementedInstruction()
		break
	case 0xf5:
		UnimplementedInstruction()
		break
	case 0xf6:
		UnimplementedInstruction()
		break
	case 0xf7:
		UnimplementedInstruction()
		break
	case 0xf8:
		UnimplementedInstruction()
		break
	case 0xf9:
		UnimplementedInstruction()
		break
	case 0xfa:
		UnimplementedInstruction()
		break
	case 0xfb:
		UnimplementedInstruction()
		break
	case 0xfc:
		UnimplementedInstruction()
		break
	case 0xfd:
		UnimplementedInstruction()
		break
	case 0xfe:
		UnimplementedInstruction()
		break
	case 0xff:
		UnimplementedInstruction()
		break
	default:
		fmt.Printf("Unknown opcode: %X", state.Memory[state.PC])
		state.Quit <- struct{}{}
	}
	return nil
}
