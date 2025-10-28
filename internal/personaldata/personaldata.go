package personaldata

import "fmt"

type Personal struct {
	Name   string  // имя человека
	Weight float64 // вес в килограммах
	Height float64 // рост в метрах
}

func (p Personal) Print() {
	fmt.Printf("Имя: %s\nВес: %.2f кг.\nРост: %.2f м.\n", p.Name, p.Weight, p.Height)
}
