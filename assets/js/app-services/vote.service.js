(function () {
    'use strict';

    angular
        .module('app')
        .factory('VoteService', VoteService);

    VoteService.$inject = ['$http'];
    function VoteService($http) {
        var service = {};

        service.SubmitBallot = SubmitBallot;

        return service;

        function SubmitBallot(ballot) {
            // console.log(ballot);
            return $http.post('/api/vote', ballot).then(handleSuccess, handleError('Error submitting ballot'));
        }

        // private functions

        function handleSuccess(res) {
            return res.data;
        }

        function handleError(error) {
            return function () {
                return { success: false, message: error };
            };
        }
    }

})();
