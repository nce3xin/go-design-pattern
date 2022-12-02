package simplefactory

import (
	"reflect"
	"testing"
)

type Args struct {
	configType string
}

func TestNewRuleConfigParser(t *testing.T) {
	tests := []struct {
		name string
		args Args
		want IRuleConfigParser
	}{
		{
			name: "json",
			args: Args{configType: "json"},
			want: JsonRuleConfigParser{},
		},
		{
			name: "yaml",
			args: Args{configType: "yaml"},
			want: YamlRuleConfigParser{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewRuleConfigParser(tt.args.configType)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRuleConfigParser() = %v, want %v", got, tt.want)
			}
		})
	}
}
