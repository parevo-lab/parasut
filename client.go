package parasut

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"golang.org/x/oauth2"
)

const (
	// BaseURL Parasüt API'nin temel URL'i
	BaseURL = "https://api.parasut.com/v4"

	// OAuth2 URL'leri
	AuthURL  = "https://api.parasut.com/oauth/authorize"
	TokenURL = "https://api.parasut.com/oauth/token"
)

// Client Parasüt API istemcisi
type Client struct {
	httpClient *http.Client
	baseURL    string
	companyID  int
	config     *oauth2.Config
	token      *oauth2.Token

	// Services
	Me                *MeService
	Accounts          *AccountsService
	BankFees          *BankFeesService
	Contacts          *ContactsService
	EArchives         *EArchivesService
	EInvoiceInboxes   *EInvoiceInboxesService
	EInvoices         *EInvoicesService
	ESMMs             *ESMMsService
	Employees         *EmployeesService
	ItemCategories    *ItemCategoriesService
	Products          *ProductsService
	PurchaseBills     *PurchaseBillsService
	Salaries          *SalariesService
	SalesInvoices     *SalesInvoicesService
	SalesOffers       *SalesOffersService
	Sharings          *SharingsService
	ShipmentDocuments *ShipmentDocumentsService
	StockMovements    *StockMovementsService
	StockUpdates      *StockUpdatesService
	Tags              *TagsService
	Taxes             *TaxesService
	TrackableJobs     *TrackableJobsService
	Transactions      *TransactionsService
	Warehouses        *WarehousesService
	Webhooks          *WebhooksService
}

// Config Parasüt API ayarları
type Config struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
	CompanyID    int
}

// NewClient yeni bir Parasüt istemcisi oluşturur
func NewClient(config *Config) *Client {
	oauth2Config := &oauth2.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		RedirectURL:  config.RedirectURL,
		Endpoint: oauth2.Endpoint{
			AuthURL:  AuthURL,
			TokenURL: TokenURL,
		},
	}

	client := &Client{
		httpClient: &http.Client{Timeout: 30 * time.Second},
		baseURL:    BaseURL,
		companyID:  config.CompanyID,
		config:     oauth2Config,
	}

	// Initialize services
	client.Me = &MeService{client: client}
	client.Accounts = &AccountsService{client: client}
	client.BankFees = &BankFeesService{client: client}
	client.Contacts = &ContactsService{client: client}
	client.EArchives = &EArchivesService{client: client}
	client.EInvoiceInboxes = &EInvoiceInboxesService{client: client}
	client.EInvoices = &EInvoicesService{client: client}
	client.ESMMs = &ESMMsService{client: client}
	client.Employees = &EmployeesService{client: client}
	client.ItemCategories = &ItemCategoriesService{client: client}
	client.Products = &ProductsService{client: client}
	client.PurchaseBills = &PurchaseBillsService{client: client}
	client.Salaries = &SalariesService{client: client}
	client.SalesInvoices = &SalesInvoicesService{client: client}
	client.SalesOffers = &SalesOffersService{client: client}
	client.Sharings = &SharingsService{client: client}
	client.ShipmentDocuments = &ShipmentDocumentsService{client: client}
	client.StockMovements = &StockMovementsService{client: client}
	client.StockUpdates = &StockUpdatesService{client: client}
	client.Tags = &TagsService{client: client}
	client.Taxes = &TaxesService{client: client}
	client.TrackableJobs = &TrackableJobsService{client: client}
	client.Transactions = &TransactionsService{client: client}
	client.Warehouses = &WarehousesService{client: client}
	client.Webhooks = &WebhooksService{client: client}

	return client
}

// AuthorizeURL OAuth2 yetkilendirme URL'ini döndürür
func (c *Client) AuthorizeURL(state string) string {
	return c.config.AuthCodeURL(state)
}

// SetTokenFromCode yetkilendirme kodundan token alır
func (c *Client) SetTokenFromCode(ctx context.Context, code string) error {
	token, err := c.config.Exchange(ctx, code)
	if err != nil {
		return err
	}
	c.token = token
	c.httpClient = c.config.Client(ctx, token)
	return nil
}

// SetTokenFromPassword e-posta ve şifre ile token alır
func (c *Client) SetTokenFromPassword(ctx context.Context, email, password string) error {
	data := url.Values{}
	data.Set("grant_type", "password")
	data.Set("client_id", c.config.ClientID)
	data.Set("client_secret", c.config.ClientSecret)
	data.Set("username", email)
	data.Set("password", password)
	data.Set("redirect_uri", c.config.RedirectURL)

	resp, err := http.PostForm(TokenURL, data)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var token oauth2.Token
	if err := json.NewDecoder(resp.Body).Decode(&token); err != nil {
		return err
	}

	c.token = &token
	c.httpClient = c.config.Client(ctx, &token)
	return nil
}

// SetToken mevcut token'ı ayarlar
func (c *Client) SetToken(token *oauth2.Token) {
	c.token = token
	c.httpClient = c.config.Client(context.Background(), token)
}

// GetToken mevcut token'ı döndürür
func (c *Client) GetToken() *oauth2.Token {
	return c.token
}

// makeRequest HTTP isteği yapar
func (c *Client) makeRequest(ctx context.Context, method, path string, body interface{}) (*http.Response, error) {
	url := fmt.Sprintf("%s/%d%s", c.baseURL, c.companyID, path)

	var reqBody io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		reqBody = bytes.NewBuffer(jsonBody)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, reqBody)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	return c.httpClient.Do(req)
}

// get GET isteği yapar
func (c *Client) get(ctx context.Context, path string, params map[string]string) (*http.Response, error) {
	if len(params) > 0 {
		values := url.Values{}
		for k, v := range params {
			values.Add(k, v)
		}
		path += "?" + values.Encode()
	}
	return c.makeRequest(ctx, "GET", path, nil)
}

// post POST isteği yapar
func (c *Client) post(ctx context.Context, path string, body interface{}) (*http.Response, error) {
	return c.makeRequest(ctx, "POST", path, body)
}

// put PUT isteği yapar
func (c *Client) put(ctx context.Context, path string, body interface{}) (*http.Response, error) {
	return c.makeRequest(ctx, "PUT", path, body)
}

// patch PATCH isteği yapar
func (c *Client) patch(ctx context.Context, path string, body interface{}) (*http.Response, error) {
	return c.makeRequest(ctx, "PATCH", path, body)
}

// delete DELETE isteği yapar
func (c *Client) delete(ctx context.Context, path string) (*http.Response, error) {
	return c.makeRequest(ctx, "DELETE", path, nil)
}

// ListParams genel liste parametreleri
type ListParams struct {
	Page     int               `json:"page,omitempty"`
	PageSize int               `json:"page_size,omitempty"`
	Sort     string            `json:"sort,omitempty"`
	Filter   map[string]string `json:"filter,omitempty"`
}

// ToMap ListParams'ı map'e çevirir
func (lp *ListParams) ToMap() map[string]string {
	params := make(map[string]string)

	if lp.Page > 0 {
		params["page[number]"] = strconv.Itoa(lp.Page)
	}
	if lp.PageSize > 0 {
		params["page[size]"] = strconv.Itoa(lp.PageSize)
	}
	if lp.Sort != "" {
		params["sort"] = lp.Sort
	}
	if lp.Filter != nil {
		for k, v := range lp.Filter {
			params[fmt.Sprintf("filter[%s]", k)] = v
		}
	}

	return params
}

// Response genel API yanıt yapısı
type Response struct {
	Data interface{} `json:"data"`
	Meta *Meta       `json:"meta,omitempty"`
}

// Meta sayfalama bilgileri
type Meta struct {
	CurrentPage int `json:"current_page"`
	TotalPages  int `json:"total_pages"`
	TotalCount  int `json:"total_count"`
}

// Error API hata yapısı
type Error struct {
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

// ErrorResponse hata yanıt yapısı
type ErrorResponse struct {
	Errors []Error `json:"errors"`
}

func (e *ErrorResponse) Error() string {
	if len(e.Errors) > 0 {
		return fmt.Sprintf("%s: %s", e.Errors[0].Title, e.Errors[0].Detail)
	}
	return "Bilinmeyen hata"
}
