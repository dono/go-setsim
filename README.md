# go-setsim
Similarity coefficients of sets (Jaccard, Dice, Simpson)

## Usage

```
func TestJaccard(t *testing.T) {
	s1 := mapset.NewSetFromSlice([]interface{}{"a", "b", "c", "d", "e"})
	s2 := mapset.NewSetFromSlice([]interface{}{"a", "b", "c", "g"})

	j := Jaccard(s1, s2)
	if j != 0.5 {
		t.Fatal("ng")
	}
}

func TestDice(t *testing.T) {
	s1 := mapset.NewSetFromSlice([]interface{}{"a", "b", "c", "d", "e", "f"})
	s2 := mapset.NewSetFromSlice([]interface{}{"a", "b", "c", "g"})

	d := Dice(s1, s2)
	if d != 0.6 {
		t.Fatal("ng")
	}
}

func TestSimpson(t *testing.T) {
	s1 := mapset.NewSetFromSlice([]interface{}{"a", "b", "c", "d", "e", "f"})
	s2 := mapset.NewSetFromSlice([]interface{}{"a", "b", "c", "g"})

	s := Simpson(s1, s2)
	if s != 0.75 {
		t.Fatal("ng")
	}
}
```