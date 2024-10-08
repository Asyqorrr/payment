// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: entity_type.sql

package db

import (
	"context"
)

const createEntity = `-- name: CreateEntity :one
INSERT INTO entity_type(
	 entity_name)
	VALUES ($1) RETURNING entity_id, entity_name
`

func (q *Queries) CreateEntity(ctx context.Context, entityName string) (*EntityType, error) {
	row := q.db.QueryRow(ctx, createEntity, entityName)
	var i EntityType
	err := row.Scan(&i.EntityID, &i.EntityName)
	return &i, err
}

const deleteEntity = `-- name: DeleteEntity :exec
DELETE FROM entity_type
	WHERE entity_id=$1
    RETURNING entity_id, entity_name
`

func (q *Queries) DeleteEntity(ctx context.Context, entityID int32) error {
	_, err := q.db.Exec(ctx, deleteEntity, entityID)
	return err
}

const findAllEntity = `-- name: FindAllEntity :many
SELECT entity_id, entity_name
	FROM entity_type
`

func (q *Queries) FindAllEntity(ctx context.Context) ([]*EntityType, error) {
	rows, err := q.db.Query(ctx, findAllEntity)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*EntityType
	for rows.Next() {
		var i EntityType
		if err := rows.Scan(&i.EntityID, &i.EntityName); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findEntityById = `-- name: FindEntityById :one
SELECT entity_id, entity_name
	FROM entity_type WHERE entity_id=$1
`

func (q *Queries) FindEntityById(ctx context.Context, entityID int32) (*EntityType, error) {
	row := q.db.QueryRow(ctx, findEntityById, entityID)
	var i EntityType
	err := row.Scan(&i.EntityID, &i.EntityName)
	return &i, err
}

const updateEntity = `-- name: UpdateEntity :one
UPDATE entity_type
	SET entity_name=$2
	WHERE entity_id=$1    
    RETURNING entity_id, entity_name
`

type UpdateEntityParams struct {
	EntityID   int32  `json:"entity_id"`
	EntityName string `json:"entity_name"`
}

func (q *Queries) UpdateEntity(ctx context.Context, arg UpdateEntityParams) (*EntityType, error) {
	row := q.db.QueryRow(ctx, updateEntity, arg.EntityID, arg.EntityName)
	var i EntityType
	err := row.Scan(&i.EntityID, &i.EntityName)
	return &i, err
}
