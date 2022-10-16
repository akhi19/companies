package services

import (
	"context"
	"errors"
	"testing"

	"github.com/akhi19/companies/pkg/company/internal"
	"github.com/akhi19/companies/pkg/company/internal/adaptors"
	"github.com/akhi19/companies/pkg/domain"
	"github.com/akhi19/companies/pkg/repository"
	mock_ports "github.com/akhi19/companies/pkg/repository/mocks"
	"github.com/golang/mock/gomock"
)

var commandService CompanyCommandService
var mockCompanyContainer *mock_ports.MockICompany

func NumPointer(num int64) *int64 {
	return &num
}

func BoolPointer(value bool) *bool {
	return &value
}

func InitCommandServiceTest(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockCompanyContainer = mock_ports.NewMockICompany(ctrl)
	companyContainer := repository.CompanyContainer{
		ICompany: mockCompanyContainer,
	}

	commandService = CompanyCommandService{
		repositoryAdaptor: adaptors.NewRepositoryAdaptor(
			companyContainer,
		),
	}
}

func TestCompanyCommandService_Add(t *testing.T) {
	InitCommandServiceTest(t)
	type args struct {
		ctx               context.Context
		addCompanyRequest internal.AddCompanyRequestDTO
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		mockFunc func(args args)
	}{
		// VALID TEST CASE
		{
			name: "Valid addition",
			args: args{
				ctx: context.Background(),
				addCompanyRequest: internal.AddCompanyRequestDTO{
					Name:        "Test",
					Description: "Test",
					Employees:   NumPointer(12),
					Registered:  BoolPointer(false),
					Type:        domain.CompanyTypeNonProfit,
				},
			},
			wantErr: false,
			mockFunc: func(args args) {
				mockCompanyContainer.EXPECT().Add(args.ctx, args.addCompanyRequest).Return(
					nil,
				).Times(1)
			},
		},

		// Database inseertion error
		{
			name: "Database insertion error",
			args: args{
				ctx: context.Background(),
				addCompanyRequest: internal.AddCompanyRequestDTO{
					Name:        "Test",
					Description: "123",
					Employees:   NumPointer(12),
					Registered:  BoolPointer(true),
					Type:        domain.CompanyTypeCooperative,
				},
			},
			wantErr: true,
			mockFunc: func(args args) {
				mockCompanyContainer.EXPECT().Add(args.ctx, args.addCompanyRequest).Return(
					errors.New(""),
				).Times(1)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc(tt.args)
			if err := commandService.Add(tt.args.ctx, tt.args.addCompanyRequest); (err != nil) != tt.wantErr {
				t.Errorf("CompanyCommandService.Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
