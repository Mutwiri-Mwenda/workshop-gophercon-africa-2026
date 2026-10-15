package pages

type (
	step struct {
		Title string
		Body  string
	}
	feature struct {
		Title string
		Body  string
	}
	installTab struct {
		ID    string
		Label string
		Code  string
	}
	faq struct {
		Q string
		A string
	}
)

// homeCopy holds every line of text on the home page. Edit it here.
var homeCopy = struct {
	Eyebrow      string
	Headline     string
	HeadlineHi   string
	Subtext      string // takes the sighting and country counts
	PrimaryCTA   string
	SecondaryCTA string
	WallLabel    string

	StripLabel string
	Strip      []string

	HowEyebrow string
	HowTitle   string
	Steps      [3]step

	FeaturesEyebrow string
	FeaturesTitle   string
	Map             feature
	Weather         feature
	Top             feature
	Years           feature
	Photos          feature
	Observers       feature

	AfricaEyebrow  string
	AfricaTitle    string
	AfricaStat     string // takes the sighting and country counts
	AfricaLive     string
	AfricaLabel    string
	AfricaFallback string
	AfricaEmpty    string

	CompareEyebrow string
	CompareTitle   string
	BeforeLabel    string
	AfterLabel     string

	SightingsEyebrow string
	SightingsTitle   string
	SightingsEmpty   string

	InstallEyebrow string
	InstallTitle   string
	InstallBody    string
	InstallTabs    []installTab
	CopyLabel      string
	CopiedLabel    string

	FAQEyebrow string
	FAQTitle   string
	FAQs       []faq

	CloserTitle string
	CloserBody  string
	CloserCTA   string
}{
	Eyebrow:      "Live from GBIF",
	Headline:     "See the wild,",
	HeadlineHi:   "as it's seen.",
	Subtext:      "%d real sightings from %d countries. Photos, places and the weather on the day.",
	PrimaryCTA:   "Explore sightings",
	SecondaryCTA: "See it in action",
	WallLabel:    "Animals recently seen",

	StripLabel: "Powered by",
	Strip:      []string{"GBIF", "Open-Meteo", "SQLite", "sqlc", "templ", "Go"},

	HowEyebrow: "How it works",
	HowTitle:   "From the field to your screen.",
	Steps: [3]step{
		{Title: "Seen", Body: "Someone in the field spots an animal and logs it with GBIF."},
		{Title: "Stored", Body: "We save the record and the weather from that day."},
		{Title: "Shown", Body: "The page is built on the server and arrives ready to read."},
	},

	FeaturesEyebrow: "Features",
	FeaturesTitle:   "Every record, brought to life.",
	Map:             feature{Title: "Every sighting, mapped", Body: "Each dot is a real animal, seen at that exact spot."},
	Weather:         feature{Title: "Weather on the day", Body: "The temperature, sky and wind at the time of each sighting."},
	Top:             feature{Title: "Most seen", Body: "See at a glance which animals get spotted most."},
	Years:           feature{Title: "Decades of records", Body: "Sightings that go back to %d."},
	Photos:          feature{Title: "Real photos, credited", Body: "Taken in the field, and credited to whoever took them."},
	Observers:       feature{Title: "%d observers", Body: "Every record has a real person behind it."},

	AfricaEyebrow:  "Live map",
	AfricaTitle:    "Seen across Africa.",
	AfricaStat:     "%d sightings · %d countries",
	AfricaLive:     "Spotted",
	AfricaLabel:    "Map of sightings across Africa",
	AfricaFallback: "The map needs JavaScript and a connection to load.",
	AfricaEmpty:    "No sightings in Africa yet.",

	CompareEyebrow: "Before and after",
	CompareTitle:   "Raw data in. A story out.",
	BeforeLabel:    "The raw GBIF record",
	AfterLabel:     "What you see here",

	SightingsEyebrow: "Sightings",
	SightingsTitle:   "The latest from the wild.",
	SightingsEmpty:   "No sightings yet. Run go run . db seed and refresh.",

	InstallEyebrow: "Run it yourself",
	InstallTitle:   "Your own copy in two minutes.",
	InstallBody:    "No API keys and no internet needed. It comes with its own copy of the data.",
	InstallTabs: []installTab{
		{ID: "macos", Label: "macOS", Code: "brew install go\ngit clone https://github.com/Mutwiri-Mwenda/workshop-gophercon-africa-2026.git\ncd workshop-gophercon-africa-2026\nmake setup\ngo run . db seed\nmake web"},
		{ID: "linux", Label: "Linux", Code: "sudo snap install go --classic\nsudo apt install git make\ngit clone https://github.com/Mutwiri-Mwenda/workshop-gophercon-africa-2026.git\ncd workshop-gophercon-africa-2026\nmake setup\ngo run . db seed\nmake web"},
		{ID: "windows", Label: "Windows", Code: "winget install GoLang.Go Git.Git ezwinports.make\ngit clone https://github.com/Mutwiri-Mwenda/workshop-gophercon-africa-2026.git\ncd workshop-gophercon-africa-2026\nmake setup\ngo run . db seed\nmake web"},
	},
	CopyLabel:   "Copy",
	CopiedLabel: "Copied",

	FAQEyebrow: "FAQ",
	FAQTitle:   "Good questions.",
	FAQs: []faq{
		{Q: "Are these sightings real?", A: "Yes. Every one is a real record from GBIF, the Global Biodiversity Information Facility, logged by a person or a museum."},
		{Q: "Where does the weather come from?", A: "From Open-Meteo's historical records, for the place and day of each sighting."},
		{Q: "Why do some sightings have no photo?", A: "Not every record comes with a photo. Those show a gradient instead, and their details are still there."},
		{Q: "Can I run my own copy?", A: "Yes. Clone the repository, run make setup, and you'll have it on localhost:8080."},
		{Q: "What is it built with?", A: "Go, with SQLite through sqlc for the data and templ for the pages. There's no JavaScript framework."},
	},

	CloserTitle: "The wild doesn't wait.",
	CloserBody:  "Hundreds of animals, seen this week and decades ago. Go and look.",
	CloserCTA:   "Explore sightings",
}
