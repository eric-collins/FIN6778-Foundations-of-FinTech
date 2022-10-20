package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

func find_hash(str string, zeros int, start int, end int, done chan int) {

	for i := start; i <= end; i++ {
		var str_nonce string = strconv.FormatInt(int64(i), 10)
		var to_hash string = str + " " + str_nonce
		var zero_checks string = strings.Repeat("0", zeros)

		h := sha256.New()
		h.Write([]byte(to_hash))
		hash := hex.EncodeToString(h.Sum(nil))
		first_chars := string(hash[0:zeros])

		// if i%50000000 == 0 {
		// 	fmt.Println(first_chars, zero_checks, i)
		// }

		if first_chars == zero_checks {
			fmt.Println(to_hash, hash, first_chars, i)
			done <- 0
			break
		}
	}
}

func main() {

	ZEROS := 8
	NUM_ROUTINES := 100
	done := make(chan int)
	//adjustment := 187173821712

	bytespace := int(math.Pow(16, float64(ZEROS+1))) // - adjustment
	iter_breaks := bytespace / NUM_ROUTINES
	iterations := make([][]int, 0)

	for i := 0; i <= NUM_ROUTINES+2; i++ {
		it_start := iter_breaks * i     // + adjustment
		it_end := iter_breaks * (i + 1) // + adjustment
		it_slice := make([]int, 0)

		it_slice = append(it_slice, it_start, it_end)

		iterations = append(iterations, it_slice)
	}

	//fmt.Println(iterations)

	var input string = "Eric Georgette Joe"

	start := time.Now()
	for _, v := range iterations {

		start_it := v[0]
		end_it := v[1]

		//fmt.Println(start_it, end_it)

		go find_hash(input, ZEROS, start_it, end_it, done)
		//fmt.Println(i)

	}

	<-done
	duration := time.Since(start)
	fmt.Println(duration)
}

//Eric Georgette Joe 28179872686
//Eric Georgette Joe 187173821712
