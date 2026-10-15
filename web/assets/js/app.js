// Sightings: the navbar's border, the install tabs, copy buttons, the
// mobile menu, and pausing looping previews while they're off screen.

(function () {
	// Navbar gains a border once the page has scrolled.
	const nav = document.querySelector('[data-nav]');
	if (nav) {
		const onScroll = function () {
			nav.classList.toggle('is-scrolled', window.scrollY > 8);
		};
		onScroll();
		window.addEventListener('scroll', onScroll, { passive: true });
	}

	// The mobile menu closes when one of its links is followed.
	document.querySelectorAll('.nav__menu').forEach(function (menu) {
		menu.addEventListener('click', function (e) {
			if (e.target.closest('a')) menu.removeAttribute('open');
		});
	});

	// Tabs, with arrow keys moving between them.
	document.querySelectorAll('[data-tabs]').forEach(function (root) {
		const tabs = Array.from(root.querySelectorAll('[role="tab"]'));
		const select = function (tab) {
			tabs.forEach(function (t) {
				const on = t === tab;
				t.setAttribute('aria-selected', String(on));
				t.tabIndex = on ? 0 : -1;
				document.getElementById(t.getAttribute('aria-controls')).hidden = !on;
			});
		};
		tabs.forEach(function (tab, i) {
			tab.addEventListener('click', function () {
				select(tab);
			});
			tab.addEventListener('keydown', function (e) {
				const step = { ArrowRight: 1, ArrowLeft: -1 }[e.key];
				if (!step) return;
				const next = tabs[(i + step + tabs.length) % tabs.length];
				select(next);
				next.focus();
			});
		});
	});

	// Copy buttons put their data-copy text on the clipboard.
	document.querySelectorAll('[data-copy]').forEach(function (btn) {
		const label = btn.textContent;
		btn.addEventListener('click', function () {
			if (!navigator.clipboard) return;
			navigator.clipboard.writeText(btn.dataset.copy).then(function () {
				btn.textContent = btn.dataset.copied;
				btn.classList.add('is-copied');
				setTimeout(function () {
					btn.textContent = label;
					btn.classList.remove('is-copied');
				}, 1800);
			});
		});
	});

	// Looping previews pause while off screen, to save the battery.
	if ('IntersectionObserver' in window) {
		const observer = new IntersectionObserver(function (entries) {
			entries.forEach(function (entry) {
				entry.target.classList.toggle('is-paused', !entry.isIntersecting);
			});
		});
		document.querySelectorAll('.demo, .strip, .step__visual, .tile__preview').forEach(function (el) {
			observer.observe(el);
		});
	}
})();
