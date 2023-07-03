package repkg

import (
	"os"
	"testing"
)

// func TestCalc(t *testing.T) {
// 	type args struct {
// 		in  *bufio.Reader
// 		out *bufio.Writer
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			Calc(tt.args.in, tt.args.out)
// 		})
// 	}
// }

func BenchmarkCalcScan(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()

	in, err := os.OpenFile("testin", os.O_RDONLY, os.ModePerm)
	if err != nil {
		b.Error(err)
	}
	defer in.Close()
	out, err := os.OpenFile("testout", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.ModePerm)
	if err != nil {
		b.Error(err)
	}
	defer out.Close()

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		CalcScan(in, out)
		b.StopTimer()
		in.Seek(0, 0)
	}
}

func BenchmarkCalcReadLine(b *testing.B) {
	b.ReportAllocs()
	b.StopTimer()

	in, err := os.OpenFile("testin", os.O_RDONLY, os.ModePerm)
	if err != nil {
		b.Error(err)
	}
	defer in.Close()
	out, err := os.OpenFile("testout", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.ModePerm)
	if err != nil {
		b.Error(err)
	}
	defer out.Close()

	for i := 0; i < b.N; i++ {
		// in := bufio.NewReader(in)
		// out := bufio.NewWriter(out)
		// out := bufio.NewWriter(&buf)
		b.StartTimer()
		CalcReadLine(in, out)
		b.StopTimer()
		in.Seek(0, 0)
	}
}
