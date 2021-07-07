import list from './list.js';

declare const pageLoad: Promise<void>;

pageLoad.then(() => {
	const base = document.getElementById("ged2web") || document.body;
	switch (window.location.pathname.split("/").pop()!.split(".").shift()) {
	default:
		list(base);
	}

});
