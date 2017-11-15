<p><a href="#!/login" class="btn btn-primary">Logout</a></p>

<!-- Rank Choice -->
<img src="/assets/img/BallotTop.png">
<ul>
    <li ng-repeat="selection in vm.rcv.voter_selections">
        <p>Choice Rank: {{$index + 1}}</p>
        <select
            ng-model="selection"
            ng-options="option for option in vm.rcv.ballot_options"
            ng-change="vm.updateRCV(selection, $index)">
        </select>
    </li>
</ul>
<p>Add write-in selection: </p>
<div class="form-group">
    <label for="writeinRvcCic">Commander in Cream (CIC): </label>
    <input type="text" name="writeinRvcCic" id="writeinRvcCic" ng-model="writeinRvcCic" />
</div>
<div class="form-group">
    <label for="writeinRvcVi">For Vice Ice: </label>
    <input type="text" name="writeinRvcVi" id="writeinRvcVi" ng-model="writeinRvcVi" />
</div>
<button class="btn btn-primary" ng-click="vm.addRCVWritein(writeinRvcCic, writeinRvcVi)">Add</button>

<p>&nbsp;</p>

<!-- Unexpired Term -->
<img src="/assets/img/ChiefDairyQueen.png">
<form name="unexpiredTerm">
    <div class="form-group" ng-repeat="ballot_option in vm.ut.ballot_options">
        <label>
            <input type="radio" ng-model="vm.ut.voter_selection" value="{{ballot_option}}">
            {{ballot_option}}
        </label><br/>
    </div>
<!-- {{vm.ut.voter_selection}} -->
</form>

<p>&nbsp;</p>

<!-- Vote For 2 -->
<img src="/assets/img/DistrictM&M.png">
<form name="voteFor2" ng-init="vf2_flags = []">
    <div class="form-group" ng-repeat="ballot_option in vm.vf2.ballot_options">
        <input type="checkbox" ng-model="vf2_flags[$index]"
            ng-change="vm.updateVF2(ballot_option, vf2_flags[$index])"
            ng-disabled="!vf2_flags[$index] && vm.vf2.voter_selections.length > 1" />
        <label for="{{ballot_option}}">{{ballot_option}}</label>
    </div>
    <p>Add write-in selection: </p>
    <div class="form-group">
        <input type="text" name="writeinVf2" id="writeinVf2" ng-model="writeinVf2" />
    </div>
    <button ng-click="vm.addVF2Writein(writeinVf2)">Add</button>
</form>
<!-- {{vm.vf2.voter_selections}} -->
<p>&nbsp;</p>

<!-- Ballot Issue -->
<img src="/assets/img/BallotIssue.png">
<form name="ballotIssue">
    <div class="form-group" ng-repeat="ballot_option in vm.bi.ballot_options">
        <label>
            <input type="radio" ng-model="vm.bi.voter_selection" value="{{ballot_option}}">
            {{ballot_option}}
        </label><br/>
    </div>
<!-- {{vm.bi.voter_selection}} -->
</form>

<p>&nbsp;</p>

<button ng-click="vm.submitVote()" class="btn btn-primary">Submit</button>
<img ng-if="vm.dataLoading" src="data:image/gif;base64,R0lGODlhEAAQAPIAAP///wAAAMLCwkJCQgAAAGJiYoKCgpKSkiH/C05FVFNDQVBFMi4wAwEAAAAh/hpDcmVhdGVkIHdpdGggYWpheGxvYWQuaW5mbwAh+QQJCgAAACwAAAAAEAAQAAADMwi63P4wyklrE2MIOggZnAdOmGYJRbExwroUmcG2LmDEwnHQLVsYOd2mBzkYDAdKa+dIAAAh+QQJCgAAACwAAAAAEAAQAAADNAi63P5OjCEgG4QMu7DmikRxQlFUYDEZIGBMRVsaqHwctXXf7WEYB4Ag1xjihkMZsiUkKhIAIfkECQoAAAAsAAAAABAAEAAAAzYIujIjK8pByJDMlFYvBoVjHA70GU7xSUJhmKtwHPAKzLO9HMaoKwJZ7Rf8AYPDDzKpZBqfvwQAIfkECQoAAAAsAAAAABAAEAAAAzMIumIlK8oyhpHsnFZfhYumCYUhDAQxRIdhHBGqRoKw0R8DYlJd8z0fMDgsGo/IpHI5TAAAIfkECQoAAAAsAAAAABAAEAAAAzIIunInK0rnZBTwGPNMgQwmdsNgXGJUlIWEuR5oWUIpz8pAEAMe6TwfwyYsGo/IpFKSAAAh+QQJCgAAACwAAAAAEAAQAAADMwi6IMKQORfjdOe82p4wGccc4CEuQradylesojEMBgsUc2G7sDX3lQGBMLAJibufbSlKAAAh+QQJCgAAACwAAAAAEAAQAAADMgi63P7wCRHZnFVdmgHu2nFwlWCI3WGc3TSWhUFGxTAUkGCbtgENBMJAEJsxgMLWzpEAACH5BAkKAAAALAAAAAAQABAAAAMyCLrc/jDKSatlQtScKdceCAjDII7HcQ4EMTCpyrCuUBjCYRgHVtqlAiB1YhiCnlsRkAAAOwAAAAAAAAAAAA==" />

<p>&nbsp;</p>

<div ng-class="{ 'alert': flash, 'alert-success': flash.type === 'success', 'alert-danger': flash.type === 'error' }" ng-if="flash" ng-bind="flash.message"></div>
