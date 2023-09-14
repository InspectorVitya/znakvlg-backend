package database

import (
	"context"
	"github.com/InspectorVitya/znakvlg-backend/internal/model"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

func (db *DB) InitCompany(ctx context.Context, company model.Company, store []model.Store) (err error) {

	query, args, err := sq.Insert("companies").
		Columns("title", "leader_name", "contract_person", "responsible_person", "phone_number", "mail", "\"comment\"", "head_company", "address", "head_other", "contract_number_ip", "agent_information", "avatar", "print_statement", "header_doc", "flat_plates", "flat_plates2", "flat_plates3", "flat_plates4", "report_module", "type_organization", "contract_number").
		Values(company.Title, company.LeaderName, company.ContractPerson, company.ResponsiblePerson, company.PhoneNumber, company.Mail, company.Comment, company.HeadCompany, company.Address, company.HeadOther, company.ContractNumber, company.AgentInformation, company.Avatar, company.PrintStatement, company.HeaderDoc, company.FlatPlates, company.FlatPlates2, company.FlatPlates3, company.FlatPlates4, company.ReportModule, company.TypeOrganization, company.ContractNumber).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar).ToSql()

	tx, err := db.conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	var id uint32
	err = tx.QueryRow(ctx, query, args...).Scan(&id)
	if err != nil {
		return err
	}

	err = db.insertStore(ctx, tx, id, store)
	if err != nil {
		return err
	}

	return err
}

func (db *DB) SelectCompanies(ctx context.Context, ids []uint32) ([]*model.CompanyInfo, error) {
	return nil, nil
}

func (db *DB) GetCompanyByID(ctx context.Context, id uint32) (model.Company, error) {
	query, args, err := sq.Select("id, title, leader_name, contract_person, responsible_person, phone_number, mail, \"comment\", head_company, address, head_other, contract_number_ip, agent_information, avatar, print_statement, header_doc, flat_plates, flat_plates2, flat_plates3, flat_plates4, report_module, type_organization, contract_number").
		From("companies").
		Where("id=$1", id).
		ToSql()
	if err != nil {
		return model.Company{}, err
	}

	//queryPlate := `select * from company_plates where company_id = ?;`
	company := model.Company{}
	err = db.conn.QueryRow(ctx, query, args...).
		Scan(&company.ID, &company.Title, &company.LeaderName, &company.ContractPerson, &company.ResponsiblePerson, &company.PhoneNumber, &company.Mail, &company.Comment, &company.HeadCompany, &company.Address, &company.HeadOther, &company.ContractNumberIp, &company.AgentInformation, &company.Avatar, &company.PrintStatement, &company.HeaderDoc, &company.FlatPlates, &company.FlatPlates2, &company.FlatPlates3, &company.FlatPlates4, &company.ReportModule, &company.TypeOrganization, &company.ContractNumber)
	if err != nil {
		return model.Company{}, err
	}

	return company, err
}
