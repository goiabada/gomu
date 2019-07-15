package gba

// InitializeROM loads the rom file and extract it's headers
func InitializeROM(romPath string) {
	romData := readROMFile(romPath)
	headers := extractHeaderData(romData[0x000:0x0E3])
	logHeaderData(headers)
	registers := new(registerSet)
	branchWithLink(headers.romEntryPoint, registers)
	branchAndExchange([]byte{0xE5, 0x0, 0x81, 0xE5}, registers)
}
