package interpreter

import "testing"

type TestCase struct {
	name  string
	stats map[string]float64
	rule  string
	want  bool
}

func TestDemo(t *testing.T) {
	stats := map[string]float64{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	testCases := []TestCase{
		{
			name:  "case1",
			stats: stats,
			rule:  "a > 1 && b > 10 && c < 5",
			want:  false,
		},
		{
			name:  "case2",
			stats: stats,
			rule:  "a > 1 && b > 1 && c < 5",
			want:  false,
		},
		{
			name:  "case3",
			stats: stats,
			rule:  "a < 5 && b > 1 && c < 10",
			want:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ari, _ := NewAlertRuleInterpreter(tc.rule)
			got := ari.Interpret(tc.stats)
			if got != tc.want {
				t.Fatalf("want %v got %v", tc.want, got)
			}
		})
	}
}
