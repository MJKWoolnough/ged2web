import type {Children} from './lib/dom.js';
import {createHTML, clearElement} from './lib/dom.js';
import {ul, li, div, span, h2, label, input, button} from './lib/html.js'
import {link, relations, nameOf} from './shared.js';
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
      person2HTML = (id: number, rel: number) => {
	if (id === 0) {
		return [];
	}
	const person = people[id];
	return div([
		span(nameOf(id)),
		" (" + relations[rel][person[4]] + ")"
	])
      },
      perPage = 20,
      paginationEnd = 3,
      paginationSurround = 3,
      processPaginationSection = (pageFn: (page: number) => void, ret: Children[], currPage: number, from: number, to: number, params: string)=> {
	if (ret.length !== 0) {
		ret.push("…");
	}
	for (let i = from; i <= to; i++) {
		if (i !== from) {
			ret.push(", ");
		}
		ret.push(currPage === i ? span((i+1)+"") : createHTML(link("list", `${params}&page=${i}`, pageFn.bind(null, i)), {"class": "pagination_link"}, (i+1)+""));
	}
      },
      pagination = (pageFn: (page: number) => void, index: number[], params: string, currPage = 0) => {
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
		if (!(page < paginationEnd || // Beginning
			page > lastPage-paginationEnd || // End
			((paginationSurround > currPage || page >= currPage-paginationSurround) && page <= currPage+paginationSurround) || // Middle
			paginationEnd > 0 && ((currPage-paginationSurround-1 == paginationEnd && page == paginationEnd) || // Merge Begining and Middle if close enough
			(currPage+paginationSurround+1 == lastPage-paginationEnd && page == lastPage-paginationEnd)))) { // Merge Middle and End if close enough
			if (page != start) {
				processPaginationSection(pageFn, ret, currPage, start, page - 1, params);
			}
			start = page + 1
		}
	}
	if (start < lastPage) {
		processPaginationSection(pageFn, ret, currPage, start, lastPage, params);
	}
	return div({"class": "pagination"}, [
		"Pages: ",
		createHTML(currPage !== 0 ? link("list", `${params}&page=${currPage-1}`, () => pageFn(currPage - 1)) : span(), {"class": "pagination_link prev"}, "Previous"),
		ret,
		createHTML(currPage !== lastPage ? link("list", `${params}&page=${currPage+1}`, () => pageFn(currPage + 1)) : span(), {"class": "pagination_link next"}, "Next"),
	]);
      },
      index2HTML = (base: HTMLDivElement, index: number[], params: string, page = 0) => {
	if (index.length === 0) {
		clearElement(base);
	}
	const max = Math.min((page + 1) * perPage, index.length),
	      list = ul({"class": "results"}),
	      pageFn = (page: number) => index2HTML(base, index, params, page);
	for (let i = page * perPage; i < max; i++) {
		const me = index[i],
		      [,,,,, childOf, ...spouseOf] = people[me],
		      [father, mother, ...siblings] = families[childOf],
		      spouseOfFams = spouseOf.map(s => families[s]);
		list.appendChild(li([
			div(span(nameOf(me))),
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
		pagination(pageFn, index, params, page),
		list,
		pagination(pageFn, index, params, page)
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
		index2HTML(d, index.sort(sortIDs), `search=${s.value}`)
	      },
	      s = input({"type": "text", "onkeypress": (e: KeyboardEvent) => e.key === "Enter" && search()});
	createHTML(clearElement(base), [
		h2("Select a Name"),
		div({"id": "indexes"}, indexes.map((_, id) => createHTML(link("list", `l=${String.fromCharCode(id+65)}`, () => index2HTML(d, indexes[id], `l=${String.fromCharCode(id+65)}`)), String.fromCharCode(id + 65)))),
		div({"id": "index_search"}, [
			label({"for": "index_search"}, "Search Terms: "),
			s,
			button({"onclick": search}, "Search"),
		]),
		d
	]);
}
