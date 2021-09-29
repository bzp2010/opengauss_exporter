"use strict";(self.webpackChunk=self.webpackChunk||[]).push([[652],{3005:function(e,t,n){n.r(t),n.d(t,{frontMatter:function(){return s},contentTitle:function(){return i},metadata:function(){return u},toc:function(){return p},default:function(){return c}});var r=n(7462),a=n(3366),o=(n(7294),n(3905)),l=["components"],s={title:"5\u5206\u949f\u4f53\u9a8c"},i=void 0,u={unversionedId:"getting-started/5min-tutorial",id:"getting-started/5min-tutorial",isDocsHomePage:!1,title:"5\u5206\u949f\u4f53\u9a8c",description:"\u6982\u8ff0",source:"@site/docs/getting-started/5min-tutorial.md",sourceDirName:"getting-started",slug:"/getting-started/5min-tutorial",permalink:"/opengauss_exporter/en/getting-started/5min-tutorial",editUrl:"https://github.com/bzp2010/opengauss_exporter/edit/main/docs/docs/getting-started/5min-tutorial.md",tags:[],version:"current",frontMatter:{title:"5\u5206\u949f\u4f53\u9a8c"},sidebar:"tutorialSidebar",previous:{title:"\u7b80\u4ecb",permalink:"/opengauss_exporter/en/"},next:{title:"\u5b89\u88c5",permalink:"/opengauss_exporter/en/getting-started/installation"}},p=[{value:"\u6982\u8ff0",id:"\u6982\u8ff0",children:[]},{value:"\u524d\u63d0\u6761\u4ef6",id:"\u524d\u63d0\u6761\u4ef6",children:[]},{value:"\u7b2c\u4e00\u6b65\uff1a\u5b89\u88c5\u5404\u4e2a\u7ec4\u4ef6",id:"\u7b2c\u4e00\u6b65\u5b89\u88c5\u5404\u4e2a\u7ec4\u4ef6",children:[]},{value:"\u7b2c\u4e8c\u6b65\uff1a\u8bbf\u95ee Prometheus UI \u5e76\u67e5\u770b\u6307\u6807",id:"\u7b2c\u4e8c\u6b65\u8bbf\u95ee-prometheus-ui-\u5e76\u67e5\u770b\u6307\u6807",children:[]}],m={toc:p};function c(e){var t=e.components,n=(0,a.Z)(e,l);return(0,o.kt)("wrapper",(0,r.Z)({},m,n,{components:t,mdxType:"MDXLayout"}),(0,o.kt)("h2",{id:"\u6982\u8ff0"},"\u6982\u8ff0"),(0,o.kt)("p",null,"\u672c\u6587\u662f openGauss \u7684\u5feb\u901f\u5165\u95e8\u6307\u5357\u7684\u5feb\u901f\u4f53\u9a8c\u90e8\u5206\u3002\u5feb\u901f\u5165\u95e8\u5206\u4e3a\u4e24\u4e2a\u6b65\u9aa4\uff1a"),(0,o.kt)("ol",null,(0,o.kt)("li",{parentName:"ol"},"\u901a\u8fc7 Docker \u5b89\u88c5\uff1a",(0,o.kt)("ul",{parentName:"li"},(0,o.kt)("li",{parentName:"ul"},"openGauss \u6570\u636e\u5e93"),(0,o.kt)("li",{parentName:"ul"},"openGauss Exporter"),(0,o.kt)("li",{parentName:"ul"},"Prometheus Server"))),(0,o.kt)("li",{parentName:"ol"},"\u7b49\u5f85\u7247\u523b\uff0c\u5373\u53ef\u901a\u8fc7 Prometheus UI \u67e5\u770b\u6307\u6807\u6570\u636e")),(0,o.kt)("h2",{id:"\u524d\u63d0\u6761\u4ef6"},"\u524d\u63d0\u6761\u4ef6"),(0,o.kt)("ul",null,(0,o.kt)("li",{parentName:"ul"},"\u5df2\u5b89\u88c5 Docker \u7684\u8fd0\u884c\u73af\u5883")),(0,o.kt)("h2",{id:"\u7b2c\u4e00\u6b65\u5b89\u88c5\u5404\u4e2a\u7ec4\u4ef6"},"\u7b2c\u4e00\u6b65\uff1a\u5b89\u88c5\u5404\u4e2a\u7ec4\u4ef6"),(0,o.kt)("ol",null,(0,o.kt)("li",{parentName:"ol"},"\u5b89\u88c5 openGauss \u6570\u636e\u5e93")),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-docker"},"docker run --name opengauss --privileged=true -d -e GS_PASSWORD=gaussdb!123 enmotech/opengauss:latest\n")),(0,o.kt)("ol",{start:2},(0,o.kt)("li",{parentName:"ol"},"\u521b\u5efa openGauss Exporter \u914d\u7f6e\u6587\u4ef6")),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-yaml"},'server:\n  http:\n    host: 0.0.0.0\n    port: 9188\n\ndata_sources:\n  - dsn: "postgresql://gaussdb:gaussdb!123@127.0.0.1:5432/postgres"\n    duration: 5s\n    max_retry: 3\n    master: true\n    enable_postgresql_exporter: true\n    enable_settings: true\n    enable_os_run_info: true\n    enable_total_memory_detail: true\n    enable_sql_count: true\n    enable_instance_time: true\n')),(0,o.kt)("p",null,"\u590d\u5236\u4ee5\u4e0a\u914d\u7f6e\u6587\u4ef6\u5b58\u50a8\u6210\u540d\u4e3a",(0,o.kt)("inlineCode",{parentName:"p"},"config.yaml"),"\u7684\u6587\u4ef6"),(0,o.kt)("ol",{start:3},(0,o.kt)("li",{parentName:"ol"},"\u542f\u52a8 openGauss Exporter")),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-docker"},"docker run -d -v $(pwd)/config.yaml:/etc/opengauss_exporter.yaml -p 9188:9188 bzp2010/opengauss_exporter -c /etc/opengauss_exporter.yaml\n")),(0,o.kt)("ol",{start:4},(0,o.kt)("li",{parentName:"ol"},"\u521b\u5efa Prometheus \u914d\u7f6e\u6587\u4ef6")),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-yaml"},'global:\n  scrape_interval: 15s\n  evaluation_interval: 15s\n\nscrape_configs:\n  - job_name: "prometheus"\n    static_configs:\n      - targets: ["localhost:9090"]\n\n  - job_name: "opengauss"\n    scrape_interval: 5s\n    static_configs:\n      - targets: ["127.0.0.1:9188"]\n')),(0,o.kt)("p",null,"\u590d\u5236\u4ee5\u4e0a\u914d\u7f6e\u6587\u4ef6\u5b58\u50a8\u6210\u540d\u4e3a",(0,o.kt)("inlineCode",{parentName:"p"},"prometheus.yaml"),"\u7684\u6587\u4ef6"),(0,o.kt)("ol",{start:5},(0,o.kt)("li",{parentName:"ol"},"\u542f\u52a8 Prometheus Server")),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-docker"},"docker run -d -p 9090:9090 -v $(pwd)/prometheus.yaml:/etc/prometheus/prometheus.yml prom/prometheus\n")),(0,o.kt)("h2",{id:"\u7b2c\u4e8c\u6b65\u8bbf\u95ee-prometheus-ui-\u5e76\u67e5\u770b\u6307\u6807"},"\u7b2c\u4e8c\u6b65\uff1a\u8bbf\u95ee Prometheus UI \u5e76\u67e5\u770b\u6307\u6807"),(0,o.kt)("ol",null,(0,o.kt)("li",{parentName:"ol"},"\u7b49\u5f85Exporter\u521d\u6b21\u522e\u524a\u53caPrometheus Server\u521d\u6b21\u91c7\u96c6\uff0c\u7ea6\u970015s"),(0,o.kt)("li",{parentName:"ol"},"\u4f7f\u7528\u6d4f\u89c8\u5668\u8bbf\u95ee ",(0,o.kt)("inlineCode",{parentName:"li"},"127.0.0.1:9090")),(0,o.kt)("li",{parentName:"ol"},"\u4f7f\u7528 PromQL \u67e5\u8be2\u6307\u6807 ",(0,o.kt)("inlineCode",{parentName:"li"},'{server:"127.0.0.1:5432"}'))))}c.isMDXComponent=!0}}]);