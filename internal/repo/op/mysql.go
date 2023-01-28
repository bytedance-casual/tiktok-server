package op


type Condition struct {
	Filed string
	Opr   string
	Value any
}

// Eq 等于
func Eq(filed string, value any) *Condition {
	return &Condition{Filed: filed, Opr: "=", Value: value}
}

// Lt 小于
func Lt(filed string, value any) *Condition {
	return &Condition{Filed: filed, Opr: "<", Value: value}
}

// Le 小于等于
func Le(filed string, value any) *Condition {
	return &Condition{Filed: filed, Opr: "<=", Value: value}
}

// Gt 大于
func Gt(filed string, value any) *Condition {
	return &Condition{Filed: filed, Opr: ">", Value: value}
}

// Ge 大于等于
func Ge(filed string, value any) *Condition {
	return &Condition{Filed: filed, Opr: "=", Value: value}
}

// Ne 不等于
func Ne(filed string, value any) *Condition {
	return &Condition{Filed: filed, Opr: "<>", Value: value}
}

// Like 模糊查询
func Like(filed string, value string) *Condition {
	return &Condition{Filed: filed, Opr: "like", Value: "%" + value + "%"}
}

var NotDeleted = &Condition{
	Filed: "delete_time",
	Opr:   "=",
	Value: "0",
}
