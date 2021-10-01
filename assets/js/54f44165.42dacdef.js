"use strict";(self.webpackChunk=self.webpackChunk||[]).push([[152],{3905:function(e,t,r){r.d(t,{Zo:function(){return p},kt:function(){return g}});var n=r(7294);function o(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function a(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,n)}return r}function i(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?a(Object(r),!0).forEach((function(t){o(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):a(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}function c(e,t){if(null==e)return{};var r,n,o=function(e,t){if(null==e)return{};var r,n,o={},a=Object.keys(e);for(n=0;n<a.length;n++)r=a[n],t.indexOf(r)>=0||(o[r]=e[r]);return o}(e,t);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);for(n=0;n<a.length;n++)r=a[n],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(o[r]=e[r])}return o}var u=n.createContext({}),l=function(e){var t=n.useContext(u),r=t;return e&&(r="function"==typeof e?e(t):i(i({},t),e)),r},p=function(e){var t=l(e.components);return n.createElement(u.Provider,{value:t},e.children)},s={inlineCode:"code",wrapper:function(e){var t=e.children;return n.createElement(n.Fragment,{},t)}},d=n.forwardRef((function(e,t){var r=e.components,o=e.mdxType,a=e.originalType,u=e.parentName,p=c(e,["components","mdxType","originalType","parentName"]),d=l(r),g=o,f=d["".concat(u,".").concat(g)]||d[g]||s[g]||a;return r?n.createElement(f,i(i({ref:t},p),{},{components:r})):n.createElement(f,i({ref:t},p))}));function g(e,t){var r=arguments,o=t&&t.mdxType;if("string"==typeof e||o){var a=r.length,i=new Array(a);i[0]=d;var c={};for(var u in t)hasOwnProperty.call(t,u)&&(c[u]=t[u]);c.originalType=e,c.mdxType="string"==typeof e?e:o,i[1]=c;for(var l=2;l<a;l++)i[l]=r[l];return n.createElement.apply(null,i)}return n.createElement.apply(null,r)}d.displayName="MDXCreateElement"},681:function(e,t,r){r.r(t),r.d(t,{frontMatter:function(){return c},contentTitle:function(){return u},metadata:function(){return l},toc:function(){return p},default:function(){return d}});var n=r(7462),o=r(3366),a=(r(7294),r(3905)),i=["components"],c={title:"\u5b89\u88c5"},u=void 0,l={unversionedId:"getting-started/installation",id:"getting-started/installation",isDocsHomePage:!1,title:"\u5b89\u88c5",description:"Docker",source:"@site/docs/getting-started/installation.md",sourceDirName:"getting-started",slug:"/getting-started/installation",permalink:"/opengauss_exporter/getting-started/installation",editUrl:"https://github.com/bzp2010/opengauss_exporter/edit/main/docs/docs/getting-started/installation.md",tags:[],version:"current",frontMatter:{title:"\u5b89\u88c5"},sidebar:"tutorialSidebar",previous:{title:"5\u5206\u949f\u4f53\u9a8c",permalink:"/opengauss_exporter/getting-started/5min-tutorial"},next:{title:"\u8f6f\u4ef6\u67b6\u6784",permalink:"/opengauss_exporter/concepts/architecture"}},p=[{value:"Docker",id:"docker",children:[]},{value:"Source",id:"source",children:[]}],s={toc:p};function d(e){var t=e.components,r=(0,o.Z)(e,i);return(0,a.kt)("wrapper",(0,n.Z)({},s,r,{components:t,mdxType:"MDXLayout"}),(0,a.kt)("h2",{id:"docker"},"Docker"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-shell"},"$ docker pull bzp2010/opengauss_exporter:v1.0.0\n$ docker run -d bzp2010/opengauss_exporter:v1.0.0\n")),(0,a.kt)("h2",{id:"source"},"Source"),(0,a.kt)("p",null,"required: Golang 1.14 +"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-shell"},"$ git clone https://github.com/bzp2010/opengauss_exporter.git\n$ cd opengauss_exporter\n$ go build -o opengauss_exporter main.go\n$ ./opengauss_exporter\n")))}d.isMDXComponent=!0}}]);