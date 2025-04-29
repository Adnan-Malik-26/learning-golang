package hello

import (
	"testing"
)

func TestSayHello(t *testing.T){
	subtests := []struct{
		items []string
		result string
	}{
		{
			result: "Hello, world!",
		},
		{
			items: []string{"Adnan"},
			result: "Hello, Adnan!",
		},
		{
			items: []string{"Adnan", "Riya"},
			result: "Hello, Adnan, Riya!",
		},

	}

	for _, st := range subtests{
		if s := Say(st.items); s != st.result {
			t.Errorf("wanted %s (%v) and got %s", st.result, st.items , s)
		}
	}
}
