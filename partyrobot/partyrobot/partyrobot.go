package partyrobot

import (
	"fmt"
	"strings"
)

func Welcome(name string) string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("Welcome to my party, %s!", name))

	return builder.String()
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age))

	return builder.String()
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("Welcome to my party, %s!\n", name))
	builder.WriteString(fmt.Sprintf("You have been assigned to table %03d. Your table is %s, exactly %.1f meters from here.\n", table, direction, distance))
	builder.WriteString(fmt.Sprintf("You will be sitting next to %s.", neighbor))

	return builder.String()
}
