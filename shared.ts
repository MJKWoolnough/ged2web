import {a} from './lib/html.js';

export const link = (page: string, params: string, fn: (e: Event) => void) => a({"href": `${page}.html?${params}`, "onclick": (e: Event) => {
	e.preventDefault();
	fn(e);
}}),
relations = [
	[
		"Parent",
		"Father",
		"Mother"
	],
	[
		"Sibling",
		"Brother",
		"Sister"
	],
	[
		"Spouse",
		"Husband",
		"Wife"
	],
	[
		"Child",
		"Son",
		"Daughter"
	],
	[
		"Pibling",
		"Uncle",
		"Aunt"
	],
	[
		"Nibling",
		"Nephew",
		"Neice"
	]
];
