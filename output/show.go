package output

import (
	"fmt"
	"github.com/fatih/color"
	"strings"
	"time"
)

func ShowDateGregorian() {
	today := time.Now().Format("Monday, January 2 2006")
	datePrint := color.New(color.FgMagenta, color.Bold).PrintFunc()

	datePrint("\n " + today + "\n")
}

func ShowDateHijri() {

}

func ShowPrayersList(prevPrayer Prayer, nextPrayer Prayer, prayers []Prayer) {
	formatter := "  %-7s : %s\n"
	prevPrayerPrint := color.New(color.FgCyan, color.Bold).PrintfFunc()
	nextPrayerPrint := color.New(color.FgGreen, color.Bold).PrintfFunc()

	for _, prayer := range prayers {
		if prevPrayer.Name == prayer.Name {
			prevPrayerPrint(formatter, prevPrayer.Name, prevPrayer.Time)

		} else if nextPrayer.Name == prayer.Name {
			nextPrayerPrint(formatter, nextPrayer.Name, nextPrayer.Time)

		} else {
			fmt.Printf(formatter, prayer.Name, prayer.Time)
		}
	}
}

func ShowPrayerLoader(prevPrayer Prayer, nextPrayer Prayer) {
	formatter := "\n Next: %s %s %s\n\n"
	nextPrayerPrint := color.New(color.FgMagenta, color.Bold).PrintfFunc()
	timeRemaining := strings.Replace(GetTimeRemaining(prevPrayer, nextPrayer).String(), "0s", "", 1)
	loadingSquares := GetLoadingSquares(prevPrayer, nextPrayer)
	nextPrayerPrint(formatter, nextPrayer.Name, loadingSquares, timeRemaining)
}

func ShowTimeRemaining(prevPrayer Prayer, nextPrayer Prayer) {
	timeRemainingPrint := color.New(color.FgYellow, color.Bold).PrintfFunc()
	timeRemaining := GetTimeRemaining(prevPrayer, nextPrayer).String()
	formatter := "\n %s%s%s\n"
	timeRemainingStr := strings.Replace(timeRemaining, "0s", "", 1)
	timeRemainingPrint(formatter, timeRemainingStr, " until ", nextPrayer.Name)

}
