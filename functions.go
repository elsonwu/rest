package rest

import (
	"io"
	"labix.org/v2/mgo/bson"
)

type Data interface{}

type errItem struct {
	Num int    `json:"num"`
	Msg string `json:"msg"`
}

func NewErrs() *Errs {
	return &Errs{}
}

type Errs []errItem

func (self *Errs) AddErr(num int, msg string) *Errs {
	for _, e := range *self {
		if e.Num == num {
			return self
		}
	}

	(*self) = append(*self, errItem{num, msg})
	return self
}

func RawDecode(raw io.Reader, out interface{}) error {
	bts, err := bson.Marshal(raw)
	if err != nil {
		return err
	}

	if err := bson.Unmarshal(bts, out); err != nil {
		return err
	}

	return nil
}
