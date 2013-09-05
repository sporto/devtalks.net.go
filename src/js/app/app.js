(function () {
	console.log('APP');

	var APP = angular.module('APP', ['ngRoute']);

	APP.controller('SomeCtrl', function ($scope) {
		$scope.title = "Hello";
	});

	APP.config(['$routeProvider', function ($routeProvider) {
		$routeProvider
			.when('/', {
				templateUrl: '/index.html',
				controller: 'IndexCtrl'
			})
			.otherwise({redirectTo: '/'});
	}]);

	APP.controller('IndexCtrl', ['$scope', function ($scope) {
		console.log('IndexCtrl');
		$scope.message = "Hello";
	}]);

}());
