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
	companyDetailsTableName = "company"

	ColumnCompanyID          domain.Column = "id"
	ColumnCompanyName        domain.Column = "name"
	ColumnCompanyDescription domain.Column = "description"
	ColumnCompanyEmployees   domain.Column = "employees"
	ColumnCompanyRegistered  domain.Column = "registered"
	ColumnCompanyType        domain.Column = "type"
	ColumnCompanyStatus      domain.Column = "status"
)

type CompanyAdaptor struct {
	sqlHandler *sql.DB
}

func NewCompanyAdaptor() *CompanyAdaptor {
	return &CompanyAdaptor{
		sqlHandler: common.GetSqlHandler(),
	}
}

func (adaptor *CompanyAdaptor) Add(
	ctx context.Context,
	companyModel models.CompanyModel,
) error {
	query, args, err := sq.Insert(companyDetailsTableName).Columns(
		string(ColumnCompanyID),
		string(ColumnCompanyName),
		string(ColumnCompanyDescription),
		string(ColumnCompanyEmployees),
		string(ColumnCompanyRegistered),
		string(ColumnCompanyType),
		string(ColumnCompanyStatus),
	).Values(
		companyModel.ID,
		companyModel.Name,
		companyModel.Description,
		companyModel.Employees,
		companyModel.Registered,
		companyModel.Type,
		companyModel.Status,
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

func (adaptor *CompanyAdaptor) Delete(
	ctx context.Context,
	name string,
) error {
	query, args, err := sq.Update(companyDetailsTableName).Where(
		sq.Eq{string(ColumnCompanyName): name},
	).Set(
		string(ColumnCompanyStatus),
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

func (adaptor *CompanyAdaptor) Update(
	ctx context.Context,
	name string,
	updateModel models.UpdateCompanyModel,
) error {
	execUpdate := false
	updateBuilder := sq.Update(companyDetailsTableName).Where(
		sq.Eq{string(ColumnCompanyName): name},
	)
	if updateModel.Description != nil {
		execUpdate = true
		updateBuilder = updateBuilder.Set(string(ColumnCompanyDescription), updateModel.Description)
	}
	if updateModel.Employees != nil {
		execUpdate = true
		updateBuilder = updateBuilder.Set(string(ColumnCompanyEmployees), updateModel.Employees)
	}
	if updateModel.Registered != nil {
		execUpdate = true
		updateBuilder = updateBuilder.Set(string(ColumnCompanyRegistered), updateModel.Registered)
	}
	if updateModel.Type != nil {
		execUpdate = true
		updateBuilder = updateBuilder.Set(string(ColumnCompanyType), updateModel.Type)
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

func (adaptor *CompanyAdaptor) GetCompanyByName(
	ctx context.Context,
	name string,
) (*models.CompanyModel, error) {
	query, args, err := sq.Select(
		string(ColumnCompanyID),
		string(ColumnCompanyName),
		string(ColumnCompanyDescription),
		string(ColumnCompanyEmployees),
		string(ColumnCompanyRegistered),
		string(ColumnCompanyType),
		string(ColumnCompanyStatus),
	).From(companyDetailsTableName).Where(sq.Eq{string(ColumnCompanyName): name}).PlaceholderFormat(sq.Dollar).ToSql()
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
	company := models.CompanyModel{}
	err = rows.Scan(
		&company.ID,
		&company.Name,
		&company.Description,
		&company.Employees,
		&company.Registered,
		&company.Type,
		&company.Status,
	)
	if err != nil {
		return nil, err
	}
	return &company, nil
}
