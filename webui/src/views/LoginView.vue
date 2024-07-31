<script>
import { Auth } from '../services/axios';
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
                
				this.id = parseInt(response.data);
				localStorage.id = this.id;
				localStorage.username = this.username;
                this.$axios.defaults.headers.common['Authorization'] = 'Bearer ' + this.id;
                this.$router.push("/home");
                
           } catch (e) {
                console.log(e)
            };
        }
	},
	mounted() {
        localStorage.clear();
	}
}
</script>

<template>
  <div class="center-wrapper">
    <div class="login-container">
      <h2>Login</h2>
      <div class="form-group">
        <label for="username">Username: </label>
        <input type="text" class="form-control" id="username" placeholder="Enter username" v-model="username">
      </div>
      <button type="submit" class="btn btn-primary" @click="login">Login</button>
    </div>
  </div>
</template>



<style scoped>
/* Wrapper styling to center the login-container */
.center-wrapper {
  display: flex;
  justify-content: center; /* Center horizontally */
  align-items: center; /* Center vertically */
  height: 100vh; /* Full height of the viewport */
  width: 100vw; /* Full width of the viewport */
  background-color: #f0f0f0; /* Optional: Add a background color for better visibility */
  position: fixed; /* Ensure it stays centered */
  top: 0;
  left: 0;
}

/* Form container styling */
.login-container {
  background-color: white;
  padding: 20px;
  border-radius: 5px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  width: 300px;
}

/* Header styling */
.login-container h2 {
  margin-bottom: 20px;
  text-align: center;
  font-size: 1.5rem;
  color: #333;
}

/* Form group styling */
.login-container .form-group {
  margin-bottom: 15px;
}

.login-container label {
  display: block;
  margin-bottom: 5px;
  font-weight: bold;
  color: #333;
}

/* Input styling */
.login-container input.form-control {
  width: 100%;
  padding: 8px;
  border: 1px solid #ccc;
  border-radius: 4px;
  box-sizing: border-box;
}

/* Button styling */
.login-container button.btn-primary {
  width: 100%;
  padding: 10px;
  background-color: #4CAF50; /* Green background */
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 16px;
  transition: background-color 0.3s; /* Smooth transition on hover */
}

/* Button hover effect */
.login-container button.btn-primary:hover {
  background-color: #45a049; /* Darker green on hover */
}
</style>
