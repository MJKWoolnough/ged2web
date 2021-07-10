import {a} from './lib/html.js';
import {people} from './gedcom.js';

export const thisPage = window.location.pathname.split("/").pop()?.split(".").shift()!,
link = (page: string, params: string, fn: (e: Event) => void) => a({"href": customPage ? `${page}.html?${params}` : `?page=${page}&${params}`, "onclick": (e: Event) => {
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
],
nameOf = (id: number) => `${people[id][0] ?? "??"} ${people[id][1] ?? "??"}`;

const customPage = ["list", "fhcalc", "tree"].includes(thisPage);
