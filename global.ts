import {amendNode, bindCustomElement} from './lib/dom.js';
import {a} from './lib/html.js';
import {goto} from './lib/router.js';
import {people} from './gedcom.js';

export interface ToString {
	toString(): string;
}

class Wrapper extends HTMLElement {
	#class = "";
	connectedCallback() {
		amendNode(document.body, {"class": [this.#class ??= this.getAttribute("class") ?? "_"]});
		document.title = `${baseTitle} - ${this.getAttribute("title") ?? ""}`;
	}
	disconnectedCallback() {
		amendNode(document.body, {"class": {[this.#class]: false}});
	}
}

const customPage = ["list", "fhcalc", "tree"].includes(window.location.pathname.split("/").pop()?.split(".").shift() ?? ""),
      baseTitle = document.title;

export const load = (module: string, params: Record<string, string | number>) => goto(modParams2URL(module, params)),
modParams2URL = (module: string, params: Record<string, string | number>) => (customPage ? `${module}.html?` : `?module=${module}&`) + Object.entries(params).map(([param, value]) => `${param}=${encodeURIComponent(value)}`).join("&"),
link = (module: string, params: Record<string, string | number>) => a({"href": modParams2URL(module, params)}),
wrapper = bindCustomElement("ged-2-web", Wrapper),
nameOf = (id: number) => `${people[id][0] ?? "?"} ${people[id][1] ?? "?"}`,
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
		"Niece"
	]
];
