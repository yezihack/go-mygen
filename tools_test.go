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
