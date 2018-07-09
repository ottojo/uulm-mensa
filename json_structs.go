package main

type Mensaplan struct {
	Weeks []struct {
		WeekNumber int `json:"weekNumber"`
		Days []struct {
			Date         string `json:"date"`
			Mensa        Place  `json:"Mensa"`
			Bistro       Place  `json:"Bistro"`
			CB           Place  `json:"CB"`
			West         Place  `json:"West"`
			Prittwitzstr Place  `json:"Prittwitzstr"`
		} `json:"days"`
	} `json:"weeks"`
}

type Meal struct {
	Category string `json:"category"`
	Meal     string `json:"meal"`
	MealRaw  string `json:"meal_raw"`
	Price    string `json:"price"`
}

// TODO implement string function
func (m Meal) String() string {
	return m.Meal
}

type Place struct {
	Meals []Meal `json:"meals"`
	Open  bool   `json:"open"`
}

type Mensaplan_Static struct {
	Weeks []struct {
		Days []struct {
			Weekday   string `json:"date"`
			Burgerbar Place  `json:"Burgerbar"`
			Diner     Place  `json:"Diner"`
		} `json:"days"`
	} `json:"weeks"`
}
