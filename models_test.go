package parasut

import (
	"encoding/json"
	"testing"
	"time"
)

func TestAccountJSON(t *testing.T) {
	// Test JSON unmarshaling
	jsonData := `{
		"id": "1",
		"type": "accounts",
		"attributes": {
			"name": "Test Account",
			"currency": "TRL",
			"account_type": "cash",
			"balance": 1000.50,
			"archived": false,
			"created_at": "2023-01-01T10:00:00Z",
			"updated_at": "2023-01-02T10:00:00Z"
		}
	}`

	var account Account
	err := json.Unmarshal([]byte(jsonData), &account)
	if err != nil {
		t.Fatalf("JSON unmarshal hatası: %v", err)
	}

	// Test attributes
	if account.ID != "1" {
		t.Errorf("ID = %s, beklenen 1", account.ID)
	}

	if account.Type != "accounts" {
		t.Errorf("Type = %s, beklenen accounts", account.Type)
	}

	if account.Attributes.Name != "Test Account" {
		t.Errorf("Name = %s, beklenen Test Account", account.Attributes.Name)
	}

	if account.Attributes.Currency != "TRL" {
		t.Errorf("Currency = %s, beklenen TRL", account.Attributes.Currency)
	}

	if account.Attributes.AccountType != "cash" {
		t.Errorf("AccountType = %s, beklenen cash", account.Attributes.AccountType)
	}

	if account.Attributes.Balance != 1000.50 {
		t.Errorf("Balance = %f, beklenen 1000.50", account.Attributes.Balance)
	}

	if account.Attributes.Archived {
		t.Error("Archived = true, beklenen false")
	}

	// Test JSON marshaling
	marshaledData, err := json.Marshal(account)
	if err != nil {
		t.Fatalf("JSON marshal hatası: %v", err)
	}

	// Tekrar unmarshal ederek test et
	var account2 Account
	err = json.Unmarshal(marshaledData, &account2)
	if err != nil {
		t.Fatalf("JSON re-unmarshal hatası: %v", err)
	}

	if account2.Attributes.Name != account.Attributes.Name {
		t.Error("Marshal/unmarshal sonrası name değişti")
	}
}

func TestContactJSON(t *testing.T) {
	jsonData := `{
		"id": "2",
		"type": "contacts",
		"attributes": {
			"name": "Test Müşteri",
			"email": "test@example.com",
			"contact_type": "company",
			"tax_number": "1234567890",
			"tax_office": "Test Tax Office",
			"city": "İstanbul",
			"address": "Test Adres",
			"phone": "+905551234567",
			"is_abroad": false,
			"archived": false,
			"account_type": "customer"
		}
	}`

	var contact Contact
	err := json.Unmarshal([]byte(jsonData), &contact)
	if err != nil {
		t.Fatalf("JSON unmarshal hatası: %v", err)
	}

	if contact.Attributes.Name != "Test Müşteri" {
		t.Errorf("Name = %s, beklenen Test Müşteri", contact.Attributes.Name)
	}

	if contact.Attributes.Email != "test@example.com" {
		t.Errorf("Email = %s, beklenen test@example.com", contact.Attributes.Email)
	}

	if contact.Attributes.ContactType != "company" {
		t.Errorf("ContactType = %s, beklenen company", contact.Attributes.ContactType)
	}

	if contact.Attributes.TaxNumber != "1234567890" {
		t.Errorf("TaxNumber = %s, beklenen 1234567890", contact.Attributes.TaxNumber)
	}

	if contact.Attributes.IsAbroad {
		t.Error("IsAbroad = true, beklenen false")
	}

	if contact.Attributes.AccountType != "customer" {
		t.Errorf("AccountType = %s, beklenen customer", contact.Attributes.AccountType)
	}
}

func TestProductJSON(t *testing.T) {
	jsonData := `{
		"id": "3",
		"type": "products",
		"attributes": {
			"code": "PROD001",
			"name": "Test Ürün",
			"vat_rate": 18.0,
			"sales_excise_duty_rate": 0.0,
			"purchase_excise_duty_rate": 0.0,
			"unit": "adet",
			"communications_tax_rate": 0.0,
			"archived": false,
			"list_price": 100.0,
			"currency": "TRL",
			"buying_price": 80.0,
			"buying_currency": "TRL",
			"inventory_tracking": true,
			"initial_stock_count": 50.0
		}
	}`

	var product Product
	err := json.Unmarshal([]byte(jsonData), &product)
	if err != nil {
		t.Fatalf("JSON unmarshal hatası: %v", err)
	}

	if product.Attributes.Code != "PROD001" {
		t.Errorf("Code = %s, beklenen PROD001", product.Attributes.Code)
	}

	if product.Attributes.Name != "Test Ürün" {
		t.Errorf("Name = %s, beklenen Test Ürün", product.Attributes.Name)
	}

	if product.Attributes.VatRate != 18.0 {
		t.Errorf("VatRate = %f, beklenen 18.0", product.Attributes.VatRate)
	}

	if product.Attributes.ListPrice != 100.0 {
		t.Errorf("ListPrice = %f, beklenen 100.0", product.Attributes.ListPrice)
	}

	if product.Attributes.Currency != "TRL" {
		t.Errorf("Currency = %s, beklenen TRL", product.Attributes.Currency)
	}

	if !product.Attributes.InventoryTracking {
		t.Error("InventoryTracking = false, beklenen true")
	}

	if product.Attributes.InitialStockCount != 50.0 {
		t.Errorf("InitialStockCount = %f, beklenen 50.0", product.Attributes.InitialStockCount)
	}
}

func TestSalesInvoiceJSON(t *testing.T) {
	jsonData := `{
		"id": "4",
		"type": "sales_invoices",
		"attributes": {
			"archived": false,
			"net_total": 100.0,
			"gross_total": 118.0,
			"withholding": 0.0,
			"total_excise_duty": 0.0,
			"total_communications_tax": 0.0,
			"total_vat": 18.0,
			"vat_withholding": 0.0,
			"total_discount": 0.0,
			"total_invoice_discount": 0.0,
			"before_taxes_total": 100.0,
			"remaining": 118.0,
			"remaining_in_trl": 118.0,
			"payment_status": "unpaid",
			"item_type": "invoice",
			"description": "Test Fatura",
			"issue_date": "2023-01-01",
			"due_date": "2023-01-31",
			"invoice_series": "A",
			"invoice_id": 1,
			"currency": "TRL",
			"exchange_rate": 1.0,
			"withholding_rate": 0.0,
			"vat_withholding_rate": 0.0,
			"invoice_discount_type": "percentage",
			"invoice_discount": 0.0,
			"is_abroad": false
		},
		"relationships": {
			"contact": {
				"id": "1",
				"type": "contacts"
			}
		}
	}`

	var invoice SalesInvoice
	err := json.Unmarshal([]byte(jsonData), &invoice)
	if err != nil {
		t.Fatalf("JSON unmarshal hatası: %v", err)
	}

	if invoice.Attributes.NetTotal != 100.0 {
		t.Errorf("NetTotal = %f, beklenen 100.0", invoice.Attributes.NetTotal)
	}

	if invoice.Attributes.GrossTotal != 118.0 {
		t.Errorf("GrossTotal = %f, beklenen 118.0", invoice.Attributes.GrossTotal)
	}

	if invoice.Attributes.TotalVat != 18.0 {
		t.Errorf("TotalVat = %f, beklenen 18.0", invoice.Attributes.TotalVat)
	}

	if invoice.Attributes.Description != "Test Fatura" {
		t.Errorf("Description = %s, beklenen Test Fatura", invoice.Attributes.Description)
	}

	if invoice.Attributes.IssueDate != "2023-01-01" {
		t.Errorf("IssueDate = %s, beklenen 2023-01-01", invoice.Attributes.IssueDate)
	}

	if invoice.Attributes.ItemType != "invoice" {
		t.Errorf("ItemType = %s, beklenen invoice", invoice.Attributes.ItemType)
	}

	if invoice.Attributes.PaymentStatus != "unpaid" {
		t.Errorf("PaymentStatus = %s, beklenen unpaid", invoice.Attributes.PaymentStatus)
	}

	// Test relationships
	if invoice.Relationships.Contact == nil {
		t.Fatal("Contact relationship nil")
	}

	if invoice.Relationships.Contact.ID != "1" {
		t.Errorf("Contact ID = %s, beklenen 1", invoice.Relationships.Contact.ID)
	}

	if invoice.Relationships.Contact.Type != "contacts" {
		t.Errorf("Contact Type = %s, beklenen contacts", invoice.Relationships.Contact.Type)
	}
}

func TestPurchaseBillJSON(t *testing.T) {
	jsonData := `{
		"id": "5",
		"type": "purchase_bills",
		"attributes": {
			"archived": false,
			"net_total": 500.0,
			"gross_total": 590.0,
			"total_vat": 90.0,
			"item_type": "bill",
			"description": "Test Gider Faturası",
			"issue_date": "2023-01-15",
			"due_date": "2023-02-15",
			"currency": "TRL",
			"supplier_name": "Test Tedarikçi",
			"supplier_tax_number": "9876543210",
			"supplier_tax_office": "Test Vergi Dairesi"
		}
	}`

	var bill PurchaseBill
	err := json.Unmarshal([]byte(jsonData), &bill)
	if err != nil {
		t.Fatalf("JSON unmarshal hatası: %v", err)
	}

	if bill.Attributes.NetTotal != 500.0 {
		t.Errorf("NetTotal = %f, beklenen 500.0", bill.Attributes.NetTotal)
	}

	if bill.Attributes.GrossTotal != 590.0 {
		t.Errorf("GrossTotal = %f, beklenen 590.0", bill.Attributes.GrossTotal)
	}

	if bill.Attributes.Description != "Test Gider Faturası" {
		t.Errorf("Description = %s, beklenen Test Gider Faturası", bill.Attributes.Description)
	}

	if bill.Attributes.SupplierName != "Test Tedarikçi" {
		t.Errorf("SupplierName = %s, beklenen Test Tedarikçi", bill.Attributes.SupplierName)
	}

	if bill.Attributes.SupplierTaxNumber != "9876543210" {
		t.Errorf("SupplierTaxNumber = %s, beklenen 9876543210", bill.Attributes.SupplierTaxNumber)
	}
}

func TestEmployeeJSON(t *testing.T) {
	now := time.Now()
	employee := Employee{
		ID:   "1",
		Type: "employees",
		Attributes: EmployeeAttributes{
			Name:      "Test Çalışan",
			Email:     "calisan@example.com",
			IBAN:      "TR123456789012345678901234",
			Archived:  false,
			CreatedAt: &now,
			UpdatedAt: &now,
		},
	}

	// Test JSON marshaling
	data, err := json.Marshal(employee)
	if err != nil {
		t.Fatalf("JSON marshal hatası: %v", err)
	}

	// Test JSON unmarshaling
	var employee2 Employee
	err = json.Unmarshal(data, &employee2)
	if err != nil {
		t.Fatalf("JSON unmarshal hatası: %v", err)
	}

	if employee2.Attributes.Name != "Test Çalışan" {
		t.Errorf("Name = %s, beklenen Test Çalışan", employee2.Attributes.Name)
	}

	if employee2.Attributes.Email != "calisan@example.com" {
		t.Errorf("Email = %s, beklenen calisan@example.com", employee2.Attributes.Email)
	}

	if employee2.Attributes.IBAN != "TR123456789012345678901234" {
		t.Errorf("IBAN = %s, beklenen TR123456789012345678901234", employee2.Attributes.IBAN)
	}
}

func TestWebhookJSON(t *testing.T) {
	webhook := Webhook{
		ID:   "1",
		Type: "webhooks",
		Attributes: WebhookAttributes{
			URL:           "https://example.com/webhook",
			Event:         "sales_invoice.created",
			IsActive:      true,
			EncryptionKey: "secret123",
		},
	}

	// Test JSON marshaling
	data, err := json.Marshal(webhook)
	if err != nil {
		t.Fatalf("JSON marshal hatası: %v", err)
	}

	// Test JSON unmarshaling
	var webhook2 Webhook
	err = json.Unmarshal(data, &webhook2)
	if err != nil {
		t.Fatalf("JSON unmarshal hatası: %v", err)
	}

	if webhook2.Attributes.URL != "https://example.com/webhook" {
		t.Errorf("URL = %s, beklenen https://example.com/webhook", webhook2.Attributes.URL)
	}

	if webhook2.Attributes.Event != "sales_invoice.created" {
		t.Errorf("Event = %s, beklenen sales_invoice.created", webhook2.Attributes.Event)
	}

	if !webhook2.Attributes.IsActive {
		t.Error("IsActive = false, beklenen true")
	}
}

func TestPaymentJSON(t *testing.T) {
	payment := Payment{
		ID:   "1",
		Type: "payments",
		Attributes: PaymentAttributes{
			Date:        "2023-01-01",
			Amount:      250.75,
			Description: "Test Ödeme",
			Currency:    "TRL",
		},
	}

	// Test JSON marshaling
	data, err := json.Marshal(payment)
	if err != nil {
		t.Fatalf("JSON marshal hatası: %v", err)
	}

	// Test JSON unmarshaling
	var payment2 Payment
	err = json.Unmarshal(data, &payment2)
	if err != nil {
		t.Fatalf("JSON unmarshal hatası: %v", err)
	}

	if payment2.Attributes.Amount != 250.75 {
		t.Errorf("Amount = %f, beklenen 250.75", payment2.Attributes.Amount)
	}

	if payment2.Attributes.Description != "Test Ödeme" {
		t.Errorf("Description = %s, beklenen Test Ödeme", payment2.Attributes.Description)
	}

	if payment2.Attributes.Currency != "TRL" {
		t.Errorf("Currency = %s, beklenen TRL", payment2.Attributes.Currency)
	}
}

func TestMeJSON(t *testing.T) {
	jsonData := `{
		"id": "1",
		"type": "me",
		"attributes": {
			"name": "Test User",
			"email": "user@example.com",
			"is_confirmed": true
		},
		"relationships": {
			"profile": {
				"id": "1",
				"type": "profiles"
			}
		}
	}`

	var me Me
	err := json.Unmarshal([]byte(jsonData), &me)
	if err != nil {
		t.Fatalf("JSON unmarshal hatası: %v", err)
	}

	if me.Attributes.Name != "Test User" {
		t.Errorf("Name = %s, beklenen Test User", me.Attributes.Name)
	}

	if me.Attributes.Email != "user@example.com" {
		t.Errorf("Email = %s, beklenen user@example.com", me.Attributes.Email)
	}

	if !me.Attributes.IsConfirmed {
		t.Error("IsConfirmed = false, beklenen true")
	}

	if me.Relationships.Profile == nil {
		t.Fatal("Profile relationship nil")
	}

	if me.Relationships.Profile.ID != "1" {
		t.Errorf("Profile ID = %s, beklenen 1", me.Relationships.Profile.ID)
	}
}

func TestStockMovementJSON(t *testing.T) {
	stockMovement := StockMovement{
		ID:   "1",
		Type: "stock_movements",
		Attributes: StockMovementAttributes{
			Date:         "2023-01-01",
			MovementType: "in",
			Quantity:     25.0,
			Description:  "Stok Giriş",
		},
	}

	// Test JSON marshaling
	data, err := json.Marshal(stockMovement)
	if err != nil {
		t.Fatalf("JSON marshal hatası: %v", err)
	}

	// Test JSON unmarshaling
	var stockMovement2 StockMovement
	err = json.Unmarshal(data, &stockMovement2)
	if err != nil {
		t.Fatalf("JSON unmarshal hatası: %v", err)
	}

	if stockMovement2.Attributes.MovementType != "in" {
		t.Errorf("MovementType = %s, beklenen in", stockMovement2.Attributes.MovementType)
	}

	if stockMovement2.Attributes.Quantity != 25.0 {
		t.Errorf("Quantity = %f, beklenen 25.0", stockMovement2.Attributes.Quantity)
	}

	if stockMovement2.Attributes.Description != "Stok Giriş" {
		t.Errorf("Description = %s, beklenen Stok Giriş", stockMovement2.Attributes.Description)
	}
}

func TestWarehouseJSON(t *testing.T) {
	warehouse := Warehouse{
		ID:   "1",
		Type: "warehouses",
		Attributes: WarehouseAttributes{
			Name:     "Ana Depo",
			City:     "İstanbul",
			District: "Kadıköy",
			Address:  "Test Adres 123",
		},
	}

	// Test JSON marshaling
	data, err := json.Marshal(warehouse)
	if err != nil {
		t.Fatalf("JSON marshal hatası: %v", err)
	}

	// Test JSON unmarshaling
	var warehouse2 Warehouse
	err = json.Unmarshal(data, &warehouse2)
	if err != nil {
		t.Fatalf("JSON unmarshal hatası: %v", err)
	}

	if warehouse2.Attributes.Name != "Ana Depo" {
		t.Errorf("Name = %s, beklenen Ana Depo", warehouse2.Attributes.Name)
	}

	if warehouse2.Attributes.City != "İstanbul" {
		t.Errorf("City = %s, beklenen İstanbul", warehouse2.Attributes.City)
	}

	if warehouse2.Attributes.District != "Kadıköy" {
		t.Errorf("District = %s, beklenen Kadıköy", warehouse2.Attributes.District)
	}
}

func TestTagJSON(t *testing.T) {
	tag := Tag{
		ID:   "1",
		Type: "tags",
		Attributes: TagAttributes{
			Name:  "Önemli",
			Color: "#FF0000",
		},
	}

	// Test JSON marshaling
	data, err := json.Marshal(tag)
	if err != nil {
		t.Fatalf("JSON marshal hatası: %v", err)
	}

	// Test JSON unmarshaling
	var tag2 Tag
	err = json.Unmarshal(data, &tag2)
	if err != nil {
		t.Fatalf("JSON unmarshal hatası: %v", err)
	}

	if tag2.Attributes.Name != "Önemli" {
		t.Errorf("Name = %s, beklenen Önemli", tag2.Attributes.Name)
	}

	if tag2.Attributes.Color != "#FF0000" {
		t.Errorf("Color = %s, beklenen #FF0000", tag2.Attributes.Color)
	}
}

func TestRelationshipDataJSON(t *testing.T) {
	relationships := SalesInvoiceRelationships{
		Contact: &RelationshipData{
			ID:   "123",
			Type: "contacts",
		},
		Details: []RelationshipData{
			{ID: "1", Type: "sales_invoice_details"},
			{ID: "2", Type: "sales_invoice_details"},
		},
	}

	// Test JSON marshaling
	data, err := json.Marshal(relationships)
	if err != nil {
		t.Fatalf("JSON marshal hatası: %v", err)
	}

	// Test JSON unmarshaling
	var relationships2 SalesInvoiceRelationships
	err = json.Unmarshal(data, &relationships2)
	if err != nil {
		t.Fatalf("JSON unmarshal hatası: %v", err)
	}

	if relationships2.Contact.ID != "123" {
		t.Errorf("Contact ID = %s, beklenen 123", relationships2.Contact.ID)
	}

	if len(relationships2.Details) != 2 {
		t.Errorf("Details sayısı = %d, beklenen 2", len(relationships2.Details))
	}

	if relationships2.Details[0].Type != "sales_invoice_details" {
		t.Errorf("Details[0] Type = %s, beklenen sales_invoice_details", relationships2.Details[0].Type)
	}
}
