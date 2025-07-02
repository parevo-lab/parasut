package parasut

import (
	"time"
)

// Account Hesap modeli
type Account struct {
	ID            string            `json:"id"`
	Type          string            `json:"type"`
	Attributes    AccountAttributes `json:"attributes"`
	Relationships interface{}       `json:"relationships,omitempty"`
}

// AccountAttributes Hesap nitelikleri
type AccountAttributes struct {
	UsedFor             string     `json:"used_for,omitempty"`
	LastUsedAt          *time.Time `json:"last_used_at,omitempty"`
	Balance             float64    `json:"balance,omitempty"`
	LastAdjustmentDate  *time.Time `json:"last_adjustment_date,omitempty"`
	BankIntegrationType string     `json:"bank_integration_type,omitempty"`
	AssociateEmail      string     `json:"associate_email,omitempty"`
	CreatedAt           *time.Time `json:"created_at,omitempty"`
	UpdatedAt           *time.Time `json:"updated_at,omitempty"`
	Name                string     `json:"name"`
	Currency            string     `json:"currency"`     // TRL, USD, EUR, GBP
	AccountType         string     `json:"account_type"` // cash, bank, sys
	BankName            string     `json:"bank_name,omitempty"`
	BankBranch          string     `json:"bank_branch,omitempty"`
	BankAccountNo       string     `json:"bank_account_no,omitempty"`
	IBAN                string     `json:"iban,omitempty"`
	Archived            bool       `json:"archived,omitempty"`
}

// AccountTransaction Hesap işlemi modeli
type AccountTransaction struct {
	ID            string                       `json:"id"`
	Type          string                       `json:"type"`
	Attributes    AccountTransactionAttributes `json:"attributes"`
	Relationships interface{}                  `json:"relationships,omitempty"`
}

// AccountTransactionAttributes Hesap işlemi nitelikleri
type AccountTransactionAttributes struct {
	Date        string     `json:"date"`
	Amount      float64    `json:"amount"`
	Description string     `json:"description,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

// BankFee Banka ücreti modeli
type BankFee struct {
	ID            string            `json:"id"`
	Type          string            `json:"type"`
	Attributes    BankFeeAttributes `json:"attributes"`
	Relationships interface{}       `json:"relationships,omitempty"`
}

// BankFeeAttributes Banka ücreti nitelikleri
type BankFeeAttributes struct {
	TotalPaid      float64    `json:"total_paid,omitempty"`
	Archived       bool       `json:"archived,omitempty"`
	Remaining      float64    `json:"remaining,omitempty"`
	RemainingInTRL float64    `json:"remaining_in_trl,omitempty"`
	CreatedAt      *time.Time `json:"created_at,omitempty"`
	UpdatedAt      *time.Time `json:"updated_at,omitempty"`
	Description    string     `json:"description"`
	Currency       string     `json:"currency"`   // TRL, USD, EUR, GBP
	IssueDate      string     `json:"issue_date"` // date format
	DueDate        string     `json:"due_date"`   // date format
	ExchangeRate   float64    `json:"exchange_rate,omitempty"`
	NetTotal       float64    `json:"net_total"`
}

// Contact Müşteri/Tedarikçi modeli
type Contact struct {
	ID            string            `json:"id"`
	Type          string            `json:"type"`
	Attributes    ContactAttributes `json:"attributes"`
	Relationships interface{}       `json:"relationships,omitempty"`
}

// ContactAttributes Müşteri/Tedarikçi nitelikleri
type ContactAttributes struct {
	Email              string     `json:"email,omitempty"`
	Name               string     `json:"name"`
	ShortName          string     `json:"short_name,omitempty"`
	ContactType        string     `json:"contact_type"` // person, company
	TaxNumber          string     `json:"tax_number,omitempty"`
	TaxOffice          string     `json:"tax_office,omitempty"`
	District           string     `json:"district,omitempty"`
	City               string     `json:"city,omitempty"`
	Address            string     `json:"address,omitempty"`
	Phone              string     `json:"phone,omitempty"`
	Fax                string     `json:"fax,omitempty"`
	IsAbroad           bool       `json:"is_abroad,omitempty"`
	Archived           bool       `json:"archived,omitempty"`
	UntrackableBalance float64    `json:"untrackable_balance,omitempty"`
	CreatedAt          *time.Time `json:"created_at,omitempty"`
	UpdatedAt          *time.Time `json:"updated_at,omitempty"`
	AccountType        string     `json:"account_type,omitempty"` // customer, supplier, both
}

// ContactTransaction Müşteri/Tedarikçi işlemi modeli
type ContactTransaction struct {
	ID            string                       `json:"id"`
	Type          string                       `json:"type"`
	Attributes    ContactTransactionAttributes `json:"attributes"`
	Relationships interface{}                  `json:"relationships,omitempty"`
}

// ContactTransactionAttributes Müşteri/Tedarikçi işlemi nitelikleri
type ContactTransactionAttributes struct {
	Date        string     `json:"date"`
	Amount      float64    `json:"amount"`
	Description string     `json:"description,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

// EArchive E-Arşiv modeli
type EArchive struct {
	ID            string             `json:"id"`
	Type          string             `json:"type"`
	Attributes    EArchiveAttributes `json:"attributes"`
	Relationships interface{}        `json:"relationships,omitempty"`
}

// EArchiveAttributes E-Arşiv nitelikleri
type EArchiveAttributes struct {
	VatWithholdingCode     string     `json:"vat_withholding_code,omitempty"`
	VatExemptionReasonCode string     `json:"vat_exemption_reason_code,omitempty"`
	VatExemptionReason     string     `json:"vat_exemption_reason,omitempty"`
	Note                   string     `json:"note,omitempty"`
	ExciseDutyCodes        []string   `json:"excise_duty_codes,omitempty"`
	InternetSale           bool       `json:"internet_sale,omitempty"`
	Shipment               bool       `json:"shipment,omitempty"`
	CreatedAt              *time.Time `json:"created_at,omitempty"`
	UpdatedAt              *time.Time `json:"updated_at,omitempty"`
}

// EInvoiceInbox E-Fatura gelen kutusu modeli
type EInvoiceInbox struct {
	ID            string                  `json:"id"`
	Type          string                  `json:"type"`
	Attributes    EInvoiceInboxAttributes `json:"attributes"`
	Relationships interface{}             `json:"relationships,omitempty"`
}

// EInvoiceInboxAttributes E-Fatura gelen kutusu nitelikleri
type EInvoiceInboxAttributes struct {
	VKN         string     `json:"vkn,omitempty"`
	InvoiceUUID string     `json:"invoice_uuid,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

// EInvoice E-Fatura modeli
type EInvoice struct {
	ID            string             `json:"id"`
	Type          string             `json:"type"`
	Attributes    EInvoiceAttributes `json:"attributes"`
	Relationships interface{}        `json:"relationships,omitempty"`
}

// EInvoiceAttributes E-Fatura nitelikleri
type EInvoiceAttributes struct {
	VatWithholdingCode     string     `json:"vat_withholding_code,omitempty"`
	VatExemptionReasonCode string     `json:"vat_exemption_reason_code,omitempty"`
	VatExemptionReason     string     `json:"vat_exemption_reason,omitempty"`
	Note                   string     `json:"note,omitempty"`
	ExciseDutyCodes        []string   `json:"excise_duty_codes,omitempty"`
	InternetSale           bool       `json:"internet_sale,omitempty"`
	Shipment               bool       `json:"shipment,omitempty"`
	CreatedAt              *time.Time `json:"created_at,omitempty"`
	UpdatedAt              *time.Time `json:"updated_at,omitempty"`
}

// ESMM E-SMM modeli
type ESMM struct {
	ID            string         `json:"id"`
	Type          string         `json:"type"`
	Attributes    ESMMAttributes `json:"attributes"`
	Relationships interface{}    `json:"relationships,omitempty"`
}

// ESMMAttributes E-SMM nitelikleri
type ESMMAttributes struct {
	VatWithholdingCode     string     `json:"vat_withholding_code,omitempty"`
	VatExemptionReasonCode string     `json:"vat_exemption_reason_code,omitempty"`
	VatExemptionReason     string     `json:"vat_exemption_reason,omitempty"`
	Note                   string     `json:"note,omitempty"`
	ExciseDutyCodes        []string   `json:"excise_duty_codes,omitempty"`
	CreatedAt              *time.Time `json:"created_at,omitempty"`
	UpdatedAt              *time.Time `json:"updated_at,omitempty"`
}

// Employee Çalışan modeli
type Employee struct {
	ID            string             `json:"id"`
	Type          string             `json:"type"`
	Attributes    EmployeeAttributes `json:"attributes"`
	Relationships interface{}        `json:"relationships,omitempty"`
}

// EmployeeAttributes Çalışan nitelikleri
type EmployeeAttributes struct {
	Name      string     `json:"name"`
	Email     string     `json:"email,omitempty"`
	IBAN      string     `json:"iban,omitempty"`
	Archived  bool       `json:"archived,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// ItemCategory Ürün kategorisi modeli
type ItemCategory struct {
	ID            string                 `json:"id"`
	Type          string                 `json:"type"`
	Attributes    ItemCategoryAttributes `json:"attributes"`
	Relationships interface{}            `json:"relationships,omitempty"`
}

// ItemCategoryAttributes Ürün kategorisi nitelikleri
type ItemCategoryAttributes struct {
	Name      string     `json:"name"`
	BgColor   string     `json:"bg_color,omitempty"`
	TextColor string     `json:"text_color,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// Product Ürün modeli
type Product struct {
	ID            string            `json:"id"`
	Type          string            `json:"type"`
	Attributes    ProductAttributes `json:"attributes"`
	Relationships interface{}       `json:"relationships,omitempty"`
}

// ProductAttributes Ürün nitelikleri
type ProductAttributes struct {
	Code                   string     `json:"code"`
	Name                   string     `json:"name"`
	VatRate                float64    `json:"vat_rate,omitempty"`
	SalesExciseDutyRate    float64    `json:"sales_excise_duty_rate,omitempty"`
	PurchaseExciseDutyRate float64    `json:"purchase_excise_duty_rate,omitempty"`
	Unit                   string     `json:"unit,omitempty"`
	CommunicationsTaxRate  float64    `json:"communications_tax_rate,omitempty"`
	Archived               bool       `json:"archived,omitempty"`
	ListPrice              float64    `json:"list_price,omitempty"`
	Currency               string     `json:"currency,omitempty"`
	BuyingPrice            float64    `json:"buying_price,omitempty"`
	BuyingCurrency         string     `json:"buying_currency,omitempty"`
	InventoryTracking      bool       `json:"inventory_tracking,omitempty"`
	InitialStockCount      float64    `json:"initial_stock_count,omitempty"`
	CreatedAt              *time.Time `json:"created_at,omitempty"`
	UpdatedAt              *time.Time `json:"updated_at,omitempty"`
}

// InventoryLevel Stok seviyesi modeli
type InventoryLevel struct {
	ID            string                   `json:"id"`
	Type          string                   `json:"type"`
	Attributes    InventoryLevelAttributes `json:"attributes"`
	Relationships interface{}              `json:"relationships,omitempty"`
}

// InventoryLevelAttributes Stok seviyesi nitelikleri
type InventoryLevelAttributes struct {
	StockCount float64    `json:"stock_count"`
	CreatedAt  *time.Time `json:"created_at,omitempty"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty"`
}

// SalesOffer Satış teklifi modeli
type SalesOffer struct {
	ID            string                  `json:"id"`
	Type          string                  `json:"type"`
	Attributes    SalesOfferAttributes    `json:"attributes"`
	Relationships SalesOfferRelationships `json:"relationships,omitempty"`
}

// SalesOfferAttributes Satış teklifi nitelikleri
type SalesOfferAttributes struct {
	Archived               bool       `json:"archived,omitempty"`
	NetTotal               float64    `json:"net_total,omitempty"`
	GrossTotal             float64    `json:"gross_total,omitempty"`
	TotalExciseDuty        float64    `json:"total_excise_duty,omitempty"`
	TotalCommunicationsTax float64    `json:"total_communications_tax,omitempty"`
	TotalVat               float64    `json:"total_vat,omitempty"`
	TotalDiscount          float64    `json:"total_discount,omitempty"`
	TotalInvoiceDiscount   float64    `json:"total_invoice_discount,omitempty"`
	BeforeTaxesTotal       float64    `json:"before_taxes_total,omitempty"`
	CreatedAt              *time.Time `json:"created_at,omitempty"`
	UpdatedAt              *time.Time `json:"updated_at,omitempty"`
	Description            string     `json:"description,omitempty"`
	IssueDate              string     `json:"issue_date"` // date format
	Currency               string     `json:"currency,omitempty"`
	ExchangeRate           float64    `json:"exchange_rate,omitempty"`
	InvoiceDiscountType    string     `json:"invoice_discount_type,omitempty"`
	InvoiceDiscount        float64    `json:"invoice_discount,omitempty"`
	Status                 string     `json:"status,omitempty"`
}

// SalesOfferRelationships Satış teklifi ilişkileri
type SalesOfferRelationships struct {
	Contact *RelationshipData  `json:"contact,omitempty"`
	Details []RelationshipData `json:"details,omitempty"`
	Tags    []RelationshipData `json:"tags,omitempty"`
}

// SalesInvoice Satış faturası modeli
type SalesInvoice struct {
	ID            string                    `json:"id"`
	Type          string                    `json:"type"`
	Attributes    SalesInvoiceAttributes    `json:"attributes"`
	Relationships SalesInvoiceRelationships `json:"relationships,omitempty"`
}

// SalesInvoiceAttributes Satış faturası nitelikleri
type SalesInvoiceAttributes struct {
	Archived               bool       `json:"archived,omitempty"`
	NetTotal               float64    `json:"net_total,omitempty"`
	GrossTotal             float64    `json:"gross_total,omitempty"`
	Withholding            float64    `json:"withholding,omitempty"`
	TotalExciseDuty        float64    `json:"total_excise_duty,omitempty"`
	TotalCommunicationsTax float64    `json:"total_communications_tax,omitempty"`
	TotalVat               float64    `json:"total_vat,omitempty"`
	VatWithholding         float64    `json:"vat_withholding,omitempty"`
	TotalDiscount          float64    `json:"total_discount,omitempty"`
	TotalInvoiceDiscount   float64    `json:"total_invoice_discount,omitempty"`
	BeforeTaxesTotal       float64    `json:"before_taxes_total,omitempty"`
	Remaining              float64    `json:"remaining,omitempty"`
	RemainingInTRL         float64    `json:"remaining_in_trl,omitempty"`
	PaymentStatus          string     `json:"payment_status,omitempty"`
	CreatedAt              *time.Time `json:"created_at,omitempty"`
	UpdatedAt              *time.Time `json:"updated_at,omitempty"`
	ItemType               string     `json:"item_type"` // invoice, estimate, cancelled, recurring_invoice, recurring_estimate, refund
	Description            string     `json:"description,omitempty"`
	IssueDate              string     `json:"issue_date"`         // date format
	DueDate                string     `json:"due_date,omitempty"` // date format
	InvoiceSeries          string     `json:"invoice_series,omitempty"`
	InvoiceID              int        `json:"invoice_id,omitempty"`
	Currency               string     `json:"currency,omitempty"`
	ExchangeRate           float64    `json:"exchange_rate,omitempty"`
	WithholdingRate        float64    `json:"withholding_rate,omitempty"`
	VatWithholdingRate     float64    `json:"vat_withholding_rate,omitempty"`
	InvoiceDiscountType    string     `json:"invoice_discount_type,omitempty"`
	InvoiceDiscount        float64    `json:"invoice_discount,omitempty"`
	BillingAddress         string     `json:"billing_address,omitempty"`
	BillingPhone           string     `json:"billing_phone,omitempty"`
	BillingFax             string     `json:"billing_fax,omitempty"`
	TaxOffice              string     `json:"tax_office,omitempty"`
	TaxNumber              string     `json:"tax_number,omitempty"`
	Country                string     `json:"country,omitempty"`
	City                   string     `json:"city,omitempty"`
	District               string     `json:"district,omitempty"`
	IsAbroad               bool       `json:"is_abroad,omitempty"`
	OrderNo                string     `json:"order_no,omitempty"`
	OrderDate              string     `json:"order_date,omitempty"`
}

// SalesInvoiceRelationships Satış faturası ilişkileri
type SalesInvoiceRelationships struct {
	Contact         *RelationshipData  `json:"contact,omitempty"`
	Details         []RelationshipData `json:"details,omitempty"`
	Payments        []RelationshipData `json:"payments,omitempty"`
	Tags            []RelationshipData `json:"tags,omitempty"`
	Sharings        []RelationshipData `json:"sharings,omitempty"`
	RecurrencePlan  *RelationshipData  `json:"recurrence_plan,omitempty"`
	ActiveEDocument *RelationshipData  `json:"active_e_document,omitempty"`
}

// RelationshipData İlişki verisi
type RelationshipData struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

// PurchaseBill Alış faturası modeli
type PurchaseBill struct {
	ID            string                    `json:"id"`
	Type          string                    `json:"type"`
	Attributes    PurchaseBillAttributes    `json:"attributes"`
	Relationships PurchaseBillRelationships `json:"relationships,omitempty"`
}

// PurchaseBillAttributes Alış faturası nitelikleri
type PurchaseBillAttributes struct {
	Archived               bool       `json:"archived,omitempty"`
	NetTotal               float64    `json:"net_total,omitempty"`
	GrossTotal             float64    `json:"gross_total,omitempty"`
	Withholding            float64    `json:"withholding,omitempty"`
	TotalExciseDuty        float64    `json:"total_excise_duty,omitempty"`
	TotalCommunicationsTax float64    `json:"total_communications_tax,omitempty"`
	TotalVat               float64    `json:"total_vat,omitempty"`
	VatWithholding         float64    `json:"vat_withholding,omitempty"`
	TotalDiscount          float64    `json:"total_discount,omitempty"`
	TotalInvoiceDiscount   float64    `json:"total_invoice_discount,omitempty"`
	BeforeTaxesTotal       float64    `json:"before_taxes_total,omitempty"`
	Remaining              float64    `json:"remaining,omitempty"`
	RemainingInTRL         float64    `json:"remaining_in_trl,omitempty"`
	PaymentStatus          string     `json:"payment_status,omitempty"`
	CreatedAt              *time.Time `json:"created_at,omitempty"`
	UpdatedAt              *time.Time `json:"updated_at,omitempty"`
	ItemType               string     `json:"item_type"` // bill, cancelled
	Description            string     `json:"description,omitempty"`
	IssueDate              string     `json:"issue_date"`         // date format
	DueDate                string     `json:"due_date,omitempty"` // date format
	InvoiceSeries          string     `json:"invoice_series,omitempty"`
	InvoiceID              string     `json:"invoice_id,omitempty"`
	Currency               string     `json:"currency,omitempty"`
	ExchangeRate           float64    `json:"exchange_rate,omitempty"`
	WithholdingRate        float64    `json:"withholding_rate,omitempty"`
	VatWithholdingRate     float64    `json:"vat_withholding_rate,omitempty"`
	InvoiceDiscountType    string     `json:"invoice_discount_type,omitempty"`
	InvoiceDiscount        float64    `json:"invoice_discount,omitempty"`
	BillingAddress         string     `json:"billing_address,omitempty"`
	BillingPhone           string     `json:"billing_phone,omitempty"`
	BillingFax             string     `json:"billing_fax,omitempty"`
	TaxOffice              string     `json:"tax_office,omitempty"`
	TaxNumber              string     `json:"tax_number,omitempty"`
	SupplierName           string     `json:"supplier_name,omitempty"`
	SupplierTaxNumber      string     `json:"supplier_tax_number,omitempty"`
	SupplierTaxOffice      string     `json:"supplier_tax_office,omitempty"`
}

// PurchaseBillRelationships Alış faturası ilişkileri
type PurchaseBillRelationships struct {
	Supplier *RelationshipData  `json:"supplier,omitempty"`
	Details  []RelationshipData `json:"details,omitempty"`
	Payments []RelationshipData `json:"payments,omitempty"`
	Tags     []RelationshipData `json:"tags,omitempty"`
}

// Salary Maaş modeli
type Salary struct {
	ID            string              `json:"id"`
	Type          string              `json:"type"`
	Attributes    SalaryAttributes    `json:"attributes"`
	Relationships SalaryRelationships `json:"relationships,omitempty"`
}

// SalaryAttributes Maaş nitelikleri
type SalaryAttributes struct {
	Archived     bool       `json:"archived,omitempty"`
	NetTotal     float64    `json:"net_total,omitempty"`
	GrossTotal   float64    `json:"gross_total,omitempty"`
	CreatedAt    *time.Time `json:"created_at,omitempty"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty"`
	Date         string     `json:"date"` // date format
	Description  string     `json:"description,omitempty"`
	Currency     string     `json:"currency,omitempty"`
	ExchangeRate float64    `json:"exchange_rate,omitempty"`
}

// SalaryRelationships Maaş ilişkileri
type SalaryRelationships struct {
	Employee *RelationshipData  `json:"employee,omitempty"`
	Payments []RelationshipData `json:"payments,omitempty"`
	Tags     []RelationshipData `json:"tags,omitempty"`
}

// Sharing Paylaşım modeli
type Sharing struct {
	ID            string            `json:"id"`
	Type          string            `json:"type"`
	Attributes    SharingAttributes `json:"attributes"`
	Relationships interface{}       `json:"relationships,omitempty"`
}

// SharingAttributes Paylaşım nitelikleri
type SharingAttributes struct {
	Name      string     `json:"name"`
	Token     string     `json:"token,omitempty"`
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// ShipmentDocument Sevkiyat belgesi modeli
type ShipmentDocument struct {
	ID            string                     `json:"id"`
	Type          string                     `json:"type"`
	Attributes    ShipmentDocumentAttributes `json:"attributes"`
	Relationships interface{}                `json:"relationships,omitempty"`
}

// ShipmentDocumentAttributes Sevkiyat belgesi nitelikleri
type ShipmentDocumentAttributes struct {
	ShipmentDate     string     `json:"shipment_date"`
	Address          string     `json:"address,omitempty"`
	ShipmentIncluded bool       `json:"shipment_included,omitempty"`
	CreatedAt        *time.Time `json:"created_at,omitempty"`
	UpdatedAt        *time.Time `json:"updated_at,omitempty"`
}

// StockUpdate Stok güncelleme modeli
type StockUpdate struct {
	ID            string                `json:"id"`
	Type          string                `json:"type"`
	Attributes    StockUpdateAttributes `json:"attributes"`
	Relationships interface{}           `json:"relationships,omitempty"`
}

// StockUpdateAttributes Stok güncelleme nitelikleri
type StockUpdateAttributes struct {
	Date        string     `json:"date"`
	StockCount  float64    `json:"stock_count"`
	UnitCost    float64    `json:"unit_cost,omitempty"`
	Description string     `json:"description,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

// Webhook Webhook modeli
type Webhook struct {
	ID            string            `json:"id"`
	Type          string            `json:"type"`
	Attributes    WebhookAttributes `json:"attributes"`
	Relationships interface{}       `json:"relationships,omitempty"`
}

// WebhookAttributes Webhook nitelikleri
type WebhookAttributes struct {
	URL           string     `json:"url"`
	Event         string     `json:"event"`
	IsActive      bool       `json:"is_active,omitempty"`
	EncryptionKey string     `json:"encryption_key,omitempty"`
	CreatedAt     *time.Time `json:"created_at,omitempty"`
	UpdatedAt     *time.Time `json:"updated_at,omitempty"`
}

// Tag Etiket modeli
type Tag struct {
	ID            string        `json:"id"`
	Type          string        `json:"type"`
	Attributes    TagAttributes `json:"attributes"`
	Relationships interface{}   `json:"relationships,omitempty"`
}

// TagAttributes Etiket nitelikleri
type TagAttributes struct {
	Name      string     `json:"name"`
	Color     string     `json:"color,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// Tax Vergi modeli
type Tax struct {
	ID            string           `json:"id"`
	Type          string           `json:"type"`
	Attributes    TaxAttributes    `json:"attributes"`
	Relationships TaxRelationships `json:"relationships,omitempty"`
}

// TaxAttributes Vergi nitelikleri
type TaxAttributes struct {
	Archived     bool       `json:"archived,omitempty"`
	NetTotal     float64    `json:"net_total,omitempty"`
	GrossTotal   float64    `json:"gross_total,omitempty"`
	CreatedAt    *time.Time `json:"created_at,omitempty"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty"`
	Date         string     `json:"date"` // date format
	Description  string     `json:"description,omitempty"`
	Currency     string     `json:"currency,omitempty"`
	ExchangeRate float64    `json:"exchange_rate,omitempty"`
}

// TaxRelationships Vergi ilişkileri
type TaxRelationships struct {
	Tags     []RelationshipData `json:"tags,omitempty"`
	Payments []RelationshipData `json:"payments,omitempty"`
}

// TrackableJob İzlenebilir iş modeli
type TrackableJob struct {
	ID            string                 `json:"id"`
	Type          string                 `json:"type"`
	Attributes    TrackableJobAttributes `json:"attributes"`
	Relationships interface{}            `json:"relationships,omitempty"`
}

// TrackableJobAttributes İzlenebilir iş nitelikleri
type TrackableJobAttributes struct {
	Status      string     `json:"status,omitempty"`
	Errors      []string   `json:"errors,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
}

// Transaction İşlem modeli
type Transaction struct {
	ID            string                `json:"id"`
	Type          string                `json:"type"`
	Attributes    TransactionAttributes `json:"attributes"`
	Relationships interface{}           `json:"relationships,omitempty"`
}

// TransactionAttributes İşlem nitelikleri
type TransactionAttributes struct {
	Date        string     `json:"date"`
	Amount      float64    `json:"amount"`
	Description string     `json:"description,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

// StockMovement Stok hareketi modeli
type StockMovement struct {
	ID            string                  `json:"id"`
	Type          string                  `json:"type"`
	Attributes    StockMovementAttributes `json:"attributes"`
	Relationships interface{}             `json:"relationships,omitempty"`
}

// StockMovementAttributes Stok hareketi nitelikleri
type StockMovementAttributes struct {
	Date         string     `json:"date"`          // date format
	MovementType string     `json:"movement_type"` // in, out
	Quantity     float64    `json:"quantity"`
	Description  string     `json:"description,omitempty"`
	CreatedAt    *time.Time `json:"created_at,omitempty"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty"`
}

// Warehouse Depo modeli
type Warehouse struct {
	ID            string              `json:"id"`
	Type          string              `json:"type"`
	Attributes    WarehouseAttributes `json:"attributes"`
	Relationships interface{}         `json:"relationships,omitempty"`
}

// WarehouseAttributes Depo nitelikleri
type WarehouseAttributes struct {
	Name      string     `json:"name"`
	City      string     `json:"city,omitempty"`
	District  string     `json:"district,omitempty"`
	Address   string     `json:"address,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// Payment Ödeme modeli
type Payment struct {
	ID            string            `json:"id"`
	Type          string            `json:"type"`
	Attributes    PaymentAttributes `json:"attributes"`
	Relationships interface{}       `json:"relationships,omitempty"`
}

// PaymentAttributes Ödeme nitelikleri
type PaymentAttributes struct {
	Date        string     `json:"date"`
	Amount      float64    `json:"amount"`
	Description string     `json:"description,omitempty"`
	Currency    string     `json:"currency,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

// Me represents current user information
type Me struct {
	ID            string          `json:"id"`
	Type          string          `json:"type"`
	Attributes    MeAttributes    `json:"attributes"`
	Relationships MeRelationships `json:"relationships,omitempty"`
}

type MeAttributes struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	IsConfirmed bool   `json:"is_confirmed"`
}

type MeRelationships struct {
	UserRoles *Relationship     `json:"user_roles,omitempty"`
	Companies *Relationship     `json:"companies,omitempty"`
	Profile   *RelationshipData `json:"profile,omitempty"`
}

type Relationship struct {
	Data []RelationshipData `json:"data,omitempty"`
}
