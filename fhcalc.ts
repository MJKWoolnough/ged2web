import {makeElement} from './lib/dom.js';
import {div, h2, h3, button, table, tbody, tr, td, ul, li} from './lib/html.js';
import {load, link, setTitle} from './ged2web.js';
import {nameOf} from './list.js';
import {people, families} from './gedcom.js';

type Connection = [number, number];

export const relations = [
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
];

const makeRoute = (connMap: Map<number, Connection>, pid: number) => {
	const route: number[] = [];
	while (pid > 0) {
		route.push(pid);
		pid = connMap.get(pid)![0];
	}
	return route.reverse();
      },
      findConn = (from: number, to: number): [number, number[], number[]] => {
	const connMap = new Map<number, Connection>([[from, [0, 0]], [to, [0, 0]]]),
	      todo: Connection[] = [[from, 1], [to, 2]];
	while (todo.length > 0) {
		const person = todo.shift()!,
		      [pid, side] = person,
		      childOf = people[pid][5];
		for (const parent of families[childOf].slice(0, 2)) {
			if (parent !== 0) {
				const existing = connMap.get(parent);
				if (existing) {
					if (existing[1] !== side) {
						return [parent, makeRoute(connMap, side === 1 ? pid : existing[0]), makeRoute(connMap, side === 1 ? existing[0] : pid)];
					}
				} else {
					connMap.set(parent, person);
					todo.push([parent, side]);
				}
			}
		}
	}
	return [0, [], []];
      },
      ordinals: Record<string, string>  = {
	      "one": "st",
	      "two": "nd",
	      "few": "rd",
	      "other": "th",
      },
      times = ["Once", "Twice", "Thrice"],
      plurals = new Intl.PluralRules("en-GB", {"type": "ordinal"}),
      getRelationship = (first: number[], second: number[], commonGender: number) => {
	const up = first.length,
	      down = second.length,
	      half = up > 0 && down > 0 && people[first[up-1]][5] != people[second[down-1]][5] ? "Half-" : "";
	switch (up) {
	case 0:
		switch (down) {
		case 0:
			return "Clone";
		case 1:
			return relations[0][commonGender];
		default:
			const greats = down - 2;
			return greats > 3 ? `${greats} x Great-Grand-` : "Great-".repeat(greats) + `Grand-${relations[0][commonGender]}`;
		}
	case 1:
		switch (down) {
		case 0:
			return relations[3][people[first[0]][4]];
		case 1:
			return `${half}${relations[1][people[first[0]][4]]}`;
		default:
			const greats = down - 2;
			return `${half}${(greats > 3 ? `${greats} x Great-Grand-` : "Great-".repeat(greats))}${relations[4][people[first[0]][4]]}`;
		}
		break;
	default:
		const greats = up - 2;
		switch (down) {
		case 0:
			return (greats > 3 ? `${greats} x Great-Grand-` : "Great-".repeat(greats) + "Grand-") + relations[3][people[first[0]][4]];
		case 1:
			return `${half}${(greats > 3 ? `${greats} x Great-` : "Great-".repeat(greats))}${relations[5][people[first[0]][4]]}`;
		default:
			const small = Math.min(up, down) - 1,
			      diff = Math.abs(up - down);
			return `${half}${small}${ordinals[plurals.select(small)]} ${half}Cousin${diff > 0 ? `, ${times[diff - 1] || `${diff} Times`} Removed` : ""}`;
		}
	}
      };

export default ({from: fromStr, to: toStr}: Record<string, string | number>) => {
	const from = typeof fromStr === "string" ? parseInt(fromStr) : fromStr,
	      to = typeof toStr === "string" ? parseInt(toStr) : toStr;
	if (from <= 0 || to <= 0 || !people[from] || !people[to]) {
		return undefined;
	}
	setTitle("Relationship Calculator");
	const [common, first, second] = findConn(from, to),
	      aname = nameOf(from),
	      bname = nameOf(to);
	return common === 0 ? h2(`No direct relationship between ${aname} and ${bname}`) : [
		div({"id": "ged2web_title"}, [
			h2(`${aname} is the ${getRelationship(first, second, people[common][4])} of ${bname}`),
			button({"onclick": () => load("fhcalc", {"from": to, "to": from})}, "Swap"),
		]),
		table({"id": "relationship"}, tbody([
			tr([aname, bname].map(name => td(h3(`Route from ${name}`)))),
			tr([first, second].map(list => td(ul(list.map(id => li(`${nameOf(id)}, who is the ${relations[3][people[id][4]]} of…`)))))),
			tr(td({"colspan": 2}, [
				div(nameOf(common)),
				h3("Common Ancestor"),
				makeElement(link("tree", {"id": common, "highlight": first.concat(second).join(".")}), "Show in Tree")
			]))
		]))
	];
}
