package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/mrdevelopgo/fitTracker2/internal/personaldata"
	"github.com/mrdevelopgo/fitTracker2/internal/spentenergy"
)

type Training struct {
	// TODO: добавить поля
	Steps                 int           // количество шагов, проделанных за тренировку
	TrainingType          string        // тип тренировки (бег или ходьба)
	Duration              time.Duration // длительность тренировки
	personaldata.Personal               // встроенная структура Personal
}

func (t *Training) Parse(datastring string) (err error) {
	// TODO: реализовать функцию

	// Разделяем строку по запятым
	parts := strings.Split(datastring, ",")

	// Проверяем, что получили 3 части
	if len(parts) != 3 {
		return errors.New("неверный формат данных: ожидается 3 значения")
	}

	// Парсим количество шагов
	stepsStr := strings.TrimSpace(parts[0])
	steps, err := strconv.Atoi(stepsStr)
	if err != nil {
		return errors.New("неверный формат количества шагов")
	}

	// Проверяем, что шаги > 0
	if steps <= 0 {
		return errors.New("неверный формат количества шагов")
	}

	t.Steps = steps

	// Сохраняем тип тренировки
	t.TrainingType = strings.TrimSpace(parts[1])

	// Парсим продолжительность
	durationStr := strings.TrimSpace(parts[2])
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return errors.New("неверный формат продолжительности")
	}

	// Проверяем, что продолжительность > 0
	if duration <= 0 {
		return errors.New("неверный формат продолжительности")
	}

	t.Duration = duration

	return nil
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию

	// Рассчитываем дистанцию
	distance := spentenergy.Distance(t.Steps, t.Height)

	// Рассчитываем среднюю скорость
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var calories float64
	var err error

	// Рассчитываем калории в зависимости от типа тренировки
	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(
			t.Steps, t.Weight, t.Height, t.Duration,
		)
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(
			t.Steps, t.Weight, t.Height, t.Duration,
		)
	default:
		return "", errors.New("неизвестный тип тренировки")
	}

	// Проверяем ошибку от функции расчета калорий
	if err != nil {
		return "", err
	}

	// Форматируем продолжительность в часах с двумя знаками после запятой
	// durationHours := t.Duration.Hours()

	// Формируем результат
	result := fmt.Sprintf(
		"Тип тренировки: %s\n"+
			"Длительность: %.2f ч.\n"+
			"Дистанция: %.2f км.\n"+
			"Скорость: %.2f км/ч\n"+
			"Сожгли калорий: %.2f\n",
		t.TrainingType,
		t.Duration.Hours(),
		distance,
		speed,
		calories,
	)

	return result, nil
}
