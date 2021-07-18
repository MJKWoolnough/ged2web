import type {Children} from './lib/dom.js';
import {createHTML, clearElement} from './lib/dom.js';
import {a} from './lib/html.js';
import list from './list.js';
import fhcalc from './fhcalc.js';
import tree from './tree.js';

declare const pageLoad: Promise<void>;

export const thisPage = window.location.pathname.split("/").pop()?.split(".").shift()!,
load = (module: string, params: Record<string, string | number>, first = false) => {
	let d: Children | undefined = undefined;
	switch (module) {
	case "tree":
		d = tree(params);
		break;
	case "fhcalc":
		d = fhcalc(params);
		break;
	case "list":
		d = list(params);
	}
	if (!first) {
		history.pushState(null, "", modParams2URL(module, params));
	}
	createHTML(clearElement(base), d || list({}));
},
link = (module: string, params: Record<string, string | number>) => a({"href": modParams2URL(module, params), "onclick": (e: Event) => {
	e.preventDefault();
	load(module, params);
}}),
setTitle = (title: string) => document.title = `${baseTitle} - ${title}`;

const customPage = ["list", "fhcalc", "tree"].includes(thisPage),
      modParams2URL = (module: string, params: Record<string, string | number>) => (customPage ? `${module}.html?` : `?module=${module}&`) + Object.entries(params).map(([param, value]) => `${param}=${encodeURIComponent(value)}`).join("&"),
      loadPage = () => {
	const params = Object.fromEntries(new URL(window.location + "").searchParams.entries());
	load(customPage ? thisPage : params["module"], params, true);
      },
      baseTitle = document.title;

let base: HTMLElement;

window.addEventListener("popstate", loadPage);

pageLoad.then(() => {
	base = document.getElementById("ged2web") || document.body;
	loadPage();
});
