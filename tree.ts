import {createHTML, clearElement} from './lib/dom.js';
import {div} from './lib/html.js';
import {nameOf} from './list.js';
import {people, families} from './gedcom.js';

const rowStart = 100,
      colStart = 50,
      rowGap = 150,
      boxWidth = 150,
      boxPadding = 50,
      colGap = boxWidth + boxPadding,
      classes = ["U", "M", "F"].map(s => `person sex_${s}`);

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
		let top = this.chosen,
		    r = 0;
		while (true) {
			const famc = people[top][5],
			      [father, mother] = families[famc]
			if (this.expanded.has(father)) {
				top = father;
			} else if (this.expanded.has(mother)) {
				top = mother;
			} else {
				if (famc) {
					new Person(this, father, 0, true);
				} else {
					const p = new Person(this, 0, 0);
					p.spouses = [new Spouse(this, p, [0, 0, top], 0)];
				}
				break;
			}
		}
		const elms = createHTML(null) as DocumentFragment;
		for (const row of this.rows) {
			for (const p of row) {
				const top = rowStart + r * rowGap,
				      left = colStart + p.col * colGap;
				if (p instanceof Person) {
					if (r > 0) {
						elms.append(div({"class": "downLeft", "style": {"top": `${top - 50}px`, "left": `${left + boxWidth / 2}px`, "width": 0, "height": 50}}));
					}
					if (p.spouses.length > 0) {
						elms.append(div({"class": "spouseLine", "style": {"top": `${top}px`, "left": `${left}px`, "width": `${(p.spouses[p.spouses.length-1].col - p.col) * colGap}px`}}));
					}
				} else if (p instanceof Spouse && p.children.length > 0) {
					if (p.col <= p.children[0].col) {
						elms.append(div({"class": "downRight", "style": {"top": `${top}px`, "left": `${left - boxPadding / 2}px`}}));
					} else {
						elms.append(div({"class": "downLeft", "style": {"top": `${top}px`, "left": `${left - boxWidth + boxPadding - boxPadding / 2}px`, "width": `${boxWidth - boxPadding}px`}}));
					}
					if (p.children.length > 1) {
						elms.append(div({"class": "downLeft", "style": {"top": `${top + rowGap - 50}px`, "left": `${colStart + p.children[0].col * colGap + boxWidth / 2}px`, "width": `${(p.children[p.children.length-1].col - p.children[0].col) * colGap}px`, "height": 0}}));
					}
				}
				const id = p.id,
				      [,, dob, dod, gender,, ...fams] = people[p.id];
				elms.append(div({"class": classes[gender] + (this.highlight.has(id) ? " highlight" : ""), "style": {"top": `${rowStart + r * rowGap}px`, "left": `${colStart + p.col * colGap}px`}, "id": this.chosen === id ? "chosen" : undefined}, [
					p.id > 0 && p.id !== this.chosen && (p instanceof Person && fams.length > 0 || p instanceof Spouse) ? div({"class": !this.expanded.has(p.id) || p instanceof Spouse ? "expand" : "collapse", "onclick": this.expand.bind(this, p.id, p instanceof Spouse)}) : [],
					div({"class": "name"}, nameOf(p.id)),
					dob ? div({"class": "dob"}, dob) : [],
					dod ? div({"class": "dod"}, dod) : [],
				]));
			}
			r++;
		}
		clearElement(this.container).append(elms);
	}
	addPerson(row: number, p: PersonBox) {
		while (this.rows.length <= row) {
			this.rows.push([]);
		}
		this.rows[row].push(p);
		if (this.rows[row].length === 1) {
			return 0;
		}
		return this.rows[row][this.rows[row].length-2].col + 1;
	}
	expand(id: number, isSpouse: boolean) {
		if (isSpouse) {
			this.chosen = id;
		} else if (this.expanded.has(id)) {
			this.expanded.delete(id);
		} else {
			this.expanded.add(id);
		}
		this.rows = [];
		this.draw();
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
	constructor(tree: Tree, id: number, row: number, forced = false) {
		super(tree, id, row);
		if (tree.expanded.has(id) || tree.chosen === id || forced) {
			const [,,,,,, ...spouses] = people[id];
			if (spouses.length > 0) {
				for (const fams of spouses) {
					this.spouses.push(new Spouse(tree, this, families[fams], row));
				}
				this.col = this.spouses[0].col - 1;
			}
		}
	}
	shift(num: number) {
		this.col += num;
		for (const spouse of this.spouses) {
			spouse.shift(num);
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
		if (children.length > 0) {
			for (const child of children) {
				this.children.push(new Person(tree, child, crow));
			}
			if (this.col < this.children[0].col) {
				this.col = this.children[0].col;
			} else if (this.children[children.length-1].col < this.col - 1)  {
				for (const child of this.children) {
					child.shift(this.col - this.children[children.length-1].col - 1);
				}
			}
		}
	}
	shift(num: number) {
		this.col += num;
		if (this.children[this.children.length-1].col < this.col - 1)  {
			for (const child of this.children) {
				child.shift(this.col - this.children[this.children.length-1].col - 1);
			}
		}
	}
}

export default function({"id": idStr, "highlight": highlightStr}: Record<string, string | number>) {
	const id = typeof idStr === "string" ? parseInt(idStr) : idStr;
	if (id <= 0 || people[id] === undefined) {
		return undefined;
	}
	return new Tree(id, ((highlightStr as string) || "").split(",").map(id => parseInt(id)).filter(id => id > 0)).container;
}
