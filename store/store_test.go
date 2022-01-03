package store

import "testing"

func TestCache(t *testing.T) {
	ts := New(DefaultExpiration, 0)
	x, found := ts.Get("TestKey")
	if found || x != "" {
		t.Error("Getting TestValue found that shouldn't exist", x)
	}
	y, found := ts.Get("TestKey1")
	if found || y != "" {
		t.Error("Getting TestValue1 found value that shouldn't exist:", y)
	}

	z, found := ts.Get("TestKey2")
	if found || z != "" {
		t.Error("Getting TestValue2 found value that shouldn't exist:", z)
	}

	ts.Set("TestKey", "TestValue", DefaultExpiration)

	a, found := ts.Get("TestKey")
	if !found {
		t.Error("TestKey not found")
	}
	if a == "" {
		t.Error("TestKey is nil")
	}

}
func TestDelete(t *testing.T) {
	ts := New(DefaultExpiration, 0)
	ts.Set("foo", "bar", DefaultExpiration)
	ts.Delete("foo")
	x, found := ts.Get("foo")
	if found {
		t.Error("foo was found, but it should have been deleted")
	}
	if x != "" {
		t.Error("x is not nil:", x)
	}
}
func TestFlush(t *testing.T) {
	ts := New(DefaultExpiration, 0)
	ts.Set("foo", "bar", DefaultExpiration)
	ts.Set("testkey", "testvalue", DefaultExpiration)
	ts.Flush()
	x, found := ts.Get("foo")
	if found {
		t.Error("foo was found, but it should have been deleted")
	}
	if x != "" {
		t.Error("x is not nil:", x)
	}
	x, found = ts.Get("testkey")
	if found {
		t.Error("testkey was found, but it should have been deleted")
	}
	if x != "" {
		t.Error("x is not nil:", x)
	}
}
