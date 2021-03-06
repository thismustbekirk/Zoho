package crm

import (
	"encoding/xml"
	"time"
)

// AccountsModule exposes the API methods that are related to the Account type
type AccountsModule CrmModule

// Accounts returns an pointer to an 'AccountsModule' which can be used to call the methods
// that can be used to interact with Zoho CRM 'Account' types
func (a *API) Accounts() *AccountsModule {
	return &AccountsModule{
		id:     a.id,
		api:    a,
		module: accountsModule,
	}
}

func (A *AccountsModule) GetMyRecords(o GetRecordsOptions) (Accounts, error) {
	A.options = o
	v, err := A.api.getMyRecords(A.module, o)
	if v == nil {
		return Accounts{}, err
	}
	return v.(Accounts), err
}

func (A *AccountsModule) GetRecords(o GetRecordsOptions) (Accounts, error) {
	A.options = o
	v, err := A.api.getRecords(A.module, o)
	if v == nil {
		return Accounts{}, err
	}
	return v.(Accounts), err
}

func (A *AccountsModule) GetRecordsByID(o GetRecordsByIdOptions) (Accounts, error) {
	A.options = o
	v, err := A.api.getRecordById(A.module, o)
	if v == nil {
		return Accounts{}, err
	}
	return v.(Accounts), err
}

func (A *AccountsModule) GetDeletedRecordIds(o GetDeletedRecordIdsOptions) (Accounts, error) {
	A.options = o
	v, err := A.api.getDeletedRecordIds(A.module, o)
	if v == nil {
		return Accounts{}, err
	}
	return v.(Accounts), err
}

func (A *AccountsModule) InsertRecords(o InsertRecordsOptions) (Accounts, error) {
	A.options = o
	v, err := A.api.insertRecords(A.module, o)
	if v == nil {
		return Accounts{}, err
	}
	return v.(Accounts), err
}

func (A *AccountsModule) UpdateRecords(o UpdateRecordsOptions) (Accounts, error) {
	A.options = o
	v, err := A.api.updateRecords(A.module, o)
	if v == nil {
		return Accounts{}, err
	}
	return v.(Accounts), err
}

func (A *AccountsModule) DeleteRecords(id string) (Accounts, error) {
	v, err := A.api.deleteRecords(A.module, id)
	if v == nil {
		return Accounts{}, err
	}
	return v.(Accounts), err
}

func (A *AccountsModule) GetSearchRecordsByPDC(o GetSearchRecordsByPDCOptions) (Accounts, error) {
	A.options = o
	v, err := A.api.getSearchRecordsByPDC(A.module, o)
	if v == nil {
		return Accounts{}, err
	}
	return v.(Accounts), err
}

func (A *AccountsModule) GetRelatedRecords(o GetRelatedRecordsOptions) (Accounts, error) {
	A.options = o
	v, err := A.api.getRelatedRecords(A.module, o)
	if v == nil {
		return Accounts{}, err
	}
	return v.(Accounts), err
}

func (A *AccountsModule) UpdateRelatedRecord(o UpdateRelatedRecordOptions) (Accounts, error) {
	A.options = o
	v, err := A.api.updateRelatedRecord(A.module, o)
	if v == nil {
		return Accounts{}, err
	}
	return v.(Accounts), err
}

func (A *AccountsModule) GetFields(kind int) (Accounts, error) {
	v, err := A.api.getFields(A.module, kind)
	if v == nil {
		return Accounts{}, err
	}
	return v.(Accounts), err
}

func (A *AccountsModule) UploadFile(o UploadFileOptions) (Accounts, error) {
	A.options = o
	v, err := A.api.uploadFile(A.module, o)
	if v == nil {
		return Accounts{}, err
	}
	return v.(Accounts), err
}

func (A *AccountsModule) DownloadFile(id string) (Accounts, error) {
	v, err := A.api.downloadFile(A.module, id)
	if v == nil {
		return Accounts{}, err
	}
	return v.(Accounts), err
}

func (A *AccountsModule) DeleteFile(id string) (Accounts, error) {
	v, err := A.api.deleteFile(A.module, id)
	if v == nil {
		return Accounts{}, err
	}
	return v.(Accounts), err
}

func (A *AccountsModule) Delink(o DelinkOptions) (Accounts, error) {
	A.options = o
	v, err := A.api.delink(A.module, o)
	if v == nil {
		return Accounts{}, err
	}
	return v.(Accounts), err
}

func (A *AccountsModule) UploadPhoto(o UploadPhotoOptions) (Accounts, error) {
	A.options = o
	v, err := A.api.uploadPhoto(A.module, o)
	if v == nil {
		return Accounts{}, err
	}
	return v.(Accounts), err
}

func (A *AccountsModule) DownloadPhoto(id string) (Accounts, error) {
	v, err := A.api.downloadPhoto(A.module, id)
	if v == nil {
		return Accounts{}, err
	}
	return v.(Accounts), err
}

func (A *AccountsModule) DeletePhoto(id string) (Accounts, error) {
	v, err := A.api.deletePhoto(A.module, id)
	if v == nil {
		return Accounts{}, err
	}
	return v.(Accounts), err
}

func (A *AccountsModule) SearchRecords(o SearchRecordsOptions) (Accounts, error) {
	A.options = o
	v, err := A.api.searchRecords(A.module, o)
	if v == nil {
		return Accounts{}, err
	}
	return v.(Accounts), err
}

// https://www.zoho.com/crm/help/api/modules-fields.html#Accounts
type Account struct {
	AccountName     string  `zoho:"Account Name"`      // 100 chars
	AccountOwner    string  `zoho:"Account Owner"`     //
	Website         string  `zoho:"Website"`           // 120 chars
	TickerSymbol    string  `zoho:"Ticker Symbol"`     // 30 chars
	ParentAccount   string  `zoho:"Parent Account"`    //
	Employees       int     `zoho:"Employees"`         // 10 digits
	Ownership       string  `zoho:"Ownership"`         //
	Industry        string  `zoho:"Industry"`          //
	AccountType     string  `zoho:"Account Type"`      //
	AccountNumber   int     `zoho:"Account Number"`    // 40 digits
	AccountSite     string  `zoho:"Account Site"`      // 30 chars
	Phone           string  `zoho:"Phone"`             // 30 chars
	Fax             string  `zoho:"Fax"`               // 30 chars
	Email           string  `zoho:"Email"`             // 100 chars
	Rating          string  `zoho:"Rating"`            //
	SicCode         int     `zoho:"SIC Code"`          // 10 digits
	AnnualRevenue   float64 `zoho:"Annual Revenue"`    // 16 digits
	BillingStreet   string  `zoho:"Billing Street"`    // 250 chars
	BillingCity     string  `zoho:"Billing City"`      // 30 chars
	BillingState    string  `zoho:"Billing State"`     // 30 chars
	BillingZipCode  string  `zoho:"Billing Zip Code"`  // 30 chars
	BillingCountry  string  `zoho:"Billing Country"`   // 30 chars
	ShippingStreet  string  `zoho:"Shipping Street"`   // 250 chars
	ShippingCity    string  `zoho:"Shipping City"`     // 30 chars
	ShippingState   string  `zoho:"Shipping State"`    // 30 chars
	ShippingZipCode string  `zoho:"Shipping Zip Code"` // 30 chars
	ShippingCountry string  `zoho:"Shipping Country"`  // 30 chars
	Description     string  `zoho:"Description"`       // 32000 chars
	Currency        string  `zoho:"Currency"`          //
	ExchangedRate   float64 `zoho:"Exchange Rate"`     //
	BillingCode     string  `zoho:"Billing Code"`      //  GetRecords
	Territories     string  `zoho:"Territories"`       //  GetRecords

	//Fields not documented but returned by some methods
	//Some of these fields may need to be stripped on 'insert' due to zoho requirements
	AccountID        int       `zoho:"ACCOUNTID,strip"`
	SmOwnerID        int       `zoho:"SMOWNERID,strip"`
	SmCreatorID      int       `zoho:"SMCREATORID,strip"`
	CreatedBy        string    `zoho:"Created By,strip"`
	ModifiedByID     int       `zoho:"MODIFIEDBY,strip"`
	ModifiedBy       string    `zoho:"Modified By,strip"`
	CreatedTime      time.Time `zoho:"Created Time,strip"`
	ModifiedTime     time.Time `zoho:"Modified Time,strip"`
	LastActivityTime time.Time `zoho:"Last Activity Time,strip"`

	ExtraFields ExtraFields
}

func (a Account) writeXML() string {
	x := XMLData{XMLName: xml.Name{Local: "Accounts"}}
	x.addRow(a, 1)
	return x.encode()
}

// String will return the XML Encoded value that will be passed to zoho
func (a Account) String() string {
	return a.writeXML()
}

// Accounts is a slice of Account
type Accounts []Account

func (A Accounts) writeXML() string {
	x := XMLData{XMLName: xml.Name{Local: "Accounts"}}
	for i := 0; i < len(A); i++ {
		x.addRow(A[i], i+1)
	}
	return x.encode()
}

// String will return the XML Encoded value that will be passed to zoho
func (A Accounts) String() string {
	return A.writeXML()
}
