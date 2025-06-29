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


const inflationRate = 2.5

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
	investmentAmount, years, expectedReturnRate := 1000.0, 10.0, 5.5

	// fmt.Print("Investment Amount: ")
	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount) // ожидаёт ввод в терминале

	// fmt.Print("Expected Return Rate: ")
	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate) // ожидаёт ввод в терминале

	// fmt.Print("Investment Horizon: ")
	outputText("Investment Horizon: ")
	fmt.Scan(&years) // ожидаёт ввод в терминале
	
	// futureValue := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	// futureRealValue := futureValue / math.Pow(1 + inflationRate/100, years)
	futureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturnRate, years)



	// fmt.Println("Future Value: ", futureValue)
	// fmt.Println("Future Value (adjusted for inflation): ", futureRealValue)
	// fmt.Printf("Future Value: %.1f\nFuture Value (adjusted for inflation): %.1f", futureValue, futureRealValue)
	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedFRV := fmt.Sprintf("Future Value (adjusted for inflation): %.1f\n", futureRealValue)

	fmt.Print(formattedFV, formattedFRV)
}


func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) (fv float64, frv float64) {
	fv = investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	frv = fv / math.Pow(1 + inflationRate/100, years)
	return fv, frv
}


// go mod init example.com/investment-calculator - инициализирует модуль

// go build - создает исполняемый файл

// go run ./investment_calculator.go - запускает программу