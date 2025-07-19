package output

import (
	"github.com/fatih/color"
	"time"
)

func ShowDate() {
	today := time.Now().Format("Monday, January 2 2006")
	datePrint := color.New(color.FgMagenta, color.Bold).PrintFunc()

	datePrint("\n" + today + "\n")
}

func GetCurrentPrayers(prayers []Prayer) (string, string) {
	now := time.Now()
	curTime := now.Format("15:04")

	for idx, prayer := range prayers {
		curPrayerName, curPrayerTime := prayer.Name, prayer.Time

		if idx+1 == len(prayers) {
			return curPrayerName, "Fajr"
		}

		nextPrayerName, nextPrayerTime := prayers[idx+1].Name, prayers[idx+1].Time

		if curTime > curPrayerTime && curTime < nextPrayerTime {
			return curPrayerName, nextPrayerName
		}
	}
	return "Error", "Error"
}
