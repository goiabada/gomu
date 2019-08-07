package arm7

import "log"

// FiqRegisters define the Registers for the FIQ CPU Mode
type FiqRegisters struct {
	// Banked Registers for Fast Interrupt Request (FIQ) mode
	R8  uint32
	R9  uint32
	R10 uint32
	R11 uint32
	R12 uint32
	R13 uint32
	R14 uint32
	// Saved Program Status Register - SPSR
	Spsr uint32
}

func (fiqRegisters *FiqRegisters) getRegister(register uint32) uint32 {
	switch register {
	case 8:
		return fiqRegisters.R8
	case 9:
		return fiqRegisters.R9
	case 10:
		return fiqRegisters.R10
	case 11:
		return fiqRegisters.R11
	case 12:
		return fiqRegisters.R12
	case 13:
		return fiqRegisters.R13
	case 14:
		return fiqRegisters.R14
	default:
		log.Println("You're are trying to access a non-existent FIQ register!")
		return 0x0
	}
}

func (fiqRegisters *FiqRegisters) setRegister(register uint32, value uint32) {
	switch register {
	case 8:
		fiqRegisters.R8 = value
	case 9:
		fiqRegisters.R9 = value
	case 10:
		fiqRegisters.R10 = value
	case 11:
		fiqRegisters.R11 = value
	case 12:
		fiqRegisters.R12 = value
	case 13:
		fiqRegisters.R13 = value
	case 14:
		fiqRegisters.R14 = value
	default:
		log.Println("You're are trying to overwrite a non-existent FIQ register!")
	}
}

func (fiqRegisters *FiqRegisters) reset(usingBIOS bool) {
	fiqRegisters.R8, fiqRegisters.R9, fiqRegisters.R10, fiqRegisters.R11 = 0x0, 0x0, 0x0, 0x0
	fiqRegisters.R12, fiqRegisters.R14 = 0x0, 0x0

	// If not booting from the BIOS
	if !usingBIOS {
		fiqRegisters.R13 = 0x03007F00
	} else {
		fiqRegisters.R13 = 0x0
	}
}
