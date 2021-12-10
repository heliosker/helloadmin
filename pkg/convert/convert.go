package convert

import "strconv"

type StrTo string

func (s StrTo) String() string {
	return string(s)
}

func (s StrTo) Int() (int, error) {
	v, e := strconv.Atoi(s.String())
	return v, e
}

func (s StrTo) MustInt() int {
	v, _ := s.Int()
	return v
}

func (s StrTo) UInt32() (uint32, error) {
	v, e := strconv.Atoi(s.String())
	return uint32(v), e
}

func (s StrTo) UInt() (uint, error) {
	v, e := strconv.Atoi(s.String())
	return uint(v), e
}

func (s StrTo) MustUInt() uint {
	v, _ := s.UInt()
	return v
}

func (s StrTo) MustUInt32() uint32 {
	v, _ := s.UInt32()
	return v
}
