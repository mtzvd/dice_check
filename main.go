package main

import (
	"fmt"
)

type config struct {
	diceNum    int
	sides      int
	threshold  int
	successNum int
	modifier   int
}

// Abs returns the modulus of a number
func Abs(number int) int {
	if number < 0 {
		return number * -1
	}

	return number
}

// isSuccess checks whether the combination is successful
func isSuccess(threshold, successNum int, combination []int) bool {
	aboveThreshold := 0
	for _, dice := range combination {
		if dice >= threshold {
			aboveThreshold++
		}
	}

	return aboveThreshold >= successNum
}

// generateModifiers returns a slice with all combinations of modifier distribution over dices
func generateModifiers(modifier, diceNum int) [][]int {
	absMod := Abs(modifier)
	if diceNum == 1 {
		// If there is one dice left, return all remaining units in it
		return [][]int{{absMod}}
	}

	// Initialization of the initial state for the first dice
	initialCombinations := make([][]int, absMod+1)
	for i := range initialCombinations {
		initialCombinations[i] = []int{i}
	}

	// Build combinations for each dice
	for box := 2; box <= diceNum; box++ {
		newCombinations := [][]int{}
		for _, combo := range initialCombinations {
			sum := 0
			for _, value := range combo {
				sum += value
			}
			for i := 0; i <= absMod-sum; i++ {
				if box == diceNum && sum+i != absMod {
					continue // For the last dice, make sure we use all units
				}

				// Add a new value to existing combinations
				newCombo := append([]int(nil), combo...)
				newCombo = append(newCombo, i)
				newCombinations = append(newCombinations, newCombo)
			}
		}
		initialCombinations = newCombinations
	}

	// Apply modification for negative modifiers
	if modifier < 0 {
		for i, combo := range initialCombinations {
			for j, value := range combo {
				initialCombinations[i][j] = -value
			}
		}
	}

	return initialCombinations
}

// generateCombinations creates all possible combinations for a given number and type of dice.
// Then sends them to a channel for processing.
func generateCombinations(diceNum, sides int, combinationsChan chan<- []int) {
	defer close(combinationsChan) // Close the channel at the end of the function

	// Initialize initial combination with empty list
	result := [][]int{{}}

	// For every dice
	for d := 0; d < diceNum; d++ {
		newResult := [][]int{}

		// Go through all current combinations
		for _, combo := range result {

			// Add all possible values of the current dice to each combination
			for s := 1; s <= sides; s++ {
				newCombo := append([]int(nil), combo...)
				newCombo = append(newCombo, s)

				// Add the new combination to the temporary result for the next level
				newResult = append(newResult, newCombo)
			}
		}

		// Update the main result with new combinations
		result = newResult
	}

	// Display combinations
	for _, combo := range result {
		fmt.Printf("Generated: %v\n", combo)
	}

	// After generating all combinations, send them to the channel
	fmt.Printf("Starting processing\n")
	for _, combo := range result {
		combinationsChan <- combo
	}
}

// summSlices summs the slides element by element, the length of the resulting slide is equal to the length of the first slide
func summSlices(a, b []int) []int {
	result := make([]int, len(a))
	for i := range a {
		result[i] = a[i] + b[i]
	}

	return result
}

// processCombination returns whether the combination can succeed with the application of the modifier
func processCombination(combination []int, modifierTable [][]int, threshold, successNum int) bool {

	// Go through all combinations of modifiers
	for _, modCombo := range modifierTable {

		// Add the current combination with the modifier
		modifiedCombo := summSlices(combination, modCombo)

		// Check whether the modified combination is successful
		if isSuccess(threshold, successNum, modifiedCombo) {
			return true // Return true if the combination is successful
		}
	}

	return false // Return false if none of the combinations was successful
}

func main() {

	// Set parameters
	cfg := config{
		threshold:  4, // dice is considered a success if at least this number is on
		successNum: 3, // number of dice on which success must be rolled (taking into account the modifier) for the result of the roll to be considered successful.
		diceNum:    3, // number of dice in the roll.
		sides:      6, // number of faces of each dice
		modifier:   1, // roll modifier (can be negative).
	}

	// Reset the counters
	allCases := 0
	succesCases := 0

	//Generate modifier table
	fmt.Printf("Generate a table of all variants of the distribution of the modifier %v on the number of dice=%v  \n", cfg.modifier, cfg.diceNum)
	modifierTable := generateModifiers(cfg.modifier, cfg.diceNum)
	fmt.Printf("Table of distribution of modifier %v on the number of dice = %v \n %v\n\n", cfg.modifier, cfg.diceNum, modifierTable)

	//Create a channel for processing combinations
	combinationsChan := make(chan []int)

	//Start the generator
	fmt.Printf("Generating all the combinations of %vd%v\n", cfg.diceNum, cfg.sides)
	go generateCombinations(cfg.diceNum, cfg.sides, combinationsChan)

	// Read combinations from the channel and process them
	for combo := range combinationsChan {
		result := processCombination(combo, modifierTable, cfg.threshold, cfg.successNum)
		fmt.Println(combo, ": ", result)
		if result {
			succesCases++
		}
		allCases++
	}

	//Display the result
	fmt.Printf("Successful rolls %v of %v\n", succesCases, allCases)
	fmt.Printf("Probability of success at \n%+v \n=%v", cfg, float32(succesCases)/float32(allCases))
}
