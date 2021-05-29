package accessdot_test

import (
	"encoding/json"
	"testing"

	"github.com/rodriez/accessdot"
	"github.com/stretchr/testify/assert"
)

func TestRead_Object_Access_Number(t *testing.T) {
	expression := "call.agent.money"
	text := `{
		"call": {
			"agent":{
				"id":1,
				"money": 323.5
			}
		}
	}`

	var source map[string]interface{}

	json.Unmarshal([]byte(text), &source)

	result := accessdot.Read(source, expression)

	assert.Equal(t, float64(323.5), result)
}

func TestRead_Object_Access_String(t *testing.T) {
	expression := "call.agent.id"
	text := `{
		"call": {
			"agent":{
				"id":"as23df"
			}
		}
	}`

	var source map[string]interface{}

	json.Unmarshal([]byte(text), &source)

	result := accessdot.Read(source, expression)

	assert.Equal(t, "as23df", result)
}

func TestRead_Access_Object_Return_Array(t *testing.T) {
	expression := "call.scoring"
	text := `{
		"call": {
			"scoring":[1, 2, 3]
		}
	}`

	var source map[string]interface{}

	json.Unmarshal([]byte(text), &source)

	result := accessdot.Read(source, expression)

	assert.Equal(t, float64(1), result.([]interface{})[0])
	assert.Equal(t, float64(2), result.([]interface{})[1])
	assert.Equal(t, float64(3), result.([]interface{})[2])
}

func TestRead_Array_Access(t *testing.T) {
	expression := "call.agents[2].id"
	text := `{
		"call": {
			"agents":[
				{
					"id":25
				},
				{
					"id":26
				},
				{
					"id":30
				}
			]
		}
	}`

	var source map[string]interface{}

	json.Unmarshal([]byte(text), &source)

	result := accessdot.Read(source, expression)

	assert.Equal(t, float64(30), result)
}

func TestRead_Array_Access_Not_Existing_Index(t *testing.T) {
	expression := "call.agents[6].id"
	text := `{
		"call": {
			"agents":[
				{
					"id":25
				},
				{
					"id":26
				},
				{
					"id":30
				}
			]
		}
	}`

	var source map[string]interface{}

	json.Unmarshal([]byte(text), &source)

	result := accessdot.Read(source, expression)

	assert.Equal(t, nil, result)
}

func TestRead_Object_Access_Date(t *testing.T) {
	expression := "call.agent.date"
	text := `{
		"call": {
			"agent":{
				"date":"2021-06-26T00:12:22.384Z"
			}
		}
	}`

	var source map[string]interface{}

	json.Unmarshal([]byte(text), &source)

	result := accessdot.Read(source, expression)

	assert.Equal(t, "2021-06-26T00:12:22.384Z", result)
}

func TestRead_Object_Access_Not_Existing_Middle_Key(t *testing.T) {
	expression := "call.agent.roles.id"
	text := `{
		"call": {
			"agent":{
				"role": "admin"
			}
		}
	}`

	var source map[string]interface{}

	json.Unmarshal([]byte(text), &source)

	result := accessdot.Read(source, expression)

	assert.Equal(t, nil, result)
}

func TestRead_Object_Access_Invalid_Key(t *testing.T) {
	expression := "call.agent.profile.roles[0].id"
	text := `{
		"call": {
			"profile":[
				"admin"
			]
		}
	}`

	var source map[string]interface{}

	json.Unmarshal([]byte(text), &source)

	result := accessdot.Read(source, expression)

	assert.Equal(t, nil, result)
}

func TestRead_Object_Access_Not_Existing_End_Key(t *testing.T) {
	expression := "call.agent.role.id"
	text := `{
		"call": {
			"agent":{
				"role": "admin"
			}
		}
	}`

	var source map[string]interface{}

	json.Unmarshal([]byte(text), &source)

	result := accessdot.Read(source, expression)

	assert.Equal(t, nil, result)
}

func TestRead_Object_Access_Start_Array(t *testing.T) {
	expression := "[0].date"
	text := `[
		{
			"date":"2021-06-26T00:12:22.384Z"
		},
		{
			"date":"2021-06-26T00:12:22.384Z"
		}
	]`

	var source []interface{}

	json.Unmarshal([]byte(text), &source)

	result := accessdot.Read(source, expression)

	assert.Equal(t, "2021-06-26T00:12:22.384Z", result)
}

func TestRead_Object_Access_Start_Array_Loop(t *testing.T) {
	expression := "[0][0][0].date"
	text := `[[[
		{
			"date":"2021-06-26T00:12:22.384Z"
		},
		{
			"date":"2021-06-26T00:12:22.384Z"
		}
	]]]`

	var source []interface{}

	json.Unmarshal([]byte(text), &source)

	result := accessdot.Read(source, expression)

	assert.Equal(t, "2021-06-26T00:12:22.384Z", result)
}

func TestRead_Object_Access_Nil_Source(t *testing.T) {
	expression := "[0][0][0].date"

	result := accessdot.Read(nil, expression)

	assert.Equal(t, nil, result)
}

func TestRead_Object_Access_String_Index(t *testing.T) {
	expression := "call.agent.[date]"
	text := `{
		"call": {
			"agent":{
				"date":"2021-06-26T00:12:22.384Z"
			}
		}
	}`

	var source map[string]interface{}

	json.Unmarshal([]byte(text), &source)

	result := accessdot.Read(source, expression)

	assert.Equal(t, nil, result)
}
