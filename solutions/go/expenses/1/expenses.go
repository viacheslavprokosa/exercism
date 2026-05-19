package expenses

import ("errors")
type Record struct {
	Day      int
	Amount   float64
	Category string
}

type DaysPeriod struct {
	From int
	To   int
}

func Filter(in []Record, predicate func(Record) bool) []Record {
	var slice []Record
	for i := 0; i < len(in); i++ {
		if predicate(in[i]) {
			slice = append(slice, in[i])
		}
	}

	return slice
}

func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(in Record) bool {
		return p.From <= in.Day && p.To >= in.Day
	}
}

func ByCategory(c string) func(Record) bool {
	return func(in Record) bool {
		return in.Category == c
	}
}

func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	var total float64
	for _, item := range Filter(in, ByDaysPeriod(p)) {
		total += item.Amount
	}
	return total
}
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	byCategory := Filter(in, ByCategory(c))
	if len(byCategory) == 0 {
		return 0.0, errors.New("unknown category entertainment")
	}
	byPeriodAndCategory := Filter(byCategory, ByDaysPeriod(p))
	
	return TotalByPeriod(byPeriodAndCategory, p), nil
}