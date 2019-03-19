package gomygen

import (
	"os"
	"testing"
)

var tools = new(Tools)

func TestTools_GenerateDir(t *testing.T) {
	path := "/data/go_mygen"
	path2, err := tools.GenerateDir(path)
	if err != nil {
		t.Error(err)
	}
	if path+string(os.PathSeparator) != path2 {
		t.Error()
	}
	t.Log(path2)
}

func TestTools_CheckFileContainsChar(t *testing.T) {
	type args struct {
		filename string
		s        string
	}
	tests := []struct {
		name string
		t    *Tools
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t := &Tools{}
			if got := t.CheckFileContainsChar(tt.args.filename, tt.args.s); got != tt.want {
				t.Errorf("Tools.CheckFileContainsChar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTools_ReadFile(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		t    *Tools
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t := &Tools{}
			if got := t.ReadFile(tt.args.filename); got != tt.want {
				t.Errorf("Tools.ReadFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTools_WriteFile(t *testing.T) {
	type args struct {
		filename string
		data     string
	}
	tests := []struct {
		name      string
		t         *Tools
		args      args
		wantCount int
		wantErr   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := &Tools{}
			gotCount, err := ts.WriteFile(tt.args.filename, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Tools.WriteFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotCount != tt.wantCount {
				t.Errorf("Tools.WriteFile() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}

func TestTools_WriteFileAppend(t *testing.T) {
	type args struct {
		filename string
		data     string
	}
	tests := []struct {
		name      string
		t         *Tools
		args      args
		wantCount int
		wantErr   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := &Tools{}
			gotCount, err := ts.WriteFileAppend(tt.args.filename, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Tools.WriteFileAppend() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotCount != tt.wantCount {
				t.Errorf("Tools.WriteFileAppend() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}

func TestTools_CreateFile(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		t    *Tools
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := &Tools{}
			if got := ts.CreateFile(tt.args.path); got != tt.want {
				t.Errorf("Tools.CreateFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTools_CreateDir(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		t    *Tools
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := &Tools{}
			if got := ts.CreateDir(tt.args.path); got != tt.want {
				t.Errorf("Tools.CreateDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTools_IsDirOrFileExist(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		t    *Tools
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := &Tools{}
			if got := ts.IsDirOrFileExist(tt.args.path); got != tt.want {
				t.Errorf("Tools.IsDirOrFileExist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTools_IsDir(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		t    *Tools
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := &Tools{}
			if got := ts.IsDir(tt.args.filename); got != tt.want {
				t.Errorf("Tools.IsDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTools_IsFile(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		t    *Tools
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := &Tools{}
			if got := ts.IsFile(tt.args.filename); got != tt.want {
				t.Errorf("Tools.IsFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTools_isFileOrDir(t *testing.T) {
	type args struct {
		filename  string
		decideDir bool
	}
	tests := []struct {
		name string
		t    *Tools
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := &Tools{}
			if got := ts.isFileOrDir(tt.args.filename, tt.args.decideDir); got != tt.want {
				t.Errorf("Tools.isFileOrDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTools_Capitalize(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		t    *Tools
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t := &Tools{}
			if got := t.Capitalize(tt.args.s); got != tt.want {
				t.Errorf("Tools.Capitalize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTools_ToUpper(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		t    *Tools
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t := &Tools{}
			if got := t.ToUpper(tt.args.s); got != tt.want {
				t.Errorf("Tools.ToUpper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTools_ToJson(t *testing.T) {
	type args struct {
		s interface{}
	}
	tests := []struct {
		name string
		t    *Tools
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t := &Tools{}
			if got := t.ToJson(tt.args.s); got != tt.want {
				t.Errorf("Tools.ToJson() = %v, want %v", got, tt.want)
			}
		})
	}
}
