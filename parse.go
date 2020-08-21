package mts

import (
	"github.com/go-ginger/models"
	"strings"
)

func iterate(data map[string]interface{}, temp *string) ([]string, []interface{}) {
	var queryItems []string
	var params []interface{}
	for k, v := range data {
		op := generateOperator(k)
		if op != nil {
			d2, success := v.([]interface{})
			if success {
				var opQueryItems []string
				for _, d := range d2 {
					d3, success := d.(map[string]interface{})
					if success {
						q, p := iterate(d3, nil)
						params = append(params, p...)
						opQueryItems = append(opQueryItems, q...)
					}
				}
				query := "(" + strings.Join(opQueryItems, ") "+*op+" (") + ")"
				queryItems = append(queryItems, query)
			}
			//queryItems = append(queryItems, *op)
		} else {
			var condition *string
			if temp != nil {
				condition = generateCondition(k, *temp, nil)
				if condition != nil {
					queryItems = append(queryItems, *condition)
					params = append(params, v)
				}
			}
			if condition == nil {
				if iv, ok := v.(map[string]interface{}); ok {
					q, p := iterate(iv, &k)
					params = append(params, p...)
					queryItems = append(queryItems, q...)
				} else if iv, ok := v.(models.Filters); ok {
					q, p := iterate(iv, &k)
					params = append(params, p...)
					queryItems = append(queryItems, q...)

				} else {
					if v != nil {
						queryItems = append(queryItems, k+"=?")
						params = append(params, v)
					} else {
						queryItems = append(queryItems, k+" IS NULL")
					}
				}
			}
		}
	}
	if queryItems != nil {
		queryItems = []string{"(" + strings.Join(queryItems, ") AND (") + ")"}
	}
	return queryItems, params
}

func Parse(query interface{}) (interface{}, []interface{}) {
	parts, params := iterate(query.(models.Filters), nil)
	var result string
	if parts != nil {
		result = "(" + strings.Join(parts, ") AND (") + ")"
	}
	return result, params
}
