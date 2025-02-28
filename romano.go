package main
//Armenta Fuentes Lobsang Leonardo
import (
	"fmt"
	"strings"
)

var romanDecimales = map[rune]int{
	'I': 1, 'V': 5, 'X': 10, 'L': 50,
	'C': 100, 'D': 500, 'M': 1000,
}

func romanoAEntero(romano string) int {
	romano = strings.ToUpper(romano)
	var result int
	prevValue := 0

	for i := len(romano) - 1; i >= 0; i-- {
		value := romanDecimales[rune(romano[i])]
		if value < prevValue {
			result -= value
		} else {
			result += value
		}
		prevValue = value
	}

	return result
}

func main() {
	numerosRomanos := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X", "XLII", "XC", "C", "CD", "D", "CM", "M", "MMMCMXCIX"}

	for _, romano := range numerosRomanos {
		fmt.Printf("%s = %d\n", romano, romanoAEntero(romano))
	}
}