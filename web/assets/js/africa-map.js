// The live map of Africa, adapted from the map block. Every sighting is a
// dim dot; a pulse hops from one to the next, lighting each up, while the
// ticker names what was seen there. With reduced motion every dot is lit
// and nothing moves.

(function () {
	const el = document.querySelector('[data-map="africa-points"]');
	if (!el) return;

	const TILES = 'https://tile.openstreetmap.org/{z}/{x}/{y}.png';
	const CREDIT = '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors';
	const AFRICA = [[-35.5, -18.5], [37.8, 52]];
	const STEP = 2600;
	const still = window.matchMedia('(prefers-reduced-motion: reduce)').matches;

	let data = [];
	try {
		data = JSON.parse(document.getElementById('africa-points').textContent) || [];
	} catch (err) {
		data = [];
	}
	if (!data.length || typeof L === 'undefined') {
		el.classList.add('is-unavailable');
		return;
	}

	const css = getComputedStyle(el);
	const accent = css.getPropertyValue('--map-marker').trim() || '#c8ff3d';
	const dim = { radius: 4, weight: 0, fillColor: accent, fillOpacity: 0.3 };
	const lit = { radius: 5, weight: 0, fillColor: accent, fillOpacity: 0.95 };

	// Built from nodes, never an HTML string: every field here is GBIF's text.
	function popup(p) {
		const root = document.createElement('div');
		root.className = 'map__popup';
		root.appendChild(document.createElement('strong')).textContent = p.name;
		root.appendChild(document.createElement('span')).textContent = [p.place, p.date].filter(Boolean).join(' · ');
		if (p.by) root.appendChild(document.createElement('span')).textContent = 'Recorded by ' + p.by;
		return root;
	}

	const map = L.map(el, {
		preferCanvas: true,
		scrollWheelZoom: false,
		zoomAnimation: !still,
		fadeAnimation: !still,
		markerZoomAnimation: !still,
	});
	L.tileLayer(TILES, { attribution: CREDIT, maxZoom: 18, minZoom: 2 }).addTo(map);
	map.fitBounds(AFRICA, { animate: false });
	map.on('focus', () => map.scrollWheelZoom.enable());
	map.on('blur', () => map.scrollWheelZoom.disable());

	const dots = data.map((p) =>
		L.circleMarker([p.lat, p.lng], still ? lit : dim)
			.bindPopup(() => popup(p))
			.addTo(map),
	);

	const ticker = document.querySelector('[data-ticker]');
	const name = ticker.querySelector('[data-ticker-name]');
	const meta = ticker.querySelector('[data-ticker-meta]');
	const by = ticker.querySelector('[data-ticker-by]');

	function tell(p) {
		name.textContent = p.name;
		meta.textContent = [p.place, p.date].filter(Boolean).join(' · ');
		by.textContent = p.by ? 'Recorded by ' + p.by : '';
		ticker.hidden = false;
	}

	if (still) {
		tell(data[0]);
		return;
	}

	const pulse = L.marker([data[0].lat, data[0].lng], {
		icon: L.divIcon({ className: 'atlas__pulse', html: '<span></span><span></span>', iconSize: [18, 18] }),
		interactive: false,
		keyboard: false,
	}).addTo(map);

	// A shuffled order, so the pulse wanders the continent rather than a list.
	const order = data.map((_, i) => i);
	for (let i = order.length - 1; i > 0; i--) {
		const j = Math.floor(Math.random() * (i + 1));
		[order[i], order[j]] = [order[j], order[i]];
	}

	let n = 0;
	function next() {
		if (n === order.length) {
			dots.forEach((d) => d.setStyle(dim));
			n = 0;
		}
		const i = order[n++];
		dots[i].setStyle(lit);
		pulse.setLatLng([data[i].lat, data[i].lng]);
		tell(data[i]);
		ticker.classList.remove('is-new');
		void ticker.offsetWidth;
		ticker.classList.add('is-new');
	}

	// Runs only while the map is on screen and the tab is in front.
	let timer = null;
	let onScreen = false;
	function run() {
		const go = onScreen && !document.hidden;
		if (go && !timer) {
			next();
			timer = setInterval(next, STEP);
		} else if (!go && timer) {
			clearInterval(timer);
			timer = null;
		}
	}
	new IntersectionObserver((entries) => {
		onScreen = entries[0].isIntersecting;
		run();
	}).observe(el);
	document.addEventListener('visibilitychange', run);
})();
