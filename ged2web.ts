import {thisPage} from './shared.js';
import list from './list.js';

declare const pageLoad: Promise<void>;

pageLoad.then(() => {
	const base = document.getElementById("ged2web") || document.body;
	switch (thisPage) {
	default:
		list(base);
	}

});
