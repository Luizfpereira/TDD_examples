package fizzbuzz_test

import (
	"testing"

	"github.com/Luizfpereira/TDD_examples/fizzbuzz"
	"github.com/stretchr/testify/suite"
)

type FizzbuzzSuite struct {
	suite.Suite
}

func (s *FizzbuzzSuite) TestOneUnchanged() {
	r := fizzbuzz.Run(1)
	s.Equal([]string{"1"}, r)
}

func (s *FizzbuzzSuite) TestFizz() {
	r := fizzbuzz.Run(3)
	s.Equal([]string{"Fizz"}, r)
}

func (s *FizzbuzzSuite) TestBuzz() {
	r := fizzbuzz.Run(5)
	s.Equal([]string{"Buzz"}, r)
}

func (s *FizzbuzzSuite) TestFizzBuzz() {
	r := fizzbuzz.Run(15)
	s.Equal([]string{"FizzBuzz"}, r)
}

func (s *FizzbuzzSuite) TestAnySlice() {
	r := fizzbuzz.Run(1, 2, 3, 4, 5, 6, 10, 13, 15, 25, 30)
	s.Equal([]string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "Buzz", "13", "FizzBuzz", "Buzz", "FizzBuzz"}, r)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(FizzbuzzSuite))
}
