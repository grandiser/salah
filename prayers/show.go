package prayers

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/grandiser/salah/apis"
	"strings"
	"time"
)

func showDate(aladhanTimes apis.AladhanAPIResponse, config Config) {
	if config.HijriDate {
		ShowDateHijri(aladhanTimes, config)
	}

	if config.GregorianDate {
		ShowDateGregorian()
	}

}

func ShowDateGregorian() {
	today := time.Now().Format("Monday, January 2 2006")
	datePrint := color.New(color.FgCyan, color.Bold).PrintFunc()

	datePrint(" " + today + "\n")
}

func ShowDateHijri(aladhanTimes apis.AladhanAPIResponse, config Config) {
	hijriDate := GetHijriDate(aladhanTimes, config)
	fmt.Println("" + hijriDate)
}
func ShowBasmalah() {
	fmt.Println("       ﷽   ")
}

func ShowPrayersList(prevPrayer Prayer, nextPrayer Prayer, prayers []Prayer) {
	prayerFormat := "%-7s : %s"
	tableFormat := "   │ %s │\n"

	fmt.Println("   ╭────────۞────────╮")
	for _, prayer := range prayers {
		if prevPrayer.Name == prayer.Name {
			formattedString := fmt.Sprintf(prayerFormat, prevPrayer.Name, prevPrayer.Time)
			boldGreen := color.New(color.FgGreen, color.Bold).SprintFunc()
			coloredString := boldGreen(formattedString)

			fmt.Printf(tableFormat, coloredString)

		} else if nextPrayer.Name == prayer.Name {
			formattedString := fmt.Sprintf(prayerFormat, nextPrayer.Name, nextPrayer.Time)
			boldMagenta := color.New(color.FgMagenta, color.Bold).SprintFunc()
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
	formatter := " Next: %s %s %s\n\n"
	nextPrayerPrint := color.New(color.FgMagenta, color.Bold).PrintfFunc()
	timeRemaining := strings.Replace(GetTimeRemaining(prevPrayer, nextPrayer).String(), "0s", "", 1)
	loadingSquares := GetLoadingSquares(prevPrayer, nextPrayer)
	nextPrayerPrint(formatter, nextPrayer.Name, loadingSquares, timeRemaining)
}
