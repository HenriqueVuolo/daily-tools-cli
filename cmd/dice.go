package cmd

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	numberOfDice  int
	numberOfFaces int
	modifier      int
	random        = rand.New(rand.NewSource(time.Now().UnixNano()))
)

var diceCmd = &cobra.Command{
	Use:   "dice",
	Short: "Roll or or more dice",
	Long: `The dice command allows you to simulate the roll of one or more dice. By default, it rolls a six-sided die (d6), but you can customize the number of dice (A), faces (X), and apply a modifier (B).

- Format: AdX + B
- Example:
  $ daily-tools-cli dice --numberOfDice 2 --numberOfFaces 20 --mod -3 => 2d20 - 3
    OR
  $ daily-tools-cli dice -A 3 -X 8 -B 1 => 3d8 + 1`,
	Run: func(cmd *cobra.Command, args []string) {
		critFmt := color.New(color.FgHiYellow, color.BgYellow)
		failFmt := color.New(color.FgHiRed, color.BgRed)
		defaultFmt := color.New(color.FgMagenta, color.BgHiWhite)

		validateDice()

		total := modifier

		for i := 0; i < numberOfDice; i++ {
			result := rollDie()
			fmtStyle := defaultFmt
			if result == numberOfFaces {
				fmtStyle = critFmt
			} else if result == 1 && numberOfFaces != 1 {
				fmtStyle = failFmt
			}
			fmtStyle.Printf(" Die %d: %d ", i+1, result)
			fmt.Println()
			total += result
		}
		fmt.Println()
		defaultFmt.Printf(" %dd%d + (%d) =  %d \n", numberOfDice, numberOfFaces, modifier, total)
	},
}

func init() {
	rootCmd.AddCommand(diceCmd)

	diceCmd.Flags().IntVarP(&numberOfDice, "numberOfDice", "A", 1, "Number of dice to be rolled")
	diceCmd.Flags().IntVarP(&numberOfFaces, "numberOfFace", "X", 6, "Number of faces on each die")
	diceCmd.Flags().IntVarP(&modifier, "mod", "B", 0, "Additive modifier to the result")
}

func rollDie() int {
	return random.Intn(numberOfFaces) + 1
}

func validateDice() {
	if numberOfDice <= 0 || numberOfDice > 1000 || numberOfFaces < 1 || numberOfFaces > 1000 {
		log.Fatalf("Invalid parameter")
	}
}
