import{p}from"./paginated-a5f6e575.js";import{r as f,a as u}from"./route-for-api-652ea8c7.js";import{t as g}from"./index-ee84452b.js";import{b as m}from"./is-44021919.js";const v=e=>!m(e)||e==="descending"?"events.descending":e==="ascending"?"events.ascending":"events.descending",h=async({namespace:e,workflowId:n,runId:t,sort:s,onStart:r,onUpdate:o,onComplete:i})=>{const a=v(s),c=await f(a,{namespace:e,workflowId:n,runId:t});return(await p(async d=>u(c,{token:d,request:fetch}),{onStart:r,onUpdate:o,onComplete:i})).history.events},A=async e=>{const{settings:n,namespace:t,accessToken:s}=e;return h(e).then(r=>g({response:r,namespace:t,settings:n,accessToken:s}))};export{h as a,A as f};
