package parasut

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Mock client oluşturucu helper
func createTestClient(handler http.HandlerFunc) *Client {
	server := httptest.NewServer(handler)
	config := &Config{
		ClientID:     "test-client-id",
		ClientSecret: "test-client-secret",
		RedirectURL:  "http://localhost:8080/callback",
		CompanyID:    123,
	}

	client := NewClient(config)
	client.baseURL = server.URL + "/v4"
	return client
}

// Generic helper metodları testleri

func TestGenericList(t *testing.T) {
	expectedAccounts := []Account{
		{
			ID:   "1",
			Type: "accounts",
			Attributes: AccountAttributes{
				Name:        "Test Account",
				Currency:    "TRL",
				AccountType: "cash",
				Balance:     1000.0,
			},
		},
	}

	client := createTestClient(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Method = %s, beklenen GET", r.Method)
		}

		if !strings.Contains(r.URL.Path, "/accounts") {
			t.Errorf("Path = %s, accounts içermeli", r.URL.Path)
		}

		// Query parametrelerini kontrol et
		if r.URL.Query().Get("page") != "1" {
			t.Errorf("page parametresi = %s, beklenen 1", r.URL.Query().Get("page"))
		}

		response := struct {
			Data []Account `json:"data"`
			Meta *Meta     `json:"meta"`
		}{
			Data: expectedAccounts,
			Meta: &Meta{
				CurrentPage: 1,
				TotalPages:  1,
				TotalCount:  1,
			},
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	ctx := context.Background()
	params := &ListParams{
		Page:     1,
		PageSize: 10,
	}

	accounts, meta, err := list[Account](client, ctx, "/accounts", params)

	if err != nil {
		t.Fatalf("list fonksiyonu hata döndü: %v", err)
	}

	if len(accounts) != 1 {
		t.Errorf("Account sayısı = %d, beklenen 1", len(accounts))
	}

	if accounts[0].Attributes.Name != "Test Account" {
		t.Errorf("Account name = %s, beklenen Test Account", accounts[0].Attributes.Name)
	}

	if meta.TotalCount != 1 {
		t.Errorf("Meta TotalCount = %d, beklenen 1", meta.TotalCount)
	}
}

func TestGenericGet(t *testing.T) {
	expectedAccount := Account{
		ID:   "1",
		Type: "accounts",
		Attributes: AccountAttributes{
			Name:        "Test Account",
			Currency:    "TRL",
			AccountType: "cash",
			Balance:     1000.0,
		},
	}

	client := createTestClient(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Method = %s, beklenen GET", r.Method)
		}

		if !strings.Contains(r.URL.Path, "/accounts/1") {
			t.Errorf("Path = %s, /accounts/1 içermeli", r.URL.Path)
		}

		response := struct {
			Data Account `json:"data"`
		}{
			Data: expectedAccount,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	ctx := context.Background()
	account, err := get[Account](client, ctx, "/accounts/1")

	if err != nil {
		t.Fatalf("get fonksiyonu hata döndü: %v", err)
	}

	if account.Attributes.Name != "Test Account" {
		t.Errorf("Account name = %s, beklenen Test Account", account.Attributes.Name)
	}
}

func TestGenericCreate(t *testing.T) {
	expectedAccount := Account{
		ID:   "1",
		Type: "accounts",
		Attributes: AccountAttributes{
			Name:        "New Account",
			Currency:    "TRL",
			AccountType: "cash",
		},
	}

	client := createTestClient(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Method = %s, beklenen POST", r.Method)
		}

		if !strings.Contains(r.URL.Path, "/accounts") {
			t.Errorf("Path = %s, /accounts içermeli", r.URL.Path)
		}

		// Request body'sini kontrol et
		var requestBody map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			t.Fatalf("Request body decode edilemedi: %v", err)
		}

		data := requestBody["data"].(map[string]interface{})
		if data["type"] != "accounts" {
			t.Errorf("Request type = %s, beklenen accounts", data["type"])
		}

		attributes := data["attributes"].(map[string]interface{})
		if attributes["name"] != "New Account" {
			t.Errorf("Request name = %s, beklenen New Account", attributes["name"])
		}

		response := struct {
			Data Account `json:"data"`
		}{
			Data: expectedAccount,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	})

	ctx := context.Background()
	attributes := AccountAttributes{
		Name:        "New Account",
		Currency:    "TRL",
		AccountType: "cash",
	}

	account, err := create[Account](client, ctx, "/accounts", "accounts", attributes, nil)

	if err != nil {
		t.Fatalf("create fonksiyonu hata döndü: %v", err)
	}

	if account.Attributes.Name != "New Account" {
		t.Errorf("Account name = %s, beklenen New Account", account.Attributes.Name)
	}
}

func TestGenericUpdate(t *testing.T) {
	expectedAccount := Account{
		ID:   "1",
		Type: "accounts",
		Attributes: AccountAttributes{
			Name:        "Updated Account",
			Currency:    "TRL",
			AccountType: "cash",
		},
	}

	client := createTestClient(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" {
			t.Errorf("Method = %s, beklenen PUT", r.Method)
		}

		if !strings.Contains(r.URL.Path, "/accounts") {
			t.Errorf("Path = %s, /accounts içermeli", r.URL.Path)
		}

		// Request body'sini kontrol et
		var requestBody map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			t.Fatalf("Request body decode edilemedi: %v", err)
		}

		data := requestBody["data"].(map[string]interface{})
		if data["id"] != "1" {
			t.Errorf("Request id = %s, beklenen 1", data["id"])
		}

		response := struct {
			Data Account `json:"data"`
		}{
			Data: expectedAccount,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	ctx := context.Background()
	attributes := AccountAttributes{
		Name:        "Updated Account",
		Currency:    "TRL",
		AccountType: "cash",
	}

	account, err := update[Account](client, ctx, "/accounts", "1", "accounts", attributes, nil)

	if err != nil {
		t.Fatalf("update fonksiyonu hata döndü: %v", err)
	}

	if account.Attributes.Name != "Updated Account" {
		t.Errorf("Account name = %s, beklenen Updated Account", account.Attributes.Name)
	}
}

func TestGenericDelete(t *testing.T) {
	client := createTestClient(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			t.Errorf("Method = %s, beklenen DELETE", r.Method)
		}

		if !strings.Contains(r.URL.Path, "/accounts/1") {
			t.Errorf("Path = %s, /accounts/1 içermeli", r.URL.Path)
		}

		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	err := deleteResource(client, ctx, "/accounts/1")

	if err != nil {
		t.Fatalf("deleteResource fonksiyonu hata döndü: %v", err)
	}
}

func TestGenericArchive(t *testing.T) {
	client := createTestClient(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PATCH" {
			t.Errorf("Method = %s, beklenen PATCH", r.Method)
		}

		if !strings.Contains(r.URL.Path, "/accounts/1/archive") {
			t.Errorf("Path = %s, /accounts/1/archive içermeli", r.URL.Path)
		}

		w.WriteHeader(http.StatusOK)
	})

	ctx := context.Background()
	err := archive(client, ctx, "/accounts/1")

	if err != nil {
		t.Fatalf("archive fonksiyonu hata döndü: %v", err)
	}
}

// Servis testleri

func TestAccountsService_List(t *testing.T) {
	client := createTestClient(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.URL.Path, "/accounts") {
			t.Errorf("Path = %s, /accounts içermeli", r.URL.Path)
		}

		response := struct {
			Data []Account `json:"data"`
			Meta *Meta     `json:"meta"`
		}{
			Data: []Account{
				{
					ID:   "1",
					Type: "accounts",
					Attributes: AccountAttributes{
						Name:     "Test Account",
						Currency: "TRL",
					},
				},
			},
			Meta: &Meta{TotalCount: 1},
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	ctx := context.Background()
	accounts, meta, err := client.Accounts.List(ctx, nil)

	if err != nil {
		t.Fatalf("Accounts.List hata döndü: %v", err)
	}

	if len(accounts) != 1 {
		t.Errorf("Account sayısı = %d, beklenen 1", len(accounts))
	}

	if meta.TotalCount != 1 {
		t.Errorf("TotalCount = %d, beklenen 1", meta.TotalCount)
	}
}

func TestAccountsService_Get(t *testing.T) {
	client := createTestClient(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.URL.Path, "/accounts/1") {
			t.Errorf("Path = %s, /accounts/1 içermeli", r.URL.Path)
		}

		response := struct {
			Data Account `json:"data"`
		}{
			Data: Account{
				ID:   "1",
				Type: "accounts",
				Attributes: AccountAttributes{
					Name:     "Test Account",
					Currency: "TRL",
				},
			},
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	ctx := context.Background()
	account, err := client.Accounts.Get(ctx, "1")

	if err != nil {
		t.Fatalf("Accounts.Get hata döndü: %v", err)
	}

	if account.ID != "1" {
		t.Errorf("Account ID = %s, beklenen 1", account.ID)
	}
}

func TestAccountsService_Create(t *testing.T) {
	client := createTestClient(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Method = %s, beklenen POST", r.Method)
		}

		if !strings.Contains(r.URL.Path, "/accounts") {
			t.Errorf("Path = %s, /accounts içermeli", r.URL.Path)
		}

		response := struct {
			Data Account `json:"data"`
		}{
			Data: Account{
				ID:   "1",
				Type: "accounts",
				Attributes: AccountAttributes{
					Name:     "New Account",
					Currency: "TRL",
				},
			},
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	})

	ctx := context.Background()
	attributes := AccountAttributes{
		Name:        "New Account",
		Currency:    "TRL",
		AccountType: "cash",
	}

	account, err := client.Accounts.Create(ctx, attributes)

	if err != nil {
		t.Fatalf("Accounts.Create hata döndü: %v", err)
	}

	if account.Attributes.Name != "New Account" {
		t.Errorf("Account name = %s, beklenen New Account", account.Attributes.Name)
	}
}

func TestContactsService_List(t *testing.T) {
	client := createTestClient(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.URL.Path, "/contacts") {
			t.Errorf("Path = %s, /contacts içermeli", r.URL.Path)
		}

		response := struct {
			Data []Contact `json:"data"`
			Meta *Meta     `json:"meta"`
		}{
			Data: []Contact{
				{
					ID:   "1",
					Type: "contacts",
					Attributes: ContactAttributes{
						Name:        "Test Contact",
						ContactType: "person",
					},
				},
			},
			Meta: &Meta{TotalCount: 1},
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	ctx := context.Background()
	contacts, _, err := client.Contacts.List(ctx, nil)

	if err != nil {
		t.Fatalf("Contacts.List hata döndü: %v", err)
	}

	if len(contacts) != 1 {
		t.Errorf("Contact sayısı = %d, beklenen 1", len(contacts))
	}

	if contacts[0].Attributes.Name != "Test Contact" {
		t.Errorf("Contact name = %s, beklenen Test Contact", contacts[0].Attributes.Name)
	}
}

func TestProductsService_List(t *testing.T) {
	client := createTestClient(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.URL.Path, "/products") {
			t.Errorf("Path = %s, /products içermeli", r.URL.Path)
		}

		response := struct {
			Data []Product `json:"data"`
			Meta *Meta     `json:"meta"`
		}{
			Data: []Product{
				{
					ID:   "1",
					Type: "products",
					Attributes: ProductAttributes{
						Code:      "PROD001",
						Name:      "Test Product",
						ListPrice: 100.0,
						Currency:  "TRL",
					},
				},
			},
			Meta: &Meta{TotalCount: 1},
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	ctx := context.Background()
	products, _, err := client.Products.List(ctx, nil)

	if err != nil {
		t.Fatalf("Products.List hata döndü: %v", err)
	}

	if len(products) != 1 {
		t.Errorf("Product sayısı = %d, beklenen 1", len(products))
	}

	if products[0].Attributes.Code != "PROD001" {
		t.Errorf("Product code = %s, beklenen PROD001", products[0].Attributes.Code)
	}
}

func TestSalesInvoicesService_Create(t *testing.T) {
	client := createTestClient(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Method = %s, beklenen POST", r.Method)
		}

		if !strings.Contains(r.URL.Path, "/sales_invoices") {
			t.Errorf("Path = %s, /sales_invoices içermeli", r.URL.Path)
		}

		// Request body'sini kontrol et
		var requestBody map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			t.Fatalf("Request body decode edilemedi: %v", err)
		}

		data := requestBody["data"].(map[string]interface{})
		if data["type"] != "sales_invoices" {
			t.Errorf("Request type = %s, beklenen sales_invoices", data["type"])
		}

		// Relationships kontrol et
		if relationships, exists := data["relationships"]; exists {
			rel := relationships.(map[string]interface{})
			if contact, exists := rel["contact"]; exists {
				contactData := contact.(map[string]interface{})
				if contactData["id"] != "1" {
					t.Errorf("Contact ID = %s, beklenen 1", contactData["id"])
				}
			}
		}

		response := struct {
			Data SalesInvoice `json:"data"`
		}{
			Data: SalesInvoice{
				ID:   "1",
				Type: "sales_invoices",
				Attributes: SalesInvoiceAttributes{
					Description: "Test Invoice",
					IssueDate:   "2023-01-01",
					NetTotal:    100.0,
				},
			},
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	})

	ctx := context.Background()
	attributes := SalesInvoiceAttributes{
		Description: "Test Invoice",
		IssueDate:   "2023-01-01",
	}

	relationships := &SalesInvoiceRelationships{
		Contact: &RelationshipData{
			ID:   "1",
			Type: "contacts",
		},
	}

	invoice, err := client.SalesInvoices.Create(ctx, attributes, relationships)

	if err != nil {
		t.Fatalf("SalesInvoices.Create hata döndü: %v", err)
	}

	if invoice.Attributes.Description != "Test Invoice" {
		t.Errorf("Invoice description = %s, beklenen Test Invoice", invoice.Attributes.Description)
	}
}

func TestSalesInvoicesService_CreatePayment(t *testing.T) {
	client := createTestClient(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Method = %s, beklenen POST", r.Method)
		}

		if !strings.Contains(r.URL.Path, "/sales_invoices/1/payments") {
			t.Errorf("Path = %s, /sales_invoices/1/payments içermeli", r.URL.Path)
		}

		response := struct {
			Data Payment `json:"data"`
		}{
			Data: Payment{
				ID:   "1",
				Type: "payments",
				Attributes: PaymentAttributes{
					Date:   "2023-01-01",
					Amount: 100.0,
				},
			},
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	})

	ctx := context.Background()
	attributes := PaymentAttributes{
		Date:   "2023-01-01",
		Amount: 100.0,
	}

	payment, err := client.SalesInvoices.CreatePayment(ctx, "1", attributes)

	if err != nil {
		t.Fatalf("SalesInvoices.CreatePayment hata döndü: %v", err)
	}

	if payment.Attributes.Amount != 100.0 {
		t.Errorf("Payment amount = %f, beklenen 100.0", payment.Attributes.Amount)
	}
}

func TestMeService_Get(t *testing.T) {
	client := createTestClient(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.URL.Path, "/me") {
			t.Errorf("Path = %s, /me içermeli", r.URL.Path)
		}

		response := struct {
			Data Me `json:"data"`
		}{
			Data: Me{
				ID:   "1",
				Type: "me",
				Attributes: MeAttributes{
					Name:        "Test User",
					Email:       "test@example.com",
					IsConfirmed: true,
				},
			},
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	ctx := context.Background()
	me, err := client.Me.Get(ctx)

	if err != nil {
		t.Fatalf("Me.Get hata döndü: %v", err)
	}

	if me.Attributes.Email != "test@example.com" {
		t.Errorf("Me email = %s, beklenen test@example.com", me.Attributes.Email)
	}
}

func TestWebhooksService_CRUD(t *testing.T) {
	// Create test
	t.Run("Create", func(t *testing.T) {
		client := createTestClient(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "POST" {
				t.Errorf("Method = %s, beklenen POST", r.Method)
			}

			response := struct {
				Data Webhook `json:"data"`
			}{
				Data: Webhook{
					ID:   "1",
					Type: "webhooks",
					Attributes: WebhookAttributes{
						URL:      "https://example.com/webhook",
						Event:    "sales_invoice.created",
						IsActive: true,
					},
				},
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(response)
		})

		ctx := context.Background()
		attributes := WebhookAttributes{
			URL:      "https://example.com/webhook",
			Event:    "sales_invoice.created",
			IsActive: true,
		}

		webhook, err := client.Webhooks.Create(ctx, attributes)

		if err != nil {
			t.Fatalf("Webhooks.Create hata döndü: %v", err)
		}

		if webhook.Attributes.URL != "https://example.com/webhook" {
			t.Errorf("Webhook URL = %s, beklenen https://example.com/webhook", webhook.Attributes.URL)
		}
	})

	// Update test
	t.Run("Update", func(t *testing.T) {
		client := createTestClient(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "PUT" {
				t.Errorf("Method = %s, beklenen PUT", r.Method)
			}

			response := struct {
				Data Webhook `json:"data"`
			}{
				Data: Webhook{
					ID:   "1",
					Type: "webhooks",
					Attributes: WebhookAttributes{
						URL:      "https://updated.com/webhook",
						Event:    "sales_invoice.updated",
						IsActive: false,
					},
				},
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		})

		ctx := context.Background()
		attributes := WebhookAttributes{
			URL:      "https://updated.com/webhook",
			Event:    "sales_invoice.updated",
			IsActive: false,
		}

		webhook, err := client.Webhooks.Update(ctx, "1", attributes)

		if err != nil {
			t.Fatalf("Webhooks.Update hata döndü: %v", err)
		}

		if webhook.Attributes.IsActive {
			t.Error("Webhook IsActive = true, beklenen false")
		}
	})
}

func TestCreatePayment(t *testing.T) {
	client := createTestClient(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Method = %s, beklenen POST", r.Method)
		}

		if !strings.Contains(r.URL.Path, "/payments") {
			t.Errorf("Path = %s, /payments içermeli", r.URL.Path)
		}

		response := struct {
			Data Payment `json:"data"`
		}{
			Data: Payment{
				ID:   "1",
				Type: "payments",
				Attributes: PaymentAttributes{
					Date:     "2023-01-01",
					Amount:   500.0,
					Currency: "TRL",
				},
			},
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	})

	ctx := context.Background()
	attributes := PaymentAttributes{
		Date:     "2023-01-01",
		Amount:   500.0,
		Currency: "TRL",
	}

	payment, err := createPayment(client, ctx, "/test-endpoint", attributes)

	if err != nil {
		t.Fatalf("createPayment hata döndü: %v", err)
	}

	if payment.Attributes.Amount != 500.0 {
		t.Errorf("Payment amount = %f, beklenen 500.0", payment.Attributes.Amount)
	}
}
