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

	/*
		old_input := strings.Join([]string{
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
	*/

	maybe := func(thing string) string {
		take := []string{
			// definitely
			"fixed point",
			"polygon",

			// definitely not
			// "hologram",
			// "tambourine",

			"boulder",
			// "fuel cell",
			"manifold",
			// "wreath",
		}

		for _, t := range take {
			if t == thing {
				return "take " + thing
			}
		}
		return ""
	}
	input := strings.Join([]string{
		"south",
		"south",
		maybe("tambourine"),
		"north",
		"north",
		"west",
		"south",
		maybe("polygon"),
		// "south",
		// "take infinite loop",
		// "north",
		"north",
		"east",
		"north",
		"west",
		maybe("boulder"),
		// "south",
		// "take escape pod",
		// "north",
		"east",
		"north",
		maybe("manifold"),
		"north",
		// "north",
		// "take photons",
		// "south",
		maybe("hologram"),
		"south",
		"west",
		maybe("fuel cell"),
		"south",
		"east",
		// "take giant electromagnet",
		"south",
		maybe("fixed point"),
		"north",
		"west",
		"north",
		"north",
		maybe("wreath"),
		"east",
		// "take molten lava",
		"east",
		"inv",
		"north",
	}, "\n") + "\n"

	/*

	   Need fixed point: without it, everything else is too light.


	*/
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
