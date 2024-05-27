package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Car struct {
	Brand     string `json:"brand"`
	Milage    string `json:"milage"`
	ADAS      string `json:"adas"`
	Automatic string `json:"automatic"`
	Petrol    string `json:"petrol"`
	Diesel    string `json:"diesel"`
}

func main() {
	// Read the csv file
	file, err := os.Open("class016/cars.csv")
	if err != nil {
		panic(err)
	}

	defer file.Chdir()

	csvReader := csv.NewReader(file)
	count := 0
	cars := []Car{}
	for {
		srtX, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return
		}
		if count == 0 {
			count++
			continue
		}

		car := Car{
			Brand:     srtX[0],
			Milage:    srtX[1],
			ADAS:      srtX[2],
			Automatic: srtX[3],
			Petrol:    srtX[4],
			Diesel:    srtX[5],
		}
		cars = append(cars, car)
	}

	fmt.Printf("cars: %v\n", cars)
	ToJSON(cars)
}

func ToJSON(cars []Car) {
	bx, err := json.Marshal(cars)
	if err != nil {
		return
	}

	os.WriteFile("output.json", bx, 0644)
}
