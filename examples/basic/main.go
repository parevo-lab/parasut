package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/parevo-lab/parasut"
)

func main() {
	// Çevre değişkenlerinden API bilgilerini al
	clientID := os.Getenv("PARASUT_CLIENT_ID")
	clientSecret := os.Getenv("PARASUT_CLIENT_SECRET")
	email := os.Getenv("PARASUT_EMAIL")
	password := os.Getenv("PARASUT_PASSWORD")
	companyID := 123 // Firma ID'nizi buraya yazın

	if clientID == "" || clientSecret == "" || email == "" || password == "" {
		log.Fatal("PARASUT_CLIENT_ID, PARASUT_CLIENT_SECRET, PARASUT_EMAIL ve PARASUT_PASSWORD çevre değişkenlerini ayarlayın")
	}

	// Client oluştur
	config := &parasut.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  "urn:ietf:wg:oauth:2.0:oob",
		CompanyID:    companyID,
	}

	client := parasut.NewClient(config)

	// E-posta ve şifre ile giriş yap
	ctx := context.Background()
	err := client.SetTokenFromPassword(ctx, email, password)
	if err != nil {
		log.Fatalf("Giriş yapılamadı: %v", err)
	}

	fmt.Println("✅ Başarıyla giriş yapıldı!")

	// Hesapları listele
	fmt.Println("\n📋 Hesaplar listeleniyor...")
	accounts, meta, err := client.Accounts.List(ctx, &parasut.ListParams{
		Page:     1,
		PageSize: 5,
	})
	if err != nil {
		log.Fatalf("Hesaplar listelenemedi: %v", err)
	}

	fmt.Printf("Toplam %d hesap bulundu (Sayfa %d/%d):\n", meta.TotalCount, meta.CurrentPage, meta.TotalPages)
	for _, account := range accounts {
		fmt.Printf("- %s (%s) - Bakiye: %.2f %s\n",
			account.Attributes.Name,
			account.Attributes.AccountType,
			account.Attributes.Balance,
			account.Attributes.Currency)
	}

	// Müşterileri listele
	fmt.Println("\n👥 Müşteriler listeleniyor...")
	contacts, meta, err := client.Contacts.List(ctx, &parasut.ListParams{
		Page:     1,
		PageSize: 5,
		Filter: map[string]string{
			"account_type": "customer",
		},
	})
	if err != nil {
		log.Fatalf("Müşteriler listelenemedi: %v", err)
	}

	fmt.Printf("Toplam %d müşteri bulundu:\n", meta.TotalCount)
	for _, contact := range contacts {
		fmt.Printf("- %s (%s)\n", contact.Attributes.Name, contact.Attributes.Email)
	}

	// Ürünleri listele
	fmt.Println("\n📦 Ürünler listeleniyor...")
	products, meta, err := client.Products.List(ctx, &parasut.ListParams{
		Page:     1,
		PageSize: 5,
	})
	if err != nil {
		log.Fatalf("Ürünler listelenemedi: %v", err)
	}

	fmt.Printf("Toplam %d ürün bulundu:\n", meta.TotalCount)
	for _, product := range products {
		fmt.Printf("- %s (%s) - Fiyat: %.2f %s\n",
			product.Attributes.Name,
			product.Attributes.Code,
			product.Attributes.ListPrice,
			product.Attributes.Currency)
	}

	// Satış faturalarını listele
	fmt.Println("\n🧾 Satış faturaları listeleniyor...")
	invoices, meta, err := client.SalesInvoices.List(ctx, &parasut.ListParams{
		Page:     1,
		PageSize: 5,
		Sort:     "-issue_date",
	})
	if err != nil {
		log.Fatalf("Satış faturaları listelenemedi: %v", err)
	}

	fmt.Printf("Toplam %d satış faturası bulundu:\n", meta.TotalCount)
	for _, invoice := range invoices {
		fmt.Printf("- %s - %s - Tutar: %.2f TL\n",
			invoice.Attributes.Description,
			invoice.Attributes.IssueDate,
			invoice.Attributes.GrossTotal)
	}

	fmt.Println("\n✨ Örnek tamamlandı!")
}
