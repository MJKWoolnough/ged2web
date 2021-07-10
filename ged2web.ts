import {createHTML, clearElement} from './lib/dom.js';
import {a} from './lib/html.js';
import {people} from './gedcom.js';
import list from './list.js';

declare const pageLoad: Promise<void>;

export const thisPage = window.location.pathname.split("/").pop()?.split(".").shift()!,
load = (module: string, params: Record<string, string | number>) => {
	switch (module) {
	case "list":
		createHTML(clearElement(base), list(params));
	}
},
link = (module: string, params: Record<string, string | number>) => a({"href": customPage ? `${module}.html?${params2String(params)}` : `?module=${module}&${params2String(params)}`, "onclick": (e: Event) => {
	e.preventDefault();
	load(module, params);
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

const customPage = ["list", "fhcalc", "tree"].includes(thisPage),
      params2String = (params: Record<string, string | number>) => Object.entries(params).map(([param, value]) => `${param}=${encodeURIComponent(value)}`).join("&");

let base: HTMLElement;

pageLoad.then(() => {
	base = document.getElementById("ged2web") || document.body;
	switch (thisPage) {
	default:
		load("list", {});
	}

});
