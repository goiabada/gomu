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
	r0  uint32
	r1  uint32
	r2  uint32
	r3  uint32
	r4  uint32
	r5  uint32
	r6  uint32
	r7  uint32
	r8  uint32
	r9  uint32
	r10 uint32
	r11 uint32
	r12 uint32
	// Stack Pointer - SP
	r13 uint32
	// Link Register - LP
	r14 uint32
	// Program Counter - PC
	r15 uint32
	// Current Program Status Register - CPSR
	cpsr uint32

	// Banked Registers for Fast Interrupt Request (FIQ) mode
	r8FIQMode  uint32
	r9FIQMode  uint32
	r10FIQMode uint32
	r11FIQMode uint32
	r12FIQMode uint32
	r13FIQMode uint32
	r14FIQMode uint32
	// Saved Program Status Register - SPSR
	spsrFiq uint32

	// Banked Supervisor Calls (SVC) mode registers
	r13SupervisorMode  uint32
	r14SupervisorMode  uint32
	spsrSupervisorMode uint32

	// Banked Abort Mode (ABT) registers
	r13AbortMode  uint32
	r14AbortMode  uint32
	spsrAbortMode uint32

	// Banked Interrupt Mode (IRQ) registers
	r13InterruptMode  uint32
	r14InterruptMode  uint32
	spsrInterruptMode uint32

	// Banked Undefined Mode registers
	r13UndefinedMode  uint32
	r14UndefinedMode  uint32
	spsrUndefinedMode uint32
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
	registers.r0, registers.r1, registers.r2, registers.r3, registers.r4, registers.r5 = 0x0, 0x0, 0x0, 0x0, 0x0, 0x0
	registers.r6, registers.r7, registers.r8, registers.r9, registers.r10 = 0x0, 0x0, 0x0, 0x0, 0x0
	registers.r11, registers.r12, registers.r14 = 0x0, 0x0, 0x0

	registers.r8FIQMode, registers.r9FIQMode, registers.r10FIQMode, registers.r11FIQMode = 0x0, 0x0, 0x0, 0x0
	registers.r12FIQMode, registers.r14FIQMode = 0x0, 0x0

	registers.r14SupervisorMode, registers.spsrSupervisorMode = 0x0, 0x0
	registers.r14AbortMode, registers.spsrAbortMode = 0x0, 0x0
	registers.r14InterruptMode, registers.spsrInterruptMode = 0x0, 0x0
	registers.r14UndefinedMode, registers.spsrUndefinedMode = 0x0, 0x0

	// If not booting from the BIOS
	if !usingBIOS {
		registers.r13, registers.r13FIQMode, registers.r13AbortMode, registers.r13UndefinedMode = 0x03007F00, 0x03007F00, 0x03007F00, 0x03007F00
		registers.r15 = 0x8000000
		registers.r13SupervisorMode = 0x03007FE0
		registers.r13InterruptMode = 0x03007FA0
		registers.cpsr = 0x5F
	} else {
		registers.r13, registers.r13FIQMode, registers.r13AbortMode, registers.r13UndefinedMode = 0x0, 0x0, 0x0, 0x0
		registers.r15 = 0x0
		registers.r13SupervisorMode = 0x0
		registers.r13InterruptMode = 0x0
		registers.cpsr = 0xD3
	}

}

func (cpu *CPU) getRegister(register uint32) uint32 {
	switch register {
	case 0:
		return cpu.registers.r0
	case 1:
		return cpu.registers.r1
	case 2:
		return cpu.registers.r2
	case 3:
		return cpu.registers.r3
	case 4:
		return cpu.registers.r4
	case 5:
		return cpu.registers.r5
	case 6:
		return cpu.registers.r6
	case 7:
		return cpu.registers.r7
	case 8:
		if cpu.cpuMode == FIQ {
			return cpu.registers.r8FIQMode
		}
		return cpu.registers.r8
	case 9:
		if cpu.cpuMode == FIQ {
			return cpu.registers.r9FIQMode
		}
		return cpu.registers.r9
	case 10:
		if cpu.cpuMode == FIQ {
			return cpu.registers.r10FIQMode
		}
		return cpu.registers.r10
	case 11:
		if cpu.cpuMode == FIQ {
			return cpu.registers.r11FIQMode
		}
		return cpu.registers.r11
	case 12:
		if cpu.cpuMode == FIQ {
			return cpu.registers.r12FIQMode
		}
		return cpu.registers.r12
	case 13:
		switch cpu.cpuMode {
		case SYS:
			return cpu.registers.r13
		case FIQ:
			return cpu.registers.r13FIQMode
		case SVC:
			return cpu.registers.r13SupervisorMode
		case ABT:
			return cpu.registers.r13AbortMode
		case IRQ:
			return cpu.registers.r13InterruptMode
		case UND:
			return cpu.registers.r13UndefinedMode
		}
	case 14:
		switch cpu.cpuMode {
		case SYS:
			return cpu.registers.r14
		case FIQ:
			return cpu.registers.r14FIQMode
		case SVC:
			return cpu.registers.r14SupervisorMode
		case ABT:
			return cpu.registers.r14AbortMode
		case IRQ:
			return cpu.registers.r14InterruptMode
		case UND:
			return cpu.registers.r14UndefinedMode
		}
	case 15:
		return cpu.registers.r15
	default:
		log.Println("You're are trying to access a non-existent register!")
	}

	return 0x0
}

func (cpu *CPU) setRegister(register uint32, value uint32) {
	switch register {
	case 0:
		cpu.registers.r0 = value
	case 1:
		cpu.registers.r1 = value
	case 2:
		cpu.registers.r2 = value
	case 3:
		cpu.registers.r3 = value
	case 4:
		cpu.registers.r4 = value
	case 5:
		cpu.registers.r5 = value
	case 6:
		cpu.registers.r6 = value
	case 7:
		cpu.registers.r7 = value
	case 8:
		if cpu.cpuMode == FIQ {
			cpu.registers.r8FIQMode = value
		} else {
			cpu.registers.r8 = value
		}
	case 9:
		if cpu.cpuMode == FIQ {
			cpu.registers.r9FIQMode = value
		} else {
			cpu.registers.r9 = value
		}
	case 10:
		if cpu.cpuMode == FIQ {
			cpu.registers.r10FIQMode = value
		} else {
			cpu.registers.r10 = value
		}
	case 11:
		if cpu.cpuMode == FIQ {
			cpu.registers.r11FIQMode = value
		} else {
			cpu.registers.r11 = value
		}
	case 12:
		if cpu.cpuMode == FIQ {
			cpu.registers.r12FIQMode = value
		} else {
			cpu.registers.r12 = value
		}
	case 13:
		switch cpu.cpuMode {
		case SYS:
			cpu.registers.r13 = value
		case FIQ:
			cpu.registers.r13FIQMode = value
		case SVC:
			cpu.registers.r13SupervisorMode = value
		case ABT:
			cpu.registers.r13AbortMode = value
		case IRQ:
			cpu.registers.r13InterruptMode = value
		case UND:
			cpu.registers.r13UndefinedMode = value
		}
	case 14:
		switch cpu.cpuMode {
		case SYS:
			cpu.registers.r14 = value
		case FIQ:
			cpu.registers.r14FIQMode = value
		case SVC:
			cpu.registers.r14SupervisorMode = value
		case ABT:
			cpu.registers.r14AbortMode = value
		case IRQ:
			cpu.registers.r14InterruptMode = value
		case UND:
			cpu.registers.r14UndefinedMode = value
		}
	case 15:
		cpu.registers.r15 = value
	default:
		log.Println("You're are trying to overwrite a non-existent register!")
	}
}

// BranchWithLink executes correspondent CPU instruction
func BranchWithLink(instruction []byte, registers *RegisterSet) {
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
	registers.r15 += 8 + 4*offset

	switch opcode {
	// Branch operation
	case 0x0:
	// Branch with link operation
	case 0x1:
		registers.r14 = registers.r15 + 4
	}

}

// BranchAndExchange executes correspondent CPU instruction
func BranchAndExchange(instruction []byte, registers *RegisterSet) {
	fixedInstruction := translateLittleEndianInstruction(instruction)

	sourceRegister := fixedInstruction & 0xF
	operation := (fixedInstruction >> 4) & 0x1

	actualRegisterValue := registers.getRegister(sourceRegister)
	registers.r15 = actualRegisterValue

	switch operation {
	// Branch
	case 0:
	// Branch and Exchange
	case 3:
		registers.r14 = registers.r15 + 0x4
	}

}
