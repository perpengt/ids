package ids

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func mockRowsToSqlRows(mockRows *sqlmock.Rows) *sql.Rows {
	randomKey := GenerateID(0).URIString()

	db, mock, _ := sqlmock.New()
	mock.ExpectQuery("select" + randomKey).WillReturnRows(mockRows)
	rows, _ := db.Query("select" + randomKey)
	return rows
}

func TestID_Scan(t *testing.T) {
	var (
		testID     = GenerateID(1)
		testResult ID
	)

	mockRows := sqlmock.NewRows([]string{"id"}).
		AddRow(testID.Bytes())

	rows := mockRowsToSqlRows(mockRows)
	rows.Next()

	err := rows.Scan(&testResult)
	if err != nil {
		t.Errorf("Scan failed: %s", err)
		return
	}

	if !Equal(testID, testResult) {
		t.Errorf("The scanned ID is different from the original ID.")
	}
}
