package logoPorts

type LogoProvider interface {
	FetchCompanyLogoFromAPI(*string) ([]byte, error)
}
