import type {Children} from './lib/dom.js';
import {createHTML, clearElement} from './lib/dom.js';
import {ul, li, div, span, h2, label, input, button} from './lib/html.js'
import {people, families} from './gedcom.js';

const indexes: number[][] = Array.from({length: 26}, () => []),
      stringSort = new Intl.Collator().compare,
      sortIDs = (a: number, b: number) => {
	const pa = people[a],
	      pb = people[b];
	if (pa[1] !== pb[1]) {
		return stringSort(pa[1], pb[1]);
	}
	if (pa[0] !== pb[0]) {
		return stringSort(pa[0], pb[0]);
	}
	return b - a;
      },
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
      person2HTML = (id: number, rel: number) => {
	if (id === 0) {
		return [];
	}
	const person = people[id];
	return div([
		span((person[0] || "??") + " " + (person[1] || "??")),
		" (" + relations[rel][person[4]] + ")"
	])
      },
      perPage = 20,
      paginationEnd = 3,
      paginationSurround = 3,
      processPaginationSection = (pageFn: (page: number) => void, ret: Children[], currPage: number, from: number, to: number) => {
	if (ret.length !== 0) {
		ret.push("…");
	}
	for (let i = from; i <= to; i++) {
		if (i !== from) {
			ret.push(", ");
		}
		ret.push(span(currPage !== i ? {"class": "pagination_link", "onclick": pageFn.bind(null, i)} : {}, (i+1)+""));
	}
      },
      pagination = (base: HTMLDivElement, index: number[], currPage = 0) => {
	const lastPage = Math.ceil(index.length / perPage) - 1,
	      ret: Children[] = [],
	      pageFn = (page: number) => index2HTML(base, index, page);
	if (currPage > lastPage) {
		currPage = lastPage;
	}
	let start = 0;
	for (let page = 0; page <= lastPage; page++) {
		if (!(page < paginationEnd || // Beginning
			page > lastPage-paginationEnd || // End
			((paginationSurround > currPage || page >= currPage-paginationSurround) && page <= currPage+paginationSurround) || // Middle
			paginationEnd > 0 && ((currPage-paginationSurround-1 == paginationEnd && page == paginationEnd) || // Merge Begining and Middle if close enough
			(currPage+paginationSurround+1 == lastPage-paginationEnd && page == lastPage-paginationEnd)))) { // Merge Middle and End if close enough
			if (page != start) {
				processPaginationSection(pageFn, ret, currPage, start, page - 1);
			}
			start = page + 1
		}
	}
	if (start < lastPage) {
		processPaginationSection(pageFn, ret, currPage, start, lastPage);
	}
	return ret;
      },
      index2HTML = (base: HTMLDivElement, index: number[], page = 0) => {
	if (index.length === 0) {
		clearElement(base);
	}
	const max = Math.min((page + 1) * perPage, index.length),
	      list = ul({"class": "results"});
	for (let i = page * perPage; i < max; i++) {
		const me = index[i],
		      [fName, lName,,,, childOf, ...spouseOf] = people[me],
		      [father, mother, ...siblings] = families[childOf],
		      spouseOfFams = spouseOf.map(s => families[s]);
		list.appendChild(li([
			div(span(`${fName} ${lName}`)),
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
	createHTML(clearElement(base), [
		pagination(base, index, page),
		list,
		pagination(base, index, page)
	]);
      };

for (let i = 0; i < people.length; i++) {
	let fl = people[i][1].charCodeAt(0);
	if (fl >= 97) {
		fl -= 32;
	}
	if (fl >= 65 && fl <=90) {
		indexes[fl - 65].push(i);
	}
}

for (const index of indexes) {
	index.sort(sortIDs);
}

export default function(base: HTMLElement) {
	const d = div(),
	      search = () => {
		const terms = s.value.toUpperCase().split(" "),
		      index: number[] = [];
		for (let i = 0; i < people.length; i++) {
			const name = `${people[i][0]} ${people[i][1]}`.toUpperCase();
			if (terms.every(term => name.includes(term))) {
				index.push(i);
			}
		}
		index2HTML(d, index.sort(sortIDs))
	      },
	      s = input({"type": "text", "onkeypress": (e: KeyboardEvent) => e.key === "Enter" && search()});
	createHTML(clearElement(base), [
		h2("Select a Name"),
		div({"id": "indexes"}, indexes.map((_, id) => span({"onclick": () => index2HTML(d, indexes[id])}, String.fromCharCode(id + 65)))),
		div({"id": "index_search"}, [
			label({"for": "index_search"}, "Search Terms"),
			s,
			button({"onclick": search}, "Search"),
		]),
		d
	]);
}