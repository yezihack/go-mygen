package gomygen

import (
	"database/sql"
	"reflect"
	"testing"
)

func TestModelS_GetTableList(t *testing.T) {
	type fields struct {
		DB *sql.DB
		T  *Tools
	}
	type args struct {
		onTable []string
	}
	tests := []struct {
		name            string
		fields          fields
		args            args
		wantTableResult map[string]string
		wantErr         bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ModelS{
				DB: tt.fields.DB,
				T:  tt.fields.T,
			}
			gotTableResult, err := m.GetTableList(tt.args.onTable...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ModelS.GetTableList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTableResult, tt.wantTableResult) {
				t.Errorf("ModelS.GetTableList() = %v, want %v", gotTableResult, tt.wantTableResult)
			}
		})
	}
}

func TestModelS_GetTableDesc(t *testing.T) {
	type fields struct {
		DB *sql.DB
		T  *Tools
	}
	type args struct {
		tableName string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantReply []*TableDesc
		wantErr   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ModelS{
				DB: tt.fields.DB,
				T:  tt.fields.T,
			}
			gotReply, err := m.GetTableDesc(tt.args.tableName)
			if (err != nil) != tt.wantErr {
				t.Errorf("ModelS.GetTableDesc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotReply, tt.wantReply) {
				t.Errorf("ModelS.GetTableDesc() = %v, want %v", gotReply, tt.wantReply)
			}
		})
	}
}
