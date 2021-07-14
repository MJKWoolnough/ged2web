import {div} from './lib/html.js';
import {people, families} from './gedcom.js';

class Tree {
	container = div();
	chosen: number;
	highlight: Set<number>;
	expanded = new Set<number>();
	rows: PersonBox[][] = [];
	constructor(id: number, highlight: number[]) {
		this.chosen = id;
		this.highlight = new Set(highlight);
		this.draw();
	}
	draw() {
		let top = this.chosen;
		while (true) {
			const famc = people[top][5],
			      [father, mother] = families[famc]
			if (this.expanded.has(father)) {
				top = father;
			} else if (this.expanded.has(mother)) {
				top = mother;
			} else {
				if (famc) {
					new Person(this, father, 0);
				} else {
					const p = new Person(this, 0, 0);
					p.spouses = [new Spouse(this, p, [0, 0, top], 0)];
				}
				break;
			}
		}
	}
	addPerson(row: number, p: PersonBox) {
		while (this.rows.length <= row) {
			this.rows.push([]);
		}
		return this.rows[row].push(p) - 1;
	}
}

abstract class PersonBox {
	id: number;
	col: number;
	constructor(tree: Tree, id: number, row: number) {
		this.id = id;
		this.col = tree.addPerson(row, this);
	}
}

class Person extends PersonBox {
	spouses: Spouse[] = [];
	constructor(tree: Tree, id: number, row = 0) {
		super(tree, id, row);
		const [,,,,,, ...spouses] = people[id];
		for (const fams of spouses) {
			this.spouses.push(new Spouse(tree, this, families[fams], row));
		}
	}
}

class Spouse extends PersonBox {
	children: Person[] = [];
	spouse: Person;
	constructor(tree: Tree, spouse: Person, fams: [number, number, ...number[]], row: number) {
		super(tree, fams[0] === spouse.id ? fams[1] : fams[0], row);
		this.spouse = spouse;
		const [,, ...children] = fams,
		      crow = row + 1;
		for (const child of children) {
			this.children.push(new Person(tree, child, crow));
		}
	}
}

export default function({idStr, highlightStr}: Record<string, string | number>) {
	new Tree(typeof idStr === "string" ? parseInt(idStr) : idStr, (highlightStr as string).split(",").map(id => parseInt(id)).filter(id => id > 0)).container;
}
