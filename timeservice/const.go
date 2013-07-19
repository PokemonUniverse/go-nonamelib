package timeservice

const (
	WEATHER_NORMAL int = iota
	WEATHER_RAIN
	WEATHER_HAIL
	WEATHER_SANDSTORM
	WEATHER_FOG
)

const (
	SUNDAY int = iota
	MONDAY
	TUESDAY
	WEDNESDAY
	THURSDAY
	FRIDAY
	SATURDAY
)

var Days = [...]string{
	"Sunday",
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
}
