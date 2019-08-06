package arm7

import "log"

// SysRegisters define the registers for the SYS CPU Mode
type SysRegisters struct {
	// General Purpose Registers
	R0  uint32
	R1  uint32
	R2  uint32
	R3  uint32
	R4  uint32
	R5  uint32
	R6  uint32
	R7  uint32
	R8  uint32
	R9  uint32
	R10 uint32
	R11 uint32
	R12 uint32
	// Stack Pointer - SP
	R13 uint32
	// Link Register - LP
	R14 uint32
	// Program Counter - PC
	R15 uint32
	// Current Program Status Register - CPSR
	Cpsr uint32
}

func (sysRegisters *SysRegisters) getRegister(register uint32) uint32 {
	switch register {
	case 0:
		return sysRegisters.R0
	case 1:
		return sysRegisters.R1
	case 2:
		return sysRegisters.R2
	case 3:
		return sysRegisters.R3
	case 4:
		return sysRegisters.R4
	case 5:
		return sysRegisters.R5
	case 6:
		return sysRegisters.R6
	case 7:
		return sysRegisters.R7
	case 8:
		return sysRegisters.R8
	case 9:
		return sysRegisters.R9
	case 10:
		return sysRegisters.R10
	case 11:
		return sysRegisters.R11
	case 12:
		return sysRegisters.R12
	case 13:
		return sysRegisters.R13
	case 14:
		return sysRegisters.R14
	case 15:
		return sysRegisters.R15
	default:
		log.Println("You're are trying to access a non-existent SYS mode register!")
		return 0x0
	}
}

func (sysRegisters *SysRegisters) setRegister(register uint32, value uint32) {
	switch register {
	case 0:
		sysRegisters.R0 = value
	case 1:
		sysRegisters.R1 = value
	case 2:
		sysRegisters.R2 = value
	case 3:
		sysRegisters.R3 = value
	case 4:
		sysRegisters.R4 = value
	case 5:
		sysRegisters.R5 = value
	case 6:
		sysRegisters.R6 = value
	case 7:
		sysRegisters.R7 = value
	case 8:
		sysRegisters.R8 = value
	case 9:
		sysRegisters.R9 = value
	case 10:
		sysRegisters.R10 = value
	case 11:
		sysRegisters.R11 = value
	case 12:
		sysRegisters.R12 = value
	case 13:
		sysRegisters.R13 = value
	case 14:
		sysRegisters.R14 = value
	case 15:
		sysRegisters.R15 = value
	default:
		log.Println("You're are trying to overwrite a non-existent SYS mode register!")
	}
}

func (sysRegisters *SysRegisters) reset(usingBIOS bool) {
	sysRegisters.R0, sysRegisters.R1, sysRegisters.R2, sysRegisters.R3, sysRegisters.R4, sysRegisters.R5 = 0x0, 0x0, 0x0, 0x0, 0x0, 0x0
	sysRegisters.R6, sysRegisters.R7, sysRegisters.R8, sysRegisters.R9, sysRegisters.R10 = 0x0, 0x0, 0x0, 0x0, 0x0
	sysRegisters.R11, sysRegisters.R12, sysRegisters.R14 = 0x0, 0x0, 0x0

	// If not booting from the BIOS
	if !usingBIOS {
		sysRegisters.R13 = 0x03007F00
		sysRegisters.R15 = 0x8000000
		sysRegisters.Cpsr = 0x5F
	} else {
		sysRegisters.R13, sysRegisters.R15 = 0x0, 0x0
		sysRegisters.Cpsr = 0xD3
	}
}
