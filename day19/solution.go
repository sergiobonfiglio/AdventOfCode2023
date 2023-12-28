package main

import (
	"fmt"
	"strconv"
	"strings"
)

func part1(input string) any {
	workflows, parts := parse(input)
	wfRes := map[string][]Part{}
	wfRes["A"] = []Part{}
	wfRes["R"] = []Part{}

	for _, part := range parts {
		curr := workflows["in"]
		for curr != nil {
			for ri := 0; ri < len(curr.rules); ri++ {
				rule := curr.rules[ri]

				nextWf := rule.eval(part)

				if nextWf == nil {
					continue
				} else if *nextWf == "A" || *nextWf == "R" {
					wfRes[*nextWf] = append(wfRes[*nextWf], part)
					curr = nil
					break
				} else {
					curr = workflows[*nextWf]
					break
				}

			}
		}

	}

	sum := 0
	for _, part := range wfRes["A"] {
		for _, v := range part {
			sum += v
		}
	}

	return sum
}

func part2(input string) any {
	workflows, _ := parse(input)

	tree := buildTree(workflows)

	_ = tree

	acceptedNodes := tree.root.DFS("A")
	_ = acceptedNodes
	cmpConds := make([]*CompCondition, len(acceptedNodes))
	for i, node := range acceptedNodes {
		curr := node
		var next *Node
		cmpConds[i] = newCompCondition()
		for curr != nil {
			if next != nil {
				if curr.cond != nil {
					tmp := curr.cond
					if curr.falseBranch == next {
						tmp = tmp.Negate()
					}

					cmpConds[i].and(tmp)
				}
			}

			next = curr
			curr = curr.prev
		}
	}

	sum := int64(0)
	for _, cmpCond := range cmpConds {
		possibilities := cmpCond.getPossible()
		sum += possibilities
	}

	return sum
}

func buildTree(workflows map[string]*Workflow) *Tree {

	startWF := workflows["in"]

	workflows["A"] = &Workflow{
		name:  "A",
		rules: nil,
	}
	workflows["R"] = &Workflow{
		name:  "R",
		rules: nil,
	}

	root := wf2Node(nil, startWF.name, startWF.rules, workflows)

	return &Tree{root}
}

func wf2Node(prev *Node, wfName string, rules []*Rule, workflows map[string]*Workflow) *Node {

	if wfName == "A" || wfName == "R" {
		return &Node{
			wfName:      wfName,
			ruleIx:      0,
			prev:        prev,
			cond:        nil,
			trueBranch:  nil,
			falseBranch: nil,
		}
	}
	if len(rules) < 1 || rules[0] == nil {
		panic(-1)
	}

	wfNode := &Node{
		wfName:      wfName,
		ruleIx:      0,
		prev:        prev,
		cond:        rules[0].condition,
		trueBranch:  nil,
		falseBranch: nil,
	}

	trueWf := workflows[rules[0].trueWF]
	if trueWf == nil {
		panic(-1)
	}
	trueNode := wf2Node(wfNode, rules[0].trueWF, trueWf.rules, workflows)
	wfNode.trueBranch = trueNode

	if len(rules) > 1 {
		wfNode.falseBranch = wf2Node(wfNode, wfName+".", rules[1:], workflows)
	} else {
		wfNode.falseBranch = &Node{
			wfName:      "R",
			ruleIx:      0,
			prev:        wfNode,
			cond:        nil,
			trueBranch:  nil,
			falseBranch: nil,
		}
	}

	return wfNode
}

type Tree struct {
	root *Node
}

func (n *Node) DFS(name string) []*Node {

	if n.wfName == name {
		return []*Node{n}
	}

	var res []*Node
	if n.trueBranch != nil {
		res = append(res, n.trueBranch.DFS(name)...)
	}
	if n.falseBranch != nil {
		res = append(res, n.falseBranch.DFS(name)...)
	}
	return res
}

type Node struct {
	wfName      string
	ruleIx      int
	prev        *Node
	cond        *Condition
	trueBranch  *Node
	falseBranch *Node
}

func parse(input string) (map[string]*Workflow, []Part) {
	workflows := map[string]*Workflow{}
	var parts []Part

	wfMode := true
	for _, line := range strings.Split(input, "\n") {

		if line == "" {
			wfMode = false
			continue
		}

		if wfMode {
			wf := parseWorkflow(line)
			workflows[wf.name] = &wf
		} else {
			parts = append(parts, parsePart(line))
		}
	}

	return workflows, parts
}

func parsePart(input string) Part {
	input = input[1 : len(input)-1] // remove {}
	catStr := strings.Split(input, ",")
	part := &Part{}
	for _, catVal := range catStr {
		valParts := strings.Split(catVal, "=")
		val, err := strconv.Atoi(valParts[1])
		if err != nil {
			panic(-1)
		}

		cat := rune(valParts[0][0])
		(*part)[cat] = val
	}

	return *part
}

func parseWorkflow(input string) Workflow {
	wfNameEnd := strings.IndexRune(input, '{')
	wfName := input[:wfNameEnd]

	rulesStr := input[wfNameEnd+1 : len(input)-1]
	rulesArray := strings.Split(rulesStr, ",")

	var rules []*Rule
	for _, ruleStr := range rulesArray {
		rules = append(rules, parseRule(ruleStr))
	}
	return Workflow{
		name:  wfName,
		rules: rules,
	}
}

func parseRule(input string) *Rule {
	ruleParts := strings.Split(input, ":")
	if len(ruleParts) == 1 {
		return &Rule{
			condition: nil,
			trueWF:    ruleParts[0],
		}
	} else {
		condStr := ruleParts[0]
		val, err := strconv.Atoi(condStr[2:])
		if err != nil {
			panic(-1)
		}
		return &Rule{
			condition: &Condition{
				cat: rune(condStr[0]),
				op:  rune(condStr[1]),
				val: val,
			},
			trueWF: ruleParts[1],
		}
	}
}

var CATEGORIES = []rune{'x', 'm', 'a', 's'}

type Part = map[rune]int

type Workflow struct {
	name  string
	rules []*Rule
}
type Operation = rune

const (
	LT Operation = '<'
	GT Operation = '>'
)

type Rule struct {
	condition *Condition //no condition == true
	trueWF    string
}
type Condition struct {
	cat rune
	op  Operation
	val int
	//isNeg bool
}

type CondInterval struct {
	lt int
	gt int
}

type CompCondition struct {
	limits map[rune]*CondInterval
}

func newCompCondition() *CompCondition {
	cond := &CompCondition{limits: map[rune]*CondInterval{}}
	for _, cat := range CATEGORIES {
		cond.limits[cat] = &CondInterval{
			lt: 4001,
			gt: 0,
		}
	}
	return cond
}
func (c *CompCondition) getPossible() int64 {

	for _, cat := range CATEGORIES {
		if c.limits[cat] == nil {
			c.limits[cat] = &CondInterval{
				lt: 4001,
				gt: 0,
			}
		}
	}

	possByCat := map[rune]int{}

	for cat, interval := range c.limits {
		catPoss := 0
		if interval.lt > interval.gt {
			catPoss = interval.lt - interval.gt - 1
		} else {
			catPoss = 4000 - interval.gt + interval.lt - 1
		}
		possByCat[cat] = catPoss
	}

	all := int64(possByCat[CATEGORIES[0]])
	for i := 1; i < len(CATEGORIES); i++ {
		all *= int64(possByCat[CATEGORIES[i]])
	}

	return all
}
func (c *CompCondition) and(c2 *Condition) {
	if _, found := c.limits[c2.cat]; !found {
		c.limits[c2.cat] = &CondInterval{
			lt: 4001,
			gt: 0,
		}
	}

	cLimit := c.limits[c2.cat]
	if c2.op == LT && c2.val < cLimit.lt {
		cLimit.lt = c2.val
	} else if c2.op == GT && c2.val > cLimit.gt {
		cLimit.gt = c2.val
	}
}

func (c *Condition) toComp() *CompCondition {
	lt := 4001
	gt := 0

	comp := &CompCondition{
		limits: map[rune]*CondInterval{},
	}
	for _, cat := range CATEGORIES {
		comp.limits[cat] = &CondInterval{
			lt: lt,
			gt: gt,
		}
	}

	if c.op == LT {
		gt = c.val - 1
	} else {
		lt = c.val + 1
	}
	comp.limits[c.cat].lt = lt
	comp.limits[c.cat].gt = gt
	return comp
}
func (c *Condition) countExcluded() int {
	if c.op == LT {
		return c.val - 1
	} else {
		return 4000 - c.val
	}
}

func (c *Condition) Negate() *Condition {
	if c.op == LT {
		return &Condition{
			cat: c.cat,
			op:  GT,
			val: c.val - 1,
		}
	} else {
		return &Condition{
			cat: c.cat,
			op:  LT,
			val: c.val + 1,
		}
	}
}

func (c *Condition) String() string {
	return fmt.Sprintf("%s%s%d", string(c.cat), string(c.op), c.val)
}

func (r *Rule) eval(part Part) *string {
	if r.condition == nil {
		return &r.trueWF
	} else {
		res := r.condition.eval(part)
		if !res {
			return nil
		} else {
			return &r.trueWF
		}
	}
}

func (c *Condition) eval(part Part) bool {

	partVal := part[c.cat]
	if c.op == LT {
		return partVal < c.val
	} else if c.op == GT {
		return partVal > c.val
	}
	panic(-1)
}
