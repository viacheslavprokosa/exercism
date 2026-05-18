package techpalace

import (
	"strings"
)

func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	line1 := strings.Repeat("*", numStarsPerLine)
	line3 := strings.Repeat("*", numStarsPerLine)
	return line1 + "\n" + welcomeMsg + "\n" + line3
}

func CleanupMessage(oldMsg string) string {
	newString := strings.Replace(oldMsg, "\n", "", -1)
	newString = strings.Replace(newString, "*", "", -1)
	newString = strings.Trim(newString, " ")
	return newString
}
