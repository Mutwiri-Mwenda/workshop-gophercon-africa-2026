# Design

A premium, cinematic dark landing page for the sightings, after magnific.com: near-black, huge
white headlines, one electric accent, big rounded cards and previews that loop.

## Colours (`web/assets/css/app.css`, `:root`)

| Property | Value | Use |
|----------|-------|-----|
| `--bg` | `#07080A` | page background, `theme-color` |
| `--surface` | `#111317` | cards and tiles |
| `--surface-2` | `#181B20` | insets inside cards |
| `--border` / `--border-strong` | white at 8% / 16% | card edges |
| `--text` | `#F5F7FA` | lettering |
| `--muted` | `#A1A8B3` | secondary lettering |
| `--accent` | `#C8FF3D` electric lime | buttons, highlights, glows: the only bright colour |
| `--on-accent` | `#0B0D00` | lettering on lime |

## Lettering

- Headlines: Inter Tight 800, tight tracking (`-0.035em`), `line-height: 1.02`. `--font-display`.
- Body: Inter 400–600. `--font-body`.
- Small labels and code: JetBrains Mono, uppercase with wide tracking. `--font-mono`.

## Shape and space

- Radii: `--radius-lg` 28px (cards, tiles), `--radius-md` 18px, `--radius-sm` 12px; buttons are pills.
- Roomy: sections `--section` (80–140px) apart, content `--wrap` 1200px wide.

## Components

- **Logo**: a dark eye on a lime rounded square, `assets/images/logo.svg`, also the favicon.
- **Navbar**: sticky, clear over the hero, blurred dark once scrolled.
- **Hero**: full-bleed, the latest sightings' photos crossfading and slowly zooming behind a
  centred headline; a glass card bottom-left names the animal, place, date and observer.
- **Tiles and cards**: `--surface`, soft border; on hover a lime edge and a soft lime glow.
- **Previews**: CSS loops only, paused off screen and frozen on the first frame for reduced motion.
- **Sightings**: a photo wall of cards, 4:3 photos, with gradient placeholders when there's none.
- Sections fade up as they scroll in (`reveal.css`/`reveal.js`), never the individual cards.

## Avoid

- A second bright colour. No pricing section: the site is free.
