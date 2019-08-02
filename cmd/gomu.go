package main

import (
	"flag"
	"os"
	"github.com/goiabada/gomu/pkg/gba"
)

type flags struct {
	cartridge string
}

func parseFlags() flags {
	cartridge := flag.String("cartridge", "", "path for GBA cartridge ROM file")
	flag.Parse()
	return flags{*cartridge}
}

func main() {
	flags := parseFlags()
	if flags.cartridge == "" {
		flag.Usage()
		os.Exit(0)
	}
	gba.InitializeROM(flags.cartridge)
}
