package forms

type FilterClause struct {
	Column       string `json:"column"`
	Op           string `json:"op"`
	UntypedValue string `json:"untypedValue"`
}

func NewFilterLike(column string, value string) FilterClause {
	return FilterClause{
		Column:       column,
		Op:           "Like",
		UntypedValue: value,
	}
}
