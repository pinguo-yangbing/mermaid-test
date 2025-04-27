package dao

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	D  = bson.D
	E  = bson.E
	M  = bson.M
	A  = bson.A
	ID = primitive.ObjectID
)

func docID(v interface{}) D {
	return D{{Key: "_id", Value: v}}
}

func ele(key string, v interface{}) E {
	return E{Key: key, Value: v}
}

func regex(field string, value string) E {
	return E{Key: field, Value: primitive.Regex{Pattern: value, Options: "i"}}
}

type DatetimeDuration struct {
	Begin, End time.Time
}

func (d DatetimeDuration) cond(fieldName string) (E, bool) {
	cond := D{}
	if !d.Begin.IsZero() {
		cond = append(cond, ele("$gte", d.Begin))
	}
	if !d.End.IsZero() {
		cond = append(cond, ele("$lte", d.End))
	}

	if len(cond) == 0 {
		return E{}, false
	}

	return ele(fieldName, cond), true
}
