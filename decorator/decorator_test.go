package decorator

import "testing"

func TestDecorator(t *testing.T) {
	var peter Peter

	var jack Jack
	jack.person = peter
	var hat Hat
	hat.person = peter

	t.Log(jack.Cost())
	jack.Show()

	t.Log(hat.Cost())
	hat.Show()
}
