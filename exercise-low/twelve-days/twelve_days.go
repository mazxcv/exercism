package twelve

import (
	"fmt"
	"strings"
)

// On the first day of Christmas my true love gave to me: a Partridge in a Pear Tree.
// On the second day of Christmas my true love gave to me: two Turtle Doves, and a Partridge in a Pear Tree.
// On the third day of Christmas my true love gave to me: three French Hens, two Turtle Doves, and a Partridge in a Pear Tree.
// On the fourth day of Christmas my true love gave to me: four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree.
// On the fifth day of Christmas my true love gave to me: five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree.
// On the sixth day of Christmas my true love gave to me: six Geese-a-Laying, five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree.
// On the seventh day of Christmas my true love gave to me: seven Swans-a-Swimming, six Geese-a-Laying, five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree.
// On the eighth day of Christmas my true love gave to me: eight Maids-a-Milking, seven Swans-a-Swimming, six Geese-a-Laying, five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree.
// On the ninth day of Christmas my true love gave to me: nine Ladies Dancing, eight Maids-a-Milking, seven Swans-a-Swimming, six Geese-a-Laying, five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree.
// On the tenth day of Christmas my true love gave to me: ten Lords-a-Leaping, nine Ladies Dancing, eight Maids-a-Milking, seven Swans-a-Swimming, six Geese-a-Laying, five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree.
// On the eleventh day of Christmas my true love gave to me: eleven Pipers Piping, ten Lords-a-Leaping, nine Ladies Dancing, eight Maids-a-Milking, seven Swans-a-Swimming, six Geese-a-Laying, five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree.
// On the twelfth day of Christmas my true love gave to me: twelve Drummers Drumming, eleven Pipers Piping, ten Lords-a-Leaping, nine Ladies Dancing, eight Maids-a-Milking, seven Swans-a-Swimming, six Geese-a-Laying, five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree.

func Verse(i int) string {
	if i < 1 || i > 12 {
		return ""
	}
	love := map[int][]string{
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
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("On the %s day of Christmas my true love gave to me: ", love[i][0]))
	for i >= 1 {
		sb.WriteString(love[i][1])

		if i == 1 {
			sb.WriteString(".")
		} else if i == 2 {
			sb.WriteString(", and ")
		} else {
			sb.WriteString(", ")
		}
		i--
	}
	return sb.String()
}

func Song() string {
	var sb strings.Builder
	for i := 1; i <= 12; i++ {
		sb.WriteString(Verse(i))
		if i != 12 {
			sb.WriteString("\n")
		}
	}
	return sb.String()
}
