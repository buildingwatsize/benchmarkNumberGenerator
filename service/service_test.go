package service_test

import (
	"benchmarkNumberGenerator/service"
	"testing"
)

func BenchmarkSolution1_GetRandomByNanosecond(b *testing.B) {
	for n := 0; n < b.N; n++ {
		randStr := service.Solution1()
		service.AddSet(randStr)
	}
	service.ClearSet()
}
func BenchmarkSolution2_randIntn(b *testing.B) {
	for n := 0; n < b.N; n++ {
		randStr := service.Solution2()
		service.AddSet(randStr)
	}
	service.ClearSet()
}
func BenchmarkSolution3_randInt63(b *testing.B) {
	for n := 0; n < b.N; n++ {
		randStr := service.Solution3()
		service.AddSet(randStr)
	}
	service.ClearSet()
}
func BenchmarkSolution4_CryptoRand(b *testing.B) {
	for n := 0; n < b.N; n++ {
		randStr := service.Solution4()
		service.AddSet(randStr)
	}
	service.ClearSet()
}
func BenchmarkSolution5_GetRandomByNanosecondShuffle(b *testing.B) {
	for n := 0; n < b.N; n++ {
		randStr := service.Solution5()
		service.AddSet(randStr)
	}
	service.ClearSet()
}
