<script>

export default {
    data: function() {
        return {
            image: null,
            imageUrl: null
        }
    },
    methods: {
        onChange(e) {
            const file = e.target.files[0];
            this.image = file;
            this.imageUrl = URL.createObjectURL(file);
        },
        async uploadImage() {

            try {
                let userID = localStorage.getItem('token');
                await this.$axios.post(`/users/${userID}/photos/`, this.image);
                this.$router.push(`/users/${userID}/profile`);
            } catch (e) {
                this.errormsg = e.toString();
            }
        },
        cancelImage() {
            this.image = null;
            this.imageUrl = null;
            this.$refs.imgInput.value = ''; 
        }
    }
}
</script>

<template>
    <div
		class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
		<h1 class="h2">New Photo</h1>
	</div>

    <ErrorMsg v-if="errormsg"></ErrorMsg>

    <div>
        <form>
            <div class="form-group mb-3">
                <label class="mb-2" for="imgInput">Upload a new photo</label>
                <input type="file" accept="image/*" class="form-control" @change="onChange" id="imgInput" ref="imgInput">
            </div>
        </form>
        <div v-if="imageUrl" class="upload-content">
            <div class="photo-container image-preview">
                <h5 class="d-flex justify-content-center">Image preview</h5>
                <img :src="imageUrl" class="img-thumbnail" alt="Image preview" />
            </div>
            <div class="btn-toolbar">
                <div class="btn-group me-2">
                    <button v-if="image" @click="uploadImage" class="btn btn-primary btn-lg">Upload Image</button>
                </div>
                <div class="btn-group me-2">
                    <button v-if="image" @click="cancelImage" class="btn btn-secondary btn-lg">Cancel</button>
                </div>
            </div>
        </div>
    </div>
</template>

<style>

    .upload-content {
        display: flex;
        flex-direction: column;
        align-items: center;
    }

    .file-input {
        margin-bottom: 20px;
        padding: 10px;
        border: 1px solid #ccc;
        border-radius: 4px;
        font-size: 16px;
    }

    .image-preview {
        max-width: 300px; /* Larghezza massima della box della foto */
    }
    
</style>
  