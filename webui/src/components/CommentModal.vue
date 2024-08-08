<script>
export default {
    props: ["isVisible", "photo"],
    data() {
        return {
            errormsg: null,
            newComment: ""
        }
    },
    emits: ['commentAdded', 'commentDeleted', 'close'], 
    methods: {
        async addComment() {
            this.errormsg = null;
            try {
                let response = await this.$axios.post(`/users/${this.photo.userID}/photos/${this.photo.photoID}/comments/`, {
                    message: this.newComment,
                });

                this.$emit('commentAdded', {
                    commentID: response.data.commentID,
                    userID: response.data.userID, 
                    username: response.data.username,
                    message: response.data.message,
                    date: response.data.date,
                });

                this.newComment = ""

                // Scroll to the bottom after the new comment is added
                this.$nextTick(() => {
                    let objDiv = document.getElementById("commentsContainer");
                    objDiv.scrollTop = objDiv.scrollHeight;
                });

            } catch (e) {
                this.errormsg = e.toString();
            }
        },
        async deleteComment(comment) {
            this.errormsg = null;
            try {
                await this.$axios.delete(`/users/${this.photo.userID}/photos/${this.photo.photoID}/comments/${comment.commentID}`);

                this.$emit('commentDeleted', comment.commentID);

            } catch (e) {
                this.errormsg = e.toString();
            }
        },
        isCommentOwner(comment) {
			return localStorage.getItem('token') == comment.userID;
		},
        visitProfile(userID) {
            this.closeModal();
			this.$router.push("/users/"+userID+"/profile");
		},
        closeModal() {
            this.$emit('close');
            this.newComment = "";
        },
        invalidComment(comment) {
            return comment.trim().length < 1 || comment.length > 2200
        },
        formatDate(inputDate) {
            const options = { year: 'numeric', month: 'short', day: 'numeric', hour: 'numeric', minute: 'numeric' };
            return new Date(inputDate).toLocaleDateString('en-US', options);
        }
    }
}
</script>

<template>
    <div v-if="isVisible" class="modal-backdrop show"></div>
    <div v-if="isVisible" class="modal show" tabindex="-1" role="dialog" aria-hidden="true">
        <div class="modal-dialog modal-dialog-centered model-dialog-scrollable" role="document">
            <div class="modal-content">
                <div class="modal-header">
                    <h5 class="modal-title">Comments {{photo.comments.length}} </h5>
                    <button type="button" class="btn-close" @click="closeModal" aria-label="Close">
                        <span aria-hidden="true"></span>
                    </button>
                </div>
                <div id="commentsContainer" class="modal-body">
                    <div class="ms-2" v-if="photo.comments.length === 0">No comments for this photo</div>
                    <ul class="list-group">
                        <li v-for="comment in photo.comments" 
                            class="list-group-item d-flex justify-content-between align-items-start mb-2">
                            <div class="ms-2 me-auto comment-content">
                                <button class="btn btn-sm fw-bold btn-outline-secondary comment-owner" @click="visitProfile(comment.userID)">@{{ comment.username }}</button>
                                <div class="p-2">{{ comment.message }}</div>
                            </div>
                            <div v-if="isCommentOwner(comment)" class="position-absolute top-0 end-0 mt-2 me-2">
                                <div class="btn btn-sm btn-danger" @click="deleteComment(comment)">
                                    <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#trash-2"/></svg>
                                </div>
                            </div>

                            <small class="position-absolute bottom-0 end-0 mb-2 me-2 comment-date">{{ formatDate(comment.date) }}</small>
                        </li>
                    </ul>
                </div>
                <div class="modal-footer">
                    <textarea 
                        v-model="newComment" 
                        class="form-control mb-2" 
                        rows="1" 
                        maxlength="2200"
                        :placeholder="'Add a comment for @'+photo.username+'...'">
                    </textarea>
                    <button class="btn btn-primary btn-sm" @click="addComment"
                    :disabled="invalidComment(newComment)">Submit</button>  
                    <button class="btn btn-secondary btn-sm" @click="closeModal">Close</button>      
                </div>
            </div>
        </div>
    </div>
</template>


<style scoped>

    .modal.show {
        display: block;
        opacity: 1;
    }

    .modal-backdrop {  
        opacity: 0.5;
    }

    .modal-dialog {
        max-width: 45%; 
        height: 80vh; /* Increased height */
    }

    .modal-content {
        max-height: 80%;
        display: flex;
        flex-direction: column;
    }

    .modal-body {
        overflow-y: auto;
        flex-grow: 1; /* Ensures the body grows to fill available space */
    }

    .list-group-item {
        padding: 5px;
        min-width: 100%;
    }

    .list-group-item:hover {
        cursor: default;
        background-color: #f8f9fa; 
    }

    .comment-owner {
        color:black
    }

    .comment-owner:hover {
        color:white
    }
    

    .comment-content {
        width: 100%; /* Adjust width as needed */
        max-width: 400px; /* Fixed width, adjust as needed */
        word-wrap: break-word;
        white-space: pre-wrap; /* Preserves line breaks in the content */
    }


    .comment-date {
        color: gray;
    }
</style>
