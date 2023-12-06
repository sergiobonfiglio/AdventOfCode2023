package day5

import (
	"AdventOfCode2023/utils"
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
	"strings"
	"unicode"
)

func readInput() []string {

	readFile, err := os.Open("day5/input.txt")
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var data []string

	for fileScanner.Scan() {

		for fileScanner.Text() != "" {
			data = append(data, fileScanner.Text())
			fileScanner.Scan()
		}

	}

	err = readFile.Close()
	if err != nil {
		panic(err)
	}

	return data
}

func SolveDay() {

	var input = readInput()

	sol1, sol2 := solveParts(input)

	fmt.Println("Part 1:", sol1)

	fmt.Println("Part 2:", sol2)
}

type Map struct {
	entries []MapEntry
}

func (m *Map) add(entry MapEntry) {
	m.entries = append(m.entries, entry)
}
func (m *Map) get(source int64) int64 {

	for _, entry := range m.entries {
		if source >= entry.sourceStart && source < entry.sourceStart+entry.length {
			return source - entry.sourceStart + entry.destStart
		}
	}
	return source
}

type Range struct {
	start  int64
	length int64
}

func (m *Map) getRangeFromRanges(sources []Range) []Range {
	var ranges []Range
	for _, sourceRange := range sources {
		ranges = append(ranges, m.getRange(sourceRange.start, sourceRange.length)...)
	}
	return ranges
}

func (m *Map) getRange(sourceStart int64, length int64) []Range {

	var ranges []Range

	mapped := int64(0)
	for _, entry := range m.entries {

		sourceEnd := sourceStart + length
		entryEnd := entry.sourceStart + entry.length
		isIntersecting := (sourceStart >= entry.sourceStart && sourceStart < entryEnd) ||
			(sourceEnd >= entry.sourceStart && sourceEnd < entryEnd) ||
			(entry.sourceStart >= sourceStart && entry.sourceStart < sourceEnd) ||
			(entryEnd >= sourceStart && entryEnd < sourceEnd)

		var interStart, interEnd int64 = -1, -1
		if isIntersecting {
			interStart = max(sourceStart, entry.sourceStart)
			interEnd = min(sourceEnd, entryEnd)
		}

		if isIntersecting {

			if interStart > sourceStart+mapped {
				//add identity mappings before intersection
				maprange := Range{
					start:  sourceStart + mapped,
					length: interStart - (sourceStart + mapped),
				}
				ranges = append(ranges, maprange)

				mapped += interStart - (sourceStart + mapped)
			}

			// add intersection mapping
			// source - entry.sourceStart + entry.destStart
			maprange := Range{
				start:  interStart + (entry.destStart - entry.sourceStart),
				length: interEnd - interStart,
			}
			ranges = append(ranges, maprange)
			mapped += interEnd - interStart
		}

	}
	// map the leftovers
	if mapped < length {
		ranges = append(ranges, Range{
			start:  sourceStart + mapped,
			length: length - mapped,
		})

	}

	//sort intersections by start
	slices.SortFunc(ranges, func(a, b Range) int {
		return cmp.Compare(a.start, b.start)
	})

	return ranges
}

func (m *Map) sort() {
	slices.SortFunc(m.entries, func(a, b MapEntry) int {
		return cmp.Compare(a.sourceStart, b.sourceStart)
	})
}

type MapEntry struct {
	destStart   int64
	sourceStart int64
	length      int64
}

func solveParts(input []string) (int64, int64) {

	var seeds []int64

	seedToSoil := &Map{}
	soilToFert := &Map{}
	fertToWater := &Map{}
	waterToLight := &Map{}
	lightToTemp := &Map{}
	tempToHum := &Map{}
	humToLoc := &Map{}

	var currMap *Map

	for i, row := range input {
		if i == 0 {
			//seeds
			parts := strings.Split(row, ":")
			seeds = utils.ToInt64Array(parts[1])
			continue
		}

		if row == "" {
			continue
		}

		if row == "seed-to-soil map:" {
			currMap = seedToSoil
		} else if row == "soil-to-fertilizer map:" {
			currMap = soilToFert
		} else if row == "fertilizer-to-water map:" {
			currMap = fertToWater
		} else if row == "water-to-light map:" {
			currMap = waterToLight
		} else if row == "light-to-temperature map:" {
			currMap = lightToTemp
		} else if row == "temperature-to-humidity map:" {
			currMap = tempToHum
		} else if row == "humidity-to-location map:" {
			currMap = humToLoc
		}

		if row != "" && unicode.IsDigit(rune(row[0])) {

			params := utils.ToInt64Array(row)
			currMap.add(MapEntry{
				destStart:   params[0],
				sourceStart: params[1],
				length:      params[2],
			})
		}
	}
	currMap.sort()

	maps := []*Map{
		seedToSoil,
		soilToFert,
		fertToWater,
		waterToLight,
		lightToTemp,
		tempToHum,
		humToLoc,
	}
	for _, m := range maps {
		m.sort()
	}

	var minLoc *int64
	for _, seed := range seeds {
		soil := seedToSoil.get(seed)
		fert := soilToFert.get(soil)
		water := fertToWater.get(fert)
		light := waterToLight.get(water)
		temp := lightToTemp.get(light)
		hum := tempToHum.get(temp)
		loc := humToLoc.get(hum)

		if minLoc == nil || loc < *minLoc {
			minLoc = &loc
		}
	}

	var minLocRange *int64
	for i := 0; i < len(seeds); i += 2 {
		seedStart := seeds[i]
		length := seeds[i+1]

		soils := seedToSoil.getRange(seedStart, length)
		ferts := soilToFert.getRangeFromRanges(soils)
		waters := fertToWater.getRangeFromRanges(ferts)
		lights := waterToLight.getRangeFromRanges(waters)
		temps := lightToTemp.getRangeFromRanges(lights)
		hums := tempToHum.getRangeFromRanges(temps)
		locs := humToLoc.getRangeFromRanges(hums)

		for _, loc := range locs {
			rangeMin := loc.start
			if minLocRange == nil || rangeMin < *minLocRange {
				minLocRange = &rangeMin
			}
		}

	}

	return *minLoc, *minLocRange
}
