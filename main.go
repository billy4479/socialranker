package main

import (
	"fmt"
	"os"
	"sort"
)

func getData() ([]float64, bool) {
	fmt.Println("Insert data:")

	data := []float64{}

	for {
		current := 0.0
		_, err := fmt.Scanf("%f", &current)
		if err != nil {
			break
		}
		data = append(data, current)
	}

	fmt.Print("Higher is better? (Y/n) ")
	answer := "Y"
	fmt.Scanf("%s", &answer)
	hib := answer != "n"
	return data, hib
}

func main() {

	data, hib := getData()
	originalData := make([]float64, len(data))
	copy(originalData, data)
	sort.Float64s(data)
	if hib {
		for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
			data[i], data[j] = data[j], data[i]
		}
	}

	last := data[0]

	firstIndex := 0

	translationTable := make(map[float64]float64)

	for i, v := range data {
		if last != v {
			diff := i - firstIndex
			avg := 0.0
			firstIndex++
			for firstIndex <= i {
				avg += float64(firstIndex)
				firstIndex++
			}

			avg /= float64(diff)

			translationTable[last] = avg
			firstIndex = i
		}
		last = v
	}

	{
		i := len(data)
		diff := i - firstIndex
		avg := 0.0
		firstIndex++
		for firstIndex <= i {
			avg += float64(firstIndex)
			firstIndex++
		}

		avg /= float64(diff)

		translationTable[last] = avg
	}

	for _, v := range originalData {
		r, ok := translationTable[v]
		if !ok {
			fmt.Printf("Error: %f is not in translationTable\n", v)
			os.Exit(1)
		}
		fmt.Println(r)
	}

}
