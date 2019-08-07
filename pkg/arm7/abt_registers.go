package arm7

import "log"

// AbtRegisters define the Registers for the ABT CPU Mode
type AbtRegisters struct {
	// Banked Supervisor Calls (Abt) mode Registers
	R13  uint32
	R14  uint32
	Spsr uint32
}

func (abtRegisters *AbtRegisters) getRegister(register uint32) uint32 {
	switch register {
	case 13:
		return abtRegisters.R13
	case 14:
		return abtRegisters.R14
	default:
		log.Println("You're are trying to access a non-existent ABT register!")
		return 0x0
	}
}

func (abtRegisters *AbtRegisters) setRegister(register uint32, value uint32) {
	switch register {
	case 13:
		abtRegisters.R13 = value
	case 14:
		abtRegisters.R14 = value
	default:
		log.Println("You're are trying to overwrite a non-existent ABT register!")
	}
}

func (abtRegisters *AbtRegisters) reset(usingBIOS bool) {
	abtRegisters.R14, abtRegisters.Spsr = 0x0, 0x0

	// If not booting from the BIOS
	if !usingBIOS {
		abtRegisters.R13 = 0x03007F00
	} else {
		abtRegisters.R13 = 0x0
	}
}
