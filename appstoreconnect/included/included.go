package included

import (
	"encoding/json"
	"fmt"

	"github.com/sse-open/go-app-store-connect/appstoreconnect/resource/inapppurchase"
	"github.com/sse-open/go-app-store-connect/appstoreconnect/resource/territories"
)

type ErrUnsupportedIncludedType struct {
	Type string
}

func (e ErrUnsupportedIncludedType) Error() string {
	return fmt.Sprintf("unsupported included type: %s", e.Type)
}

type Included struct {
	Type     string
	TypeData interface{}
}

type UnmarshallableIncludedType interface {
	Unmarshal(b []byte) (interface{}, error)
}

type GenericUnmarshallableIncludedType[T any] struct{}

func (guit GenericUnmarshallableIncludedType[T]) Unmarshal(b []byte) (interface{}, error) {
	var v T
	err := json.Unmarshal(b, &v)
	return v, err
}

var includedTypeToStructTypeMap = map[string]UnmarshallableIncludedType{
	"territories":              GenericUnmarshallableIncludedType[territories.Territory]{},
	"inAppPurchasePricePoints": GenericUnmarshallableIncludedType[inapppurchase.InAppPurchasePricePoint]{},
}

func UnmarshalInclude(b []byte) (string, interface{}, error) {
	var typeRef struct {
		Type string `json:"type"`
	}

	err := json.Unmarshal(b, &typeRef)
	if err != nil {
		return "", nil, err
	}

	typeName := typeRef.Type

	if structType, ok := includedTypeToStructTypeMap[typeName]; ok {
		typeData, err := structType.Unmarshal(b)
		if err != nil {
			return "", nil, err
		}
		return typeName, typeData, nil
	}

	return "", nil, ErrUnsupportedIncludedType{Type: typeName}
}

func (i *Included) UnmarshalJSON(b []byte) error {
	typeName, typeData, err := UnmarshalInclude(b)
	i.Type = typeName
	i.TypeData = typeData

	return err
}
