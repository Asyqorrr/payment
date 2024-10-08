// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: user_accounts.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUserAccount = `-- name: CreateUserAccount :one
INSERT INTO user_accounts(
	usac_account_no, usac_balance, usac_created_on, usac_buty_id, usac_user_id)
	VALUES (
	$1, $2, $3, $4, $5)
    RETURNING usac_id, usac_account_no, usac_balance, usac_created_on, usac_buty_id, usac_user_id
`

type CreateUserAccountParams struct {
	UsacAccountNo *string     `json:"usac_account_no"`
	UsacBalance   *float64    `json:"usac_balance"`
	UsacCreatedOn pgtype.Date `json:"usac_created_on"`
	UsacButyID    *int32      `json:"usac_buty_id"`
	UsacUserID    *int32      `json:"usac_user_id"`
}

func (q *Queries) CreateUserAccount(ctx context.Context, arg CreateUserAccountParams) (*UserAccount, error) {
	row := q.db.QueryRow(ctx, createUserAccount,
		arg.UsacAccountNo,
		arg.UsacBalance,
		arg.UsacCreatedOn,
		arg.UsacButyID,
		arg.UsacUserID,
	)
	var i UserAccount
	err := row.Scan(
		&i.UsacID,
		&i.UsacAccountNo,
		&i.UsacBalance,
		&i.UsacCreatedOn,
		&i.UsacButyID,
		&i.UsacUserID,
	)
	return &i, err
}

const deleteUserAccount = `-- name: DeleteUserAccount :exec
DELETE FROM user_accounts
    WHERE usac_user_id = $1
    RETURNING usac_id, usac_account_no, usac_balance, usac_created_on, usac_buty_id, usac_user_id
`

func (q *Queries) DeleteUserAccount(ctx context.Context, usacUserID *int32) error {
	_, err := q.db.Exec(ctx, deleteUserAccount, usacUserID)
	return err
}

const findUserAccountByAccno = `-- name: FindUserAccountByAccno :one
SELECT ua.usac_id, usac_account_no, usac_balance, usac_created_on, u.user_name, be.buty_name
FROM user_accounts as ua
join users as u on ua.usac_user_id = u.user_id
join business_entity as be on ua.usac_buty_id = be.buty_id
where ua.usac_account_no = $1
`

type FindUserAccountByAccnoRow struct {
	UsacID        int32       `json:"usac_id"`
	UsacAccountNo *string     `json:"usac_account_no"`
	UsacBalance   *float64    `json:"usac_balance"`
	UsacCreatedOn pgtype.Date `json:"usac_created_on"`
	UserName      *string     `json:"user_name"`
	ButyName      *string     `json:"buty_name"`
}

func (q *Queries) FindUserAccountByAccno(ctx context.Context, usacAccountNo *string) (*FindUserAccountByAccnoRow, error) {
	row := q.db.QueryRow(ctx, findUserAccountByAccno, usacAccountNo)
	var i FindUserAccountByAccnoRow
	err := row.Scan(
		&i.UsacID,
		&i.UsacAccountNo,
		&i.UsacBalance,
		&i.UsacCreatedOn,
		&i.UserName,
		&i.ButyName,
	)
	return &i, err
}

const getAllUserAccount = `-- name: GetAllUserAccount :many
SELECT ua.usac_id, usac_account_no, usac_balance, usac_created_on, u.user_name, be.buty_name
FROM user_accounts as ua
join users as u on ua.usac_user_id = u.user_id
join business_entity as be on ua.usac_buty_id = be.buty_id
`

type GetAllUserAccountRow struct {
	UsacID        int32       `json:"usac_id"`
	UsacAccountNo *string     `json:"usac_account_no"`
	UsacBalance   *float64    `json:"usac_balance"`
	UsacCreatedOn pgtype.Date `json:"usac_created_on"`
	UserName      *string     `json:"user_name"`
	ButyName      *string     `json:"buty_name"`
}

func (q *Queries) GetAllUserAccount(ctx context.Context) ([]*GetAllUserAccountRow, error) {
	rows, err := q.db.Query(ctx, getAllUserAccount)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetAllUserAccountRow
	for rows.Next() {
		var i GetAllUserAccountRow
		if err := rows.Scan(
			&i.UsacID,
			&i.UsacAccountNo,
			&i.UsacBalance,
			&i.UsacCreatedOn,
			&i.UserName,
			&i.ButyName,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateBalance = `-- name: UpdateBalance :one
UPDATE user_accounts
	SET  usac_balance=$2
	WHERE usac_account_no=$1
	RETURNING usac_id, usac_account_no, usac_balance, usac_created_on, usac_buty_id, usac_user_id
`

type UpdateBalanceParams struct {
	UsacAccountNo *string  `json:"usac_account_no"`
	UsacBalance   *float64 `json:"usac_balance"`
}

func (q *Queries) UpdateBalance(ctx context.Context, arg UpdateBalanceParams) (*UserAccount, error) {
	row := q.db.QueryRow(ctx, updateBalance, arg.UsacAccountNo, arg.UsacBalance)
	var i UserAccount
	err := row.Scan(
		&i.UsacID,
		&i.UsacAccountNo,
		&i.UsacBalance,
		&i.UsacCreatedOn,
		&i.UsacButyID,
		&i.UsacUserID,
	)
	return &i, err
}
