package main

import (
	"fmt"
	"strconv"
	"strings"
)

func part1_new(input string) any {
	sum := 0
	for lineIx, line := range strings.Split(input, "\n") {

		if line == "" {
			continue
		}

		_ = lineIx
		fields := strings.Fields(line)
		springs, groupsList := fields[0], fields[1]

		var groupsInt []int
		for _, s := range strings.Split(groupsList, ",") {
			n, err := strconv.Atoi(s)
			if err != nil {
				panic(1)
			}
			groupsInt = append(groupsInt, n)
		}

		validCount := solve2(springs, groupsInt)

		sum += validCount
	}

	return sum
}
func part1(input string) any {

	sum := 0
	maxVisited := 0
	maxLine := ""
	for lineIx, line := range strings.Split(input, "\n") {

		if line == "" {
			continue
		}

		fields := strings.Fields(line)
		springs, groupsList := fields[0], fields[1]

		var groupsInt []int
		for _, s := range strings.Split(groupsList, ",") {
			n, err := strconv.Atoi(s)
			if err != nil {
				panic(1)
			}
			groupsInt = append(groupsInt, n)
		}

		valids := map[string]bool{}
		visited := map[string]bool{}
		solve(springs, groupsInt, 0, visited, valids)

		validCount := 0
		for _, v := range valids {
			if v {
				//fmt.Printf("==>valid:%s\n", k)
				validCount++
			}
		}
		sum += validCount
		fmt.Printf("%d: %s [visited: %d, valids: %d]\n", lineIx, line, len(visited), validCount)

		if len(visited) > maxVisited {
			maxVisited = len(visited)
			maxLine = line
		}
	}

	fmt.Printf("WORST: %s | visited: %d\n", maxLine, maxVisited)
	return sum
}

func inpKey(input string, groups []int) string {
	grpStr := ""
	for _, grp := range groups {
		grpStr += strconv.Itoa(grp) + ","
	}
	return input + "_" + grpStr
}

func solve3(input string, groups []int, rec int, visited map[string]int) int {

	//fmt.Printf("%strying: %s\n", strings.Repeat("-", rec), input)
	key := inpKey(input, groups)

	if val, ok := visited[key]; ok {
		return val
	}

	valid := isValid(input, groups)
	if valid {
		return 1
	}

	counts := count(input)
	qmCount, hashCount := counts['?'], counts['#']

	expectedHashes := sumArray(groups)

	if qmCount+hashCount < expectedHashes || hashCount > expectedHashes {
		//impossible
		fmt.Printf("!!!IMPOSSIBLE\n")
		return 0
	}

	missingHashes := expectedHashes - hashCount
	missingDots := qmCount - missingHashes
	_ = missingDots

	cmpFirst := false
	grp := groups[0]
	i := 0
	for !cmpFirst {
		hs := 0
		qm := 0

		if input[i] == '#' {
			hs++
		} else if input[i] == '?' {
			qm++
			//branch off
		} else if input[i] == '.' {
			if hs >= grp {

			}
			// if not complete
			// - if
		}

		if qm+hs == grp {
			//TODO check visited
			return 1 + solve3(input[i+1:], groups[1:], rec+1, visited)
		}

	}

	return 0
}

func solve(input string, groups []int, rec int, visited map[string]bool, valids map[string]bool) {

	fmt.Printf("%strying: %s\n", strings.Repeat(" ", rec), input)

	visited[input] = true

	counts := count(input)
	qmCount, hashCount := counts['?'], counts['#']
	expectedHashes := sumArray(groups)

	if hashCount == expectedHashes {
		valid := isValid(input, groups)
		if valid {
			valids[input] = true
			return
		}
		return
	}

	if qmCount+hashCount < expectedHashes || hashCount > expectedHashes {
		//impossible
		fmt.Printf("!!!IMPOSSIBLE\n")
		return
	}

	missingHashes := expectedHashes - hashCount
	missingDots := qmCount - missingHashes
	_ = missingDots
	var qmIndexes []int
	for ix, s := range input {
		if s == '?' {
			qmIndexes = append(qmIndexes, ix)
		}
	}

	/*	if len(groups) > 1 && rec == 0 {

		restSum := 0
		for j := 1; j < len(groups); j++ {
			restSum += groups[j]
		}
		minOccupiedSpace := restSum + len(groups) - 1
		subValid := map[string]bool{}
		fmt.Printf("subvalid1:[%s] | %v...\n", input[0:len(input)-minOccupiedSpace], groups[0:1])
		solve(input[0:len(input)-minOccupiedSpace], groups[0:1], rec+1, map[string]bool{}, subValid)
		fmt.Printf("subvalid1:[%s]==>%d\n", input[0:len(input)-minOccupiedSpace], countTrue(subValid))

		subValid = map[string]bool{}
		restStr := "." + input[len(input)-minOccupiedSpace+1:]
		restGrp := groups[1:]
		fmt.Printf("subvalid2:[%s] | %v\n", restStr, restGrp)
		solve(restStr, restGrp, rec+1, map[string]bool{}, subValid)
		fmt.Printf("subvalid2:[%s]==>%d\n", restStr, countTrue(subValid))
	}*/

	//for i, group := range groups {
	//	hs := 0
	//	for i := 0; i < len(input); i++ {
	//		if input[i] == '#' {
	//			hs++
	//		} else if input[i] == '?' {
	//			//branch off
	//
	//		}
	//	}
	//}

	for j, ix := range qmIndexes {
		//if missingDots > 0 && len(qmIndexes)-j >= missingDots {
		//	strHash := input[0:ix] + "." + input[ix+1:]
		//	if !visited[strHash] {
		//		solve(strHash, groups, rec+1, visited, valids)
		//	}
		//}
		if missingHashes > 0 && len(qmIndexes)-j >= missingHashes {
			strHash := input[0:ix] + "#" + input[ix+1:]
			if !visited[strHash] {
				solve(strHash, groups, rec+1, visited, valids)
			}
		}

	}

	/*

		.??..??...?##. 1,1,3
			.#?..??...?##. 1,3 - 1
				.#...#?...?##. 3 - 2
		 	..?..??...?##. 1,1,3


		?? => 0=2, 1=0
		??? => 0=2, 1=1


					3#, 9?
					?###???????? 3,2,1 ==> 10

					7?
					??????? 2,1 =>
				  	##.


			.??..??...?##. 1,1,3
			1 + .#?..??...###.;1
			1 +



					1,3,1,6 = 11 | 7#, 8?
					=> .#.###.#.######
					   _______________
					\.#[\.?]

								?#?#?#?#?#?#?#? 1,3,1,6


							1	 3	   1	6
							.?#? ?#?#? ?#? ?#?#?#?.


						.???.###. 1,1,3
						1,1		3
						.???. .###.


						.??..??...?##. 1,1,3
						1
						.??.	.??.	.?##.
	*/

}

type Part struct {
	start int
	end   int
	group int
}

func solve2(str string, groups []int) int {
	padded := "." + str + "."
	start := 0

	var combs []int
	var parts []*Part
	for _, group := range groups {

		end := start + 1
		currStr := ""
		isCompleteGroup := false
		for !isCompleteGroup {
			end++
			currStr = padded[start:end]

			last := currStr[len(currStr)-1]

			validStart := currStr[0] == '.' || currStr[0] == '?'
			validEnd := last == '.' || last == '?'

			middleStr := currStr[1 : len(currStr)-1]
			cntMiddle := count(middleStr)

			hasEnough := (cntMiddle['?'] + cntMiddle['#']) >= group
			validMiddle := hasEnough

			if cntMiddle['.'] > 0 && len(parts) > 0 {
				//extend previous
				dotIx := strings.Index(middleStr, ".")

				if dotIx == -1 {
					panic(1)
				}

				prev := parts[len(parts)-1]
				prev.end += dotIx + 1

				start = prev.end
				end = start + 1
				//continue
			} else {

				nextIsNotDot := end == len(padded) || padded[end] != '.'

				isCompleteGroup = validStart && validEnd && validMiddle && nextIsNotDot
			}

		}

		parts = append(parts, &Part{start, end, group})

		start = end - 1

	}
	//fmt.Printf("str: %d-%d=%s\n", start, end, currStr)
	//
	//valids := map[string]bool{}
	//currStr = "." + currStr[1:len(currStr)-1] + "."
	//solve(currStr, []int{group}, 0, map[string]bool{}, valids)
	//validCnt := 0
	//for _, v := range valids {
	//	if v {
	//		validCnt++
	//	}
	//}
	//combs = append(combs, validCnt)
	//fmt.Printf("group %d, str: %s |=> %d \n", group, currStr, validCnt)

	for _, part := range parts {
		currStr := padded[part.start:part.end]
		fmt.Printf("str: %d-%d=%s\n", part.start, part.end, currStr)

		valids := map[string]bool{}
		currStr = "." + currStr[1:len(currStr)-1] + "."
		solve(currStr, []int{part.group}, 0, map[string]bool{}, valids)
		validCnt := 0
		for _, v := range valids {
			if v {
				validCnt++
			}
		}
		combs = append(combs, validCnt)
		fmt.Printf("group %d, str: %s |=> %d \n\n", part.group, currStr, validCnt)
	}

	tot := 1
	for _, comb := range combs {
		tot *= comb
	}
	return tot
}

func countTrue(x map[string]bool) int {
	cnt := 0
	for _, v := range x {
		if v {
			cnt++
		}
	}
	return cnt
}

func sumArray(x []int) int {
	sum := 0
	for _, i := range x {
		sum += i
	}
	return sum
}
func count(x string) map[rune]int {
	res := map[rune]int{}
	for _, c := range x {
		res[c]++
	}
	return res
}

func isValid(x string, groups []int) bool {

	var foundGroups []int
	currDmgGrp := 0
	lastSym := '*'
	paddedX := "." + x + "."
	//paddedX = strings.Replace(paddedX, "?", ".", -1)
	for _, s := range paddedX {
		if s == '?' {
			//fmt.Printf("unexpected ?\n")
			//return false
			s = '.'
		}

		if s == '#' && lastSym == '.' {
			currDmgGrp = 1
		} else if s == '.' && lastSym == '#' {
			foundGroups = append(foundGroups, currDmgGrp)
			currDmgGrp = 0
		} else if s == '#' && lastSym == '#' {
			currDmgGrp++
		}

		lastSym = s
	}

	if len(foundGroups) != len(groups) {
		return false
	}

	for i := 0; i < len(foundGroups); i++ {
		if foundGroups[i] != groups[i] {
			return false
		}
	}
	return true
}

func getFixedGroups(input string) []int {
	var fixedGroups []int
	currDmgGrp := 0
	lastSym := '*'
	for _, s := range input {
		if s == '#' && lastSym == '.' {
			currDmgGrp = 1
		} else if s == '.' && lastSym == '#' {
			fixedGroups = append(fixedGroups, currDmgGrp)
			currDmgGrp = 0
		} else if s == '#' && lastSym == '#' {
			currDmgGrp++
		}

		lastSym = s
	}

	if currDmgGrp > 0 {
		fixedGroups = append(fixedGroups, currDmgGrp)
	}
	fmt.Printf("fixed in %s: %v\n", input, fixedGroups)

	return fixedGroups
}

func part2(input string) any {

	sum := 0
	for lineIx, line := range strings.Split(input, "\n") {

		if line == "" {
			continue
		}

		_ = lineIx
		fields := strings.Fields(line)
		springs, groupsList := fields[0], fields[1]

		springs = strings.Repeat(springs+"?", 5)
		springs = springs[:len(springs)-1] // remove last ?

		groupsList = strings.Repeat(groupsList+",", 5)
		groupsList = groupsList[:len(groupsList)-1]

		var groupsInt []int
		for _, s := range strings.Split(groupsList, ",") {
			n, err := strconv.Atoi(s)
			if err != nil {
				panic(1)
			}
			groupsInt = append(groupsInt, n)
		}

		//validCount := solve2(springs, groupsInt)
		valids := map[string]bool{}
		visited := map[string]bool{}
		solve(springs, groupsInt, 0, visited, valids)

		validCount := 0
		for _, v := range valids {
			if v {
				validCount++
			}
		}
		sum += validCount
		fmt.Printf("%d: %s [visited: %d, valids: %d]\n", lineIx, springs, len(visited), validCount)

	}

	return sum
}
