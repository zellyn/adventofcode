package main

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/zellyn/adventofcode/math"
	"github.com/zellyn/adventofcode/util"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf

type pulse bool

const (
	HIGH pulse = true
	LOW  pulse = false
)

func (p pulse) String() string {
	if p == HIGH {
		return "high"
	}
	return "low"
}

type state bool

const (
	ON  state = true
	OFF state = false
)

func (s state) String() string {
	if s == ON {
		return "1"
	}
	return "0"
}

type namedPulse struct {
	source string
	target string
	pulse  pulse
}

// component interface

type component interface {
	getName() string
	getTargets() []string
	connectInput(string)
	reset()
	pulse(namedPulse) []namedPulse
	freshCopy() component
}

// broadcaster

type broadcaster struct {
	targets []string
}

func (b *broadcaster) getName() string {
	return "broadcaster"
}

func (b *broadcaster) getTargets() []string {
	return b.targets
}

func (b *broadcaster) connectInput(name string) {
	// do nothing
}

func (b *broadcaster) reset() {
	// do nothing
}

func (b *broadcaster) pulse(np namedPulse) []namedPulse {
	return sendToAll(np.pulse, b)
}

func (b *broadcaster) freshCopy() component {
	return &broadcaster{targets: slices.Clone(b.targets)}
}

// button

type button struct {
	targets []string
}

func (b *button) getName() string {
	return "button"
}

func (b *button) getTargets() []string {
	return []string{"broadcaster"}
}

func (b *button) connectInput(name string) {
	panic(fmt.Sprintf("Component %q trying to connect to button as an input", name))
}

func (b *button) reset() {
	// do nothing
}

func (b *button) pulse(np namedPulse) []namedPulse {
	return []namedPulse{
		{
			source: "button",
			target: "broadcaster",
			pulse:  LOW,
		},
	}
}

func (b *button) freshCopy() component {
	return &button{targets: slices.Clone(b.targets)}
}

// flipflop

type flipflop struct {
	name    string
	targets []string
	state   state
}

func (f *flipflop) String() string {
	return f.name + ":" + f.state.String()
}

func (f *flipflop) getName() string {
	return f.name
}

func (f *flipflop) getTargets() []string {
	return f.targets
}

func (f *flipflop) connectInput(name string) {
	// do nothing
}

func (f *flipflop) reset() {
	f.state = OFF
}

func (f *flipflop) pulse(np namedPulse) []namedPulse {
	if np.pulse == HIGH {
		return nil
	}

	if f.state == OFF {
		f.state = ON
		return sendToAll(HIGH, f)
	}

	f.state = OFF
	return sendToAll(LOW, f)
}

func (f *flipflop) freshCopy() component {
	return &flipflop{name: f.name, targets: slices.Clone(f.targets), state: OFF}
}

// sink

type sink struct {
	name      string
	highCount int
	lowCount  int
}

func (s *sink) getName() string {
	return s.name
}

func (s *sink) getTargets() []string {
	// no targets
	return nil
}

func (s *sink) connectInput(name string) {
	// do nothing
}

func (s *sink) reset() {
	s.highCount = 0
	s.lowCount = 0
}

func (s *sink) pulse(np namedPulse) []namedPulse {
	if np.pulse == HIGH {
		s.highCount++
	} else {
		s.lowCount++
	}
	return nil
}

func (s *sink) freshCopy() component {
	return &sink{name: s.name}
}

// conjunction

type conjunction struct {
	name    string
	targets []string
	inputs  map[string]pulse
	sent    bool
	last    pulse
}

func (f *conjunction) getName() string {
	return f.name
}

func (c *conjunction) getTargets() []string {
	return c.targets
}

func (c *conjunction) connectInput(name string) {
	c.inputs[name] = LOW
}

func (c *conjunction) reset() {
	for name := range c.inputs {
		c.inputs[name] = LOW
	}
}

func (c *conjunction) pulse(np namedPulse) []namedPulse {
	// printf("conjunction %s got %s pulse from %s\n", c.name, np.pulse, np.source)
	c.inputs[np.source] = np.pulse
	// printf("  state after updating = %v\n", c.inputs)

	for _, value := range c.inputs {
		if value == LOW {
			// printf("  emitting HIGH\n")
			c.last = HIGH
			c.sent = true
			return sendToAll(HIGH, c)
		}
	}

	// printf("  emitting LOW\n")
	c.last = LOW
	c.sent = true
	return sendToAll(LOW, c)
}

func (c *conjunction) freshCopy() component {
	return &conjunction{
		name:    c.name,
		targets: slices.Clone(c.targets),
		inputs:  make(map[string]pulse),
	}
}

// circuit

type circuit struct {
	// For debugging
	debug            bool
	pulseRecord      []string
	button           component
	highCount        int
	lowCount         int
	sinceButtonCount int
	buttonPresses    int

	components   []component
	componentMap map[string]component
	taps         map[namedPulse]tapFn
}

func (c *circuit) reset() {
	c.pulseRecord = c.pulseRecord[:0]
	c.highCount = 0
	c.lowCount = 0
	c.sinceButtonCount = 0
	c.buttonPresses = 0
}

func (c *circuit) pushButton() {
	c.buttonPresses++
	c.sinceButtonCount = 0
	queue := c.button.pulse(namedPulse{})

	for len(queue) > 0 {
		p := queue[0]
		tap := c.taps[namedPulse{target: p.target, pulse: p.pulse}]
		if tap != nil {
			tap(tapInfo{
				np:                               p,
				pulsesBefore:                     c.highCount + c.lowCount,
				pulsesBeforeSinceLastButtonPress: c.sinceButtonCount,
				buttonPresses:                    c.buttonPresses,
			})
		}
		if p.pulse == HIGH {
			c.highCount++
		} else {
			c.lowCount++
		}
		c.sinceButtonCount++
		if c.debug {
			c.pulseRecord = append(c.pulseRecord, fmt.Sprintf("%s -%s-> %s", p.source, p.pulse, p.target))
		}
		target, ok := c.componentMap[p.target]
		if !ok {
			panic("unkown target: " + p.target)
		}
		pulses := target.pulse(p)
		queue = append(queue[1:], pulses...)
	}
}

func (c *circuit) allFlipflops() string {
	var pieces []string

	for _, comp := range c.components {
		if f, ok := comp.(*flipflop); ok {
			pieces = append(pieces, f.name+":"+f.state.String())
		}
	}

	return strings.Join(pieces, " ")
}

func (c *circuit) addComponent(comp component) {
	name := comp.getName()
	c.components = append(c.components, comp)
	c.componentMap[name] = comp
	if name == "button" {
		c.button = comp
	}
}

func (c *circuit) subcircuit(name string) *circuit {
	circ := &circuit{
		componentMap: make(map[string]component),
		taps:         make(map[namedPulse]tapFn),
	}
	circ.addComponent(&broadcaster{
		targets: []string{name},
	})

	toAdd := []string{name}
	for len(toAdd) > 0 {
		name := toAdd[0]
		toAdd = toAdd[1:]
		if circ.componentMap[name] != nil {
			continue
		}
		comp := c.componentMap[name].freshCopy()
		toAdd = append(toAdd, comp.getTargets()...)
		circ.addComponent(comp)
	}

	circ.connectComponents()

	return circ
}

type tapInfo struct {
	np                               namedPulse
	pulsesBefore                     int
	pulsesBeforeSinceLastButtonPress int
	buttonPresses                    int
}

type tapFn func(tapInfo)

func (c *circuit) addTap(name string, wantedPulse pulse, f tapFn) {
	c.taps[namedPulse{target: name, pulse: wantedPulse}] = f
}

func sendToAll(p pulse, source component) []namedPulse {
	targets := source.getTargets()
	res := make([]namedPulse, 0, len(targets))
	for _, target := range targets {
		res = append(res, namedPulse{
			source: source.getName(),
			target: target,
			pulse:  p,
		})
	}
	return res
}

func parseComponent(input string) (component, error) {
	typeAndName, targetString, ok := strings.Cut(input, " -> ")
	if !ok {
		return nil, fmt.Errorf("Weird component spec: %q", input)
	}

	targets := strings.Split(targetString, ", ")

	if typeAndName == "broadcaster" {
		return &broadcaster{
			targets: targets,
		}, nil
	}

	if typeAndName[0] == '%' {
		return &flipflop{
			name:    typeAndName[1:],
			targets: targets,
			state:   OFF,
		}, nil
	}

	if typeAndName[0] == '&' {
		return &conjunction{
			name:    typeAndName[1:],
			targets: targets,
			inputs:  make(map[string]pulse),
		}, nil
	}

	return nil, fmt.Errorf("type %q not implemented", typeAndName)
}

func (c *circuit) connectComponents() {
	for _, comp := range c.components {
		name := comp.getName()
		c.componentMap[name] = comp
		if name == "button" {
			c.button = comp
		}
	}

	if c.button == nil {
		c.addComponent(&button{})
	}

	unseen := make(map[string]bool)
	for _, comp := range c.components {
		for _, target := range comp.getTargets() {
			if c.componentMap[target] == nil {
				unseen[target] = true
			}
		}
	}

	// Handle any targets that aren't defined.
	for name := range unseen {
		c.addComponent(&sink{name: name})
	}

	for name, comp := range c.componentMap {
		for _, targetName := range comp.getTargets() {
			target := c.componentMap[targetName]
			target.connectInput(name)
		}
	}
}

func parse(inputs []string) (*circuit, error) {
	c := &circuit{
		componentMap: make(map[string]component, len(inputs)),
		taps:         make(map[namedPulse]tapFn),
	}
	var err error

	c.components, err = util.MapE(inputs, parseComponent)
	if err != nil {
		return c, err
	}

	c.connectComponents()
	return c, nil
}

func part1(inputs []string) (int, error) {
	circ, err := parse(inputs)
	if err != nil {
		return 0, err
	}
	for i := 0; i < 1000; i++ {
		circ.pushButton()
	}
	return circ.highCount * circ.lowCount, nil
}

func part2(inputs []string) (int, error) {
	circ, err := parse(inputs)
	if err != nil {
		return 0, err
	}
	subcircuitNames := circ.componentMap["broadcaster"].getTargets()
	subcircuits := util.Map(subcircuitNames, func(s string) *circuit { return circ.subcircuit(s) })

	cycles := make([]int, len(subcircuitNames))

	for i, sc := range subcircuits {
		sc.addTap("rx", LOW, func(ti tapInfo) {
			cycles[i] = ti.buttonPresses
		})

		for cycles[i] == 0 {
			sc.pushButton()
		}
	}

	return math.MultiLCM(cycles...), nil
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
