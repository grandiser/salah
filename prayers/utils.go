package prayers

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/grandiser/salah/apis"
	"log"
	"math"
	"strings"
	"time"
)

func GetHijriDate(aladhanTimes apis.AladhanAPIResponse, config Config) string {
	// BUG: API Return Date Values Incorrect. Do not use until fixed.
	// Fix request made to AladhanAPI team

	var hijri string
	dayNum := aladhanTimes.Data.Date.Hijri.Day

	if config.UseArabic {
		arMonth := aladhanTimes.Data.Date.Hijri.Month.Ar
		year := aladhanTimes.Data.Date.Hijri.Year
		arWeekday := aladhanTimes.Data.Date.Hijri.Weekday.Ar
		hijri = fmt.Sprintf("%s، %s %s %s", arWeekday, dayNum, arMonth, year)
	} else {
		enMonth := aladhanTimes.Data.Date.Hijri.Month.En
		year := aladhanTimes.Data.Date.Hijri.Year
		enWeekday := aladhanTimes.Data.Date.Hijri.Weekday.En
		format := " %s, %s %s %s"
		hijri = fmt.Sprintf(format, enWeekday, dayNum, enMonth, year)
	}
	return color.New(color.FgCyan, color.Bold).Sprintf(hijri)
}

func ConvertStringToTime(timeStr string) time.Time {
	timeFormat := "15:04"

	timeTime, err := time.Parse(timeFormat, timeStr)
	if err != nil {
		log.Fatalln(err)
	}
	return timeTime
}

func GetCurrentPrayers(prayers []Prayer) (prevPrayer Prayer, nextPrayer Prayer) {
	now := time.Now()
	curTime := now.Format("15:04")

	for idx, prayer := range prayers {
		prevName, prevTime := prayer.Name, prayer.Time

		if idx+1 == len(prayers) {
			prevPrayer = Prayer{prevName, prevTime}
			nextPrayer = Prayer{"Fajr", prayers[0].Time}
			return prevPrayer, nextPrayer
		}

		nextName, nextTime := prayers[idx+1].Name, prayers[idx+1].Time

		if curTime > prevTime && curTime < nextTime {
			prevPrayer = Prayer{prevName, prevTime}
			nextPrayer = Prayer{nextName, nextTime}
			return prevPrayer, nextPrayer
		}
	}
	// fallback
	return prayers[0], prayers[1]
}

func GetLoadingSquares(prevPrayer Prayer, nextPrayer Prayer) string {
	loaded := "▣"
	unloaded := "▢"

	timeBetween := CalculateTimeDiff(prevPrayer, prevPrayer.Time, nextPrayer.Time)
	timeRemaining := GetTimeRemaining(prevPrayer, nextPrayer)
	nSquares := int(math.Ceil((float64(timeRemaining) / float64(timeBetween)) * 10))

	loadingSquares := strings.Repeat(loaded, 10-nSquares) + strings.Repeat(unloaded, nSquares)
	return loadingSquares
}

func CalculateTimeDiff(prevPrayer Prayer, prevTimeStr string, nextTimeStr string) time.Duration {
	prevTime := ConvertStringToTime(prevTimeStr)
	nextTime := ConvertStringToTime(nextTimeStr)

	timeDiff := nextTime.Sub(prevTime)

	if prevPrayer.Name == "Isha" && nextTime.Before(prevTime) {
		// Next prayer is tomorrow, add 24 hours
		tomorrow := nextTime.Add(24 * time.Hour)
		timeDiff = tomorrow.Sub(prevTime)
	} else {
		timeDiff = nextTime.Sub(prevTime)
	}

	return timeDiff
}

func GetTimeRemaining(prevPrayer Prayer, nextPrayer Prayer) time.Duration {
	timeFormat := "15:04"

	nowTimeStr := time.Now().Format(timeFormat)
	timeRemaining := CalculateTimeDiff(prevPrayer, nowTimeStr, nextPrayer.Time)

	return timeRemaining
}
