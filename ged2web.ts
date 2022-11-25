import {amendNode, bindElement, clearNode} from './lib/dom.js';
import {a, ns} from './lib/html.js';
import pageLoad from './lib/load.js';
import {goto, router} from './lib/router.js';
import fhcalc from './fhcalc.js';
import list from './list.js';
import tree from './tree.js';

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

export const load = (module: string, params: Record<string, string | number>) => goto(modParams2URL(module, params)),
link = (module: string, params: Record<string, string | number>) => a({"href": modParams2URL(module, params)}),
wrapper = bindElement(ns, "ged-2-web");

const customPage = ["list", "fhcalc", "tree"].includes(window.location.pathname.split("/").pop()?.split(".").shift()!),
      modParams2URL = (module: string, params: Record<string, string | number>) => (customPage ? `${module}.html?` : `?module=${module}&`) + Object.entries(params).map(([param, value]) => `${param}=${encodeURIComponent(value)}`).join("&"),
      baseTitle = document.title;

customElements.define("ged-2-web", Wrapper);

pageLoad.then(() => {
	clearNode(document.getElementById("ged2web") ?? document.body, router().add("tree.html?id=:id&highlight=:highlight", tree).add("?module=tree&id=:id&highlight=:highlight", tree).add("fhcalc.html?from=:from&to=:to", fhcalc).add("?module=fhcalc&from=:from&to=:to", fhcalc).add("list.html?l=:l&q=:q&p=:p", list).add("?module=list&l=:l&q=:q&p=:p", list).add("", list));
});
