app.controller('loginController', function($scope, $http, $location, authService){
    $scope.user = {};
    $scope.user.username = "";
    $scope.user.password = "";
    $scope.user.$error = {};

    $scope.login = function(){
        if($scope.user.username === ""){
            $scope.user.$error.unReq = true;
        } else {
            $scope.user.$error.unReq = false;
        }

        if($scope.user.password === ""){
            $scope.user.$error.passReq = true;
        } else {
            $scope.user.$error.passReq = false;
        }

        $http.post('/login', $scope.user).success(function(data){
            console.log(data);
            if(data !== "FALSE") {
                $scope.auth = true;
                authService.user = $scope.user;
                $location.url("dashboard");
            } else {
                
            }
        });
    };
});
