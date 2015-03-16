package rtda

import (
	rtc "github.com/zxh0/jvm.go/jvmgo/jvm/rtda/class"
	"testing"
)

func TestGetSet(t *testing.T) {
	vars := newLocalVars(6)
	vars.SetRef(0, nil)
	vars.SetRef(1, &rtc.Obj{})
	vars.SetInt(2, -37)
	vars.SetLong(3, 0xabcd1234ff)
	vars.SetFloat(4, 3.14)
	vars.SetDouble(5, -2.71828)
	//vars.SetInt(6, 0)

	if x := vars.GetRef(0); x != nil {
		t.Errorf("not nil: %v", x)
	}
	if x := vars.GetRef(1); x == nil {
		t.Errorf("nil!")
	}
	if x := vars.GetInt(2); x != -37 {
		t.Errorf("int:%v", x)
	}
	if x := vars.GetLong(3); x != 0xabcd1234ff {
		t.Errorf("long:%v", x)
	}
	if x := vars.GetFloat(4); x != 3.14 {
		t.Errorf("float:%v", x)
	}
	if x := vars.GetDouble(5); x != -2.71828 {
		t.Errorf("double:%v", x)
	}
}
