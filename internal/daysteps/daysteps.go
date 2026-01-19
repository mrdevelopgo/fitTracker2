package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/mrdevelopgo/fitTracker2/internal/personaldata"
	"github.com/mrdevelopgo/fitTracker2/internal/spentenergy"
)

type DaySteps struct {
	// TODO: добавить поля
	Steps                 int           // количество шагов
	Duration              time.Duration // длительность прогулки
	personaldata.Personal               // встроенная структура Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	// TODO: реализовать функцию

	// Разделяем строку по запятым
	parts := strings.Split(datastring, ",")

	// Проверяем, что получили 2 части (шаги и продолжительность)
	if len(parts) != 2 {
		return errors.New("invalid data format: expected 2 values")
	}

	// Парсим количество шагов (первый элемент)
	stepsStr := parts[0]
	steps, err := strconv.Atoi(stepsStr)
	if err != nil {
		return err
	}

	// Проверяем, что шаги > 0
	if steps <= 0 {
		return errors.New("steps must be positive")
	}

	ds.Steps = steps

	// Парсим продолжительность (второй элемент), убираем пробелы
	durationStr := strings.TrimSpace(parts[1])
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return err
	}

	// Проверяем, что продолжительность > 0
	if duration <= 0 {
		return fmt.Errorf("duration must be positive")
	}

	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: реализовать функцию

	// Рассчитываем дистанцию
	distance := spentenergy.Distance(ds.Steps, ds.Height)

	// Для прогулок используем функцию WalkingSpentCalories
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)

	// Проверяем ошибку от функции расчета калорий
	if err != nil {
		return "", err
	}

	// Формируем результат
	result := fmt.Sprintf(
		"Количество шагов: %d.\n"+
			"Дистанция составила %.2f км.\n"+
			"Вы сожгли %.2f ккал.\n",
		ds.Steps,
		distance,
		calories,
	)

	return result, nil
}
