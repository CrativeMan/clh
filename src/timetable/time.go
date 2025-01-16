package timetable

import (
	"fmt"
	"time"
)

const lessonLength time.Duration = 45 * time.Minute

func GetTime() {
	t := time.Now()
	lessons := []time.Time{
		time.Date(t.Year(), t.Month(), t.Day(), 8, 5, 0, 0, t.Location()),
		time.Date(t.Year(), t.Month(), t.Day(), 8, 50, 0, 0, t.Location()),
		time.Date(t.Year(), t.Month(), t.Day(), 9, 55, 0, 0, t.Location()),
		time.Date(t.Year(), t.Month(), t.Day(), 10, 40, 0, 0, t.Location()),
		time.Date(t.Year(), t.Month(), t.Day(), 11, 45, 0, 0, t.Location()),
		time.Date(t.Year(), t.Month(), t.Day(), 12, 30, 0, 0, t.Location()),
		time.Date(t.Year(), t.Month(), t.Day(), 13, 15, 0, 0, t.Location()),
		time.Date(t.Year(), t.Month(), t.Day(), 14, 0, 0, 0, t.Location()),
		time.Date(t.Year(), t.Month(), t.Day(), 14, 45, 0, 0, t.Location()),
	}

	for i, lesson := range lessons {
		if compareTime(t, lesson) {
			printLesson(t, lesson, i+1)
			return
		}
	}

	for _, lesson := range lessons {
		if t.Before(lesson) {
			fmt.Println("Time till next lesson: ", lesson.Sub(t).Round(time.Second))
			return
		}
	}

	fmt.Println("All lessons are over for today.")
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
