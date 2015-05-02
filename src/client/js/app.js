var app = angular.module('cube-builder', ['ngMaterial', 'ngRoute', 'ngMessages']);

app.config(['$routeProvider', function($routeProvider){
    $routeProvider.when('/login', {
        templateUrl: 'views/login.html',
        controller: 'loginController'
    }).when('/dashboard', {
        templateUrl: 'views/dashboard.html',
        controller: 'dashboardController'
    }).otherwise({
        redirectTo: '/login'
    });
}]);

app.service('authService', function(){
    var user = this;
});

app.controller('rootController', ['$scope', '$mdSidenav', function($scope, $mdSidenav, $message){

}]);
