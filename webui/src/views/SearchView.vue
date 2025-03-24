<script>
import axios, { getId } from "../services/axios.js";

export default {
  data() {
    return {
      errormsg: null,
      loading: false,
      currentUsername: null,
      searchUsername: null,
      searched: null,
      bannedUsers: new Set(), 
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

        // Get the current user's ID
        const currentUserId = getId();

        // Get the user profile by username to get their ID
        const profileResponse = await axios.get("/profiles", {
          params: { username: this.searchUsername },
        });

        if (profileResponse.status === 200) {
          const userProfile = profileResponse.data;

          if (userProfile) {
            const targetUserId = userProfile.user.user_id;

            // Check if the current user is banned by the target user
            const banListResponse = await axios.get(`/profiles/${targetUserId}/bans/${currentUserId}`);
            if (banListResponse.status === 200) {
              const isBanned = banListResponse.data.is_banned;

              if (isBanned) {
                this.errormsg = "You are banned by this user.";
                this.loading = false;
                return;
              }
            } else {
              this.errormsg = `Unexpected response status from ban check: ${banListResponse.status}`;
              this.loading = false;
              return;
            }

            // Proceed with searching the profile if not banned
            this.searched = userProfile;
            this.$router.push(`/profiles/${targetUserId}/profile`);
          } else {
            this.errormsg = "User not found or incomplete response data.";
          }
        } else {
          this.errormsg = `Unexpected response status: ${profileResponse.status}`;
        }
      } catch (e) {
        this.errormsg = e.response && e.response.data ? e.response.data : e.toString();
      }

      this.loading = false;
    },
  },
  mounted() {
    this.currentUsername = localStorage.getItem("username");
  },
};
</script>

<template>
  <div class="search-page">
    <div class="search-container">
      <input
        v-model="searchUsername"
        placeholder="Search for a user..."
        class="search-input"
      />
      <button @click="searchUsers" class="search-button">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          width="20"
          height="20"
          fill="currentColor"
          class="bi bi-search"
          viewBox="0 0 16 16"
        >
          <path
            d="M11.742 10.344a6.5 6.5 0 1 0-1.397 1.398h-.001l3.85 3.85a1 1 0 0 0 1.415-1.415l-3.85-3.85zm-5.262 0a5.5 5.5 0 1 1 0-11 5.5 5.5 0 0 1 0 11z"
          />
        </svg>
      </button>
    </div>

    <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

    <div v-if="loading" class="loading-indicator">
      <div class="spinner"></div> Loading...
    </div>
  </div>
</template>

<style scoped>
/* Global Styles */
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
  background-color: #f1f1f1;
  color: #333;
}

/* Container Styling */
.search-page {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  background-color: #fafafa;
}

.search-container {
  display: flex;
  align-items: center;
  gap: 10px;
  max-width: 600px;
  width: 100%;
}

/* Input Styling */
.search-input {
  flex-grow: 1;
  padding: 14px;
  font-size: 16px;
  border: 2px solid #d1d1d1;
  border-radius: 50px;
  outline: none;
  background-color: #fff;
  transition: all 0.3s ease;
}

.search-input:focus {
  border-color: #2a9d8f;
  box-shadow: 0 0 8px rgba(42, 157, 143, 0.4);
}

/* Button Styling */
.search-button {
  background-color: #2a9d8f;
  color: white;
  border: none;
  border-radius: 50%;
  padding: 12px;
  cursor: pointer;
  transition: all 0.3s ease;
  display: flex;
  justify-content: center;
  align-items: center;
}

.search-button svg {
  width: 20px;
  height: 20px;
}

.search-button:hover {
  background-color: #264653;
}

/* Error Message */
.error-msg {
  margin-top: 15px;
  color: #e76f51;
  font-size: 14px;
  text-align: center;
  font-style: italic;
}

/* Loading indicator */
.loading-indicator {
  display: flex;
  justify-content: center;
  align-items: center;
  margin-top: 20px;
}

.spinner {
  border: 4px solid #f3f3f3;
  border-top: 4px solid #2a9d8f;
  border-radius: 50%;
  width: 30px;
  height: 30px;
  animation: spin 1s linear infinite;
}

/* Spinner animation */
@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}
</style>
