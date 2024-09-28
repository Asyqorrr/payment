// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type BusinessEntity struct {
	ButyID     int32   `json:"buty_id"`
	ButyName   *string `json:"buty_name"`
	ButyCode   *string `json:"buty_code"`
	ButyEntyID *int32  `json:"buty_enty_id"`
}

type EntityType struct {
	EntityID   int32  `json:"entity_id"`
	EntityName string `json:"entity_name"`
}

type PaymentTransaction struct {
	PatrxNo         string      `json:"patrx_no"`
	PatrxCreatedOn  pgtype.Date `json:"patrx_created_on"`
	PatrxDebet      *float64    `json:"patrx_debet"`
	PatrxCredit     *float64    `json:"patrx_credit"`
	PatrxNotes      *string     `json:"patrx_notes"`
	PatrxAcctnoFrom *string     `json:"patrx_acctno_from"`
	PatrxAcctnoTo   *string     `json:"patrx_acctno_to"`
	PatrxPatrxRef   *string     `json:"patrx_patrx_ref"`
	PatrxTratyID    *int32      `json:"patrx_traty_id"`
}

type TransactionType struct {
	TratyID   int32   `json:"traty_id"`
	TratyName *string `json:"traty_name"`
}

type User struct {
	UserID        int32       `json:"user_id"`
	UserPassword  *string     `json:"user_password"`
	UserName      *string     `json:"user_name"`
	UserEmail     *string     `json:"user_email"`
	UserHandphone *string     `json:"user_handphone"`
	UserCreated   interface{} `json:"user_created"`
}

type UserAccount struct {
	UsacID        int32       `json:"usac_id"`
	UsacAccountNo *string     `json:"usac_account_no"`
	UsacBalance   *float64    `json:"usac_balance"`
	UsacCreatedOn pgtype.Date `json:"usac_created_on"`
	UsacButyID    *int32      `json:"usac_buty_id"`
	UsacUserID    *int32      `json:"usac_user_id"`
}
