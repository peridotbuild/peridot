function b(){}const X=t=>t;function Tt(t,e){for(const n in e)t[n]=e[n];return t}function St(t){return t&&typeof t=="object"&&typeof t.then=="function"}function ft(t){return t()}function ot(){return Object.create(null)}function T(t){t.forEach(ft)}function O(t){return typeof t=="function"}function ie(t,e){return t!=t?e==e:t!==e||t&&typeof t=="object"||typeof t=="function"}let P;function se(t,e){return P||(P=document.createElement("a")),P.href=e,t===P.href}function jt(t){return Object.keys(t).length===0}function _t(t,...e){if(t==null)return b;const n=t.subscribe(...e);return n.unsubscribe?()=>n.unsubscribe():n}function re(t){let e;return _t(t,n=>e=n)(),e}function oe(t,e,n){t.$$.on_destroy.push(_t(e,n))}function ce(t,e,n,i){if(t){const s=dt(t,e,n,i);return t[0](s)}}function dt(t,e,n,i){return t[1]&&i?Tt(n.ctx.slice(),t[1](i(e))):n.ctx}function le(t,e,n,i){if(t[2]&&i){const s=t[2](i(n));if(e.dirty===void 0)return s;if(typeof s=="object"){const o=[],r=Math.max(e.dirty.length,s.length);for(let a=0;a<r;a+=1)o[a]=e.dirty[a]|s[a];return o}return e.dirty|s}return e.dirty}function ae(t,e,n,i,s,o){if(s){const r=dt(e,n,i,o);t.p(r,s)}}function ue(t){if(t.ctx.length>32){const e=[],n=t.ctx.length/32;for(let i=0;i<n;i++)e[i]=-1;return e}return-1}function fe(t){const e={};for(const n in t)n[0]!=="$"&&(e[n]=t[n]);return e}function _e(t,e){const n={};e=new Set(e);for(const i in t)!e.has(i)&&i[0]!=="$"&&(n[i]=t[i]);return n}function de(t){const e={};for(const n in t)e[n]=!0;return e}function he(t){return t==null?"":t}function me(t,e,n){return t.set(n),e}const Mt=(t,e)=>Object.prototype.hasOwnProperty.call(t,e);function pe(t){return t&&O(t.destroy)?t.destroy:b}const ht=typeof window<"u";let Y=ht?()=>window.performance.now():()=>Date.now(),Z=ht?t=>requestAnimationFrame(t):b;const A=new Set;function mt(t){A.forEach(e=>{e.c(t)||(A.delete(e),e.f())}),A.size!==0&&Z(mt)}function tt(t){let e;return A.size===0&&Z(mt),{promise:new Promise(n=>{A.add(e={c:t,f:n})}),abort(){A.delete(e)}}}let I=!1;function Dt(){I=!0}function Ot(){I=!1}function Ht(t,e,n,i){for(;t<e;){const s=t+(e-t>>1);n(s)<=i?t=s+1:e=s}return t}function Pt(t){if(t.hydrate_init)return;t.hydrate_init=!0;let e=t.childNodes;if(t.nodeName==="HEAD"){const c=[];for(let l=0;l<e.length;l++){const f=e[l];f.claim_order!==void 0&&c.push(f)}e=c}const n=new Int32Array(e.length+1),i=new Int32Array(e.length);n[0]=-1;let s=0;for(let c=0;c<e.length;c++){const l=e[c].claim_order,f=(s>0&&e[n[s]].claim_order<=l?s+1:Ht(1,s,_=>e[n[_]].claim_order,l))-1;i[c]=n[f]+1;const u=f+1;n[u]=c,s=Math.max(u,s)}const o=[],r=[];let a=e.length-1;for(let c=n[s]+1;c!=0;c=i[c-1]){for(o.push(e[c-1]);a>=c;a--)r.push(e[a]);a--}for(;a>=0;a--)r.push(e[a]);o.reverse(),r.sort((c,l)=>c.claim_order-l.claim_order);for(let c=0,l=0;c<r.length;c++){for(;l<o.length&&r[c].claim_order>=o[l].claim_order;)l++;const f=l<o.length?o[l]:null;t.insertBefore(r[c],f)}}function pt(t,e){t.appendChild(e)}function yt(t){if(!t)return document;const e=t.getRootNode?t.getRootNode():t.ownerDocument;return e&&e.host?e:t.ownerDocument}function Rt(t){const e=J("style");return zt(yt(t),e),e.sheet}function zt(t,e){return pt(t.head||t,e),e.sheet}function Bt(t,e){if(I){for(Pt(t),(t.actual_end_child===void 0||t.actual_end_child!==null&&t.actual_end_child.parentNode!==t)&&(t.actual_end_child=t.firstChild);t.actual_end_child!==null&&t.actual_end_child.claim_order===void 0;)t.actual_end_child=t.actual_end_child.nextSibling;e!==t.actual_end_child?(e.claim_order!==void 0||e.parentNode!==t)&&t.insertBefore(e,t.actual_end_child):t.actual_end_child=e.nextSibling}else(e.parentNode!==t||e.nextSibling!==null)&&t.appendChild(e)}function Lt(t,e,n){t.insertBefore(e,n||null)}function qt(t,e,n){I&&!n?Bt(t,e):(e.parentNode!==t||e.nextSibling!=n)&&t.insertBefore(e,n||null)}function N(t){t.parentNode.removeChild(t)}function ye(t,e){for(let n=0;n<t.length;n+=1)t[n]&&t[n].d(e)}function J(t){return document.createElement(t)}function ge(t,e){const n={};for(const i in t)Mt(t,i)&&e.indexOf(i)===-1&&(n[i]=t[i]);return n}function gt(t){return document.createElementNS("http://www.w3.org/2000/svg",t)}function et(t){return document.createTextNode(t)}function be(){return et(" ")}function we(){return et("")}function ct(t,e,n,i){return t.addEventListener(e,n,i),()=>t.removeEventListener(e,n,i)}function xe(t){return function(e){return e.preventDefault(),t.call(this,e)}}function $e(t){return function(e){return e.stopPropagation(),t.call(this,e)}}function bt(t,e,n){n==null?t.removeAttribute(e):t.getAttribute(e)!==n&&t.setAttribute(e,n)}function ve(t,e){const n=Object.getOwnPropertyDescriptors(t.__proto__);for(const i in e)e[i]==null?t.removeAttribute(i):i==="style"?t.style.cssText=e[i]:i==="__value"?t.value=t[i]=e[i]:n[i]&&n[i].set?t[i]=e[i]:bt(t,i,e[i])}function ke(t,e,n){e in t?t[e]=typeof t[e]=="boolean"&&n===""?!0:n:bt(t,e,n)}function Wt(t){return Array.from(t.childNodes)}function wt(t){t.claim_info===void 0&&(t.claim_info={last_index:0,total_claimed:0})}function xt(t,e,n,i,s=!1){wt(t);const o=(()=>{for(let r=t.claim_info.last_index;r<t.length;r++){const a=t[r];if(e(a)){const c=n(a);return c===void 0?t.splice(r,1):t[r]=c,s||(t.claim_info.last_index=r),a}}for(let r=t.claim_info.last_index-1;r>=0;r--){const a=t[r];if(e(a)){const c=n(a);return c===void 0?t.splice(r,1):t[r]=c,s?c===void 0&&t.claim_info.last_index--:t.claim_info.last_index=r,a}}return i()})();return o.claim_order=t.claim_info.total_claimed,t.claim_info.total_claimed+=1,o}function $t(t,e,n,i){return xt(t,s=>s.nodeName===e,s=>{const o=[];for(let r=0;r<s.attributes.length;r++){const a=s.attributes[r];n[a.name]||o.push(a.name)}o.forEach(r=>s.removeAttribute(r))},()=>i(e))}function Ee(t,e,n){return $t(t,e,n,J)}function Ce(t,e,n){return $t(t,e,n,gt)}function Ft(t,e){return xt(t,n=>n.nodeType===3,n=>{const i=""+e;if(n.data.startsWith(i)){if(n.data.length!==i.length)return n.splitText(i.length)}else n.data=i},()=>et(e),!0)}function Ae(t){return Ft(t," ")}function lt(t,e,n){for(let i=n;i<t.length;i+=1){const s=t[i];if(s.nodeType===8&&s.textContent.trim()===e)return i}return t.length}function Ne(t,e){const n=lt(t,"HTML_TAG_START",0),i=lt(t,"HTML_TAG_END",n);if(n===i)return new at(void 0,e);wt(t);const s=t.splice(n,i-n+1);N(s[0]),N(s[s.length-1]);const o=s.slice(1,s.length-1);for(const r of o)r.claim_order=t.claim_info.total_claimed,t.claim_info.total_claimed+=1;return new at(o,e)}function Te(t,e){e=""+e,t.wholeText!==e&&(t.data=e)}function Se(t,e){t.value=e==null?"":e}function je(t,e,n,i){n===null?t.style.removeProperty(e):t.style.setProperty(e,n,i?"important":"")}function Me(t,e){for(let n=0;n<t.options.length;n+=1){const i=t.options[n];if(i.__value===e){i.selected=!0;return}}t.selectedIndex=-1}function De(t){const e=t.querySelector(":checked")||t.options[0];return e&&e.__value}let R;function Gt(){if(R===void 0){R=!1;try{typeof window<"u"&&window.parent&&window.parent.document}catch{R=!0}}return R}function Oe(t,e){getComputedStyle(t).position==="static"&&(t.style.position="relative");const i=J("iframe");i.setAttribute("style","display: block; position: absolute; top: 0; left: 0; width: 100%; height: 100%; overflow: hidden; border: 0; opacity: 0; pointer-events: none; z-index: -1;"),i.setAttribute("aria-hidden","true"),i.tabIndex=-1;const s=Gt();let o;return s?(i.src="data:text/html,<script>onresize=function(){parent.postMessage(0,'*')}<\/script>",o=ct(window,"message",r=>{r.source===i.contentWindow&&e()})):(i.src="about:blank",i.onload=()=>{o=ct(i.contentWindow,"resize",e)}),pt(t,i),()=>{(s||o&&i.contentWindow)&&o(),N(i)}}function He(t,e,n){t.classList[n?"add":"remove"](e)}function vt(t,e,{bubbles:n=!1,cancelable:i=!1}={}){const s=document.createEvent("CustomEvent");return s.initCustomEvent(t,n,i,e),s}function Pe(t,e){const n=[];let i=0;for(const s of e.childNodes)if(s.nodeType===8){const o=s.textContent.trim();o===`HEAD_${t}_END`?(i-=1,n.push(s)):o===`HEAD_${t}_START`&&(i+=1,n.push(s))}else i>0&&n.push(s);return n}class It{constructor(e=!1){this.is_svg=!1,this.is_svg=e,this.e=this.n=null}c(e){this.h(e)}m(e,n,i=null){this.e||(this.is_svg?this.e=gt(n.nodeName):this.e=J(n.nodeName),this.t=n,this.c(e)),this.i(i)}h(e){this.e.innerHTML=e,this.n=Array.from(this.e.childNodes)}i(e){for(let n=0;n<this.n.length;n+=1)Lt(this.t,this.n[n],e)}p(e){this.d(),this.h(e),this.i(this.a)}d(){this.n.forEach(N)}}class at extends It{constructor(e,n=!1){super(n),this.e=this.n=null,this.l=e}c(e){this.l?this.n=this.l:super.c(e)}i(e){for(let n=0;n<this.n.length;n+=1)qt(this.t,this.n[n],e)}}function Re(t,e){return new t(e)}const q=new Map;let W=0;function Jt(t){let e=5381,n=t.length;for(;n--;)e=(e<<5)-e^t.charCodeAt(n);return e>>>0}function Kt(t,e){const n={stylesheet:Rt(e),rules:{}};return q.set(t,n),n}function nt(t,e,n,i,s,o,r,a=0){const c=16.666/i;let l=`{
`;for(let m=0;m<=1;m+=c){const g=e+(n-e)*o(m);l+=m*100+`%{${r(g,1-g)}}
`}const f=l+`100% {${r(n,1-n)}}
}`,u=`__svelte_${Jt(f)}_${a}`,_=yt(t),{stylesheet:d,rules:h}=q.get(_)||Kt(_,t);h[u]||(h[u]=!0,d.insertRule(`@keyframes ${u} ${f}`,d.cssRules.length));const y=t.style.animation||"";return t.style.animation=`${y?`${y}, `:""}${u} ${i}ms linear ${s}ms 1 both`,W+=1,u}function F(t,e){const n=(t.style.animation||"").split(", "),i=n.filter(e?o=>o.indexOf(e)<0:o=>o.indexOf("__svelte")===-1),s=n.length-i.length;s&&(t.style.animation=i.join(", "),W-=s,W||Qt())}function Qt(){Z(()=>{W||(q.forEach(t=>{const{ownerNode:e}=t.stylesheet;e&&N(e)}),q.clear())})}function ze(t,e,n,i){if(!e)return b;const s=t.getBoundingClientRect();if(e.left===s.left&&e.right===s.right&&e.top===s.top&&e.bottom===s.bottom)return b;const{delay:o=0,duration:r=300,easing:a=X,start:c=Y()+o,end:l=c+r,tick:f=b,css:u}=n(t,{from:e,to:s},i);let _=!0,d=!1,h;function y(){u&&(h=nt(t,0,1,r,o,a,u)),o||(d=!0)}function m(){u&&F(t,h),_=!1}return tt(g=>{if(!d&&g>=c&&(d=!0),d&&g>=l&&(f(1,0),m()),!_)return!1;if(d){const $=g-c,k=0+1*a($/r);f(k,1-k)}return!0}),y(),f(0,1),m}function Be(t){const e=getComputedStyle(t);if(e.position!=="absolute"&&e.position!=="fixed"){const{width:n,height:i}=e,s=t.getBoundingClientRect();t.style.position="absolute",t.style.width=n,t.style.height=i,Ut(t,s)}}function Ut(t,e){const n=t.getBoundingClientRect();if(e.left!==n.left||e.top!==n.top){const i=getComputedStyle(t),s=i.transform==="none"?"":i.transform;t.style.transform=`${s} translate(${e.left-n.left}px, ${e.top-n.top}px)`}}let M;function v(t){M=t}function C(){if(!M)throw new Error("Function called outside component initialization");return M}function Le(t){C().$$.on_mount.push(t)}function qe(t){C().$$.after_update.push(t)}function We(t){C().$$.on_destroy.push(t)}function Fe(){const t=C();return(e,n,{cancelable:i=!1}={})=>{const s=t.$$.callbacks[e];if(s){const o=vt(e,n,{cancelable:i});return s.slice().forEach(r=>{r.call(t,o)}),!o.defaultPrevented}return!0}}function Ge(t,e){return C().$$.context.set(t,e),e}function Ie(t){return C().$$.context.get(t)}function Je(t){return C().$$.context.has(t)}function Ke(t,e){const n=t.$$.callbacks[e.type];n&&n.slice().forEach(i=>i.call(this,e))}const j=[],ut=[],B=[],U=[],kt=Promise.resolve();let V=!1;function Et(){V||(V=!0,kt.then(it))}function Qe(){return Et(),kt}function D(t){B.push(t)}function Ue(t){U.push(t)}const Q=new Set;let z=0;function it(){const t=M;do{for(;z<j.length;){const e=j[z];z++,v(e),Vt(e.$$)}for(v(null),j.length=0,z=0;ut.length;)ut.pop()();for(let e=0;e<B.length;e+=1){const n=B[e];Q.has(n)||(Q.add(n),n())}B.length=0}while(j.length);for(;U.length;)U.pop()();V=!1,Q.clear(),v(t)}function Vt(t){if(t.fragment!==null){t.update(),T(t.before_update);const e=t.dirty;t.dirty=[-1],t.fragment&&t.fragment.p(t.ctx,e),t.after_update.forEach(D)}}let S;function Ct(){return S||(S=Promise.resolve(),S.then(()=>{S=null})),S}function G(t,e,n){t.dispatchEvent(vt(`${e?"intro":"outro"}${n}`))}const L=new Set;let E;function Xt(){E={r:0,c:[],p:E}}function Yt(){E.r||T(E.c),E=E.p}function st(t,e){t&&t.i&&(L.delete(t),t.i(e))}function At(t,e,n,i){if(t&&t.o){if(L.has(t))return;L.add(t),E.c.push(()=>{L.delete(t),i&&(n&&t.d(1),i())}),t.o(e)}else i&&i()}const Nt={duration:0};function Ve(t,e,n){let i=e(t,n),s=!1,o,r,a=0;function c(){o&&F(t,o)}function l(){const{delay:u=0,duration:_=300,easing:d=X,tick:h=b,css:y}=i||Nt;y&&(o=nt(t,0,1,_,u,d,y,a++)),h(0,1);const m=Y()+u,g=m+_;r&&r.abort(),s=!0,D(()=>G(t,!0,"start")),r=tt($=>{if(s){if($>=g)return h(1,0),G(t,!0,"end"),c(),s=!1;if($>=m){const k=d(($-m)/_);h(k,1-k)}}return s})}let f=!1;return{start(){f||(f=!0,F(t),O(i)?(i=i(),Ct().then(l)):l())},invalidate(){f=!1},end(){s&&(c(),s=!1)}}}function Xe(t,e,n){let i=e(t,n),s=!0,o;const r=E;r.r+=1;function a(){const{delay:c=0,duration:l=300,easing:f=X,tick:u=b,css:_}=i||Nt;_&&(o=nt(t,1,0,l,c,f,_));const d=Y()+c,h=d+l;D(()=>G(t,!1,"start")),tt(y=>{if(s){if(y>=h)return u(0,1),G(t,!1,"end"),--r.r||T(r.c),!1;if(y>=d){const m=f((y-d)/l);u(1-m,m)}}return s})}return O(i)?Ct().then(()=>{i=i(),a()}):a(),{end(c){c&&i.tick&&i.tick(1,0),s&&(o&&F(t,o),s=!1)}}}function Ye(t,e){const n=e.token={};function i(s,o,r,a){if(e.token!==n)return;e.resolved=a;let c=e.ctx;r!==void 0&&(c=c.slice(),c[r]=a);const l=s&&(e.current=s)(c);let f=!1;e.block&&(e.blocks?e.blocks.forEach((u,_)=>{_!==o&&u&&(Xt(),At(u,1,1,()=>{e.blocks[_]===u&&(e.blocks[_]=null)}),Yt())}):e.block.d(1),l.c(),st(l,1),l.m(e.mount(),e.anchor),f=!0),e.block=l,e.blocks&&(e.blocks[o]=l),f&&it()}if(St(t)){const s=C();if(t.then(o=>{v(s),i(e.then,1,e.value,o),v(null)},o=>{if(v(s),i(e.catch,2,e.error,o),v(null),!e.hasCatch)throw o}),e.current!==e.pending)return i(e.pending,0),!0}else{if(e.current!==e.then)return i(e.then,1,e.value,t),!0;e.resolved=t}}function Ze(t,e,n){const i=e.slice(),{resolved:s}=t;t.current===t.then&&(i[t.value]=s),t.current===t.catch&&(i[t.error]=s),t.block.p(i,n)}const tn=typeof window<"u"?window:typeof globalThis<"u"?globalThis:global;function en(t,e){t.d(1),e.delete(t.key)}function Zt(t,e){At(t,1,1,()=>{e.delete(t.key)})}function nn(t,e){t.f(),Zt(t,e)}function sn(t,e,n,i,s,o,r,a,c,l,f,u){let _=t.length,d=o.length,h=_;const y={};for(;h--;)y[t[h].key]=h;const m=[],g=new Map,$=new Map;for(h=d;h--;){const p=u(s,o,h),w=n(p);let x=r.get(w);x?i&&x.p(p,e):(x=l(w,p),x.c()),g.set(w,m[h]=x),w in y&&$.set(w,Math.abs(h-y[w]))}const k=new Set,rt=new Set;function K(p){st(p,1),p.m(a,f),r.set(p.key,p),f=p.first,d--}for(;_&&d;){const p=m[d-1],w=t[_-1],x=p.key,H=w.key;p===w?(f=p.first,_--,d--):g.has(H)?!r.has(x)||k.has(x)?K(p):rt.has(H)?_--:$.get(x)>$.get(H)?(rt.add(x),K(p)):(k.add(H),_--):(c(w,r),_--)}for(;_--;){const p=t[_];g.has(p.key)||c(p,r)}for(;d;)K(m[d-1]);return m}function rn(t,e){const n={},i={},s={$$scope:1};let o=t.length;for(;o--;){const r=t[o],a=e[o];if(a){for(const c in r)c in a||(i[c]=1);for(const c in a)s[c]||(n[c]=a[c],s[c]=1);t[o]=a}else for(const c in r)s[c]=1}for(const r in i)r in n||(n[r]=void 0);return n}function on(t){return typeof t=="object"&&t!==null?t:{}}function cn(t,e,n){const i=t.$$.props[e];i!==void 0&&(t.$$.bound[i]=n,n(t.$$.ctx[i]))}function ln(t){t&&t.c()}function an(t,e){t&&t.l(e)}function te(t,e,n,i){const{fragment:s,after_update:o}=t.$$;s&&s.m(e,n),i||D(()=>{const r=t.$$.on_mount.map(ft).filter(O);t.$$.on_destroy?t.$$.on_destroy.push(...r):T(r),t.$$.on_mount=[]}),o.forEach(D)}function ee(t,e){const n=t.$$;n.fragment!==null&&(T(n.on_destroy),n.fragment&&n.fragment.d(e),n.on_destroy=n.fragment=null,n.ctx=[])}function ne(t,e){t.$$.dirty[0]===-1&&(j.push(t),Et(),t.$$.dirty.fill(0)),t.$$.dirty[e/31|0]|=1<<e%31}function un(t,e,n,i,s,o,r,a=[-1]){const c=M;v(t);const l=t.$$={fragment:null,ctx:[],props:o,update:b,not_equal:s,bound:ot(),on_mount:[],on_destroy:[],on_disconnect:[],before_update:[],after_update:[],context:new Map(e.context||(c?c.$$.context:[])),callbacks:ot(),dirty:a,skip_bound:!1,root:e.target||c.$$.root};r&&r(l.root);let f=!1;if(l.ctx=n?n(t,e.props||{},(u,_,...d)=>{const h=d.length?d[0]:_;return l.ctx&&s(l.ctx[u],l.ctx[u]=h)&&(!l.skip_bound&&l.bound[u]&&l.bound[u](h),f&&ne(t,u)),_}):[],l.update(),f=!0,T(l.before_update),l.fragment=i?i(l.ctx):!1,e.target){if(e.hydrate){Dt();const u=Wt(e.target);l.fragment&&l.fragment.l(u),u.forEach(N)}else l.fragment&&l.fragment.c();e.intro&&st(t.$$.fragment),te(t,e.target,e.anchor,e.customElement),Ot(),it()}v(c)}class fn{$destroy(){ee(this,1),this.$destroy=b}$on(e,n){if(!O(n))return b;const i=this.$$.callbacks[e]||(this.$$.callbacks[e]=[]);return i.push(n),()=>{const s=i.indexOf(n);s!==-1&&i.splice(s,1)}}$set(e){this.$$set&&!jt(e)&&(this.$$.skip_bound=!0,this.$$set(e),this.$$.skip_bound=!1)}}export{Ze as $,rn as A,on as B,ee as C,Tt as D,Qe as E,b as F,_t as G,T as H,O as I,ce as J,ae as K,ue as L,le as M,Bt as N,oe as O,Fe as P,Ie as Q,ve as R,fn as S,He as T,ct as U,$e as V,fe as W,Ke as X,Ye as Y,Se as Z,pe as _,be as a,ut as a0,ye as a1,he as a2,_e as a3,me as a4,tn as a5,se as a6,re as a7,cn as a8,Ue as a9,gt as aa,Ce as ab,sn as ac,Be as ad,Ut as ae,ze as af,D as ag,Ve as ah,Xe as ai,nn as aj,xe as ak,de as al,Pe as am,We as an,Je as ao,Zt as ap,ge as aq,Oe as ar,en as as,Me as at,De as au,at as av,Ne as aw,ke as ax,qt as b,Ae as c,Yt as d,we as e,st as f,Xt as g,N as h,un as i,Ge as j,qe as k,J as l,Ee as m,Wt as n,Le as o,bt as p,je as q,et as r,ie as s,At as t,Ft as u,Te as v,Re as w,ln as x,an as y,te as z};
