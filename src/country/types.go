package country

type (
	Currency struct {
		Code   string `json:"code"`
		Name   string `json:"name"`
		Symbol string `json:"symbol"`
	}

	Language struct {
		Name       string `json:"name"`
		NativeName string `json:"nativeName"`
	}

	Translation struct {
		DE string `json:"de"`
		ES string `json:"es"`
		FR string `json:"fr"`
		JA string `json:"ja"`
		BR string `json:"br"`
		PT string `json:"pt"`
		NL string `json:"nl"`
		HR string `json:"hr"`
		FA string `json:"fa"`
	}

	Country struct {
		Name           string      `json:"name"`
		Alpha2Code     string      `json:alpha2Code"`
		Alpha3Code     string      `json:"alpha3Code"`
		Capital        string      `json:"capital"`
		TopLevelDomain []string    `json:"topLevelDomain"`
		CallingCodes   []string    `json:"callingCodes"`
		Region         string      `json:"region"`
		LatLong        []float64   `json:"latlng"`
		NativeName     string      `json:"nativeName"`
		Currencies     []Currency  `json:"currencies"`
		Languages      []Language  `json:"languages"`
		Translations   Translation `json:"translations"`
		Flag           string      `json:"flag"`
	}
)