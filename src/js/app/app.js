(function () {
	console.log('APP');
	APP = angular.module('APP', []);

	APP.controller('videos.showCtrl', function ($scope) {
		$scope.message = "Hello";
	});
	
}());