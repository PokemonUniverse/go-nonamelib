package timeservice

import (
	"math/rand"
	"time"

	"github.com/PokemonUniverse/nonamelib/log"
)

type TimeService struct {
	isRunning         bool
	lastWeatherUpdate int64
	forcedWeather     int

	Day     int
	Hour    int
	Minutes int

	Weather int
}

func SYS_TIME() int64 {
	timeNano := float64(time.Now().UnixNano())
	return int64(timeNano * 0.000001) // Return milliseconds
}

func NewTimeService() *TimeService {
	timeService := &TimeService{isRunning: false,
		lastWeatherUpdate: SYS_TIME(),
		forcedWeather:     9,
		Day:               0,
		Hour:              0,
		Minutes:           0,
		Weather:           0}

	return timeService
}

func (t *TimeService) Start() {
	if !t.isRunning {
		// Calculate in-game time from real world time
		// this way we always have the right in-game time
		t.calculateTimeFromSystem()

		// Generate random starting weather
		t.generateWeather()

		t.isRunning = true

		t.run()
		log.Info("TimeService", "Start", "TimeService started")
	}
}

func (t *TimeService) Stop() {
	if t.isRunning {
		log.Info("TimeService", "Stop", "Stopping TimeService")
		t.isRunning = false
	}
}

func (t *TimeService) IsNight() bool {
	return (t.Hour >= 20 || t.Hour < 6)
}

func (t *TimeService) SetForcedWeather(_weather int) {
	t.forcedWeather = _weather
	t.lastWeatherUpdate = 0
}

func (t *TimeService) WeatherToStr() (str string) {
	str = "Normal"
	switch t.Weather {
	case 1:
		str = "Rain"
	case 2:
		str = "Hail"
	case 3:
		str = "Fog"
	case 4:
		str = "Sandstorm"
	}

	return
}

func (t *TimeService) generateWeather() {
	weather := t.forcedWeather
	if weather == 9 {
		weather = rand.Intn(10)
	}

	switch weather {
	case 0:
		t.Weather = WEATHER_NORMAL
	case 1:
		t.Weather = WEATHER_RAIN
	case 2:
		t.Weather = WEATHER_HAIL
	case 3:
		t.Weather = WEATHER_FOG
	case 4:
		t.Weather = WEATHER_SANDSTORM
	default:
		t.Weather = WEATHER_NORMAL
	}

	log.Info("TimeService", "Weather", "INFO: Weather type changed to: %v", t.WeatherToStr())
}

func (t *TimeService) calculateTimeFromSystem() {
	currentTime := time.Now()
	currentDay := currentTime.Weekday()
	currentHour := currentTime.Hour()
	currentMinute := currentTime.Minute()

	if currentDay == time.Monday {
		if currentHour < 12 {
			t.Day = int(time.Monday)
		} else {
			t.Day = int(time.Tuesday)
		}
	} else if currentDay == time.Tuesday {
		if currentHour < 12 {
			t.Day = int(time.Wednesday)
		} else {
			t.Day = int(time.Thursday)
		}
	} else if currentDay == time.Wednesday {
		if currentHour < 12 {
			t.Day = int(time.Friday)
		} else {
			t.Day = int(time.Saturday)
		}
	} else if currentDay == time.Thursday {
		if currentHour < 12 {
			t.Day = int(time.Sunday)
		} else {
			t.Day = int(time.Monday)
		}
	} else if currentDay == time.Friday {
		if currentHour < 12 {
			t.Day = int(time.Tuesday)
		} else {
			t.Day = int(time.Wednesday)
		}
	} else if currentDay == time.Saturday {
		if currentHour < 12 {
			t.Day = int(time.Thursday)
		} else {
			t.Day = int(time.Friday)
		}
	} else {
		if currentHour < 12 {
			t.Day = int(time.Saturday)
		} else {
			t.Day = int(time.Sunday)
		}
	}

	if currentHour >= 12 {
		currentHour = currentHour - 12
	}
	t.Hour = currentHour * 2

	if currentMinute >= 30 {
		currentMinute = currentMinute - 30

		// 30 min real world time, increment in-game hour by 1
		t.Hour++
	}
	t.Minutes = currentMinute * 2

	log.Info("TimeService", "INFO", "Current time is %v @ %d:%d", Days[t.Day], t.Hour, t.Minutes)
}

func (t *TimeService) run() {
	if t.isRunning {

		// In-game clock runs twice as fast as real time
		go func() {
			time.Sleep(time.Second * 30)
			t.run()
		}()

		currentMin := t.Minutes
		if currentMin == 59 {
			currentMin = 0
			currentHour := t.Hour

			if currentHour == 23 {
				currentHour = 0

				t.incrementDay()
			} else {
				currentHour++
			}

			t.Hour = currentHour
		} else {
			currentMin++
		}
		t.Minutes = currentMin

		weatherUpdate := SYS_TIME() - t.lastWeatherUpdate
		if weatherUpdate > 3600000 {
			t.generateWeather()
			t.lastWeatherUpdate = SYS_TIME()
		}

	} else {
		log.Info("TimeService", "Run", "TimeServie stopped")
	}
}

func (t *TimeService) incrementDay() {
	currentDay := t.Day

	if currentDay == 6 {
		currentDay = 0
	} else {
		currentDay++
	}

	t.Day = currentDay
}
