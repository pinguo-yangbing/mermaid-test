package server

import (
	"fmt"
	"net/http"
	"reflect"

	khttp "github.com/go-kratos/kratos/v2/transport/http"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var jsonMarshalOptions = protojson.MarshalOptions{
	EmitUnpopulated: true,
	UseEnumNumbers:  false,
}

func invoke(svc interface{}) khttp.HandlerFunc {
	rv := reflect.ValueOf(svc)
	if reflect.Indirect(rv).Kind() == reflect.Interface {
		rv = rv.Elem()
	}

	return func(ctx khttp.Context) error {
		name := ctx.Vars().Get("name")
		rm := rv.MethodByName(name)
		if !rm.IsValid() {
			return ctx.String(http.StatusBadRequest, fmt.Sprintf("Invalid name %q", name))
		}

		in := reflect.New(rm.Type().In(1).Elem()).Interface()
		if err := ctx.Bind(in); err != nil {
			return err
		}

		res := rm.Call([]reflect.Value{
			reflect.ValueOf(ctx),
			reflect.ValueOf(in),
		})

		if err := res[1].Interface(); err != nil {
			return ctx.String(http.StatusInternalServerError, fmt.Sprintf("%v", err))
		}

		data, err := jsonMarshalOptions.Marshal(res[0].Interface().(protoreflect.ProtoMessage))
		if err != nil {
			return ctx.String(http.StatusInternalServerError, "JSON 格式化响应数据出错")
		}

		return ctx.Blob(http.StatusOK, "application/json", data)
	}
}
