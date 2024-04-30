<script>
export default {
	data: function() {
		return {
            username: null,
		}
	},
	methods: {
        login() {

            console.log("Username: "+this.username);
        }
	},
	mounted() {
	}
}
</script>

<template>
	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h2 class="h2">Login</h2>
		</div>
            
                <div class="form-group">
                    <label for="username">Username: </label>
                    <input type="text" class="form-control" id="username"  placeholder="Enter username" v-model="username">
                </div>
                <br>
                <button type="submit" class="btn btn-primary" @click="login">Login</button>
	</div>
</template>

<style>
</style>
