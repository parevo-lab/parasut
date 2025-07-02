package parasut

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"golang.org/x/oauth2"
)

func TestNewClient(t *testing.T) {
	config := &Config{
		ClientID:     "test-client-id",
		ClientSecret: "test-client-secret",
		RedirectURL:  "http://localhost:8080/callback",
		CompanyID:    123,
	}

	client := NewClient(config)

	if client == nil {
		t.Fatal("NewClient beklenmeyen şekilde nil döndü")
	}

	if client.baseURL != BaseURL {
		t.Errorf("baseURL = %s, beklenen %s", client.baseURL, BaseURL)
	}

	if client.companyID != config.CompanyID {
		t.Errorf("companyID = %d, beklenen %d", client.companyID, config.CompanyID)
	}

	if client.config.ClientID != config.ClientID {
		t.Errorf("config.ClientID = %s, beklenen %s", client.config.ClientID, config.ClientID)
	}

	// Servislerin başlatıldığını kontrol et
	if client.Me == nil {
		t.Error("Me servisi başlatılmadı")
	}
	if client.Accounts == nil {
		t.Error("Accounts servisi başlatılmadı")
	}
	if client.Contacts == nil {
		t.Error("Contacts servisi başlatılmadı")
	}
}

func TestClient_AuthorizeURL(t *testing.T) {
	config := &Config{
		ClientID:     "test-client-id",
		ClientSecret: "test-client-secret",
		RedirectURL:  "http://localhost:8080/callback",
		CompanyID:    123,
	}

	client := NewClient(config)
	state := "test-state"

	authURL := client.AuthorizeURL(state)

	if authURL == "" {
		t.Fatal("AuthorizeURL boş string döndü")
	}

	parsedURL, err := url.Parse(authURL)
	if err != nil {
		t.Fatalf("AuthorizeURL geçersiz URL döndü: %v", err)
	}

	if !strings.Contains(parsedURL.Host, "api.parasut.com") {
		t.Errorf("AuthorizeURL yanlış host içeriyor: %s", parsedURL.Host)
	}

	queryParams := parsedURL.Query()
	if queryParams.Get("state") != state {
		t.Errorf("state parametresi = %s, beklenen %s", queryParams.Get("state"), state)
	}
}

func TestClient_SetTokenFromPassword(t *testing.T) {
	// Mock token response
	tokenResponse := map[string]interface{}{
		"access_token":  "test-access-token",
		"token_type":    "Bearer",
		"expires_in":    3600,
		"refresh_token": "test-refresh-token",
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/oauth/token" {
			t.Errorf("Yanlış endpoint çağrıldı: %s", r.URL.Path)
		}

		if r.Method != "POST" {
			t.Errorf("Yanlış HTTP metodu: %s, beklenen POST", r.Method)
		}

		// Form verilerini kontrol et
		err := r.ParseForm()
		if err != nil {
			t.Fatalf("Form parse edilemedi: %v", err)
		}

		if r.Form.Get("grant_type") != "password" {
			t.Errorf("grant_type = %s, beklenen password", r.Form.Get("grant_type"))
		}

		if r.Form.Get("username") != "test@example.com" {
			t.Errorf("username = %s, beklenen test@example.com", r.Form.Get("username"))
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tokenResponse)
	}))
	defer server.Close()

	config := &Config{
		ClientID:     "test-client-id",
		ClientSecret: "test-client-secret",
		RedirectURL:  "http://localhost:8080/callback",
		CompanyID:    123,
	}

	client := NewClient(config)

	// Token URL'ini test sunucusuna yönlendir (const olduğu için config üzerinden yapıyoruz)

	// Test için geçici çözüm - client'ın config'ini güncelle
	client.config.Endpoint.TokenURL = server.URL + "/oauth/token"

	ctx := context.Background()
	err := client.SetTokenFromPassword(ctx, "test@example.com", "test-password")

	if err != nil {
		t.Fatalf("SetTokenFromPassword hata döndü: %v", err)
	}

	if client.token == nil {
		t.Fatal("Token ayarlanmadı")
	}

	if client.token.AccessToken != "test-access-token" {
		t.Errorf("Access token = %s, beklenen test-access-token", client.token.AccessToken)
	}
}

func TestClient_SetToken(t *testing.T) {
	config := &Config{
		ClientID:     "test-client-id",
		ClientSecret: "test-client-secret",
		RedirectURL:  "http://localhost:8080/callback",
		CompanyID:    123,
	}

	client := NewClient(config)

	token := &oauth2.Token{
		AccessToken:  "test-access-token",
		TokenType:    "Bearer",
		RefreshToken: "test-refresh-token",
		Expiry:       time.Now().Add(time.Hour),
	}

	client.SetToken(token)

	if client.token != token {
		t.Error("Token doğru şekilde ayarlanmadı")
	}

	if client.httpClient == nil {
		t.Error("HTTP client güncellenmedi")
	}
}

func TestClient_GetToken(t *testing.T) {
	config := &Config{
		ClientID:     "test-client-id",
		ClientSecret: "test-client-secret",
		RedirectURL:  "http://localhost:8080/callback",
		CompanyID:    123,
	}

	client := NewClient(config)

	// Token ayarlanmadan önce
	if client.GetToken() != nil {
		t.Error("Token ayarlanmadan önce nil dönmeli")
	}

	// Token ayarla
	token := &oauth2.Token{
		AccessToken: "test-access-token",
	}
	client.SetToken(token)

	// Token'ı al
	retrievedToken := client.GetToken()
	if retrievedToken != token {
		t.Error("GetToken yanlış token döndü")
	}
}

func TestClient_makeRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Request'i kontrol et
		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("Content-Type = %s, beklenen application/json", r.Header.Get("Content-Type"))
		}

		if r.Header.Get("Accept") != "application/json" {
			t.Errorf("Accept = %s, beklenen application/json", r.Header.Get("Accept"))
		}

		expectedPath := "/v4/123/test-path"
		if r.URL.Path != expectedPath {
			t.Errorf("Path = %s, beklenen %s", r.URL.Path, expectedPath)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "ok"}`))
	}))
	defer server.Close()

	config := &Config{
		ClientID:     "test-client-id",
		ClientSecret: "test-client-secret",
		RedirectURL:  "http://localhost:8080/callback",
		CompanyID:    123,
	}

	client := NewClient(config)
	client.baseURL = server.URL + "/v4" // Test sunucusunu kullan

	ctx := context.Background()
	resp, err := client.makeRequest(ctx, "GET", "/test-path", nil)

	if err != nil {
		t.Fatalf("makeRequest hata döndü: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status code = %d, beklenen %d", resp.StatusCode, http.StatusOK)
	}

	resp.Body.Close()
}

func TestClient_get(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Method = %s, beklenen GET", r.Method)
		}

		// Query parametrelerini kontrol et
		if r.URL.Query().Get("page") != "1" {
			t.Errorf("page parametresi = %s, beklenen 1", r.URL.Query().Get("page"))
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"data": []}`))
	}))
	defer server.Close()

	config := &Config{
		ClientID:     "test-client-id",
		ClientSecret: "test-client-secret",
		RedirectURL:  "http://localhost:8080/callback",
		CompanyID:    123,
	}

	client := NewClient(config)
	client.baseURL = server.URL + "/v4"

	ctx := context.Background()
	params := map[string]string{
		"page": "1",
	}

	resp, err := client.get(ctx, "/test", params)

	if err != nil {
		t.Fatalf("get hata döndü: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status code = %d, beklenen %d", resp.StatusCode, http.StatusOK)
	}

	resp.Body.Close()
}

func TestClient_post(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Method = %s, beklenen POST", r.Method)
		}

		// Request body'sini kontrol et
		var body map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			t.Fatalf("Request body decode edilemedi: %v", err)
		}

		if body["test"] != "data" {
			t.Errorf("Request body yanlış: %v", body)
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"id": "123"}`))
	}))
	defer server.Close()

	config := &Config{
		ClientID:     "test-client-id",
		ClientSecret: "test-client-secret",
		RedirectURL:  "http://localhost:8080/callback",
		CompanyID:    123,
	}

	client := NewClient(config)
	client.baseURL = server.URL + "/v4"

	ctx := context.Background()
	body := map[string]interface{}{
		"test": "data",
	}

	resp, err := client.post(ctx, "/test", body)

	if err != nil {
		t.Fatalf("post hata döndü: %v", err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Status code = %d, beklenen %d", resp.StatusCode, http.StatusCreated)
	}

	resp.Body.Close()
}

func TestListParams_ToMap(t *testing.T) {
	tests := []struct {
		name   string
		params ListParams
		want   map[string]string
	}{
		{
			name:   "Boş parametreler",
			params: ListParams{},
			want:   map[string]string{},
		},
		{
			name: "Sayfa parametreleri",
			params: ListParams{
				Page:     1,
				PageSize: 10,
			},
			want: map[string]string{
				"page":      "1",
				"page_size": "10",
			},
		},
		{
			name: "Tüm parametreler",
			params: ListParams{
				Page:     2,
				PageSize: 20,
				Sort:     "name",
				Filter: map[string]string{
					"status": "active",
					"type":   "customer",
				},
			},
			want: map[string]string{
				"page":      "2",
				"page_size": "20",
				"sort":      "name",
				"status":    "active",
				"type":      "customer",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.params.ToMap()

			if len(got) != len(tt.want) {
				t.Errorf("ToMap() döndürülen map uzunluğu = %d, beklenen %d", len(got), len(tt.want))
			}

			for key, expectedValue := range tt.want {
				if actualValue, exists := got[key]; !exists {
					t.Errorf("ToMap() '%s' anahtarı bulunamadı", key)
				} else if actualValue != expectedValue {
					t.Errorf("ToMap()[%s] = %s, beklenen %s", key, actualValue, expectedValue)
				}
			}
		})
	}
}

func TestErrorResponse_Error(t *testing.T) {
	tests := []struct {
		name      string
		errorResp ErrorResponse
		want      string
	}{
		{
			name: "Tek hata",
			errorResp: ErrorResponse{
				Errors: []Error{
					{Title: "Validation Error", Detail: "Name is required"},
				},
			},
			want: "Validation Error: Name is required",
		},
		{
			name: "Çoklu hatalar",
			errorResp: ErrorResponse{
				Errors: []Error{
					{Title: "Validation Error", Detail: "Name is required"},
					{Title: "Authorization Error", Detail: "Invalid token"},
				},
			},
			want: "Validation Error: Name is required; Authorization Error: Invalid token",
		},
		{
			name: "Boş hatalar",
			errorResp: ErrorResponse{
				Errors: []Error{},
			},
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.errorResp.Error()
			if got != tt.want {
				t.Errorf("ErrorResponse.Error() = %s, beklenen %s", got, tt.want)
			}
		})
	}
}

func TestClient_HTTPMethods(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		expectedMethod string
		expectedStatus int
	}{
		{"PUT method", "put", "PUT", http.StatusOK},
		{"PATCH method", "patch", "PATCH", http.StatusOK},
		{"DELETE method", "delete", "DELETE", http.StatusNoContent},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != tt.expectedMethod {
					t.Errorf("Method = %s, beklenen %s", r.Method, tt.expectedMethod)
				}
				w.WriteHeader(tt.expectedStatus)
			}))
			defer server.Close()

			config := &Config{
				ClientID:     "test-client-id",
				ClientSecret: "test-client-secret",
				RedirectURL:  "http://localhost:8080/callback",
				CompanyID:    123,
			}

			client := NewClient(config)
			client.baseURL = server.URL + "/v4"

			ctx := context.Background()
			var resp *http.Response
			var err error

			switch tt.method {
			case "put":
				resp, err = client.put(ctx, "/test", map[string]string{"test": "data"})
			case "patch":
				resp, err = client.patch(ctx, "/test", map[string]string{"test": "data"})
			case "delete":
				resp, err = client.delete(ctx, "/test")
			}

			if err != nil {
				t.Fatalf("%s hata döndü: %v", tt.method, err)
			}

			if resp.StatusCode != tt.expectedStatus {
				t.Errorf("Status code = %d, beklenen %d", resp.StatusCode, tt.expectedStatus)
			}

			resp.Body.Close()
		})
	}
}
