package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	// разбиваем строку на части
	parts := strings.Split(datastring, ",")

	// проверяем что частей ровно 3
	if len(parts) != 3 {
		return errors.New("invalid data format: expected 'steps,training_type,duration'")
	}

	// переводим шаги в число
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return err
	}
	if steps <= 0 {
		return errors.New("steps count must be positive")
	}
	t.Steps = steps

	// сохраняем тип тренировки
	t.TrainingType = parts[1]

	// переводим время в правильный формат
	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return err
	}
	if duration <= 0 {
		return errors.New("duration must be positive")
	}
	t.Duration = duration

	return nil
}

func (t Training) ActionInfo() (string, error) {
	// считаем расстояние
	distance := spentenergy.Distance(t.Steps, t.Height)

	// считаем скорость
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var calories float64
	var err error

	// проверяем тип тренировки и считаем калории
	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", errors.New("unknown training type")
	}

	// если была ошибка при подсчете калорий
	if err != nil {
		return "", err
	}

	// делаем красивую строку с информацией
	durationInHours := t.Duration.Hours()
	result := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType, durationInHours, distance, meanSpeed, calories)

	return result, nil
}

func (t Training) Print() {
	t.Personal.Print()
}
