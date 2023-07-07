package simple_factory

import "testing"

func TestNewApi(t *testing.T) {
	api := NewApi(1)
	if s := api.say("pxy"); s != "hi pxy" {
		t.Fatal("get api error")
	}
	api = NewApi(2)
	if s := api.say("tlf"); s != "hello tlf" {
		t.Fatal("get api error")
	}
}
