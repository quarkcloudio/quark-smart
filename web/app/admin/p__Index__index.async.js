"use strict";(self.webpackChunk=self.webpackChunk||[]).push([[500],{50565:function(V,S,a){a.r(S),a.d(S,{default:function(){return K}});var I=a(15009),l=a.n(I),T=a(99289),Z=a.n(T),z=a(5574),x=a.n(z),v=a(67294),d=a(80854),R=a(94740),$=a(3321),G=a(57953),M=a(97857),C=a.n(M),N=a(91373),D=a(46401),E=a(94785),r=a(85893),O={api:""},U=function(f){var e=(0,d.useLocation)(),u=E.Z.parse(e.search),c=C()(C()({},O),f),s=c.api,h=(0,v.useState)(""),p=x()(h,2),m=p[0],i=p[1],g=(0,d.useModel)("pageLoading"),o=g.setPageLoading,A=function(){var n=Z()(l()().mark(function L(){var j,P;return l()().wrap(function(t){for(;;)switch(t.prev=t.next){case 0:if(s){t.next=3;break}return i("\u8BF7\u8BBE\u7F6E\u63A5\u53E3\uFF01"),t.abrupt("return");case 3:return j={},Object.keys(u).forEach(function(y){y!=="api"&&(j[y]=u[y])}),o(!0),t.next=8,(0,N.U)({url:s,data:j});case 8:P=t.sent,i(P),o(!1);case 11:case"end":return t.stop()}},L)}));return function(){return n.apply(this,arguments)}}();return(0,v.useEffect)(function(){A()},[s,u.timestamp]),(0,r.jsx)(D.Z,{body:m})},B=U,W=a(33852),F=a(82925),H={loading:"loading___z87bD"},J=function(){var f=(0,d.useLocation)(),e=E.Z.parse(f.search),u=(0,v.useState)(String),c=x()(u,2),s=c[0],h=c[1],p=function(){var m=Z()(l()().mark(function i(){var g,o;return l()().wrap(function(n){for(;;)switch(n.prev=n.next){case 0:return n.next=2,(0,d.request)("./config.json");case 2:g=n.sent,o=g.api.default,e!=null&&e.api&&(o=e.api),h(o);case 6:case"end":return n.stop()}},i)}));return function(){return m.apply(this,arguments)}}();return(0,v.useEffect)(function(){p()},[f.search]),(0,r.jsx)(R.ZP,{locale:F.Z,children:(0,r.jsx)($.Z,{children:s?(0,r.jsx)(B,{api:s}):(0,r.jsx)("div",{className:H.loading,children:(0,r.jsx)(G.Z,{})})})})},K=J}}]);
