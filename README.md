# Paraşüt API v4 Golang SDK

Bu paket Paraşüt API v4'ü kullanmak için tasarlanmış kapsamlı bir Golang SDK'sıdır.

## Kurulum

```bash
go get github.com/parevo-lab/parasut
```

## Hızlı Başlangıç

### 1. Client Oluşturma

```go
package main

import (
    "context"
    "log"
    
    "github.com/parevo-lab/parasut"
)

func main() {
    config := &parasut.Config{
        ClientID:     "your-client-id",
        ClientSecret: "your-client-secret", 
        RedirectURL:  "urn:ietf:wg:oauth:2.0:oob",
        CompanyID:    123, // Firma ID'niz
    }
    
    client := parasut.NewClient(config)
    
    // E-posta ve şifre ile giriş
    ctx := context.Background()
    err := client.SetTokenFromPassword(ctx, "email@example.com", "password")
    if err != nil {
        log.Fatal(err)
    }
    
    // Artık API'yi kullanabilirsiniz
}
```

### 2. OAuth2 Akışı ile Giriş

```go
// 1. Kullanıcıyı yetkilendirme URL'ine yönlendirin
authURL := client.AuthorizeURL("state")
fmt.Println("Şu URL'ye gidin:", authURL)

// 2. Kullanıcı yetkilendirme kodunu aldıktan sonra
err := client.SetTokenFromCode(ctx, "authorization-code")
if err != nil {
    log.Fatal(err)
}
```

## API Kullanımı

### Hesaplar (Accounts)

```go
// Hesapları listele
accounts, meta, err := client.Accounts.List(ctx, &parasut.ListParams{
    Page: 1,
    PageSize: 10,
    Filter: map[string]string{
        "name": "Kasa",
    },
})

// Yeni hesap oluştur
account, err := client.Accounts.Create(ctx, parasut.AccountAttributes{
    Name:        "Yeni Hesap",
    Currency:    "TRL",
    AccountType: "cash",
})

// Hesap detayı getir
account, err := client.Accounts.Get(ctx, "account-id")

// Hesap güncelle
account, err := client.Accounts.Update(ctx, "account-id", parasut.AccountAttributes{
    Name: "Güncellenmiş Hesap",
})

// Hesap sil
err := client.Accounts.Delete(ctx, "account-id")
```

### Müşteri/Tedarikçiler (Contacts)

```go
// Müşterileri listele
contacts, meta, err := client.Contacts.List(ctx, &parasut.ListParams{
    Filter: map[string]string{
        "account_type": "customer",
    },
})

// Yeni müşteri oluştur
contact, err := client.Contacts.Create(ctx, parasut.ContactAttributes{
    Name:        "Müşteri Adı",
    Email:       "musteri@example.com",
    ContactType: "company",
    AccountType: "customer",
})

// Müşteri detayı getir
contact, err := client.Contacts.Get(ctx, "contact-id")

// Müşteri güncelle
contact, err := client.Contacts.Update(ctx, "contact-id", parasut.ContactAttributes{
    Name: "Güncellenmiş Müşteri",
})
```

### Ürünler (Products)

```go
// Ürünleri listele
products, meta, err := client.Products.List(ctx, &parasut.ListParams{
    Sort: "name",
})

// Yeni ürün oluştur
product, err := client.Products.Create(ctx, parasut.ProductAttributes{
    Code:     "PRD001",
    Name:     "Ürün Adı",
    VatRate:  18.0,
    ListPrice: 100.0,
    Currency: "TRL",
})

// Ürün detayı getir
product, err := client.Products.Get(ctx, "product-id")

// Ürün güncelle
product, err := client.Products.Update(ctx, "product-id", parasut.ProductAttributes{
    Name: "Güncellenmiş Ürün",
    ListPrice: 120.0,
})
```

### Satış Faturaları (Sales Invoices)

```go
// Satış faturalarını listele
invoices, meta, err := client.SalesInvoices.List(ctx, &parasut.ListParams{
    Filter: map[string]string{
        "item_type": "invoice",
    },
})

// Yeni satış faturası oluştur
invoice, err := client.SalesInvoices.Create(ctx, 
    parasut.SalesInvoiceAttributes{
        ItemType:     "invoice",
        Description:  "Test Faturası",
        IssueDate:    "2023-12-01",
        DueDate:      "2023-12-31",
        Currency:     "TRL",
    },
    &parasut.SalesInvoiceRelationships{
        Contact: &parasut.RelationshipData{
            ID:   "contact-id",
            Type: "contacts",
        },
    },
)

// Fatura detayı getir
invoice, err := client.SalesInvoices.Get(ctx, "invoice-id")

// Fatura iptal et
err := client.SalesInvoices.Cancel(ctx, "invoice-id")

// Fatura arşivle
err := client.SalesInvoices.Archive(ctx, "invoice-id")
```

### Alış Faturaları (Purchase Bills)

```go
// Alış faturalarını listele
bills, meta, err := client.PurchaseBills.List(ctx, &parasut.ListParams{
    Sort: "-issue_date",
})

// Yeni alış faturası oluştur
bill, err := client.PurchaseBills.Create(ctx,
    parasut.PurchaseBillAttributes{
        ItemType:     "bill",
        Description:  "Test Alış Faturası",
        IssueDate:    "2023-12-01",
        Currency:     "TRL",
    },
    &parasut.PurchaseBillRelationships{
        Supplier: &parasut.RelationshipData{
            ID:   "supplier-id",
            Type: "contacts",
        },
    },
)

// Fatura detayı getir
bill, err := client.PurchaseBills.Get(ctx, "bill-id")
```

### Banka Ücretleri (Bank Fees)

```go
// Banka ücretlerini listele
bankFees, meta, err := client.BankFees.List(ctx, nil)

// Yeni banka ücreti oluştur
bankFee, err := client.BankFees.Create(ctx, parasut.BankFeeAttributes{
    Description:  "Banka Komisyonu",
    Currency:     "TRL",
    IssueDate:    "2023-12-01",
    DueDate:      "2023-12-31",
    NetTotal:     50.0,
})

// Banka ücreti arşivle
err := client.BankFees.Archive(ctx, "bank-fee-id")

// Banka ücreti arşivden çıkar
err := client.BankFees.Unarchive(ctx, "bank-fee-id")
```

## Desteklenen Modüller

- ✅ **Accounts** (Hesaplar) - Tam CRUD desteği
- ✅ **BankFees** (Banka Ücretleri) - Tam CRUD desteği
- ✅ **Contacts** (Müşteri/Tedarikçiler) - Tam CRUD desteği
- ✅ **Products** (Ürünler) - Tam CRUD desteği
- ✅ **SalesInvoices** (Satış Faturaları) - Tam CRUD desteği
- ✅ **PurchaseBills** (Alış Faturaları) - Tam CRUD desteği
- 🚧 **EArchives** (E-Arşiv) - Planlanan
- 🚧 **EInvoices** (E-Faturalar) - Planlanan
- 🚧 **ESMMs** (E-SMM) - Planlanan
- 🚧 **Employees** (Çalışanlar) - Planlanan
- 🚧 **ItemCategories** (Ürün Kategorileri) - Planlanan
- 🚧 **Salaries** (Maaşlar) - Planlanan
- 🚧 **SalesOffers** (Satış Teklifleri) - Planlanan
- 🚧 **Sharings** (Paylaşımlar) - Planlanan
- 🚧 **ShipmentDocuments** (Sevkiyat Belgeleri) - Planlanan
- 🚧 **StockMovements** (Stok Hareketleri) - Planlanan
- 🚧 **StockUpdates** (Stok Güncellemeleri) - Planlanan
- 🚧 **Tags** (Etiketler) - Planlanan
- 🚧 **Taxes** (Vergiler) - Planlanan
- 🚧 **TrackableJobs** (İzlenebilir İşler) - Planlanan
- 🚧 **Transactions** (İşlemler) - Planlanan
- 🚧 **Warehouses** (Depolar) - Planlanan
- 🚧 **Webhooks** (Webhooklar) - Planlanan

## Hata Yönetimi

```go
// API hataları ErrorResponse türünde döner
accounts, meta, err := client.Accounts.List(ctx, nil)
if err != nil {
    if apiErr, ok := err.(*parasut.ErrorResponse); ok {
        for _, e := range apiErr.Errors {
            fmt.Printf("Hata: %s - %s\n", e.Title, e.Detail)
        }
    } else {
        fmt.Printf("Sistem hatası: %v\n", err)
    }
}
```

## Filtreleme ve Sayfalama

```go
params := &parasut.ListParams{
    Page:     1,
    PageSize: 25,
    Sort:     "-created_at", // Azalan sıralama için - kullanın
    Filter: map[string]string{
        "name":     "Arama terimi",
        "currency": "TRL",
    },
}

accounts, meta, err := client.Accounts.List(ctx, params)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Toplam sayfa: %d\n", meta.TotalPages)
fmt.Printf("Toplam kayıt: %d\n", meta.TotalCount)
fmt.Printf("Mevcut sayfa: %d\n", meta.CurrentPage)
```

## Token Yönetimi

```go
// Token'ı kaydet
token := client.GetToken()
// Token'ı veritabanına veya dosyaya kaydedin

// Token'ı geri yükle
client.SetToken(savedToken)
```

## Gereksinimler

- Go 1.21 veya üzeri
- Paraşüt API v4 erişim bilgileri

## Lisans

MIT License

## Katkıda Bulunma

1. Fork edin
2. Feature branch oluşturun (`git checkout -b feature/amazing-feature`)
3. Değişikliklerinizi commit edin (`git commit -m 'Add some amazing feature'`)
4. Branch'i push edin (`git push origin feature/amazing-feature`)
5. Pull Request oluşturun

## Destek

Sorunlarınız için [GitHub Issues](https://github.com/parevo-lab/parasut/issues) kullanabilirsiniz.

## Yazar

[Parevo Lab](https://github.com/parevo-lab) 