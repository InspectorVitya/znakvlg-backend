package model

import (
	"time"
)

type Users struct {
	ID          string     `json:"id,omitempty" db:"id"`
	Login       string     `json:"login,omitempty" db:"login"`
	Email       string     `json:"email,omitempty" db:"email"`
	Password    string     `json:"password" db:"password"`
	Blocked     bool       `json:"blocked,omitempty" db:"blocked"`
	RoleID      uint8      `json:"roleId,omitempty" db:"role_id"`
	Comment     string     `json:"comment,omitempty" db:"comment"`
	PhoneNumber string     `json:"phoneNumber,omitempty" db:"phone_number"`
	Name        string     `json:"name,omitempty" db:"name"`
	Patronymic  string     `json:"patronymic,omitempty" db:"patronymic"`
	SurName     string     `json:"surname,omitempty" db:"surname"`
	LastLogin   *time.Time `json:"authDate,omitempty" db:"last_login"`
	Address     string     `json:"address,omitempty" db:"address"`
	WorkPlace   WorkPlace  `json:"workPlace,omitempty"`
}

type WorkPlace []uint32

type Company struct {
	ID                uint32 `json:"id,omitempty" db:"id"`
	Title             string `json:"title" db:"title"`
	LeaderName        string `json:"leader_name" db:"leader_name"`
	ContractPerson    string `json:"contract_person" db:"contract_person"`
	ResponsiblePerson string `json:"responsible_person" db:"responsible_person"`
	PhoneNumber       string `json:"phone_number" db:"phone_number"`
	Mail              string `json:"mail" db:"mail"`
	PrintStatement    bool   `json:"print_statement" db:"print_statement"`
	Comment           string `json:"comment" db:"comment"`
	HeadCompany       string `json:"head_company" db:"head_company"`
	Address           string `json:"address" db:"address"`
	HeaderDoc         bool   `json:"header_doc" db:"header_doc"`
	HeadOther         string `json:"head_other" db:"head_other"`
	ReportModule      bool   `json:"report_module" db:"report_module"`
	TypeOrganization  bool   `json:"type_organization" db:"type_organization"`
	ContractNumberIp  string `json:"contract_number_ip" db:"contract_number_ip"`
	ContractNumber    string `json:"contract_number" db:"contract_number"`
	AgentInformation  string `json:"agent_information" db:"agent_information"`
	Avatar            string `json:"avatar" db:"avatar"`
	FlatPlates        bool   `json:"flat_plates" db:"flat_plates"`
	FlatPlates2       bool   `json:"flat_plates2" db:"flat_plates2"`
	FlatPlates3       bool   `json:"flat_plates3" db:"flat_plates3"`
	FlatPlates4       bool   `json:"flat_plates4" db:"flat_plates4"`
}

type CompanyName struct {
	ID    uint32 `json:"id" db:"id"`
	Title string `json:"title" db:"title"`
}

type Store struct {
	CompanyID   uint32 `json:"companyID,omitempty" db:"company_id"`
	PlateTypeID uint8  `json:"plateTypeID" db:"plate_type_id"`
	Count       int64  `json:"count" db:"count"`
	Used        bool   `json:"used" db:"used"`
}

type Products struct {
	ID                                           uint32    `json:"id,omitempty" db:"id"`
	LicensePlate                                 string    `json:"license_plate,omitempty" db:"license_plate"`
	IssuedDoc                                    string    `json:"issued_doc" db:"issued_doc"`
	RepresentativeDocumentIssuedDoc              string    `json:"representative_document_issued_doc" db:"representative_document_issued_doc"`
	IndividualDocumentIssuedDoc                  string    `json:"individual_document_issued_doc" db:"individual_document_issued_doc"`
	SerialNumber                                 string    `json:"serial_number" db:"serial_number"`
	DateCreate                                   time.Time `json:"date_create" db:"date_create"`
	DateIssuing                                  time.Time `json:"date_issuing" db:"date_issuing"`
	Document                                     string    `json:"document" db:"document"`
	DateIssuingTransportDocument                 string    `json:"date_issuing_transport_document" db:"date_issuing_transport_document"`
	CountPlate                                   uint8     `json:"count_plate" db:"count_plate"`
	IndividualDocumentOtherType                  string    `json:"individual_document_other_type" db:"individual_document_other_type"`
	IndividualDocumentSurname                    string    `json:"individual_document_surname" db:"individual_document_surname"`
	IndividualDocumentName                       string    `json:"individual_document_name" db:"individual_document_name"`
	IndividualDocumentPatronymic                 string    `json:"individual_document_patronymic" db:"individual_document_patronymic"`
	IndividualDocumentSerialNumber               string    `json:"individual_document_serial_number" db:"individual_document_serial_number"`
	IndividualDocumentDateIssuing                string    `json:"individual_document_date_issuing" db:"individual_document_date_issuing"`
	EntityDocumentName                           string    `json:"entity_document_name" db:"entity_document_name"`
	EntityDocumentInn                            string    `json:"entity_document_inn" db:"entity_document_inn"`
	EntityDocumentOrgn                           string    `json:"entity_document_orgn" db:"entity_document_orgn"`
	RepresentativeTypeOwnerID                    uint8     `json:"representative_type_owner_id" db:"representative_type_owner_id"`
	RepresentativeIndividualDocumentOtherType    string    `json:"representative_individual_document_other_type" db:"representative_individual_document_other_type"`
	RepresentativeIndividualDocumentSurname      string    `json:"representative_individual_document_surname" db:"representative_individual_document_surname"`
	RepresentativeIndividualDocumentName         string    `json:"representative_individual_document_name" db:"representative_individual_document_name"`
	RepresentativeIndividualDocumentPatronymic   string    `json:"representative_individual_document_patronymic" db:"representative_individual_document_patronymic"`
	RepresentativeIndividualDocumentSerialNumber string    `json:"representative_individual_document_serial_number" db:"representative_individual_document_serial_number"`
	RepresentativeIndividualDocumentDateIssuing  string    `json:"representative_individual_document_date_issuing" db:"representative_individual_document_date_issuing"`
	RepresentativeEntityDocumentName             string    `json:"representative_entity_document_name" db:"representative_entity_document_name"`
	RepresentativeEntityDocumentInn              string    `json:"representative_entity_document_inn" db:"representative_entity_document_inn"`
	RepresentativeEntityDocumentOrgn             string    `json:"representative_entity_document_orgn" db:"representative_entity_document_orgn"`
	RepresentativeDocumentTypeOwnerID            uint8     `json:"representative_document_type_ownerId" db:"representative_document_type_ownerId"`
	CompanyID                                    uint32    `json:"companyId" db:"company_id"`
	PlateTypeID                                  uint8     `json:"plateTypeId" db:"plate_type_id"`
	TypeDocumentOwnerID                          *uint8    `json:"typeDocumentOwnerId" db:"type_document_owner_id"`
	TypeDocumentTransportID                      uint8     `json:"typeDocumentTransportId" db:"type_document_transport_id"`
	OwnerTypeID                                  uint8     `json:"ownerTypeId" db:"owner_type_id"`
	StatusProductID                              uint8     `json:"statusProductId" db:"status_product_id"`
}

type CompanyInfo struct {
	ID             uint32                  `json:"id" db:"id"`
	Title          string                  `json:"title" db:"title"`
	Address        string                  `json:"address" db:"address"`
	Comment        string                  `json:"comment" db:"comment"`
	LastAuths      string                  `json:"lastAuth" db:"last_auth"`
	AddAll         string                  `json:"addAll"`
	AuditAll       string                  `json:"auditAll"`
	UserCount      int                     `json:"userCount"`
	Store          map[string]StoreCompany `json:"store"`
	PlateInfoCheck struct {
		WaitCheck   uint32 `json:"waitCheck"`
		WaitRecheck uint32 `json:"waitRecheck"`
	} `json:"plateInfoCheck"`
	PlateWaitAudit struct {
		WaitAll uint32 `json:"waitAll"`
		WaitDay uint32 `json:"waitDay"`
	} `json:"plateWaitAudit"`
	PlateChecked struct {
		CheckedAll uint32 `json:"checkedAll"`
		CheckedDay uint32 `json:"checkedDay"`
	} `json:"plateChecked"`
	PlateDefect struct {
		DefectAll uint32 `json:"defectAll"`
		DefectDay uint32 `json:"defectDay"`
	} `json:"plateDefect"`
}

type StoreCompany struct {
	Count        uint32 `json:"count"`
	PlateAll     uint32 `json:"plateAll"`
	PlateLastDay uint32 `json:"plateLastDay"`
}

type DefectProduct struct {
	ID           uint32    `json:"id" db:"id"`
	LicensePlate string    `json:"licensePlate" db:"license_plate"`
	CountPlate   int32     `json:"countPlate" db:"count_plate"`
	DateCreate   time.Time `json:"dateCreate" db:"date_create"`
	CompanyId    int32     `json:"companyId" db:"company_id"`
	PlateTypeId  uint8     `json:"plateTypeId" db:"plate_type_id"`
}

type PlateTypes struct {
	ID          uint32 `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Mask        string `json:"mask" db:"mask"`
	Placeholder string `json:"placeholder" db:"placeholder"`
	Regexp      string `json:"regexp" db:"regexp"`
	ShortTitle  string `json:"shortTitle" db:"short_title"`
}
