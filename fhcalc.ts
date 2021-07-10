import {relations} from './shared.js';
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
      };
