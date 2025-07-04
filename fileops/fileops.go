package fileops
import (
	"fmt"
	"os"
	"errors"
	"strconv"
)
func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return 1000, errors.New("Failed to find file.")
	}
	valueText := string(data) // байты в строку
	value, err := strconv.ParseFloat(valueText, 64) // строку в число с плавающей точкой
	
	if err != nil {
		return 1000, errors.New("Failed to parse stored value")
	}

	return value, nil
}
func WriteFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}