package facade

import "testing"

func TestNewAModule(t *testing.T) {
	api := NewApi()
	api.Test()
}
