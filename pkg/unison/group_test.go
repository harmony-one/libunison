package unison

import (
	"bytes"
	"testing"
)

func TestNewGroupTag(t *testing.T) {
	successfulTest := func(t *testing.T, src interface{}, expected []byte) {
		actual := NewGroupTag(src)
		if bytes.Compare((*actual)[:], expected) != 0 {
			t.Errorf("NewGroupTag(%+v) should be %+v but is %+v",
				src, expected, actual)
		}
	}
    hashHarmony := []byte{
		0xcc, 0xfe, 0x55, 0xdf, 0xa0, 0xfe, 0x96, 0x39,
		0x10, 0xdc, 0x39, 0x49, 0xaf, 0x46, 0x23, 0x9e,
	}
	t.Run("HarmonyString", func(t *testing.T) {
		successfulTest(t, "Harmony", hashHarmony)
	})
    t.Run("HarmonyBytes", func(t *testing.T) {
		successfulTest(t, []byte("Harmony"), hashHarmony)
	})
    panickingTest := func(t *testing.T, src interface{}) {
    	var tag *GroupTag
    	defer func() {
    		if x := recover(); x == nil {
                t.Errorf("should have panicked but returned %+v", tag)
			} else {
				t.Logf("panicked as expected, with: %+v", x)
			}
		}()
    	tag = NewGroupTag(src)
	}
	t.Run("PanicOnInt", func(t *testing.T) {
		panickingTest(t, 3)
	})
    t.Run("PanicOnNil", func(t *testing.T) {
    	panickingTest(t, nil)
	})
}

func TestNewGroupID(t *testing.T) {
	successfulTest := func(t *testing.T, src interface{}, expected []byte) {
		id := NewGroupID(src)
		var actual *GroupTag
		switch v := id.(type) {
		case *GroupTag:
            actual = v
		default:
            t.Logf("returned group ID of an unexpected type: %#v", id)
			t.FailNow()
		}
		if bytes.Compare((*actual)[:], expected) != 0 {
			t.Errorf("NewGroupID(%+v) should be %+v but is %+v",
				src, expected, actual)
		}
	}
	hashHarmony := []byte{
		0xcc, 0xfe, 0x55, 0xdf, 0xa0, 0xfe, 0x96, 0x39,
		0x10, 0xdc, 0x39, 0x49, 0xaf, 0x46, 0x23, 0x9e,
	}
	t.Run("HarmonyString", func(t *testing.T) {
		successfulTest(t, "Harmony", hashHarmony)
	})
	t.Run("HarmonyBytes", func(t *testing.T) {
		successfulTest(t, []byte("Harmony"), hashHarmony)
	})
	panickingTest := func(t *testing.T, src interface{}) {
		var id GroupID
		defer func() {
			if x := recover(); x == nil {
				t.Errorf("should have panicked but returned %+v", id)
			} else {
				t.Logf("panicked as expected, with: %+v", x)
			}
		}()
		id = NewGroupTag(src)
	}
	t.Run("PanicOnInt", func(t *testing.T) {
		panickingTest(t, 3)
	})
	t.Run("PanicOnNil", func(t *testing.T) {
		panickingTest(t, nil)
	})
}