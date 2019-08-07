package arm7

import "log"

// UndRegisters define the Registers for the UND CPU Mode
type UndRegisters struct {
	// Banked Undefined Mode Registers
	R13  uint32
	R14  uint32
	Spsr uint32
}

func (undRegisters *UndRegisters) getRegister(register uint32) uint32 {
	switch register {
	case 13:
		return undRegisters.R13
	case 14:
		return undRegisters.R14
	default:
		log.Println("You're are trying to access a non-existent UND register!")
		return 0x0
	}
}

func (undRegisters *UndRegisters) setRegister(register uint32, value uint32) {
	switch register {
	case 13:
		undRegisters.R13 = value
	case 14:
		undRegisters.R14 = value
	default:
		log.Println("You're are trying to overwrite a non-existent UND register!")
	}
}

func (undRegisters *UndRegisters) reset(usingBIOS bool) {
	undRegisters.R14, undRegisters.Spsr = 0x0, 0x0

	// If not booting from the BIOS
	if !usingBIOS {
		undRegisters.R13 = 0x03007F00
	} else {
		undRegisters.R13 = 0x0
	}
}
