<script>

export default {
  data() {
    return {
      username: null,
      id: null,
      usernameValidation: new RegExp('^[a-z][a-z0-9]*$'),
    }
  },
  methods: {

    async login() {
      try {
        if (!this.usernameValidation.test(this.username)) {
          alert("Invalid username: it must start with letters and contain only lowercase letters and numbers.");
          return;
        }
        if (this.username.length < 3 || this.username.length > 13) {
          alert("Invalid username: it must contain a minimum of 3 characters and a maximum of 13 characters.");
          return;
        }

        let response = await this.$axios.post('/session', {
          "username": this.username,
        });

        this.id = parseInt(response.data);
        localStorage.setItem("id", this.id);
        localStorage.setItem("username", this.username);

        this.$router.push("/home");

      } catch (e) {
        console.error('Login failed:', e);
        alert('Login failed. Please try again.');
      }
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
  background-color: #e0e0e0; /* Light grey background typical of Windows XP */
  position: fixed; /* Ensure it stays centered */
  top: 0;
  left: 0;
}

/* Form container styling */
.login-container {
  background-color: #ffffff; /* White background for the login box */
  padding: 20px;
  border: 2px solid #0033cc; /* Blue border to mimic Windows XP */
  border-radius: 4px; /* Rounded corners */
  box-shadow: 2px 2px 5px rgba(0, 0, 0, 0.2); /* Subtle shadow for 3D effect */
  width: 300px;
  font-family: 'Courier New', Courier, monospace; /* Retro font */
}

/* Header styling */
.login-container h2 {
  margin-bottom: 20px;
  text-align: center;
  font-size: 1.5rem;
  color: #0033cc; /* Blue text to match Windows XP color scheme */
  font-weight: bold;
  text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.3); /* Retro text shadow effect */
}

/* Form group styling */
.login-container .form-group {
  margin-bottom: 15px;
}

.login-container label {
  display: block;
  margin-bottom: 5px;
  font-weight: bold;
  color: #000000; /* Black text for labels */
  font-size: 14px;
}

/* Input styling */
.login-container input.form-control {
  width: 100%;
  padding: 8px;
  border: 1px solid #0033cc; /* Blue border */
  border-radius: 4px; /* Rounded corners */
  background-color: #f0f0f0; /* Light grey background inside input */
  box-sizing: border-box;
  font-family: 'Courier New', Courier, monospace; /* Retro font inside input */
}

/* Button styling */
.login-container button.btn-primary {
  width: 100%;
  padding: 10px;
  background-color: #0033cc; /* Windows XP classic blue */
  color: white;
  border: 1px solid #0033cc; /* Same color as the background for a uniform look */
  border-radius: 4px;
  cursor: pointer;
  font-size: 16px;
  transition: background-color 0.3s; /* Smooth transition on hover */
  font-family: 'Courier New', Courier, monospace; /* Retro font for button */
}

/* Button hover effect */
.login-container button.btn-primary:hover {
  background-color: #002a80; /* Darker blue on hover */
  box-shadow: 0 0 5px 2px rgba(0, 0, 0, 0.2); /* Subtle shadow on hover */
}
</style>
