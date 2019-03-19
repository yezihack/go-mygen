package gomygen

import (
	"database/sql"
	"reflect"
	"testing"
)

func TestModelS_Find(t *testing.T) {
	type fields struct {
		DB *sql.DB
		T  *Tools
	}
	type args struct {
		sql  string
		args []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []map[string]interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ModelS{
				DB: tt.fields.DB,
				T:  tt.fields.T,
			}
			got, err := m.Find(tt.args.sql, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ModelS.Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ModelS.Find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestModelS_First(t *testing.T) {
	type fields struct {
		DB *sql.DB
		T  *Tools
	}
	type args struct {
		sql  string
		args []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ModelS{
				DB: tt.fields.DB,
				T:  tt.fields.T,
			}
			got, err := m.First(tt.args.sql, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ModelS.First() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ModelS.First() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestModelS_Pluck(t *testing.T) {
	type fields struct {
		DB *sql.DB
		T  *Tools
	}
	type args struct {
		sql  string
		name string
		args []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ModelS{
				DB: tt.fields.DB,
				T:  tt.fields.T,
			}
			got, err := m.Pluck(tt.args.sql, tt.args.name, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ModelS.Pluck() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ModelS.Pluck() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestModelS_Update(t *testing.T) {
	type fields struct {
		DB *sql.DB
		T  *Tools
	}
	type args struct {
		sql  string
		args []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ModelS{
				DB: tt.fields.DB,
				T:  tt.fields.T,
			}
			got, err := m.Update(tt.args.sql, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ModelS.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ModelS.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestModelS_Delete(t *testing.T) {
	type fields struct {
		DB *sql.DB
		T  *Tools
	}
	type args struct {
		sql  string
		args []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ModelS{
				DB: tt.fields.DB,
				T:  tt.fields.T,
			}
			got, err := m.Delete(tt.args.sql, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ModelS.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ModelS.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestModelS_Insert(t *testing.T) {
	type fields struct {
		DB *sql.DB
		T  *Tools
	}
	type args struct {
		sql  string
		args []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &ModelS{
				DB: tt.fields.DB,
				T:  tt.fields.T,
			}
			got, err := m.Insert(tt.args.sql, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ModelS.Insert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ModelS.Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}
