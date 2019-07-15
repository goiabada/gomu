package gba

import "encoding/binary"

func translateLittleEndianInstruction(instruction []byte) uint32 {
	return binary.LittleEndian.Uint32(instruction)
}
