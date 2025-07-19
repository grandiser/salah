package output

import (
	"fmt"
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

func ShowPrayerTimes(curPrayer string, nextPrayer string, prayers []Prayer) {
	formatter := "   %-7s : %s\n"
	curPrayerPrint := color.New(color.FgGreen, color.Bold).PrintfFunc()
	nextPrayerPrint := color.New(color.FgYellow, color.Bold).PrintfFunc()

	for _, prayer := range prayers {
		if curPrayer == prayer.Name {
			curPrayerPrint(formatter, curPrayer, prayer.Time)

		} else if nextPrayer == prayer.Name {
			nextPrayerPrint(formatter, nextPrayer, prayer.Time)

		} else {
			fmt.Printf(formatter, prayer.Name, prayer.Time)
		}
	}
	fmt.Printf("\n")
}
