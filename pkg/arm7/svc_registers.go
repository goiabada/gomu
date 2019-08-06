package arm7

import "log"

// SvcRegisters define the registers for the SVC CPU Mode
type SvcRegisters struct {
	// Banked Supervisor Calls (SVC) mode registers
	R13  uint32
	R14  uint32
	Spsr uint32
}

func (svcRegisters *SvcRegisters) getRegister(register uint32) uint32 {
	switch register {
	case 13:
		return svcRegisters.R13
	case 14:
		return svcRegisters.R14
	default:
		log.Println("You're are trying to access a non-existent SVC register!")
		return 0x0
	}
}

func (svcRegisters *SvcRegisters) setRegister(register uint32, value uint32) {
	switch register {
	case 13:
		svcRegisters.R13 = value
	case 14:
		svcRegisters.R14 = value
	default:
		log.Println("You're are trying to overwrite a non-existent SVC register!")
	}
}

func (svcRegisters *SvcRegisters) reset(usingBIOS bool) {
	svcRegisters.R14, svcRegisters.Spsr = 0x0, 0x0

	// If not booting from the BIOS
	if !usingBIOS {
		svcRegisters.R13 = 0x03007FE0
	} else {
		svcRegisters.R13 = 0x0
	}
}
