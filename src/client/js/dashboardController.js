app.controller('dashboardController', function($scope, $location, $mdSidenav, authService){
    $scope.user = authService.user;
    console.log($scope.user);
    // Make sure someone is logged in
    // (function(){
    //     if($scope.user === undefined || $scope.user.length === 0)
    //         $location.url('login');
    // })();

    $scope.toggleMenu = function(){
        console.log("Toggle");
        $mdSidenav('menu').toggle();
    };

    // Holds the dashboardItems which contain info on the various collections
    $scope.dashItems = [];

    $scope.buildDashboard = function(){
        $http.get('/get_dash/' + $scope.user.Username).success(function(data){
            console.log(data);
        });
    };

    $scope.createCollection = function() {
        console.log("I SHould change something");
        $location.url('createCollection');
    };


});
