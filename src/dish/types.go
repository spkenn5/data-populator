package dish

type (
	Dish struct {
		ID            int
		Name          string
		Description   string
		MenuAppeared  int
		TimesAppeared int
		FirstAppeared string
		LastAppeared  string
		LowPrice      float64
		HighPrice     float64
	}
)
