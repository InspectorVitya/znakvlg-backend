package repository

import (
	"context"
	"github.com/InspectorVitya/znakvlg-backend/internal/model"
	"github.com/InspectorVitya/znakvlg-backend/internal/storage"
	sq "github.com/Masterminds/squirrel"
)

func (r Repository) InsertCompany(ctx context.Context, q storage.Query, company model.Company) (uint32, error) {

	query, args, err := sq.Insert("companies").
		Columns("title", "leader_name", "contract_person", "responsible_person", "phone_number", "mail", "\"comment\"", "head_company", "address", "head_other", "contract_number_ip", "agent_information", "avatar", "print_statement", "header_doc", "flat_plates", "flat_plates2", "flat_plates3", "flat_plates4", "report_module", "type_organization", "contract_number").
		Values(company.Title, company.LeaderName, company.ContractPerson, company.ResponsiblePerson, company.PhoneNumber, company.Mail, company.Comment, company.HeadCompany, company.Address, company.HeadOther, company.ContractNumber, company.AgentInformation, company.Avatar, company.PrintStatement, company.HeaderDoc, company.FlatPlates, company.FlatPlates2, company.FlatPlates3, company.FlatPlates4, company.ReportModule, company.TypeOrganization, company.ContractNumber).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return 0, err
	}

	var id uint32
	err = q.QueryOne(ctx, &id, "insert company", query, args...)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r Repository) SelectCompanies(ctx context.Context, ids []uint32) ([]*model.CompanyInfo, error) {
	return nil, nil
}

func (r Repository) SelectCompanyByID(ctx context.Context, q storage.Query, id uint32) (model.Company, error) {
	query, args, err := sq.Select("id, title, leader_name, contract_person, responsible_person, phone_number, mail, \"comment\", head_company, address, head_other, contract_number_ip, agent_information, avatar, print_statement, header_doc, flat_plates, flat_plates2, flat_plates3, flat_plates4, report_module, type_organization, contract_number").
		From("companies").
		Where("id=$1", id).
		ToSql()
	if err != nil {
		return model.Company{}, err
	}

	company := model.Company{}
	err = q.QueryOne(ctx, &company, "select company by id", query, args...)
	if err != nil {
		return model.Company{}, err
	}

	return company, err
}
