
<script>
export default {
    props: ["photo", "comments", "likes", "owner"],
    data: function() {
        return {
            localLikes: [],
            localComments: [],
            errormsg: null,
            showModal: false
        }
    },
    emits: ['photoDeleted'], 
    methods: {
        async likePost(photo) {
            this.errormsg = null;
            try {
                let liked = this.isLiked(photo);
                let token = localStorage.getItem('token');
                if (!liked) {
                    await this.$axios.put(`/users/${photo.userID}/photos/${photo.photoID}/likes/${token}`);
                    this.localLikes.push({ userID: token });
                } else {
                    await this.$axios.delete(`/users/${photo.userID}/photos/${photo.photoID}/likes/${token}`);
                    this.localLikes = this.localLikes.filter(user => user.userID != token);
                }

            } catch (e) {
                this.errormsg = e.toString();
            }
        },
        async deletePhoto(photo) {
            try {
                await this.$axios.delete(`/users/${photo.userID}/photos/${photo.photoID}`);
                this.$emit('photoDeleted', photo.photoID);
            } catch (e) {
                this.errormsg = e.toString();
            }
        },
        handleCommentAdded(newComment) {
            this.localComments.push(newComment)
        },
        handleCommentDeleted(commentID) {
            this.localComments = this.localComments.filter(comment => comment.commentID != commentID);
        },
        isLiked() {
            let token = localStorage.getItem('token');
            return this.localLikes.some(user => user.userID == token);
        },
        formatDate(inputDate) {
            const options = { year: 'numeric', month: 'short', day: 'numeric', hour: 'numeric', minute: 'numeric' };
            const formattedDate = new Date(inputDate).toLocaleDateString('en-US', options);
            return formattedDate;
        },
        toggleModal() {
            this.showModal = true;
            document.body.classList.add("modal-open");
        },
        closeModal() {
            this.showModal = false;
            document.body.classList.remove("modal-open");
        }
    },

    mounted() {
        this.localLikes = this.likes;
        this.localComments = this.comments;
    }
};
</script>

<template>
    <div class="photo-container">
        <img :src="'data:image/jpeg;base64,'+photo.image" alt="Image">
        <div v-if="owner" class="delete-button-container">
            <div class="btn btn-sm btn-danger" @click="deletePhoto(photo)">
                Delete
                <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#trash-2"/></svg>
            </div>
        </div>
        <div class="d-flex align-items-center justify-content-between">
            <div>
                <div v-if="!owner" id="like-counter" :class="['btn btn-sm', isLiked() ? 'btn-outline-danger' : 'btn-outline-secondary']" @click="likePost(photo)">
                    {{ localLikes.length }}
                    <svg class="feather" :class="{'liked': isLiked()}">
                        <use href="/feather-sprite-v4.29.0.svg#heart"/>
                    </svg>
                </div>
                <button v-else id="like-counter" class="btn btn-outline-secondary btn-sm me" @click="likePost(photo)" disabled>
                    {{ localLikes.length }}
                    <svg class="feather" :class="{'liked': isLiked(photo)}">
                        <use href="/feather-sprite-v4.29.0.svg#heart"/>
                    </svg>
                </button>
                
                <button id="comment-counter" class="btn btn-outline-secondary btn-sm" @click="toggleModal">
                    {{ localComments.length }}
                    <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#message-circle"/></svg>
                </button>
            </div>
        </div>
        <small id="photo-date" class="photo-text"> {{ formatDate(photo.date) }} </small>
        <CommentModal
            :isVisible="showModal"
            :photo="photo"
            :comments="localComments"
            @commentAdded="handleCommentAdded" 
            @commentDeleted="handleCommentDeleted"
            @close="closeModal">
        </CommentModal>
    </div>
</template>


<style>
    .photo-container {
        text-align: left;
        width: 100%;
        height: 300px; 
        border-radius: 8px; 
        overflow: hidden; 
        margin-bottom: 45px;
    }

    .photo-container img {
        width: 100%;
        height: 100%;
        object-fit: cover; 
        border-radius: 8px; 
    }

    .photo-text {
        position: absolute;
        padding: 4px 8px;
        border-radius: 4px; 
        font-size: 14px; 
        color: #333; 
        font-family: Arial, Helvetica, sans-serif;
    }

    #photo-date {
        bottom:8px;
        right: 8px;
        color:gray
    }

    #like-counter {
        position: absolute;
        bottom: 8px; 
        left: 8px; 
    }

    #comment-counter {
        position: absolute;
        bottom: 8px; 
        left: 60px;
    }

    .delete-button-container {
        position: absolute;
        top: 8px;
        right: 8px;
        border-radius: 4px; 
    }

    .feather.liked use {
        fill: red;
    }

    body.modal-open {
        overflow: hidden;
    }
</style>
               