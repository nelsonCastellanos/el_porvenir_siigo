package siigo_model

import (
	siigo_customer "el_porvenir.com/cloudfunction/siigo/model/siigo/customer"
	siigo_invoice "el_porvenir.com/cloudfunction/siigo/model/siigo/invoice"
	siigo_product "el_porvenir.com/cloudfunction/siigo/model/siigo/product"
)

type SiigoData struct {
	Taxes         []siigo_invoice.TaxesSiigo
	PaymentsTypes []siigo_invoice.PaymentTypesSiigo
	DocumentType  []siigo_invoice.DocumentTypesSiigo
	AccountGroup  []siigo_customer.AccountGroupSiigo
	Products      []siigo_product.ProductSiigo
	Customers     []siigo_customer.CustomerSiigo
	Invoices      []siigo_invoice.InvoiceSiigo
}
