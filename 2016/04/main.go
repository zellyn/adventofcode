package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
)

type room struct {
	raw          string
	name         string
	sector       int
	seenChecksum string
}

var roomRe = regexp.MustCompile(`^([a-z-]+)-([0-9]+)\[([a-z]+)\]$`)

func parseRoom(input string) (room, error) {
	match := roomRe.FindStringSubmatch(input)
	result := room{
		raw: input,
	}
	if match == nil {
		return room{}, fmt.Errorf("weird room: %q", input)
	}
	result.name = match[1]
	var err error
	result.sector, err = strconv.Atoi(match[2])
	if err != nil {
		return room{}, fmt.Errorf("error parsing room %q: %w", input, err)
	}
	result.seenChecksum = match[3]

	return result, nil
}

func (r room) checksum() string {
	countMap := map[rune]int{}
	for _, ch := range r.name {
		switch {
		case 'a' <= ch && ch <= 'z':
			countMap[ch]++
		case ch == '-':
			// do nothing
		default:
			panic(fmt.Sprintf("weird char %c in room name %q", ch, r.name))
		}
	}
	var counts lettercounts
	for ch, count := range countMap {
		counts = append(counts, lettercount{letter: ch, count: count})
	}
	sort.Sort(counts)
	sum := ""
	for _, lc := range counts {
		sum += string(lc.letter)
	}
	if len(sum) <= 5 {
		return sum
	}
	return sum[:5]
}

func (r room) decrypt() string {
	result := ""
	for _, ch := range r.name {
		if ch == '-' {
			result += " "
		} else {
			letter := rune((int(ch)-'a'+r.sector)%26 + 'a')
			result += string(letter)
		}
	}
	return result
}

type lettercount struct {
	letter rune
	count  int
}

type lettercounts []lettercount

func (a lettercounts) Len() int      { return len(a) }
func (a lettercounts) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a lettercounts) Less(i, j int) bool {
	if a[i].count == a[j].count {
		return a[i].letter < a[j].letter
	}
	return a[i].count > a[j].count
}

func parseRooms(lines []string) ([]room, error) {
	var result []room
	for _, line := range lines {
		r, err := parseRoom(line)
		if err != nil {
			return nil, err
		}
		result = append(result, r)
	}
	return result, nil
}

func validSectorSum(lines []string) (int, error) {
	sum := 0
	rooms, err := parseRooms(lines)
	if err != nil {
		return 0, err
	}
	for _, room := range rooms {
		if room.seenChecksum == room.checksum() {
			sum += room.sector
		}
	}
	return sum, nil
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
