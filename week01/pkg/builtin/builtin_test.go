package builtin

import "testing"

func BenchmarkArraySum(b *testing.B) {
	a := [10]int{}
	for i := 0; i < 10; i++ {
		a[i] = i
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = AarraySum(a)
	}
}

func BenchmarkSliceSum(b *testing.B) {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = SliceSum(s)
	}
}

func BenchmarkCountChars(b *testing.B){
	str :="hello world 12345566771fgd"

	b.ResetTimer()
	for i:=0;i < b.N;i++{
		_=CountChars(str)
	}
	
}