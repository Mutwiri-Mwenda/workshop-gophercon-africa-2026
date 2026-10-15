package layout

// Link is a labelled href, used for the nav and the footer.
type Link struct {
	Label string
	Href  string
}

// site holds the header and footer text shared by every page. Edit it here.
var site = struct {
	Name        string
	Tagline     string
	Description string
	Nav         []Link
	CTA         Link
	FooterCols  []struct {
		Title string
		Links []Link
	}
	GitHub  Link
	Socials []Link
	Credit  string
}{
	Name:        "Sightings",
	Tagline:     "Real wildlife, recorded by real people.",
	Description: "Real wildlife sightings from across the world, with photos, places and the weather on the day.",
	Nav: []Link{
		{Label: "Features", Href: "/#features"},
		{Label: "How it works", Href: "/#how"},
		{Label: "Sightings", Href: "/#sightings"},
		{Label: "FAQ", Href: "/#faq"},
	},
	CTA: Link{Label: "Explore sightings", Href: "/#sightings"},
	FooterCols: []struct {
		Title string
		Links []Link
	}{
		{Title: "Explore", Links: []Link{
			{Label: "Features", Href: "/#features"},
			{Label: "How it works", Href: "/#how"},
			{Label: "Sightings", Href: "/#sightings"},
			{Label: "Run it yourself", Href: "/#install"},
			{Label: "FAQ", Href: "/#faq"},
		}},
		{Title: "Data", Links: []Link{
			{Label: "GBIF", Href: "https://www.gbif.org"},
			{Label: "Open-Meteo", Href: "https://open-meteo.com"},
		}},
		{Title: "Built with", Links: []Link{
			{Label: "Go", Href: "https://go.dev"},
			{Label: "sqlc", Href: "https://sqlc.dev"},
			{Label: "templ", Href: "https://templ.guide"},
		}},
	},
	GitHub: Link{Label: "GitHub", Href: "https://github.com/Mutwiri-Mwenda/workshop-gophercon-africa-2026"},
	// Add your own, e.g. {Label: "X", Href: "https://x.com/you"}. Empty hrefs are skipped.
	Socials: []Link{
		{Label: "X", Href: ""},
		{Label: "LinkedIn", Href: ""},
	},
	Credit: "Built at GopherCon Africa 2026, Nairobi. Occurrence data from GBIF, weather from Open-Meteo.",
}
