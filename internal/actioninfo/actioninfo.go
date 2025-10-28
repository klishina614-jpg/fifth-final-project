package actioninfo

import (
	"log"
)

// интерфейс для работы с данными
type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	// проходим по всем строкам с данными
	for _, dataString := range dataset {
		// разбираем строку
		err := dp.Parse(dataString)
		if err != nil {
			// если ошибка - пишем в лог и идем дальше
			log.Printf("Ошибка парсинга данных '%s': %v", dataString, err)
			continue
		}

		// получаем информацию об активности
		info, err := dp.ActionInfo()
		if err != nil {
			// если ошибка - пишем в лог и идем дальше
			log.Printf("Ошибка формирования информации: %v", err)
			continue
		}

		// выводим информацию
		log.Println(info)
	}
}
