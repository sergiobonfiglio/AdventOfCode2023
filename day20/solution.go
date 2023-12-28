package main

import (
	"AdventOfCode2023/utils"
	"fmt"
	"strings"
)

func part1(input string) any {
	modules, broadcaster := initModules(input)

	highsCount := 0
	lowCounts := 0
	for i := 0; i < 1000; i++ {
		buttonPulse := Pulse{
			val:  L,
			from: "button",
			to:   broadcaster.label(),
		}
		//fmt.Println(&buttonPulse)
		curPulses := broadcaster.process(buttonPulse)

		lowCounts++ // button pulse

		for len(curPulses) > 0 {
			var nextPulses []*Pulse
			for _, pulse := range curPulses {

				//fmt.Println(pulse)

				if pulse.val == H {
					highsCount++
				} else {
					lowCounts++
				}

				nextPulses = append(nextPulses, modules[pulse.to].process(*pulse)...)
			}
			curPulses = nextPulses
		}
		//fmt.Println("----------")
	}
	fmt.Printf("H: %d, L:%d\n", highsCount, lowCounts)

	return highsCount * lowCounts
}

func part2(input string) any {
	modules, broadcaster := initModules(input)

	buttonPresses := 0

	var rxInputs []string
	var rxInputs2 []string
	rxInputs = append(rxInputs, modules["rx"].inputs()...)
	for _, rxInput := range rxInputs {
		rxInputs2 = append(rxInputs2, modules[rxInput].inputs()...)
	}

	//for _, s := range rxInputs2 {
	//	fmt.Printf("input2: %s\n", s)
	//}

	cycles := map[string]int{}
	allCyclesFound := false

	for !allCyclesFound {
		buttonPulse := Pulse{
			val:  L,
			from: "button",
			to:   broadcaster.label(),
		}
		buttonPresses++
		curPulses := broadcaster.process(buttonPulse)

		for len(curPulses) > 0 {
			var nextPulses []*Pulse

			for _, pulse := range curPulses {

				for _, rxIn := range rxInputs2 {
					if pulse.val == H && pulse.from == rxIn {
						if _, found := cycles[pulse.from]; !found {
							cycles[pulse.from] = buttonPresses
						}
					}
				}

				allCyclesFound = true
				for _, rxIn := range rxInputs2 {
					if _, found := cycles[rxIn]; !found {
						allCyclesFound = false
						break
					}
				}
				nextPulses = append(nextPulses, modules[pulse.to].process(*pulse)...)
			}
			curPulses = nextPulses
		}
	}

	var values []int64
	for _, v := range cycles {
		values = append(values, int64(v))
		//fmt.Printf("%s: %d\n", in, v)
	}
	lcm := utils.LCM(values[0], values[1], values[2:]...)

	return lcm
}

func initModules(input string) (map[string]Module, Module) {
	modules := map[string]Module{}
	var broadcaster Module
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		parts := strings.Split(line, " -> ")

		outputs := strings.Split(parts[1], ", ")

		//var module Module
		if parts[0] == "broadcaster" {
			broadcaster = newBroadcast(parts[0], outputs)
		} else {
			modType, modName := rune(parts[0][0]), parts[0][1:]
			modules[modName] = newModule(modType, modName, outputs)
		}

		//accounts for output-only modules
		for _, output := range outputs {
			if _, found := modules[output]; !found {
				modules[output] = newBroadcast(output, []string{})
			}
		}
	}

	for _, module := range modules {
		for _, out := range module.outputs() {
			modules[out].addInput(module.label())
		}
	}
	return modules, broadcaster
}

type Pulse struct {
	val  PulseVal
	from string
	to   string
}

var boolIntMap = map[bool]int{false: 0, true: 1}

func (p *Pulse) String() string {
	return fmt.Sprintf("%s %d -> %s", p.from, boolIntMap[p.val], p.to)
}
func newPulse(val PulseVal, from string, to string) *Pulse {
	return &Pulse{
		val:  val,
		from: from,
		to:   to,
	}
}

type PulseVal = bool

const (
	H PulseVal = true
	L PulseVal = false
)

func newModule(modType rune, name string, outputs []string) Module {
	if modType == '%' {
		return newFlipFlop(name, outputs)
	} else if modType == '&' {
		return newConjunction(name, outputs)
	} else {
		return newBroadcast(name, outputs)
	}
}

type Module interface {
	label() string
	outputs() []string
	process(pulse Pulse) []*Pulse
	addInput(m string)
	inputs() []string
	getType() rune
}

type CommonModule struct {
	_label   string
	_outputs []string
	_inputs  []string
	_type    rune
}

func newCommonModule(label string, outputs []string, _type rune) CommonModule {
	return CommonModule{
		_label:   label,
		_outputs: outputs,
		_type:    _type,
	}

}
func (c *CommonModule) label() string {
	return c._label
}

func (c *CommonModule) outputs() []string {
	return c._outputs
}

func (c *CommonModule) addInput(m string) {
	c._inputs = append(c._inputs, m)
}

func (c *CommonModule) inputs() []string {
	return c._inputs
}

func (c *CommonModule) getType() rune {
	return c._type
}

type BroadcastModule struct {
	CommonModule
}

var _ Module = new(BroadcastModule)

func newBroadcast(label string, outputs []string) *BroadcastModule {
	return &BroadcastModule{newCommonModule(label, outputs, 'b')}
}
func (b *BroadcastModule) process(pulse Pulse) []*Pulse {
	var outs []*Pulse
	for _, out := range b._outputs {
		outs = append(outs, &Pulse{
			val:  pulse.val,
			from: b._label,
			to:   out,
		})
	}

	return outs
}

type FlipFlop struct {
	CommonModule
	isOn bool
}

var _ Module = new(FlipFlop)

func newFlipFlop(label string, outputs []string) *FlipFlop {
	return &FlipFlop{
		CommonModule: newCommonModule(label, outputs, '%'),
		isOn:         false,
	}
}

func (f *FlipFlop) process(pulse Pulse) []*Pulse {
	var outs []*Pulse
	if pulse.val == L {
		var res PulseVal
		if f.isOn {
			res = L
		} else {
			res = H
		}
		f.isOn = !f.isOn

		for _, out := range f._outputs {
			outs = append(outs, &Pulse{
				val:  res,
				from: f._label,
				to:   out,
			})
		}
	}
	return outs
}

type Conjunction struct {
	CommonModule
	lastPulse map[string]PulseVal
}

var _ Module = new(Conjunction)

func newConjunction(label string, outputs []string) *Conjunction {
	lastPulse := map[string]PulseVal{}

	return &Conjunction{
		CommonModule: newCommonModule(label, outputs, '&'),
		lastPulse:    lastPulse,
	}
}

func (c *Conjunction) addInput(m string) {
	c.CommonModule.addInput(m)
	c.lastPulse[m] = L
}

func (c *Conjunction) process(pulse Pulse) []*Pulse {

	var outs []*Pulse

	c.lastPulse[pulse.from] = pulse.val
	allHigh := true
	for _, p := range c.lastPulse {
		if p == L {
			allHigh = false
			break
		}
	}

	retVal := H
	if allHigh {
		retVal = L
	}

	for _, out := range c._outputs {
		outs = append(outs, &Pulse{
			val:  retVal,
			from: c._label,
			to:   out,
		})
	}

	return outs
}
