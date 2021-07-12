import type {Children} from './lib/dom.js';
import {createHTML, clearElement} from './lib/dom.js';
import {a} from './lib/html.js';
import {people} from './gedcom.js';
import list from './list.js';
import fhcalc from './fhcalc.js';

declare const pageLoad: Promise<void>;

export const thisPage = window.location.pathname.split("/").pop()?.split(".").shift()!,
load = (module: string, params: Record<string, string | number>) => {
	let d: Children | undefined = undefined;
	switch (module) {
	case "fhcalc":
		d = fhcalc(params);
		break;
	case "list":
		d = list(params);
	}
	createHTML(clearElement(base), d || list({}));
},
link = (module: string, params: Record<string, string | number>) => a({"href": customPage ? `${module}.html?${params2String(params)}` : `?module=${module}&${params2String(params)}`, "onclick": (e: Event) => {
	e.preventDefault();
	load(module, params);
}}),
nameOf = (id: number) => `${people[id][0] ?? "??"} ${people[id][1] ?? "??"}`;

const customPage = ["list", "fhcalc", "tree"].includes(thisPage),
      params2String = (params: Record<string, string | number>) => Object.entries(params).map(([param, value]) => `${param}=${encodeURIComponent(value)}`).join("&");

let base: HTMLElement;

pageLoad.then(() => {
	base = document.getElementById("ged2web") || document.body;
	const params = Object.fromEntries(new URL(window.location + "").searchParams.entries());
	load(customPage ? thisPage : params["module"], params);

});
