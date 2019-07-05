package ids

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
)

var _ driver.Valuer = (*ID)(nil)
var _ sql.Scanner = (*ID)(nil)

func (id *ID) Value() (driver.Value, error) {
	return id.Bytes(), nil
}

func (id *ID) Scan(src interface{}) error {
	data, ok := src.([]byte)
	if !ok {
		return errors.New("ids: unsupported type")
	}

	newID := (*ID)(&data)
	if err := newID.Valid(); err != nil {
		return fmt.Errorf("ids: %s", err)
	}

	*id = *newID
	return nil
}
