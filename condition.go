package mts

//type void struct{}
//var member void

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

func generateCondition(op, key string, value interface{}) *string {
	c, exists := conditions[op]
	if exists {
		query := key
		query += c
		if value == nil {
			value = "?"
		}
		query += value.(string)
		return &query
	}
	return nil
}