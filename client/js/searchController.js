app.controller("searchController", function($scope, $http){
	$scope.game = "magic";

	$scope.possibleCompletions = [];

	$scope.search = function(input) {
		$http.get("/search/" + $scope.game + "/" + input).success(function(data){
			console.log(data);
			$scope.possibleCompletions = data;
		});
	};


});
