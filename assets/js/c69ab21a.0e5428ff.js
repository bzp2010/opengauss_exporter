"use strict";(self.webpackChunk=self.webpackChunk||[]).push([[89],{3905:function(t,e,a){a.d(e,{Zo:function(){return p},kt:function(){return d}});var n=a(7294);function l(t,e,a){return e in t?Object.defineProperty(t,e,{value:a,enumerable:!0,configurable:!0,writable:!0}):t[e]=a,t}function i(t,e){var a=Object.keys(t);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(t);e&&(n=n.filter((function(e){return Object.getOwnPropertyDescriptor(t,e).enumerable}))),a.push.apply(a,n)}return a}function _(t){for(var e=1;e<arguments.length;e++){var a=null!=arguments[e]?arguments[e]:{};e%2?i(Object(a),!0).forEach((function(e){l(t,e,a[e])})):Object.getOwnPropertyDescriptors?Object.defineProperties(t,Object.getOwnPropertyDescriptors(a)):i(Object(a)).forEach((function(e){Object.defineProperty(t,e,Object.getOwnPropertyDescriptor(a,e))}))}return t}function r(t,e){if(null==t)return{};var a,n,l=function(t,e){if(null==t)return{};var a,n,l={},i=Object.keys(t);for(n=0;n<i.length;n++)a=i[n],e.indexOf(a)>=0||(l[a]=t[a]);return l}(t,e);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(t);for(n=0;n<i.length;n++)a=i[n],e.indexOf(a)>=0||Object.prototype.propertyIsEnumerable.call(t,a)&&(l[a]=t[a])}return l}var s=n.createContext({}),o=function(t){var e=n.useContext(s),a=e;return t&&(a="function"==typeof t?t(e):_(_({},e),t)),a},p=function(t){var e=o(t.components);return n.createElement(s.Provider,{value:e},t.children)},m={inlineCode:"code",wrapper:function(t){var e=t.children;return n.createElement(n.Fragment,{},e)}},c=n.forwardRef((function(t,e){var a=t.components,l=t.mdxType,i=t.originalType,s=t.parentName,p=r(t,["components","mdxType","originalType","parentName"]),c=o(a),d=l,u=c["".concat(s,".").concat(d)]||c[d]||m[d]||i;return a?n.createElement(u,_(_({ref:e},p),{},{components:a})):n.createElement(u,_({ref:e},p))}));function d(t,e){var a=arguments,l=e&&e.mdxType;if("string"==typeof t||l){var i=a.length,_=new Array(i);_[0]=c;var r={};for(var s in e)hasOwnProperty.call(e,s)&&(r[s]=e[s]);r.originalType=t,r.mdxType="string"==typeof t?t:l,_[1]=r;for(var o=2;o<i;o++)_[o]=a[o];return n.createElement.apply(null,_)}return n.createElement.apply(null,a)}c.displayName="MDXCreateElement"},7928:function(t,e,a){a.r(e),a.d(e,{frontMatter:function(){return r},contentTitle:function(){return s},metadata:function(){return o},toc:function(){return p},default:function(){return c}});var n=a(7462),l=a(3366),i=(a(7294),a(3905)),_=["components"],r={title:"\u522e\u524a\u5668"},s=void 0,o={unversionedId:"concepts/scrapers",id:"concepts/scrapers",isDocsHomePage:!1,title:"\u522e\u524a\u5668",description:"\u522e\u524a\u5668\u5217\u8868",source:"@site/docs/concepts/scrapers.md",sourceDirName:"concepts",slug:"/concepts/scrapers",permalink:"/opengauss_exporter/concepts/scrapers",editUrl:"https://github.com/bzp2010/opengauss_exporter/edit/main/docs/docs/concepts/scrapers.md",tags:[],version:"current",frontMatter:{title:"\u522e\u524a\u5668"},sidebar:"tutorialSidebar",previous:{title:"\u8f6f\u4ef6\u67b6\u6784",permalink:"/opengauss_exporter/concepts/architecture"},next:{title:"\u914d\u7f6e\u6587\u4ef6",permalink:"/opengauss_exporter/reference/config-file"}},p=[{value:"\u522e\u524a\u5668\u5217\u8868",id:"\u522e\u524a\u5668\u5217\u8868",children:[{value:"\u6838\u5fc3",id:"\u6838\u5fc3",children:[]},{value:"PostgreSQL Exporter",id:"postgresql_exporter",children:[]},{value:"PG_SETTING",id:"pg_setting",children:[]},{value:"GS_OS_RUN_INFO",id:"gs_os_run_info",children:[]},{value:"GS_SQL_COUNT",id:"gs_sql_count",children:[]},{value:"GS_INSTANCE_TIME",id:"gs_instance_time",children:[]},{value:"GS_TOTAL_MEMORY_DETAIL",id:"gs_total_memory_detail",children:[]}]}],m={toc:p};function c(t){var e=t.components,a=(0,l.Z)(t,_);return(0,i.kt)("wrapper",(0,n.Z)({},m,a,{components:e,mdxType:"MDXLayout"}),(0,i.kt)("h2",{id:"\u522e\u524a\u5668\u5217\u8868"},"\u522e\u524a\u5668\u5217\u8868"),(0,i.kt)("h3",{id:"\u6838\u5fc3"},"\u6838\u5fc3"),(0,i.kt)("h4",{id:"og_up"},"og_up"),(0,i.kt)("p",null,"\u6570\u636e\u5e93\u5b9e\u4f8b\u662f\u5426\u6b63\u5e38\u8fd0\u884c\u3002"),(0,i.kt)("h4",{id:"og_version"},"og_version"),(0,i.kt)("p",null,"\u6570\u636e\u5e93\u7248\u672c\u53f7\uff0c\u5728 label \u4e2d\u8fdb\u884c\u6807\u8bb0\uff0c\u4f8b\u5982"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-text"},'{og="2.1.0",pg="9.2.4",server="xxxx"}\n')),(0,i.kt)("h3",{id:"postgresql_exporter"},"PostgreSQL Exporter"),(0,i.kt)("p",null,"\u8fd9\u4e2a\u522e\u524a\u5668\u5b9e\u73b0\u4e86 PostgreSQL \u539f\u7248 Exporter \u7684\u522e\u524a\u529f\u80fd\uff0c\u5b83\u4ece ",(0,i.kt)("inlineCode",{parentName:"p"},"pg_")," \u7684\u4e00\u4e9b\u7cfb\u7edf\u89c6\u56fe\u4e2d\u83b7\u53d6\u8fd0\u884c\u6570\u636e\u3002"),(0,i.kt)("h4",{id:"pg_stat_bgwriter"},"pg_stat_bgwriter"),(0,i.kt)("p",null,"\u8be5\u89c6\u56fe\u663e\u793a\u5173\u4e8e\u540e\u7aef\u5199\u8fdb\u7a0b\u6d3b\u52a8\u7684\u7edf\u8ba1\u4fe1\u606f\u3002",(0,i.kt)("br",{parentName:"p"}),"\n","\u7ef4\u5ea6\uff1a\u6570\u636e\u5e93\u8282\u70b9"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},"pg_stat_bgwriter_buffers_alloc"),(0,i.kt)("li",{parentName:"ul"},"pg_stat_bgwriter_buffers_backend"),(0,i.kt)("li",{parentName:"ul"},"pg_stat_bgwriter_buffers_backend_fsync"),(0,i.kt)("li",{parentName:"ul"},"pg_stat_bgwriter_buffers_checkpoint"),(0,i.kt)("li",{parentName:"ul"},"pg_stat_bgwriter_buffers_clean"),(0,i.kt)("li",{parentName:"ul"},"pg_stat_bgwriter_checkpoint_sync_time"),(0,i.kt)("li",{parentName:"ul"},"pg_stat_bgwriter_checkpoint_write_time"),(0,i.kt)("li",{parentName:"ul"},"pg_stat_bgwriter_checkpoints_req"),(0,i.kt)("li",{parentName:"ul"},"pg_stat_bgwriter_checkpoints_timed"),(0,i.kt)("li",{parentName:"ul"},"pg_stat_bgwriter_maxwritten_clean"),(0,i.kt)("li",{parentName:"ul"},"pg_stat_bgwriter_stats_reset")),(0,i.kt)("div",{className:"admonition admonition-note alert alert--secondary"},(0,i.kt)("div",{parentName:"div",className:"admonition-heading"},(0,i.kt)("h5",{parentName:"div"},(0,i.kt)("span",{parentName:"h5",className:"admonition-icon"},(0,i.kt)("svg",{parentName:"span",xmlns:"http://www.w3.org/2000/svg",width:"14",height:"16",viewBox:"0 0 14 16"},(0,i.kt)("path",{parentName:"svg",fillRule:"evenodd",d:"M6.3 5.69a.942.942 0 0 1-.28-.7c0-.28.09-.52.28-.7.19-.18.42-.28.7-.28.28 0 .52.09.7.28.18.19.28.42.28.7 0 .28-.09.52-.28.7a1 1 0 0 1-.7.3c-.28 0-.52-.11-.7-.3zM8 7.99c-.02-.25-.11-.48-.31-.69-.2-.19-.42-.3-.69-.31H6c-.27.02-.48.13-.69.31-.2.2-.3.44-.31.69h1v3c.02.27.11.5.31.69.2.2.42.31.69.31h1c.27 0 .48-.11.69-.31.2-.19.3-.42.31-.69H8V7.98v.01zM7 2.3c-3.14 0-5.7 2.54-5.7 5.68 0 3.14 2.56 5.7 5.7 5.7s5.7-2.55 5.7-5.7c0-3.15-2.56-5.69-5.7-5.69v.01zM7 .98c3.86 0 7 3.14 7 7s-3.14 7-7 7-7-3.12-7-7 3.14-7 7-7z"}))),"\u53c2\u8003\u6587\u6863")),(0,i.kt)("div",{parentName:"div",className:"admonition-content"},(0,i.kt)("p",{parentName:"div"},(0,i.kt)("a",{parentName:"p",href:"https://www.opengauss.org/zh/docs/2.0.1/docs/Developerguide/PG_STAT_BGWRITER.html"},"pg_stat_bgwriter")))),(0,i.kt)("h4",{id:"pg_stat_database"},"pg_stat_database"),(0,i.kt)("p",null,"\u8be5\u89c6\u56fe\u5c06\u5305\u542bopenGauss\u4e2d\u6bcf\u4e2a\u6570\u636e\u5e93\u7684\u6570\u636e\u5e93\u7edf\u8ba1\u4fe1\u606f\u3002",(0,i.kt)("br",{parentName:"p"}),"\n","\u7ef4\u5ea6\uff1a\u6570\u636e\u5e93\u8282\u70b9 - \u6570\u636e\u5e93"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},"pg_stat_database_blk_read_time"),(0,i.kt)("li",{parentName:"ul"},"pg_stat_database_blk_write_time"),(0,i.kt)("li",{parentName:"ul"},"pg_stat_database_blks_hit"),(0,i.kt)("li",{parentName:"ul"},"pg_stat_database_blks_read"),(0,i.kt)("li",{parentName:"ul"},"pg_stat_database_conflicts"),(0,i.kt)("li",{parentName:"ul"},"pg_stat_database_conflicts_confl_bufferpin"),(0,i.kt)("li",{parentName:"ul"},"pg_stat_database_conflicts_confl_deadlock"),(0,i.kt)("li",{parentName:"ul"},"pg_stat_database_conflicts_confl_lock"),(0,i.kt)("li",{parentName:"ul"},"pg_stat_database_conflicts_confl_snapshot"),(0,i.kt)("li",{parentName:"ul"},"pg_stat_database_conflicts_confl_tablespace"),(0,i.kt)("li",{parentName:"ul"},"pg_stat_database_deadlocks"),(0,i.kt)("li",{parentName:"ul"},"pg_stat_database_numbackends"),(0,i.kt)("li",{parentName:"ul"},"pg_stat_database_stats_reset"),(0,i.kt)("li",{parentName:"ul"},"pg_stat_database_temp_bytes"),(0,i.kt)("li",{parentName:"ul"},"pg_stat_database_temp_files"),(0,i.kt)("li",{parentName:"ul"},"pg_stat_database_tup_deleted"),(0,i.kt)("li",{parentName:"ul"},"pg_stat_database_tup_fetched"),(0,i.kt)("li",{parentName:"ul"},"pg_stat_database_tup_inserted"),(0,i.kt)("li",{parentName:"ul"},"pg_stat_database_tup_returned"),(0,i.kt)("li",{parentName:"ul"},"pg_stat_database_tup_updated"),(0,i.kt)("li",{parentName:"ul"},"pg_stat_database_xact_commit"),(0,i.kt)("li",{parentName:"ul"},"pg_stat_database_xact_rollback")),(0,i.kt)("div",{className:"admonition admonition-note alert alert--secondary"},(0,i.kt)("div",{parentName:"div",className:"admonition-heading"},(0,i.kt)("h5",{parentName:"div"},(0,i.kt)("span",{parentName:"h5",className:"admonition-icon"},(0,i.kt)("svg",{parentName:"span",xmlns:"http://www.w3.org/2000/svg",width:"14",height:"16",viewBox:"0 0 14 16"},(0,i.kt)("path",{parentName:"svg",fillRule:"evenodd",d:"M6.3 5.69a.942.942 0 0 1-.28-.7c0-.28.09-.52.28-.7.19-.18.42-.28.7-.28.28 0 .52.09.7.28.18.19.28.42.28.7 0 .28-.09.52-.28.7a1 1 0 0 1-.7.3c-.28 0-.52-.11-.7-.3zM8 7.99c-.02-.25-.11-.48-.31-.69-.2-.19-.42-.3-.69-.31H6c-.27.02-.48.13-.69.31-.2.2-.3.44-.31.69h1v3c.02.27.11.5.31.69.2.2.42.31.69.31h1c.27 0 .48-.11.69-.31.2-.19.3-.42.31-.69H8V7.98v.01zM7 2.3c-3.14 0-5.7 2.54-5.7 5.68 0 3.14 2.56 5.7 5.7 5.7s5.7-2.55 5.7-5.7c0-3.15-2.56-5.69-5.7-5.69v.01zM7 .98c3.86 0 7 3.14 7 7s-3.14 7-7 7-7-3.12-7-7 3.14-7 7-7z"}))),"\u53c2\u8003\u6587\u6863")),(0,i.kt)("div",{parentName:"div",className:"admonition-content"},(0,i.kt)("p",{parentName:"div"},(0,i.kt)("a",{parentName:"p",href:"https://www.opengauss.org/zh/docs/2.0.1/docs/Developerguide/PG_STAT_DATABASE.html"},"pg_stat_database")))),(0,i.kt)("h4",{id:"pg_stat_database_conflicts"},"pg_stat_database_conflicts"),(0,i.kt)("p",null,"\u8be5\u89c6\u56fe\u663e\u793a\u6570\u636e\u5e93\u51b2\u7a81\u72b6\u6001\u7684\u7edf\u8ba1\u4fe1\u606f\u3002",(0,i.kt)("br",{parentName:"p"}),"\n","\u7ef4\u5ea6\uff1a\u6570\u636e\u5e93\u8282\u70b9 - \u6570\u636e\u5e93"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},"pg_stat_database_conflicts"),(0,i.kt)("li",{parentName:"ul"},"pg_stat_database_conflicts_confl_bufferpin"),(0,i.kt)("li",{parentName:"ul"},"pg_stat_database_conflicts_confl_deadlock"),(0,i.kt)("li",{parentName:"ul"},"pg_stat_database_conflicts_confl_lock"),(0,i.kt)("li",{parentName:"ul"},"pg_stat_database_conflicts_confl_snapshot"),(0,i.kt)("li",{parentName:"ul"},"pg_stat_database_conflicts_confl_tablespace")),(0,i.kt)("div",{className:"admonition admonition-note alert alert--secondary"},(0,i.kt)("div",{parentName:"div",className:"admonition-heading"},(0,i.kt)("h5",{parentName:"div"},(0,i.kt)("span",{parentName:"h5",className:"admonition-icon"},(0,i.kt)("svg",{parentName:"span",xmlns:"http://www.w3.org/2000/svg",width:"14",height:"16",viewBox:"0 0 14 16"},(0,i.kt)("path",{parentName:"svg",fillRule:"evenodd",d:"M6.3 5.69a.942.942 0 0 1-.28-.7c0-.28.09-.52.28-.7.19-.18.42-.28.7-.28.28 0 .52.09.7.28.18.19.28.42.28.7 0 .28-.09.52-.28.7a1 1 0 0 1-.7.3c-.28 0-.52-.11-.7-.3zM8 7.99c-.02-.25-.11-.48-.31-.69-.2-.19-.42-.3-.69-.31H6c-.27.02-.48.13-.69.31-.2.2-.3.44-.31.69h1v3c.02.27.11.5.31.69.2.2.42.31.69.31h1c.27 0 .48-.11.69-.31.2-.19.3-.42.31-.69H8V7.98v.01zM7 2.3c-3.14 0-5.7 2.54-5.7 5.68 0 3.14 2.56 5.7 5.7 5.7s5.7-2.55 5.7-5.7c0-3.15-2.56-5.69-5.7-5.69v.01zM7 .98c3.86 0 7 3.14 7 7s-3.14 7-7 7-7-3.12-7-7 3.14-7 7-7z"}))),"\u53c2\u8003\u6587\u6863")),(0,i.kt)("div",{parentName:"div",className:"admonition-content"},(0,i.kt)("p",{parentName:"div"},(0,i.kt)("a",{parentName:"p",href:"https://www.opengauss.org/zh/docs/2.0.1/docs/Developerguide/PG_STAT_DATABASE_CONFLICTS.html"},"pg_stat_database_conflicts")))),(0,i.kt)("h4",{id:"pg_locks"},"pg_locks"),(0,i.kt)("p",null,"\u8be5\u89c6\u56fe\u5b58\u50a8\u5404\u6253\u5f00\u4e8b\u52a1\u6240\u6301\u6709\u7684\u9501\u4fe1\u606f\u3002",(0,i.kt)("br",{parentName:"p"}),"\n","\u7ef4\u5ea6\uff1a\u6570\u636e\u5e93\u8282\u70b9 - \u6570\u636e\u5e93 - \u6307\u6807\u7c7b\u578b"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},"pg_locks_count",(0,i.kt)("ul",{parentName:"li"},(0,i.kt)("li",{parentName:"ul"},"accessexclusivelock"),(0,i.kt)("li",{parentName:"ul"},"accesssharelock"),(0,i.kt)("li",{parentName:"ul"},"exclusivelock"),(0,i.kt)("li",{parentName:"ul"},"rowexclusivelock"),(0,i.kt)("li",{parentName:"ul"},"rowsharelock"),(0,i.kt)("li",{parentName:"ul"},"sharelock"),(0,i.kt)("li",{parentName:"ul"},"sharerowexclusivelock"),(0,i.kt)("li",{parentName:"ul"},"shareupdateexclusivelock"),(0,i.kt)("li",{parentName:"ul"},"sireadlock")))),(0,i.kt)("div",{className:"admonition admonition-note alert alert--secondary"},(0,i.kt)("div",{parentName:"div",className:"admonition-heading"},(0,i.kt)("h5",{parentName:"div"},(0,i.kt)("span",{parentName:"h5",className:"admonition-icon"},(0,i.kt)("svg",{parentName:"span",xmlns:"http://www.w3.org/2000/svg",width:"14",height:"16",viewBox:"0 0 14 16"},(0,i.kt)("path",{parentName:"svg",fillRule:"evenodd",d:"M6.3 5.69a.942.942 0 0 1-.28-.7c0-.28.09-.52.28-.7.19-.18.42-.28.7-.28.28 0 .52.09.7.28.18.19.28.42.28.7 0 .28-.09.52-.28.7a1 1 0 0 1-.7.3c-.28 0-.52-.11-.7-.3zM8 7.99c-.02-.25-.11-.48-.31-.69-.2-.19-.42-.3-.69-.31H6c-.27.02-.48.13-.69.31-.2.2-.3.44-.31.69h1v3c.02.27.11.5.31.69.2.2.42.31.69.31h1c.27 0 .48-.11.69-.31.2-.19.3-.42.31-.69H8V7.98v.01zM7 2.3c-3.14 0-5.7 2.54-5.7 5.68 0 3.14 2.56 5.7 5.7 5.7s5.7-2.55 5.7-5.7c0-3.15-2.56-5.69-5.7-5.69v.01zM7 .98c3.86 0 7 3.14 7 7s-3.14 7-7 7-7-3.12-7-7 3.14-7 7-7z"}))),"\u53c2\u8003\u6587\u6863")),(0,i.kt)("div",{parentName:"div",className:"admonition-content"},(0,i.kt)("p",{parentName:"div"},(0,i.kt)("a",{parentName:"p",href:"https://www.opengauss.org/zh/docs/2.0.1/docs/Developerguide/PG_LOCKS.html"},"pg_locks")))),(0,i.kt)("h4",{id:"pg_stat_activity"},"pg_stat_activity"),(0,i.kt)("p",null,"\u8be5\u89c6\u56fe\u663e\u793a\u548c\u5f53\u524d\u7528\u6237\u67e5\u8be2\u76f8\u5173\u7684\u4fe1\u606f\u3002",(0,i.kt)("br",{parentName:"p"}),"\n","\u7ef4\u5ea6\uff1a\u6570\u636e\u5e93\u8282\u70b9 - \u6570\u636e\u5e93 - \u6307\u6807\u7c7b\u578b"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},"pg_stat_activity_count",(0,i.kt)("ul",{parentName:"li"},(0,i.kt)("li",{parentName:"ul"},"active"),(0,i.kt)("li",{parentName:"ul"},"disabled"),(0,i.kt)("li",{parentName:"ul"},"fastpath function call"),(0,i.kt)("li",{parentName:"ul"},"idle"),(0,i.kt)("li",{parentName:"ul"},"idle in transaction"),(0,i.kt)("li",{parentName:"ul"},"idle in transaction (aborted)")))),(0,i.kt)("div",{className:"admonition admonition-note alert alert--secondary"},(0,i.kt)("div",{parentName:"div",className:"admonition-heading"},(0,i.kt)("h5",{parentName:"div"},(0,i.kt)("span",{parentName:"h5",className:"admonition-icon"},(0,i.kt)("svg",{parentName:"span",xmlns:"http://www.w3.org/2000/svg",width:"14",height:"16",viewBox:"0 0 14 16"},(0,i.kt)("path",{parentName:"svg",fillRule:"evenodd",d:"M6.3 5.69a.942.942 0 0 1-.28-.7c0-.28.09-.52.28-.7.19-.18.42-.28.7-.28.28 0 .52.09.7.28.18.19.28.42.28.7 0 .28-.09.52-.28.7a1 1 0 0 1-.7.3c-.28 0-.52-.11-.7-.3zM8 7.99c-.02-.25-.11-.48-.31-.69-.2-.19-.42-.3-.69-.31H6c-.27.02-.48.13-.69.31-.2.2-.3.44-.31.69h1v3c.02.27.11.5.31.69.2.2.42.31.69.31h1c.27 0 .48-.11.69-.31.2-.19.3-.42.31-.69H8V7.98v.01zM7 2.3c-3.14 0-5.7 2.54-5.7 5.68 0 3.14 2.56 5.7 5.7 5.7s5.7-2.55 5.7-5.7c0-3.15-2.56-5.69-5.7-5.69v.01zM7 .98c3.86 0 7 3.14 7 7s-3.14 7-7 7-7-3.12-7-7 3.14-7 7-7z"}))),"\u53c2\u8003\u6587\u6863")),(0,i.kt)("div",{parentName:"div",className:"admonition-content"},(0,i.kt)("p",{parentName:"div"},(0,i.kt)("a",{parentName:"p",href:"https://www.opengauss.org/zh/docs/2.0.1/docs/Developerguide/PG_STAT_ACTIVITY.html"},"pg_stat_activity")))),(0,i.kt)("h4",{id:"pg_stat_replication"},"pg_stat_replication"),(0,i.kt)("p",null,"\u8be5\u89c6\u56fe\u7528\u4e8e\u63cf\u8ff0\u65e5\u5fd7\u540c\u6b65\u72b6\u6001\u4fe1\u606f\uff0c\u5982\u53d1\u8d77\u7aef\u53d1\u9001\u65e5\u5fd7\u4f4d\u7f6e\uff0c\u6536\u7aef\u63a5\u6536\u65e5\u5fd7\u4f4d\u7f6e\u7b49\u3002",(0,i.kt)("br",{parentName:"p"}),"\n",(0,i.kt)("strong",{parentName:"p"},"\u5907\u6ce8\uff1a\u5f53\u524d\u89c6\u56fe\u5728 openGauss \u4e2d\u4e0d\u53ef\u7528")),(0,i.kt)("h4",{id:"pg_replication_slots"},"pg_replication_slots"),(0,i.kt)("p",null,"\u8be5\u89c6\u56fe\u67e5\u770b\u590d\u5236\u8282\u70b9\u7684\u4fe1\u606f\u3002",(0,i.kt)("br",{parentName:"p"}),"\n",(0,i.kt)("strong",{parentName:"p"},"\u5907\u6ce8\uff1a\u5f53\u524d\u89c6\u56fe\u5728 openGauss \u4e2d\u4e0d\u53ef\u7528")),(0,i.kt)("h4",{id:"pg_stat_archiver"},"pg_stat_archiver"),(0,i.kt)("p",null,(0,i.kt)("strong",{parentName:"p"},"\u5907\u6ce8\uff1a\u5f53\u524d\u89c6\u56fe\u5728 openGauss \u4e2d\u4e0d\u53ef\u7528")),(0,i.kt)("h3",{id:"pg_setting"},"PG_SETTING"),(0,i.kt)("p",null,"\u8fd9\u4e2a\u522e\u524a\u5668\u5b9e\u73b0\u4e86\u4ece openGauss \u6570\u636e\u5e93\u4e2d\u83b7\u53d6\u6570\u636e\u5e93\u914d\u7f6e\u7684\u529f\u80fd\uff0c\u5b83\u4ece ",(0,i.kt)("inlineCode",{parentName:"p"},"pg_settings")," \u89c6\u56fe\u4e2d\u6293\u53d6\u5168\u90e8\u7684\u6570\u636e\u5e93\u914d\u7f6e\u3002"),(0,i.kt)("div",{className:"admonition admonition-note alert alert--secondary"},(0,i.kt)("div",{parentName:"div",className:"admonition-heading"},(0,i.kt)("h5",{parentName:"div"},(0,i.kt)("span",{parentName:"h5",className:"admonition-icon"},(0,i.kt)("svg",{parentName:"span",xmlns:"http://www.w3.org/2000/svg",width:"14",height:"16",viewBox:"0 0 14 16"},(0,i.kt)("path",{parentName:"svg",fillRule:"evenodd",d:"M6.3 5.69a.942.942 0 0 1-.28-.7c0-.28.09-.52.28-.7.19-.18.42-.28.7-.28.28 0 .52.09.7.28.18.19.28.42.28.7 0 .28-.09.52-.28.7a1 1 0 0 1-.7.3c-.28 0-.52-.11-.7-.3zM8 7.99c-.02-.25-.11-.48-.31-.69-.2-.19-.42-.3-.69-.31H6c-.27.02-.48.13-.69.31-.2.2-.3.44-.31.69h1v3c.02.27.11.5.31.69.2.2.42.31.69.31h1c.27 0 .48-.11.69-.31.2-.19.3-.42.31-.69H8V7.98v.01zM7 2.3c-3.14 0-5.7 2.54-5.7 5.68 0 3.14 2.56 5.7 5.7 5.7s5.7-2.55 5.7-5.7c0-3.15-2.56-5.69-5.7-5.69v.01zM7 .98c3.86 0 7 3.14 7 7s-3.14 7-7 7-7-3.12-7-7 3.14-7 7-7z"}))),"note")),(0,i.kt)("div",{parentName:"div",className:"admonition-content"},(0,i.kt)("p",{parentName:"div"},"\u6307\u6807\u6570\u91cf\u8fc7\u591a\uff0c\u6b64\u5904\u4e0d\u8fdb\u884c\u5217\u4e3e"))),(0,i.kt)("h3",{id:"gs_os_run_info"},"GS_OS_RUN_INFO"),(0,i.kt)("p",null,"\u8fd9\u4e2a\u522e\u524a\u5668\u4ece ",(0,i.kt)("inlineCode",{parentName:"p"},"GS_OS_RUN_INFO")," \u7cfb\u7edf\u89c6\u56fe\u4e2d\u83b7\u53d6\u6570\u636e\u5e93\u670d\u52a1\u5668\u8fd0\u884c\u60c5\u51b5\u3002",(0,i.kt)("br",{parentName:"p"}),"\n","\u7ef4\u5ea6\uff1a\u6570\u636e\u5e93\u8282\u70b9"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},"gs_os_run_info_avg_busy_time"),(0,i.kt)("li",{parentName:"ul"},"gs_os_run_info_avg_idle_time "),(0,i.kt)("li",{parentName:"ul"},"gs_os_run_info_avg_iowait_time"),(0,i.kt)("li",{parentName:"ul"},"gs_os_run_info_avg_nice_time"),(0,i.kt)("li",{parentName:"ul"},"gs_os_run_info_avg_sys_time"),(0,i.kt)("li",{parentName:"ul"},"gs_os_run_info_avg_user_time"),(0,i.kt)("li",{parentName:"ul"},"gs_os_run_info_busy_time"),(0,i.kt)("li",{parentName:"ul"},"gs_os_run_info_idle_time"),(0,i.kt)("li",{parentName:"ul"},"gs_os_run_info_iowait_time"),(0,i.kt)("li",{parentName:"ul"},"gs_os_run_info_load"),(0,i.kt)("li",{parentName:"ul"},"gs_os_run_info_nice_time"),(0,i.kt)("li",{parentName:"ul"},"gs_os_run_info_num_cpu_cores"),(0,i.kt)("li",{parentName:"ul"},"gs_os_run_info_num_cpu_sockets"),(0,i.kt)("li",{parentName:"ul"},"gs_os_run_info_num_cpus"),(0,i.kt)("li",{parentName:"ul"},"gs_os_run_info_physical_memory_bytes"),(0,i.kt)("li",{parentName:"ul"},"gs_os_run_info_sys_time"),(0,i.kt)("li",{parentName:"ul"},"gs_os_run_info_user_time"),(0,i.kt)("li",{parentName:"ul"},"gs_os_run_info_vm_page_in_bytes"),(0,i.kt)("li",{parentName:"ul"},"gs_os_run_info_vm_page_out_bytes")),(0,i.kt)("div",{className:"admonition admonition-note alert alert--secondary"},(0,i.kt)("div",{parentName:"div",className:"admonition-heading"},(0,i.kt)("h5",{parentName:"div"},(0,i.kt)("span",{parentName:"h5",className:"admonition-icon"},(0,i.kt)("svg",{parentName:"span",xmlns:"http://www.w3.org/2000/svg",width:"14",height:"16",viewBox:"0 0 14 16"},(0,i.kt)("path",{parentName:"svg",fillRule:"evenodd",d:"M6.3 5.69a.942.942 0 0 1-.28-.7c0-.28.09-.52.28-.7.19-.18.42-.28.7-.28.28 0 .52.09.7.28.18.19.28.42.28.7 0 .28-.09.52-.28.7a1 1 0 0 1-.7.3c-.28 0-.52-.11-.7-.3zM8 7.99c-.02-.25-.11-.48-.31-.69-.2-.19-.42-.3-.69-.31H6c-.27.02-.48.13-.69.31-.2.2-.3.44-.31.69h1v3c.02.27.11.5.31.69.2.2.42.31.69.31h1c.27 0 .48-.11.69-.31.2-.19.3-.42.31-.69H8V7.98v.01zM7 2.3c-3.14 0-5.7 2.54-5.7 5.68 0 3.14 2.56 5.7 5.7 5.7s5.7-2.55 5.7-5.7c0-3.15-2.56-5.69-5.7-5.69v.01zM7 .98c3.86 0 7 3.14 7 7s-3.14 7-7 7-7-3.12-7-7 3.14-7 7-7z"}))),"\u53c2\u8003\u6587\u6863")),(0,i.kt)("div",{parentName:"div",className:"admonition-content"},(0,i.kt)("p",{parentName:"div"},(0,i.kt)("a",{parentName:"p",href:"https://www.opengauss.org/zh/docs/2.0.1/docs/Developerguide/GS_OS_RUN_INFO.html"},"GS_OS_RUN_INFO")))),(0,i.kt)("h3",{id:"gs_sql_count"},"GS_SQL_COUNT"),(0,i.kt)("p",null,"\u8fd9\u4e2a\u522e\u524a\u5668\u4ece ",(0,i.kt)("inlineCode",{parentName:"p"},"GS_SQL_COUNT")," \u7cfb\u7edf\u89c6\u56fe\u4e2d\u83b7\u53d6\u6570\u636e\u5e93\u4e2d\u6267\u884c\u7684SQL\u8bed\u53e5\u8ba1\u6570\u60c5\u51b5\u3002",(0,i.kt)("br",{parentName:"p"}),"\n","\u7ef4\u5ea6\uff1a\u6570\u636e\u5e93\u8282\u70b9 - \u6570\u636e\u5e93"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},"gs_sql_count_select_count"),(0,i.kt)("li",{parentName:"ul"},"gs_sql_count_insert_count"),(0,i.kt)("li",{parentName:"ul"},"gs_sql_count_update_count"),(0,i.kt)("li",{parentName:"ul"},"gs_sql_count_delete_count"),(0,i.kt)("li",{parentName:"ul"},"gs_sql_count_dcl_count"),(0,i.kt)("li",{parentName:"ul"},"gs_sql_count_ddl_count"),(0,i.kt)("li",{parentName:"ul"},"gs_sql_count_dml_count"),(0,i.kt)("li",{parentName:"ul"},"gs_sql_count_mergeinto_count"),(0,i.kt)("li",{parentName:"ul"},"gs_sql_count_avg_delete_elapse"),(0,i.kt)("li",{parentName:"ul"},"gs_sql_count_avg_insert_elapse"),(0,i.kt)("li",{parentName:"ul"},"gs_sql_count_avg_select_elapse"),(0,i.kt)("li",{parentName:"ul"},"gs_sql_count_avg_update_elapse"),(0,i.kt)("li",{parentName:"ul"},"gs_sql_count_total_delete_elapse"),(0,i.kt)("li",{parentName:"ul"},"gs_sql_count_total_insert_elapse"),(0,i.kt)("li",{parentName:"ul"},"gs_sql_count_total_select_elapse"),(0,i.kt)("li",{parentName:"ul"},"gs_sql_count_total_update_elapse"),(0,i.kt)("li",{parentName:"ul"},"gs_sql_count_max_delete_elapse"),(0,i.kt)("li",{parentName:"ul"},"gs_sql_count_max_insert_elapse"),(0,i.kt)("li",{parentName:"ul"},"gs_sql_count_max_select_elapse"),(0,i.kt)("li",{parentName:"ul"},"gs_sql_count_max_update_elapse"),(0,i.kt)("li",{parentName:"ul"},"gs_sql_count_min_delete_elapse"),(0,i.kt)("li",{parentName:"ul"},"gs_sql_count_min_insert_elapse"),(0,i.kt)("li",{parentName:"ul"},"gs_sql_count_min_select_elapse"),(0,i.kt)("li",{parentName:"ul"},"gs_sql_count_min_update_elapse")),(0,i.kt)("div",{className:"admonition admonition-note alert alert--secondary"},(0,i.kt)("div",{parentName:"div",className:"admonition-heading"},(0,i.kt)("h5",{parentName:"div"},(0,i.kt)("span",{parentName:"h5",className:"admonition-icon"},(0,i.kt)("svg",{parentName:"span",xmlns:"http://www.w3.org/2000/svg",width:"14",height:"16",viewBox:"0 0 14 16"},(0,i.kt)("path",{parentName:"svg",fillRule:"evenodd",d:"M6.3 5.69a.942.942 0 0 1-.28-.7c0-.28.09-.52.28-.7.19-.18.42-.28.7-.28.28 0 .52.09.7.28.18.19.28.42.28.7 0 .28-.09.52-.28.7a1 1 0 0 1-.7.3c-.28 0-.52-.11-.7-.3zM8 7.99c-.02-.25-.11-.48-.31-.69-.2-.19-.42-.3-.69-.31H6c-.27.02-.48.13-.69.31-.2.2-.3.44-.31.69h1v3c.02.27.11.5.31.69.2.2.42.31.69.31h1c.27 0 .48-.11.69-.31.2-.19.3-.42.31-.69H8V7.98v.01zM7 2.3c-3.14 0-5.7 2.54-5.7 5.68 0 3.14 2.56 5.7 5.7 5.7s5.7-2.55 5.7-5.7c0-3.15-2.56-5.69-5.7-5.69v.01zM7 .98c3.86 0 7 3.14 7 7s-3.14 7-7 7-7-3.12-7-7 3.14-7 7-7z"}))),"\u53c2\u8003\u6587\u6863")),(0,i.kt)("div",{parentName:"div",className:"admonition-content"},(0,i.kt)("p",{parentName:"div"},(0,i.kt)("a",{parentName:"p",href:"https://www.opengauss.org/zh/docs/2.0.1/docs/Developerguide/GS_SQL_COUNT.html"},"GS_SQL_COUNT")))),(0,i.kt)("h3",{id:"gs_instance_time"},"GS_INSTANCE_TIME"),(0,i.kt)("p",null,"\u8fd9\u4e2a\u522e\u524a\u5668\u4ece ",(0,i.kt)("inlineCode",{parentName:"p"},"GS_INSTANCE_TIME")," \u7cfb\u7edf\u89c6\u56fe\u4e2d\u83b7\u53d6\u5f53\u524d\u6570\u636e\u5e93\u5b9e\u4f8b\u4e2a\u9636\u6bb5\u8fd0\u884c\u8017\u65f6\u72b6\u6001\u3002",(0,i.kt)("br",{parentName:"p"}),"\n","\u7ef4\u5ea6\uff1a\u6570\u636e\u5e93\u8282\u70b9 - \u6570\u636e\u5e93"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},"gs_instance_time_cpu_time"),(0,i.kt)("li",{parentName:"ul"},"gs_instance_time_data_io_time"),(0,i.kt)("li",{parentName:"ul"},"gs_instance_time_db_time"),(0,i.kt)("li",{parentName:"ul"},"gs_instance_time_execution_time"),(0,i.kt)("li",{parentName:"ul"},"gs_instance_time_net_send_time"),(0,i.kt)("li",{parentName:"ul"},"gs_instance_time_parse_time"),(0,i.kt)("li",{parentName:"ul"},"gs_instance_time_pl_compilation_time"),(0,i.kt)("li",{parentName:"ul"},"gs_instance_time_pl_execution_time"),(0,i.kt)("li",{parentName:"ul"},"gs_instance_time_plan_time"),(0,i.kt)("li",{parentName:"ul"},"gs_instance_time_rewrite_time")),(0,i.kt)("div",{className:"admonition admonition-note alert alert--secondary"},(0,i.kt)("div",{parentName:"div",className:"admonition-heading"},(0,i.kt)("h5",{parentName:"div"},(0,i.kt)("span",{parentName:"h5",className:"admonition-icon"},(0,i.kt)("svg",{parentName:"span",xmlns:"http://www.w3.org/2000/svg",width:"14",height:"16",viewBox:"0 0 14 16"},(0,i.kt)("path",{parentName:"svg",fillRule:"evenodd",d:"M6.3 5.69a.942.942 0 0 1-.28-.7c0-.28.09-.52.28-.7.19-.18.42-.28.7-.28.28 0 .52.09.7.28.18.19.28.42.28.7 0 .28-.09.52-.28.7a1 1 0 0 1-.7.3c-.28 0-.52-.11-.7-.3zM8 7.99c-.02-.25-.11-.48-.31-.69-.2-.19-.42-.3-.69-.31H6c-.27.02-.48.13-.69.31-.2.2-.3.44-.31.69h1v3c.02.27.11.5.31.69.2.2.42.31.69.31h1c.27 0 .48-.11.69-.31.2-.19.3-.42.31-.69H8V7.98v.01zM7 2.3c-3.14 0-5.7 2.54-5.7 5.68 0 3.14 2.56 5.7 5.7 5.7s5.7-2.55 5.7-5.7c0-3.15-2.56-5.69-5.7-5.69v.01zM7 .98c3.86 0 7 3.14 7 7s-3.14 7-7 7-7-3.12-7-7 3.14-7 7-7z"}))),"\u53c2\u8003\u6587\u6863")),(0,i.kt)("div",{parentName:"div",className:"admonition-content"},(0,i.kt)("p",{parentName:"div"},(0,i.kt)("a",{parentName:"p",href:"https://www.opengauss.org/zh/docs/2.0.1/docs/Developerguide/GS_INSTANCE_TIME.html"},"GS_INSTANCE_TIME")))),(0,i.kt)("h3",{id:"gs_total_memory_detail"},"GS_TOTAL_MEMORY_DETAIL"),(0,i.kt)("p",null,"\u8fd9\u4e2a\u522e\u524a\u5668\u4ece ",(0,i.kt)("inlineCode",{parentName:"p"},"GS_TOTAL_MEMORY_DETAIL")," \u7cfb\u7edf\u89c6\u56fe\u4e2d\u83b7\u53d6\u5f53\u524d\u6570\u636e\u5e93\u5b9e\u4f8b\u5185\u5b58\u4f7f\u7528\u72b6\u6001\u3002",(0,i.kt)("br",{parentName:"p"}),"\n","\u7ef4\u5ea6\uff1a\u6570\u636e\u5e93\u8282\u70b9 - \u6570\u636e\u5e93"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},"gs_total_memory_detail_backend_used_memory"),(0,i.kt)("li",{parentName:"ul"},"gs_total_memory_detail_cstore_used_memory"),(0,i.kt)("li",{parentName:"ul"},"gs_total_memory_detail_dynamic_peak_memory"),(0,i.kt)("li",{parentName:"ul"},"gs_total_memory_detail_dynamic_peak_shrctx"),(0,i.kt)("li",{parentName:"ul"},"gs_total_memory_detail_dynamic_used_memory"),(0,i.kt)("li",{parentName:"ul"},"gs_total_memory_detail_dynamic_used_shrctx"),(0,i.kt)("li",{parentName:"ul"},"gs_total_memory_detail_gpu_dynamic_peak_memory"),(0,i.kt)("li",{parentName:"ul"},"gs_total_memory_detail_gpu_dynamic_used_memory"),(0,i.kt)("li",{parentName:"ul"},"gs_total_memory_detail_gpu_max_dynamic_memory"),(0,i.kt)("li",{parentName:"ul"},"gs_total_memory_detail_max_backend_memory"),(0,i.kt)("li",{parentName:"ul"},"gs_total_memory_detail_max_cstore_memory"),(0,i.kt)("li",{parentName:"ul"},"gs_total_memory_detail_max_dynamic_memory"),(0,i.kt)("li",{parentName:"ul"},"gs_total_memory_detail_max_process_memory"),(0,i.kt)("li",{parentName:"ul"},"gs_total_memory_detail_max_sctpcomm_memory"),(0,i.kt)("li",{parentName:"ul"},"gs_total_memory_detail_max_shared_memory"),(0,i.kt)("li",{parentName:"ul"},"gs_total_memory_detail_other_used_memory"),(0,i.kt)("li",{parentName:"ul"},"gs_total_memory_detail_pooler_conn_memory"),(0,i.kt)("li",{parentName:"ul"},"gs_total_memory_detail_pooler_freeconn_memory"),(0,i.kt)("li",{parentName:"ul"},"gs_total_memory_detail_process_used_memory"),(0,i.kt)("li",{parentName:"ul"},"gs_total_memory_detail_sctpcomm_peak_memory"),(0,i.kt)("li",{parentName:"ul"},"gs_total_memory_detail_sctpcomm_used_memory"),(0,i.kt)("li",{parentName:"ul"},"gs_total_memory_detail_shared_used_memory"),(0,i.kt)("li",{parentName:"ul"},"gs_total_memory_detail_storage_compress_memory"),(0,i.kt)("li",{parentName:"ul"},"gs_total_memory_detail_udf_reserved_memory")),(0,i.kt)("div",{className:"admonition admonition-note alert alert--secondary"},(0,i.kt)("div",{parentName:"div",className:"admonition-heading"},(0,i.kt)("h5",{parentName:"div"},(0,i.kt)("span",{parentName:"h5",className:"admonition-icon"},(0,i.kt)("svg",{parentName:"span",xmlns:"http://www.w3.org/2000/svg",width:"14",height:"16",viewBox:"0 0 14 16"},(0,i.kt)("path",{parentName:"svg",fillRule:"evenodd",d:"M6.3 5.69a.942.942 0 0 1-.28-.7c0-.28.09-.52.28-.7.19-.18.42-.28.7-.28.28 0 .52.09.7.28.18.19.28.42.28.7 0 .28-.09.52-.28.7a1 1 0 0 1-.7.3c-.28 0-.52-.11-.7-.3zM8 7.99c-.02-.25-.11-.48-.31-.69-.2-.19-.42-.3-.69-.31H6c-.27.02-.48.13-.69.31-.2.2-.3.44-.31.69h1v3c.02.27.11.5.31.69.2.2.42.31.69.31h1c.27 0 .48-.11.69-.31.2-.19.3-.42.31-.69H8V7.98v.01zM7 2.3c-3.14 0-5.7 2.54-5.7 5.68 0 3.14 2.56 5.7 5.7 5.7s5.7-2.55 5.7-5.7c0-3.15-2.56-5.69-5.7-5.69v.01zM7 .98c3.86 0 7 3.14 7 7s-3.14 7-7 7-7-3.12-7-7 3.14-7 7-7z"}))),"\u53c2\u8003\u6587\u6863")),(0,i.kt)("div",{parentName:"div",className:"admonition-content"},(0,i.kt)("p",{parentName:"div"},(0,i.kt)("a",{parentName:"p",href:"https://www.opengauss.org/zh/docs/2.0.1/docs/Developerguide/GS_TOTAL_MEMORY_DETAIL.html"},"GS_TOTAL_MEMORY_DETAIL")))))}c.isMDXComponent=!0}}]);