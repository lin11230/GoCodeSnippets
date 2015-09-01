package gotest

import "testing"

func Test_Division_1(t *testing.T) {
	if i, e := Division(6, 2); i != 3 || e != nil {
		//try a unit test on fucntion
		t.Error("除法函數測試沒通過")
	} else {
		t.Log("第一個測試成功了")
	}
}

func Test_Division_2(t *testing.T) {
	if _, e := Division(6, 0); e == nil {
		//try a unit test on function
		t.Error("Division did not work as expected.")
	} else {
		t.Log("one test passed.", e)
	}
}
