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
      perPage = 20,
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
      pagination = (index: number[], page = 0) => {
	      return [index.length + "", page + ""];
      },
      index2HTML = (index: number[], page = 0) => {
	if (index.length === 0) {
		return [];
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
	return [
		pagination(index, page),
		list,
		pagination(index, page)
	];
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
		index.sort(sortIDs);
		createHTML(clearElement(d), index2HTML(index))
	      },
	      s = input({"type": "text", "onkeypress": (e: KeyboardEvent) => e.key === "Enter" && search()});
	createHTML(clearElement(base), [
		h2("Select a Name"),
		div({"id": "indexes"}, indexes.map((_, id) => span({"onclick": () => createHTML(clearElement(d), index2HTML(indexes[id]))}, String.fromCharCode(id + 65)))),
		div({"id": "index_search"}, [
			label({"for": "index_search"}, "Search Terms"),
			s,
			button({"onclick": search}, "Search"),
		]),
		d
	]);
}
