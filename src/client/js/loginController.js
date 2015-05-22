app.controller('loginController', function($scope, $http, $location, authService){
    $scope.user = {};
    $scope.user.username = "";
    $scope.user.password = "";
    $scope.user.email = "";
    $scope.user.$error = {};

    $scope.checkIfEmpty = function() {
        var somethingMissing = false;
        if($scope.user.username === ""){
            $scope.user.$error.unReq = true;
            somethingMissing = true;
        } else {
            $scope.user.$error.unReq = false;
        }

        if($scope.user.password === ""){
            $scope.user.$error.passReq = true;
            somethingMissing = true;
        } else {
            $scope.user.$error.passReq = false;
        }
        
        return somethingMissing;
    };

    $scope.login = function(){
        
        if($scope.user.$error.hasOwnProperty("loginFailed")){
            $scope.user.$error.loginFailed = false;
        }
        
        var somethingIsEmpty = $scope.checkIfEmpty();
        if(somethingIsEmpty){
            return;
        }
        
        console.log("logging in");

        $http.post('/login', angular.toJson($scope.user)).success(function(data){
            if(data !== 403) {
                console.log("Login sucess");
                $scope.user = data[0];
                authService.user = $scope.user;
                document.cookie = "token=" + data._id;
                $location.url("dashboard");
            }
        }).error(function(data){
            $scope.user.$error.loginFailed = true;
        });
    };

    $scope.signup = function(){
        var somethingIsEmpty = $scope.checkIfEmpty();
        console.log("Something");
        console.log(somethingIsEmpty);

        if($scope.user.email === "" || $scope.user.email.indexOf("@") === -1 || $scope.user.email.indexOf(".") === -1){
            $scope.user.$error.emailReq = true;
            return;
        } else {
            $scope.user.$error.emailReq = false;
        }
        
        if(somethingIsEmpty){
            return;
        }

        console.log("POSTING");
        $http.post('/signup', angular.toJson($scope.user)).success(function(data, status){
            console.log(data);
            if(data !== 403){
                $scope.user = data[0];
                authService.user = $scope.user;
                document.cookie = "token=" + data._id;
                $location.url("dashboard");
            }
        }).error(function(){
            $scope.user.$error.unTaken = true;
        });;
    };
});
