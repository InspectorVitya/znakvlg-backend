package repository

import (
	"context"
	"fmt"
	"github.com/InspectorVitya/znakvlg-backend/internal/model"
	"github.com/InspectorVitya/znakvlg-backend/internal/storage"
	"github.com/InspectorVitya/znakvlg-backend/pkg/helpers"
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

type companiesInfo struct {
	ID             uint32 `db:"id"`
	Title          string `db:"title"`
	Address        string `db:"address"`
	Comment        string `db:"comment"`
	PlateTitle     string `db:"title_plate"`
	Count          string `db:"count"`
	WaitAudit      string `db:"wait_audit"`
	WaitAuditToday string `db:"wait_audit_today"`
	Audit          string `db:"audit"`
	AuditToday     string `db:"audit_today"`
	Deffect        string `db:"deffect"`
	DefectToday    string `db:"deffect_today"`
	WaitCheck      string `db:"wait_check"`
	WaitRecheck    string `db:"wait_recheck"`
	PlateUsed      string `db:"plate_used"`
	PlateUsedToday string `db:"plate_used_today"`
}

func (r Repository) SelectCompanies(ctx context.Context, q storage.Query, ids []uint32) ([]*model.CompanyInfo, error) {
	startDay, endDay := helpers.GetStartAndEndDay()
	rows, err := q.QueryAll(ctx, "select company info", selectInfoCompanies, startDay, endDay)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	//res := make(map[uint32]*model.CompanyInfo)
	for rows.Next() {
		var comp companiesInfo
		err := rows.StructScan(&comp)
		if err != nil {
			return nil, err
		}
		fmt.Println(comp)
	}
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
