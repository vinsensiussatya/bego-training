package util

import (
	"fmt"
	"strconv"
	"time"
)

// GetYearMonth returns the year and month of the given date.
func GetYearMonth(addMonth int) (int, error) {
	theDate := BeginningOfMonth(time.Now()).AddDate(0, addMonth, 0) // use beginning of month to avoid 31th date problem with AddDate
	yearMonthStr := fmt.Sprintf("%d%02d", theDate.Year(), theDate.Month())
	return strconv.Atoi(yearMonthStr)
}

func GetStartOfTheDay(theDay time.Time) time.Time {
	return time.Date(theDay.Year(), theDay.Month(), theDay.Day(), 0, 0, 0, 0, theDay.Location())
}

func GetEndOfTheDay(theDay time.Time) time.Time {
	return time.Date(theDay.Year(), theDay.Month(), theDay.Day(), 23, 59, 59, 999, theDay.Location())
}

func YearMonthToTime(yearMonth int) time.Time {
	year, month := yearMonth/100, yearMonth%100
	return time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
}

func MonthSeries(startDate time.Time, endDate time.Time) (monthSeries []string) {
	const layout = "Jan 06"

	monthSeries = append(monthSeries, startDate.Format(layout))
	for startDate.Before(endDate) {
		startDate = startDate.AddDate(0, 1, 0)
		monthSeries = append(monthSeries, startDate.Format(layout))
	}

	return
}

func DaySeries(startDate time.Time, endDate time.Time, rangeDay int) (series []string) {
	const layout = "Jan 02"

	for startDate.Before(endDate) {
		if rangeDay > 1 {
			series = append(series, fmt.Sprintf("%s - %s", startDate.Format(layout), startDate.AddDate(0, 0, rangeDay-1).Format(layout)))
		} else {
			series = append(series, startDate.Format(layout))
		}
		startDate = startDate.AddDate(0, 0, rangeDay)
	}

	return
}

func GetMondayOfTheWeek(theDate time.Time) time.Time {
	return theDate.AddDate(0, 0, -int(theDate.Weekday())+1)
}

// GetDiffMonthSeries return diff months from theYearMonth from startPeriod
func GetDiffMonthSeries(startPeriod time.Time, theYearMonth int) int {
	diffMonth := 0
	theDate := YearMonthToTime(theYearMonth)
	diffYear := theDate.Year() - startPeriod.Year()
	if diffYear > 0 {
		diffMonth = diffYear * 12
	}

	diffMonth += int(theDate.Month()) - int(startPeriod.Month())
	return diffMonth
}

// GetDiffDay return diff days from theDate from startPeriod, per dayGap
func GetDiffDay(startPeriod time.Time, theDate time.Time, dayGap int) int {
	dayDiff := theDate.Sub(startPeriod).Hours() / 24
	return int(dayDiff) / dayGap
}

func GetNumberMonth(month string) string {
	numberMonth := ""
	switch month {
	case "Jan":
		numberMonth = "01"
	case "Feb":
		numberMonth = "02"
	case "Mar":
		numberMonth = "03"
	case "Apr":
		numberMonth = "04"
	case "May":
		numberMonth = "05"
	case "Jun":
		numberMonth = "06"
	case "Jul":
		numberMonth = "07"
	case "Aug":
		numberMonth = "08"
	case "Sep":
		numberMonth = "09"
	case "Oct":
		numberMonth = "10"
	case "Nov":
		numberMonth = "11"
	case "Dec":
		numberMonth = "12"
	}

	return numberMonth
}

func GetMax(a []float64) (max float64) {
	max = a[0]
	for _, value := range a {
		if value > max {
			max = value
		}
	}

	return max
}

func BeginningOfMonth(timeDate time.Time) time.Time {
	return time.Date(timeDate.Year(), timeDate.Month(), 1, 0, 0, 0, 0, timeDate.Location())
}

func EndOfMonth(timeDate time.Time) time.Time {
	return BeginningOfMonth(timeDate).AddDate(0, 1, -1)
}

func GetNumberFullNameMonth(month string) (numMonth int) {
	switch month {
	case "January":
		numMonth = 1
	case "February":
		numMonth = 2
	case "March":
		numMonth = 3
	case "April":
		numMonth = 4
	case "May":
		numMonth = 5
	case "June":
		numMonth = 6
	case "July":
		numMonth = 7
	case "August":
		numMonth = 8
	case "September":
		numMonth = 9
	case "October":
		numMonth = 10
	case "November":
		numMonth = 11
	case "December":
		numMonth = 12
	default:
		num := time.Now().Month()
		return int(num)
	}
	return numMonth
}

// ConvertTimezone convert location for provided time.
//
// Example:
// 2022-01-01T00:00:00Z, UTC+7 --> 2021-12-31T17:00:00Z
//
// It's different from time.Time.In():
//
// 2022-01-01T00:00:00Z, UTC+7 -> 2022-01-01T07:00:00+07:00
func ConvertTimezone(t time.Time, location *time.Location) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), location).In(t.Location())
}

func ParseLocation(hourOffset int) *time.Location {
	zone := "UTC"
	if hourOffset != 0 {
		zone = fmt.Sprintf("UTC%+d", hourOffset)
	}
	return time.FixedZone(zone, hourOffset*60*60)
}
