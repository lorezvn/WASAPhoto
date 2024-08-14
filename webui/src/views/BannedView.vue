<template>
	<div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
		<h1 class="h2 null-element">-</h1>
	</div>

	<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

    <div class="profile-info mb-4">
        <div class="d-flex justify-content-center text-center">
            <div class="profile-counter">
                <div class="counter-number null-element">-</div>
                <div class="counter-text">photos</div>
            </div>
            <div class="profile-counter">
                <div class="counter-number null-element">-</div>
                <div class="counter-text">followers</div>
            </div>
            <div class="profile-counter">
                <div class="counter-number null-element">-</div>
                <div class="counter-text">following</div>
            </div>
        </div>

        <div v-if="!owner" class="btn-toolbar justify-content-center mt-3">
            <div class="btn-group me-2">
                <button class="btn btn-primary btn-sm" disabled>
                    Follow
                    <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#user-plus"/></svg>
                </button>
            </div>
            <div class="btn-group me-2">
                <button class="btn btn-primary btn-sm" disabled>
                    Ban
                    <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#lock"/></svg>
                </button>
            </div>
        </div>
    </div>

    <div class="photos-section mt-4">
        <div class="no-photos d-flex flex-column align-items-center justify-content-center pt-3 pb-2 border-top">
            <svg id="no-photos-icon" class="feather"><use href="/feather-sprite-v4.29.0.svg#camera-off"/></svg>
            <p class="no-photos-text">No photos</p>
        </div>
    </div>

</template>

<style scoped>

    .null-element {
        color:white;
    }
    
</style>