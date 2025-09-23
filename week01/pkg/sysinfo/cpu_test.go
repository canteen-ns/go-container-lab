package sysinfo


import "testing"

func TestCPUCores(t *testing.T) {
	n, err := CPUCores()
	if err != nil {
		t.Fatal(err)
	}
	if n == 0 {
		t.Fatal("cores should > 0")
	}
	t.Logf("detected %d logical cores", n)
}