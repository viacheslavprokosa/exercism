package sorting

import(
    "fmt"
    "strconv"
)

func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %.1f", float32(nb.Number()))
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

func ExtractFancyNumber(fnb FancyNumberBox) int {
	switch fnb.(type) {
	case FancyNumber:
		val, err := strconv.Atoi(fnb.Value())
		if err != nil {
			return 0
		}
		return val
	default:
		return 0
	}
}
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	switch fnb.(type) {
	case FancyNumber:
		val, err := strconv.Atoi(fnb.Value())
		if err != nil {
			return "This is a fancy box containing the number 0.0"
		}
		return fmt.Sprintf("This is a fancy box containing the number %.1f", float32(val))
	default:
		return "This is a fancy box containing the number 0.0"
	}
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i any) string {
	switch i.(type) {
	case float64:
		return DescribeNumber(i.(float64))
	case int:
		return DescribeNumber(float64(i.(int)))
	case NumberBox:
		return DescribeNumberBox(i.(NumberBox))
	case FancyNumberBox:
		return DescribeFancyNumberBox(i.(FancyNumberBox))
	default:
		return "Return to sender"
	}
}
