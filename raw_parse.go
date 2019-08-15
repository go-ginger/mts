package mts

import (
	"fmt"
	"strings"
)



func generateRawCondition(op, key string, value interface{}) *string {
	c, exists := conditions[op]
	if exists {
		query := key
		query += c
		query += value.(string)
		return &query
	}
	return nil
}

func iterateRaw(data map[string]interface{}, temp *string) []string {
	var queryItems []string
	for k, v := range data {
		op := generateOperator(k)
		if op != nil {
			queryItems = append(queryItems, *op)
		} else if temp != nil {
			condition := generateRawCondition(k, *temp, v)
			if condition != nil {
				queryItems = append(queryItems, *condition)
			}
		} else {
			switch v.(type) {
			case string:
				queryItems = append(queryItems, k+" = '"+v.(string) + "'")
			default:
				queryItems = append(queryItems, k+" = "+fmt.Sprintf("%v", v))
			}
		}
	}
	return queryItems
}

func RawParse(query interface{}) interface{} {
	parts := iterateRaw(query.(map[string]interface{}), nil)
	result := "(" + strings.Join(parts, ") AND (") + ")"
	return result
}
