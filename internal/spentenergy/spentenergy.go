package spentenergy

import (
	"errors"
	"time"
)

// константы для расчетов
const (
	mInKm                      = 1000 // метров в километре
	minInH                     = 60   // минут в часе
	stepLengthCoefficient      = 0.45 // коэффициент для длины шага
	walkingCaloriesCoefficient = 0.5  // коэффициент для ходьбы
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// проверяем что все данные правильные
	if steps <= 0 {
		return 0, errors.New("steps count must be positive")
	}
	if weight <= 0 {
		return 0, errors.New("weight must be positive")
	}
	if height <= 0 {
		return 0, errors.New("height must be positive")
	}
	if duration <= 0 {
		return 0, errors.New("duration must be positive")
	}

	// считаем скорость
	meanSpeed := MeanSpeed(steps, height, duration)

	// переводим время в минуты
	durationInMinutes := duration.Minutes()

	// считаем калории
	calories := (weight * meanSpeed * durationInMinutes) / minInH

	// применяем коэффициент для ходьбы
	walkingCalories := calories * walkingCaloriesCoefficient

	return walkingCalories, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// проверяем что все данные правильные
	if steps <= 0 {
		return 0, errors.New("steps count must be positive")
	}
	if weight <= 0 {
		return 0, errors.New("weight must be positive")
	}
	if height <= 0 {
		return 0, errors.New("height must be positive")
	}
	if duration <= 0 {
		return 0, errors.New("duration must be positive")
	}

	// считаем скорость
	meanSpeed := MeanSpeed(steps, height, duration)

	// переводим время в минуты
	durationInMinutes := duration.Minutes()

	// считаем калории
	calories := (weight * meanSpeed * durationInMinutes) / minInH

	return calories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// проверяем что время больше нуля
	if duration <= 0 {
		return 0
	}

	// проверяем что шаги не отрицательные
	if steps < 0 {
		return 0
	}

	// считаем расстояние
	distance := Distance(steps, height)

	// считаем скорость в км/ч
	durationInHours := duration.Hours()
	return distance / durationInHours
}

func Distance(steps int, height float64) float64 {
	// считаем длину шага
	stepLength := height * stepLengthCoefficient

	// считаем расстояние в метрах
	distanceInMeters := float64(steps) * stepLength

	// переводим в километры
	return distanceInMeters / mInKm
}
