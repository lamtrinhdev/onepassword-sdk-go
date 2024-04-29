// Code generated by typeshare 1.9.2. DO NOT EDIT.
package onepassword

type ItemCategory string

const (
	ItemCategoryLogin                ItemCategory = "Login"
	ItemCategorySecureNote           ItemCategory = "SecureNote"
	ItemCategoryCreditCard           ItemCategory = "CreditCard"
	ItemCategoryCryptoWallet         ItemCategory = "CryptoWallet"
	ItemCategoryIdentity             ItemCategory = "Identity"
	ItemCategoryPassword             ItemCategory = "Password"
	ItemCategoryDocument             ItemCategory = "Document"
	ItemCategoryApiCredentials       ItemCategory = "ApiCredentials"
	ItemCategoryBankAccount          ItemCategory = "BankAccount"
	ItemCategoryDatabase             ItemCategory = "Database"
	ItemCategoryDriverLicense        ItemCategory = "DriverLicense"
	ItemCategoryEmail                ItemCategory = "Email"
	ItemCategoryMedicalRecord        ItemCategory = "MedicalRecord"
	ItemCategoryMembership           ItemCategory = "Membership"
	ItemCategoryOutdoorLicense       ItemCategory = "OutdoorLicense"
	ItemCategoryPassport             ItemCategory = "Passport"
	ItemCategoryRewards              ItemCategory = "Rewards"
	ItemCategoryRouter               ItemCategory = "Router"
	ItemCategoryServer               ItemCategory = "Server"
	ItemCategorySshKey               ItemCategory = "SshKey"
	ItemCategorySocialSecurityNumber ItemCategory = "SocialSecurityNumber"
	ItemCategorySoftwareLicense      ItemCategory = "SoftwareLicense"
	ItemCategoryPerson               ItemCategory = "Person"
	ItemCategoryUnsupported          ItemCategory = "Unsupported"
)

type ItemFieldType string

const (
	ItemFieldTypeText        ItemFieldType = "Text"
	ItemFieldTypeConcealed   ItemFieldType = "Concealed"
	ItemFieldTypeUnsupported ItemFieldType = "Unsupported"
)

type ItemField struct {
	ID        string        `json:"id"`
	Title     string        `json:"title"`
	SectionID *string       `json:"section_id,omitempty"`
	FieldType ItemFieldType `json:"field_type"`
	Value     string        `json:"value"`
}
type ItemSection struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}
type Item struct {
	ID       string        `json:"id"`
	Title    string        `json:"title"`
	Category ItemCategory  `json:"category"`
	VaultID  string        `json:"vault_id"`
	Fields   []ItemField   `json:"fields"`
	Sections []ItemSection `json:"sections"`
}
