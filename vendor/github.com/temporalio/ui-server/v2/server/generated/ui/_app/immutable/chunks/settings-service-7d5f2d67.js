import{s as d}from"./settings-c94c37e6.js";import{r as m,a as f,g as E}from"./route-for-api-652ea8c7.js";function A(){var o;return(o={VITE_API:"",VITE_MODE:"development",BASE_URL:"/",MODE:"production",DEV:!1,PROD:!0}.VITE_TEMPORAL_UI_BUILD_TARGET)!=null?o:null}const u=/(tmprl\.cloud|tmprl-test\.cloud)$/,O=async(o=fetch)=>{var c,i,l,e;const n=await m("settings"),a=await f(n,{request:o}),r=A(),t={auth:{enabled:!!((c=a==null?void 0:a.Auth)!=null&&c.Enabled),options:(i=a==null?void 0:a.Auth)==null?void 0:i.Options},baseUrl:E(),codec:{endpoint:(l=a==null?void 0:a.Codec)==null?void 0:l.Endpoint,passAccessToken:(e=a==null?void 0:a.Codec)==null?void 0:e.PassAccessToken},defaultNamespace:(a==null?void 0:a.DefaultNamespace)||"default",disableWriteActions:!!(a!=null&&a.DisableWriteActions)||!1,showTemporalSystemNamespace:a==null?void 0:a.ShowTemporalSystemNamespace,notifyOnNewVersion:a==null?void 0:a.NotifyOnNewVersion,feedbackURL:a==null?void 0:a.FeedbackURL,runtimeEnvironment:{get isCloud(){return r?r==="cloud":u.test(window.location.hostname)},get isLocal(){return r?r==="local":u.test(window.location.hostname)},envOverride:Boolean(r)},version:a==null?void 0:a.Version};return d.set(t),t};export{O as f,u as i};
