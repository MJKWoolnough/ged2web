import {createHTML, clearElement} from './lib/dom.js';
import {div, h2, h3, button, table, tbody, tr, td, ul, li} from './lib/html.js';
import {relations, link, nameOf} from './shared.js';
import {people, families} from './gedcom.js';

type Connection = [number, number];

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
		      [,,,,, childOf] = people[pid];
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
      times = ["once", "twice", "thrice"],
      plurals = new Intl.PluralRules("en-GB", {"type": "ordinal"}),
      getRelationship = (first: number[], second: number[], commonGender: number) => {
	const up = first.length,
	      down = second.length;
	let relationship = "";
	if (up > 0 && down > 0 && people[first[up-1]][5] != people[second[down]][5]) {
		relationship = "Half-";
	}
	switch (up) {
	case 0:
		switch (down) {
		case 0:
			return "Clone";
		case 1:
			break;
		default:
			const greats = down - 2;
			relationship += greats > 3 ? `${greats} x Great-Grand-` : "Great-".repeat(greats) + "Grand-";
		}
		relationship += relations[0][commonGender];
		break;
	case 1:
		switch (down) {
		case 0:
			return relations[3][people[first[0]][4]];
		case 1:
			relationship += relations[1][people[first[0]][4]];
			break;
		default:
			const greats = down - 2;
			relationship += (greats > 3 ? `${greats} x Great-Grand-` : "Great-".repeat(greats)) + relations[4][people[first[0]][4]];
		}
		break;
	default:
		const greats = up - 2;
		switch (down) {
		case 0:
			return (greats > 3 ? `${greats} x Great-Grand-` : "Great-".repeat(greats) + "Grand-") + relations[3][people[first[0]][4]];
		case 1:
			relationship += (greats > 3 ? `${greats} x Great-` : "Great-".repeat(greats)) + relations[5][people[first[0]][4]];
			break;
		default:
			const small = Math.min(up, down) - 1,
			      diff = Math.abs(up - down);
			relationship += `${small}${ordinals[plurals.select(small)]} Cousin${diff > 0 ? `, ${diff < 4 ? times[diff] : `${diff} times`}` : ""}}`;
		}
	}
	return relationship;
      };

export default function (base: HTMLElement, from: number, to: number) {
	const [common, first, second] = findConn(from, to),
	      aname = nameOf(from),
	      bname = nameOf(to);
	createHTML(clearElement(base), common === 0 ? h2(`No direct relationship betwen ${aname} and ${bname}`) : [
		h2(`${aname} is the ${getRelationship(first, second, people[common][4])} of ${bname}`),
		button("Swap"),
		table(tbody([
			tr(([[aname, first], [bname, second]] as [string, number[]][]).map(([name, list]) => td([
				h3(`Route from ${name}`),
				ul(list.map(id => li(`${nameOf(id)}, who is the ${relations[3][people[id][4]]} of…`))),
			]))),
			tr({"colspan": 2}, td([
				div(nameOf(common)),
				h3("Common Ancestor"),
				link("tree", `id=${first}&highlight=${first.concat(common, second).join(",")}`, () => {})
			]))
		]))
	]);
}