import{d as r}from"./index-d8ed1bec.js";import{p as a}from"./stores-7f07ccfb.js";import{p as o}from"./persist-store-c25f6403.js";import{s as n}from"./settings-c94c37e6.js";import{c as i,i as p}from"./version-check-3a7b1db8.js";const f=r([i],([e])=>e==null?void 0:e.serverVersion),S=r([n],([e])=>e==null?void 0:e.version),V=o("autoRefreshWorkflow","off"),h=o("eventView","feed"),E=o("expandAllEvents","false"),d=o("eventFilterSort","descending"),g=o("eventShowElapsed","false"),x=r([a],([e])=>e.url.searchParams.get("category")),m=r([f,n],([e,s])=>s.runtimeEnvironment.isCloud?!0:p(e,"1.16.0")),y=r([d,m],([e,s])=>{let t="ascending";return s&&(t=e),t});export{V as a,d as b,y as c,x as d,h as e,g as f,E as g,m as s,f as t,S as u};
