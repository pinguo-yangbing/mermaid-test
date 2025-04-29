package dao

import (
	"github.com/pinguo-icc/kratos-library/v2/base"
	pmongo "github.com/pinguo-icc/kratos-library/v2/mongo"
	"go.mongodb.org/mongo-driver/mongo"
)

type ExampleType string

const (
	ExampleTypeManual ExampleType = "manual"
	ExampleTypeAuto   ExampleType = "auto"
)

type Example struct {
	base.BaseDO `bson:"inline"`
	Name        string          `bson:"name"`
	Age         int             `bson:"age"`
	Address     *ExampleAddress `bson:"address"`
	Type        ExampleType     `bson:"type"`
}

type ExampleAddress struct {
	Province string `bson:"province"`
	City     string `bson:"city"`
	District string `bson:"district"`
}

type ExampleDAO struct {
	*base.BaseDAOImpl[Example]
}

// NewExampleDAO 创建ExampleDAO对象
func NewExampleDAO(db *mongo.Database) (*ExampleDAO, error) {
	ud := &ExampleDAO{
		BaseDAOImpl: &base.BaseDAOImpl[Example]{
			Coll: pmongo.NewCollectionSelector(db, "example"),
		},
	}

	return ud, nil
}
