app.controller("dashboardController",function(o,e,n,l){o.user=l.user,console.log(o.user),o.toggleMenu=function(){console.log("Toggle"),n("menu").toggle()},o.dashItems=[],o.buildDashboard=function(){$http.get("/get_dash/"+o.user.Username).success(function(o){console.log(o)})},o.createCollection=function(){console.log("I SHould change something"),e.url("createCollection")}});