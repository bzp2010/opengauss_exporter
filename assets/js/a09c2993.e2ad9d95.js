"use strict";(self.webpackChunk=self.webpackChunk||[]).push([[128],{8495:function(e,t,r){r.r(t),r.d(t,{frontMatter:function(){return s},contentTitle:function(){return o},metadata:function(){return l},toc:function(){return p},default:function(){return m}});var n=r(7462),u=r(3366),i=(r(7294),r(3905)),a=["components"],s={title:"\u7b80\u4ecb",slug:"/"},o=void 0,l={unversionedId:"introduction",id:"introduction",isDocsHomePage:!1,title:"\u7b80\u4ecb",description:"openGauss \u6570\u636e\u5e93",source:"@site/docs/introduction.md",sourceDirName:".",slug:"/",permalink:"/opengauss_exporter/",editUrl:"https://github.com/bzp2010/opengauss_exporter/edit/main/docs/docs/introduction.md",tags:[],version:"current",frontMatter:{title:"\u7b80\u4ecb",slug:"/"},sidebar:"tutorialSidebar",next:{title:"5\u5206\u949f\u4f53\u9a8c",permalink:"/opengauss_exporter/getting-started/5min-tutorial"}},p=[{value:"openGauss \u6570\u636e\u5e93",id:"opengauss",children:[{value:"\u4e3b\u8981\u7279\u70b9",id:"opengauss-characteristic",children:[]}]},{value:"Prometheus",id:"prometheus",children:[{value:"Metric",id:"prometheus-metric",children:[]},{value:"Server",id:"server",children:[]},{value:"Exporter",id:"exporter",children:[]}]},{value:"openGauss Exporter",id:"opengauss-exporter",children:[]}],c={toc:p};function m(e){var t=e.components,r=(0,u.Z)(e,a);return(0,i.kt)("wrapper",(0,n.Z)({},c,r,{components:t,mdxType:"MDXLayout"}),(0,i.kt)("h2",{id:"opengauss"},"openGauss \u6570\u636e\u5e93"),(0,i.kt)("p",null,"openGauss\u662f\u4e00\u6b3e\u5f00\u6e90\u7684\u5173\u7cfb\u578b\u6570\u636e\u5e93,\u91c7\u7528\u5ba2\u6237\u7aef/\u670d\u52a1\u5668\uff0c\u5355\u8fdb\u7a0b\u591a\u7ebf\u7a0b\u67b6\u6784\uff0c\u652f\u6301\u5355\u673a\u548c\u4e00\u4e3b\u591a\u5907\u90e8\u7f72\u65b9\u5f0f\uff0c\u5907\u673a\u53ef\u8bfb\uff0c\u652f\u6301\u53cc\u673a\u9ad8\u53ef\u7528\u548c\u8bfb\u6269\u5c55\u3002"),(0,i.kt)("h3",{id:"opengauss-characteristic"},"\u4e3b\u8981\u7279\u70b9"),(0,i.kt)("h4",{id:"opengauss-characteristic-performance"},"\u9ad8\u6027\u80fd"),(0,i.kt)("p",null,"\u63d0\u4f9b\u4e86\u9762\u5411\u591a\u6838\u67b6\u6784\u7684\u5e76\u53d1\u63a7\u5236\u6280\u672f\u7ed3\u5408\u9cb2\u9e4f\u786c\u4ef6\u4f18\u5316\uff0c\u5728\u4e24\u8def\u9cb2\u9e4f\u4e0bTPCC Benchmark\u8fbe\u6210\u6027\u80fd150\u4e07tpmc\u3002\n\u9488\u5bf9\u5f53\u524d\u786c\u4ef6\u591a\u6838numa\u7684\u67b6\u6784\u8d8b\u52bf\uff0c \u5728\u5185\u6838\u5173\u952e\u7ed3\u6784\u4e0a\u91c7\u7528\u4e86Numa-Aware\u7684\u6570\u636e\u7ed3\u6784\u3002\n\u63d0\u4f9bSql-bypass\u667a\u80fd\u5feb\u901f\u5f15\u64ce\u6280\u672f\u3002"),(0,i.kt)("h4",{id:"opengauss-characteristic-availability"},"\u9ad8\u53ef\u7528"),(0,i.kt)("p",null,"\u652f\u6301\u4e3b\u5907\u540c\u6b65\uff0c\u5f02\u6b65\u4ee5\u53ca\u7ea7\u8054\u5907\u673a\u591a\u79cd\u90e8\u7f72\u6a21\u5f0f\u3002\n\u6570\u636e\u9875CRC\u6821\u9a8c\uff0c\u635f\u574f\u6570\u636e\u9875\u901a\u8fc7\u5907\u673a\u81ea\u52a8\u4fee\u590d\u3002\n\u5907\u673a\u5e76\u884c\u6062\u590d\uff0c10\u79d2\u5185\u53ef\u5347\u4e3b\u63d0\u4f9b\u670d\u52a1\u3002"),(0,i.kt)("h4",{id:"opengauss-characteristic-security"},"\u9ad8\u5b89\u5168"),(0,i.kt)("p",null,"\u652f\u6301\u5168\u5bc6\u6001\u8ba1\u7b97\uff0c\u8bbf\u95ee\u63a7\u5236\u3001\u52a0\u5bc6\u8ba4\u8bc1\u3001\u6570\u636e\u5e93\u5ba1\u8ba1\u3001\u52a8\u6001\u6570\u636e\u8131\u654f\u7b49\u5b89\u5168\u7279\u6027\uff0c\u63d0\u4f9b\u5168\u65b9\u4f4d\u7aef\u5230\u7aef\u7684\u6570\u636e\u5b89\u5168\u4fdd\u62a4\u3002"),(0,i.kt)("h4",{id:"opengauss-characteristic-maintenance"},"\u6613\u8fd0\u7ef4"),(0,i.kt)("p",null,"\u57fa\u4e8eAI\u7684\u667a\u80fd\u53c2\u6570\u8c03\u4f18\u548c\u7d22\u5f15\u63a8\u8350\uff0c\u63d0\u4f9bAI\u81ea\u52a8\u53c2\u6570\u63a8\u8350\u3002\n\u6162SQL\u8bca\u65ad\uff0c\u591a\u7ef4\u6027\u80fd\u81ea\u76d1\u63a7\u89c6\u56fe\uff0c\u5b9e\u65bd\u638c\u63a7\u7cfb\u7edf\u7684\u6027\u80fd\u8868\u73b0\u3002\n\u63d0\u4f9b\u5728\u7ebf\u81ea\u5b66\u4e60\u7684SQL\u65f6\u95f4\u9884\u6d4b\u3002"),(0,i.kt)("h4",{id:"opengauss-characteristic-open"},"\u5168\u5f00\u653e"),(0,i.kt)("p",null,"\u91c7\u7528\u6728\u5170\u5bbd\u677e\u8bb8\u53ef\u8bc1\u534f\u8bae\uff0c\u5141\u8bb8\u5bf9\u4ee3\u7801\u81ea\u7531\u4fee\u6539\uff0c\u4f7f\u7528\uff0c\u5f15\u7528\u3002\n\u6570\u636e\u5e93\u5185\u6838\u80fd\u529b\u5168\u5f00\u653e\u3002\n\u63d0\u4f9b\u4e30\u5bcc\u7684\u4f19\u4f34\u8ba4\u8bc1\uff0c\u57f9\u8bad\u4f53\u7cfb\u548c\u9ad8\u6821\u8bfe\u7a0b\u3002\nopenGauss\u76f8\u6bd4\u5176\u4ed6\u5f00\u6e90\u6570\u636e\u5e93\u4e3b\u8981\u6709\u591a\u5b58\u50a8\u6a21\u5f0f\uff0cNUMA\u5316\u5185\u6838\u7ed3\u6784\u548c\u9ad8\u53ef\u7528\u7b49\u4ea7\u54c1\u7279\u70b9\u3002"),(0,i.kt)("h2",{id:"prometheus"},"Prometheus"),(0,i.kt)("p",null,"Prometheus \u662f\u7531 Soundcloud \u4ee5\u5f00\u6e90\u8f6f\u4ef6\u7684\u5f62\u5f0f\u8fdb\u884c\u53d1\u5e03\u7684\u76d1\u63a7\u548c\u544a\u8b66\u8f6f\u4ef6\uff0c\u8bb8\u591a\u516c\u53f8\u548c\u7ec4\u7ec7\u90fd\u91c7\u7528\u4e86 Prometheus \u4f5c\u4e3a\u5176\u76d1\u63a7\u544a\u8b66\u5de5\u5177\u3002Prometheus \u4e8e 2016 \u5e74 5 \u6708\u52a0\u5165 CNCF \u57fa\u91d1\u4f1a\uff0c\u6210\u4e3a\u7ee7 Kubernetes \u4e4b\u540e\u7684\u7b2c\u4e8c\u4e2a CNCF \u6258\u7ba1\u9879\u76ee\u3002\nPrometheus \u4e3b\u8981\u5305\u62ec Server \u548c\u65f6\u5e8f\u6570\u636e\u5e93\u4e24\u90e8\u5206\uff0cServer\u8d1f\u8d23\u5bf9\u5916\u63d0\u4f9b\u67e5\u8be2\u3001\u5199\u5165API\uff0c\u65f6\u5e8f\u6570\u636e\u5e93\u8d1f\u8d23\u4ee5\u65f6\u95f4\u987a\u5e8f\u5b58\u50a8\u91c7\u96c6\u6765\u7684 Metric \u6307\u6807\u6570\u636e\u3002"),(0,i.kt)("h3",{id:"prometheus-metric"},"Metric"),(0,i.kt)("p",null,"Prometheus\u91c7\u96c6\u5230\u7684\u76d1\u63a7\u6570\u636e\u5747\u4ee5 Metric\uff08\u6307\u6807\uff09\u5f62\u5f0f\u4fdd\u5b58\u5728\u65f6\u5e8f\u6570\u636e\u5e93\u4e2d\uff08TSDB\uff09"),(0,i.kt)("h4",{id:"metric-\u6307\u6807\u683c\u5f0f"},"Metric \u6307\u6807\u683c\u5f0f"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-text"},"<metric name>{<label key>=<label value>, ...}\n")),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},"metric name: \u76d1\u63a7\u6307\u6807\u540d\u79f0"),(0,i.kt)("li",{parentName:"ul"},"label key = label value: \u6807\u7b7e\u952e\u503c",(0,i.kt)("br",{parentName:"li"}),"\u901a\u8fc7 label \u7cfb\u7edf\u53ef\u4ee5\u6807\u8bb0\u6570\u636e\uff08eg. server=127.0.0.1:5432\uff09")),(0,i.kt)("h4",{id:"metric-\u6307\u6807\u7c7b\u578b"},"Metric \u6307\u6807\u7c7b\u578b"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},"Counter \u8ba1\u6570\u503c\uff0c\u7d2f\u52a0\u6570\u636e"),(0,i.kt)("li",{parentName:"ul"},"Gauge \u8ba1\u91cf\u503c\uff0c\u53ef\u589e\u53ef\u51cf"),(0,i.kt)("li",{parentName:"ul"},"Histogram \u76f4\u65b9\u56fe"),(0,i.kt)("li",{parentName:"ul"},"Summary \u6458\u8981")),(0,i.kt)("h3",{id:"server"},"Server"),(0,i.kt)("p",null,"Prometheus Server \u76f4\u63a5\u4ece\u76d1\u63a7\u76ee\u6807\u4e2d\u6216\u8005\u95f4\u63a5\u901a\u8fc7\u63a8\u9001\u7f51\u5173\u6765\u62c9\u53d6\u76d1\u63a7\u6307\u6807\uff0c\u5e76\u5728\u672c\u5730\u5b58\u50a8\u6240\u6709\u6293\u53d6\u5230\u7684\u6837\u672c\u6570\u636e\u3002"),(0,i.kt)("h3",{id:"exporter"},"Exporter"),(0,i.kt)("p",null,"\u5728 Prometheus Server \u4e2d Pull \u6a21\u5f0f\u4f7f\u7528\u7684\u7ec4\u4ef6\u3002\u5b83\u4ece\u76ee\u6807\u670d\u52a1\u4e2d\u63d0\u53d6 Metric \u6307\u6807\u6570\u636e\uff0c\u5e76\u5c06\u5176\u7ec4\u88c5\u4e3a Prometheus \u53ef\u8bc6\u522b\u7684\u683c\u5f0f\uff0c\u901a\u8fc7\u5bf9\u5916\u5f00\u653e\u7684 API \u4e3a Prometheus \u522e\u524a\u5668\u91c7\u96c6\u5b58\u50a8\u3002"),(0,i.kt)("h2",{id:"opengauss-exporter"},"openGauss Exporter"),(0,i.kt)("p",null,"openGauss Exporter \u662f\u4e3a openGauss \u6570\u636e\u5e93\u5b9e\u73b0\u7684 Prometheus Exporter\uff0c\u5b83\u901a\u8fc7 SQL \u4ece\u6570\u636e\u5e93\u8282\u70b9\u4e2d\u8bfb\u53d6\u8fd0\u884c\u73af\u5883\u6307\u6807\u7b49\u4fe1\u606f\uff0c\u5e76\u7ec4\u88c5\u6210 Prometheus Metric \u683c\u5f0f\uff0c\u5411\u5916\u63d0\u4f9b\u91c7\u96c6\u63a5\u53e3\u3002"),(0,i.kt)("p",null,"\u5b83\u53ef\u4ee5\u540c\u65f6\u8fde\u63a5\u591a\u4e2a openGauss \u8282\u70b9\uff0c\u4ee5\u4e00\u5b9a\u7684\u65f6\u95f4\u95f4\u9694\u91c7\u96c6\u8fd0\u884c\u6570\u636e\u3002\u8fd9\u4e9b\u6570\u636e\u5c06\u88ab\u4e34\u65f6\u5b58\u50a8\u5728\u5185\u5b58\u7f13\u5b58\u4e2d\uff0c\u4ee5\u4f9b\u91c7\u96c6 API \u8c03\u7528\u3002"))}m.isMDXComponent=!0}}]);