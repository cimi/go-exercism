package twelve

import (
	"fmt"
	"strings"
)

const verse = "On the %s day of Christmas my true love gave to me, %s."

var gifts = [12]string{
	"twelve Drummers Drumming",
	"eleven Pipers Piping",
	"ten Lords-a-Leaping",
	"nine Ladies Dancing",
	"eight Maids-a-Milking",
	"seven Swans-a-Swimming",
	"six Geese-a-Laying",
	"five Gold Rings",
	"four Calling Birds",
	"three French Hens",
	"two Turtle Doves",
	"and a Partridge in a Pear Tree",
}

var days = [12]string{
	"first",
	"second",
	"third",
	"fourth",
	"fifth",
	"sixth",
	"seventh",
	"eighth",
	"ninth",
	"tenth",
	"eleventh",
	"twelfth",
}

func Gifts(day int) string {
	if day == 0 {
		return "a Partridge in a Pear Tree"
	}
	return strings.Join(gifts[(len(gifts)-day-1):], ", ")
}

func Verse(day int) string {
	day = day - 1
	return fmt.Sprintf(verse, days[day], Gifts(day))
}

func Song() string {
	var sb strings.Builder
	for i := 1; i <= len(gifts); i++ {
		sb.WriteString(Verse(i))
		sb.WriteString("\n")
	}
	return sb.String()
}
