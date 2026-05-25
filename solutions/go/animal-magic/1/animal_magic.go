package chance

import "math/rand/v2"

func RollADie() int {
	return rand.IntN(19)+1
}

func GenerateWandEnergy() float64 {
	return rand.Float64() * 12.0
}

func ShuffleAnimals() []string {
	var animals = []string{
		"ant",
		"beaver",
		"cat",
		"dog",
		"elephant",
		"fox",
		"giraffe",
		"hedgehog",
	}
	rand.Shuffle(len(animals), func(i, j int) {
		animals[i], animals[j] = animals[j], animals[i]
	})
	return animals
}

