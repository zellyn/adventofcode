package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/zellyn/adventofcode/stringset"
)

type info struct {
	ingredients stringset.S
	allergens   stringset.S
}

func (i info) String() string {
	return fmt.Sprintf("[ingredients:%v allergens:%v", i.ingredients, i.allergens)
}

func parse(inputs []string) (infos []info, ingredients stringset.S, allergens stringset.S, err error) {
	allergens = stringset.New()
	ingredients = stringset.New()
	for _, input := range inputs {
		if !strings.HasSuffix(input, ")") {
			return nil, nil, nil, fmt.Errorf("weird input: %q", input)
		}
		input = input[:len(input)-1]
		parts := strings.Split(input, " (contains ")
		if len(parts) != 2 {
			return nil, nil, nil, fmt.Errorf("weird input: %q", input)
		}
		ing := stringset.New(strings.Split(parts[0], " ")...)
		all := stringset.New(strings.Split(parts[1], ", ")...)
		ingredients.AddAll(ing)
		allergens.AddAll(all)
		infos = append(infos, info{
			ingredients: ing,
			allergens:   all,
		})
	}

	return infos, ingredients, allergens, nil
}

func part1(inputs []string) (int, string, error) {
	infos, ingredients, allergens, err := parse(inputs)
	if err != nil {
		return 0, "", err
	}

	// Everything is possible at first
	possible := make(map[string]stringset.S)
	for k := range ingredients {
		possible[k] = allergens.Copy()
	}

	for _, inf := range infos {
		for ing := range ingredients {
			for all := range possible[ing] {
				if inf.allergens[all] && !inf.ingredients[ing] {
					delete(possible[ing], all)
				}
			}
		}
	}
	safe := stringset.New()
	for k, v := range possible {
		if len(v) == 0 {
			safe[k] = true
		}
	}

	count := 0
	for _, inf := range infos {
		for ing := range inf.ingredients {
			if safe[ing] {
				count++
			}
		}
	}

	for ing, poss := range possible {
		if len(poss) == 0 {
			delete(possible, ing)
		}
	}

	// known maps allergen to ingredient
	known := make(map[string]string)

	for len(possible) > 0 {
		for ing, poss := range possible {
			if len(poss) == 1 {
				for p := range poss {
					known[p] = ing
				}
				delete(possible, ing)
			}
			for p := range poss {
				_, ok := known[p]
				if ok {
					delete(poss, p)
				}
			}
		}
	}

	var values []string
	for _, s := range allergens.Keys() {
		values = append(values, known[s])
	}

	return count, strings.Join(values, ","), nil
}

func run() error {
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
}
