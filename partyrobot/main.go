package main

import (
	"fmt"
	"partyrobot/partyrobot"
)

func main() {
	fmt.Sprintf(partyrobot.Welcome("Muchiri"))
	fmt.Sprintf(partyrobot.HappyBirthday("John", 30))
	fmt.Sprintf(partyrobot.AssignTable("Ken", 022, "Joy", "on the right", 100.0))
}
