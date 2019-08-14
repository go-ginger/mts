package mts

import (
	"strings"
)

func iterate(data map[string]interface{}, temp *string) ([]string, *[]interface{}) {
	var queryItems []string
	var params []interface{}
	for k, v := range data {
		op := generateOperator(k)
		if op != nil {
			queryItems = append(queryItems, *op)
		} else if temp != nil {
			condition := generateCondition(k, *temp, v)
			if condition != nil {
				queryItems = append(queryItems, *condition)
			}
		} else {
			queryItems = append(queryItems, k+" = ?")
		}
		params = append(params, v)
	}
	return queryItems, &params
}

func Parse(query interface{}) (interface{}, *[]interface{}) {
	parts, params := iterate(query.(map[string]interface{}), nil)
	result := "(" + strings.Join(parts, ") AND (") + ")"
	return result, params
}
