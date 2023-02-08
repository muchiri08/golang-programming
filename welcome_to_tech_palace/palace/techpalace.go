package palace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	starLine := strings.Repeat("*", numStarsPerLine)
	result := starLine + "\n"
	result += welcomeMsg + "\n"
	result += starLine

	return result
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	cleanedString := strings.Replace(oldMsg, "*", "", -1)
	return strings.TrimSpace(cleanedString)
}
