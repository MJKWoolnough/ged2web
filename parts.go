package main

// File automatically generated with ./embed.sh

const (
	htmlStart = "<html lang=\"en\"><head><title>Ged2Web</title><meta charset=\"UTF-8\" /><script type=\"module\">"
	htmlEnd   = "</script><style type=\"text/css\">#indexes a:not(:first-child){padding-left:.5em}#indexes a:not(:last-child){border-right:1px solid #000;padding-right:.5em}.pagination .prev{padding-right:1em}.pagination .next{padding-left:1em}.results{list-style:none;margin:20px 0;padding:0;overflow:hidden;clear:left}.results>li{padding-bottom:10px;border-bottom:1px dashed #000;margin-bottom:10px;overflow:hidden}.results>li>div:first-child{float:left;width:200px}.results>li>div:last-child{overflow:hidden}.results>li>div:last-child>div{display:inline;margin-right:20px;overflow:hidden}#relationship{border-collapse:separate;border-spacing:5px 0;margin-left:auto;margin-right:auto}#relationship ul{list-style:none;margin:0;padding:0}#relationship td:first-child{background-color:rgba(0,255,0,0.5);text-align:right}#relationship td:last-child{background-color:rgba(255,0,0,0.5);text-align:left}#relationship tr:last-child td{background-color:rgba(255,255,0,0.5);text-align:center}#relationship td{vertical-align:bottom;padding:0 20px;width:50%}#ged2web_title{text-align:center}.person{position:absolute;width:150px;left:50px;border:1px solid #000;text-align:center;word-wrap:break-word}.highlight{border-color:#080}.sex_U{background-color:#aaa}.sex_M{background-color:#aaf}.sex_F{background-color:#faa}.spouseLine{position:absolute;border-top:5px solid #000;border-bottom:5px solid #000;height:5px;width:200px;left:50px;z-index:-1}.downLeft{position:absolute;margin-top:10px;height:100px;width:100px;border-right:5px solid #000;border-bottom:5px solid #000}.downRight{position:absolute;margin-top:10px;height:100px;width:100px;border-left:5px solid #000;border-bottom:5px solid #000}.chosen{border-color:#f00 !important}.dob,.dod{font-size:.5em}.collapse,.expand{width:7px;height:7px;border-width:0 1px 1px 0;border-color:#000;border-style:solid;float:left;background-color:#fff;cursor:pointer;cursor:hand}.collapse{background-color:#000}</style></head><body></body></html>"
	modStart  = "export const people = ["
	modMid    = "], families = ["
	modEnd    = "]"
	jsStart   = "class e{constructor(e,t){this.container=u(),this.rows=[],this.chosen=e,this.highlight=new Set(t),this.expanded=new Set(t),this.draw()}draw(e=0,t=0){let i=this.chosen,o=0,r=0;for(;;){const e=S[i][5],[t,o]=A[e];if(this.expanded.has(t))i=t;else{if(!this.expanded.has(o)){if(e)new s(this,t,0,!0);else{const e=new s(this,0,0);e.spouses=[new n(this,e,[0,0,i],0)]}break}i=o}}const l=a(null);for(const t of this.rows){for(const i of t){const t=X+o*J,c=z+i.col*Y;e&&i.id===e&&(r=c),i instanceof s?(o>0&&l.append(u({class:\"downLeft\",style:{top:t-50+\"px\",left:`${c+K/2}px`,width:0,height:\"50px\"}})),i.spouses.length>0&&l.append(u({class:\"spouseLine\",style:{top:`${t}px`,left:`${c}px`,width:(i.spouses[i.spouses.length-1].col-i.col)*Y+\"px\"}}))):i instanceof n&&i.children.length>0&&(l.append(u(i.col<=i.children[0].col?{class:\"downRight\",style:{top:`${t}px`,left:c-Q/2+\"px\"}}:{class:\"downLeft\",style:{top:`${t}px`,left:c-K+Q/2+\"px\",width:K-Q+\"px\"}})),i.children.length>1&&l.append(u({class:\"downLeft\",style:{top:t+J-50+\"px\",left:`${z+i.children[0].col*Y+K/2}px`,width:(i.children[i.children.length-1].col-i.children[0].col)*Y+\"px\",height:0}})));const a=i.id,[,,h,d,p,,...f]=S[i.id],g=i instanceof n;l.append(u({class:Z[p]+(this.highlight.has(a)?\" highlight\":\"\")+(this.chosen===i.id?\" chosen\":\"\"),style:{top:`${X+o*J}px`,left:`${z+i.col*Y}px`}},[i.id>0&&i.id!==this.chosen&&(i instanceof s&&f.length>0||g)?u({class:!this.expanded.has(i.id)||g?\"expand\":\"collapse\",onclick:this.expand.bind(this,i.id,g)}):[],u({class:\"name\"},E(i.id)),h?u({class:\"dob\"},h):[],d?u({class:\"dod\"},d):[]]))}o++}h(this.container).append(l),e&&window.scroll({left:r+t})}addPerson(e,t){const s=this.rows[e],n=s?.[s?.length-1];return s?(s.push(n.next=t),n.col+1):(this.rows.push([t]),0)}expand(e,t,s){t?this.chosen=e:this.expanded.has(e)?this.expanded.delete(e):this.expanded.add(e),this.rows=[],this.draw(e,window.scrollX-s.target.offsetParent.offsetLeft)}}class t{constructor(e,t,s){this.id=t,this.col=e.addPerson(s,this)}}class s extends t{constructor(e,t,s,i=!1){if(super(e,t,s),this.spouses=[],e.expanded.has(t)||e.chosen===t||i){const[,,,,,,...i]=S[t];if(i.length>0){for(const t of i)this.spouses.push(new n(e,this,A[t],s));this.col=this.spouses[0].col-1}}}shift(e){this.col+=e,this.next&&this.next.col<=this.col&&this.next.shift(this.col-this.next.col+1)}}class n extends t{constructor(e,t,n,i){super(e,n[0]===t.id?n[1]:n[0],i),this.children=[];const[,,...o]=n,r=i+1;if(o.length>0){for(const t of o)this.children.push(new s(e,t,r));for(let e=this.children.length-1;e>=0;e--){const t=this.children[e];t.col<this.col-1&&t.next&&(t.col=t.next.col-1)}this.col<this.children[0].col&&(this.col=this.children[0].col),this.shift(0)}}shift(e){if(this.col+=e,0!==this.children.length)for(;this.children[this.children.length-1].col<this.col-1;)this.children[0].shift(this.col-this.children[this.children.length-1].col-1);this.next&&this.next.col<=this.col&&this.next.shift(this.col-this.next.col+1)}}let i,o,r=0,l=\"\";const c=(e,t)=>{if(\"string\"==typeof t)e.appendChild(document.createTextNode(t));else if(Array.isArray(t))for(const s of t)c(e,s);else if(t instanceof Node)e.appendChild(t);else if(t instanceof NodeList)for(const s of t)e.appendChild(s)},a=(e,t,s)=>{const n=\"string\"==typeof e?document.createElementNS(\"http://www.w3.org/1999/xhtml\",e):e instanceof Node?e:document.createDocumentFragment();if(!(\"string\"==typeof t||t instanceof Array||t instanceof NodeList||t instanceof Node)&&(\"object\"!=typeof s||s instanceof Array||s instanceof Node||s instanceof NodeList)||([t,s]=[s,t]),\"object\"==typeof t&&n instanceof Element)for(const[e,s]of Object.entries(t))if(s instanceof Function){const t={};let i=e;e:for(;;){switch(i.charAt(0)){case\"1\":t.once=!0;break;case\"C\":t.capture=!0;break;case\"P\":t.passive=!0;break;default:break e}i=i.slice(1)}i.startsWith(\"on\")&&n.addEventListener(i.substr(2),s,t)}else if(\"class\"===e&&(s instanceof Array||s instanceof DOMTokenList)&&s.length>0)n.classList.add(...s);else if(\"style\"===e&&\"object\"==typeof s&&(n instanceof HTMLElement||n instanceof SVGElement))for(const e in s)void 0===s[e]?n.style.removeProperty(e):n.style.setProperty(e,s[e]);else\"string\"==typeof s||\"number\"==typeof s?n.setAttribute(e,s):\"boolean\"==typeof s?n.toggleAttribute(e,s):void 0===s&&n.hasAttribute(e)&&n.removeAttribute(e);return\"string\"==typeof s?n.textContent=s:s&&(s instanceof Array||s instanceof Node||s instanceof NodeList)&&c(n,s),n},h=e=>{for(;null!==e.lastChild;)e.removeChild(e.lastChild);return e},[d,p,f,u,g,m,w,x,$,y,b,C,v,L,k,N]=\"a button datalist div h2 h3 input label li option span table tbody td tr ul\".split(\" \").map((e=>a.bind(null,e))),S=["
	jsMid     = "],A=["
	jsEnd     = "],G=[[\"Parent\",\"Father\",\"Mother\"],[\"Sibling\",\"Brother\",\"Sister\"],[\"Spouse\",\"Husband\",\"Wife\"],[\"Child\",\"Son\",\"Daughter\"],[\"Pibling\",\"Uncle\",\"Aunt\"],[\"Nibling\",\"Nephew\",\"Neice\"]],P=(e,t)=>{const s=[];for(;t>0;)s.push(t),t=e.get(t)[0];return s.reverse()},T={one:\"st\",two:\"nd\",few:\"rd\",other:\"th\"},j=[\"Once\",\"Twice\",\"Thrice\"],M=new Intl.PluralRules(\"en-GB\",{type:\"ordinal\"}),_=(e,t,s)=>{const n=e.length,i=t.length,o=n>0&&i>0&&S[e[n-1]][5]!=S[t[i-1]][5]?\"Half-\":\"\";switch(n){case 0:switch(i){case 0:return\"Clone\";case 1:return G[0][s];default:const e=i-2;return e>3?`${e} x Great-Grand-`:\"Great-\".repeat(e)+`Grand-${G[0][s]}`}case 1:switch(i){case 0:return G[3][S[e[0]][4]];case 1:return`${o}${G[1][S[e[0]][4]]}`;default:const t=i-2;return`${o}${t>3?`${t} x Great-Grand-`:\"Great-\".repeat(t)}${G[4][S[e[0]][4]]}`}break;default:const t=n-2;switch(i){case 0:return(t>3?`${t} x Great-Grand-`:\"Great-\".repeat(t)+\"Grand-\")+G[3][S[e[0]][4]];case 1:return`${o}${t>3?`${t} x Great-`:\"Great-\".repeat(t)}${G[5][S[e[0]][4]]}`;default:const s=Math.min(n,i)-1,r=Math.abs(n-i);return`${o}${s}${T[M.select(s)]} ${o}Cousin${r>0?`, ${j[r-1]||`${r} Times`} Removed`:\"\"}`}}},E=e=>`${S[e][0]??\"?\"} ${S[e][1]??\"?\"}`,I=Array.from({length:26},(()=>[])),O=(new Intl.Collator).compare,R=(e,t)=>{const[s=\"\",n=\"\"]=S[e],[i=\"\",o=\"\"]=S[t];return n!==o?O(n,o):s!==i?O(s,i):t-e},U=(e,t)=>0===e?[]:u([a(te(\"tree\",{id:e}),E(e)),\" (\"+G[t][S[e][4]]+\")\"]),F=(e,t,s,n,i)=>{0!==e.length&&e.push(\"…\");for(let o=s;o<=n;o++)o!==s&&e.push(\", \"),e.push(t===o?b(o+1+\"\"):a(te(\"list\",Object.assign({p:o},i)),{class:\"pagination_link\"},o+1+\"\"))},D=(e,t,s=0)=>{const n=Math.ceil(e.length/20)-1,i=[];if(0===n)return[];s>n&&(s=n);let o=0;for(let e=0;e<=n;e++)e<3||e>n-3||(3>s||e>=s-3)&&e<=s+3||s-3-1==3&&3==e||s+3+1==n-3&&e==n-3||(e!=o&&F(i,s,o,e-1,t),o=e+1);return o<n&&F(i,s,o,n,t),u({class:\"pagination\"},[\"Pages: \",a(0!==s?te(\"list\",Object.assign({p:s-1},t)):b(),{class:\"pagination_link prev\"},\"Previous\"),i,a(s!==n?te(\"list\",Object.assign({p:s+1},t)):b(),{class:\"pagination_link next\"},\"Next\")])},q=(e,t,s,n=0)=>{const i=Math.min(20*(n+1),t.length),o=N({class:\"results\"}),l=[];for(let e=20*n;e<i;e++){const s=t[e],[,,,,,n,...i]=S[s],[c,h,...d]=A[n],f=i.map((e=>A[e])),g=p({onclick:()=>{if(r===s){r=0;for(const e of l)e.innerText=\"+\"}else if(0===r){r=s;for(const e of l)e.innerText=\"=\";g.innerText=\"-\"}else ee(\"fhcalc\",{from:r,to:s})}},0===r?\"+\":r===s?\"-\":\"=\");l.push(g),o.appendChild($([u([a(te(\"tree\",{id:s}),E(s)),g]),u([U(c,0),U(h,0),d.filter((e=>e!==s)).map((e=>U(e,1))),f.map((([e,t,...n])=>[U(e!==s?e:t,2),n.filter((e=>e!==s)).map((e=>U(e,3)))]))])]))}a(e,[D(t,s,n),o,D(t,s,n)])},B=new Map,H=()=>ee(\"list\",{q:W.value}),W=w({type:\"text\",list:\"treeNames\",onkeypress:e=>\"Enter\"===e.key&&H()}),V=({l:e,q:t,d_p:s=0})=>{const n=u(),o=Math.max(0,\"string\"==typeof s?parseInt(s)||0:s);if(se(\"List\"),\"string\"==typeof t){W.value=t,se(\"Search\");const e=W.value.toUpperCase().split(\" \").sort(),s=e.join(\" \");let i=[];if(B.has(s))i=B.get(s);else{for(let t=0;t<S.length;t++){const s=`${S[t][0]||\"\"} ${S[t][1]||\"\"}`.toUpperCase();e.every((e=>s.includes(e)))&&i.push(t)}B.set(s,i)}q(n,i.sort(R),{q:t},o)}else if(\"string\"==typeof e){const t=e.toUpperCase().charCodeAt(0);t>=65&&t<=90&&(se(`List - ${e}`),q(n,I[t-65],{l:e},o))}return[i||(i=u({id:\"ged2web_title\"},[g(\"Select a Name\"),u({id:\"indexes\"},I.map(((e,t)=>a(te(\"list\",{l:String.fromCharCode(t+65)}),String.fromCharCode(t+65))))),u({id:\"index_search\"},[f({id:\"treeNames\"},S.map((([e=\"\",t=\"\"])=>e&&t?y({value:`${e} ${t}`}):[]))),x({for:\"index_search\"},\"Search Terms: \"),W,p({onclick:H},\"Search\")])])),n]},X=100,z=50,J=150,K=150,Q=50,Y=K+Q,Z=[\"U\",\"M\",\"F\"].map((e=>`person sex_${e}`)),ee=(t,s,n=!1)=>{let i,r=\"list\";switch(t){case\"tree\":i=(({id:t,highlight:s})=>{const n=\"string\"==typeof t?parseInt(t):t;if(!(n<=0||void 0===S[n]))return se(`Family Tree - ${E(n)}`),new e(n,(s||\"\").split(\".\").map((e=>parseInt(e))).filter((e=>e>0))).container})(s),r=t;break;case\"fhcalc\":i=(({from:e,to:t})=>{const s=\"string\"==typeof e?parseInt(e):e,n=\"string\"==typeof t?parseInt(t):t;if(s<=0||n<=0||!S[s]||!S[n])return;se(\"Relationship Calculator\");const[i,o,r]=((e,t)=>{const s=new Map([[e,[0,0]],[t,[0,0]]]),n=[[e,1],[t,2]];for(;n.length>0;){const e=n.shift(),[t,i]=e,[,,,,,o]=S[t];for(const r of A[o].slice(0,2))if(0!==r){const o=s.get(r);if(o){if(o[1]!==i)return[r,P(s,1===i?t:o[0]),P(s,1===i?o[0]:t)]}else s.set(r,e),n.push([r,i])}}return[0,[],[]]})(s,n),l=E(s),c=E(n);return 0===i?g(`No direct relationship betwen ${l} and ${c}`):[u({id:\"ged2web_title\"},[g(`${l} is the ${_(o,r,S[i][4])} of ${c}`),p({onclick:()=>ee(\"fhcalc\",{from:n,to:s})},\"Swap\")]),C({id:\"relationship\"},v([k([l,c].map((e=>L(m(`Route from ${e}`))))),k([o,r].map((e=>L(N(e.map((e=>$(`${E(e)}, who is the ${G[3][S[e][4]]} of…`)))))))),k(L({colspan:2},[u(E(i)),m(\"Common Ancestor\"),a(te(\"tree\",{id:i,highlight:o.concat(r).join(\".\")}),\"Show in Tree\")]))]))]})(s),r=t;break;case\"list\":i=V(s)}n||history.pushState(null,\"\",oe(t,s)),l?document.body.classList.replace(l,l=\"ged2web_\"+r):document.body.classList.add(l=\"ged2web_\"+r),a(h(o),i||V({}))},te=(e,t)=>d({href:oe(e,t),onclick:s=>{s.preventDefault(),ee(e,t)}}),se=e=>document.title=`${le} - ${e}`,ne=()=>window.location.pathname.split(\"/\").pop()?.split(\".\").shift(),ie=[\"list\",\"fhcalc\",\"tree\"].includes(ne()),oe=(e,t)=>(ie?`${e}.html?`:`?module=${e}&`)+Object.entries(t).map((([e,t])=>`${e}=${encodeURIComponent(t)}`)).join(\"&\"),re=()=>{const e=Object.fromEntries(new URL(window.location+\"\").searchParams.entries());ee(ie?ne():e.module,e,!0)},le=document.title;for(let e=0;e<S.length;e++){let t=(S[e][1]??\"\").charCodeAt(0);t>=97&&(t-=32),t>=65&&t<=90&&I[t-65].push(e)}for(const e of I)e.sort(R);window.addEventListener(\"popstate\",re),(\"complete\"==document.readyState?Promise.resolve():new Promise((e=>globalThis.addEventListener(\"load\",e,{once:!0})))).then((()=>{o=document.getElementById(\"ged2web\")||document.body,re()}));"
)
