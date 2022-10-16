package adaptors

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"

	"github.com/akhi19/companies/pkg/common"
	"github.com/akhi19/companies/pkg/domain"
	"github.com/akhi19/companies/pkg/repository/internal/models"
)

const (
	userDetailsTableName = "users"

	ColumnUserID       domain.Column = "id"
	ColumnUsername     domain.Column = "name"
	ColumnUserEmail    domain.Column = "email"
	ColumnUserPassword domain.Column = "password"
	ColumnUserRole     domain.Column = "role"
	ColumnUserStatus   domain.Column = "status"
)

type UserAdaptor struct {
	sqlHandler *sql.DB
}

func NewUserAdaptor() *UserAdaptor {
	return &UserAdaptor{
		sqlHandler: common.GetSqlHandler(),
	}
}

func (adaptor *UserAdaptor) Add(
	ctx context.Context,
	userModel models.UserModel,
) error {
	query, args, err := sq.Insert(userDetailsTableName).Columns(
		string(ColumnUserID),
		string(ColumnUsername),
		string(ColumnUserPassword),
		string(ColumnUserRole),
		string(ColumnUserEmail),
		string(ColumnUserStatus),
	).Values(
		userModel.ID,
		userModel.Name,
		userModel.Password,
		userModel.Role,
		userModel.Email,
		userModel.Status,
	).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return err
	}

	_, err = adaptor.sqlHandler.ExecContext(
		ctx,
		query,
		args...,
	)
	return err
}

func (adaptor *UserAdaptor) Delete(
	ctx context.Context,
	name string,
) error {
	query, args, err := sq.Update(userDetailsTableName).Where(
		sq.Eq{string(ColumnUsername): name},
	).Set(
		string(ColumnUserStatus),
		domain.EntityStatusInactive,
	).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return err
	}

	_, err = adaptor.sqlHandler.ExecContext(
		ctx,
		query,
		args...,
	)
	return err
}

func (adaptor *UserAdaptor) Update(
	ctx context.Context,
	name string,
	updateModel models.UpdateUserModel,
) error {
	execUpdate := false
	updateBuilder := sq.Update(userDetailsTableName).Where(
		sq.Eq{string(ColumnUsername): name},
	)
	if updateModel.Name != nil {
		execUpdate = true
		updateBuilder = updateBuilder.Set(string(ColumnUsername), updateModel.Name)
	}
	if updateModel.Role != nil {
		execUpdate = true
		updateBuilder = updateBuilder.Set(string(ColumnUserRole), updateModel.Role)
	}
	if !execUpdate {
		return nil
	}
	query, args, err := updateBuilder.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return err
	}
	_, err = adaptor.sqlHandler.ExecContext(
		ctx,
		query,
		args...,
	)
	return err
}

func (adaptor *UserAdaptor) GetUserByEmail(
	ctx context.Context,
	email string,
) (*models.UserModel, error) {
	query, args, err := sq.Select(
		string(ColumnUserID),
		string(ColumnUsername),
		string(ColumnUserEmail),
		string(ColumnUserPassword),
		string(ColumnUserRole),
		string(ColumnUserStatus),
	).From(userDetailsTableName).Where(sq.Eq{string(ColumnUserEmail): email}).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := adaptor.sqlHandler.QueryContext(
		ctx,
		query,
		args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, nil
	}
	user := models.UserModel{}
	err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role, &user.Status)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
