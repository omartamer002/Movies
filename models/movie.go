package models

type Movie struct {
	ID       string    `json:"id"`
	Title    string    `json:"title"`
	ISBN     string    `json:"isbn"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var Movies = []Movie{
	{
		ID:    "1",
		Title: "Catch me if you can",
		ISBN:  "978-3-16-148410-0",
		Director: &Director{
			FirstName: "Steven",
			LastName:  "Spielberg",
		},
	},
	{
		ID:    "2",
		Title: "Suicide Squad",
		ISBN:  "978-3-16-148410-0",
		Director: &Director{
			FirstName: "David",
			LastName:  "Ayre",
		},
	},
}
