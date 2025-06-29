package main

/* 
import "fmt"
import "math"
^ можно объединить в один импорт >>
*/
import (
	"fmt"
	"math"
)
func main() {
	// 1 Способ.
	/*
	var investmentAmount = 1000
	var expectedReturnRate = 5.5
	var years = 10

	важно соблюдать соответствие типов, участвующих в операции
	float64 приводит к числу с плавающей запятой
	var futureValue = float64(investmentAmount) * math.Pow(1 + expectedReturnRate / 100, float64(years))
	*/
	



	// 2 Способ.
	/*
	так же можно сразу в объявлении задать тип переменной и объединить несколько объявлений в одно >>
	если тип не указывается, то можно опустить ключевое слово var, используя := >>
	*/
	// var investmentAmount, years float64 = 1000, 10
	// expectedReturnRate := 5.5 


	// 3 Способ.
	// ещё более простой путь - сразу присвоить переменным значения с плавающей запятой, однако это делает код менее читабельным >>
	const inflationRate = 2.5
	investmentAmount, years, expectedReturnRate := 1000.0, 10.0, 5.5

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount) // ожидаёт ввод в терминале

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate) // ожидаёт ввод в терминале

	fmt.Print("Investment Horizon: ")
	fmt.Scan(&years) // ожидаёт ввод в терминале
	
	futureValue := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	futureRealValue := futureValue / math.Pow(1 + inflationRate/100, years)

	// fmt.Println("Future Value: ", futureValue)
	// fmt.Println("Future Value (adjusted for inflation): ", futureRealValue)
	// fmt.Printf("Future Value: %.1f\nFuture Value (adjusted for inflation): %.1f", futureValue, futureRealValue)
	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedFRV := fmt.Sprintf("Future Value (adjusted for inflation): %.1f\n", futureRealValue)

	fmt.Print(formattedFV, formattedFRV)
}








// go mod init example.com/investment-calculator - инициализирует модуль

// go build - создает исполняемый файл

// go run ./investment_calculator.go - запускает программу