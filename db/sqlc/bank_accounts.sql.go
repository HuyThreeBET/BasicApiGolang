// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: bank_accounts.sql

package sqlc

import (
	"context"
)

const addBankAccountBalance = `-- name: AddBankAccountBalance :one
UPDATE bank_accounts
SET balance = balance + $2
WHERE card_number = $3 AND currency = $1
RETURNING username, card_number, balance, currency, created_at, update_at
`

type AddBankAccountBalanceParams struct {
	Currency   string `json:"currency"`
	Amount     int32  `json:"amount"`
	CardNumber string `json:"card_number"`
}

func (q *Queries) AddBankAccountBalance(ctx context.Context, arg AddBankAccountBalanceParams) (BankAccount, error) {
	row := q.db.QueryRowContext(ctx, addBankAccountBalance, arg.Currency, arg.Amount, arg.CardNumber)
	var i BankAccount
	err := row.Scan(
		&i.Username,
		&i.CardNumber,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
		&i.UpdateAt,
	)
	return i, err
}

const createBankAccount = `-- name: CreateBankAccount :one
INSERT INTO bank_accounts (
  username, card_number, currency, balance
) VALUES (
  $1, $2, $3, $4
) RETURNING username, card_number, balance, currency, created_at, update_at
`

type CreateBankAccountParams struct {
	Username   string `json:"username"`
	CardNumber string `json:"card_number"`
	Currency   string `json:"currency"`
	Balance    int32  `json:"balance"`
}

func (q *Queries) CreateBankAccount(ctx context.Context, arg CreateBankAccountParams) (BankAccount, error) {
	row := q.db.QueryRowContext(ctx, createBankAccount,
		arg.Username,
		arg.CardNumber,
		arg.Currency,
		arg.Balance,
	)
	var i BankAccount
	err := row.Scan(
		&i.Username,
		&i.CardNumber,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
		&i.UpdateAt,
	)
	return i, err
}

const deleteBankAccount = `-- name: DeleteBankAccount :exec
DELETE FROM bank_accounts
WHERE card_number = $1
`

func (q *Queries) DeleteBankAccount(ctx context.Context, cardNumber string) error {
	_, err := q.db.ExecContext(ctx, deleteBankAccount, cardNumber)
	return err
}

const deleteBankAccountByCardNumberAndUserName = `-- name: DeleteBankAccountByCardNumberAndUserName :exec
DELETE FROM bank_accounts
WHERE card_number = $1 AND username = $2
`

type DeleteBankAccountByCardNumberAndUserNameParams struct {
	CardNumber string `json:"card_number"`
	Username   string `json:"username"`
}

func (q *Queries) DeleteBankAccountByCardNumberAndUserName(ctx context.Context, arg DeleteBankAccountByCardNumberAndUserNameParams) error {
	_, err := q.db.ExecContext(ctx, deleteBankAccountByCardNumberAndUserName, arg.CardNumber, arg.Username)
	return err
}

const getBankAccountByUserNameAndCurrency = `-- name: GetBankAccountByUserNameAndCurrency :one
SELECT username, card_number, balance, currency, created_at, update_at FROM bank_accounts
WHERE username = $1 AND currency = $2
`

type GetBankAccountByUserNameAndCurrencyParams struct {
	Username string `json:"username"`
	Currency string `json:"currency"`
}

func (q *Queries) GetBankAccountByUserNameAndCurrency(ctx context.Context, arg GetBankAccountByUserNameAndCurrencyParams) (BankAccount, error) {
	row := q.db.QueryRowContext(ctx, getBankAccountByUserNameAndCurrency, arg.Username, arg.Currency)
	var i BankAccount
	err := row.Scan(
		&i.Username,
		&i.CardNumber,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
		&i.UpdateAt,
	)
	return i, err
}

const getBankAccountByUserNameAndCurrencyForUpdate = `-- name: GetBankAccountByUserNameAndCurrencyForUpdate :one
SELECT username, card_number, balance, currency, created_at, update_at FROM bank_accounts
WHERE username = $1 AND currency = $2
FOR UPDATE
`

type GetBankAccountByUserNameAndCurrencyForUpdateParams struct {
	Username string `json:"username"`
	Currency string `json:"currency"`
}

func (q *Queries) GetBankAccountByUserNameAndCurrencyForUpdate(ctx context.Context, arg GetBankAccountByUserNameAndCurrencyForUpdateParams) (BankAccount, error) {
	row := q.db.QueryRowContext(ctx, getBankAccountByUserNameAndCurrencyForUpdate, arg.Username, arg.Currency)
	var i BankAccount
	err := row.Scan(
		&i.Username,
		&i.CardNumber,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
		&i.UpdateAt,
	)
	return i, err
}

const getCardNumberByCardNumberAndUsername = `-- name: GetCardNumberByCardNumberAndUsername :one
SELECT card_number FROM bank_accounts
WHERE username = $1 AND card_number = $2
`

type GetCardNumberByCardNumberAndUsernameParams struct {
	Username   string `json:"username"`
	CardNumber string `json:"card_number"`
}

func (q *Queries) GetCardNumberByCardNumberAndUsername(ctx context.Context, arg GetCardNumberByCardNumberAndUsernameParams) (string, error) {
	row := q.db.QueryRowContext(ctx, getCardNumberByCardNumberAndUsername, arg.Username, arg.CardNumber)
	var card_number string
	err := row.Scan(&card_number)
	return card_number, err
}

const getCurrencyByCardNumberAndUsername = `-- name: GetCurrencyByCardNumberAndUsername :one
SELECT currency FROM bank_accounts
WHERE username = $1 AND card_number = $2
`

type GetCurrencyByCardNumberAndUsernameParams struct {
	Username   string `json:"username"`
	CardNumber string `json:"card_number"`
}

func (q *Queries) GetCurrencyByCardNumberAndUsername(ctx context.Context, arg GetCurrencyByCardNumberAndUsernameParams) (string, error) {
	row := q.db.QueryRowContext(ctx, getCurrencyByCardNumberAndUsername, arg.Username, arg.CardNumber)
	var currency string
	err := row.Scan(&currency)
	return currency, err
}

const listBankAccounts = `-- name: ListBankAccounts :many
SELECT username, card_number, balance, currency, created_at, update_at FROM bank_accounts
WHERE username = $1
ORDER BY username
`

func (q *Queries) ListBankAccounts(ctx context.Context, username string) ([]BankAccount, error) {
	rows, err := q.db.QueryContext(ctx, listBankAccounts, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []BankAccount{}
	for rows.Next() {
		var i BankAccount
		if err := rows.Scan(
			&i.Username,
			&i.CardNumber,
			&i.Balance,
			&i.Currency,
			&i.CreatedAt,
			&i.UpdateAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listBankAccountsByUsername = `-- name: ListBankAccountsByUsername :many
SELECT username, card_number, balance, currency, created_at, update_at FROM bank_accounts
WHERE username = $1
ORDER BY username
`

func (q *Queries) ListBankAccountsByUsername(ctx context.Context, username string) ([]BankAccount, error) {
	rows, err := q.db.QueryContext(ctx, listBankAccountsByUsername, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []BankAccount{}
	for rows.Next() {
		var i BankAccount
		if err := rows.Scan(
			&i.Username,
			&i.CardNumber,
			&i.Balance,
			&i.Currency,
			&i.CreatedAt,
			&i.UpdateAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
