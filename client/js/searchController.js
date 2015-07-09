app.controller("searchController", function($scope, $http){
	$scope.game = "magic";

	$scope.possibleCompletions = [];

	$scope.search = function(input) {
		$http.get("/api/v1/cardSearch?game=" + $scope.game + "&cardName=" + input).success(function(data){
			console.log(data);
			$scope.possibleCompletions = data;
		});
	};


});
