package main

//ring die: 0: blank 1: opp&stress 2: opp 3: pass&stress 4: pass 5: crit&stress
//skill die: 0-1: blank 2-4: opp 5-6: pass&stress 7-8: pass 9: pass&opp 10: crit&stress 11: crit

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
)

func rollDieSingle(die string) string {
	if die == "r" {
		rawRoll := rand.Int63n(6)
		roll := strconv.FormatInt(rawRoll, 10) //roll is 0-5
		roll = (roll) + "r"
		if roll == "5r" {
			newRoll := rollDieSingle("r")
			roll = roll + newRoll
		}
		return roll

	} else {
		rawRoll := rand.Int63n(12)
		roll := strconv.FormatInt(rawRoll, 10) //roll is 0-11
		roll = (roll) + "s"
		if roll == "10s" || roll == "11s" {
			newRoll := rollDieSingle("1s")
			roll = roll + newRoll
		}
		return roll
	}
}

func parseAndRoll(cmd string) {
	expRing := regexp.MustCompile(`\dr`)
	expSkill := regexp.MustCompile(`\ds`)
	splitRing := expRing.FindStringSubmatch(cmd)
	splitSkill := expSkill.FindStringSubmatch(cmd)
	fmt.Println(splitRing)

	ringTrimmed := strings.TrimSuffix(splitRing[0], "r")
	ringNum, err := strconv.Atoi(ringTrimmed)
	if err != nil {
		fmt.Println(err)
		return
	}
	for range ringNum {
		roll := rollDieSingle("r")
		DicePool = append(DicePool, roll)
	}

	fmt.Println(splitSkill)
	skillTrimmed := strings.TrimSuffix(splitSkill[0], "s")
	skillNum, err := strconv.Atoi(skillTrimmed)
	if err != nil {
		fmt.Println(err)
		return
	}
	for range skillNum {
		roll := rollDieSingle("s")
		DicePool = append(DicePool, roll)
	}
	pushRoll(DicePool)
}

func pushRoll(pool []string) {
	for i := range pool {
		fmt.Println(pool[i])
	}
	DicePool = nil
}
