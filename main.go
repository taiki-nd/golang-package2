package main

import (
	"fmt"
	"regexp"
	"sort"
	"time"
)

const (
	/*
		c1 = iota
		c2 = iota
		c3 = iota
	*/

	c1 = iota
	c2 //省略可能
	c3
)

const (
	_      = iota
	KB int = 1 << (10 * iota)
	MB
	GB
)

func main() {
	//time
	/*
	   const (
	   	Layout      = "01/02 03:04:05PM '06 -0700" // The reference time, in numerical order.
	   	ANSIC       = "Mon Jan _2 15:04:05 2006"
	   	UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
	   	RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
	   	RFC822      = "02 Jan 06 15:04 MST"
	   	RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
	   	RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
	   	RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
	   	RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
	   	RFC3339     = "2006-01-02T15:04:05Z07:00"
	   	RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	   	Kitchen     = "3:04PM"
	   	// Handy time stamps.
	   	Stamp      = "Jan _2 15:04:05"
	   	StampMilli = "Jan _2 15:04:05.000"
	   	StampMicro = "Jan _2 15:04:05.000000"
	   	StampNano  = "Jan _2 15:04:05.000000000"
	   )
	*/
	fmt.Println("---time---")
	t := time.Now()
	fmt.Println(t)
	fmt.Println(t.Format(time.RFC3339))
	fmt.Println(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())

	//regex
	fmt.Println("---regex---")
	match, _ := regexp.MatchString("a([a-z]+)e", "apple")
	fmt.Println(match)
	match2, _ := regexp.MatchString("a([a-z]+)e", "app0le")
	fmt.Println(match2)

	r := regexp.MustCompile("a([a-z]+)e")
	ms := r.MatchString("apple")
	fmt.Println(ms)

	r2 := regexp.MustCompile("^/(view|edit|save)/([a-zA-Z0-9]+)$")
	fs := r2.FindString("/view/test")
	fmt.Println(fs)
	fss := r2.FindStringSubmatch("/view/test")
	fmt.Println(fss, fss[0], fss[1], fss[2])

	//sort
	fmt.Println("---sort---")
	i := []int{5, 3, 2, 8, 7}
	s := []string{"d", "a", "f"}
	p := []struct {
		Name string
		Age  int
	}{
		{"Nancy", 20},
		{"Vera", 40},
		{"Mike", 30},
		{"Bob", 50},
	}
	fmt.Println(i, s, p)
	sort.Ints(i)
	sort.Strings(s)
	sort.Slice(p, func(i, j int) bool { return p[i].Name < p[j].Name }) //less関数
	fmt.Println(i, s, p)
	sort.Slice(p, func(i, j int) bool { return p[i].Age < p[j].Age })
	fmt.Println(i, s, p)

	//iota
	fmt.Println("---iota---")
	fmt.Println(c1, c2, c3)
	fmt.Println(KB, MB, GB)
}
