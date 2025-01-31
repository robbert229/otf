// Code generated by pggen. DO NOT EDIT.

package pggen

import (
	"context"
	"fmt"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
)

const insertVariableSQL = `INSERT INTO variables (
    variable_id,
    key,
    value,
    description,
    category,
    sensitive,
    hcl,
    version_id
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
);`

type InsertVariableParams struct {
	VariableID  pgtype.Text
	Key         pgtype.Text
	Value       pgtype.Text
	Description pgtype.Text
	Category    pgtype.Text
	Sensitive   bool
	HCL         bool
	VersionID   pgtype.Text
}

// InsertVariable implements Querier.InsertVariable.
func (q *DBQuerier) InsertVariable(ctx context.Context, params InsertVariableParams) (pgconn.CommandTag, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "InsertVariable")
	cmdTag, err := q.conn.Exec(ctx, insertVariableSQL, params.VariableID, params.Key, params.Value, params.Description, params.Category, params.Sensitive, params.HCL, params.VersionID)
	if err != nil {
		return cmdTag, fmt.Errorf("exec query InsertVariable: %w", err)
	}
	return cmdTag, err
}

// InsertVariableBatch implements Querier.InsertVariableBatch.
func (q *DBQuerier) InsertVariableBatch(batch genericBatch, params InsertVariableParams) {
	batch.Queue(insertVariableSQL, params.VariableID, params.Key, params.Value, params.Description, params.Category, params.Sensitive, params.HCL, params.VersionID)
}

// InsertVariableScan implements Querier.InsertVariableScan.
func (q *DBQuerier) InsertVariableScan(results pgx.BatchResults) (pgconn.CommandTag, error) {
	cmdTag, err := results.Exec()
	if err != nil {
		return cmdTag, fmt.Errorf("exec InsertVariableBatch: %w", err)
	}
	return cmdTag, err
}

const findVariableSQL = `SELECT *
FROM variables
WHERE variable_id = $1
;`

type FindVariableRow struct {
	VariableID  pgtype.Text `json:"variable_id"`
	Key         pgtype.Text `json:"key"`
	Value       pgtype.Text `json:"value"`
	Description pgtype.Text `json:"description"`
	Category    pgtype.Text `json:"category"`
	Sensitive   bool        `json:"sensitive"`
	HCL         bool        `json:"hcl"`
	VersionID   pgtype.Text `json:"version_id"`
}

// FindVariable implements Querier.FindVariable.
func (q *DBQuerier) FindVariable(ctx context.Context, variableID pgtype.Text) (FindVariableRow, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "FindVariable")
	row := q.conn.QueryRow(ctx, findVariableSQL, variableID)
	var item FindVariableRow
	if err := row.Scan(&item.VariableID, &item.Key, &item.Value, &item.Description, &item.Category, &item.Sensitive, &item.HCL, &item.VersionID); err != nil {
		return item, fmt.Errorf("query FindVariable: %w", err)
	}
	return item, nil
}

// FindVariableBatch implements Querier.FindVariableBatch.
func (q *DBQuerier) FindVariableBatch(batch genericBatch, variableID pgtype.Text) {
	batch.Queue(findVariableSQL, variableID)
}

// FindVariableScan implements Querier.FindVariableScan.
func (q *DBQuerier) FindVariableScan(results pgx.BatchResults) (FindVariableRow, error) {
	row := results.QueryRow()
	var item FindVariableRow
	if err := row.Scan(&item.VariableID, &item.Key, &item.Value, &item.Description, &item.Category, &item.Sensitive, &item.HCL, &item.VersionID); err != nil {
		return item, fmt.Errorf("scan FindVariableBatch row: %w", err)
	}
	return item, nil
}

const updateVariableByIDSQL = `UPDATE variables
SET
    key = $1,
    value = $2,
    description = $3,
    category = $4,
    sensitive = $5,
    version_id = $6,
    hcl = $7
WHERE variable_id = $8
RETURNING variable_id
;`

type UpdateVariableByIDParams struct {
	Key         pgtype.Text
	Value       pgtype.Text
	Description pgtype.Text
	Category    pgtype.Text
	Sensitive   bool
	VersionID   pgtype.Text
	HCL         bool
	VariableID  pgtype.Text
}

// UpdateVariableByID implements Querier.UpdateVariableByID.
func (q *DBQuerier) UpdateVariableByID(ctx context.Context, params UpdateVariableByIDParams) (pgtype.Text, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "UpdateVariableByID")
	row := q.conn.QueryRow(ctx, updateVariableByIDSQL, params.Key, params.Value, params.Description, params.Category, params.Sensitive, params.VersionID, params.HCL, params.VariableID)
	var item pgtype.Text
	if err := row.Scan(&item); err != nil {
		return item, fmt.Errorf("query UpdateVariableByID: %w", err)
	}
	return item, nil
}

// UpdateVariableByIDBatch implements Querier.UpdateVariableByIDBatch.
func (q *DBQuerier) UpdateVariableByIDBatch(batch genericBatch, params UpdateVariableByIDParams) {
	batch.Queue(updateVariableByIDSQL, params.Key, params.Value, params.Description, params.Category, params.Sensitive, params.VersionID, params.HCL, params.VariableID)
}

// UpdateVariableByIDScan implements Querier.UpdateVariableByIDScan.
func (q *DBQuerier) UpdateVariableByIDScan(results pgx.BatchResults) (pgtype.Text, error) {
	row := results.QueryRow()
	var item pgtype.Text
	if err := row.Scan(&item); err != nil {
		return item, fmt.Errorf("scan UpdateVariableByIDBatch row: %w", err)
	}
	return item, nil
}

const deleteVariableByIDSQL = `DELETE
FROM variables
WHERE variable_id = $1
RETURNING *
;`

type DeleteVariableByIDRow struct {
	VariableID  pgtype.Text `json:"variable_id"`
	Key         pgtype.Text `json:"key"`
	Value       pgtype.Text `json:"value"`
	Description pgtype.Text `json:"description"`
	Category    pgtype.Text `json:"category"`
	Sensitive   bool        `json:"sensitive"`
	HCL         bool        `json:"hcl"`
	VersionID   pgtype.Text `json:"version_id"`
}

// DeleteVariableByID implements Querier.DeleteVariableByID.
func (q *DBQuerier) DeleteVariableByID(ctx context.Context, variableID pgtype.Text) (DeleteVariableByIDRow, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "DeleteVariableByID")
	row := q.conn.QueryRow(ctx, deleteVariableByIDSQL, variableID)
	var item DeleteVariableByIDRow
	if err := row.Scan(&item.VariableID, &item.Key, &item.Value, &item.Description, &item.Category, &item.Sensitive, &item.HCL, &item.VersionID); err != nil {
		return item, fmt.Errorf("query DeleteVariableByID: %w", err)
	}
	return item, nil
}

// DeleteVariableByIDBatch implements Querier.DeleteVariableByIDBatch.
func (q *DBQuerier) DeleteVariableByIDBatch(batch genericBatch, variableID pgtype.Text) {
	batch.Queue(deleteVariableByIDSQL, variableID)
}

// DeleteVariableByIDScan implements Querier.DeleteVariableByIDScan.
func (q *DBQuerier) DeleteVariableByIDScan(results pgx.BatchResults) (DeleteVariableByIDRow, error) {
	row := results.QueryRow()
	var item DeleteVariableByIDRow
	if err := row.Scan(&item.VariableID, &item.Key, &item.Value, &item.Description, &item.Category, &item.Sensitive, &item.HCL, &item.VersionID); err != nil {
		return item, fmt.Errorf("scan DeleteVariableByIDBatch row: %w", err)
	}
	return item, nil
}
