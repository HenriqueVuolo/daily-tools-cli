package cmd

import (
	"log"
	"math/rand"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	length       int
	uppercase    bool
	number       bool
	special      bool
	lowercaseSet = "abcdedfghijklmnopqrstuvwxyz"
	uppercaseSet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialSet   = "!@#$%&*()+=<>?"
	numberSet    = "0123456789"
)

var passwordCmd = &cobra.Command{
	Use:   "password",
	Short: "Generate a random password",
	Run: func(cmd *cobra.Command, args []string) {
		if length < 1 || length > 50 {
			log.Fatalf("Invalid length")
		}

		random := rand.New(rand.NewSource(time.Now().UnixNano()))
		charSet := lowercaseSet
		password := ""

		if uppercase {
			charSet += uppercaseSet
		}

		if special {
			charSet += specialSet
		}

		if number {
			charSet += numberSet
		}

		for i := 0; i < length; i++ {
			charIndex := random.Intn(len(charSet))
			password += string(charSet[charIndex])
		}

		color.New(color.Bold).Println(password)
	},
}

func init() {
	rootCmd.AddCommand(passwordCmd)
	passwordCmd.Flags().IntVarP(&length, "length", "l", 8, "Password length")
	passwordCmd.Flags().BoolVarP(&uppercase, "uppercase", "u", false, "Include uppercase characters")
	passwordCmd.Flags().BoolVarP(&number, "number", "n", false, "Include numbers")
	passwordCmd.Flags().BoolVarP(&special, "special", "s", false, "Include special characters")
}
