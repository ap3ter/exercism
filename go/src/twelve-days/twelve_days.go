package twelve

import (
	"fmt"
)

type dayVerse struct {
	day  string
	gift string
}

// Hold the day and the gift for the day
var verses = map[int]dayVerse{
	1:  {"first", "a Partridge in a Pear Tree"},
	2:  {"second", "two Turtle Doves"},
	3:  {"third", "three French Hens"},
	4:  {"fourth", "four Calling Birds"},
	5:  {"fifth", "five Gold Rings"},
	6:  {"sixth", "six Geese-a-Laying"},
	7:  {"seventh", "seven Swans-a-Swimming"},
	8:  {"eighth", "eight Maids-a-Milking"},
	9:  {"ninth", "nine Ladies Dancing"},
	10: {"tenth", "ten Lords-a-Leaping"},
	11: {"eleventh", "eleven Pipers Piping"},
	12: {"twelfth", "twelve Drummers Drumming"},
}

// Verse creates the verse for the day
func Verse(d int) string {
	v, ok := verses[d]
	if !ok {
		panic(fmt.Errorf("Wrong day %v should be between 1 and 12", d))
	}
	// Make gifts for all days
	gifts := ""
	for i := d; i >= 1; i-- {
		and := ""
		if d > 1 && i == 1 {
			and = "and "
		}
		gifts += ", " + and + verses[i].gift
	}
	// Make the verse by combining the gifts and rest of words
	verse := fmt.Sprintf("On the %v day of Christmas my true love gave to me%v.", v.day, gifts)
	return verse
}

// Song generates the twelve day song
func Song() string {
	var song string
	for i := 1; i <= 12; i++ {
		song += Verse(i) + "\n"
	}
	return song
}
