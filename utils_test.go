package gomygen

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestClean(t *testing.T) {
	fmt.Println(strings.Repeat("曾经沧海难为水，除却巫山不是云", 2000))
	time.Sleep(time.Millisecond * 700)
	Clean()
}

func TestCreateDir(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateDir(tt.args.path); got != tt.want {
				t.Errorf("CreateDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWriteFile(t *testing.T) {
	type args struct {
		path string
		data string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := WriteFile(tt.args.path, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("WriteFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWriteAppendFile(t *testing.T) {
	type args struct {
		path string
		data string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := WriteAppendFile(tt.args.path, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("WriteAppendFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestShowCmdHelp(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ShowCmdHelp()
		})
	}
}

func TestGetExeRootDir(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetExeRootDir(); got != tt.want {
				t.Errorf("GetExeRootDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRootPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRootPath(tt.args.path); got != tt.want {
				t.Errorf("GetRootPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubStr(t *testing.T) {
	type args struct {
		s      string
		pos    int
		length int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SubStr(tt.args.s, tt.args.pos, tt.args.length); got != tt.want {
				t.Errorf("SubStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrMsg(t *testing.T) {
	type args struct {
		msg string
		err error
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ErrMsg(tt.args.msg, tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ErrMsg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGofmt(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Gofmt(tt.args.path); got != tt.want {
				t.Errorf("Gofmt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetOs(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOs(); got != tt.want {
				t.Errorf("GetOs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExecCommand(t *testing.T) {
	type args struct {
		name string
		args []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExecCommand(tt.args.name, tt.args.args...); got != tt.want {
				t.Errorf("ExecCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatField(t *testing.T) {
	type args struct {
		field   string
		formats []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatField(tt.args.field, tt.args.formats); got != tt.want {
				t.Errorf("FormatField() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInArrayString(t *testing.T) {
	type args struct {
		str string
		arr []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InArrayString(tt.args.str, tt.args.arr); got != tt.want {
				t.Errorf("InArrayString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckCharDoSpecial(t *testing.T) {
	type args struct {
		s    string
		char byte
		regs string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckCharDoSpecial(tt.args.s, tt.args.char, tt.args.regs); got != tt.want {
				t.Errorf("CheckCharDoSpecial() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckCharDoSpecialArr(t *testing.T) {
	type args struct {
		s    string
		char byte
		reg  string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckCharDoSpecialArr(tt.args.s, tt.args.char, tt.args.reg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckCharDoSpecialArr() = %v, want %v", got, tt.want)
			}
		})
	}
}
