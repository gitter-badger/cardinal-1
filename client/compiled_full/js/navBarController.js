app.controller('navBarController', function($scope, authService){
    $scope.user = authService.user;
    console.log("Naggigator engage");
});
