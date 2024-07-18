<script>
export default {
	data: function() {
		return {
            username: null,
            id: null,
			usernameValidation: new RegExp('^[a-z][a-z0-9]*$'),
		}
	},
	methods: {
        async login() {
			try {
                if (!this.usernameValidation.test(this.username)) throw "Invalid username, it must start with letters and contain only lowercase letters and numbers"
                if (this.username.length < 3 || this.username.length > 13) throw "Invalid username, it must contains mininum 3 characters and maximum 13 characters"
                let response = await this.$axios.post('/session', {
                   "username": this.username,
                });
                console.log(response)
				this.id = response.data;
                console.log("Id is",this.id);
				localStorage.uid = this.id;
				localStorage.username = this.username;
                    
                this.$router.push("/home");
                
           } catch (e) {
                console.log("ciao simo")
                console.log(e)
                //document.getElementsByTagName("input")[0].style.outline = "auto";
                //document.getElementsByTagName("input")[0].style.outlineColor = "red";
            };
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
