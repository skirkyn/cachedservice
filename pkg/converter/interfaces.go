package converter


type GenericConverter interface {
	FromMap(map[string]interface{}) (interface{}, error)
	FromBinaryJson([]byte) (interface{}, error)
	FromMapSlice([]map[string]interface{}) ([]interface{}, error)
	FromBinaryJsonSlice([]interface{}) ([]interface{}, error)
}


