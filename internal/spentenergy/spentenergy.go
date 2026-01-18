package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	// Проверяем входные параметры на корректность
	if steps <= 0 {
		return 0, errors.New("steps must be positive")
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

	// Рассчитываем среднюю скорость
	meanSpeed := MeanSpeed(steps, height, duration)

	// Переводим продолжительность в минуты
	minutes := duration.Minutes()

	// Рассчитываем количество калорий по формуле
	calories := (weight * meanSpeed * minutes) / minInH

	// Умножаем на коэффициент для ходьбы
	calories *= walkingCaloriesCoefficient

	return calories, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию расчёта калорий потраченных при беге
	// Проверяем входные параметры на корректность
	if steps <= 0 {
		return 0, errors.New("steps must be positive")
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

	// Рассчитываем среднюю скорость
	meanSpeed := MeanSpeed(steps, height, duration)

	// Переводим продолжительность в минуты
	minutes := duration.Minutes()

	// Рассчитываем количество калорий по формуле
	calories := (weight * meanSpeed * minutes) / minInH

	return calories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию расчёта средней скорости в км/ч
	// Проверяем, что продолжительность больше 0
	if duration <= 0 {
		return 0
	}

	// Рассчитываем дистанцию
	distance := Distance(steps, height)

	// Переводим продолжительность в часы
	hours := duration.Hours()

	// Рассчитываем среднюю скорость
	speed := distance / hours

	return speed
}

func Distance(steps int, height float64) float64 {
	// TODO: реализовать функцию расчёта пройденной дистанции в км
	// Рассчитываем длину шага в метрах
	stepLength := height * stepLengthCoefficient

	// Рассчитываем дистанцию в метрах и переводим в километры
	distanceInMeters := float64(steps) * stepLength
	distanceInKm := distanceInMeters / mInKm

	return distanceInKm
}
