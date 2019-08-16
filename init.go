package mts

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
