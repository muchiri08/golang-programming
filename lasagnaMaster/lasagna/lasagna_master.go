package lasagna

func PreparationTime(slice []string, time int) int {
	if time <= 0 {
		time = 2
	}
	return len(slice) * time
}

func Quantities(slice []string) (noodles int, sauce float64) {
	noodles, sauce = 0, 0

	for i := 0; i < len(slice); i++ {
		if slice[i] == "noodles" {
			noodles++
		} else if slice[i] == "sauce" {
			sauce++
		}
	}
	return noodles * 50, sauce * 0.2
}

func AddSecretIngredient(friendIngredients []string, myIngredients []string) {
	secretIngredient := friendIngredients[len(friendIngredients)-1]
	myIngredients[len(myIngredients)-1] = secretIngredient
}

func ScaleRecipe(slice []float64, num int) (amt []float64) {
	for _, v := range slice {
		amt = append(amt, v*float64(num)/2)
	}
	return amt
}
