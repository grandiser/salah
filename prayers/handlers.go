package prayers

import (
	"fmt"

	"github.com/01walid/goarabic"
	"github.com/grandiser/salah/apis"
)

type Prayer struct {
	Name string
	Time string
}

func localizedName(englishName string, useArabic bool) string {
	if !useArabic {
		return englishName
	}

	arabicNames := map[string]string{
		"Fajr":    "الفجر",
		"Sunrise": "الشروق",
		"Dhuhr":   "الظهر",
		"Asr":     "العصر",
		"Maghrib": "المغرب",
		"Isha":    "العشاء",
	}

	name, exists := arabicNames[englishName]
	if !exists {
		name = englishName
	}

	if userOS == "windows" {
		name = goarabic.Reverse(name)
	}

	return name
}

func AladhanHandler(aladhanTimes apis.AladhanAPIResponse, config Config) {
	if config.Compact {
		SingleAladhan(aladhanTimes, config)
	} else {
		ListAladhan(aladhanTimes, config)
	}
}

func ListAladhan(aladhanTimes apis.AladhanAPIResponse, config Config) {
	prayers := []Prayer{
		{localizedName("Fajr", config.UseArabic), aladhanTimes.Data.Timings.Fajr},
		{localizedName("Sunrise", config.UseArabic), aladhanTimes.Data.Timings.Sunrise},
		{localizedName("Dhuhr", config.UseArabic), aladhanTimes.Data.Timings.Dhuhr},
		{localizedName("Asr", config.UseArabic), aladhanTimes.Data.Timings.Asr},
		{localizedName("Maghrib", config.UseArabic), aladhanTimes.Data.Timings.Maghrib},
		{localizedName("Isha", config.UseArabic), aladhanTimes.Data.Timings.Isha},
	}

	prevPrayer, nextPrayer := GetCurrentPrayers(prayers)

	fmt.Printf("\n")
	ShowBasmalah()
	ShowPrayersList(prevPrayer, nextPrayer, prayers, config.UseArabic)
	showDate(aladhanTimes, config)
	ShowPrayerLoader(prevPrayer, nextPrayer, config.UseArabic)
	fmt.Printf("\n")
}

func SingleAladhan(aladhanTimes apis.AladhanAPIResponse, config Config) {
	prayers := []Prayer{
		{localizedName("Fajr", config.UseArabic), aladhanTimes.Data.Timings.Fajr},
		{localizedName("Sunrise", config.UseArabic), aladhanTimes.Data.Timings.Sunrise},
		{localizedName("Dhuhr", config.UseArabic), aladhanTimes.Data.Timings.Dhuhr},
		{localizedName("Asr", config.UseArabic), aladhanTimes.Data.Timings.Asr},
		{localizedName("Maghrib", config.UseArabic), aladhanTimes.Data.Timings.Maghrib},
		{localizedName("Isha", config.UseArabic), aladhanTimes.Data.Timings.Isha},
	}
	prevPrayer, nextPrayer := GetCurrentPrayers(prayers)

	fmt.Printf("\n")
	showDate(aladhanTimes, config)
	ShowPrayerLoader(prevPrayer, nextPrayer, config.UseArabic)
	fmt.Printf("\n")
}
