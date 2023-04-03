package dechutil

import (
	"math/rand"
	"strings"
	"time"
)

// Random Min to Max-1
func RandomInt(min, max int) int {
	seed := rand.NewSource(time.Now().UnixNano())
	rand2 := rand.New(seed)
	return rand2.Intn(max-min) + min
}

// Random Min to Max-1
func RandomFloat(min, max float64, decimal int) float64 {
	seed := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(seed)

	result := min + rand.Float64()*(max-min)
	result = RoundFloat(result, decimal)
	return result
}

func RandomString(num int, hasLower, hasUpper, hasNumber, isRemoveSameChar bool) string {
	charsetLower := "abcdefghijklmnopqrstuvwxyz"
	charsetUpper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	charsetNumber := "1234567890"
	charsetSame := "1lkKcCoO0pPquUxXyYzZ"

	var sb strings.Builder
	if hasLower {
		sb.WriteString(charsetLower)
	}

	if hasUpper {
		sb.WriteString(charsetUpper)
	}

	if hasNumber {
		sb.WriteString(charsetNumber)
	}

	charset := sb.String()

	if isRemoveSameChar {
		for _, v := range charsetSame {
			charset = strings.ReplaceAll(charset, string(v), "")
		}
	}

	seed := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(seed)

	result := ""
	ch := make(chan string)
	for {
		go randomChar(rand, charset, ch)
		result = result + string(<-ch)

		if len(result) >= num {
			close(ch)
			break
		}
	}

	/*
		for {
			if len(result) >= num {
				return result
			}

			c := charset[rand.Intn(len(charset))]
			result = result + string(c)
		}
	*/
	return result
}

func randomChar(rand *rand.Rand, charset string, ch chan string) {
	c := charset[rand.Intn(len(charset))]
	ch <- string(c)
}
