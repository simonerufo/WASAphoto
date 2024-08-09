<script>
import axios,{getId} from "../services/axios.js";

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
        const profileResponse = await axios.get("/profile", {
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
  <div class="stream-page">
    <div class="search-container">
      <input
        v-model="searchUsername"
        placeholder="Enter username"
        class="search-input"
      />
      <button @click="searchUsers" class="search-button">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          width="16"
          height="16"
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

    <div v-if="loading">Loading...</div>
  </div>
</template>

<style scoped>
/* Centering the container */
.search-container {
  display: flex;
  justify-content: center;
  align-items: center;
  margin-top: 20px;
}

/* Styling the search input */
.search-input {
  width: 300px; /* Increased width for the search bar */
  padding: 10px; /* Padding inside the input */
  border: 2px solid #0033cc; /* Consistent border color */
  border-radius: 5px 0 0 5px; /* Rounded corners on the left side */
  font-size: 16px; /* Larger font size */
  outline: none;
}

/* Styling the search button */
.search-button {
  padding: 10px;
  background-color: #0033cc; /* Blue background */
  border: 2px solid #0033cc; /* Matching border color */
  border-radius: 0 5px 5px 0; /* Rounded corners on the right side */
  cursor: pointer;
  display: flex;
  justify-content: center;
  align-items: center;
}

.search-button svg {
  fill: #fff; /* White color for the search icon */
}

/* Adjusting hover effects */
.search-button:hover {
  background-color: #0055ff;
}

</style>
