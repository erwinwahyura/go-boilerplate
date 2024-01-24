package utils

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"time"
)

var (
	LocJakarta *time.Location

	TIME_FORMAT_ONLY_NUMBER  = "20060102150405"
	TIME_FORMAT_STANDARD     = "2006-01-02 15:04:05"
	TIME_FORMAT_STANDARD_ISO = "2006-01-02T15:04:05Z"
	TIME_IN_MINUTE           = "minute"
	TIME_IN_HOUR             = "hour"
)

func init() {
	LocJakarta = MustLoadLocation("Asia/Jakarta")
}

// PUBLIC METHOD
func MustLoadLocation(name string) *time.Location {
	l, err := time.LoadLocation(name)
	if err != nil {
		panic(fmt.Sprintf("time util: could not load location %s: %s", name, err.Error()))
	}
	return l
}

func TimeNow() time.Time {
	return time.Now().In(LocJakarta)
}

// TimeLapse struct
type TimeLapse struct {
	StartTime time.Time
}

// StartTimer func
func StartTimer() TimeLapse {
	return TimeLapse{
		StartTime: time.Now(),
	}
}

// GetTimeLapseSecond function to get timelapse in second
func (t TimeLapse) GetTimeLapseSecond() float64 {
	return time.Since(t.StartTime).Seconds()
}

// GetTimeNowIDN function to get time now UTC + 7
func GetTimeNowIDN() time.Time {
	nowUTC := time.Now().UTC()
	nowIDN := nowUTC.Add(time.Hour * 7)
	return nowIDN
}

// ConvertStringToTime convert string to time
func ConvertStringToTime(datetime string, format string) (time.Time, error) {
	return time.Parse(format, datetime)
}

// ConvertToTargetTimezone convert to target timezone
func ConvertToTargetTimezone(datetime string, currTimezone int32, targetTimezone int32) (time.Time, error) {
	timezone := "UTC"
	if currTimezone < 0 {
		timezone += "-%d"
	} else {
		timezone += "+%d"
	}
	tz := fmt.Sprintf(timezone, currTimezone)
	loc := time.FixedZone(tz, int(currTimezone)*60*60)
	arrTime, err := time.ParseInLocation(TIME_FORMAT_STANDARD, datetime, loc)
	if err != nil {
		return time.Now(), err
	}

	timezone = "UTC"
	if targetTimezone < 0 {
		timezone += "-%d"
	} else {
		timezone += "+%d"
	}
	tz = fmt.Sprintf(timezone, targetTimezone)
	return arrTime.In(time.FixedZone(tz, int(targetTimezone)*60*60)), nil
}

// ConvertTimeToMillisecond time to millisecond
func ConvertTimeToMillisecond(times time.Time) int64 {
	return times.UnixNano() / int64(time.Millisecond)
}

// GetCurrentTimeGMT7 func to get GMT+7 timezone
func GetCurrentTimeGMT7() (time.Time, error) {
	t := time.Now()
	// get the location
	location, err := time.LoadLocation("Asia/Jakarta")
	if err == nil {
		t = t.In(location)
	}
	// this should give you time in location
	return t, err
}

// TimeToFormat convert time to format
func TimeToFormat(datetime time.Time, targetFormat string) (time.Time, error) {
	formatString := datetime.Format(targetFormat)
	return time.Parse(targetFormat, formatString)
}

type (
	Timer struct {
		InitRunTime time.Time
		LastRunTime time.Time
	}
)

var (
	reStr2Seconds        = regexp.MustCompile(`^(?i)([-+])?(\d+)([a-z])$`)
	ErrDateInvalidFormat = fmt.Errorf("invalid date format")
	ErrDateInvalidRange  = fmt.Errorf("invalid date range, start date greater than end date")
)

func init() {
	LocJakarta = MustLoadLocation("Asia/Jakarta")
}

func NewTimer() *Timer {
	return &Timer{}
}

func (t *Timer) GetElapsedTime() float64 {
	if iz := t.LastRunTime.IsZero(); iz {
		panic("Timer has not been started yet")
	}

	lastRun := t.LastRunTime

	elTime := time.Since(lastRun).Seconds()
	t.LastRunTime = time.Now()

	return elTime
}

func (t *Timer) GetTotalRunTime() float64 {
	if iz := t.InitRunTime.IsZero(); iz {
		panic("Timer has not been started yet")
	}

	initRun := t.InitRunTime

	ttlTime := time.Since(initRun).Seconds()

	return ttlTime
}

func (t *Timer) SetNow() {
	now := time.Now()

	if iz := t.InitRunTime.IsZero(); iz {
		t.InitRunTime = now
	}
	t.LastRunTime = now
}

func (t *Timer) StartTimer() {
	now := time.Now()

	t.InitRunTime = now
	t.LastRunTime = now
}

func GetUnixUTC(t time.Time) (int64, error) {
	// if time not defined then return
	if iz := t.IsZero(); iz {
		return 0, errors.New("invalid time format, variable not initialized")
	}
	// if time zone is already utc (which usually come from database without timezone) then substract by 7
	if zone, _ := t.Zone(); zone == "UTC" {
		t = t.Add(time.Duration(-7) * time.Hour)
	}

	res := t.UTC().Unix()
	return res, nil
}

func Str2Seconds(s string) int {
	subMatchList := reStr2Seconds.FindStringSubmatch(s)
	if len(subMatchList) > 0 {
		sign := subMatchList[1]
		num, err := strconv.Atoi(subMatchList[2])
		dtType := subMatchList[3]

		if err == nil {
			var secRatio int
			switch dtType {
			case "s":
				secRatio = 1
			case "m":
				secRatio = 60
			case "h":
				secRatio = 3600
			case "d":
				secRatio = 86400
			case "w":
				secRatio = 604800
			case "M":
				secRatio = 2592000
			case "Y":
				secRatio = 31104000
			}

			if secRatio > 0 {
				duration := num * secRatio
				if sign == "-" {
					duration *= -1
				}

				return duration
			}
		}
	}

	return 0
}

func Str2Date(s string) string {

	year := s[:4]
	month := s[4:6]
	day := s[6:8]
	hour := s[8:10]
	minute := s[10:12]

	return day + "-" + month + "-" + year + " Pukul " + hour + ":" + minute + " WIB"
}

func UnixNano2Unix(utNano int64) int64 {
	return utNano / int64(time.Second)
}

func UnixNano2UnixTime(utNano int64) time.Time {
	return time.Unix(UnixNano2Unix(utNano), 0)
}

// get diff time in second
func GetDifferentTime(current time.Time) int64 {
	unixTime := current.Unix()
	diffTime := unixTime - time.Now().Unix()

	return diffTime
}

func GetDifferentTimeInMillis(current time.Time) int64 {
	unixTime := current.UnixNano()
	diffTime := unixTime - time.Now().UnixNano()
	diffTime = diffTime / 1e6

	return diffTime
}

// check if now is after "time"
// format: yyyy-mm-dd hh:ii:ss
func IsNowAfterTime(timeString string) bool {
	if tm, err := time.Parse(TIME_FORMAT_STANDARD, timeString); err == nil {
		if SetTimezoneToLocal(time.Now()).After(tm.Add(time.Duration(-7) * time.Hour)) {
			return true
		}
	}
	return false
}

// LimitDateRange limit a date range in the format of YYYY-MM-DD.
// By default, the start date is ensured to be within 1 month range from the end date.
// If maxRangeInMonths is specified, then start date will be ensured to within maxRangeInMonths months from end date.
func LimitDateRange(startDateStr, endDateStr string, maxRangeInMonths int) (startDateClean, endDateClean string, e error) {

	var startDate time.Time
	var endDate time.Time
	var err error

	// startDate and endDate must be in YYYY-MM-DD
	// If either of one them is not, return error invalid format
	endDate, err = time.Parse("2006-01-02", endDateStr)
	if err != nil {
		return startDateClean, endDateClean, ErrDateInvalidFormat
	}
	startDate, e = time.Parse("2006-01-02", startDateStr)
	if e != nil {
		return startDateClean, endDateClean, ErrDateInvalidFormat
	}

	lastNMonths := endDate.AddDate(0, -1, 0)
	if maxRangeInMonths > 0 {
		lastNMonths = endDate.AddDate(0, -1*maxRangeInMonths, 0)
	}

	if endDate.Before(startDate) {
		return startDateClean, endDateClean, ErrDateInvalidRange
	}

	if startDate.Before(lastNMonths) {
		startDate = lastNMonths
	}

	startDateClean = startDate.Format("2006-01-02")
	endDateClean = endDate.Format("2006-01-02")

	return startDateClean, endDateClean, nil
}

func DateTimeFormat(t time.Time, useShortMonth bool, useTz bool) string {
	month := t.Month()
	var monthString string

	MonthMei := "Mei"

	if useShortMonth {
		switch month {
		case 1:
			monthString = "Jan"
		case 2:
			monthString = "Feb"
		case 3:
			monthString = "Mar"
		case 4:
			monthString = "Apr"
		case 5:
			monthString = MonthMei
		case 6:
			monthString = "Jun"
		case 7:
			monthString = "Jul"
		case 8:
			monthString = "Ags"
		case 9:
			monthString = "Sep"
		case 10:
			monthString = "Okt"
		case 11:
			monthString = "Nov"
		case 12:
			monthString = "Des"
		}
	} else {
		switch month {
		case 1:
			monthString = "Januari"
		case 2:
			monthString = "Februari"
		case 3:
			monthString = "Maret"
		case 4:
			monthString = "April"
		case 5:
			monthString = MonthMei
		case 6:
			monthString = "Juni"
		case 7:
			monthString = "Juli"
		case 8:
			monthString = "Agustus"
		case 9:
			monthString = "September"
		case 10:
			monthString = "Oktober"
		case 11:
			monthString = "November"
		case 12:
			monthString = "Desember"
		}
	}

	tzString := "WIB"
	var layout string
	if useTz {
		layout = fmt.Sprintf("02 %s 2006, 15:04 %s", monthString, tzString)
	} else {
		layout = fmt.Sprintf("02 %s 2006 pukul 15:04", monthString)
	}

	return t.Format(layout)
}

func IndonesianMonth(month int) string {
	months := make(map[int]string)
	months[1] = "Januari"
	months[2] = "Februari"
	months[3] = "Maret"
	months[4] = "April"
	months[5] = "Mei"
	months[6] = "Juni"
	months[7] = "Juli"
	months[8] = "Agustus"
	months[9] = "September"
	months[10] = "Oktober"
	months[11] = "November"
	months[12] = "Desember"

	return months[month]
}

func GetLocalWeekday(d time.Weekday) string {
	switch d {
	case 0:
		return "Minggu"
	case 1:
		return "Senin"
	case 2:
		return "Selasa"
	case 3:
		return "Rabu"
	case 4:
		return "Kamis"
	case 5:
		return "Jumat"
	case 6:
		return "Sabtu"
	}
	return "UNDEFINED"
}

func GetIndonesianDate(t time.Time) string {
	d := GetLocalWeekday(t.Weekday())
	m := IndonesianMonth(int(t.Month()))
	// m := GetLocalMonth(t.Month())

	clockStr := t.Format("15:04")
	localDate := fmt.Sprintf("%v %v %v %v pukul %v", d, t.Day(), m, t.Year(), clockStr)
	return localDate
}

// get unix timestamp for next scheduler
// t: current time
// scheduleHour: time the scheduler to be run
func GetScheduleUnix(t time.Time, scheduleHour int) int64 {
	var scheduleUnixTime int64
	hr, _, _ := t.Clock()
	hourToSchedule := (24 + scheduleHour - hr) % 24 // duration to schedule in hours
	if hourToSchedule == 0 {
		hourToSchedule = 24
	}
	scheduleUnixTime = time.Date(t.Year(), t.Month(), t.Day(), hr+hourToSchedule, 0, 0, 0, t.Location()).Unix()

	return scheduleUnixTime
}

func GetLastNDate(n int) []time.Time {
	days := make([]time.Time, 0)
	curr := time.Now()
	curr = time.Date(curr.Year(), curr.Month(), curr.Day(), 0, 0, 0, 0, time.UTC)
	for i := 0; i < n; i++ {
		days = append([]time.Time{curr}, days...)
		curr = curr.AddDate(0, 0, -1)
	}

	return days
}

func GenerateRangeDateQuery(dates []time.Time) string {
	var startDate, endDate time.Time
	var dateClause string

	totalDate := len(dates)
	if totalDate == 0 {
		return ""
	}

	startDate = dates[0]

	if totalDate == 1 {
		date := startDate.Format("20060102")
		dateClause = fmt.Sprintf("and date_id = %s", date)
	} else {
		endDate = dates[totalDate-1]

		startDateStr := startDate.Format("20060102")
		endDateStr := endDate.Format("20060102")

		dateClause = fmt.Sprintf("and (date_id >= %s and date_id <= %s)", startDateStr, endDateStr)
	}

	return dateClause
}

func GetIntersectedDate(startDateStr, endDateStr, layout string, n int) ([]time.Time, error) {
	intersectedDays := make([]time.Time, 0)
	lastNDays := GetLastNDate(n)

	endDate, err := time.Parse(layout, endDateStr)
	if err != nil {
		return intersectedDays, err
	}

	startDate, err := time.Parse(layout, startDateStr)
	if err != nil {
		return intersectedDays, err
	}

	if endDate.Before(lastNDays[0]) {
		return intersectedDays, nil
	}

	for _, day := range lastNDays {
		if day.After(endDate) {
			break
		}
		if day.Before(startDate) {
			continue
		}
		intersectedDays = append(intersectedDays, day)
	}

	return intersectedDays, nil
}

func SetTimeToEndOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
}

func SetTimezoneToLocal(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), LocJakarta)
}

// GetNextDayStartTime gives the next day's start time based on the current time input
func GetNextDayStartTime(t time.Time) time.Time {
	timeStr := t.Format("2006-01-02")
	currTime, _ := time.Parse("2006-01-02", timeStr)
	nextDayStartTime := currTime.Add(time.Duration(24) * time.Hour)

	return nextDayStartTime
}

// GetDayStartTime returns start time of a day given time as input
func GetDayStartTime(t time.Time) time.Time {
	newTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
	return SetTimezoneToLocal(newTime)
}

func InTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}

const DateStatsFormat string = "20060102"

// ParseDateStatistic parse date in integer form to time form
func ParseDateStatistic(i int64) time.Time {
	dateStr := strconv.FormatInt(i, 10)
	t, _ := time.Parse(DateStatsFormat, dateStr)
	return t
}

// FormatDateStatistic format date into integer form
func FormatDateStatistic(t time.Time) int64 {
	i, _ := strconv.ParseInt(t.Format(DateStatsFormat), 10, 64)
	return i
}

// GenerateExpireTime in hour / minutes
func GenerateExpireTime(timeDuration time.Duration, timeInWord string, timeFormat string) string {
	if timeInWord == TIME_IN_MINUTE {
		return TimeNow().Add(time.Minute * timeDuration).Format(timeFormat)
	}
	// default timeInWord is in hours
	return TimeNow().Add(time.Hour * timeDuration).Format(timeFormat)
}
