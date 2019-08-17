package arm7

import "log"

// IrqRegisters define the Registers for the IRQ CPU Mode
type IrqRegisters struct {
	// Banked Interrupt Mode (IRQ) Registers
	R13  uint32
	R14  uint32
	Spsr uint32
}

func (irqRegisters *IrqRegisters) getRegister(register uint32) uint32 {
	switch register {
	case 13:
		return irqRegisters.R13
	case 14:
		return irqRegisters.R14
	default:
		log.Println("You're are trying to access a non-existent IRQ register!")
		return 0x0
	}
}

func (irqRegisters *IrqRegisters) setRegister(register uint32, value uint32) {
	switch register {
	case 13:
		irqRegisters.R13 = value
	case 14:
		irqRegisters.R14 = value
	default:
		log.Println("You're are trying to overwrite a non-existent IRQ register!")
	}
}

func (irqRegisters *IrqRegisters) reset(usingBIOS bool) {
	irqRegisters.R14, irqRegisters.Spsr = 0x0, 0x0

	// If not booting from the BIOS
	if !usingBIOS {
		irqRegisters.R13 = 0x03007FA0
	} else {
		irqRegisters.R13 = 0x0
	}
}
