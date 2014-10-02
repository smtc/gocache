package gocache

import "testing"

func TestCache(t *testing.T) {
	i := "100"
	c := GetCache()
	c.Set("i", i, 0)

	i2, suc := c.Get("i")
	if suc != true {
		t.Fatal("i should existed")
	}
	if i2.(string) != "100" {
		t.Fatal("i should be 100")
	}

	c.Store("s", "10000", 0)
	var s string
	suc, err := c.Retrieve("s", &s)
	if suc == false {
		t.Fatal("i should existed")
	}
	if err != nil {
		t.Fatal(err.Error())
	}
	if s != "10000" {
		t.Fatal("s should be 10000")
	}

	c.Delete("i")

	i2, suc = c.Get("i")
	if suc == true {
		t.Fatal("i should be deleted")
	}

}
