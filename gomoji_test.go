package gomoji_test

import (
	"iglou-eu/gomoji"
	"testing"
)

func TestGet(t *testing.T) {
	test := []struct {
		msg  string
		name string
		expt *gomoji.Emoji
	}{
		{"empty name", "", nil},
		{"not available, because it's too dangerous", "a proton pack", nil},
		{"is that you, slimer?", "ghost", &gomoji.List[110]},
	}

	for i, v := range test {
		res := gomoji.Get(v.name)
		if res != v.expt {
			t.Errorf(
				"Test No%d (%s) for the emoji '%s' has failed\nExpected => %#v\nResult => %#v",
				i,
				v.msg,
				v.name,
				v.expt,
				res,
			)
		}
	}
}

func TestInfo(t *testing.T) {
	test := []struct {
		msg   string
		emoji string
		expt  *gomoji.Emoji
	}{
		{"empty emoji", "", nil},
		{"not regular, what about the Twinkie?", "∹", nil},
		{"almost a twinkie", "🥮", &gomoji.List[752]},
	}

	for i, v := range test {
		res := gomoji.Info(v.emoji)
		if res != v.expt {
			t.Errorf(
				"Test No%d (%s) for the emoji '%s' has failed\nExpected => %#v\nResult => %#v",
				i,
				v.msg,
				v.emoji,
				v.expt,
				res,
			)
		}
	}
}

func TestByKeyword(t *testing.T) {
	test := []struct {
		msg  string
		key  string
		size int
	}{
		{"empty key", "", 0},
		{"are you the keymaster? nop!", "Keymaster", 0},
		{"janine phone", "phone", 9},
	}

	for i, v := range test {
		res := gomoji.ByKeyword(v.key)
		if len(res) != v.size {
			t.Errorf(
				"Test No%d (%s) for the key '%s' has failed\nExpected size => %d | Gived size => %d\nResult => %#v",
				i,
				v.msg,
				v.key,
				v.size,
				len(res),
				res,
			)
		}
	}
}

func TestByCategory(t *testing.T) {
	test := []struct {
		msg  string
		cat  string
		size int
	}{
		{"empty cat", "", 0},
		{"not a cat! a dog ?", "schrodinger", 0},
		{"fun with flags", "Flags", 269},
	}

	for i, v := range test {
		res := gomoji.ByCategory(v.cat)
		if len(res) != v.size {
			t.Errorf(
				"Test No%d (%s) for the cat '%s' has failed\nExpected size => %d | Gived size => %d\nResult => %#v",
				i,
				v.msg,
				v.cat,
				v.size,
				len(res),
				res,
			)
		}
	}
}

func TestCategorysed(t *testing.T) {
	categorysed := gomoji.Categorysed()

	if len(categorysed) != int(gomoji.ListCategorySize) {
		t.Errorf(
			"Test Categorysed has failed\nExpected size => %d | Gived size => %d\nResult => %#v",
			gomoji.ListCategorySize,
			len(categorysed),
			categorysed,
		)
	}
}
