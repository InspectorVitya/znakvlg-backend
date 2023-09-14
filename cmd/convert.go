package main

import (
	"fmt"
	"github.com/InspectorVitya/znakvlg-backend/internal/model"
	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"log"
	"strings"
	"time"
)

func main() {
	db, err := sqlx.Connect("mysql", getDBURL("localhost", "znakvlg", "root", "12345678", "3306"))
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	db2, err := sqlx.Connect("pgx", "postgres://gen_user:znakvlg12345@80.90.185.138:5432/default_db")
	if err != nil {
		log.Fatal(err)
	}
	err = db2.Ping()
	if err != nil {
		log.Fatal(err)
	}
	//comp := company(db)
	//InsertCompany(db2, comp)
	//product(db, db2)
	store(db, db2)
}

func InsertCompany(db *sqlx.DB, comp []model.Company) {
	query := "INSERT INTO companies (id, title, leader_name, contract_person, responsible_person, phone_number, mail, comment, head_company, address, head_other, contract_number_ip, agent_information, avatar, print_statement, header_doc, flat_plates, flat_plates2, flat_plates3, flat_plates4, report_module, type_organization, contract_number) OVERRIDING SYSTEM VALUE VALUES(:id, :title, :leader_name, :contract_person, :responsible_person, :phone_number, :mail, :comment, :head_company, :address, :head_other, :contract_number_ip, :agent_information, :avatar, :print_statement, :header_doc, :flat_plates, :flat_plates2, :flat_plates3, :flat_plates4, :report_module, :type_organization, :contract_number)"
	for _, v := range comp {
		_, err := db.NamedExec(query, &v)
		fmt.Println(v.ContractNumber)
		if err != nil {
			log.Fatal(err)
		}
	}

}

type Company struct {
	Id                int32   `db:"id"`
	Title             *string `db:"title"`
	LeaderName        *string `db:"leader_name"`
	ContactPerson     *string `db:"contact_person"`
	ResponsiblePerson *string `db:"responsible_person"`
	PhoneNumber       *string `db:"phone_number"`
	Mail              *string `db:"mail"`
	PrintStatement    bool    `db:"print_statement"`
	Comment           *string `db:"comment"`
	HeadCompany       *string `db:"head_company"`
	Address           *string `db:"adress"`
	HeaderDocBool     bool    `db:"header_doc_bool"`
	HeadOther         *string `db:"head_other"`
	FlatPlates        bool    `db:"flat_plates"`
	FlatPlates2       bool    `db:"flat_plates2"`
	ReportPerm        bool    `db:"report_perm"`
	TypeOrganization  bool    `db:"type_organization"`
	ContractNumberIp  *string `db:"contract_number_ip"`
	ContractNumber    *string `db:"contract_number"`
	AgentInformation  *string `db:"agent_information"`
	Avatar            *string `db:"avatar"`
	FlatPlates3       bool    `db:"flat_plates3"`
	FlatPlates4       bool    `db:"flat_plates4"`
}

func company(db *sqlx.DB) []model.Company {
	res := make([]Company, 0)
	err := db.Select(&res, "SELECT * FROM companies;")
	if err != nil {
		log.Fatal(err)
	}
	l := make([]model.Company, len(res))
	for i := range res {
		l[i].ID = uint32(res[i].Id)
		l[i].Title = stringNil(res[i].Title)
		l[i].LeaderName = stringNil(res[i].LeaderName)
		l[i].ContractPerson = stringNil(res[i].ContactPerson)
		l[i].ResponsiblePerson = stringNil(res[i].ResponsiblePerson)
		l[i].PhoneNumber = stringNil(res[i].PhoneNumber)
		l[i].Mail = stringNil(res[i].Mail)
		l[i].PrintStatement = res[i].PrintStatement
		l[i].Comment = stringNil(res[i].Comment)
		l[i].HeadCompany = stringNil(res[i].HeadCompany)
		l[i].Address = stringNil(res[i].Address)
		l[i].HeaderDoc = res[i].HeaderDocBool
		l[i].HeadOther = stringNil(res[i].HeadOther)
		l[i].ReportModule = res[i].ReportPerm
		l[i].TypeOrganization = res[i].TypeOrganization
		l[i].ContractNumberIp = stringNil(res[i].ContractNumber)
		l[i].ContractNumber = stringNil(res[i].ContractNumber)
		l[i].AgentInformation = stringNil(res[i].AgentInformation)
		l[i].Avatar = stringNil(res[i].Avatar)
		l[i].FlatPlates = res[i].FlatPlates
		l[i].FlatPlates2 = res[i].FlatPlates2
		l[i].FlatPlates3 = res[i].FlatPlates3
		l[i].FlatPlates4 = res[i].FlatPlates4
	}

	return l
}

type DefectProduct struct {
	Id           int32  `db:"id"`
	LicensePlate string `db:"license_plate"`
	CountPlate   int32  `db:"count_plate"`
	DateCreate   string `db:"date_create"`
	CompanyId    int32  `db:"companyId"`
	PlateTypeId  int32  `db:"plateTypeId"`
}

func defect(db, db2 *sqlx.DB) {
	def := make([]DefectProduct, 0)
	err := db.Select(&def, "SELECT id, license_plate, count_plate, date_create, companyId, plateTypeId FROM defect_products;")
	if err != nil {
		log.Fatal(err)
	}
	res := make([]model.DefectProduct, len(def))
	for i, v := range def {
		res[i].ID = uint32(v.Id)
		res[i].PlateTypeId = uint8(v.PlateTypeId)
		res[i].CountPlate = v.CountPlate
		res[i].DateCreate, err = time.Parse("02.01.2006", v.DateCreate) //14.07.2020
		if err != nil {
			log.Fatal(err)
		}
		res[i].CompanyId = v.CompanyId
		res[i].LicensePlate = v.LicensePlate
	}
	query := "insert into defect_products (id, license_plate, count_plate, date_create, company_id, plate_type_id) values (:id, :license_plate, :count_plate, :date_create, :company_id, :plate_type_id)"
	for _, v := range res {
		_, err := db2.NamedExec(query, &v)
		if err != nil {
			log.Fatal(err)
		}
	}
}

type Product struct {
	Id                                           int32   `db:"id"`
	LicensePlate                                 *string `db:"license_plate"`
	IssuedDoc                                    *string `db:"issued_doc"`
	RepresentativeDocumentIssuedDoc              *string `db:"representative_document_issued_doc"`
	IndividualDocumentIssuedDoc                  *string `db:"individual_document_issued_doc"`
	SerialNumber                                 *string `db:"serial_number"`
	DateCreate                                   *string `db:"date_create"`
	DateIssuing                                  *string `db:"date_issuing"`
	Document                                     *string `db:"document"`
	DateIssuingTransportDocument                 *string `db:"date_issuing_transport_document"`
	CountPlate                                   *int32  `db:"count_plate"`
	IndividualDocumentOtherType                  *string `db:"individual_document_other_type"`
	IndividualDocumentSurname                    *string `db:"individual_document_surname"`
	IndividualDocumentName                       *string `db:"individual_document_name"`
	IndividualDocumentPatronymic                 *string `db:"individual_document_patronymic"`
	IndividualDocumentSerialNumber               *string `db:"individual_document_serial_number"`
	IndividualDocumentDateIssuing                *string `db:"individual_document_date_issuing"`
	EntityDocumentName                           *string `db:"entity_document_name"`
	EntityDocumentInn                            *string `db:"entity_document_inn"`
	EntityDocumentOrgn                           *string `db:"entity_document_orgn"`
	RepresentativeTypeOwnerId                    *int32  `db:"representative_type_owner_id"`
	RepresentativeIndividualDocumentOtherType    *string `db:"representative_individual_document_other_type"`
	RepresentativeIndividualDocumentSurname      *string `db:"representative_individual_document_surname"`
	RepresentativeIndividualDocumentName         *string `db:"representative_individual_document_name"`
	RepresentativeIndividualDocumentPatronymic   *string `db:"representative_individual_document_patronymic"`
	RepresentativeIndividualDocumentSerialNumber *string `db:"representative_individual_document_serial_number"`
	RepresentativeIndividualDocumentDateIssuing  *string `db:"representative_individual_document_date_issuing"`
	RepresentativeEntityDocumentName             *string `db:"representative_entity_document_name"`
	RepresentativeEntityDocumentInn              *string `db:"representative_entity_document_inn"`
	RepresentativeEntityDocumentOrgn             *string `db:"representative_entity_document_orgn"`
	RepresentativeDocumentTypeOwnerId            *int32  `db:"representative_document_type_ownerId"`
	CompanyId                                    *int32  `db:"companyId"`
	PlateTypeId                                  *int32  `db:"plateTypeId"`
	TypeDocumentOwnerId                          *uint8  `db:"typeDocumentOwnerId"`
	TypeDocumentTransportId                      *int32  `db:"typeDocumentTransportId"`
	OwnerTypeId                                  *int32  `db:"ownerTypeId"`
	StatusProductId                              *int32  `db:"statusProductId"`
}

func product(db, db2 *sqlx.DB) {
	prod := make([]Product, 0)
	err := db.Select(&prod, "SELECT id, license_plate, issued_doc, representative_document_issued_doc, individual_document_issued_doc, serial_number, date_create, date_issuing, document, date_issuing_transport_document, count_plate, individual_document_other_type, individual_document_surname, individual_document_name, individual_document_patronymic, individual_document_serial_number, individual_document_date_issuing, entity_document_name, entity_document_inn, entity_document_orgn, representative_type_owner_id, representative_individual_document_other_type, representative_individual_document_surname, representative_individual_document_name, representative_individual_document_patronymic, representative_individual_document_serial_number, representative_individual_document_date_issuing, representative_entity_document_name, representative_entity_document_inn, representative_entity_document_orgn, representative_document_type_ownerId, companyId, plateTypeId, typeDocumentOwnerId, typeDocumentTransportId, ownerTypeId, statusProductId FROM products;")
	if err != nil {
		log.Fatal(err)
	}
	res := make([]model.Products, 0, len(prod))
	for _, v := range prod {
		p := model.Products{
			ID:                                           uint32(v.Id),
			LicensePlate:                                 stringNil(v.LicensePlate),
			IssuedDoc:                                    stringNil(v.IssuedDoc),
			RepresentativeDocumentIssuedDoc:              stringNil(v.RepresentativeDocumentIssuedDoc),
			IndividualDocumentIssuedDoc:                  stringNil(v.IndividualDocumentIssuedDoc),
			SerialNumber:                                 stringNil(v.SerialNumber),
			Document:                                     stringNil(v.Document),
			DateIssuingTransportDocument:                 stringNil(v.DateIssuingTransportDocument),
			CountPlate:                                   uint8(intNil(v.CountPlate)),
			IndividualDocumentOtherType:                  stringNil(v.IndividualDocumentOtherType),
			IndividualDocumentSurname:                    stringNil(v.IndividualDocumentSurname),
			IndividualDocumentName:                       stringNil(v.IndividualDocumentName),
			IndividualDocumentPatronymic:                 stringNil(v.IndividualDocumentPatronymic),
			IndividualDocumentSerialNumber:               stringNil(v.IndividualDocumentSerialNumber),
			IndividualDocumentDateIssuing:                stringNil(v.IndividualDocumentDateIssuing),
			EntityDocumentName:                           stringNil(v.EntityDocumentName),
			EntityDocumentInn:                            stringNil(v.EntityDocumentInn),
			EntityDocumentOrgn:                           stringNil(v.EntityDocumentOrgn),
			RepresentativeTypeOwnerID:                    uint8(intNil(v.RepresentativeDocumentTypeOwnerId)),
			RepresentativeIndividualDocumentOtherType:    stringNil(v.RepresentativeIndividualDocumentOtherType),
			RepresentativeIndividualDocumentSurname:      stringNil(v.RepresentativeIndividualDocumentSurname),
			RepresentativeIndividualDocumentName:         stringNil(v.RepresentativeIndividualDocumentName),
			RepresentativeIndividualDocumentPatronymic:   stringNil(v.RepresentativeIndividualDocumentPatronymic),
			RepresentativeIndividualDocumentSerialNumber: stringNil(v.RepresentativeIndividualDocumentSerialNumber),
			RepresentativeIndividualDocumentDateIssuing:  stringNil(v.RepresentativeIndividualDocumentDateIssuing),
			RepresentativeEntityDocumentName:             stringNil(v.RepresentativeEntityDocumentName),
			RepresentativeEntityDocumentInn:              stringNil(v.RepresentativeEntityDocumentInn),
			RepresentativeEntityDocumentOrgn:             stringNil(v.RepresentativeEntityDocumentOrgn),
			RepresentativeDocumentTypeOwnerID:            uint8(intNil(v.TypeDocumentTransportId)),
			CompanyID:                                    uint32(intNil(v.CompanyId)),
			PlateTypeID:                                  uint8(intNil(v.PlateTypeId)),
			TypeDocumentOwnerID:                          v.TypeDocumentOwnerId,
			TypeDocumentTransportID:                      uint8(intNil(v.TypeDocumentTransportId)),
			OwnerTypeID:                                  uint8(intNil(v.OwnerTypeId)),
			StatusProductID:                              uint8(intNil(v.StatusProductId)),
		}
		if strings.Split(stringNil(v.DateCreate), ".")[2] == "23" || strings.Split(stringNil(v.DateCreate), ".")[2] == "22" || strings.Split(stringNil(v.DateCreate), ".")[2] == "21" {
			str := strings.Split(stringNil(v.DateCreate), ".")
			if str[2] == "23" {
				str[2] = "2023"
			}
			if str[2] == "22" {
				str[2] = "2022"
			}
			if str[2] == "21" {
				str[2] = "2021"
			}
			s := strings.Join(str, ".")
			v.DateCreate = &s
		}
		DateCreate, err := time.Parse("02.01.2006", stringNil(v.DateCreate)) //14.07.2020
		if err != nil {
			log.Println(err, "create")
		}
		if strings.Split(stringNil(v.DateIssuing), ".")[2] == "23" || strings.Split(stringNil(v.DateIssuing), ".")[2] == "22" || strings.Split(stringNil(v.DateIssuing), ".")[2] == "21" {
			str := strings.Split(stringNil(v.DateIssuing), ".")
			if str[2] == "23" {
				str[2] = "2023"
			}
			if str[2] == "22" {
				str[2] = "2022"
			}
			if str[2] == "21" {
				str[2] = "2021"
			}
			s := strings.Join(str, ".")
			v.DateIssuing = &s
		}
		DateIssuing, err := time.Parse("02.01.2006", stringNil(v.DateIssuing)) //14.07.2020
		if err != nil {
			if v.DateIssuing == nil {
				DateIssuing = time.Now().AddDate(-2, 0, 0)
			} else {
				log.Println(err, "iss")
			}

		}
		p.DateIssuing = DateIssuing
		p.DateCreate = DateCreate
		res = append(res, p)
	}
	//query := "insert into products (id, license_plate, issued_doc, representative_document_issued_doc, individual_document_issued_doc, serial_number, date_create, date_issuing, document, date_issuing_transport_document, count_plate, individual_document_other_type, individual_document_surname, individual_document_name, individual_document_patronymic, individual_document_serial_number, individual_document_date_issuing, entity_document_name, entity_document_inn, entity_document_orgn, representative_type_owner_id, representative_individual_document_other_type, representative_individual_document_surname, representative_individual_document_name, representative_individual_document_patronymic, representative_individual_document_serial_number, representative_individual_document_date_issuing, representative_entity_document_name, representative_entity_document_inn, representative_entity_document_orgn, representative_document_type_ownerId, company_id, plate_type_id, type_document_owner_id, type_document_transport_id, owner_type_id, status_product_id) values (:id, :license_plate, :issued_doc, :representative_document_issued_doc, :individual_document_issued_doc, :serial_number, :date_create, :date_issuing, :document, :date_issuing_transport_document, :count_plate, :individual_document_other_type, :individual_document_surname, :individual_document_name, :individual_document_patronymic, :individual_document_serial_number, :individual_document_date_issuing, :entity_document_name, :entity_document_inn, :entity_document_orgn, :representative_type_owner_id, :representative_individual_document_other_type, :representative_individual_document_surname, :representative_individual_document_name, :representative_individual_document_patronymic, :representative_individual_document_serial_number, :representative_individual_document_date_issuing, :representative_entity_document_name, :representative_entity_document_inn, :representative_entity_document_orgn, :representative_document_type_ownerId, :company_id, :plate_type_id, :type_document_owner_id, :type_document_transport_id, :owner_type_id, :status_product_id)"
	//for i := range res {
	//	_, err = db2.NamedExec(query, &res[i])
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	fmt.Println(i)
	//}
	qInsert := sq.Insert("products").Columns("id", "license_plate", "issued_doc", "representative_document_issued_doc", "individual_document_issued_doc", "serial_number", "date_create", "date_issuing", "document", "date_issuing_transport_document", "count_plate", "individual_document_other_type", "individual_document_surname", "individual_document_name", "individual_document_patronymic", "individual_document_serial_number", "individual_document_date_issuing", "entity_document_name", "entity_document_inn", "entity_document_orgn", "representative_type_owner_id", "representative_individual_document_other_type", "representative_individual_document_surname", "representative_individual_document_name", "representative_individual_document_patronymic", "representative_individual_document_serial_number", "representative_individual_document_date_issuing", "representative_entity_document_name", "representative_entity_document_inn", "representative_entity_document_orgn", "representative_document_type_ownerId", "company_id", "plate_type_id", "type_document_owner_id", "type_document_transport_id", "owner_type_id", "status_product_id")
	d := qInsert
	count := 0
	sum := 0
	for i, v := range res {
		count++
		sum++
		d = d.Values(v.ID, v.LicensePlate, v.IssuedDoc, v.RepresentativeDocumentIssuedDoc, v.IndividualDocumentIssuedDoc, v.SerialNumber, v.DateCreate, v.DateIssuing, v.Document, v.DateIssuingTransportDocument, v.CountPlate, v.IndividualDocumentOtherType, v.IndividualDocumentSurname, v.IndividualDocumentName, v.IndividualDocumentPatronymic, v.IndividualDocumentSerialNumber, v.IndividualDocumentDateIssuing, v.EntityDocumentName, v.EntityDocumentInn, v.EntityDocumentOrgn, v.RepresentativeTypeOwnerID, v.RepresentativeIndividualDocumentOtherType, v.RepresentativeIndividualDocumentSurname, v.RepresentativeIndividualDocumentName, v.RepresentativeIndividualDocumentPatronymic, v.RepresentativeIndividualDocumentSerialNumber, v.RepresentativeIndividualDocumentDateIssuing, v.RepresentativeEntityDocumentName, v.RepresentativeEntityDocumentInn, v.RepresentativeEntityDocumentOrgn, v.RepresentativeDocumentTypeOwnerID, v.CompanyID, v.PlateTypeID, v.TypeDocumentOwnerID, v.TypeDocumentTransportID, v.OwnerTypeID, v.StatusProductID)
		if count == 800 {
			sqlInsert, args, err := d.ToSql()
			if err != nil {
				log.Fatal(err)
			}
			_, err = db2.Exec(sqlInsert, args...)
			if err != nil {
				log.Fatal(err)
			}
			count = 0
			d = qInsert
			fmt.Println(sum)
			continue
		}
		if i == len(res)-1 {
			fmt.Println(sum)
			sqlInsert, args, err := d.ToSql()
			if err != nil {
				log.Fatal(err)
			}
			_, err = db2.Exec(sqlInsert, args...)
			if err != nil {
				log.Fatal(err)
			}
			count = 0
			d = qInsert
			continue
		}

	}
}

type stores struct {
	Id          int `db:"id"`
	Count       int `db:"count"`
	CompanyID   int `db:"companyId"`
	PlateTypeId int `db:"plateTypeId"`
}

func store(db, db2 *sqlx.DB) {
	sql := "SELECT id, count, companyId, plateTypeId FROM stores;"
	prod := make([]stores, 0)
	err := db.Select(&prod, sql)
	if err != nil {
		log.Fatal(err)
	}
	queryStore := sq.Insert("store").Columns("count", "company_id", "plate_type_id", "used").PlaceholderFormat(sq.Dollar)
	for _, v := range prod {

		queryStore = queryStore.Values(v.Count, uint32(v.CompanyID), uint8(v.PlateTypeId), true)
	}
	sqlInsert, args, err := queryStore.ToSql()
	if err != nil {
		log.Fatal(err)
	}
	_, err = db2.Exec(sqlInsert, args...)
	if err != nil {
		log.Fatal(err)
	}
}

func stringNil(str *string) string {
	if str == nil {
		return ""
	}
	return *str
}

func intNil(in *int32) int32 {
	if in == nil {
		return 0
	}
	return *in
}

func getDBURL(host, nameDB, user, psw, port string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, psw, host, port, nameDB)
}
