package mts

import "fmt"

//type void struct{}
//var member void

var conditions map[string]string
var boolConditions map[string]string

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
	boolConditions = map[string]string{
		"$ne": " IS NOT ",
	}
}

func generateCondition(op, key string, value interface{}) (cond *string, handleValue bool) {
	if b, ok := value.(bool); ok {
		c, exists := boolConditions[op]
		if exists {
			query := key
			query += c
			query += fmt.Sprintf("%t", b)
			return &query, false
		}
	}
	c, exists := conditions[op]
	if exists {
		query := key
		query += c + "?"
		return &query, true
	}
	return nil, false
}
