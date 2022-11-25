import type {Children} from './lib/dom.js';
import {amendNode, clearNode} from './lib/dom.js';
import {a} from './lib/html.js';
import pageLoad from './lib/load.js';
import fhcalc from './fhcalc.js';
import list from './list.js';
import tree from './tree.js';

export const load = (module: string, params: Record<string, string | number>, first = false) => {
	let d: Children | undefined = undefined,
	    c = "list";
	switch (module) {
	case "tree":
		d = tree(params);
		c = module;
		break;
	case "fhcalc":
		d = fhcalc(params);
		c = module;
		break;
	case "list":
		d = list(params);
	}
	if (!first) {
		history.pushState(null, "", modParams2URL(module, params));
	}
	amendNode(document.body, {"class": {[lastClass]: false, [lastClass = "ged2web_" + c]: true}});
	clearNode(base, d || list({}));
},
link = (module: string, params: Record<string, string | number>) => a({"href": modParams2URL(module, params), "onclick": (e: Event) => {
	e.preventDefault();
	load(module, params);
}}),
setTitle = (title: string) => document.title = `${baseTitle} - ${title}`;

const basePage = () => window.location.pathname.split("/").pop()?.split(".").shift()!,
      customPage = ["list", "fhcalc", "tree"].includes(basePage()),
      modParams2URL = (module: string, params: Record<string, string | number>) => (customPage ? `${module}.html?` : `?module=${module}&`) + Object.entries(params).map(([param, value]) => `${param}=${encodeURIComponent(value)}`).join("&"),
      loadPage = () => {
	const params = Object.fromEntries(new URL(window.location + "").searchParams.entries());
	load(customPage ? basePage() : params["module"], params, true);
      },
      baseTitle = document.title;

let base: HTMLElement,
    lastClass = "_";

amendNode(window, {"onpopstate": loadPage});

pageLoad.then(() => {
	base = document.getElementById("ged2web") || document.body;
	loadPage();
});
