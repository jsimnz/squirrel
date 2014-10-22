package squirrel

import (
	"fmt"
)

type distinctOnPart part

func newDistinctOnPart(column string) Sqlizer {
	return &distinctOnPart{pred: column}
}

func (p distinctOnPart) ToSql() (sql string, args []interface{}, err error) {
	switch pred := p.pred.(type) {
	case nil:
		// no-op
	case string:
		sql = pred
		args = p.args
	default:
		err = fmt.Errorf("expected string, not %T", pred)
	}
	return
}
