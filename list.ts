import type {ToString} from './global.js';
import {amendNode, clearNode} from './lib/dom.js';
import {button, datalist, div, h2, input, label, li, option, ul} from './lib/html.js';
import {checkInt} from './lib/misc.js';
import pagination from './lib/pagination.js';
import {families, people} from './gedcom.js';
import {link, load, modParams2URL, nameOf, relations, wrapper} from './global.js';

const indexes: number[][] = Array.from({length: 26}, () => []),
      stringSort = new Intl.Collator().compare,
      sortIDs = (a: number, b: number) => {
	const [paf = "", pas = ""] = people[a],
	      [pbf = "", pbs = ""] = people[b];

	return pas !== pbs ? stringSort(pas, pbs) : paf !== pbf ? stringSort(paf, pbf) : b - a;
      },
      person2HTML = (id: number, rel: number) => id === 0 ? [] : div([
	amendNode(link("tree", {id}), nameOf(id)),
	" (" + relations[rel][people[id][4]] + ")"
      ]),
      perPage = 20,
      paginationEnd = 3,
      paginationSurround = 3,
      buttons: [number, HTMLButtonElement][] = [],
      index2HTML = (base: HTMLDivElement, index: number[], params: Record<string, string>, page = 0) => {
	const max = Math.min((page + 1) * perPage, index.length),
	      list = ul({"class": "results"}),
	      pParams = {"href": (page: number) => modParams2URL("list", Object.assign(params, {"p": page})), "end": paginationEnd, "surround": paginationSurround, page, "total": Math.ceil(index.length / perPage) - 1};

	for (let i = page * perPage; i < max; i++) {
		const me = index[i],
		      [,,,,, childOf, ...spouseOf] = people[me],
		      [father, mother, ...siblings] = families[childOf],
		      spouseOfFams = spouseOf.map(s => families[s]),
		      c = button({"onclick": () => {
			if (chosen === me) {
				chosen = 0;
				for (const [, button] of buttons) {
					clearNode(button, "+");
				}
			} else if (chosen === 0) {
				chosen = me;
				for (const [id, button] of buttons) {
					clearNode(button, id === me ? "-" : "=");
				}
			} else {
				load("fhcalc", {from: chosen, to: me});
			}
		      }}, chosen === 0 ? "+" : chosen === me ? "-" : "=");

		buttons.push([me, c]);
		amendNode(list, li([
			div([
				amendNode(link("tree", {"id": me}), nameOf(me)),
				c
			]),
			div([
				person2HTML(father, 0),
				person2HTML(mother, 0),
				siblings.filter(id => id !== me).map(id => person2HTML(id, 1)),
				spouseOfFams.map(([husband, wife, ...children]) => [
					person2HTML(husband !== me ? husband : wife, 2),
					children.filter(id => id !== me).map(id => person2HTML(id, 3))
				])
			])
		]));
	}

	amendNode(base, [
		pagination(pParams),
		list,
		pagination(pParams)
	]);
      },
      searchCache = new Map<string, number[]>(),
      treeNames = datalist({"id": "treeNames"});

let chosen = 0;

for (let i = 0; i < people.length; i++) {
	let fl = (people[i][1] ?? "").charCodeAt(0);
	if (fl >= 97) {
		fl -= 32;
	}

	if (fl >= 65 && fl <=90) {
		indexes[fl - 65].push(i);
	}
}

for (const index of indexes) {
	index.sort(sortIDs);

	for (const id of index) {
		const [fname, lname] = people[id];

		if (fname && lname) {
			amendNode(treeNames, option({"value": `${fname} ${lname}`}));
		}
	}
}

export default (attrs: Record<string, ToString>) => {
	if (!treeNames.parentNode) {
		amendNode(document.body, treeNames);
	}

	const l = (attrs["l"] ?? "") + "",
	      q = (attrs["q"] ?? "") + "",
	      d = div(),
	      page = checkInt(parseInt(attrs["p"] + ""), 0),
	      search = () => load("list", {"q": s.value}),
	      s = input({"type": "text", "list": "treeNames", "onkeypress": (e: KeyboardEvent) => e.key === "Enter" && search()});

	let title = "List";

	if (q) {
		s.value = q;
		title = "Search";

		const terms = s.value.toUpperCase().split(" ").sort(),
		      jterms = terms.join(" ");

		let index: number[] = [];

		if (searchCache.has(jterms)) {
			index = searchCache.get(jterms)!;
		} else {
			for (let i = 0; i < people.length; i++) {
				const name = `${people[i][0] || ""} ${people[i][1] || ""}`.toUpperCase();

				if (terms.every(term => name.includes(term))) {
					index.push(i);
				}
			}

			searchCache.set(jterms, index);
		}

		index2HTML(d, index.sort(sortIDs), {q}, page)
	} else if (l) {
		const cc = l.toUpperCase().charCodeAt(0);

		if (cc >= 65 && cc <= 90) {
			title = `List - ${l}`;

			index2HTML(d, indexes[cc-65], {l}, page)
		}
	}
	return wrapper({title, "class": "ged2web_list"}, [
		div({"id": "ged2web_title"}, [
			h2("Select a Name"),
			div({"id": "indexes"}, indexes.map((_, id) => amendNode(link("list", {"l": String.fromCharCode(id+65)}), String.fromCharCode(id+65)))),
			div({"id": "index_search"}, [
				label({"for": "index_search"}, "Search Terms: "),
				s,
				button({"onclick": search}, "Search")
			])
		]),
		d
	]);
}
