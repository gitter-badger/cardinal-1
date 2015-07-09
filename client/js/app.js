var app = angular.module('card-manager', ['ngMaterial', 'ngRoute', 'ngMessages']);

app.config(['$routeProvider', function($routeProvider){
    $routeProvider.when('/login', {
        templateUrl: 'views/login.html',
        controller: 'loginController'
    }).when('/dashboard', {
        templateUrl: 'views/dashboard.html',
        controller: 'dashboardController'
    }).when('/createCollection',{
        templateUrl: 'views/createCollection.html',
        controller: 'searchController'
    }).otherwise({
        redirectTo: '/login'
    });
}]);

app.service('authService', function(){
    var user = this;
});

app.service('cookieService', function(){
    getCookie = function(name){
        var re = new RegExp(name + "=([^;]+)");
        var value = re.exec(document.cookie);
        return (value !== null) ? unescape(value[1]) : null;
    };

    setCookie = function(name, value){

    };

    return {
        getCookie: getCookie,
        setCookie: setCookie
    };
});

app.controller('rootController', ['$scope', '$mdSidenav', function($scope, $mdSidenav, $message){
    console.log("Root loaded");
}]);
