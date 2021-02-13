package providers

import (
	"log"
	"sort"
)

type Provider struct {
	ProviderIndex *ProviderIndex
}

type ProviderIndex struct {
	Providers    []string
	ProvidersMap map[string]string
}

func NewProvider() (*Provider, error) {
	initProviders()
	m := make(map[string]string)
	m["github"] = "Github"
	m["google"] = "Google"

	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	provider := &Provider{
		ProviderIndex: &ProviderIndex{Providers: keys, ProvidersMap: m},
	}
	return provider, nil
}


// valid the received access token
// return true if access token is valid otherwise return false
func (p *Provider) TokenValidation(accessToken string, provider string) bool {
	log.Println("Try to login with access token: ", accessToken, " from ", provider, "provider")
	return true
}


func initProviders()  {
	//goth.UseProviders(
	//	google.New("723796931510-uq7o9uroma092fok3cbkqtiglicm5tg9.apps.googleusercontent.com", "oD1sTf009H9kLS01n1h-d_iC", "http://localhost:4061/auth/google/callback", "email", "profile"),
	//	github.New("ce63dad6ba4f327bbd10", "a9cc172c085688a09656ae67a57afa6109b18cbd", "http://localhost:4061/auth/github/callback"),
	//)
}

