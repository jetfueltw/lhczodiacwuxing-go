package lhczodiacwuxing_test

import (
	"testing"

	"github.com/jetfueltw/lhczodiacwuxing-go"
	"github.com/stretchr/testify/assert"
)

func TestGetZodiacNumbers(t *testing.T) {
	assert := assert.New(t)

	zodiacNumbers, err := lhczodiacwuxing.GetZodiacNumbers(2020)
	if !assert.NoError(err) {
		return
	}
	assert.Equal(map[string][]int{
		"rat":     []int{1, 13, 25, 37, 49},
		"ox":      []int{12, 24, 36, 48},
		"tiger":   []int{11, 23, 35, 47},
		"rabbit":  []int{10, 22, 34, 46},
		"dragon":  []int{9, 21, 33, 45},
		"snake":   []int{8, 20, 32, 44},
		"horse":   []int{7, 19, 31, 43},
		"goat":    []int{6, 18, 30, 42},
		"monkey":  []int{5, 17, 29, 41},
		"rooster": []int{4, 16, 28, 40},
		"dog":     []int{3, 15, 27, 39},
		"pig":     []int{2, 14, 26, 38},
	}, zodiacNumbers)
}

func TestGetNumberZodiacs(t *testing.T) {
	assert := assert.New(t)

	numberZodiacs, err := lhczodiacwuxing.GetNumberZodiacs(2020)
	if !assert.NoError(err) {
		return
	}
	assert.Equal(map[int]string{
		1:  "rat",
		2:  "pig",
		3:  "dog",
		4:  "rooster",
		5:  "monkey",
		6:  "goat",
		7:  "horse",
		8:  "snake",
		9:  "dragon",
		10: "rabbit",
		11: "tiger",
		12: "ox",
		13: "rat",
		14: "pig",
		15: "dog",
		16: "rooster",
		17: "monkey",
		18: "goat",
		19: "horse",
		20: "snake",
		21: "dragon",
		22: "rabbit",
		23: "tiger",
		24: "ox",
		25: "rat",
		26: "pig",
		27: "dog",
		28: "rooster",
		29: "monkey",
		30: "goat",
		31: "horse",
		32: "snake",
		33: "dragon",
		34: "rabbit",
		35: "tiger",
		36: "ox",
		37: "rat",
		38: "pig",
		39: "dog",
		40: "rooster",
		41: "monkey",
		42: "goat",
		43: "horse",
		44: "snake",
		45: "dragon",
		46: "rabbit",
		47: "tiger",
		48: "ox",
		49: "rat",
	}, numberZodiacs)
}

func TestGetDutyZodiac(t *testing.T) {
	assert := assert.New(t)

	testcases := map[int]string{
		2020: "rat",
		2021: "ox",
		2022: "tiger",
		2023: "rabbit",
		2029: "rooster",
		2030: "dog",
		2035: "rabbit",
		2036: "dragon",
	}
	for year, zodiac := range testcases {
		dutyZodiac, err := lhczodiacwuxing.GetDutyZodiac(year)
		if !assert.NoError(err) {
			return
		}
		assert.Equal(zodiac, dutyZodiac)
	}
}
