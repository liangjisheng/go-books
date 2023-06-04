package scylladb__test

import (
	"fmt"
	"github.com/scylladb/go-set/iset"
	"github.com/scylladb/go-set/strset"
	"testing"
)

func TestStringSet(t *testing.T) {
	s1 := strset.New("1", "2", "2")
	s2 := strset.New("2", "3", "2")
	fmt.Println("s1", s1)
	fmt.Println("s2", s2)

	fmt.Println("intersection", strset.Intersection(s1, s2))
	fmt.Println("union", strset.Union(s1, s2))
	fmt.Println("difference", strset.Difference(s1, s2))
	fmt.Println("SymmetricDifference", strset.SymmetricDifference(s1, s2))
}

func TestIntSet(t *testing.T) {
	i1 := iset.New(1, 2, 2)
	i2 := iset.New(2, 3, 2)
	fmt.Println("i1", i1)
	fmt.Println("i2", i2)

	fmt.Println("intersection", iset.Intersection(i1, i2))
	fmt.Println("union", iset.Union(i1, i2))
	fmt.Println("difference", iset.Difference(i1, i2))
	fmt.Println("SymmetricDifference", iset.SymmetricDifference(i1, i2))
}
