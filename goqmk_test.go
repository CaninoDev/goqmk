package goqmk

import (
	"reflect"
	"testing"
)

func TestAllKeyboardsList(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AllKeyboardsList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AllKeyboardsList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDownloadHex(t *testing.T) {
	type args struct {
		keyboard string
		keymap   string
	}
	tests := []struct {
		name string
		args args
		want error
	}{

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DownloadHex(tt.args.keyboard, tt.args.keymap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DownloadHex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetBootLoaderType(t *testing.T) {
	type args struct {
		keyboard string
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
			if got := GetBootLoaderType(tt.args.keyboard); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBootLoaderType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKeyboardData(t *testing.T) {
	type args struct {
		keyboard string
	}
	tests := []struct {
		name string
		args args
		want Keyboard
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KeyboardData(tt.args.keyboard); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KeyboardData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKeymaps(t *testing.T) {
	type args struct {
		keyboard string
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
			if got := Keymaps(tt.args.keyboard); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Keymaps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLayouts(t *testing.T) {
	type args struct {
		keyboard string
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
			if got := Layouts(tt.args.keyboard); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Layouts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queryQMK(t *testing.T) {
	type args struct {
		keyboard string
	}
	tests := []struct {
		name string
		args args
		want map[string]Keyboard
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := queryQMK(tt.args.keyboard); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("queryQMK() = %v, want %v", got, tt.want)
			}
		})
	}
}