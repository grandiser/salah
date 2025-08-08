package prayers

import (
	"fmt"
	"strings"
	"time"

	"github.com/grandiser/salah/apis"
)

func showDate(aladhanTimes apis.AladhanAPIResponse, config Config) {
	var todayDate string

	if config.HijriDate {
		todayDate = GetHijriDate(aladhanTimes, config)
	} else {
		todayDate = time.Now().Format("Monday, January 2 2006")
	}
	formattedDate := DateFormatter(todayDate)
	DatePrinter(formattedDate)
}

func ShowBasmalah() {
	basmalah := "﷽"
	formattedBasmalah := BasmalahFormatter(basmalah)
	BasmalahPrinter(formattedBasmalah)
}

func ShowPrayersList(prevPrayer Prayer, nextPrayer Prayer, prayers []Prayer, useArabic bool) {
	fmt.Println("   ╭────────۞────────╮")
	for _, prayer := range prayers {
		isPrev := prevPrayer.Name == prayer.Name
		isNext := nextPrayer.Name == prayer.Name
		formattedPrayer := PrayerFormatter(prayer.Name, prayer.Time, useArabic)
		coloredPrayer := PrayerColorer(formattedPrayer, isPrev, isNext)
		formattedTableLine := TableFormatter(coloredPrayer)
		fmt.Printf(formattedTableLine)
	}
	fmt.Println("   ╰────────۞────────╯\n")
}

func ShowPrayerLoader(prevPrayer Prayer, nextPrayer Prayer, useArabic bool) {
	timeRemaining := strings.Replace(GetTimeRemaining(prevPrayer, nextPrayer).String(), "0s", "", 1)
	loadingSquares := GetLoadingSquares(prevPrayer, nextPrayer)
	formattedLoader := LoaderFormatter(nextPrayer.Name, loadingSquares, timeRemaining, useArabic)
	LoaderPrinter(formattedLoader)
}
