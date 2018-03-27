package a2conf_test

import (
	"io"
	"io/ioutil"
	"strings"
	"testing"

	. "github.com/bearx3f/a2conf"
)

var (
	opt = FormatOption{
		Tab2Space: true,
		TabSize:   4,
	}

	test1In, _  = ioutil.ReadFile("./test/test1_input.txt")
	test1Out, _ = ioutil.ReadFile("./test/test1_output.txt")

	test2In, _  = ioutil.ReadFile("./test/test2_input.txt")
	test2Out, _ = ioutil.ReadFile("./test/test2_output.txt")

	test3In, _  = ioutil.ReadFile("./test/test3_input.txt")
	test3Out, _ = ioutil.ReadFile("./test/test3_output.txt")

	test4In, _  = ioutil.ReadFile("./test/test4_input.txt")
	test4Out, _ = ioutil.ReadFile("./test/test4_output.txt")
)

func TestFormat(t *testing.T) {
	type args struct {
		in     io.Reader
		fmtOpt FormatOption
	}
	tests := []struct {
		name       string
		args       args
		wantOutput string
		wantErr    bool
	}{
		{
			name: "Test1",
			args: args{
				strings.NewReader(string(test1In)),
				opt,
			},
			wantOutput: string(test1Out),
			wantErr:    false,
		},
		{
			name: "Test2",
			args: args{
				strings.NewReader(string(test2In)),
				opt,
			},
			wantOutput: string(test2Out),
			wantErr:    false,
		},
		{
			name: "Test3",
			args: args{
				strings.NewReader(string(test3In)),
				opt,
			},
			wantOutput: string(test3Out),
			wantErr:    false,
		},
		{
			name: "Test4",
			args: args{
				strings.NewReader(string(test4In)),
				opt,
			},
			wantOutput: string(test4Out),
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOutput, err := Format(tt.args.in, tt.args.fmtOpt)
			if (err != nil) != tt.wantErr {
				t.Errorf("Format() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOutput != tt.wantOutput {
				t.Errorf("Format() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
