<script>
import axios from "../services/axios.js"

export default {
  data() {
    return {
      errormsg: null,
      loading: false,
      currentUsername: null,
      searchUsername: null,
      searched: null,
    };
  },
  methods: {
    async searchUsers() {
      this.loading = true;
      this.errormsg = null;
      try {
        if (!this.searchUsername) {
          throw new Error("Please enter a username.");
        }

        // Fetch the user profile by username using the /profile endpoint
        const path = `/profile`;
        console.log(`Making API call to: ${path}`);
        let response = await axios.get(path, {
          params: { username: this.searchUsername }
        });

        if (response.status === 200) {
          // Update the searched result
		  
          this.searched = response.data;
		  console.log(this.searched)
          if (this.searched) {
            // Redirect to the user's profile page using the retrieved userID
            this.$router.push(`/profiles/${this.searched.user.user_id}/profile`);
          } else {
            this.errormsg = "User not found or incomplete response data.";
          }
        } else {
          // Handle non-200 responses
          this.errormsg = `Unexpected response status: ${response.status}`;
        }
      } catch (e) {
        // Handle errors from the API or other sources
        this.errormsg = e.response && e.response.data ? e.response.data : e.toString();
      }
      this.loading = false;
    }
  },
  mounted() {
    this.currentUsername = localStorage.getItem('username');
  }
}
</script>

<template>
  <div class="stream-page">
    <div>
      <br>
      <input v-model="searchUsername" placeholder="Enter username" />
      <button @click="searchUsers">Search</button>
    </div>
    
    <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
    
    <div v-if="loading">Loading...</div>
  </div>
</template>

<style>
</style>
