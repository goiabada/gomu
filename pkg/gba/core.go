package gba

import "../arm7"

// InitializeROM loads the rom file and extract it's headers
func InitializeROM(romPath string) {
	romData := readROMFile(romPath)
	headers := extractHeaderData(romData[0x000:0x0E3])
	logHeaderData(headers)

	cpu := new(arm7.CPU)
	cpu.Registers.Reset(false)
	// cpu.BranchWithLink(headers.romEntryPoint)
	// cpu.BranchAndExchange([]byte{0xE5, 0x0, 0x81, 0xE5})
}
