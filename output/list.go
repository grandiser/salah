package output

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/grandiser/salah/times"
)

type Prayer struct {
	Name string
	Time string
}

func OutputListAladhan(aladhanTimes times.AladhanAPIResponse) {
	fajr := &aladhanTimes.Data.Timings.Fajr
	sunrise := &aladhanTimes.Data.Timings.Sunrise
	dhuhr := &aladhanTimes.Data.Timings.Dhuhr
	asr := &aladhanTimes.Data.Timings.Asr
	maghrib := &aladhanTimes.Data.Timings.Maghrib
	isha := &aladhanTimes.Data.Timings.Isha

	prayers := []Prayer{
		{"Fajr", *fajr},
		{"Sunrise", *sunrise},
		{"Dhuhr", *dhuhr},
		{"Asr", *asr},
		{"Maghrib", *maghrib},
		{"Isha", *isha},
	}

	curPrayer, nextPrayer := GetCurrentPrayers(prayers)

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
