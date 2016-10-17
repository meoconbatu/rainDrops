package raindrops

import "strconv"

const testVersion = 2

var rainDrops = []struct {
	factor int
	tune   string
}{
	{3, "Pling"},
	{5, "Plang"},
	{7, "Plong"},
}

func Convert(input int) string {
	var output string
	rainDropFactors := ContainsRainDrops(input)
	if len(rainDropFactors) == 0 {
		output = strconv.Itoa(input)
	} else {
		output = RainDrops(rainDropFactors)
	}
	return output
}
func ContainsRainDrops(input int) []int {
	rainDropFactors := []int{}
	for _, rainDrop := range rainDrops {
		if input%rainDrop.factor == 0 {
			rainDropFactors = append(rainDropFactors, rainDrop.factor)
		}
	}
	return rainDropFactors
}

func RainDrops(factors []int) string {
	var output string
	for i := 0; i < len(factors); i++ {
		for _, rainDrop := range rainDrops {
			if rainDrop.factor == factors[i] {
				output += rainDrop.tune
			}
		}
	}
	return output
}
