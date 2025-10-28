package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	// разбиваем строку на части
	parts := strings.Split(datastring, ",")

	// проверяем что частей ровно 2
	if len(parts) != 2 {
		return errors.New("invalid data format: expected 'steps,duration'")
	}

	// переводим шаги в число
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return err
	}
	if steps <= 0 {
		return errors.New("steps count must be positive")
	}
	ds.Steps = steps

	// переводим время в правильный формат
	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return err
	}
	if duration <= 0 {
		return errors.New("duration must be positive")
	}
	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	// считаем расстояние
	distance := spentenergy.Distance(ds.Steps, ds.Height)

	// считаем калории (используем функцию для ходьбы)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	// делаем красивую строку с информацией
	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps, distance, calories)

	return result, nil
}

func (ds DaySteps) Print() {
	ds.Personal.Print()
}
