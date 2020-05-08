package service

import (
	cryptoRand "crypto/rand"
	"fmt"
	"math/big"
	"math/rand"
	"strconv"
	"time"
)

var (
	times = 1000000 // default value (1M)
	Set   = make(map[string]bool)
)

func AddSet(key string) {
	Set[key] = true
}
func ClearSet() {
	Set = map[string]bool{}
}

func Run(numberTimes int) {
	if numberTimes != 0 {
		times = numberTimes
	}
	fmt.Println("Test:", times, "items")
	startTime := time.Now()
	doneTime := time.Now().Sub(startTime)

	// ################
	// #### SOLUTION 1 - GetRandomByNanosecond
	// ################
	startTime = time.Now()

	fmt.Println("SOLUTION 1 - GetRandomByNanosecond")
	fmt.Println("Example SOLUTION 1:", Solution1())
	for i := 0; i < times; i++ {
		randStr := Solution1()
		// fmt.Println(randStr)
		AddSet(randStr)
	}

	fmt.Println("Solution 1: Unique:", len(Set), "items. UniquePercent:", (100 * float64(len(Set)) / float64(times)), "%")
	ClearSet()

	doneTime = time.Now().Sub(startTime)
	fmt.Println("Time Usage:", doneTime)

	// ################
	// #### SOLUTION 2 - rand.Intn()
	// ################
	startTime = time.Now()

	fmt.Println("SOLUTION 2 - rand.Intn()")
	fmt.Println("Example SOLUTION 2:", Solution2())
	for i := 0; i < times; i++ {
		randStr := Solution2()
		// fmt.Println(randStr)
		AddSet(randStr)
	}

	fmt.Println("Solution 2: Unique:", len(Set), "items. UniquePercent:", (100 * float64(len(Set)) / float64(times)), "%")
	ClearSet()

	doneTime = time.Now().Sub(startTime)
	fmt.Println("Time Usage:", doneTime)

	// ################
	// #### SOLUTION 3 - rand.Int63()
	// ################
	startTime = time.Now()

	fmt.Println("SOLUTION 3 - rand.Int63()")
	fmt.Println("Example SOLUTION 3:", Solution3())
	for i := 0; i < times; i++ {
		randStr := Solution3()
		// fmt.Println(randStr)
		AddSet(randStr)
	}

	fmt.Println("Solution 3: Unique:", len(Set), "items. UniquePercent:", (100 * float64(len(Set)) / float64(times)), "%")
	ClearSet()

	doneTime = time.Now().Sub(startTime)
	fmt.Println("Time Usage:", doneTime)

	// ################
	// #### SOLUTION 4 - CryptoRand
	// ################
	startTime = time.Now()

	fmt.Println("SOLUTION 4 - CryptoRand")
	fmt.Println("Example SOLUTION 4:", Solution4())
	for i := 0; i < times; i++ {
		randStr := Solution4()
		// fmt.Println(randStr)
		AddSet(randStr)
	}

	fmt.Println("Solution 4: Unique:", len(Set), "items. UniquePercent:", (100 * float64(len(Set)) / float64(times)), "%")
	ClearSet()

	doneTime = time.Now().Sub(startTime)
	fmt.Println("Time Usage:", doneTime)

	// ################
	// #### SOLUTION 5 - SOLUTION 1 + SHUFFLE
	// ################
	startTime = time.Now()

	fmt.Println("SOLUTION 5 - SOLUTION 1 (GetRandomByNanosecond) + SHUFFLE")
	fmt.Println("Example SOLUTION 5:", Solution5())
	for i := 0; i < times; i++ {
		randStr := Solution5()
		// fmt.Println(randStr)
		AddSet(randStr)
	}

	fmt.Println("Solution 5: Unique:", len(Set), "items. UniquePercent:", (100 * float64(len(Set)) / float64(times)), "%")
	ClearSet()

	doneTime = time.Now().Sub(startTime)
	fmt.Println("Time Usage:", doneTime)
}

func Solution1() string {
	timeran := time.Now().UnixNano() / int64(time.Nanosecond) // Get Time as nano second
	timestr := strconv.FormatInt(timeran, 10)                 // Convert to String

	ninestr := timestr[6 : len(timestr)-4] // Strip last 9-digit

	return ninestr
}

func Solution2() string {
	randItoa := strconv.Itoa(rand.Intn(100000000))
	result := fmt.Sprintf("%09v", randItoa)
	return result
}

func Solution3() string {
	randFormatInt := strconv.FormatInt(rand.Int63(), 10)[:9]
	result := fmt.Sprintf("%09v", randFormatInt)
	return result
}

func Solution4() string {
	bigIntResult, _ := cryptoRand.Int(cryptoRand.Reader, big.NewInt(100000000))
	result := fmt.Sprintf("%09v", bigIntResult)
	return result
}

func Solution5() string {
	solution1String := Solution1()
	byteString := []byte(solution1String)
	// rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(byteString), func(i, j int) {
		byteString[i], byteString[j] = byteString[j], byteString[i]
	})
	strShuffle := string(byteString)
	return strShuffle
}
