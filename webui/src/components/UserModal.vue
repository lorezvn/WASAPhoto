<script>
export default {
    props: ["isVisible", "title", "users"],
    emits: ['close'], 
    methods: {
        visitProfile(userID) {
            this.closeModal();
			this.$router.push("/users/"+userID+"/profile");
		},
        closeModal() {
            this.$emit('close');
        },
    }
}
</script>

<template>
    <div v-if="isVisible" class="modal-backdrop show"></div>
    <div v-if="isVisible" class="modal show" tabindex="-1" role="dialog" aria-hidden="true">
        <div class="modal-dialog modal-dialog-centered model-dialog-scrollable" role="document">
            <div class="modal-content">
                <div class="modal-header">
                    <h5 class="modal-title fw-bold">{{ title }} </h5>
                    <button type="button" class="btn-close" @click="closeModal" aria-label="Close">
                        <span aria-hidden="true"></span>
                    </button>
                </div>
                <div class="modal-body">
                    <div class="ms-2" v-if="users.length === 0">No users found</div>
                    <ul v-if=users.length class="list-group list-group-flush">
                        <li v-for="user in users" 
                            class="list-group-item d-flex justify-content-between align-items-center" 
                            @click="visitProfile(user.userID)" 
                            style="cursor: pointer;">
                            {{ user.username }}
                        </li>
                    </ul>
                </div>
                <div class="modal-footer">
                    <button class="btn btn-secondary btn-sm" @click="closeModal">Close</button>      
                </div>
            </div>
        </div>
    </div>
</template>


<style scoped>

    .modal-dialog {
        max-width: 30%; 
        height: 80vh;
    }

    .modal-content {
        max-height: 60%;
        display: flex;
        flex-direction: column;
    }

</style>
