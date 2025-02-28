package main
//Armenta Fuentes Lobsang Leonardo
import "fmt"

var decimalARoman = []struct {
	Value  int
	Symbol string
}{
	{1000, "M"}, {900, "CM"}, {500, "D"}, {400, "CD"},
	{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"},
	{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"},
}

func enteroARomano(num int) string {
	result := ""
	for _, entry := range decimalARoman {
		for num >= entry.Value {
			result += entry.Symbol
			num -= entry.Value
		}
	}
	return result
}

func main() {
	numerosEnteros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 42, 90, 100, 400, 500, 900, 1000, 3999}

	for _, num := range numerosEnteros {
		fmt.Printf("%d = %s\n", num, enteroARomano(num))
	}
}