package store

import (
	"fmt"
	"reflect"
	"testing"
)

func TestEvolvest_Del(t *testing.T) {
	type fields struct {
		Nodes map[string]string
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantVal string
		wantErr bool
	}{
		{
			name: "key exits",
			fields: fields{
				Nodes: map[string]string{
					"hello": "world",
				},
			},
			args: args{
				key: "hello",
			},
			wantVal: "world",
			wantErr: false,
		},
		{
			name: "key not exits",
			fields: fields{
				Nodes: map[string]string{
					"hello": "world",
				},
			},
			args: args{
				key: "hello123",
			},
			wantVal: "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(tt.name)
			e := &Evolvest{
				Nodes: tt.fields.Nodes,
			}
			gotVal, err := e.Del(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Del() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotVal, tt.wantVal) {
				t.Errorf("Del() gotVal = %v, want %v", gotVal, tt.wantVal)
			}
		})
	}
}

func TestEvolvest_Get(t *testing.T) {
	type fields struct {
		Nodes map[string]string
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantVal string
		wantErr bool
	}{
		{
			name: "key exits",
			fields: fields{
				Nodes: map[string]string{
					"hello": "world",
				},
			},
			args: args{
				key: "hello",
			},
			wantVal: "world",
			wantErr: false,
		},
		{
			name: "key not exits",
			fields: fields{
				Nodes: map[string]string{
					"hello": "world",
				},
			},
			args: args{
				key: "hello123",
			},
			wantVal: "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(tt.name)
			e := &Evolvest{
				Nodes: tt.fields.Nodes,
			}
			gotVal, err := e.Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotVal, tt.wantVal) {
				t.Errorf("Get() gotVal = %v, want %v", gotVal, tt.wantVal)
			}
		})
	}
}

func TestEvolvest_Set(t *testing.T) {
	type fields struct {
		Nodes map[string]string
	}
	type args struct {
		key string
		val string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantOldVal string
		wantExist  bool
	}{
		{
			name: "key exit",
			fields: fields{
				Nodes: map[string]string{
					"hello": "world",
				},
			},
			args: args{
				key: "hello",
				val: "123",
			},
			wantOldVal: "world",
			wantExist:  true,
		},
		{
			name: "key not exit",
			fields: fields{
				Nodes: map[string]string{
					"hello": "world",
				},
			},
			args: args{
				key: "hello123",
				val: "123",
			},
			wantOldVal: "",
			wantExist:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(tt.name)
			e := &Evolvest{
				Nodes: tt.fields.Nodes,
			}
			gotOldVal, goExist := e.Set(tt.args.key, tt.args.val)
			if goExist != tt.wantExist {
				t.Errorf("Set() exit = %v, wantExist %v", goExist, tt.wantExist)
				return
			}
			if !reflect.DeepEqual(gotOldVal, tt.wantOldVal) {
				t.Errorf("Set() gotOldVal = %v, want %v", gotOldVal, tt.wantOldVal)
			}
		})
	}
}

func TestNewEvolvest(t *testing.T) {
	tests := []struct {
		name string
		want *Evolvest
	}{
		{
			name: "normal",
			want: &Evolvest{Nodes: make(map[string]string, 17)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(tt.name)
			if got := NewEvolvest(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEvolvest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEvolvest_Save(t *testing.T) {
	type fields struct {
		Nodes map[string]string
	}
	tests := []struct {
		name     string
		fields   fields
		wantData []byte
		wantErr  bool
	}{
		{
			name: "nil",
			fields: fields{
				Nodes: nil,
			},
			wantData: []byte(`{"nodes":null}`),
			wantErr:  false,
		},
		{
			name: "empty",
			fields: fields{
				Nodes: map[string]string{},
			},
			wantData: []byte(`{"nodes":{}}`),
			wantErr:  false,
		},
		{
			name: "have values",
			fields: fields{
				Nodes: map[string]string{
					"hello": "world",
				},
			},
			wantData: []byte(`{"nodes":{"hello":"world"}}`),
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(tt.name)
			e := &Evolvest{
				Nodes: tt.fields.Nodes,
			}
			gotData, err := e.Save()
			fmt.Printf("gotData: %s\n", string(gotData))
			if (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("Save() gotData = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}

func TestEvolvest_Load(t *testing.T) {
	type fields struct {
		Nodes map[string]string
	}
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "nil",
			fields: fields{
				Nodes: map[string]string{
					"abc": "123",
				},
			},
			args: args{
				data: nil,
			},
			wantErr: true,
		},
		{
			name: "empty",
			fields: fields{
				Nodes: map[string]string{
					"abc": "123",
				},
			},
			args: args{
				data: []byte("{}"),
			},
			wantErr: false,
		},
		{
			name: "not empty",
			fields: fields{
				Nodes: map[string]string{
					"abc":   "123",
					"hello": "456",
				},
			},
			args: args{
				data: []byte(`{"Nodes":{"hello":"world"}}`),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(tt.name)
			e := &Evolvest{
				Nodes: tt.fields.Nodes,
			}
			if err := e.Load(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Load() error = %v, wantErr %v", err, tt.wantErr)
			}
			d, err := e.Save()
			fmt.Printf("data:%s, err:%v\n", string(d), err)
		})
	}
}