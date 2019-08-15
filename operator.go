package mts

var operators map[string]string

func init() {
	operators = map[string]string{
		"$and": " AND ",
		"$or":  " OR ",
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