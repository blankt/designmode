package factory_method

import "testing"

func TestFactoryMethod(t *testing.T) {
	scBun := new(SCBun)
	dbBun := new(DBBun)

	scCabbageBun := scBun.generate("SCCabbageBun")
	scCabbageBun.create()
	scPotatoBun := scBun.generate("SCPotatoBun")
	scPotatoBun.create()

	dbCabbageBun := dbBun.generate("DBCabbageBun")
	dbCabbageBun.create()
	dbPotatoBun := dbBun.generate("DBPotatoBun")
	dbPotatoBun.create()
}
