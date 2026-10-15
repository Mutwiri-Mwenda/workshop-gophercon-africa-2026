package pages

import (
	"encoding/json"
	"fmt"
	"maps"
	"slices"
	"strings"
	"workshop/internal/domain/sighting"
	"workshop/web/views/components"

	"github.com/a-h/templ"
)

type (
	homeProps struct {
		Summary   sighting.Summary
		Hero      []sighting.Sighting
		Dots      []dot
		Spark     spark
		Top       []bar
		Weather   []sighting.Sighting
		Photo     *sighting.Sighting
		Observers []string
		Sample    *sighting.Sighting
		Raw       string
	}
	// dot is a sighting's position on a 360 by 180 equirectangular map.
	dot struct{ X, Y string }
	bar struct {
		Name  string
		Count int
		Width string
	}
	spark struct {
		Line, Area string
		From, To   int
	}
)

// newHomeProps works out everything the home page shows, once.
func newHomeProps(sightings []sighting.Sighting) homeProps {
	p := homeProps{Summary: sighting.Summarise(sightings)}

	withPhoto := slices.DeleteFunc(slices.Clone(sightings), func(s sighting.Sighting) bool {
		_, ok := components.Cover(s)
		return !ok
	})
	p.Hero = withPhoto[:min(5, len(withPhoto))]
	if len(p.Hero) == 0 {
		p.Hero = sightings[:min(5, len(sightings))]
	}
	if len(withPhoto) > 5 {
		p.Photo = &withPhoto[5]
	} else if len(withPhoto) > 0 {
		p.Photo = &withPhoto[0]
	}

	for _, s := range sightings {
		c := s.Location.Coordinates
		if c.Latitude == 0 && c.Longitude == 0 {
			continue
		}
		p.Dots = append(p.Dots, dot{
			X: fmt.Sprintf("%.1f", c.Longitude+180),
			Y: fmt.Sprintf("%.1f", 90-c.Latitude),
		})
	}

	p.Spark = newSpark(sighting.CountByYear(sightings))

	top := sighting.Top(sighting.CountBy(sightings, func(s sighting.Sighting) string {
		return components.DisplayName(s.Species)
	}), 5)
	for _, t := range top {
		p.Top = append(p.Top, bar{Name: t.Name, Count: t.Count, Width: fmt.Sprintf("%.1f", float64(t.Count)/float64(top[0].Count)*100)})
	}

	for _, s := range sightings {
		if !s.Weather.ObservedAt.IsZero() && s.Weather.Condition.Description != "" {
			p.Weather = append(p.Weather, s)
		}
		if len(p.Weather) == 3 {
			break
		}
	}

	seen := map[string]bool{}
	for _, s := range sightings {
		if s.RecordedBy != "" && !seen[s.RecordedBy] {
			seen[s.RecordedBy] = true
			p.Observers = append(p.Observers, s.RecordedBy)
		}
		if len(p.Observers) == 4 {
			break
		}
	}

	if i := slices.IndexFunc(withPhoto, func(s sighting.Sighting) bool {
		return s.RecordedBy != "" && s.Location.Locality != ""
	}); i >= 0 {
		p.Sample = &withPhoto[i]
	} else if len(withPhoto) > 0 {
		p.Sample = &withPhoto[0]
	}
	if p.Sample != nil {
		p.Raw = rawRecord(*p.Sample)
	}

	return p
}

// newSpark draws sightings per year as an SVG path on a 300 by 80 box.
func newSpark(counts map[int]int) spark {
	if len(counts) < 2 {
		return spark{}
	}
	years := slices.Sorted(maps.Keys(counts))
	from, to := years[0], years[len(years)-1]
	peak := slices.Max(slices.Collect(maps.Values(counts)))

	var b strings.Builder
	for y := from; y <= to; y++ {
		x := float64(y-from) / float64(to-from) * 300
		h := 76 - float64(counts[y])/float64(peak)*68
		cmd := "L"
		if y == from {
			cmd = "M"
		}
		fmt.Fprintf(&b, "%s%.1f %.1f ", cmd, x, h)
	}
	line := strings.TrimSpace(b.String())
	return spark{Line: line, Area: line + " L300 80 L0 80 Z", From: from, To: to}
}

// rawRecord renders a sighting back as the shape of a GBIF occurrence.
func rawRecord(s sighting.Sighting) string {
	type media struct {
		Type       string `json:"type"`
		Identifier string `json:"identifier"`
	}
	rec := struct {
		Key              int64   `json:"key"`
		ScientificName   string  `json:"scientificName"`
		EventDate        string  `json:"eventDate"`
		DecimalLatitude  float64 `json:"decimalLatitude"`
		DecimalLongitude float64 `json:"decimalLongitude"`
		CountryCode      string  `json:"countryCode"`
		Locality         string  `json:"locality,omitempty"`
		RecordedBy       string  `json:"recordedBy,omitempty"`
		BasisOfRecord    string  `json:"basisOfRecord,omitempty"`
		Media            []media `json:"media,omitempty"`
	}{
		Key:              s.GBIFKey,
		ScientificName:   s.Species.ScientificName,
		EventDate:        s.HappenedAt.Format("2006-01-02T15:04:05"),
		DecimalLatitude:  s.Location.Latitude,
		DecimalLongitude: s.Location.Longitude,
		CountryCode:      s.Location.CountryCode.String(),
		Locality:         s.Location.Locality,
		RecordedBy:       s.RecordedBy,
		BasisOfRecord:    s.BasisOfRecord,
	}
	if m, ok := components.Cover(s); ok {
		rec.Media = []media{{Type: m.Type, Identifier: m.URL}}
	}
	out, _ := json.MarshalIndent(rec, "", "  ")
	return string(out)
}

// temperature formats a reading as whole degrees Celsius.
func temperature(c float64) string {
	return fmt.Sprintf("%.0f°C", c)
}

// stepDelay staggers the how-it-works steps as they scroll in.
func stepDelay(i int) templ.SafeCSS {
	return templ.SafeCSS(fmt.Sprintf("--reveal-delay: %dms", i*120))
}

// loopClass picks the crossfade timing for a list of n items, up to five.
func loopClass(n int) string {
	if n < 2 {
		return ""
	}
	return fmt.Sprintf("loop--%d", min(n, 5))
}

// tabIndex keeps only the first tab in the page's tab order.
func tabIndex(i int) string {
	if i == 0 {
		return "0"
	}
	return "-1"
}

// initial returns the first letter of a name, for an avatar.
func initial(name string) string {
	for _, r := range name {
		return strings.ToUpper(string(r))
	}
	return "?"
}
