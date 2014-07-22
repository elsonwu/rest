package rest

import (
	"io"

	"gopkg.in/mgo.v2/bson"
)

type Data interface{}

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
