package mts

import (
	"strings"
)

func generateCondition(op, key string, value interface{}) *string {
	c, exists := conditions[op]
	if exists {
		query := key
		query += c
		query += "?"
		return &query
	}
	return nil
}

func iterate(data map[string]interface{}, temp *string) ([]string, []interface{}) {
	var queryItems []string
	var params []interface{}
	for k, v := range data {
		op := generateOperator(k)
		if op != nil {
			queryItems = append(queryItems, *op)
		} else if temp != nil {
			condition := generateRawCondition(k, *temp, v)
			if condition != nil {
				queryItems = append(queryItems, *condition)
			}
		}
		d2, success := v.(map[string]interface{})
		if success {
			iterate(d2, &k)
		} else {
			queryItems = append(queryItems, k+"=?")
			params = append(params, v)
		}
	}
	return queryItems, params
}

func Parse(query interface{}) (interface{}, []interface{}) {
	parts, params := iterate(query.(map[string]interface{}), nil)
	result := "(" + strings.Join(parts, ") AND (") + ")"
	return result, params
}
