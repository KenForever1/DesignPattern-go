package factory_test

import (
	"DesignPattern-go/factory"
	"reflect"
	"testing"
)

func TestNewIRuleConfigParser(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name string
		args args
		want factory.IRuleConfigParser
	}{
		{
			name: "json",
			args: args{t: "json"},
			want: factory.JsonRuleConfigParser{},
		},
		{
			name: "yaml",
			args: args{t: "yaml"},
			want: factory.YamlRuleConfigParser{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := factory.NewIRuleConfigParser(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIRuleConfigParser()=%v, want %v", got, tt.want)
			}
		})
	}
}
