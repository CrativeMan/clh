package timetable

import (
	"fmt"
	"time"
)

const lessonLength time.Duration = 45 * time.Minute

func GetTime() {

	t := time.Now()
	beginFirstLesson := time.Date(t.Year(), t.Month(), t.Day(), 8, 5, 0, 0, t.Location())
	beginSecondLesson := time.Date(t.Year(), t.Month(), t.Day(), 8, 50, 0, 0, t.Location())
	beginThirLesson := time.Date(t.Year(), t.Month(), t.Day(), 9, 55, 0, 0, t.Location())
	beginFourthLesson := time.Date(t.Year(), t.Month(), t.Day(), 10, 40, 0, 0, t.Location())
	beginFifthLesson := time.Date(t.Year(), t.Month(), t.Day(), 11, 45, 0, 0, t.Location())
	beginSixthLesson := time.Date(t.Year(), t.Month(), t.Day(), 12, 30, 0, 0, t.Location())
	beginSeventhLesson := time.Date(t.Year(), t.Month(), t.Day(), 13, 15, 0, 0, t.Location())
	beginEightthLesson := time.Date(t.Year(), t.Month(), t.Day(), 14, 0, 0, 0, t.Location())
	beginNinethLesson := time.Date(t.Year(), t.Month(), t.Day(), 14, 45, 0, 0, t.Location())

	if compareTime(t, beginFirstLesson) {
		printLesson(t, beginFirstLesson, 1)
	} else if compareTime(t, beginSecondLesson) {
		printLesson(t, beginSecondLesson, 2)
	} else if compareTime(t, beginThirLesson) {
		printLesson(t, beginThirLesson, 3)
	} else if compareTime(t, beginFourthLesson) {
		printLesson(t, beginFourthLesson, 4)
	} else if compareTime(t, beginFifthLesson) {
		printLesson(t, beginFifthLesson, 5)
	} else if compareTime(t, beginSixthLesson) {
		printLesson(t, beginSixthLesson, 6)
	} else if compareTime(t, beginSeventhLesson) {
		printLesson(t, beginSeventhLesson, 7)
	} else if compareTime(t, beginEightthLesson) {
		printLesson(t, beginEightthLesson, 8)
	} else if compareTime(t, beginNinethLesson) {
		printLesson(t, beginNinethLesson, 9)
	} else {
		println("No lesson")
	}

}

func printLesson(current time.Time, begin time.Time, lesson int) {
	fmt.Println("============ Current Lesson ============")
	fmt.Println("  Current lesson: ", lesson)
	fmt.Println("  Current time: ", current.Format("15:04:05"))
	fmt.Println("  Begin of lesson: ", begin.Format("15:04:05"))
	fmt.Println("  End of lesson: ", begin.Add(lessonLength).Format("15:04:05"))
	fmt.Println("  Time left in lesson: ", begin.Add(lessonLength).Sub(current).Round(time.Second))
}

func compareTime(t time.Time, t2 time.Time) bool {
	return t.After(t2) && t.Before(t2.Add(lessonLength))
}
