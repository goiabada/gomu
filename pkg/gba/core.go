package gba

import "github.com/goiabada/gomu/pkg/arm7"

// InitializeROM loads the rom file and extract it's headers
func InitializeROM(romPath string) {
	romData := readROMFile(romPath)
	headers := extractHeaderData(romData[0x000:0x0E3])
	logHeaderData(headers)
	registers := new(arm7.RegisterSet)
	arm7.BranchWithLink(headers.romEntryPoint, registers)
	arm7.BranchAndExchange([]byte{0xE5, 0x0, 0x81, 0xE5}, registers)
}
