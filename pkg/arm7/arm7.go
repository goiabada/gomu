package arm7

import (
	"log"
)

// CPU defines the processor with all relevant info and registers
type CPU struct {
	cpuMode         int8
	instructionMode int8
	registers       RegisterSet
}

// RegisterSet envelops all different registers from every CPU Mode
type RegisterSet struct {
	sysRegisters SysRegisters
	fiqRegisters FiqRegisters
	svcRegisters SvcRegisters
	abtRegisters AbtRegisters
	irqRegisters IrqRegisters
	undRegisters UndRegisters
}

// Constants for defining the CPU modes
const (
	USR int8 = iota
	SYS
	FIQ
	SVC
	ABT
	IRQ
	UND
)

// Constants for defining the current ARM instruction mode
const (
	ARM int8 = iota
	THUMB
)

func (registers *RegisterSet) reset(usingBIOS bool) {
	registers.sysRegisters.reset(usingBIOS)
	registers.fiqRegisters.reset(usingBIOS)
	registers.svcRegisters.reset(usingBIOS)
	registers.abtRegisters.reset(usingBIOS)
	registers.irqRegisters.reset(usingBIOS)
	registers.undRegisters.reset(usingBIOS)
}

func (cpu *CPU) getRegister(register uint32) uint32 {
	switch register {
	case 0, 1, 2, 3, 4, 5, 6, 7, 15:
		return cpu.registers.sysRegisters.getRegister(register)
	case 8, 9, 10, 11, 12:
		if cpu.cpuMode == FIQ {
			return cpu.registers.fiqRegisters.getRegister(register)
		}
		return cpu.registers.sysRegisters.getRegister(register)
	case 13, 14:
		switch cpu.cpuMode {
		case SYS:
			return cpu.registers.sysRegisters.getRegister(register)
		case FIQ:
			return cpu.registers.fiqRegisters.getRegister(register)
		case SVC:
			return cpu.registers.svcRegisters.getRegister(register)
		case ABT:
			return cpu.registers.abtRegisters.getRegister(register)
		case IRQ:
			return cpu.registers.irqRegisters.getRegister(register)
		case UND:
			return cpu.registers.undRegisters.getRegister(register)
		}
	default:
		log.Println("You're are trying to access a non-existent register!")
	}
	return 0x0
}

func (cpu *CPU) setRegister(register uint32, value uint32) {
	switch register {
	case 0, 1, 2, 3, 4, 5, 6, 7, 15:
		cpu.registers.sysRegisters.setRegister(register, value)
	case 8, 9, 10, 11, 12:
		if cpu.cpuMode == FIQ {
			cpu.registers.fiqRegisters.setRegister(register, value)
		}
		cpu.registers.sysRegisters.setRegister(register, value)
	case 13, 14:
		switch cpu.cpuMode {
		case SYS:
			cpu.registers.sysRegisters.setRegister(register, value)
		case FIQ:
			cpu.registers.fiqRegisters.setRegister(register, value)
		case SVC:
			cpu.registers.svcRegisters.setRegister(register, value)
		case ABT:
			cpu.registers.abtRegisters.setRegister(register, value)
		case IRQ:
			cpu.registers.irqRegisters.setRegister(register, value)
		case UND:
			cpu.registers.undRegisters.setRegister(register, value)
		}
	default:
		log.Println("You're are trying to overwrite a non-existent register!")
	}
}

// BranchWithLink executes correspondent CPU instruction
func (cpu *CPU) BranchWithLink(instruction []byte) {
	// First, we correct the byte order of the opcode
	fixedOInstruction := translateLittleEndianInstruction(instruction)

	// We grab the offset
	offset := fixedOInstruction & 0xFFFFFF

	// If the offset is negative (24th offset bit is 1), do a complement of 2.
	if offset>>23&0x1 == 0x1 {
		offset = -offset
	}

	// For the opcode, we shift right 24 bits and grab just the first bit
	opcode := (fixedOInstruction >> 24) & 0x1
	/// Update the program counter
	cpu.setRegister(15, cpu.getRegister(15)+8+4*offset)

	switch opcode {
	// Branch operation
	case 0x0:
	// Branch with link operation
	case 0x1:
		cpu.setRegister(14, cpu.getRegister(15)+4)
	}

}

// BranchAndExchange executes correspondent CPU instruction
func (cpu *CPU) BranchAndExchange(instruction []byte) {
	fixedInstruction := translateLittleEndianInstruction(instruction)

	sourceRegister := fixedInstruction & 0xF
	operation := (fixedInstruction >> 4) & 0x1

	actualRegisterValue := cpu.getRegister(sourceRegister)
	cpu.setRegister(15, actualRegisterValue)

	switch operation {
	// Branch
	case 0:
	// Branch and Exchange
	case 3:
		cpu.setRegister(14, cpu.getRegister(15)+0x4)
	}
}
