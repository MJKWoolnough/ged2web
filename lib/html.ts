import type {DOMBind} from './dom.js';
import {createHTML} from './dom.js';

export {createHTML};

export const [a, button, datalist, div, h2, h3, input, label, li, option, span, table, tbody, td, tr, ul] = "a button datalist div h2 h3 input label li option span table tbody td tr ul".split(" ").map(e => createHTML.bind(null, e)) as [DOMBind<HTMLElementTagNameMap["a"]>, DOMBind<HTMLElementTagNameMap["button"]>, DOMBind<HTMLElementTagNameMap["datalist"]>, DOMBind<HTMLElementTagNameMap["div"]>, DOMBind<HTMLElementTagNameMap["h2"]>, DOMBind<HTMLElementTagNameMap["h3"]>, DOMBind<HTMLElementTagNameMap["input"]>, DOMBind<HTMLElementTagNameMap["label"]>, DOMBind<HTMLElementTagNameMap["li"]>, DOMBind<HTMLElementTagNameMap["option"]>, DOMBind<HTMLElementTagNameMap["span"]>, DOMBind<HTMLElementTagNameMap["table"]>, DOMBind<HTMLElementTagNameMap["tbody"]>, DOMBind<HTMLElementTagNameMap["td"]>, DOMBind<HTMLElementTagNameMap["tr"]>, DOMBind<HTMLElementTagNameMap["ul"]>];
