package setsim

import (
	mapset "github.com/deckarep/golang-set"
)

func Jaccard(a, b mapset.Set) float64 {
	return float64(a.Intersect(b).Cardinality()) / float64(a.Union(b).Cardinality())
}

func Dice(a, b mapset.Set) float64 {
	return 2.0 * float64(a.Intersect(b).Cardinality()) / (float64(a.Cardinality() + b.Cardinality()))
}

func Simpson(a, b mapset.Set) float64 {
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	return float64(a.Intersect(b).Cardinality()) / float64(min(a.Cardinality(), b.Cardinality()))
}