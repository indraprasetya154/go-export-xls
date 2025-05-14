package helper

import (
	"strconv"
	"time"
)

type Data struct {
	Name      string    `json:"name"`
	Category  string    `json:"category"`
	CreatedAt time.Time `json:"created_at"`
}

func GenerateData() []Data {
	data := make([]Data, 10)
	for i := 0; i < 10; i++ {
		data[i] = Data{
			Name:      "Item",
			Category:  "Category " + strconv.Itoa(i),
			CreatedAt: time.Now(),
		}
	}
	return data
}

func GenerateColumn(colNum int) string {
	result := ""
	for colNum > 0 {
		colNum--
		result = string(rune('A'+colNum%26)) + result
		colNum /= 26
	}
	return result
}
