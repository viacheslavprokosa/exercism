package lasagnamaster

func PreparationTime(layers []string, timePerLayer int) int {
	if timePerLayer==0 {
		timePerLayer = 2
	}
	return len(layers) * timePerLayer
}

func Quantities(components []string) (int, float64) {
	var sauce float64
	var noodles int
	for _, component := range components {
		if component == "sauce" {
			sauce += 0.2
		}
		if component == "noodles" {
			noodles += 50
		}
	}
	return noodles, sauce
}

func AddSecretIngredient(friendsList []string, myList []string) {
	secretIngredient := friendsList[len(friendsList)-1]
	myList[len(myList)-1] = secretIngredient
}

func ScaleRecipe(quantities []float64, needed int) []float64 {
	coefficient := float64(needed) / float64(2)
	desiredQuantities := make([]float64, len(quantities))
	for i, quantity := range quantities {
		desiredQuantities[i] = quantity * coefficient
	}
	return desiredQuantities
}


// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
