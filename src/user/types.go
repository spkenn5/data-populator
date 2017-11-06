package user

type (
	Name struct {
		Title string `json:"title"`
		First string `json:"first"`
		Last  string `json:"last"`
	}

	Location struct {
		Street string `json:"street"`
		City   string `json:"city"`
		State  string `json:"state"`
	}

	Login struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Salt     string `json:"salt"`
		MD5      string `json:"md5"`
	}

	Picture struct {
		Large     string `json:"large"`
		Medium    string `json:"medium"`
		Thumbnail string `json:"thumbnail"`
	}

	User struct {
		Gender      string   `json:"gender"`
		Name        Name     `json:"name"`
		Location    Location `json:"location"`
		Email       string   `json:"email"`
		Login       Login    `json:"login"`
		DOB         string   `json:"dob"`
		CreatedTime string   `json:"registered"`
		Phone       string   `json:"phone"`
		PictureURL  Picture  `json:"picture"`
		Nationality string   `json:"nat"`
	}

	Results struct {
		Results []User `json:"results"`
	}
)