/*
*

	@author:
	@date : 2023/10/25
*/
package myTime

import (
	"strconv"
	"time"
)

// 时间字符串转时间戳 timestr格式 yyyy-mm-dd hh:ii:ss timestamp int
func TimestrToTimestamp(timestr string) int {
	t, _ := time.ParseInLocation(time.UnixDate, timestr, time.Local)
	return int(t.Unix())
}

// 时间戳转时间字符串
func TimestampToTimestr(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format(time.UnixDate)
}

// n小时前后 1h：1小时后  -1h：1小时前（当前时间为基准）
func NHourLastAfterTimestamp(str string) int {
	nowT := time.Now()
	hh, _ := time.ParseDuration(str)
	t := nowT.Add(hh)
	return int(t.Unix())
}

// 根据秒数获取相差 天时分秒
func SecondsDifferDayHoursMinutesSeconds(seconds int) (day, hour, minute, second int) {
	day = seconds / 60 / 60 / 24 % 365
	hour = seconds / 60 / 60 % 24
	minute = seconds / 60 % 60
	second = seconds % 60
	return
}

// 根据时间字符串获取年月日int
func TimestrGetYearMonthDayInt(timestr string) (year, month, day int) {
	t, _ := time.ParseInLocation(time.UnixDate, timestr, time.Local)
	year = t.Year()
	month, _ = strconv.Atoi(t.Format("01"))
	day = t.Day()
	return
}

// 根据时间戳获取年月日int
func TimestampGetYearMonthDayInt(timestamp int) (year, month, day int) {
	t := time.Unix(int64(timestamp), 0)
	year = t.Year()
	month, _ = strconv.Atoi(t.Format("01"))
	day = t.Day()
	return
}

// 今天开始时间与结束时间戳
func CurDayStartEndTimestamp() (StartTimeStamp, EndTimeStamp int) {
	nowT := time.Now()
	CurYear, CurMonth, CurDay := nowT.Date()
	StartTimeStamp = int(time.Date(CurYear, CurMonth, CurDay, 0, 0, 0, 0, time.Local).Unix())
	EndTimeStamp = int(time.Date(CurYear, CurMonth, CurDay, 23, 59, 59, 0, time.Local).Unix())
	return
}

// 今天开始时间与结束时间戳(毫秒)
func CurDayStartEndTimestampMilli() (StartTimeStamp, EndTimeStamp int) {
	nowT := time.Now()
	CurYear, CurMonth, CurDay := nowT.Date()
	StartTimeStamp = int(time.Date(CurYear, CurMonth, CurDay, 0, 0, 0, 0, time.Local).UnixMilli())
	EndTimeStamp = int(time.Date(CurYear, CurMonth, CurDay, 23, 59, 59, 0, time.Local).UnixMilli())
	return
}

// 昨天开始与结束时间戳
func YesterdayStartEndTimestamp() (StartTimeStamp, EndTimeStamp int) {
	nowT := time.Now()
	yesterT := nowT.AddDate(0, 0, -1)
	StartTimeStamp = int(time.Date(yesterT.Year(), yesterT.Month(), yesterT.Day(), 0, 0, 0, 0, time.Local).Unix())
	EndTimeStamp = int(time.Date(yesterT.Year(), yesterT.Month(), yesterT.Day(), 23, 59, 59, 0, time.Local).Unix())
	return
}

// 某天时间字符串获取其当天开始与结束时间戳 timestr格式 yyyy-mm-dd hh:ii:ss
func DayTimestrToStartEndTimestamp(timestr string) (StartTimeStamp, EndTimeStamp int) {
	t, _ := time.ParseInLocation(time.UnixDate, timestr, time.Local)
	CurYear, CurMonth, CurDay := t.Date()
	StartTimeStamp = int(time.Date(CurYear, CurMonth, CurDay, 0, 0, 0, 0, time.Local).Unix())
	EndTimeStamp = int(time.Date(CurYear, CurMonth, CurDay, 23, 59, 59, 0, time.Local).Unix())
	return
}

// 某天时间字戳获取其当天开始与结束时间戳
func DayTimestampToStartEndTimestamp(timestamp int) (StartTimeStamp, EndTimeStamp int) {
	t := time.Unix(int64(timestamp), 0)
	CurYear, CurMonth, CurDay := t.Date()
	StartTimeStamp = int(time.Date(CurYear, CurMonth, CurDay, 0, 0, 0, 0, time.Local).Unix())
	EndTimeStamp = int(time.Date(CurYear, CurMonth, CurDay, 23, 59, 59, 0, time.Local).Unix())
	return
}

// 这个月开始与结束时间戳
func CurMonthStartEndTimestamp() (StartTimeStamp, EndTimeStamp int) {
	t := time.Now()
	CurYear, CurMonth, _ := t.Date()
	StartTimeStamp = int(time.Date(CurYear, CurMonth, 1, 0, 0, 0, 0, time.Local).Unix())
	EndTimeStamp = int(time.Date(CurYear, CurMonth+1, 1, 0, 0, 0, -1, time.Local).Unix())
	return
}

// 上个月的开始与结束时间戳
func LastMonthStartEndTimestamp() (StartTimeStamp, EndTimeStamp int) {
	t := time.Now()
	CurYear, CurMonth, _ := t.Date()
	StartTimeStamp = int(time.Date(CurYear, CurMonth-1, 1, 0, 0, 0, 0, time.Local).Unix())
	EndTimeStamp = int(time.Date(CurYear, CurMonth, 1, 0, 0, 0, -1, time.Local).Unix())
	return
}

// 今年某个月的开始与结束时间戳
func SomeMonthStartEndTimestamp(m time.Month) (StartTimeStamp, EndTimeStamp int) {
	t := time.Now()
	CurYear, _, _ := t.Date()
	StartTimeStamp = int(time.Date(CurYear, m, 1, 0, 0, 0, 0, time.Local).Unix())
	EndTimeStamp = int(time.Date(CurYear, m+1, 1, 0, 0, 0, -1, time.Local).Unix())
	return
}

// 当前时间转字符串 // yyyymmddhhmmss
func NowTimeToString() string {
	return time.Now().Format("20060102150405")
}

// 返回 yyyymmdd
func GetUnixTimeDateStr(timestamp int64) string {
	return time.Unix(timestamp, 0).Format("20060102")
}
