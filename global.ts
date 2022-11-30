import {amendNode, bindElement} from './lib/dom.js';
import {a, ns} from './lib/html.js';
import {goto} from './lib/router.js';
import {people} from './gedcom.js';

export interface ToString {
	toString(): string;
}

class Wrapper extends HTMLElement {
	connectedCallback() {
		amendNode(document.body, {"class": [this.getAttribute("class") ?? "_"]});
		document.title = `${baseTitle} - ${this.getAttribute("title") ?? ""}`;
	}
	disconnectedCallback() {
		amendNode(document.body, {"class": {[this.getAttribute("class") ?? "_"]: false}});
	}
}

customElements.define("ged-2-web", Wrapper);

const customPage = ["list", "fhcalc", "tree"].includes(window.location.pathname.split("/").pop()?.split(".").shift()!),
      modParams2URL = (module: string, params: Record<string, string | number>) => (customPage ? `${module}.html?` : `?module=${module}&`) + Object.entries(params).map(([param, value]) => `${param}=${encodeURIComponent(value)}`).join("&"),
      baseTitle = document.title;

export const load = (module: string, params: Record<string, string | number>) => goto(modParams2URL(module, params)),
link = (module: string, params: Record<string, string | number>) => a({"href": modParams2URL(module, params)}),
wrapper = bindElement(ns, "ged-2-web"),
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
