package rest

import (
	"io"

	"labix.org/v2/mgo/bson"
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
