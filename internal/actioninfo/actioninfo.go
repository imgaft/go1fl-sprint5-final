package actioninfo

import (
	"fmt"
	"log"
	"strings"
)

// DataParser определяет методы разбора и описания активности.
type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

// Info обрабатывает набор записей активности и выводит результаты.
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		if err := dp.Parse(data); err != nil {
			log.Println(err)
			continue
		}
		info, err := dp.ActionInfo()
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Print(info)
		if !strings.HasSuffix(info, "\n") {
			fmt.Println()
		}
	}
}
