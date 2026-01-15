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
		return errors.New("неверный формат данных? отсутствует шаги и продолжительность)")
	}

	// Парсим количество шагов (первый элемент)
	steps, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return errors.New("неверный формат количества шагов")
	}
	ds.Steps = steps

	// Парсим продолжительность (второй элемент)
	duration, err := time.ParseDuration(strings.TrimSpace(parts[1]))
	if err != nil {
		return errors.New("неверный формат продолжительности")
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
			"Вы сожгли %.2f ккал.",
		ds.Steps,
		distance,
		calories,
	)

	return result, nil
}
