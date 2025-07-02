package parasut

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// Generic helper metodları - Kod tekrarını önlemek için

// list generic list metodu
func list[T any](c *Client, ctx context.Context, endpoint string, params *ListParams) ([]T, *Meta, error) {
	var queryParams map[string]string
	if params != nil {
		queryParams = params.ToMap()
	}

	resp, err := c.get(ctx, endpoint, queryParams)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	var response struct {
		Data []T   `json:"data"`
		Meta *Meta `json:"meta"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, nil, err
	}

	return response.Data, response.Meta, nil
}

// get generic get metodu
func get[T any](c *Client, ctx context.Context, endpoint string) (*T, error) {
	resp, err := c.get(ctx, endpoint, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response struct {
		Data T `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response.Data, nil
}

// create generic create metodu
func create[T any](c *Client, ctx context.Context, endpoint, resourceType string, attributes interface{}, relationships interface{}) (*T, error) {
	body := map[string]interface{}{
		"data": map[string]interface{}{
			"type":       resourceType,
			"attributes": attributes,
		},
	}

	if relationships != nil {
		body["data"].(map[string]interface{})["relationships"] = relationships
	}

	resp, err := c.post(ctx, endpoint, body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response struct {
		Data T `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response.Data, nil
}

// update generic update metodu
func update[T any](c *Client, ctx context.Context, endpoint, id, resourceType string, attributes interface{}, relationships interface{}) (*T, error) {
	body := map[string]interface{}{
		"data": map[string]interface{}{
			"id":         id,
			"type":       resourceType,
			"attributes": attributes,
		},
	}

	if relationships != nil {
		body["data"].(map[string]interface{})["relationships"] = relationships
	}

	resp, err := c.put(ctx, endpoint, body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response struct {
		Data T `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response.Data, nil
}

// action generic action metodu (patch)
func action[T any](c *Client, ctx context.Context, endpoint string, body interface{}) (*T, error) {
	resp, err := c.patch(ctx, endpoint, body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if body == nil {
		// Sadece action, response döndürmeyen
		return nil, nil
	}

	var response struct {
		Data T `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response.Data, nil
}

// deleteResource generic delete metodu
func deleteResource(c *Client, ctx context.Context, endpoint string) error {
	resp, err := c.delete(ctx, endpoint)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

// getPDF generic PDF metodu
func getPDF(c *Client, ctx context.Context, endpoint string) ([]byte, error) {
	resp, err := c.get(ctx, endpoint, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var pdfData []byte
	if err := json.NewDecoder(resp.Body).Decode(&pdfData); err != nil {
		return nil, err
	}

	return pdfData, nil
}

// Archive action
func archive(c *Client, ctx context.Context, endpoint string) error {
	_, err := action[interface{}](c, ctx, endpoint+"/archive", nil)
	return err
}

// Unarchive action
func unarchive(c *Client, ctx context.Context, endpoint string) error {
	_, err := action[interface{}](c, ctx, endpoint+"/unarchive", nil)
	return err
}

// Cancel action
func cancel(c *Client, ctx context.Context, endpoint string) error {
	_, err := action[interface{}](c, ctx, endpoint+"/cancel", nil)
	return err
}

// Recover action
func recover(c *Client, ctx context.Context, endpoint string) error {
	_, err := action[interface{}](c, ctx, endpoint+"/recover", nil)
	return err
}

// CreatePayment generic payment metodu
func createPayment(c *Client, ctx context.Context, endpoint string, attributes PaymentAttributes) (*Payment, error) {
	return create[Payment](c, ctx, endpoint+"/payments", "payments", attributes, nil)
}

// SERVIS TANIMLARI

// MeService Kullanıcı bilgileri servisi
type MeService struct {
	client *Client
}

func (s *MeService) Get(ctx context.Context, includeParams ...string) (*Me, error) {
	// /me endpoint doesn't use company_id, so we need to make request differently
	endpoint := "/me"
	if len(includeParams) > 0 {
		endpoint += "?include=" + includeParams[0]
	}

	// Direct request to base URL without company_id
	url := fmt.Sprintf("https://api.parasut.com/v4%s", endpoint)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := s.client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response struct {
		Data Me `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response.Data, nil
}

// AccountsService Hesaplar servisi
type AccountsService struct {
	client *Client
}

func (s *AccountsService) List(ctx context.Context, params *ListParams) ([]Account, *Meta, error) {
	return list[Account](s.client, ctx, "/accounts", params)
}

func (s *AccountsService) Get(ctx context.Context, id string) (*Account, error) {
	return get[Account](s.client, ctx, fmt.Sprintf("/accounts/%s", id))
}

func (s *AccountsService) Create(ctx context.Context, attributes AccountAttributes) (*Account, error) {
	return create[Account](s.client, ctx, "/accounts", "accounts", attributes, nil)
}

func (s *AccountsService) Update(ctx context.Context, id string, attributes AccountAttributes) (*Account, error) {
	return update[Account](s.client, ctx, fmt.Sprintf("/accounts/%s", id), id, "accounts", attributes, nil)
}

func (s *AccountsService) Delete(ctx context.Context, id string) error {
	return deleteResource(s.client, ctx, fmt.Sprintf("/accounts/%s", id))
}

func (s *AccountsService) GetTransactions(ctx context.Context, accountID string) ([]AccountTransaction, *Meta, error) {
	return list[AccountTransaction](s.client, ctx, fmt.Sprintf("/accounts/%s/transactions", accountID), nil)
}

func (s *AccountsService) CreateDebitTransaction(ctx context.Context, accountID string, attributes AccountTransactionAttributes) (*AccountTransaction, error) {
	return create[AccountTransaction](s.client, ctx, fmt.Sprintf("/accounts/%s/debit_transactions", accountID), "account_transactions", attributes, nil)
}

func (s *AccountsService) CreateCreditTransaction(ctx context.Context, accountID string, attributes AccountTransactionAttributes) (*AccountTransaction, error) {
	return create[AccountTransaction](s.client, ctx, fmt.Sprintf("/accounts/%s/credit_transactions", accountID), "account_transactions", attributes, nil)
}

// BankFeesService Banka ücretleri servisi
type BankFeesService struct {
	client *Client
}

func (s *BankFeesService) List(ctx context.Context, params *ListParams) ([]BankFee, *Meta, error) {
	return list[BankFee](s.client, ctx, "/bank_fees", params)
}

func (s *BankFeesService) Get(ctx context.Context, id string) (*BankFee, error) {
	return get[BankFee](s.client, ctx, fmt.Sprintf("/bank_fees/%s", id))
}

func (s *BankFeesService) Create(ctx context.Context, attributes BankFeeAttributes) (*BankFee, error) {
	return create[BankFee](s.client, ctx, "/bank_fees", "bank_fees", attributes, nil)
}

func (s *BankFeesService) Update(ctx context.Context, id string, attributes BankFeeAttributes) (*BankFee, error) {
	return update[BankFee](s.client, ctx, fmt.Sprintf("/bank_fees/%s", id), id, "bank_fees", attributes, nil)
}

func (s *BankFeesService) Archive(ctx context.Context, id string) error {
	return archive(s.client, ctx, fmt.Sprintf("/bank_fees/%s", id))
}

func (s *BankFeesService) Unarchive(ctx context.Context, id string) error {
	return unarchive(s.client, ctx, fmt.Sprintf("/bank_fees/%s", id))
}

func (s *BankFeesService) CreatePayment(ctx context.Context, bankFeeID string, attributes PaymentAttributes) (*Payment, error) {
	return createPayment(s.client, ctx, fmt.Sprintf("/bank_fees/%s", bankFeeID), attributes)
}

// ContactsService Müşteri/Tedarikçi servisi
type ContactsService struct {
	client *Client
}

func (s *ContactsService) List(ctx context.Context, params *ListParams) ([]Contact, *Meta, error) {
	return list[Contact](s.client, ctx, "/contacts", params)
}

func (s *ContactsService) Get(ctx context.Context, id string) (*Contact, error) {
	return get[Contact](s.client, ctx, fmt.Sprintf("/contacts/%s", id))
}

func (s *ContactsService) Create(ctx context.Context, attributes ContactAttributes) (*Contact, error) {
	return create[Contact](s.client, ctx, "/contacts", "contacts", attributes, nil)
}

func (s *ContactsService) Update(ctx context.Context, id string, attributes ContactAttributes) (*Contact, error) {
	return update[Contact](s.client, ctx, fmt.Sprintf("/contacts/%s", id), id, "contacts", attributes, nil)
}

func (s *ContactsService) Delete(ctx context.Context, id string) error {
	return deleteResource(s.client, ctx, fmt.Sprintf("/contacts/%s", id))
}

func (s *ContactsService) GetDebitTransactions(ctx context.Context, contactID string) ([]ContactTransaction, *Meta, error) {
	return list[ContactTransaction](s.client, ctx, fmt.Sprintf("/contacts/%s/contact_debit_transactions", contactID), nil)
}

func (s *ContactsService) GetCreditTransactions(ctx context.Context, contactID string) ([]ContactTransaction, *Meta, error) {
	return list[ContactTransaction](s.client, ctx, fmt.Sprintf("/contacts/%s/contact_credit_transactions", contactID), nil)
}

// ProductsService Ürünler servisi
type ProductsService struct {
	client *Client
}

func (s *ProductsService) List(ctx context.Context, params *ListParams) ([]Product, *Meta, error) {
	return list[Product](s.client, ctx, "/products", params)
}

func (s *ProductsService) Get(ctx context.Context, id string) (*Product, error) {
	return get[Product](s.client, ctx, fmt.Sprintf("/products/%s", id))
}

func (s *ProductsService) Create(ctx context.Context, attributes ProductAttributes) (*Product, error) {
	return create[Product](s.client, ctx, "/products", "products", attributes, nil)
}

func (s *ProductsService) Update(ctx context.Context, id string, attributes ProductAttributes) (*Product, error) {
	return update[Product](s.client, ctx, fmt.Sprintf("/products/%s", id), id, "products", attributes, nil)
}

func (s *ProductsService) Delete(ctx context.Context, id string) error {
	return deleteResource(s.client, ctx, fmt.Sprintf("/products/%s", id))
}

func (s *ProductsService) GetInventoryLevels(ctx context.Context, productID string) ([]InventoryLevel, *Meta, error) {
	return list[InventoryLevel](s.client, ctx, fmt.Sprintf("/products/%s/inventory_levels", productID), nil)
}

// SalesInvoicesService Satış faturaları servisi
type SalesInvoicesService struct {
	client *Client
}

func (s *SalesInvoicesService) List(ctx context.Context, params *ListParams) ([]SalesInvoice, *Meta, error) {
	return list[SalesInvoice](s.client, ctx, "/sales_invoices", params)
}

func (s *SalesInvoicesService) Get(ctx context.Context, id string) (*SalesInvoice, error) {
	return get[SalesInvoice](s.client, ctx, fmt.Sprintf("/sales_invoices/%s", id))
}

func (s *SalesInvoicesService) Create(ctx context.Context, attributes SalesInvoiceAttributes, relationships *SalesInvoiceRelationships) (*SalesInvoice, error) {
	return create[SalesInvoice](s.client, ctx, "/sales_invoices", "sales_invoices", attributes, relationships)
}

func (s *SalesInvoicesService) Update(ctx context.Context, id string, attributes SalesInvoiceAttributes, relationships *SalesInvoiceRelationships) (*SalesInvoice, error) {
	return update[SalesInvoice](s.client, ctx, fmt.Sprintf("/sales_invoices/%s", id), id, "sales_invoices", attributes, relationships)
}

func (s *SalesInvoicesService) Cancel(ctx context.Context, id string) error {
	return cancel(s.client, ctx, fmt.Sprintf("/sales_invoices/%s", id))
}

func (s *SalesInvoicesService) Recover(ctx context.Context, id string) error {
	return recover(s.client, ctx, fmt.Sprintf("/sales_invoices/%s", id))
}

func (s *SalesInvoicesService) Archive(ctx context.Context, id string) error {
	return archive(s.client, ctx, fmt.Sprintf("/sales_invoices/%s", id))
}

func (s *SalesInvoicesService) Unarchive(ctx context.Context, id string) error {
	return unarchive(s.client, ctx, fmt.Sprintf("/sales_invoices/%s", id))
}

func (s *SalesInvoicesService) CreatePayment(ctx context.Context, invoiceID string, attributes PaymentAttributes) (*Payment, error) {
	return createPayment(s.client, ctx, fmt.Sprintf("/sales_invoices/%s", invoiceID), attributes)
}

func (s *SalesInvoicesService) ConvertToInvoice(ctx context.Context, invoiceID string) (*SalesInvoice, error) {
	return action[SalesInvoice](s.client, ctx, fmt.Sprintf("/sales_invoices/%s/convert_to_invoice", invoiceID), map[string]interface{}{})
}

func (s *SalesInvoicesService) GetPDF(ctx context.Context, id string) ([]byte, error) {
	return getPDF(s.client, ctx, fmt.Sprintf("/sales_invoices/%s/pdf", id))
}

// PurchaseBillsService Alış faturaları servisi
type PurchaseBillsService struct {
	client *Client
}

func (s *PurchaseBillsService) List(ctx context.Context, params *ListParams) ([]PurchaseBill, *Meta, error) {
	return list[PurchaseBill](s.client, ctx, "/purchase_bills", params)
}

func (s *PurchaseBillsService) Get(ctx context.Context, id string) (*PurchaseBill, error) {
	return get[PurchaseBill](s.client, ctx, fmt.Sprintf("/purchase_bills/%s", id))
}

func (s *PurchaseBillsService) Create(ctx context.Context, attributes PurchaseBillAttributes, relationships *PurchaseBillRelationships) (*PurchaseBill, error) {
	return create[PurchaseBill](s.client, ctx, "/purchase_bills", "purchase_bills", attributes, relationships)
}

func (s *PurchaseBillsService) Update(ctx context.Context, id string, attributes PurchaseBillAttributes, relationships *PurchaseBillRelationships) (*PurchaseBill, error) {
	return update[PurchaseBill](s.client, ctx, fmt.Sprintf("/purchase_bills/%s", id), id, "purchase_bills", attributes, relationships)
}

func (s *PurchaseBillsService) CreatePayment(ctx context.Context, billID string, attributes PaymentAttributes) (*Payment, error) {
	return createPayment(s.client, ctx, fmt.Sprintf("/purchase_bills/%s", billID), attributes)
}

func (s *PurchaseBillsService) Cancel(ctx context.Context, billID string) error {
	return cancel(s.client, ctx, fmt.Sprintf("/purchase_bills/%s", billID))
}

func (s *PurchaseBillsService) Recover(ctx context.Context, billID string) error {
	return recover(s.client, ctx, fmt.Sprintf("/purchase_bills/%s", billID))
}

func (s *PurchaseBillsService) Archive(ctx context.Context, billID string) error {
	return archive(s.client, ctx, fmt.Sprintf("/purchase_bills/%s", billID))
}

func (s *PurchaseBillsService) Unarchive(ctx context.Context, billID string) error {
	return unarchive(s.client, ctx, fmt.Sprintf("/purchase_bills/%s", billID))
}

func (s *PurchaseBillsService) GetPDF(ctx context.Context, id string) ([]byte, error) {
	return getPDF(s.client, ctx, fmt.Sprintf("/purchase_bills/%s/pdf", id))
}

// EmployeesService Çalışanlar servisi
type EmployeesService struct {
	client *Client
}

func (s *EmployeesService) List(ctx context.Context, params *ListParams) ([]Employee, *Meta, error) {
	return list[Employee](s.client, ctx, "/employees", params)
}

func (s *EmployeesService) Get(ctx context.Context, id string) (*Employee, error) {
	return get[Employee](s.client, ctx, fmt.Sprintf("/employees/%s", id))
}

func (s *EmployeesService) Create(ctx context.Context, attributes EmployeeAttributes) (*Employee, error) {
	return create[Employee](s.client, ctx, "/employees", "employees", attributes, nil)
}

func (s *EmployeesService) Update(ctx context.Context, id string, attributes EmployeeAttributes) (*Employee, error) {
	return update[Employee](s.client, ctx, fmt.Sprintf("/employees/%s", id), id, "employees", attributes, nil)
}

func (s *EmployeesService) Archive(ctx context.Context, id string) error {
	return archive(s.client, ctx, fmt.Sprintf("/employees/%s", id))
}

func (s *EmployeesService) Unarchive(ctx context.Context, id string) error {
	return unarchive(s.client, ctx, fmt.Sprintf("/employees/%s", id))
}

// SalariesService Maaşlar servisi
type SalariesService struct {
	client *Client
}

func (s *SalariesService) List(ctx context.Context, params *ListParams) ([]Salary, *Meta, error) {
	return list[Salary](s.client, ctx, "/salaries", params)
}

func (s *SalariesService) Get(ctx context.Context, id string) (*Salary, error) {
	return get[Salary](s.client, ctx, fmt.Sprintf("/salaries/%s", id))
}

func (s *SalariesService) Create(ctx context.Context, attributes SalaryAttributes, relationships *SalaryRelationships) (*Salary, error) {
	return create[Salary](s.client, ctx, "/salaries", "salaries", attributes, relationships)
}

func (s *SalariesService) Update(ctx context.Context, id string, attributes SalaryAttributes, relationships *SalaryRelationships) (*Salary, error) {
	return update[Salary](s.client, ctx, fmt.Sprintf("/salaries/%s", id), id, "salaries", attributes, relationships)
}

func (s *SalariesService) Archive(ctx context.Context, id string) error {
	return archive(s.client, ctx, fmt.Sprintf("/salaries/%s", id))
}

func (s *SalariesService) Unarchive(ctx context.Context, id string) error {
	return unarchive(s.client, ctx, fmt.Sprintf("/salaries/%s", id))
}

func (s *SalariesService) CreatePayment(ctx context.Context, salaryID string, attributes PaymentAttributes) (*Payment, error) {
	return createPayment(s.client, ctx, fmt.Sprintf("/salaries/%s", salaryID), attributes)
}

// TaxesService Vergiler servisi
type TaxesService struct {
	client *Client
}

func (s *TaxesService) List(ctx context.Context, params *ListParams) ([]Tax, *Meta, error) {
	return list[Tax](s.client, ctx, "/taxes", params)
}

func (s *TaxesService) Get(ctx context.Context, id string) (*Tax, error) {
	return get[Tax](s.client, ctx, fmt.Sprintf("/taxes/%s", id))
}

func (s *TaxesService) Create(ctx context.Context, attributes TaxAttributes, relationships *TaxRelationships) (*Tax, error) {
	return create[Tax](s.client, ctx, "/taxes", "taxes", attributes, relationships)
}

func (s *TaxesService) Update(ctx context.Context, id string, attributes TaxAttributes, relationships *TaxRelationships) (*Tax, error) {
	return update[Tax](s.client, ctx, fmt.Sprintf("/taxes/%s", id), id, "taxes", attributes, relationships)
}

func (s *TaxesService) Archive(ctx context.Context, id string) error {
	return archive(s.client, ctx, fmt.Sprintf("/taxes/%s", id))
}

func (s *TaxesService) Unarchive(ctx context.Context, id string) error {
	return unarchive(s.client, ctx, fmt.Sprintf("/taxes/%s", id))
}

func (s *TaxesService) CreatePayment(ctx context.Context, taxID string, attributes PaymentAttributes) (*Payment, error) {
	return createPayment(s.client, ctx, fmt.Sprintf("/taxes/%s", taxID), attributes)
}

// TagsService Etiketler servisi
type TagsService struct {
	client *Client
}

func (s *TagsService) List(ctx context.Context, params *ListParams) ([]Tag, *Meta, error) {
	return list[Tag](s.client, ctx, "/tags", params)
}

func (s *TagsService) Get(ctx context.Context, id string) (*Tag, error) {
	return get[Tag](s.client, ctx, fmt.Sprintf("/tags/%s", id))
}

func (s *TagsService) Create(ctx context.Context, attributes TagAttributes) (*Tag, error) {
	return create[Tag](s.client, ctx, "/tags", "tags", attributes, nil)
}

func (s *TagsService) Update(ctx context.Context, id string, attributes TagAttributes) (*Tag, error) {
	return update[Tag](s.client, ctx, fmt.Sprintf("/tags/%s", id), id, "tags", attributes, nil)
}

func (s *TagsService) Delete(ctx context.Context, id string) error {
	return deleteResource(s.client, ctx, fmt.Sprintf("/tags/%s", id))
}

// WarehousesService Depolar servisi
type WarehousesService struct {
	client *Client
}

func (s *WarehousesService) List(ctx context.Context, params *ListParams) ([]Warehouse, *Meta, error) {
	return list[Warehouse](s.client, ctx, "/warehouses", params)
}

func (s *WarehousesService) Get(ctx context.Context, id string) (*Warehouse, error) {
	return get[Warehouse](s.client, ctx, fmt.Sprintf("/warehouses/%s", id))
}

func (s *WarehousesService) Create(ctx context.Context, attributes WarehouseAttributes) (*Warehouse, error) {
	return create[Warehouse](s.client, ctx, "/warehouses", "warehouses", attributes, nil)
}

func (s *WarehousesService) Update(ctx context.Context, id string, attributes WarehouseAttributes) (*Warehouse, error) {
	return update[Warehouse](s.client, ctx, fmt.Sprintf("/warehouses/%s", id), id, "warehouses", attributes, nil)
}

func (s *WarehousesService) Delete(ctx context.Context, id string) error {
	return deleteResource(s.client, ctx, fmt.Sprintf("/warehouses/%s", id))
}

// StockMovementsService Stok hareketleri servisi
type StockMovementsService struct {
	client *Client
}

func (s *StockMovementsService) List(ctx context.Context, params *ListParams) ([]StockMovement, *Meta, error) {
	return list[StockMovement](s.client, ctx, "/stock_movements", params)
}

// WebhooksService Webhooks servisi
type WebhooksService struct {
	client *Client
}

func (s *WebhooksService) List(ctx context.Context, params *ListParams) ([]Webhook, *Meta, error) {
	return list[Webhook](s.client, ctx, "/webhooks", params)
}

func (s *WebhooksService) Get(ctx context.Context, id string) (*Webhook, error) {
	return get[Webhook](s.client, ctx, fmt.Sprintf("/webhooks/%s", id))
}

func (s *WebhooksService) Create(ctx context.Context, attributes WebhookAttributes) (*Webhook, error) {
	return create[Webhook](s.client, ctx, "/webhooks", "webhooks", attributes, nil)
}

func (s *WebhooksService) Update(ctx context.Context, id string, attributes WebhookAttributes) (*Webhook, error) {
	return update[Webhook](s.client, ctx, fmt.Sprintf("/webhooks/%s", id), id, "webhooks", attributes, nil)
}

func (s *WebhooksService) Delete(ctx context.Context, id string) error {
	return deleteResource(s.client, ctx, fmt.Sprintf("/webhooks/%s", id))
}

// E-Document Services

// EArchivesService E-Arşiv servisi
type EArchivesService struct {
	client *Client
}

func (s *EArchivesService) List(ctx context.Context, params *ListParams) ([]EArchive, *Meta, error) {
	return list[EArchive](s.client, ctx, "/e_archives", params)
}

func (s *EArchivesService) Get(ctx context.Context, id string) (*EArchive, error) {
	return get[EArchive](s.client, ctx, fmt.Sprintf("/e_archives/%s", id))
}

func (s *EArchivesService) GetPDF(ctx context.Context, id string) ([]byte, error) {
	return getPDF(s.client, ctx, fmt.Sprintf("/e_archives/%s/pdf", id))
}

// EInvoiceInboxesService E-Fatura gelen kutusu servisi
type EInvoiceInboxesService struct {
	client *Client
}

func (s *EInvoiceInboxesService) List(ctx context.Context, params *ListParams) ([]EInvoiceInbox, *Meta, error) {
	return list[EInvoiceInbox](s.client, ctx, "/e_invoice_inboxes", params)
}

// EInvoicesService E-Fatura servisi
type EInvoicesService struct {
	client *Client
}

func (s *EInvoicesService) List(ctx context.Context, params *ListParams) ([]EInvoice, *Meta, error) {
	return list[EInvoice](s.client, ctx, "/e_invoices", params)
}

func (s *EInvoicesService) Get(ctx context.Context, id string) (*EInvoice, error) {
	return get[EInvoice](s.client, ctx, fmt.Sprintf("/e_invoices/%s", id))
}

func (s *EInvoicesService) Create(ctx context.Context, attributes EInvoiceAttributes) (*EInvoice, error) {
	return create[EInvoice](s.client, ctx, "/e_invoices", "e_invoices", attributes, nil)
}

func (s *EInvoicesService) GetPDF(ctx context.Context, id string) ([]byte, error) {
	return getPDF(s.client, ctx, fmt.Sprintf("/e_invoices/%s/pdf", id))
}

// ESMMsService E-SMM servisi
type ESMMsService struct {
	client *Client
}

func (s *ESMMsService) List(ctx context.Context, params *ListParams) ([]ESMM, *Meta, error) {
	return list[ESMM](s.client, ctx, "/e_smms", params)
}

func (s *ESMMsService) Get(ctx context.Context, id string) (*ESMM, error) {
	return get[ESMM](s.client, ctx, fmt.Sprintf("/e_smms/%s", id))
}

func (s *ESMMsService) Create(ctx context.Context, attributes ESMMAttributes) (*ESMM, error) {
	return create[ESMM](s.client, ctx, "/e_smms", "e_smms", attributes, nil)
}

func (s *ESMMsService) GetPDF(ctx context.Context, id string) ([]byte, error) {
	return getPDF(s.client, ctx, fmt.Sprintf("/e_smms/%s.pdf", id))
}

// ItemCategoriesService Ürün kategorileri servisi
type ItemCategoriesService struct {
	client *Client
}

func (s *ItemCategoriesService) List(ctx context.Context, params *ListParams) ([]ItemCategory, *Meta, error) {
	return list[ItemCategory](s.client, ctx, "/item_categories", params)
}

func (s *ItemCategoriesService) Get(ctx context.Context, id string) (*ItemCategory, error) {
	return get[ItemCategory](s.client, ctx, fmt.Sprintf("/item_categories/%s", id))
}

func (s *ItemCategoriesService) Create(ctx context.Context, attributes ItemCategoryAttributes) (*ItemCategory, error) {
	return create[ItemCategory](s.client, ctx, "/item_categories", "item_categories", attributes, nil)
}

func (s *ItemCategoriesService) Update(ctx context.Context, id string, attributes ItemCategoryAttributes) (*ItemCategory, error) {
	return update[ItemCategory](s.client, ctx, fmt.Sprintf("/item_categories/%s", id), id, "item_categories", attributes, nil)
}

func (s *ItemCategoriesService) Delete(ctx context.Context, id string) error {
	return deleteResource(s.client, ctx, fmt.Sprintf("/item_categories/%s", id))
}

// SalesOffersService Satış teklifleri servisi
type SalesOffersService struct {
	client *Client
}

func (s *SalesOffersService) List(ctx context.Context, params *ListParams) ([]SalesOffer, *Meta, error) {
	return list[SalesOffer](s.client, ctx, "/sales_offers", params)
}

func (s *SalesOffersService) Get(ctx context.Context, id string) (*SalesOffer, error) {
	return get[SalesOffer](s.client, ctx, fmt.Sprintf("/sales_offers/%s", id))
}

func (s *SalesOffersService) Create(ctx context.Context, attributes SalesOfferAttributes, relationships *SalesOfferRelationships) (*SalesOffer, error) {
	return create[SalesOffer](s.client, ctx, "/sales_offers", "sales_offers", attributes, relationships)
}

func (s *SalesOffersService) Update(ctx context.Context, id string, attributes SalesOfferAttributes, relationships *SalesOfferRelationships) (*SalesOffer, error) {
	return update[SalesOffer](s.client, ctx, fmt.Sprintf("/sales_offers/%s", id), id, "sales_offers", attributes, relationships)
}

func (s *SalesOffersService) Delete(ctx context.Context, id string) error {
	return deleteResource(s.client, ctx, fmt.Sprintf("/sales_offers/%s", id))
}

func (s *SalesOffersService) Archive(ctx context.Context, id string) error {
	return archive(s.client, ctx, fmt.Sprintf("/sales_offers/%s", id))
}

func (s *SalesOffersService) Unarchive(ctx context.Context, id string) error {
	return unarchive(s.client, ctx, fmt.Sprintf("/sales_offers/%s", id))
}

func (s *SalesOffersService) GetPDF(ctx context.Context, id string) ([]byte, error) {
	return getPDF(s.client, ctx, fmt.Sprintf("/sales_offers/%s/pdf", id))
}

func (s *SalesOffersService) GetDetails(ctx context.Context, id string) (*SalesOffer, error) {
	return get[SalesOffer](s.client, ctx, fmt.Sprintf("/sales_offers/%s/details", id))
}

func (s *SalesOffersService) UpdateStatus(ctx context.Context, id string, status string) (*SalesOffer, error) {
	body := map[string]interface{}{
		"data": map[string]interface{}{
			"id":   id,
			"type": "sales_offers",
			"attributes": map[string]interface{}{
				"status": status,
			},
		},
	}
	return action[SalesOffer](s.client, ctx, fmt.Sprintf("/sales_offers/%s/update_status", id), body)
}

// SharingsService Paylaşımlar servisi
type SharingsService struct {
	client *Client
}

func (s *SharingsService) List(ctx context.Context, params *ListParams) ([]Sharing, *Meta, error) {
	return list[Sharing](s.client, ctx, "/sharings", params)
}

func (s *SharingsService) Create(ctx context.Context, attributes SharingAttributes) (*Sharing, error) {
	return create[Sharing](s.client, ctx, "/sharings", "sharings", attributes, nil)
}

// ShipmentDocumentsService Sevkiyat belgeleri servisi
type ShipmentDocumentsService struct {
	client *Client
}

func (s *ShipmentDocumentsService) List(ctx context.Context, params *ListParams) ([]ShipmentDocument, *Meta, error) {
	return list[ShipmentDocument](s.client, ctx, "/shipment_documents", params)
}

func (s *ShipmentDocumentsService) Get(ctx context.Context, id string) (*ShipmentDocument, error) {
	return get[ShipmentDocument](s.client, ctx, fmt.Sprintf("/shipment_documents/%s", id))
}

func (s *ShipmentDocumentsService) Create(ctx context.Context, attributes ShipmentDocumentAttributes) (*ShipmentDocument, error) {
	return create[ShipmentDocument](s.client, ctx, "/shipment_documents", "shipment_documents", attributes, nil)
}

func (s *ShipmentDocumentsService) Update(ctx context.Context, id string, attributes ShipmentDocumentAttributes) (*ShipmentDocument, error) {
	return update[ShipmentDocument](s.client, ctx, fmt.Sprintf("/shipment_documents/%s", id), id, "shipment_documents", attributes, nil)
}

func (s *ShipmentDocumentsService) Delete(ctx context.Context, id string) error {
	return deleteResource(s.client, ctx, fmt.Sprintf("/shipment_documents/%s", id))
}

// StockUpdatesService Stok güncellemeleri servisi
type StockUpdatesService struct {
	client *Client
}

func (s *StockUpdatesService) Create(ctx context.Context, attributes StockUpdateAttributes) (*StockUpdate, error) {
	return create[StockUpdate](s.client, ctx, "/stock_updates", "stock_updates", attributes, nil)
}

// TrackableJobsService İzlenebilir işler servisi
type TrackableJobsService struct {
	client *Client
}

func (s *TrackableJobsService) Get(ctx context.Context, id string) (*TrackableJob, error) {
	return get[TrackableJob](s.client, ctx, fmt.Sprintf("/trackable_jobs/%s", id))
}

// TransactionsService İşlemler servisi
type TransactionsService struct {
	client *Client
}

func (s *TransactionsService) Get(ctx context.Context, id string) (*Transaction, error) {
	return get[Transaction](s.client, ctx, fmt.Sprintf("/transactions/%s", id))
}

func (s *TransactionsService) Update(ctx context.Context, id string, attributes TransactionAttributes) (*Transaction, error) {
	return update[Transaction](s.client, ctx, fmt.Sprintf("/transactions/%s", id), id, "transactions", attributes, nil)
}

func (s *TransactionsService) Delete(ctx context.Context, id string) error {
	return deleteResource(s.client, ctx, fmt.Sprintf("/transactions/%s", id))
}
