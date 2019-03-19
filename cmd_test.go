package gomygen

import (
	"testing"
)

func Test_setFormat(t *testing.T) {
	setFormat()
}

func TestCmd(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Cmd()
		})
	}
}

func TestCommands(t *testing.T) {
	type args struct {
		DbConn DBConfig
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
			if err := Commands(tt.args.DbConn); (err != nil) != tt.wantErr {
				t.Errorf("Commands() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
