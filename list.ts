import type {Children} from './lib/dom.js';
import {amendNode} from './lib/dom.js';
import {datalist, option, ul, li, div, span, h2, label, input, button} from './lib/html.js'
import {load, link, setTitle} from './ged2web.js';
import {people, families} from './gedcom.js';
import {relations} from './fhcalc.js';

export const nameOf = (id: number) => `${people[id][0] ?? "?"} ${people[id][1] ?? "?"}`;

const indexes: number[][] = Array.from({length: 26}, () => []),
      stringSort = new Intl.Collator().compare,
      sortIDs = (a: number, b: number) => {
	const [paf = "", pas = ""] = people[a],
	      [pbf = "", pbs = ""] = people[b];
	if (pas !== pbs) {
		return stringSort(pas, pbs);
	}
	if (paf !== pbf) {
		return stringSort(paf, pbf);
	}
	return b - a;
      },
      person2HTML = (id: number, rel: number) => id === 0 ? [] : div([
		amendNode(link("tree", {id}), nameOf(id)),
		" (" + relations[rel][people[id][4]] + ")"
      ]),
      perPage = 20,
      paginationEnd = 3,
      paginationSurround = 3,
      processPaginationSection = (ret: Children[], currPage: number, from: number, to: number, params: Record<string, string>)=> {
	if (ret.length !== 0) {
		ret.push("…");
	}
	for (let p = from; p <= to; p++) {
		if (p !== from) {
			ret.push(", ");
		}
		ret.push(currPage === p ? span((p+1)+"") : amendNode(link("list", Object.assign({p}, params)), {"class": "pagination_link"}, (p+1)+""));
	}
      },
      pagination = (index: number[], params: Record<string, string>, currPage = 0) => {
	const lastPage = Math.ceil(index.length / perPage) - 1,
	      ret: Children[] = [];
	if (lastPage === 0) {
		return [];
	}
	if (currPage > lastPage) {
		currPage = lastPage;
	}
	let start = 0;
	for (let page = 0; page <= lastPage; page++) {
		if (!(page < paginationEnd || page > lastPage-paginationEnd || ((paginationSurround > currPage || page >= currPage-paginationSurround) && page <= currPage+paginationSurround) || paginationEnd > 0 && ((currPage-paginationSurround-1 == paginationEnd && page == paginationEnd) || (currPage+paginationSurround+1 == lastPage-paginationEnd && page == lastPage-paginationEnd)))) {
			if (page != start) {
				processPaginationSection(ret, currPage, start, page - 1, params);
			}
			start = page + 1
		}
	}
	if (start < lastPage) {
		processPaginationSection(ret, currPage, start, lastPage, params);
	}
	return div({"class": "pagination"}, [
		"Pages: ",
		amendNode(currPage !== 0 ? link("list", Object.assign({"p": currPage-1}, params)) : span(), {"class": "pagination_link prev"}, "Previous"),
		ret,
		amendNode(currPage !== lastPage ? link("list", Object.assign({"p": currPage+1}, params)) : span(), {"class": "pagination_link next"}, "Next"),
	]);
      },
      index2HTML = (base: HTMLDivElement, index: number[], params: Record<string, string>, page = 0) => {
	const max = Math.min((page + 1) * perPage, index.length),
	      list = ul({"class": "results"}),
	      buttons: HTMLButtonElement[] = [];
	for (let i = page * perPage; i < max; i++) {
		const me = index[i],
		      [,,,,, childOf, ...spouseOf] = people[me],
		      [father, mother, ...siblings] = families[childOf],
		      spouseOfFams = spouseOf.map(s => families[s]),
		      c = button({"onclick": () => {
			if (chosen === me) {
				chosen = 0;
				for (const button of buttons) {
					button.innerText = "+";
				}
			} else if (chosen === 0) {
				chosen = me;
				for (const button of buttons) {
					button.innerText = "=";
				}
				c.innerText = "-";
			} else {
				load("fhcalc", {from: chosen, to: me});
			}
		      }}, chosen === 0 ? "+" : chosen === me ? "-" : "=");
		buttons.push(c);
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
					children.filter(id => id !== me).map(id => person2HTML(id, 3)),
				])
			])
		]));
	}
	amendNode(base, [
		pagination(index, params, page),
		list,
		pagination(index, params, page)
	]);
      },
      searchCache = new Map<string, number[]>(),
      search = () => load("list", {"q": s.value}),
      s = input({"type": "text", "list": "treeNames", "onkeypress": (e: KeyboardEvent) => e.key === "Enter" && search()}),
      treeNames = datalist({"id": "treeNames"});

let chosen = 0,
    head: HTMLDivElement;

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

export default ({l, q, p = 0}: Record<string, string | number>) => {
	const d = div(),
	      page = Math.max(0, typeof p === "string" ? parseInt(p) || 0 : p);
	setTitle("List");
	if (typeof q === "string") {
		s.value = q;
		setTitle("Search");
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
	} else if (typeof l === "string") {
		const cc = l.toUpperCase().charCodeAt(0);
		if (cc >= 65 && cc <= 90) {
			setTitle(`List - ${l}`);
			index2HTML(d, indexes[cc-65], {l}, page)
		}
	}
	return [
		head ? head : head = div({"id": "ged2web_title"}, [
			h2("Select a Name"),
			div({"id": "indexes"}, indexes.map((_, id) => amendNode(link("list", {"l": String.fromCharCode(id+65)}), String.fromCharCode(id+65)))),
			div({"id": "index_search"}, [
				treeNames,
				label({"for": "index_search"}, "Search Terms: "),
				s,
				button({"onclick": search}, "Search"),
			]),
		]),
		d
	];
}
