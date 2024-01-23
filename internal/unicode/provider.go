package unicode_client

type UnicodeProviderClient struct {
	Username string
}

func NewUnicodeProviderClient(username string) *UnicodeProviderClient {
	return &UnicodeProviderClient{Username: username}
}

type UnicodeData struct {
	Char     string `json:"char" tfsdk:"unicode_char"`
	Category string `tfsdk:"unicode_category"`
	Block    string `tfsdk:"unicode_block"`
	Name     string `tfsdk:"unicode_name"`
}

func (u *UnicodeProviderClient) GetUnicodeCharData(unicodeChar string) (*UnicodeData, error) {
	return &UnicodeData{
		Char:     unicodeChar,
		Category: "Ll",
		Block:    "Basic Latin",
		Name:     "LATIN SMALL LETTER A",
	}, nil
}
