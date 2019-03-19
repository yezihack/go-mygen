package gomygen

import (
	"sync"
	"testing"

	"github.com/yezihack/colorlog"
)

var lg *Logic

func init() {
	DbConn := DBConfig{
		Host:   "localhost",
		Port:   3308,
		Name:   "root",
		Pass:   "123456",
		DBName: "kindled",
	}
	db, err := InitDB(DbConn)
	if db == nil || err != nil {
		panic(err)
	}
	colorlog.Info("数据库连接成功")
	lg = new(Logic)
	lg.DB = NewDB(db)
}

func TestLogic_CreateMarkdown(t *testing.T) {
	t.Log("ok")
}
func BenchmarkLogic_CreateMarkdown(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lg.CreateMarkdown()
	}
}

func TestLogic_CreateEntity(t *testing.T) {
	type fields struct {
		T    *Tools
		DB   *ModelS
		Path string
		Once sync.Once
	}
	type args struct {
		formatList []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logic{
				T:    tt.fields.T,
				DB:   tt.fields.DB,
				Path: tt.fields.Path,
				Once: tt.fields.Once,
			}
			if err := l.CreateEntity(tt.args.formatList); (err != nil) != tt.wantErr {
				t.Errorf("Logic.CreateEntity() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLogic_CreateCURD(t *testing.T) {
	type fields struct {
		T    *Tools
		DB   *ModelS
		Path string
		Once sync.Once
	}
	type args struct {
		formatList []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logic{
				T:    tt.fields.T,
				DB:   tt.fields.DB,
				Path: tt.fields.Path,
				Once: tt.fields.Once,
			}
			if err := l.CreateCURD(tt.args.formatList); (err != nil) != tt.wantErr {
				t.Errorf("Logic.CreateCURD() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLogic_GetMysqlDir(t *testing.T) {
	type fields struct {
		T    *Tools
		DB   *ModelS
		Path string
		Once sync.Once
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logic{
				T:    tt.fields.T,
				DB:   tt.fields.DB,
				Path: tt.fields.Path,
				Once: tt.fields.Once,
			}
			if got := l.GetMysqlDir(); got != tt.want {
				t.Errorf("Logic.GetMysqlDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogic_GetRoot(t *testing.T) {
	type fields struct {
		T    *Tools
		DB   *ModelS
		Path string
		Once sync.Once
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logic{
				T:    tt.fields.T,
				DB:   tt.fields.DB,
				Path: tt.fields.Path,
				Once: tt.fields.Once,
			}
			if got := l.GetRoot(); got != tt.want {
				t.Errorf("Logic.GetRoot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogic_GenerateDBStructure(t *testing.T) {
	type fields struct {
		T    *Tools
		DB   *ModelS
		Path string
		Once sync.Once
	}
	type args struct {
		tableName    string
		tableComment string
		path         string
		tableDesc    []*TableDesc
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logic{
				T:    tt.fields.T,
				DB:   tt.fields.DB,
				Path: tt.fields.Path,
				Once: tt.fields.Once,
			}
			if err := l.GenerateDBStructure(tt.args.tableName, tt.args.tableComment, tt.args.path, tt.args.tableDesc); (err != nil) != tt.wantErr {
				t.Errorf("Logic.GenerateDBStructure() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLogic_GenerateDBEntity(t *testing.T) {
	type fields struct {
		T    *Tools
		DB   *ModelS
		Path string
		Once sync.Once
	}
	type args struct {
		req *EntityReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logic{
				T:    tt.fields.T,
				DB:   tt.fields.DB,
				Path: tt.fields.Path,
				Once: tt.fields.Once,
			}
			if err := l.GenerateDBEntity(tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("Logic.GenerateDBEntity() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLogic_GenerateCURDFile(t *testing.T) {
	type fields struct {
		T    *Tools
		DB   *ModelS
		Path string
		Once sync.Once
	}
	type args struct {
		tableName    string
		tableComment string
		tableDesc    []*TableDesc
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logic{
				T:    tt.fields.T,
				DB:   tt.fields.DB,
				Path: tt.fields.Path,
				Once: tt.fields.Once,
			}
			if err := l.GenerateCURDFile(tt.args.tableName, tt.args.tableComment, tt.args.tableDesc); (err != nil) != tt.wantErr {
				t.Errorf("Logic.GenerateCURDFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLogic_GenerateExample(t *testing.T) {
	type fields struct {
		T    *Tools
		DB   *ModelS
		Path string
		Once sync.Once
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logic{
				T:    tt.fields.T,
				DB:   tt.fields.DB,
				Path: tt.fields.Path,
				Once: tt.fields.Once,
			}
			l.GenerateExample(tt.args.name)
		})
	}
}

func TestLogic_GenerateTableList(t *testing.T) {
	type fields struct {
		T    *Tools
		DB   *ModelS
		Path string
		Once sync.Once
	}
	type args struct {
		list []*TableList
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logic{
				T:    tt.fields.T,
				DB:   tt.fields.DB,
				Path: tt.fields.Path,
				Once: tt.fields.Once,
			}
			if err := l.GenerateTableList(tt.args.list); (err != nil) != tt.wantErr {
				t.Errorf("Logic.GenerateTableList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLogic_GenerateSQL(t *testing.T) {
	type fields struct {
		T    *Tools
		DB   *ModelS
		Path string
		Once sync.Once
	}
	type args struct {
		info         *SqlInfo
		tableComment string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logic{
				T:    tt.fields.T,
				DB:   tt.fields.DB,
				Path: tt.fields.Path,
				Once: tt.fields.Once,
			}
			if err := l.GenerateSQL(tt.args.info, tt.args.tableComment); (err != nil) != tt.wantErr {
				t.Errorf("Logic.GenerateSQL() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLogic_GenerateMarkdown(t *testing.T) {
	type fields struct {
		T    *Tools
		DB   *ModelS
		Path string
		Once sync.Once
	}
	type args struct {
		data *MarkDownData
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Logic{
				T:    tt.fields.T,
				DB:   tt.fields.DB,
				Path: tt.fields.Path,
				Once: tt.fields.Once,
			}
			if err := l.GenerateMarkdown(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Logic.GenerateMarkdown() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
