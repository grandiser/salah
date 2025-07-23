package prayers

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/grandiser/salah/apis"
	"strings"
	"time"
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

func ShowPrayersList(prevPrayer Prayer, nextPrayer Prayer, prayers []Prayer) {
	prayerFormat := "%-7s : %s"
	tableFormat := "   │ %s │\n"
	boldGreen := color.New(color.FgHiCyan, color.Bold).SprintFunc()
	boldMagenta := color.New(color.FgHiBlue, color.Bold).SprintFunc()

	fmt.Println("   ╭────────۞────────╮")
	for _, prayer := range prayers {
		if prevPrayer.Name == prayer.Name {
			formattedString := fmt.Sprintf(prayerFormat, prevPrayer.Name, prevPrayer.Time)
			coloredString := boldGreen(formattedString)

			fmt.Printf(tableFormat, coloredString)

		} else if nextPrayer.Name == prayer.Name {
			formattedString := fmt.Sprintf(prayerFormat, nextPrayer.Name, nextPrayer.Time)
			coloredString := boldMagenta(formattedString)

			fmt.Printf(tableFormat, coloredString)

		} else {
			prayerString := fmt.Sprintf(prayerFormat, prayer.Name, prayer.Time)
			fmt.Printf(tableFormat, prayerString)
		}
	}
	fmt.Println("   ╰────────۞────────╯\n")
}

func ShowPrayerLoader(prevPrayer Prayer, nextPrayer Prayer) {
	timeRemaining := strings.Replace(GetTimeRemaining(prevPrayer, nextPrayer).String(), "0s", "", 1)
	loadingSquares := GetLoadingSquares(prevPrayer, nextPrayer)
	formattedLoader := LoaderFormatter(nextPrayer.Name, loadingSquares, timeRemaining)
	LoaderPrinter(formattedLoader)
}
