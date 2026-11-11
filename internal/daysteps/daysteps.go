package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

// DaySteps хранит сведения о дневной активности пользователя.
type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// Parse разбирает количество шагов и продолжительность активности.
func (ds *DaySteps) Parse(datastring string) error {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return fmt.Errorf("неверный формат данных")
	}
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return err
	}
	if steps <= 0 {
		return fmt.Errorf("количество шагов должно быть положительным")
	}
	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return err
	}
	if duration <= 0 {
		return fmt.Errorf("продолжительность должна быть положительной")
	}
	ds.Steps, ds.Duration = steps, duration
	return nil
}

// ActionInfo возвращает сведения о дневной активности.
func (ds DaySteps) ActionInfo() (string, error) {
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps,
		spentenergy.Distance(ds.Steps, ds.Height),
		calories,
	), nil
}
