<template>
	<div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
		<h1 class="h2">Not Found</h1>
	</div>

	<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

    <div>
        <div style="float:left;" class="d-flex flex-column mb-4">
            <div class="d-flex justify-content-center">
                <div id="counter">
                    <div class="follow-number">0</div>
                    <div class="follow-text">photos</div>
                </div>
                <div id="counter">
                    <div class="follow-number">0</div>
                    <div class="follow-text">followers</div>
                </div>
                <div id="counter">
                    <div class="follow-number">0</div>
                    <div class="follow-text">following</div>
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

        <div style="clear:both; height:40vh;" class="d-flex flex-column align-items-center justify-content-center pt-3 pb-2 mt-3 border-top">
            <svg id="icon-no-post" class="feather"><use href="/feather-sprite-v4.29.0.svg#camera-off"/></svg>
        </div>
    </div>

</template>

<style>
</style>