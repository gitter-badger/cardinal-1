app.controller('dashboardController', function($scope, $location, $mdSidenav, authService){
   $scope.user = authService.user;
    (function(){
        // if($scope.user === undefined || $scope.user.length === 0){
            // $location.url('login');
        // }
    })();
});
