import {clearNode} from './lib/dom.js';
import pageLoad from './lib/load.js';
import {router} from './lib/router.js';
import fhcalc from './fhcalc.js';
import list from './list.js';
import tree from './tree.js';

pageLoad.then(() => clearNode(document.getElementById("ged2web") ?? document.body, router().add("tree.html?id=:id&highlight=:highlight", tree).add("?module=tree&id=:id&highlight=:highlight", tree).add("fhcalc.html?from=:from&to=:to", fhcalc).add("?module=fhcalc&from=:from&to=:to", fhcalc).add("list.html?l=:l&q=:q&p=:p", list).add("?module=list&l=:l&q=:q&p=:p", list).add("", list)));
