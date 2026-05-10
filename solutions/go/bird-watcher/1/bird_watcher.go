package birdwatcher

func TotalBirdCount(birdsPerDay []int) int {
	var count int
	for i := 0; i < len(birdsPerDay); i++ {
		count += birdsPerDay[i]
	}
	return count
}

func BirdsInWeek(birdsPerDay []int, week int) int {
	if week == 0 {
		return 0
	}
	minDay := (week - 1) * 7
	maxDay := minDay + 6
	totalDays := len(birdsPerDay)
	var count int
	if totalDays < maxDay {
		maxDay = totalDays
	}
	for i := minDay; i <= maxDay; i++ {
		count += birdsPerDay[i]
	}
	return count
}

func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i += 2 {
		birdsPerDay[i] = birdsPerDay[i] + 1
	}
	return birdsPerDay
}

