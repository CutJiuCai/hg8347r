package hg8347r

import "testing"

func TestListDevices(t *testing.T) {
	router := New("http://localhost:8888", "user", "<your_password>")
	d := router.ListDevices()
	t.Log(d)
}
