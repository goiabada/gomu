package gba

func InitializeROM(romPath string) {
	romData := readROMFile(romPath)
	headers := extractHeaderData(romData[0x000:0x0E3])
	logHeaderData(headers)
}
