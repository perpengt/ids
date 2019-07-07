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
	if src == nil {
		*id = ID{}
		return nil
	}

	data, ok := src.([]byte)
	if !ok {
		return errors.New("ids: unsupported type")
	}

	if len(data) != 8 {
		return errors.New("ids: invalid id format")
	}

	// Create new ID
	newID := &ID{data[0], data[1], data[2], data[3], data[4], data[5], data[6], data[7]}
	if err := newID.Valid(); err != nil {
		return fmt.Errorf("ids: %s", err)
	}

	*id = *newID
	return nil
}
