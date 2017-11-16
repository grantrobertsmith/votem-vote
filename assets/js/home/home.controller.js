(function () {
    'use strict';

    angular
        .module('app')
        .controller('HomeController', HomeController);

    HomeController.$inject = ['VoteService', 'UserService', '$rootScope', '$timeout', '$scope', '$route', 'FlashService'];
    function HomeController(VoteService, UserService, $rootScope, $timeout, $scope, $route, FlashService) {
        const EMPTY_SELECTION_CANDIDACY = `-- SELECT CANDIDATES --`;

        var vm = this;

        vm.user = null;
        vm.allUsers = [];
        vm.rcv = {
            voter_selections: [],
            ballot_options: [
                EMPTY_SELECTION_CANDIDACY,
                `Democratic Ticket:
    Reese WithoutASpoon for CIC
    Cherry Garia for Vice Ice`,
                `Republican Ticket:
    Choco 'Chip' Dough for CIC
    Carmela Coney for Vice Ice`,
                `Independent Ticket:
    Magic Browny for CIC
    Phish Food for Vice Ice`,
            ]
        };
        vm.ut = {
            voter_selection: null,
            ballot_options: [
                'YES',
                'NO',
                'ABSTAIN (I choose not to vote.)'
            ]
        };
        vm.vf2 = {
            voter_selections: [],
            ballot_options: [
                'P. Nut Butter (REPUBLIAN)',
                'Marsh Mallow (DEMOCRAT)',
                'Cream C. Kol (INDEPENDENT)'
            ]
        };
        vm.bi = {
            voter_selection: null,
            ballot_options: [
                'YES',
                'NO',
                'ABSTAIN (I choose not to vote.)'
            ]
        };

        vm.updateRCV = updateRCV;
        vm.addRCVWritein = addRCVWritein;
        vm.updateVF2 = updateVF2;
        vm.addVF2Writein = addVF2Writein;
        vm.submitVote = submitVote;

        initController();

        function initController() {
            loadCurrentUser();
            initRCVBallot();
            initUTBallot();
            initVoteFor2Ballot();
            initBIBallot();
        }

        function loadCurrentUser() {
            vm.user = $rootScope.globals.currentUser;
        }

        function initRCVBallot() {
            vm.rcv.voter_selections.push(vm.rcv.ballot_options[0]);
        }

        function initUTBallot() {
            vm.ut.voter_selection = vm.ut.ballot_options[2];
        }

        function initVoteFor2Ballot() {}

        function initBIBallot() {
            vm.bi.voter_selection = vm.bi.ballot_options[2];
        }

        function updateRCV(selection, index) {
            var oldLocation = vm.rcv.voter_selections.indexOf(selection);
            // var oldValue = vm.rcv.voter_selections[index];
            vm.rcv.voter_selections[index] = selection;

            if (selection === EMPTY_SELECTION_CANDIDACY && vm.rcv.voter_selections.length > 1) {
                // We removed a selection from the list
                vm.rcv.voter_selections.splice(index, 1);

            } else if (oldLocation !== -1) {
                // We reorder the list to reflect the new rank selected for the selected contest option
                vm.rcv.voter_selections.splice(oldLocation, 1);

                // if (oldValue !== EMPTY_SELECTION_CANDIDACY) {
                //     vm.rcv.voter_selections.splice(oldLocation, 0, oldValue);
                // }
            }

            if (vm.rcv.voter_selections[vm.rcv.voter_selections.length - 1] !== EMPTY_SELECTION_CANDIDACY) {
                // We added another selection to the list
                $timeout(function () { // Force it to happen on a seperate iteration so the UI updates
                    // Clearly this should be done in some way idomatic of angular, I just don't know what that is right now and it's taking too much time to figure it out
                    vm.rcv.voter_selections.push(EMPTY_SELECTION_CANDIDACY);
                }, 1);
            }
        }

        function addRCVWritein(writeinRvcCic, writeinRvcVi) {
            var writeinString = `Write-In Ticket:
    ${writeinRvcCic} for CIC
    ${writeinRvcVi} for Vice Ice`

            if (vm.rcv.ballot_options.indexOf(writeinString) === -1) {
                vm.rcv.ballot_options.push(writeinString);
            }

            // Clear out the write-in form so it can be used again
            $scope.writeinRvcVi = $scope.writeinRvcCic = '';
        }

        function updateVF2(selection, isSelected) {
            if (isSelected) {
                vm.vf2.voter_selections.push(selection);
    
            } else {
                var index = vm.vf2.voter_selections.indexOf(selection);
                vm.vf2.voter_selections.splice(index, 1);
            }
        }
    
        function addVF2Writein(writeinVf2) {
            if (vm.vf2.ballot_options.indexOf(writeinVf2) === -1) {
                vm.vf2.ballot_options.push(writeinVf2);
            }
    
            // Clear out the write-in form so it can be used again
            $scope.writeinVf2 = '';
        }

        function submitVote() {
            // remove placeholder used for UI
            vm.rcv.voter_selections.splice(vm.rcv.voter_selections.length - 1, 1);

            vm.dataLoading = true;

            VoteService.SubmitBallot({
                voter_email: $rootScope.globals.currentUser.email,
                rcv: vm.rcv.voter_selections,
                unexpired_term: vm.ut.voter_selection,
                vote_for_2: vm.vf2.voter_selections,
                ballot_issue: vm.bi.voter_selection

            }).then(function (response) {
                if (response.success) {
                    FlashService.Success('You voted!', true);
                    $route.reload();

                } else {
                    FlashService.Error(response.message);
                    vm.dataLoading = false;
                }
            });
        }
    }

})();
