package gba

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
)

type cartridgeHeader struct {
	/*
		Space for a single 32bit ARM opcode that redirects to the actual startaddress of the cartridge,
		this should be usually a "B <start>" instruction.
	*/
	romEntryPoint []byte
	/*
		Contains the Nintendo logo which is displayed during the boot procedure. Cartridge won't work if this data is missing or modified.
		A copy of the compression data is stored in the BIOS, the GBA will compare this data and lock-up itself if the BIOS data
		isn't exactly the same as in the cartridge.
	*/
	nintendoLogo []byte
	/*
		Space for the game title, padded with 00h (if less than 12 chars).
	*/
	gameTitle []byte
	/*
		The first character (U) is usually "A" or "B", the second/third characters (TT) are usually an
		abbreviation of the game title and the fourth character indicates destination/language.
	*/
	gameCode []byte
	/*
		Identifies the (commercial) developer. For example, "01"=Nintendo.
		This guuy should be treated as binary always, I think.
	*/
	makerCode []byte
	/*
		Must be 96h (150 dec). Required
	*/
	fixedValue byte
	/*
		Identifies the required hardware. Should be 00h for current GBA models.
	*/
	mainUnitCode byte
	/*
		Normally, this entry should be zero
	*/
	deviceType byte
	/*
		Version number of the game. Usually zero.
	*/
	softwareVersion byte
	/*
		Header checksum, cartridge won't work if incorrect. Calculate as such:
		chk=0:for i=0A0h to 0BCh:chk=chk-[i]:next:chk=(chk-19h) and 0FFh
	*/
	complementCheck byte
	/*
		This entry is used only if the GBA has been booted by using Normal or Multiplay transfer mode.
		Typically deposit a ARM-32bit "B <start>" branch opcode at this location, which is pointing to your actual initialization procedure.
	*/
	ramEntryPoint []byte
	/*
		Indicaties the used multiboot transfer mode: 01h -> Joybus mode, 02h -> Normal mode, 03h -> Multiplay mode.
		Initially zero.
	*/
	bootMode byte
	/*
		If the GBA has been booted in Normal or Multiplay mode, this byte becomes overwritten by the slave ID number
		of the local GBA (that'd be always 01h for normal mode). Initially as 00h.
	*/
	slaveID byte
	/*
		If the GBA has been booted by using Joybus transfer mode, then the entry point is located at this address.
	*/
	joybusEntryPoint []byte
}

func readROMFile(romPath string) []byte {
	absRomPath, _ := filepath.Abs(romPath)

	log.Println("Loading rom file ...")
	romFile, err := os.Open(absRomPath)
	if err != nil {
		log.Fatal(err)
	}

	stats, statsErr := romFile.Stat()
	if statsErr != nil {
		log.Fatal(statsErr)
	}

	size := stats.Size()
	log.Println("ROM with", size/1024, "KB loaded")

	romBytes := make([]byte, size)
	reader := bufio.NewReader(romFile)
	_, err = reader.Read(romBytes)

	return romBytes
}

func extractHeaderData(romData []byte) *cartridgeHeader {
	header := new(cartridgeHeader)

	header.romEntryPoint = romData[0x000:0x003]
	header.gameTitle = romData[0x0A0:0x0AB]
	header.gameCode = romData[0x0AC:0x0AF]
	header.makerCode = romData[0x0B0:0x0B1]
	header.fixedValue = romData[0x0B2]
	header.mainUnitCode = romData[0x0B3]
	header.deviceType = romData[0x0B4]
	header.softwareVersion = romData[0x0BC]
	header.complementCheck = romData[0x0BD]
	header.ramEntryPoint = romData[0x0C0:0x0C3]
	header.bootMode = romData[0x0C4]
	header.slaveID = romData[0x0C5]
	header.joybusEntryPoint = romData[0x0E0:0x0E3]

	return header
}

func logHeaderData(header *cartridgeHeader) {
	log.Println("Reading cartridge readers...")
	log.Println("ROM Entry Point: ", header.romEntryPoint)
	log.Println("Game Title: ", string(header.gameTitle))
	log.Println("Game Code: ", string(header.gameCode))
	log.Println("Maker Code: ", string(header.makerCode))
	log.Println("Fixed Value: ", header.fixedValue)
	log.Println("Main unit code: ", header.mainUnitCode)
	log.Println("Device type: ", header.deviceType)
	log.Println("Software version: ", header.softwareVersion)
	log.Println("Complement Check: ", header.complementCheck)
	log.Println("RAM Entry Point: ", header.ramEntryPoint)
	log.Println("Boot mode: ", header.bootMode)
	log.Println("Slave ID Number: ", header.slaveID)
	log.Println("Joybus Entry Point: ", header.joybusEntryPoint)
}
