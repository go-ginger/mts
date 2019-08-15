package mts

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
