"use strict";(self.webpackChunk=self.webpackChunk||[]).push([[355],{3905:function(e,r,t){t.d(r,{Zo:function(){return l},kt:function(){return m}});var n=t(7294);function o(e,r,t){return r in e?Object.defineProperty(e,r,{value:t,enumerable:!0,configurable:!0,writable:!0}):e[r]=t,e}function i(e,r){var t=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);r&&(n=n.filter((function(r){return Object.getOwnPropertyDescriptor(e,r).enumerable}))),t.push.apply(t,n)}return t}function c(e){for(var r=1;r<arguments.length;r++){var t=null!=arguments[r]?arguments[r]:{};r%2?i(Object(t),!0).forEach((function(r){o(e,r,t[r])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(t)):i(Object(t)).forEach((function(r){Object.defineProperty(e,r,Object.getOwnPropertyDescriptor(t,r))}))}return e}function a(e,r){if(null==e)return{};var t,n,o=function(e,r){if(null==e)return{};var t,n,o={},i=Object.keys(e);for(n=0;n<i.length;n++)t=i[n],r.indexOf(t)>=0||(o[t]=e[t]);return o}(e,r);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(n=0;n<i.length;n++)t=i[n],r.indexOf(t)>=0||Object.prototype.propertyIsEnumerable.call(e,t)&&(o[t]=e[t])}return o}var p=n.createContext({}),u=function(e){var r=n.useContext(p),t=r;return e&&(t="function"==typeof e?e(r):c(c({},r),e)),t},l=function(e){var r=u(e.components);return n.createElement(p.Provider,{value:r},e.children)},f={inlineCode:"code",wrapper:function(e){var r=e.children;return n.createElement(n.Fragment,{},r)}},s=n.forwardRef((function(e,r){var t=e.components,o=e.mdxType,i=e.originalType,p=e.parentName,l=a(e,["components","mdxType","originalType","parentName"]),s=u(t),m=o,d=s["".concat(p,".").concat(m)]||s[m]||f[m]||i;return t?n.createElement(d,c(c({ref:r},l),{},{components:t})):n.createElement(d,c({ref:r},l))}));function m(e,r){var t=arguments,o=r&&r.mdxType;if("string"==typeof e||o){var i=t.length,c=new Array(i);c[0]=s;var a={};for(var p in r)hasOwnProperty.call(r,p)&&(a[p]=r[p]);a.originalType=e,a.mdxType="string"==typeof e?e:o,c[1]=a;for(var u=2;u<i;u++)c[u]=t[u];return n.createElement.apply(null,c)}return n.createElement.apply(null,t)}s.displayName="MDXCreateElement"},4921:function(e,r,t){t.r(r),t.d(r,{frontMatter:function(){return a},contentTitle:function(){return p},metadata:function(){return u},toc:function(){return l},default:function(){return s}});var n=t(7462),o=t(3366),i=(t(7294),t(3905)),c=["components"],a={title:"HTTP \u63a5\u53e3"},p=void 0,u={unversionedId:"reference/http-api",id:"reference/http-api",isDocsHomePage:!1,title:"HTTP \u63a5\u53e3",description:"GET /metrics",source:"@site/docs/reference/http-api.md",sourceDirName:"reference",slug:"/reference/http-api",permalink:"/opengauss_exporter/en/reference/http-api",editUrl:"https://github.com/bzp2010/opengauss_exporter/edit/main/docs/docs/reference/http-api.md",tags:[],version:"current",frontMatter:{title:"HTTP \u63a5\u53e3"},sidebar:"tutorialSidebar",previous:{title:"\u914d\u7f6e\u6587\u4ef6",permalink:"/opengauss_exporter/en/reference/config-file"},next:{title:"\u547d\u4ee4\u884c",permalink:"/opengauss_exporter/en/reference/cli"}},l=[{value:"GET <code>/metrics</code>",id:"get-metrics",children:[]},{value:"GET <code>/refresh</code>",id:"get-refresh",children:[]}],f={toc:l};function s(e){var r=e.components,t=(0,o.Z)(e,c);return(0,i.kt)("wrapper",(0,n.Z)({},f,t,{components:r,mdxType:"MDXLayout"}),(0,i.kt)("h2",{id:"get-metrics"},"GET ",(0,i.kt)("inlineCode",{parentName:"h2"},"/metrics")),(0,i.kt)("p",null,"\u8f93\u51fa Prometheus \u683c\u5f0f\u7684\u6307\u6807\u6570\u636e"),(0,i.kt)("h2",{id:"get-refresh"},"GET ",(0,i.kt)("inlineCode",{parentName:"h2"},"/refresh")),(0,i.kt)("p",null,"\u6e05\u9664 Metrics \u5185\u5b58\u7f13\u5b58\u6570\u636e"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},"\u54cd\u5e94",(0,i.kt)("ul",{parentName:"li"},(0,i.kt)("li",{parentName:"ul"},"\u6210\u529f\uff1arefresh success"),(0,i.kt)("li",{parentName:"ul"},"\u5931\u8d25\uff1arefresh failed: {\u5931\u8d25\u539f\u56e0}")))))}s.isMDXComponent=!0}}]);