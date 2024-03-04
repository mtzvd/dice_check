Dice Roll Success Calculator

## Introduction
This Dice Roll Success Calculator is a Go program designed to calculate the probability of achieving a specific success condition in a dice roll scenario. It can accommodate various configurations, including the number of dice, the number of sides on each die, a success threshold, the number of dice that must meet or exceed this threshold, and a modifier that can be applied to each roll.

## Features
- **Flexible Configuration:** Set up the dice roll simulation with custom values for dice number, dice sides, success threshold, required number of successful dice, and roll modifiers.
- **Modifier Distribution:** Analyzes all possible distributions of a given modifier across all dice in a roll.
- **Success Calculation:** Calculates the probability of a roll being successful under the given conditions.
- **Comprehensive Analysis:** Generates all possible dice combinations based on the configuration and evaluates each for success.

## Installation
To use this program, ensure you have Go installed on your system. Clone the repository to your local machine and navigate to the program's directory.

```bash
git clone <repository-url>
cd <program-directory>
```

## Usage

1. **Configuration:** Modify the `main` function in the program to set your desired configuration. For example:
   ```go
   cfg := config{
       threshold:  4,
       successNum: 3,
       diceNum:    3,
       sides:      6,
       modifier:   1,
   }
   ```
   This configuration sets up a scenario with 3 six-sided dice, a success threshold of 4, at least 3 dice must meet or exceed this threshold, and a modifier of +1 to each roll.

2. **Running the Program:** Compile and run the program using the Go command.
   ```bash
   go run main.go
   ```
   The program will output the process of generating dice combinations, the distribution of the modifier, and the final probability of success based on the provided configuration.

3. **Interpreting Output:** The output will include detailed logs of each dice combination and whether it meets the success criteria. Finally, it will display the total number of successful combinations out of all possible combinations, along with the probability of success.

## Example Output
```
Generate a table of all variants of the distribution of the modifier 1 on the number of dice=3
Table of distribution of modifier 1 on the number of dice = 3
 [[0 0 1] [0 1 0] [1 0 0]]

Generating all the combinations of 3d6
...
[1 2 3] :  false
[6 5 4] :  true
Successful rolls 100 of 216
Probability of success at 
{diceNum:3 sides:6 threshold:4 successNum:3 modifier:1} 
=0.46296296
```

This output demonstrates the program's process and final probability calculation, providing a clear indication of how likely a given dice roll configuration is to succeed under the specified conditions.
