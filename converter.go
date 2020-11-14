package cachedservice

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type Converter struct {
	entityProvider      func() interface{}
	entitySliceProvider func() []interface{}
}

func NewConverter(entityProvider func() interface{}) *Converter {
	return &Converter{entityProvider: entityProvider}
}

func (cnv Converter) FromMap(theMap map[string]interface{}) (interface{}, error) {
	entity := cnv.entityProvider()
	err := mapstructure.Decode(theMap, &entity)
	return entity, err

}
func (cnv Converter) FromBinaryJson(str []byte) (interface{}, error) {
	theMap := make(map[string]interface{}, 0)
	err := json.Unmarshal(str, &theMap)
	if err == nil {
		return cnv.FromMap(theMap)

	} else {
		return nil, err
	}
}
func (cnv Converter) FromMapSlice(theSlice []map[string]interface{}) ([]interface{}, error) {
	res := make([]interface{}, 0)
	errSlice := make([]error, 0)
	for _, r := range theSlice {
		current, err := cnv.FromMap(r)
		if err == nil {
			res = append(res, current)
		} else {
			errSlice = append(errSlice, err)
		}
	}
	if len(errSlice) == 0 {
		return res, nil
	}
	return res, fmt.Errorf("errors %+q\n", errSlice)
}
func (cnv Converter) FromBinaryJsonSlice(theSlice []interface{}) ([]interface{}, error) {
	res := make([]interface{}, 0)
	errSlice := make([]error, 0)

	for _, r := range theSlice {
		if r != nil {
			current, err := cnv.FromBinaryJson([]byte(r.(string)))
			if err == nil {
				res = append(res, current)
			} else {
				errSlice = append(errSlice, err)
			}
		}

	}
	if len(errSlice) == 0 {
		return res, nil
	}
	return res, fmt.Errorf("errors %+q\n", errSlice)
}
