package lhczodiacwuxing

import (
	"fmt"
)

var (
	yearZodiacNumbers = map[int]map[string][]int{
		2017: map[string][]int{
			"rat":     []int{10, 22, 34, 46},
			"ox":      []int{9, 21, 33, 45},
			"tiger":   []int{8, 20, 32, 44},
			"rabbit":  []int{7, 19, 31, 43},
			"dragon":  []int{6, 18, 30, 42},
			"snake":   []int{5, 17, 29, 41},
			"horse":   []int{4, 16, 28, 40},
			"goat":    []int{3, 15, 27, 39},
			"monkey":  []int{2, 14, 26, 38},
			"rooster": []int{1, 13, 25, 37, 49},
			"dog":     []int{12, 24, 36, 48},
			"pig":     []int{11, 23, 35, 47},
		},
		2018: map[string][]int{
			"rat":     []int{11, 23, 35, 47},
			"ox":      []int{10, 22, 34, 46},
			"tiger":   []int{9, 21, 33, 45},
			"rabbit":  []int{8, 20, 32, 44},
			"dragon":  []int{7, 19, 31, 43},
			"snake":   []int{6, 18, 30, 42},
			"horse":   []int{5, 17, 29, 41},
			"goat":    []int{4, 16, 28, 40},
			"monkey":  []int{3, 15, 27, 39},
			"rooster": []int{2, 14, 26, 38},
			"dog":     []int{1, 13, 25, 37, 49},
			"pig":     []int{12, 24, 36, 48},
		},
		2019: map[string][]int{
			"rat":     []int{12, 24, 36, 48},
			"ox":      []int{11, 23, 35, 47},
			"tiger":   []int{10, 22, 34, 46},
			"rabbit":  []int{9, 21, 33, 45},
			"dragon":  []int{8, 20, 32, 44},
			"snake":   []int{7, 19, 31, 43},
			"horse":   []int{6, 18, 30, 42},
			"goat":    []int{5, 17, 29, 41},
			"monkey":  []int{4, 16, 28, 40},
			"rooster": []int{3, 15, 27, 39},
			"dog":     []int{2, 14, 26, 38},
			"pig":     []int{1, 13, 25, 37, 49},
		},
		2020: map[string][]int{
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
		},
		2021: map[string][]int{
			"rat":     []int{2, 14, 26, 38},
			"ox":      []int{1, 13, 25, 37, 49},
			"tiger":   []int{12, 24, 36, 48},
			"rabbit":  []int{11, 23, 35, 47},
			"dragon":  []int{10, 22, 34, 46},
			"snake":   []int{9, 21, 33, 45},
			"horse":   []int{8, 20, 32, 44},
			"goat":    []int{7, 19, 31, 43},
			"monkey":  []int{6, 18, 30, 42},
			"rooster": []int{5, 17, 29, 41},
			"dog":     []int{4, 16, 28, 40},
			"pig":     []int{3, 15, 27, 39},
		},
		2022: map[string][]int{
			"rat":     []int{3, 15, 27, 39},
			"ox":      []int{2, 14, 26, 38},
			"tiger":   []int{1, 13, 25, 37, 49},
			"rabbit":  []int{12, 24, 36, 48},
			"dragon":  []int{11, 23, 35, 47},
			"snake":   []int{10, 22, 34, 46},
			"horse":   []int{9, 21, 33, 45},
			"goat":    []int{8, 20, 32, 44},
			"monkey":  []int{7, 19, 31, 43},
			"rooster": []int{6, 18, 30, 42},
			"dog":     []int{5, 17, 29, 41},
			"pig":     []int{4, 16, 28, 40},
		},
		2023: map[string][]int{
			"rat":     []int{4, 16, 28, 40},
			"ox":      []int{3, 15, 27, 39},
			"tiger":   []int{2, 14, 26, 38},
			"rabbit":  []int{1, 13, 25, 37, 49},
			"dragon":  []int{12, 24, 36, 48},
			"snake":   []int{11, 23, 35, 47},
			"horse":   []int{10, 22, 34, 46},
			"goat":    []int{9, 21, 33, 45},
			"monkey":  []int{8, 20, 32, 44},
			"rooster": []int{7, 19, 31, 43},
			"dog":     []int{6, 18, 30, 42},
			"pig":     []int{5, 17, 29, 41},
		},
		2024: map[string][]int{
			"rat":     []int{5, 17, 29, 41},
			"ox":      []int{4, 16, 28, 40},
			"tiger":   []int{3, 15, 27, 39},
			"rabbit":  []int{2, 14, 26, 38},
			"dragon":  []int{1, 13, 25, 37, 49},
			"snake":   []int{12, 24, 36, 48},
			"horse":   []int{11, 23, 35, 47},
			"goat":    []int{10, 22, 34, 46},
			"monkey":  []int{9, 21, 33, 45},
			"rooster": []int{8, 20, 32, 44},
			"dog":     []int{7, 19, 31, 43},
			"pig":     []int{6, 18, 30, 42},
		},
		2025: map[string][]int{
			"rat":     []int{6, 18, 30, 42},
			"ox":      []int{5, 17, 29, 41},
			"tiger":   []int{4, 16, 28, 40},
			"rabbit":  []int{3, 15, 27, 39},
			"dragon":  []int{2, 14, 26, 38},
			"snake":   []int{1, 13, 25, 37, 49},
			"horse":   []int{12, 24, 36, 48},
			"goat":    []int{11, 23, 35, 47},
			"monkey":  []int{10, 22, 34, 46},
			"rooster": []int{9, 21, 33, 45},
			"dog":     []int{8, 20, 32, 44},
			"pig":     []int{7, 19, 31, 43},
		},
		2026: map[string][]int{
			"rat":     []int{7, 19, 31, 43},
			"ox":      []int{6, 18, 30, 42},
			"tiger":   []int{5, 17, 29, 41},
			"rabbit":  []int{4, 16, 28, 40},
			"dragon":  []int{3, 15, 27, 39},
			"snake":   []int{2, 14, 26, 38},
			"horse":   []int{1, 13, 25, 37, 49},
			"goat":    []int{12, 24, 36, 48},
			"monkey":  []int{11, 23, 35, 47},
			"rooster": []int{10, 22, 34, 46},
			"dog":     []int{9, 21, 33, 45},
			"pig":     []int{8, 20, 32, 44},
		},
		2027: map[string][]int{
			"rat":     []int{8, 20, 32, 44},
			"ox":      []int{7, 19, 31, 43},
			"tiger":   []int{6, 18, 30, 42},
			"rabbit":  []int{5, 17, 29, 41},
			"dragon":  []int{4, 16, 28, 40},
			"snake":   []int{3, 15, 27, 39},
			"horse":   []int{2, 14, 26, 38},
			"goat":    []int{1, 13, 25, 37, 49},
			"monkey":  []int{12, 24, 36, 48},
			"rooster": []int{11, 23, 35, 47},
			"dog":     []int{10, 22, 34, 46},
			"pig":     []int{9, 21, 33, 45},
		},
		2028: map[string][]int{
			"rat":     []int{9, 21, 33, 45},
			"ox":      []int{8, 20, 32, 44},
			"tiger":   []int{7, 19, 31, 43},
			"rabbit":  []int{6, 18, 30, 42},
			"dragon":  []int{5, 17, 29, 41},
			"snake":   []int{4, 16, 28, 40},
			"horse":   []int{3, 15, 27, 39},
			"goat":    []int{2, 14, 26, 38},
			"monkey":  []int{1, 13, 25, 37, 49},
			"rooster": []int{12, 24, 36, 48},
			"dog":     []int{11, 23, 35, 47},
			"pig":     []int{10, 22, 34, 46},
		},
		2029: map[string][]int{
			"rat":     []int{10, 22, 34, 46},
			"ox":      []int{9, 21, 33, 45},
			"tiger":   []int{8, 20, 32, 44},
			"rabbit":  []int{7, 19, 31, 43},
			"dragon":  []int{6, 18, 30, 42},
			"snake":   []int{5, 17, 29, 41},
			"horse":   []int{4, 16, 28, 40},
			"goat":    []int{3, 15, 27, 39},
			"monkey":  []int{2, 14, 26, 38},
			"rooster": []int{1, 13, 25, 37, 49},
			"dog":     []int{12, 24, 36, 48},
			"pig":     []int{11, 23, 35, 47},
		},
		2030: map[string][]int{
			"rat":     []int{11, 23, 35, 47},
			"ox":      []int{10, 22, 34, 46},
			"tiger":   []int{9, 21, 33, 45},
			"rabbit":  []int{8, 20, 32, 44},
			"dragon":  []int{7, 19, 31, 43},
			"snake":   []int{6, 18, 30, 42},
			"horse":   []int{5, 17, 29, 41},
			"goat":    []int{4, 16, 28, 40},
			"monkey":  []int{3, 15, 27, 39},
			"rooster": []int{2, 14, 26, 38},
			"dog":     []int{1, 13, 25, 37, 49},
			"pig":     []int{12, 24, 36, 48},
		},
		2031: map[string][]int{
			"rat":     []int{12, 24, 36, 48},
			"ox":      []int{11, 23, 35, 47},
			"tiger":   []int{10, 22, 34, 46},
			"rabbit":  []int{9, 21, 33, 45},
			"dragon":  []int{8, 20, 32, 44},
			"snake":   []int{7, 19, 31, 43},
			"horse":   []int{6, 18, 30, 42},
			"goat":    []int{5, 17, 29, 41},
			"monkey":  []int{4, 16, 28, 40},
			"rooster": []int{3, 15, 27, 39},
			"dog":     []int{2, 14, 26, 38},
			"pig":     []int{1, 13, 25, 37, 49},
		},
		2032: map[string][]int{
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
		},
		2033: map[string][]int{
			"rat":     []int{2, 14, 26, 38},
			"ox":      []int{1, 13, 25, 37, 49},
			"tiger":   []int{12, 24, 36, 48},
			"rabbit":  []int{11, 23, 35, 47},
			"dragon":  []int{10, 22, 34, 46},
			"snake":   []int{9, 21, 33, 45},
			"horse":   []int{8, 20, 32, 44},
			"goat":    []int{7, 19, 31, 43},
			"monkey":  []int{6, 18, 30, 42},
			"rooster": []int{5, 17, 29, 41},
			"dog":     []int{4, 16, 28, 40},
			"pig":     []int{3, 15, 27, 39},
		},
		2034: map[string][]int{
			"rat":     []int{3, 15, 27, 39},
			"ox":      []int{2, 14, 26, 38},
			"tiger":   []int{1, 13, 25, 37, 49},
			"rabbit":  []int{12, 24, 36, 48},
			"dragon":  []int{11, 23, 35, 47},
			"snake":   []int{10, 22, 34, 46},
			"horse":   []int{9, 21, 33, 45},
			"goat":    []int{8, 20, 32, 44},
			"monkey":  []int{7, 19, 31, 43},
			"rooster": []int{6, 18, 30, 42},
			"dog":     []int{5, 17, 29, 41},
			"pig":     []int{4, 16, 28, 40},
		},
		2035: map[string][]int{
			"rat":     []int{4, 16, 28, 40},
			"ox":      []int{3, 15, 27, 39},
			"tiger":   []int{2, 14, 26, 38},
			"rabbit":  []int{1, 13, 25, 37, 49},
			"dragon":  []int{12, 24, 36, 48},
			"snake":   []int{11, 23, 35, 47},
			"horse":   []int{10, 22, 34, 46},
			"goat":    []int{9, 21, 33, 45},
			"monkey":  []int{8, 20, 32, 44},
			"rooster": []int{7, 19, 31, 43},
			"dog":     []int{6, 18, 30, 42},
			"pig":     []int{5, 17, 29, 41},
		},
		2036: map[string][]int{
			"rat":     []int{5, 17, 29, 41},
			"ox":      []int{4, 16, 28, 40},
			"tiger":   []int{3, 15, 27, 39},
			"rabbit":  []int{2, 14, 26, 38},
			"dragon":  []int{1, 13, 25, 37, 49},
			"snake":   []int{12, 24, 36, 48},
			"horse":   []int{11, 23, 35, 47},
			"goat":    []int{10, 22, 34, 46},
			"monkey":  []int{9, 21, 33, 45},
			"rooster": []int{8, 20, 32, 44},
			"dog":     []int{7, 19, 31, 43},
			"pig":     []int{6, 18, 30, 42},
		},
	}
	yearNumberZodiacs = map[int]map[int]string{}
)

// GetZodiacNumbers 取得 x 年的生肖號碼
func GetZodiacNumbers(year int) (map[string][]int, error) {
	zodiacNumbers, ok := yearZodiacNumbers[year]
	if !ok {
		return nil, fmt.Errorf("Not configured zodiac config: %d year", year)
	}

	return zodiacNumbers, nil
}

// GetNumberZodiacs 取得 x 年的號碼生肖
func GetNumberZodiacs(year int) (map[int]string, error) {
	_, ok := yearNumberZodiacs[year]
	if !ok {
		zodiacNumbers, err := GetZodiacNumbers(year)
		if err != nil {
			return nil, err
		}

		numberZodiacs := map[int]string{}
		for zodiac, numbers := range zodiacNumbers {
			for _, number := range numbers {
				numberZodiacs[number] = zodiac
			}
		}

		yearNumberZodiacs[year] = numberZodiacs
	}

	return yearNumberZodiacs[year], nil
}

// GetDutyZodiac 取得 x 年的生肖
func GetDutyZodiac(year int) (string, error) {
	numberZodiacs, err := GetNumberZodiacs(year)
	if err != nil {
		return "", err
	}

	return numberZodiacs[1], nil
}
