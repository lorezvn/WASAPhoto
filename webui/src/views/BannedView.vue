<script>


export default {
	data: function() {
		return {
			errormsg: null,
		}
	},

	watch: {
        '$route.params.userID': {
            handler(newUserID) {
                if (newUserID !== this.userID) {
                    this.userID = newUserID;
                    this.getUserProfile();
                }
            },
        }
    },
}
</script>

<template>
	<div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
		<h1 class="h2">You are banned</h1>
	</div>

	<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

    <div>
        <div class="d-flex flex-md-nowrap align-items-left">
            <div id="counter">
                <div class="follow-number">
                    0
                </div>
                <div class="follow-text">
                    photos
                </div>
            </div>
            <div id="counter">
                <div class="follow-number">
                    0
                </div>
                <div class="follow-text">
                    followers
                </div>
            </div>
            <div id="counter">
                <div class="follow-number">
                    0
                </div>
                <div class="follow-text">
                    following
                </div>
            </div>
        </div>

        <div v-if="!owner" style="margin-left: -910px;" class="btn-toolbar justify-content-center mt-3">
            <div class="btn-group me-2">
                <button v-if="!isFollowing" class="btn btn-primary btn-sm" @click="followUser" disabled>
                    Follow
                    <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#user-plus"/></svg>
                </button>
            </div>
            <div class="btn-group me-2">
                <button v-if="!isBanned" class="btn btn-primary btn-sm" @click="banUser" disabled>
                    Ban
                    <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#lock"/></svg>
                </button>
            </div>
        </div>
    </div>
    <div style="margin-top: 20px;">
        <h5>No photos posted</h5>
    </div>
</template>