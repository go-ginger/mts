package mts

import (
	"fmt"
	"strings"
)

//type void struct{}
//var member void

var operators map[string]string
var conditions map[string]string

func init() {
	operators = map[string]string{
		"$and": " AND ",
		"$or":  " OR ",
	}
	conditions = map[string]string{
		"$lt":  " < ",
		"$lte": " <= ",
		"$gt":  " > ",
		"$gte": " >= ",
		"$ne":  " <> ",
	}
}

func generateOperator(op string) *string {
	c, exists := operators[op]
	if exists {
		query := c
		return &query
	}
	return nil
}

func generateCondition(op, key string, value interface{}) *string {
	c, exists := conditions[op]
	if exists {
		query := key
		query += c
		query += value.(string)
		return &query
	}
	return nil
}

func iterate(data map[string]interface{}, temp *string) []string {
	var queryItems []string
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

func Parse(query interface{}) interface{} {
	parts := iterate(query.(map[string]interface{}), nil)
	result := "(" + strings.Join(parts, ") AND (") + ")"
	return result
}
