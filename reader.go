package accessdot

func Read(src interface{}, e string) interface{} {
	if src == nil {
		return nil
	}

	reader := NewExpressionReader(".[]", "EOF")
	list := NewParser(reader).Parse(e)
	return doRead(src, list)
}

func doRead(data interface{}, list []Access) interface{} {
	if len(list) == 0 {
		return data
	}

	switch data.(type) {
	case map[string]interface{}:
		return readMap(data.(map[string]interface{}), list)
	case []interface{}:
		return readArray(data.([]interface{}), list)
	default:
		return readValue(data, list)
	}
}

func readValue(data interface{}, list []Access) interface{} {
	if len(list) == 0 {
		return data
	}

	return nil
}

func readArray(arr []interface{}, list []Access) interface{} {
	if len(arr)-1 >= list[0].Index {
		return doRead(arr[list[0].Index], list[1:])
	}

	return nil
}

func readMap(dic map[string]interface{}, list []Access) interface{} {
	if len(list) > 0 && list[0].Type == "key" {
		if v, ok := dic[list[0].Key]; ok {
			return doRead(v, list[1:])
		}
	}

	return nil
}
