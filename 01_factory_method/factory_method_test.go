package factorymethod

import (
	"reflect"
	"testing"
)

type Args struct {
	configType string
}

func TestNewIRuleConfigParserFactory(t *testing.T) {
	testCases := []struct {
		name string
		args Args
		want IRuleConfigParserFactory
	}{
		{
			name: "json",
			args: Args{configType: "json"},
			want: JsonRuleConfigParserFactory{},
		},
		{
			name: "yaml",
			args: Args{configType: "yaml"},
			want: YamlRuleConfigParserFactory{},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := NewIRuleConfigParserFactory(tt.args.configType)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIRuleConfigParserFactory() = %v, want %v", got, tt.want)
			}
		})
	}
}
