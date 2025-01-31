// Code generated by pggen. DO NOT EDIT.

package pggen

import (
	"context"
	"fmt"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
)

const upsertOrganizationTokenSQL = `INSERT INTO organization_tokens (
    organization_token_id,
    created_at,
    organization_name,
    expiry
) VALUES (
    $1,
    $2,
    $3,
    $4
) ON CONFLICT (organization_name) DO UPDATE
  SET created_at            = $2,
      organization_token_id = $1,
      expiry                = $4;`

type UpsertOrganizationTokenParams struct {
	OrganizationTokenID pgtype.Text
	CreatedAt           pgtype.Timestamptz
	OrganizationName    pgtype.Text
	Expiry              pgtype.Timestamptz
}

// UpsertOrganizationToken implements Querier.UpsertOrganizationToken.
func (q *DBQuerier) UpsertOrganizationToken(ctx context.Context, params UpsertOrganizationTokenParams) (pgconn.CommandTag, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "UpsertOrganizationToken")
	cmdTag, err := q.conn.Exec(ctx, upsertOrganizationTokenSQL, params.OrganizationTokenID, params.CreatedAt, params.OrganizationName, params.Expiry)
	if err != nil {
		return cmdTag, fmt.Errorf("exec query UpsertOrganizationToken: %w", err)
	}
	return cmdTag, err
}

// UpsertOrganizationTokenBatch implements Querier.UpsertOrganizationTokenBatch.
func (q *DBQuerier) UpsertOrganizationTokenBatch(batch genericBatch, params UpsertOrganizationTokenParams) {
	batch.Queue(upsertOrganizationTokenSQL, params.OrganizationTokenID, params.CreatedAt, params.OrganizationName, params.Expiry)
}

// UpsertOrganizationTokenScan implements Querier.UpsertOrganizationTokenScan.
func (q *DBQuerier) UpsertOrganizationTokenScan(results pgx.BatchResults) (pgconn.CommandTag, error) {
	cmdTag, err := results.Exec()
	if err != nil {
		return cmdTag, fmt.Errorf("exec UpsertOrganizationTokenBatch: %w", err)
	}
	return cmdTag, err
}

const findOrganizationTokensByNameSQL = `SELECT *
FROM organization_tokens
WHERE organization_name = $1;`

type FindOrganizationTokensByNameRow struct {
	OrganizationTokenID pgtype.Text        `json:"organization_token_id"`
	CreatedAt           pgtype.Timestamptz `json:"created_at"`
	OrganizationName    pgtype.Text        `json:"organization_name"`
	Expiry              pgtype.Timestamptz `json:"expiry"`
}

// FindOrganizationTokensByName implements Querier.FindOrganizationTokensByName.
func (q *DBQuerier) FindOrganizationTokensByName(ctx context.Context, organizationName pgtype.Text) ([]FindOrganizationTokensByNameRow, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "FindOrganizationTokensByName")
	rows, err := q.conn.Query(ctx, findOrganizationTokensByNameSQL, organizationName)
	if err != nil {
		return nil, fmt.Errorf("query FindOrganizationTokensByName: %w", err)
	}
	defer rows.Close()
	items := []FindOrganizationTokensByNameRow{}
	for rows.Next() {
		var item FindOrganizationTokensByNameRow
		if err := rows.Scan(&item.OrganizationTokenID, &item.CreatedAt, &item.OrganizationName, &item.Expiry); err != nil {
			return nil, fmt.Errorf("scan FindOrganizationTokensByName row: %w", err)
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("close FindOrganizationTokensByName rows: %w", err)
	}
	return items, err
}

// FindOrganizationTokensByNameBatch implements Querier.FindOrganizationTokensByNameBatch.
func (q *DBQuerier) FindOrganizationTokensByNameBatch(batch genericBatch, organizationName pgtype.Text) {
	batch.Queue(findOrganizationTokensByNameSQL, organizationName)
}

// FindOrganizationTokensByNameScan implements Querier.FindOrganizationTokensByNameScan.
func (q *DBQuerier) FindOrganizationTokensByNameScan(results pgx.BatchResults) ([]FindOrganizationTokensByNameRow, error) {
	rows, err := results.Query()
	if err != nil {
		return nil, fmt.Errorf("query FindOrganizationTokensByNameBatch: %w", err)
	}
	defer rows.Close()
	items := []FindOrganizationTokensByNameRow{}
	for rows.Next() {
		var item FindOrganizationTokensByNameRow
		if err := rows.Scan(&item.OrganizationTokenID, &item.CreatedAt, &item.OrganizationName, &item.Expiry); err != nil {
			return nil, fmt.Errorf("scan FindOrganizationTokensByNameBatch row: %w", err)
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("close FindOrganizationTokensByNameBatch rows: %w", err)
	}
	return items, err
}

const findOrganizationTokensByIDSQL = `SELECT *
FROM organization_tokens
WHERE organization_token_id = $1;`

type FindOrganizationTokensByIDRow struct {
	OrganizationTokenID pgtype.Text        `json:"organization_token_id"`
	CreatedAt           pgtype.Timestamptz `json:"created_at"`
	OrganizationName    pgtype.Text        `json:"organization_name"`
	Expiry              pgtype.Timestamptz `json:"expiry"`
}

// FindOrganizationTokensByID implements Querier.FindOrganizationTokensByID.
func (q *DBQuerier) FindOrganizationTokensByID(ctx context.Context, organizationTokenID pgtype.Text) (FindOrganizationTokensByIDRow, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "FindOrganizationTokensByID")
	row := q.conn.QueryRow(ctx, findOrganizationTokensByIDSQL, organizationTokenID)
	var item FindOrganizationTokensByIDRow
	if err := row.Scan(&item.OrganizationTokenID, &item.CreatedAt, &item.OrganizationName, &item.Expiry); err != nil {
		return item, fmt.Errorf("query FindOrganizationTokensByID: %w", err)
	}
	return item, nil
}

// FindOrganizationTokensByIDBatch implements Querier.FindOrganizationTokensByIDBatch.
func (q *DBQuerier) FindOrganizationTokensByIDBatch(batch genericBatch, organizationTokenID pgtype.Text) {
	batch.Queue(findOrganizationTokensByIDSQL, organizationTokenID)
}

// FindOrganizationTokensByIDScan implements Querier.FindOrganizationTokensByIDScan.
func (q *DBQuerier) FindOrganizationTokensByIDScan(results pgx.BatchResults) (FindOrganizationTokensByIDRow, error) {
	row := results.QueryRow()
	var item FindOrganizationTokensByIDRow
	if err := row.Scan(&item.OrganizationTokenID, &item.CreatedAt, &item.OrganizationName, &item.Expiry); err != nil {
		return item, fmt.Errorf("scan FindOrganizationTokensByIDBatch row: %w", err)
	}
	return item, nil
}

const deleteOrganiationTokenByNameSQL = `DELETE
FROM organization_tokens
WHERE organization_name = $1
RETURNING organization_token_id;`

// DeleteOrganiationTokenByName implements Querier.DeleteOrganiationTokenByName.
func (q *DBQuerier) DeleteOrganiationTokenByName(ctx context.Context, organizationName pgtype.Text) (pgtype.Text, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "DeleteOrganiationTokenByName")
	row := q.conn.QueryRow(ctx, deleteOrganiationTokenByNameSQL, organizationName)
	var item pgtype.Text
	if err := row.Scan(&item); err != nil {
		return item, fmt.Errorf("query DeleteOrganiationTokenByName: %w", err)
	}
	return item, nil
}

// DeleteOrganiationTokenByNameBatch implements Querier.DeleteOrganiationTokenByNameBatch.
func (q *DBQuerier) DeleteOrganiationTokenByNameBatch(batch genericBatch, organizationName pgtype.Text) {
	batch.Queue(deleteOrganiationTokenByNameSQL, organizationName)
}

// DeleteOrganiationTokenByNameScan implements Querier.DeleteOrganiationTokenByNameScan.
func (q *DBQuerier) DeleteOrganiationTokenByNameScan(results pgx.BatchResults) (pgtype.Text, error) {
	row := results.QueryRow()
	var item pgtype.Text
	if err := row.Scan(&item); err != nil {
		return item, fmt.Errorf("scan DeleteOrganiationTokenByNameBatch row: %w", err)
	}
	return item, nil
}
