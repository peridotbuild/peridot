const P=!0,p=(e,r)=>{const t=r instanceof URLSearchParams;return r&&!t&&(r=new URLSearchParams(r)),r?`${e}?${r}`:e},o={VITE_API:"",VITE_MODE:"development",BASE_URL:"/",MODE:"production",DEV:!1,PROD:!0}.VITE_PUBLIC_PATH||"";function l(e){return e.replace(/%/g,"%25").replace(/,/g,"%2C").replace(/\//g,"%2F").replace(/\\/g,"%5C").replace(/\?/g,"%3F").replace(/:/g,"%3A").replace(/@/g,"%40").replace(/&/g,"%26").replace(/=/g,"%3D").replace(/\+/g,"%2B").replace(/\$/g,"%24").replace(/#/g,"%23")}function d(e){return e.replace(/%2C/g,",").replace(/%2F/g,"/").replace(/%5C/g,"\\").replace(/%3F/g,"?").replace(/%3A/g,":").replace(/%40/g,"@").replace(/%26/g,"&").replace(/%3D/g,"=").replace(/%2B/g,"+").replace(/%24/g,"$").replace(/%23/g,"#").replace(/%25/g,"%")}const F=e=>e==="feed"||e==="compact"||e==="json",m=()=>`${o}/namespaces`,c=({namespace:e})=>`${o}/namespaces/${e}`,w=()=>`${o}/select-namespace`,f=e=>`${c(e)}/workflows`,U=({namespace:e,query:r,search:t})=>p(f({namespace:e}),{query:r,search:t}),k=e=>`${c(e)}/archival`,a=({workflow:e,run:r,...t})=>{const s=l(e);return`${f(t)}/${s}/${r}`},h=e=>`${c(e)}/schedules`,S=({namespace:e})=>`${h({namespace:e})}/new`,E=({scheduleId:e,namespace:r})=>{const t=l(e);return`${h({namespace:r})}/${t}`},R=({view:e,queryParams:r,...t})=>{const s=`${a(t)}/history`;return!e||!F(e)?`${s}/feed`:p(`${s}/${e}`,r)},L=e=>`${a(e)}/workers`,A=e=>{const r=l(e.queue);return`${c({namespace:e.namespace})}/task-queues/${r}`},I=e=>`${a(e)}/stack-trace`,v=e=>`${a(e)}/query`,W=e=>`${a(e)}/pending-activities`,C=e=>{var i;const{settings:r,searchParams:t,originUrl:s}=e,n=new URL(`${o}/auth/sso`,r.baseUrl);let u=(i=r.auth.options)!=null?i:[];return u=[...u,"returnUrl"],u.forEach(g=>{const $=t.get(g);$&&n.searchParams.set(g,$)}),!n.searchParams.get("returnUrl")&&s&&n.searchParams.set("returnUrl",s),n.toString()},D=(e="",r=!0)=>{if(r){const t=new URL(`${o}/login`,window.location.origin);return t.searchParams.set("returnUrl",window.location.href),e&&t.searchParams.set("error",e),t.toString()}return`${o}/login`},b=({importType:e,view:r})=>e==="events"&&r?`${o}/import/${e}/namespace/workflow/run/history/${r}`:`${o}/import/${e}`;export{k as a,P as b,m as c,h as d,D as e,b as f,R as g,L as h,W as i,I as j,v as k,U as l,d as m,a as n,A as o,o as p,w as q,f as r,C as s,p as t,E as u,S as v,c as w};
