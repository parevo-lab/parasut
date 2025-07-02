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

### Kullanıcı Bilgileri (Me)

```go
// Kullanıcı bilgilerini getir
me, err := client.Me.Get(ctx)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Kullanıcı: %s\n", me.Data.Attributes.Name)
```

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

// Hesap işlemlerini getir
transactions, meta, err := client.Accounts.GetTransactions(ctx, "account-id")

// Borç işlemi oluştur
transaction, err := client.Accounts.CreateDebitTransaction(ctx, "account-id", parasut.AccountTransactionAttributes{
    Description: "Borç işlemi",
    Debit:       100.0,
    Date:        "2023-12-01",
})

// Alacak işlemi oluştur
transaction, err := client.Accounts.CreateCreditTransaction(ctx, "account-id", parasut.AccountTransactionAttributes{
    Description: "Alacak işlemi",
    Credit:      100.0,
    Date:        "2023-12-01",
})
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

// Müşteri sil
err := client.Contacts.Delete(ctx, "contact-id")

// Müşteri borç işlemlerini getir
debitTransactions, meta, err := client.Contacts.GetDebitTransactions(ctx, "contact-id")

// Müşteri alacak işlemlerini getir
creditTransactions, meta, err := client.Contacts.GetCreditTransactions(ctx, "contact-id")
```

### Ürünler (Products)

```go
// Ürünleri listele
products, meta, err := client.Products.List(ctx, &parasut.ListParams{
    Sort: "name",
})

// Yeni ürün oluştur
product, err := client.Products.Create(ctx, parasut.ProductAttributes{
    Code:      "PRD001",
    Name:      "Ürün Adı",
    VatRate:   18.0,
    ListPrice: 100.0,
    Currency:  "TRL",
})

// Ürün detayı getir
product, err := client.Products.Get(ctx, "product-id")

// Ürün güncelle
product, err := client.Products.Update(ctx, "product-id", parasut.ProductAttributes{
    Name:      "Güncellenmiş Ürün",
    ListPrice: 120.0,
})

// Ürün sil
err := client.Products.Delete(ctx, "product-id")

// Ürün stok seviyelerini getir
inventoryLevels, meta, err := client.Products.GetInventoryLevels(ctx, "product-id")
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
        ItemType:    "invoice",
        Description: "Test Faturası",
        IssueDate:   "2023-12-01",
        DueDate:     "2023-12-31",
        Currency:    "TRL",
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

// Fatura güncelle
invoice, err := client.SalesInvoices.Update(ctx, "invoice-id",
    parasut.SalesInvoiceAttributes{
        Description: "Güncellenmiş Fatura",
    },
    nil,
)

// Fatura iptal et
err := client.SalesInvoices.Cancel(ctx, "invoice-id")

// Fatura kurtar
err := client.SalesInvoices.Recover(ctx, "invoice-id")

// Fatura arşivle
err := client.SalesInvoices.Archive(ctx, "invoice-id")

// Fatura arşivden çıkar
err := client.SalesInvoices.Unarchive(ctx, "invoice-id")

// Faturaya ödeme ekle
payment, err := client.SalesInvoices.CreatePayment(ctx, "invoice-id", parasut.PaymentAttributes{
    Date:   "2023-12-01",
    Amount: 100.0,
})

// Faturayı faturaya dönüştür
invoice, err := client.SalesInvoices.ConvertToInvoice(ctx, "invoice-id")

// Fatura PDF'ini getir
pdfData, err := client.SalesInvoices.GetPDF(ctx, "invoice-id")
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
        ItemType:    "bill",
        Description: "Test Alış Faturası",
        IssueDate:   "2023-12-01",
        Currency:    "TRL",
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

// Fatura güncelle
bill, err := client.PurchaseBills.Update(ctx, "bill-id",
    parasut.PurchaseBillAttributes{
        Description: "Güncellenmiş Alış Faturası",
    },
    nil,
)

// Faturaya ödeme ekle
payment, err := client.PurchaseBills.CreatePayment(ctx, "bill-id", parasut.PaymentAttributes{
    Date:   "2023-12-01",
    Amount: 100.0,
})

// Fatura iptal et
err := client.PurchaseBills.Cancel(ctx, "bill-id")

// Fatura kurtar
err := client.PurchaseBills.Recover(ctx, "bill-id")

// Fatura arşivle
err := client.PurchaseBills.Archive(ctx, "bill-id")

// Fatura arşivden çıkar
err := client.PurchaseBills.Unarchive(ctx, "bill-id")

// Fatura PDF'ini getir
pdfData, err := client.PurchaseBills.GetPDF(ctx, "bill-id")
```

### Banka Ücretleri (Bank Fees)

```go
// Banka ücretlerini listele
bankFees, meta, err := client.BankFees.List(ctx, nil)

// Yeni banka ücreti oluştur
bankFee, err := client.BankFees.Create(ctx, parasut.BankFeeAttributes{
    Description: "Banka Komisyonu",
    Currency:    "TRL",
    IssueDate:   "2023-12-01",
    DueDate:     "2023-12-31",
    NetTotal:    50.0,
})

// Banka ücreti detayı getir
bankFee, err := client.BankFees.Get(ctx, "bank-fee-id")

// Banka ücreti güncelle
bankFee, err := client.BankFees.Update(ctx, "bank-fee-id", parasut.BankFeeAttributes{
    Description: "Güncellenmiş Banka Komisyonu",
})

// Banka ücreti arşivle
err := client.BankFees.Archive(ctx, "bank-fee-id")

// Banka ücreti arşivden çıkar
err := client.BankFees.Unarchive(ctx, "bank-fee-id")

// Banka ücretine ödeme ekle
payment, err := client.BankFees.CreatePayment(ctx, "bank-fee-id", parasut.PaymentAttributes{
    Date:   "2023-12-01",
    Amount: 50.0,
})
```

### Çalışanlar (Employees)

```go
// Çalışanları listele
employees, meta, err := client.Employees.List(ctx, nil)

// Yeni çalışan oluştur
employee, err := client.Employees.Create(ctx, parasut.EmployeeAttributes{
    Name:  "Çalışan Adı",
    Email: "calisan@example.com",
})

// Çalışan detayı getir
employee, err := client.Employees.Get(ctx, "employee-id")

// Çalışan güncelle
employee, err := client.Employees.Update(ctx, "employee-id", parasut.EmployeeAttributes{
    Name: "Güncellenmiş Çalışan",
})

// Çalışan arşivle
err := client.Employees.Archive(ctx, "employee-id")

// Çalışan arşivden çıkar
err := client.Employees.Unarchive(ctx, "employee-id")
```

### Maaşlar (Salaries)

```go
// Maaşları listele
salaries, meta, err := client.Salaries.List(ctx, nil)

// Yeni maaş oluştur
salary, err := client.Salaries.Create(ctx,
    parasut.SalaryAttributes{
        Description: "Ocak 2023 Maaşı",
        Date:        "2023-01-31",
        NetTotal:    5000.0,
    },
    &parasut.SalaryRelationships{
        Employee: &parasut.RelationshipData{
            ID:   "employee-id",
            Type: "employees",
        },
    },
)

// Maaş detayı getir
salary, err := client.Salaries.Get(ctx, "salary-id")

// Maaş güncelle
salary, err := client.Salaries.Update(ctx, "salary-id",
    parasut.SalaryAttributes{
        NetTotal: 5500.0,
    },
    nil,
)

// Maaş arşivle
err := client.Salaries.Archive(ctx, "salary-id")

// Maaş arşivden çıkar
err := client.Salaries.Unarchive(ctx, "salary-id")

// Maaşa ödeme ekle
payment, err := client.Salaries.CreatePayment(ctx, "salary-id", parasut.PaymentAttributes{
    Date:   "2023-01-31",
    Amount: 5000.0,
})
```

### Vergiler (Taxes)

```go
// Vergileri listele
taxes, meta, err := client.Taxes.List(ctx, nil)

// Yeni vergi oluştur
tax, err := client.Taxes.Create(ctx,
    parasut.TaxAttributes{
        Description: "KDV Beyannamesi",
        Date:        "2023-12-31",
        NetTotal:    1000.0,
    },
    &parasut.TaxRelationships{
        // İlişkiler burada tanımlanır
    },
)

// Vergi detayı getir
tax, err := client.Taxes.Get(ctx, "tax-id")

// Vergi güncelle
tax, err := client.Taxes.Update(ctx, "tax-id",
    parasut.TaxAttributes{
        NetTotal: 1200.0,
    },
    nil,
)

// Vergi arşivle
err := client.Taxes.Archive(ctx, "tax-id")

// Vergi arşivden çıkar
err := client.Taxes.Unarchive(ctx, "tax-id")

// Vergiye ödeme ekle
payment, err := client.Taxes.CreatePayment(ctx, "tax-id", parasut.PaymentAttributes{
    Date:   "2023-12-31",
    Amount: 1000.0,
})
```

### Etiketler (Tags)

```go
// Etiketleri listele
tags, meta, err := client.Tags.List(ctx, nil)

// Yeni etiket oluştur
tag, err := client.Tags.Create(ctx, parasut.TagAttributes{
    Name: "Önemli",
    Color: "#FF0000",
})

// Etiket detayı getir
tag, err := client.Tags.Get(ctx, "tag-id")

// Etiket güncelle
tag, err := client.Tags.Update(ctx, "tag-id", parasut.TagAttributes{
    Name: "Çok Önemli",
})

// Etiket sil
err := client.Tags.Delete(ctx, "tag-id")
```

### Depolar (Warehouses)

```go
// Depoları listele
warehouses, meta, err := client.Warehouses.List(ctx, nil)

// Yeni depo oluştur
warehouse, err := client.Warehouses.Create(ctx, parasut.WarehouseAttributes{
    Name: "Ana Depo",
    City: "İstanbul",
})

// Depo detayı getir
warehouse, err := client.Warehouses.Get(ctx, "warehouse-id")

// Depo güncelle
warehouse, err := client.Warehouses.Update(ctx, "warehouse-id", parasut.WarehouseAttributes{
    Name: "Güncellenmiş Ana Depo",
})

// Depo sil
err := client.Warehouses.Delete(ctx, "warehouse-id")
```

### Stok Hareketleri (Stock Movements)

```go
// Stok hareketlerini listele
stockMovements, meta, err := client.StockMovements.List(ctx, &parasut.ListParams{
    Sort: "-date",
})
```

### Stok Güncellemeleri (Stock Updates)

```go
// Stok güncellemesi oluştur
stockUpdate, err := client.StockUpdates.Create(ctx, parasut.StockUpdateAttributes{
    ProductId:   "product-id",
    WarehouseId: "warehouse-id",
    Quantity:    100,
})
```

### Ürün Kategorileri (Item Categories)

```go
// Ürün kategorilerini listele
categories, meta, err := client.ItemCategories.List(ctx, nil)

// Yeni kategori oluştur
category, err := client.ItemCategories.Create(ctx, parasut.ItemCategoryAttributes{
    Name: "Elektronik",
})

// Kategori detayı getir
category, err := client.ItemCategories.Get(ctx, "category-id")

// Kategori güncelle
category, err := client.ItemCategories.Update(ctx, "category-id", parasut.ItemCategoryAttributes{
    Name: "Elektronik Ürünler",
})

// Kategori sil
err := client.ItemCategories.Delete(ctx, "category-id")
```

### Satış Teklifleri (Sales Offers)

```go
// Satış tekliflerini listele
offers, meta, err := client.SalesOffers.List(ctx, nil)

// Yeni satış teklifi oluştur
offer, err := client.SalesOffers.Create(ctx,
    parasut.SalesOfferAttributes{
        Description: "Test Teklifi",
        IssueDate:   "2023-12-01",
        ExpiryDate:  "2023-12-31",
    },
    &parasut.SalesOfferRelationships{
        Contact: &parasut.RelationshipData{
            ID:   "contact-id",
            Type: "contacts",
        },
    },
)

// Teklif detayı getir
offer, err := client.SalesOffers.Get(ctx, "offer-id")

// Teklif güncelle
offer, err := client.SalesOffers.Update(ctx, "offer-id",
    parasut.SalesOfferAttributes{
        Description: "Güncellenmiş Teklif",
    },
    nil,
)

// Teklif sil
err := client.SalesOffers.Delete(ctx, "offer-id")

// Teklif arşivle
err := client.SalesOffers.Archive(ctx, "offer-id")

// Teklif arşivden çıkar
err := client.SalesOffers.Unarchive(ctx, "offer-id")

// Teklif PDF'ini getir
pdfData, err := client.SalesOffers.GetPDF(ctx, "offer-id")

// Teklif detaylarını getir
offer, err := client.SalesOffers.GetDetails(ctx, "offer-id")

// Teklif durumunu güncelle
offer, err := client.SalesOffers.UpdateStatus(ctx, "offer-id", "accepted")
```

### E-Arşiv (E-Archives)

```go
// E-arşiv belgelerini listele
eArchives, meta, err := client.EArchives.List(ctx, nil)

// E-arşiv belge detayı getir
eArchive, err := client.EArchives.Get(ctx, "e-archive-id")

// E-arşiv PDF'ini getir
pdfData, err := client.EArchives.GetPDF(ctx, "e-archive-id")
```

### E-Fatura Gelen Kutusu (E-Invoice Inboxes)

```go
// E-fatura gelen kutusunu listele
inboxes, meta, err := client.EInvoiceInboxes.List(ctx, nil)
```

### E-Faturalar (E-Invoices)

```go
// E-faturaları listele
eInvoices, meta, err := client.EInvoices.List(ctx, nil)

// E-fatura detayı getir
eInvoice, err := client.EInvoices.Get(ctx, "e-invoice-id")

// Yeni e-fatura oluştur
eInvoice, err := client.EInvoices.Create(ctx, parasut.EInvoiceAttributes{
    VknTckn:     "1234567890",
    InvoiceType: "SATIS",
})

// E-fatura PDF'ini getir
pdfData, err := client.EInvoices.GetPDF(ctx, "e-invoice-id")
```

### E-SMM (Electronic Cargo Waybill)

```go
// E-SMM belgelerini listele
esmms, meta, err := client.ESMMs.List(ctx, nil)

// E-SMM belge detayı getir
esmm, err := client.ESMMs.Get(ctx, "esmm-id")

// Yeni E-SMM oluştur
esmm, err := client.ESMMs.Create(ctx, parasut.ESMMAttributes{
    CarrierVknTckn: "1234567890",
    CarrierTitle:   "Kargo Şirketi",
})

// E-SMM PDF'ini getir
pdfData, err := client.ESMMs.GetPDF(ctx, "esmm-id")
```

### Paylaşımlar (Sharings)

```go
// Paylaşımları listele
sharings, meta, err := client.Sharings.List(ctx, nil)

// Yeni paylaşım oluştur
sharing, err := client.Sharings.Create(ctx, parasut.SharingAttributes{
    ShareableType: "sales_invoices",
    ShareableId:   "invoice-id",
    SharedWithId:  "user-id",
})
```

### Sevkiyat Belgeleri (Shipment Documents)

```go
// Sevkiyat belgelerini listele
documents, meta, err := client.ShipmentDocuments.List(ctx, nil)

// Sevkiyat belgesi detayı getir
document, err := client.ShipmentDocuments.Get(ctx, "document-id")

// Yeni sevkiyat belgesi oluştur
document, err := client.ShipmentDocuments.Create(ctx, parasut.ShipmentDocumentAttributes{
    ShipmentDate: "2023-12-01",
    VehiclePlate: "34 ABC 123",
})

// Sevkiyat belgesi güncelle
document, err := client.ShipmentDocuments.Update(ctx, "document-id", parasut.ShipmentDocumentAttributes{
    VehiclePlate: "34 XYZ 789",
})

// Sevkiyat belgesi sil
err := client.ShipmentDocuments.Delete(ctx, "document-id")
```

### İzlenebilir İşler (Trackable Jobs)

```go
// İzlenebilir iş detayı getir
job, err := client.TrackableJobs.Get(ctx, "job-id")
```

### İşlemler (Transactions)

```go
// İşlem detayı getir
transaction, err := client.Transactions.Get(ctx, "transaction-id")

// İşlem güncelle
transaction, err := client.Transactions.Update(ctx, "transaction-id", parasut.TransactionAttributes{
    Description: "Güncellenmiş İşlem",
})

// İşlem sil
err := client.Transactions.Delete(ctx, "transaction-id")
```

### Webhooklar (Webhooks)

```go
// Webhookları listele
webhooks, meta, err := client.Webhooks.List(ctx, nil)

// Webhook detayı getir
webhook, err := client.Webhooks.Get(ctx, "webhook-id")

// Yeni webhook oluştur
webhook, err := client.Webhooks.Create(ctx, parasut.WebhookAttributes{
    URL:    "https://example.com/webhook",
    Events: []string{"sales_invoice.created", "contact.updated"},
})

// Webhook güncelle
webhook, err := client.Webhooks.Update(ctx, "webhook-id", parasut.WebhookAttributes{
    URL: "https://example.com/new-webhook",
})

// Webhook sil
err := client.Webhooks.Delete(ctx, "webhook-id")
```

## Desteklenen Modüller

- ✅ **Me** (Kullanıcı Bilgileri) - Tam destek
- ✅ **Accounts** (Hesaplar) - Tam CRUD desteği + İşlemler
- ✅ **BankFees** (Banka Ücretleri) - Tam CRUD desteği + Ödeme
- ✅ **Contacts** (Müşteri/Tedarikçiler) - Tam CRUD desteği + İşlemler
- ✅ **Products** (Ürünler) - Tam CRUD desteği + Stok Seviyeleri
- ✅ **SalesInvoices** (Satış Faturaları) - Tam CRUD desteği + Ödeme + PDF
- ✅ **PurchaseBills** (Alış Faturaları) - Tam CRUD desteği + Ödeme + PDF
- ✅ **Employees** (Çalışanlar) - CRUD + Arşiv desteği
- ✅ **Salaries** (Maaşlar) - CRUD + Arşiv + Ödeme desteği
- ✅ **Taxes** (Vergiler) - CRUD + Arşiv + Ödeme desteği
- ✅ **Tags** (Etiketler) - Tam CRUD desteği
- ✅ **Warehouses** (Depolar) - Tam CRUD desteği
- ✅ **StockMovements** (Stok Hareketleri) - Listeleme desteği
- ✅ **StockUpdates** (Stok Güncellemeleri) - Oluşturma desteği
- ✅ **ItemCategories** (Ürün Kategorileri) - Tam CRUD desteği
- ✅ **SalesOffers** (Satış Teklifleri) - Tam CRUD + Arşiv + PDF + Durum Güncelleme
- ✅ **EArchives** (E-Arşiv) - Listeleme + Detay + PDF
- ✅ **EInvoiceInboxes** (E-Fatura Gelen Kutusu) - Listeleme desteği
- ✅ **EInvoices** (E-Faturalar) - CRUD + PDF desteği
- ✅ **ESMMs** (E-SMM) - CRUD + PDF desteği
- ✅ **Sharings** (Paylaşımlar) - Listeleme + Oluşturma desteği
- ✅ **ShipmentDocuments** (Sevkiyat Belgeleri) - Tam CRUD desteği
- ✅ **TrackableJobs** (İzlenebilir İşler) - Detay getirme desteği
- ✅ **Transactions** (İşlemler) - Detay + Güncelleme + Silme desteği
- ✅ **Webhooks** (Webhooklar) - Tam CRUD desteği

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