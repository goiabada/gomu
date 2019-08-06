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

// RegisterSet defines virtual CPU registers
type RegisterSet struct {
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

	// Banked Registers for Fast Interrupt Request (FIQ) mode
	R8FIQMode  uint32
	R9FIQMode  uint32
	R10FIQMode uint32
	R11FIQMode uint32
	R12FIQMode uint32
	R13FIQMode uint32
	R14FIQMode uint32
	// Saved Program Status Register - SPSR
	SpsrFiq uint32

	// Banked Supervisor Calls (SVC) mode registers
	R13SupervisorMode  uint32
	R14SupervisorMode  uint32
	SpsrSupervisorMode uint32

	// Banked Abort Mode (ABT) registers
	R13AbortMode  uint32
	R14AbortMode  uint32
	SpsrAbortMode uint32

	// Banked Interrupt Mode (IRQ) registers
	R13InterruptMode  uint32
	R14InterruptMode  uint32
	SpsrInterruptMode uint32

	// Banked Undefined Mode registers
	R13UndefinedMode  uint32
	R14UndefinedMode  uint32
	SpsrUndefinedMode uint32
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
	registers.R0, registers.R1, registers.R2, registers.R3, registers.R4, registers.R5 = 0x0, 0x0, 0x0, 0x0, 0x0, 0x0
	registers.R6, registers.R7, registers.R8, registers.R9, registers.R10 = 0x0, 0x0, 0x0, 0x0, 0x0
	registers.R11, registers.R12, registers.R14 = 0x0, 0x0, 0x0

	registers.R8FIQMode, registers.R9FIQMode, registers.R10FIQMode, registers.R11FIQMode = 0x0, 0x0, 0x0, 0x0
	registers.R12FIQMode, registers.R14FIQMode = 0x0, 0x0

	registers.R14SupervisorMode, registers.SpsrSupervisorMode = 0x0, 0x0
	registers.R14AbortMode, registers.SpsrAbortMode = 0x0, 0x0
	registers.R14InterruptMode, registers.SpsrInterruptMode = 0x0, 0x0
	registers.R14UndefinedMode, registers.SpsrUndefinedMode = 0x0, 0x0

	// If not booting from the BIOS
	if !usingBIOS {
		registers.R13, registers.R13FIQMode, registers.R13AbortMode, registers.R13UndefinedMode = 0x03007F00, 0x03007F00, 0x03007F00, 0x03007F00
		registers.R15 = 0x8000000
		registers.R13SupervisorMode = 0x03007FE0
		registers.R13InterruptMode = 0x03007FA0
		registers.Cpsr = 0x5F
	} else {
		registers.R13, registers.R13FIQMode, registers.R13AbortMode, registers.R13UndefinedMode = 0x0, 0x0, 0x0, 0x0
		registers.R15 = 0x0
		registers.R13SupervisorMode = 0x0
		registers.R13InterruptMode = 0x0
		registers.Cpsr = 0xD3
	}

}

func (cpu *CPU) getRegister(register uint32) uint32 {
	switch register {
	case 0:
		return cpu.registers.R0
	case 1:
		return cpu.registers.R1
	case 2:
		return cpu.registers.R2
	case 3:
		return cpu.registers.R3
	case 4:
		return cpu.registers.R4
	case 5:
		return cpu.registers.R5
	case 6:
		return cpu.registers.R6
	case 7:
		return cpu.registers.R7
	case 8:
		if cpu.cpuMode == FIQ {
			return cpu.registers.R8FIQMode
		}
		return cpu.registers.R8
	case 9:
		if cpu.cpuMode == FIQ {
			return cpu.registers.R9FIQMode
		}
		return cpu.registers.R9
	case 10:
		if cpu.cpuMode == FIQ {
			return cpu.registers.R10FIQMode
		}
		return cpu.registers.R10
	case 11:
		if cpu.cpuMode == FIQ {
			return cpu.registers.R11FIQMode
		}
		return cpu.registers.R11
	case 12:
		if cpu.cpuMode == FIQ {
			return cpu.registers.R12FIQMode
		}
		return cpu.registers.R12
	case 13:
		switch cpu.cpuMode {
		case SYS:
			return cpu.registers.R13
		case FIQ:
			return cpu.registers.R13FIQMode
		case SVC:
			return cpu.registers.R13SupervisorMode
		case ABT:
			return cpu.registers.R13AbortMode
		case IRQ:
			return cpu.registers.R13InterruptMode
		case UND:
			return cpu.registers.R13UndefinedMode
		}
	case 14:
		switch cpu.cpuMode {
		case SYS:
			return cpu.registers.R14
		case FIQ:
			return cpu.registers.R14FIQMode
		case SVC:
			return cpu.registers.R14SupervisorMode
		case ABT:
			return cpu.registers.R14AbortMode
		case IRQ:
			return cpu.registers.R14InterruptMode
		case UND:
			return cpu.registers.R14UndefinedMode
		}
	case 15:
		return cpu.registers.R15
	default:
		log.Println("You're are trying to access a non-existent register!")
	}

	return 0x0
}

func (cpu *CPU) setRegister(register uint32, value uint32) {
	switch register {
	case 0:
		cpu.registers.R0 = value
	case 1:
		cpu.registers.R1 = value
	case 2:
		cpu.registers.R2 = value
	case 3:
		cpu.registers.R3 = value
	case 4:
		cpu.registers.R4 = value
	case 5:
		cpu.registers.R5 = value
	case 6:
		cpu.registers.R6 = value
	case 7:
		cpu.registers.R7 = value
	case 8:
		if cpu.cpuMode == FIQ {
			cpu.registers.R8FIQMode = value
		} else {
			cpu.registers.R8 = value
		}
	case 9:
		if cpu.cpuMode == FIQ {
			cpu.registers.R9FIQMode = value
		} else {
			cpu.registers.R9 = value
		}
	case 10:
		if cpu.cpuMode == FIQ {
			cpu.registers.R10FIQMode = value
		} else {
			cpu.registers.R10 = value
		}
	case 11:
		if cpu.cpuMode == FIQ {
			cpu.registers.R11FIQMode = value
		} else {
			cpu.registers.R11 = value
		}
	case 12:
		if cpu.cpuMode == FIQ {
			cpu.registers.R12FIQMode = value
		} else {
			cpu.registers.R12 = value
		}
	case 13:
		switch cpu.cpuMode {
		case SYS:
			cpu.registers.R13 = value
		case FIQ:
			cpu.registers.R13FIQMode = value
		case SVC:
			cpu.registers.R13SupervisorMode = value
		case ABT:
			cpu.registers.R13AbortMode = value
		case IRQ:
			cpu.registers.R13InterruptMode = value
		case UND:
			cpu.registers.R13UndefinedMode = value
		}
	case 14:
		switch cpu.cpuMode {
		case SYS:
			cpu.registers.R14 = value
		case FIQ:
			cpu.registers.R14FIQMode = value
		case SVC:
			cpu.registers.R14SupervisorMode = value
		case ABT:
			cpu.registers.R14AbortMode = value
		case IRQ:
			cpu.registers.R14InterruptMode = value
		case UND:
			cpu.registers.R14UndefinedMode = value
		}
	case 15:
		cpu.registers.R15 = value
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
	cpu.registers.R15 = actualRegisterValue

	switch operation {
	// Branch
	case 0:
	// Branch and Exchange
	case 3:
		cpu.setRegister(14, cpu.getRegister(15)+0x4)
	}
}
