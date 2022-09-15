package util

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"
const numeric = "1234567890"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generate a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomIntString(n int) string {
	var sb strings.Builder
	k := len(numeric)

	for i := 0; i < n; i++ {
		c := numeric[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomPhoneNumber generate random phone number
func RandomPhoneNumber() string {
	return RandomIntString(3) + "-" + RandomIntString(3) + "-" + RandomIntString(4)
}

// RandomEmail generate a random email address
func RandomEmail() string {
	name := RandomString(5)
	return name + "@email.com"
}

// RandomStreetAddress1 generate random street address
func RandomStreetAddress1() string {
	return RandomIntString(4) + " " + RandomString(8) + " Street"
}

// RandomStreetAddress2 generate random street address
func RandomStreetAddress2() string {
	return "P.O. Box " + RandomIntString(5)
}

// RandomState select random state from list
func RandomState() string {
	states := []string{"AZ", "CA", "TX", "OH", "NM", "MA", "NY", "WV", "SD", "ND", "WI", "IA"}
	n := len(states)
	return states[rand.Intn(n)]
}

func randomIntRange(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

// RandomDate generate random date
func RandomDate() string {
	month := rand.Intn(12)
	if month == 0 {
		month += 1
	}

	// get current year
	t := time.Now()
	currYear := t.Year()

	return fmt.Sprint(month) + "/" + strconv.Itoa(randomIntRange(10, 28)) + "/" + strconv.Itoa(randomIntRange(currYear-90, currYear))
}

// RandomAppointmentTime generate random time minute increment from defined slice
func RandomAppointmentTime() string {
	minutes := []string{"00", "15", "30", "45"}
	hour := randomIntRange(8, 23)

	return strconv.Itoa(hour) + ":" + minutes[rand.Intn(len(minutes))]
}

// RandomPath generate random string path name
func RandomPath() string {
	pathName := RandomString(4) + "/" + RandomString(5) + "/" + RandomString(4)
	return pathName
}
