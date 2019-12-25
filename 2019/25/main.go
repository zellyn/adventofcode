package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/zellyn/adventofcode/2019/intcode"
)

func run() error {
	program, err := intcode.ReadProgram("input")
	if err != nil {
		return err
	}

	input := strings.Join([]string{
		"north",
		"north",
		"north",
		"take mutex",
		"south",
		"south",
		"east",
		// "take escape pod",
		"north",
		"take loom",
		"south",
		"west",
		"south",
		"east",
		"take semiconductor",
		"east",
		"take ornament",
		"north",
		// "take photons",
		"west",
		// "take infinite loop",
		"west",
		// "take giant electromagnet",
		"east",
		"east",
		"south",
		"west",  // Kitchen
		"west",  // Hull Breach
		"west",  // Engineering
		"south", // Crew Quarters
		"north", // Engineering
		"west",  // Hallway
		"take sand",
		"south", // Sick Bay
		"east",  // Warp Drive Maintenance
		"take asterisk",
		"north", // Stables
		"take wreath",
		"south", // Warp Drive Maintenance
		"west",  // Sick Bay
		"north", // Hallway
		"north", // Science Lab
		"take dark matter",
		"east", // Security Checkpoint

		// "drop sand",
		"drop semiconductor",
		// "drop mutex",
		"drop asterisk",
		// "drop wreath",

		"drop dark matter",
		// "drop loom",
		"drop ornament",
		"inv",
		"east",
	}, "\n") + "\n"

	_, writes, err := intcode.RunProgram(program, intcode.FromAscii(input), false)
	fmt.Println(intcode.ToAscii(writes))

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
