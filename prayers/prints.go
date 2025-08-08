package prayers

import (
	"fmt"
	"runtime"
	"strings"
	"unicode/utf8"

	"github.com/01walid/goarabic"
	"github.com/fatih/color"
)

// Color codes list : https://www.ditig.com/256-colors-cheat-sheet

var userOS string

func init() {
	userOS = runtime.GOOS
}

func BasmalahFormatter(basmalah string) string {
	switch userOS {

	case "darwin":
		return fmt.Sprintf("%-7s%s%-3s\n", "", basmalah, "")
	case "linux":
		return fmt.Sprintf("%-6s%s%-3s\n", "", basmalah, "")
	case "windows":
		basmalah = goarabic.Reverse("بِسْمِ ٱللّٰهِ ٱلرَّحْمَٰنِ ٱلرَّحِيم")
		return fmt.Sprintf("%-2s%s\n", "", basmalah)
	default:
		return ""
	}
}

func BasmalahPrinter(basmalahSprint string) {
	// Choose colors here
	switch userOS {

	case "darwin":
		fmt.Printf(basmalahSprint)
	case "linux":
		whitePrint := color.New(color.FgHiWhite, color.Bold).PrintFunc()
		whitePrint(basmalahSprint)
	case "windows":
		whitePrint := color.New(color.FgHiWhite, color.Bold).PrintFunc()
		whitePrint(basmalahSprint)
		return

	default:
		return
	}
}

func DateFormatter(todayDate string) string {
	switch userOS {

	case "darwin":
		return fmt.Sprintf("%-1s%s\n", "", todayDate)
	case "linux":
		return fmt.Sprintf("%-1s%s\n", "", todayDate)
	case "windows":
		return fmt.Sprintf("%-1s%s\n", "", todayDate)

	default:
		return ""
	}
}

func DatePrinter(dateSprint string) {
	// Choose colors here
	switch userOS {

	case "darwin":
		fmt.Printf(boldColor256(157, dateSprint))
	case "linux":
		fmt.Printf(boldColor256(157, dateSprint))
	case "windows":
		fmt.Printf(boldColor256(157, dateSprint))
	default:
		return
	}
}

func LoaderFormatter(nextPrayerName string, loadingSquares string, timeRemaining string, useArabic bool) string {
	if useArabic {
		switch userOS {

		case "darwin":
			return fmt.Sprintf("%s %s %s %s %s\n", "", timeRemaining, loadingSquares, "القادم:", nextPrayerName)
		case "linux":
			return fmt.Sprintf("%s %s %s %s %s\n", "", timeRemaining, loadingSquares, "القادم:", nextPrayerName)
		case "windows":
			arNext := goarabic.Reverse("القادم:")
			return fmt.Sprintf("%s %s %s %s %s\n", "", timeRemaining, loadingSquares, nextPrayerName, arNext)
		default:
			return ""
		}
	} else {

		switch userOS {

		case "darwin":
			return fmt.Sprintf("%-1sNext: %s %s %s\n", "", nextPrayerName, loadingSquares, timeRemaining)
		case "linux":
			return fmt.Sprintf("%-1sNext: %s %s %s\n", "", nextPrayerName, loadingSquares, timeRemaining)
		case "windows":
			return fmt.Sprintf("%-1sNext: %s %s %s\n", "", nextPrayerName, loadingSquares, timeRemaining)
		default:
			return ""
		}
	}
}

func LoaderPrinter(loaderSprint string) {
	// Choose colors here
	switch userOS {

	case "darwin":
		fmt.Printf(boldColor256(153, loaderSprint))
	case "linux":
		fmt.Printf(boldColor256(153, loaderSprint))
	case "windows":
		fmt.Printf(boldColor256(153, loaderSprint))
	default:
		return
	}
}

func PrayerFormatter(prayerName string, prayerTime string, useArabic bool) string {
	// When Arabic is enabled, show time on the left and name on the right
	if useArabic {
		const columnWidth = 7
		paddedName := padRightAlign(prayerName, columnWidth)
		return fmt.Sprintf("%s : %s", prayerTime, paddedName)
	} else {
		// Default (non-Arabic): name on the left, time on the right
		switch userOS {
		case "darwin":
			return fmt.Sprintf("%-7s : %s", prayerName, prayerTime)
		case "linux":
			return fmt.Sprintf("%-7s : %s", prayerName, prayerTime)
		case "windows":
			return fmt.Sprintf("%-7s : %s", prayerName, prayerTime)
		default:
			return ""
		}
	}

}

func padRightAlign(s string, width int) string {
	// Count visible runes, not bytes
	runeCount := utf8.RuneCountInString(s)
	if runeCount >= width {
		return s
	}
	padding := width - runeCount
	return strings.Repeat(" ", padding) + s
}

func PrayerColorer(prayerSprint string, isPrev bool, isNext bool) string {
	// Change Colors here
	switch userOS {
	case "darwin":
		if isPrev {
			return fmt.Sprintf(boldColor256(153, prayerSprint))
		} else if isNext {
			return fmt.Sprintf(boldColor256(102, prayerSprint))
		} else {
			return fmt.Sprintf(boldColor256(102, prayerSprint))
		}
	case "windows":
		if isPrev {
			return fmt.Sprintf(boldColor256(153, prayerSprint))
		} else if isNext {
			return fmt.Sprintf(boldColor256(102, prayerSprint))
		} else {
			return fmt.Sprintf(boldColor256(102, prayerSprint))
		}
	case "linux":
		if isPrev {
			return fmt.Sprintf(boldColor256(153, prayerSprint))
		} else if isNext {
			//return fmt.Sprintf(boldColor256(15, prayerSprint))
			return prayerSprint
		} else {
			return prayerSprint
		}

	default:
		return ""
	}
}

func TableFormatter(coloredPrayer string) string {
	switch userOS {

	case "darwin":
		return fmt.Sprintf("   │ %s │\n", coloredPrayer)
	case "linux":
		return fmt.Sprintf("   │ %s │\n", coloredPrayer)
	case "windows":
		return fmt.Sprintf("   │ %s │\n", coloredPrayer)
	default:
		return ""
	}
}
