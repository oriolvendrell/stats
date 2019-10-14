package stats

import (
	"math"
	"testing"
)

// func TestOneSampleZTest(t *testing.T) {
// 	type args struct {
// 		sample []float64
// 		pop    pd.Normal
// 		alpha  float64
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want bool
// 	}{
// 		{"Normal case", args{[]float64{0.1, 0.02, -0.3, 0.47, 0.015, 0.21, -0.32, -0.05, -0.1, 0.15, 0.17, 0.08, -0.125}, pd.NewNormal(), 0.05}, true},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := OneSampleZTest(tt.args.sample, tt.args.pop, tt.args.alpha); got != tt.want {
// 				t.Errorf("OneSampleZTest() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func TestOneSampleTTest(t *testing.T) {
	type args struct {
		sample  []float64
		popmean float64
		alpha   float64
		tails   int
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 float64
	}{
		{"Normal case", args{[]float64{0.1, 0.02, -0.3, 0.47, 0.015, 0.21, -0.32, -0.05, -0.1, 0.15, 0.17, 0.08, -0.125}, 0.0, 0.05, 1}, true, 1.7823},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := OneSampleTTest(tt.args.sample, tt.args.popmean, tt.args.alpha, tt.args.tails)
			if got != tt.want {
				t.Errorf("OneSampleTTest() got = %v, want %v", got, tt.want)
			}
			if math.Abs(got1-tt.want1) > 1e-3 {
				t.Errorf("OneSampleTTest() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
