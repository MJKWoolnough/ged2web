package main

const (
	htmlStart = "<html><head><title>Ged2Web</title><script type=\"module\">"
	htmlEnd   = "</script><style type=\"text/css\">#indexes a:not(:first-child){padding-left:.5em}#indexes a:not(:last-child){border-right:1px solid #000;padding-right:.5em}.pagination .prev{padding-right:1em}.pagination .next{padding-left:1em}.results{list-style:none;margin:20px 0;padding:0;overflow:hidden;clear:left}.results>li{padding-bottom:10px;border-bottom:1px dashed #000;margin-bottom:10px;overflow:hidden}.results>li>div:first-child{float:left;width:200px}.results>li>div:last-child{overflow:hidden}.results>li>div:last-child>div{display:inline;margin-right:20px;overflow:hidden}#relationship{border-collapse:separate;border-spacing:5px 0;margin-left:auto;margin-right:auto}#relationship ul{list-style:none;margin:0;padding:0}#relationship td:first-child{background-color:rgba(0,255,0,0.5);text-align:right}#relationship td:last-child{background-color:rgba(255,0,0,0.5);text-align:left}#relationship tr:last-child td{background-color:rgba(255,255,0,0.5);text-align:center}#relationship td{vertical-align:bottom;padding:0 20px;width:50%}#ged2web_title{text-align:center}.person{position:absolute;width:150px;left:50px;border:1px solid #000;text-align:center;word-wrap:break-word}.highlight{border-color:#080}.sex_U{background-color:#aaa}.sex_M{background-color:#aaf}.sex_F{background-color:#faa}.spouseLine{position:absolute;border-top:5px solid #000;border-bottom:5px solid #000;height:5px;width:200px;left:50px;z-index:-1}.downLeft{position:absolute;margin-top:10px;height:100px;width:100px;border-right:5px solid #000;border-bottom:5px solid #000}.downRight{position:absolute;margin-top:10px;height:100px;width:100px;border-left:5px solid #000;border-bottom:5px solid #000}.chosen{border-color:#f00 !important}.dob,.dod{font-size:.5em}.collapse,.expand{width:7px;height:7px;border-width:0 1px 1px 0;border-color:#000;border-style:solid;float:left;background-color:#fff;cursor:pointer;cursor:hand}.collapse{background-color:#000}</style></head><body></body></html>"
	modStart  = "export const people = ["
	modMid    = "], families = ["
	modEnd    = "]"
	jsStart   = "class e{constructor(e,t){this.container=g(),this.rows=[],this.chosen=e,this.highlight=new Set(t),this.expanded=new Set(t),this.draw()}draw(e=0,t=0){let i=this.chosen,o=0,r=0;for(;;){const e=A[i][5],[t,o]=G[e];if(this.expanded.has(t))i=t;else{if(!this.expanded.has(o)){if(e)new s(this,t,0,!0);else{const e=new s(this,0,0);e.spouses=[new n(this,e,[0,0,i],0)]}break}i=o}}const l=h(null);for(const t of this.rows){for(const i of t){const t=X+o*J,c=z+i.col*Y;e&&i.id===e&&(r=c),i instanceof s?(o>0&&l.append(g({class:\"downLeft\",style:{top:t-50+\"px\",left:`${c+K/2}px`,width:0,height:50}})),i.spouses.length>0&&l.append(g({class:\"spouseLine\",style:{top:`${t}px`,left:`${c}px`,width:(i.spouses[i.spouses.length-1].col-i.col)*Y+\"px\"}}))):i instanceof n&&i.children.length>0&&(l.append(g(i.col<=i.children[0].col?{class:\"downRight\",style:{top:`${t}px`,left:c-Q/2+\"px\"}}:{class:\"downLeft\",style:{top:`${t}px`,left:c-K+Q-Q/2+\"px\",width:K-Q+\"px\"}})),i.children.length>1&&l.append(g({class:\"downLeft\",style:{top:t+J-50+\"px\",left:`${z+i.children[0].col*Y+K/2}px`,width:(i.children[i.children.length-1].col-i.children[0].col)*Y+\"px\",height:0}})));const a=i.id,[,,h,d,p,,...f]=A[i.id],u=i instanceof n;l.append(g({class:Z[p]+(this.highlight.has(a)?\" highlight\":\"\")+(this.chosen===i.id?\" chosen\":\"\"),style:{top:`${X+o*J}px`,left:`${z+i.col*Y}px`}},[i.id>0&&i.id!==this.chosen&&(i instanceof s&&f.length>0||u)?g({class:!this.expanded.has(i.id)||u?\"expand\":\"collapse\",onclick:this.expand.bind(this,i.id,u)}):[],g({class:\"name\"},I(i.id)),h?g({class:\"dob\"},h):[],d?g({class:\"dod\"},d):[]]))}o++}d(this.container).append(l),e&&window.scroll({left:r+t})}addPerson(e,t){const s=this.rows[e],n=s?.[s?.length-1];return s?(s.push(n.next=t),n.col+1):(this.rows.push([t]),0)}expand(e,t,s){t?this.chosen=e:this.expanded.has(e)?this.expanded.delete(e):this.expanded.add(e),this.rows=[],this.draw(e,window.scrollX-s.target.offsetParent.offsetLeft)}}class t{constructor(e,t,s){this.id=t,this.col=e.addPerson(s,this)}}class s extends t{constructor(e,t,s,i=!1){if(super(e,t,s),this.spouses=[],e.expanded.has(t)||e.chosen===t||i){const[,,,,,,...i]=A[t];if(i.length>0){for(const t of i)this.spouses.push(new n(e,this,G[t],s));this.col=this.spouses[0].col-1}}}shift(e){this.col+=e,this.next&&this.next.col<=this.col&&this.next.shift(this.col-this.next.col+1)}}class n extends t{constructor(e,t,n,i){super(e,n[0]===t.id?n[1]:n[0],i),this.children=[];const[,,...o]=n,r=i+1;if(o.length>0){for(const t of o)this.children.push(new s(e,t,r));for(let e=this.children.length-1;e>=0;e--){const t=this.children[e];t.col<this.col-1&&t.next&&(t.col=t.next.col-1)}this.col<this.children[0].col&&(this.col=this.children[0].col),this.shift(0)}}shift(e){if(this.col+=e,0!==this.children.length)for(;this.children[this.children.length-1].col<this.col-1;)this.children[0].shift(this.col-this.children[this.children.length-1].col-1);this.next&&this.next.col<=this.col&&this.next.shift(this.col-this.next.col+1)}}function i({l:e,q:t,p:s=0}){const n=g(),i=Math.max(0,\"string\"==typeof s?parseInt(s)||0:s);if(se(\"List\"),\"string\"==typeof t){V.value=t,se(\"Search\");const e=V.value.toUpperCase().split(\" \").sort(),s=e.join(\" \");let o=[];if(H.has(s))o=H.get(s);else{for(let t=0;t<A.length;t++){const s=`${A[t][0]||\"\"} ${A[t][1]||\"\"}`.toUpperCase();e.every((e=>s.includes(e)))&&o.push(t)}H.set(s,o)}B(n,o.sort(U),{q:t},i)}else if(\"string\"==typeof e){const t=e.toUpperCase().charCodeAt(0);t>=65&&t<=90&&(se(`List - ${e}`),B(n,O[t-65],{l:e},i))}return[o||(o=g({id:\"ged2web_title\"},[m(\"Select a Name\"),g({id:\"indexes\"},O.map(((e,t)=>h(te(\"list\",{l:String.fromCharCode(t+65)}),String.fromCharCode(t+65))))),g({id:\"index_search\"},[u({id:\"treeNames\"},A.map((([e=\"\",t=\"\"])=>e&&t?b({value:`${e} ${t}`}):[]))),$({for:\"index_search\"},\"Search Terms: \"),V,f({onclick:W},\"Search\")])])),n]}let o,r,l=0,c=\"\";const a=(e,t)=>{if(\"string\"==typeof t)e.appendChild(document.createTextNode(t));else if(Array.isArray(t))for(const s of t)a(e,s);else if(t instanceof Node)e.appendChild(t);else if(t instanceof NodeList)for(const s of t)e.appendChild(s)},h=(e,t,s)=>{const n=\"string\"==typeof e?document.createElementNS(\"http://www.w3.org/1999/xhtml\",e):e instanceof Node?e:document.createDocumentFragment();if(!(\"string\"==typeof t||t instanceof Array||t instanceof NodeList||t instanceof Node)&&(\"object\"!=typeof s||s instanceof Array||s instanceof Node||s instanceof NodeList)||([t,s]=[s,t]),\"object\"==typeof t&&n instanceof Element)for(const[e,s]of Object.entries(t))if(s instanceof Function){const t={};let i=e;e:for(;;){switch(i.charAt(0)){case\"1\":t.once=!0;break;case\"C\":t.capture=!0;break;case\"P\":t.passive=!0;break;default:break e}i=i.slice(1)}i.startsWith(\"on\")&&n.addEventListener(i.substr(2),s,t)}else if(\"class\"===e&&(s instanceof Array||s instanceof DOMTokenList)&&s.length>0)n.classList.add(...s);else if(\"style\"===e&&\"object\"==typeof s&&(n instanceof HTMLElement||n instanceof SVGElement))for(const e in s)void 0===s[e]?n.style.removeProperty(e):n.style.setProperty(e,s[e]);else\"string\"==typeof s||\"number\"==typeof s?n.setAttribute(e,s):\"boolean\"==typeof s?n.toggleAttribute(e,s):void 0===s&&n.hasAttribute(e)&&n.removeAttribute(e);return\"string\"==typeof s?n.textContent=s:s&&(s instanceof Array||s instanceof Node||s instanceof NodeList)&&a(n,s),n},d=e=>{for(;null!==e.lastChild;)e.removeChild(e.lastChild);return e},[p,f,u,g,m,w,x,$,y,b,C,v,L,k,N,S]=\"a button datalist div h2 h3 input label li option span table tbody td tr ul\".split(\" \").map((e=>h.bind(null,e))),A=[],G=["
	jsMid     = "],P=["
	jsEnd     = "[\"Parent\",\"Father\",\"Mother\"],[\"Sibling\",\"Brother\",\"Sister\"],[\"Spouse\",\"Husband\",\"Wife\"],[\"Child\",\"Son\",\"Daughter\"],[\"Pibling\",\"Uncle\",\"Aunt\"],[\"Nibling\",\"Nephew\",\"Neice\"]],T=(e,t)=>{const s=[];for(;t>0;)s.push(t),t=e.get(t)[0];return s.reverse()},j={one:\"st\",two:\"nd\",few:\"rd\",other:\"th\"},M=[\"Once\",\"Twice\",\"Thrice\"],E=new Intl.PluralRules(\"en-GB\",{type:\"ordinal\"}),_=(e,t,s)=>{const n=e.length,i=t.length,o=n>0&&i>0&&A[e[n-1]][5]!=A[t[i-1]][5]?\"Half-\":\"\";switch(n){case 0:switch(i){case 0:return\"Clone\";case 1:return P[0][s];default:const e=i-2;return e>3?`${e} x Great-Grand-`:\"Great-\".repeat(e)+`Grand-${P[0][s]}`}case 1:switch(i){case 0:return P[3][A[e[0]][4]];case 1:return`${o}${P[1][A[e[0]][4]]}`;default:const t=i-2;return`${o}${t>3?`${t} x Great-Grand-`:\"Great-\".repeat(t)}${P[4][A[e[0]][4]]}`}break;default:const t=n-2;switch(i){case 0:return(t>3?`${t} x Great-Grand-`:\"Great-\".repeat(t)+\"Grand-\")+P[3][A[e[0]][4]];case 1:return`${o}${t>3?`${t} x Great-`:\"Great-\".repeat(t)}${P[5][A[e[0]][4]]}`;default:const s=Math.min(n,i)-1,r=Math.abs(n-i);return`${o}${s}${j[E.select(s)]} ${o}Cousin${r>0?`, ${M[r-1]||`${r} Times`} Removed`:\"\"}`}}},I=e=>`${A[e][0]??\"?\"} ${A[e][1]??\"?\"}`,O=Array.from({length:26},(()=>[])),R=(new Intl.Collator).compare,U=(e,t)=>{const[s=\"\",n=\"\"]=A[e],[i=\"\",o=\"\"]=A[t];return n!==o?R(n,o):s!==i?R(s,i):t-e},F=(e,t)=>0===e?[]:g([h(te(\"tree\",{id:e}),I(e)),\" (\"+P[t][A[e][4]]+\")\"]),D=(e,t,s,n,i)=>{0!==e.length&&e.push(\"…\");for(let o=s;o<=n;o++)o!==s&&e.push(\", \"),e.push(t===o?C(o+1+\"\"):h(te(\"list\",Object.assign({p:o},i)),{class:\"pagination_link\"},o+1+\"\"))},q=(e,t,s=0)=>{const n=Math.ceil(e.length/20)-1,i=[];if(0===n)return[];s>n&&(s=n);let o=0;for(let e=0;e<=n;e++)e<3||e>n-3||(3>s||e>=s-3)&&e<=s+3||s-3-1==3&&3==e||s+3+1==n-3&&e==n-3||(e!=o&&D(i,s,o,e-1,t),o=e+1);return o<n&&D(i,s,o,n,t),g({class:\"pagination\"},[\"Pages: \",h(0!==s?te(\"list\",Object.assign({p:s-1},t)):C(),{class:\"pagination_link prev\"},\"Previous\"),i,h(s!==n?te(\"list\",Object.assign({p:s+1},t)):C(),{class:\"pagination_link next\"},\"Next\")])},B=(e,t,s,n=0)=>{const i=Math.min(20*(n+1),t.length),o=S({class:\"results\"}),r=[];for(let e=20*n;e<i;e++){const s=t[e],[,,,,,n,...i]=A[s],[c,a,...d]=G[n],p=i.map((e=>G[e])),u=f({onclick:()=>{if(l===s){l=0;for(const e of r)e.innerText=\"+\"}else if(0===l){l=s;for(const e of r)e.innerText=\"=\";u.innerText=\"-\"}else ee(\"fhcalc\",{from:l,to:s})}},0===l?\"+\":l===s?\"-\":\"=\");r.push(u),o.appendChild(y([g([h(te(\"tree\",{id:s}),I(s)),u]),g([F(c,0),F(a,0),d.filter((e=>e!==s)).map((e=>F(e,1))),p.map((([e,t,...n])=>[F(e!==s?e:t,2),n.filter((e=>e!==s)).map((e=>F(e,3)))]))])]))}h(e,[q(t,s,n),o,q(t,s,n)])},H=new Map,W=()=>ee(\"list\",{q:V.value}),V=x({type:\"text\",list:\"treeNames\",onkeypress:e=>\"Enter\"===e.key&&W()}),X=100,z=50,J=150,K=150,Q=50,Y=K+Q,Z=[\"U\",\"M\",\"F\"].map((e=>`person sex_${e}`)),ee=(t,s,n=!1)=>{let o,l=\"list\";switch(t){case\"tree\":o=function({id:t,highlight:s}){const n=\"string\"==typeof t?parseInt(t):t;if(!(n<=0||void 0===A[n]))return se(`Family Tree - ${I(n)}`),new e(n,(s||\"\").split(\".\").map((e=>parseInt(e))).filter((e=>e>0))).container}(s),l=t;break;case\"fhcalc\":o=function({from:e,to:t}){const s=\"string\"==typeof e?parseInt(e):e,n=\"string\"==typeof t?parseInt(t):t;if(s<=0||n<=0||!A[s]||!A[n])return;se(\"Relationship Calculator\");const[i,o,r]=((e,t)=>{const s=new Map([[e,[0,0]],[t,[0,0]]]),n=[[e,1],[t,2]];for(;n.length>0;){const e=n.shift(),[t,i]=e,[,,,,,o]=A[t];for(const r of G[o].slice(0,2))if(0!==r){const o=s.get(r);if(o){if(o[1]!==i)return[r,T(s,1===i?t:o[0]),T(s,1===i?o[0]:t)]}else s.set(r,e),n.push([r,i])}}return[0,[],[]]})(s,n),l=I(s),c=I(n);return 0===i?m(`No direct relationship betwen ${l} and ${c}`):[g({id:\"ged2web_title\"},[m(`${l} is the ${_(o,r,A[i][4])} of ${c}`),f({onclick:()=>ee(\"fhcalc\",{from:n,to:s})},\"Swap\")]),v({id:\"relationship\"},L([N([l,c].map((e=>k(w(`Route from ${e}`))))),N([o,r].map((e=>k(S(e.map((e=>y(`${I(e)}, who is the ${P[3][A[e][4]]} of…`)))))))),N(k({colspan:2},[g(I(i)),w(\"Common Ancestor\"),h(te(\"tree\",{id:i,highlight:o.concat(r).join(\".\")}),\"Show in Tree\")]))]))]}(s),l=t;break;case\"list\":o=i(s)}n||history.pushState(null,\"\",oe(t,s)),c?document.body.classList.replace(c,c=\"ged2web_\"+l):document.body.classList.add(c=\"ged2web_\"+l),h(d(r),o||i({}))},te=(e,t)=>p({href:oe(e,t),onclick:s=>{s.preventDefault(),ee(e,t)}}),se=e=>document.title=`${le} - ${e}`,ne=()=>window.location.pathname.split(\"/\").pop()?.split(\".\").shift(),ie=[\"list\",\"fhcalc\",\"tree\"].includes(ne()),oe=(e,t)=>(ie?`${e}.html?`:`?module=${e}&`)+Object.entries(t).map((([e,t])=>`${e}=${encodeURIComponent(t)}`)).join(\"&\"),re=()=>{const e=Object.fromEntries(new URL(window.location+\"\").searchParams.entries());ee(ie?ne():e.module,e,!0)},le=document.title;for(let e=0;e<A.length;e++){let t=(A[e][1]??\"\").charCodeAt(0);t>=97&&(t-=32),t>=65&&t<=90&&O[t-65].push(e)}for(const e of O)e.sort(U);window.addEventListener(\"popstate\",re),(\"complete\"==document.readyState?Promise.resolve():new Promise((e=>globalThis.addEventListener(\"load\",e,{once:!0})))).then((()=>{r=document.getElementById(\"ged2web\")||document.body,re()}));"
)