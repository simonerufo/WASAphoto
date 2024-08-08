<script>
import axios from "../services/axios.js";

export default {
  data() {
    return {
      errormsg: null,
      loading: false,
      currentUsername: null,
      userID: 0,
      userPhotos: [], // This will hold the photos from the main user
      followedUsers: [], // List of users followed by the main user
      followedUsernames: {} // Mapping of user IDs to usernames
    };
  },
  methods: {
    async refresh() {
      this.loading = true;
      this.errormsg = null;
      try {
        // Fetch user information and followed users
        await this.fetchUserInfo();
        // Fetch user photos after getting the main data
        await this.fetchUserPhotos();
        // Sort photos by timestamp after fetching
        this.sortPhotosByTimestamp();
      } catch (e) {
        this.errormsg = e.toString();
      }
      this.loading = false;
    },
    formatTimestamp(timestamp) {
      const date = new Date(timestamp);
      const options = {
        year: 'numeric',
        month: 'short',
        day: '2-digit',
        hour: '2-digit',
        minute: '2-digit',
        second: '2-digit',
      };
      return date.toLocaleDateString('en-US', options);
    },
    sortPhotosByTimestamp() {
      this.userPhotos.sort((a, b) => new Date(b.timestamp) - new Date(a.timestamp));
    },
    async fetchUserInfo() {
      try {
        this.userID = localStorage.getItem("id");
        this.currentUsername = localStorage.getItem('username');
        await this.fetchFollowedUsers(this.userID); // Fetch followed users
      } catch (e) {
        this.errormsg = e.toString();
      }
    },
    async fetchFollowedUsers(userId) {
      try {
        const response = await axios.get(`/profiles/${userId}/following`);
        console.log("Followed users:", response.data);
        this.followedUsers = response.data;
      } catch (e) {
        this.errormsg = e.toString();
      }
    },
    async fetchUsername(userID) {
      let path = `/profiles/${userID}/profile`;
      try {
        let response = await axios.get(path);
        return response.data.user.username;
      } catch (error) {
        console.error(error);
        return "Unknown";
      }
    },
    async fetchUserPhotos() {
      try {
        this.userPhotos = []; // Clear previous photos
        this.followedUsernames = {}; // Clear previous usernames

        // Fetch the stream for the main user
        console.log(`Fetching photos for main user ID ${this.userID}`);
        const path = `/profiles/${this.userID}/stream`;
        const response = await axios.get(path);
        console.log(`Response data for main user ID ${this.userID}:`, response.data);

        // Check the structure of response.data
        if (Array.isArray(response.data)) {
          this.userPhotos = response.data;

          // Fetch usernames for each photo
          for (let photo of this.userPhotos) {
            const username = await this.fetchUsername(photo.user_id);
            this.followedUsernames = { ...this.followedUsernames, [photo.user_id]: username };
          }
        } else {
          console.warn(`Unexpected data format for the main user ${this.currentUsername}:`, response.data);
        }
      } catch (e) {
        this.errormsg = e.toString();
      }
    },
  },
  async mounted() {
    await this.refresh();
  },
};
</script>

<template>
  <div class="stream-page">
    <div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
      <h2>{{ currentUsername }}'s Stream</h2>
    </div>

    <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

    <div v-if="loading">Loading...</div>

    <div v-if="userPhotos.length > 0">
      <div v-for="photo in userPhotos" :key="photo.photo_id" class="col-md-4 mb-4">
        <img :src="photo.image" class="img-fluid photo" :alt="followedUsernames[photo.user_id] || 'User photo'" @error="handleImageError">
        <div class="photo-info">
          <p class="photo-username"><strong>{{ followedUsernames[photo.user_id] || 'Loading...' }}</strong> </p>
          <p class="photo-timestamp"><strong>Timestamp:</strong> {{ formatTimestamp(photo.timestamp) }}</p>
          <p class="photo-caption"><strong>Caption:</strong> {{ photo.caption }}</p>
        </div>
      </div>
    </div>
    <div v-else>No photos available.</div>
  </div>
</template>

<style>
.photo {
  width: 100%;
  height: auto;
  object-fit: cover;
}
.photo-info {
  margin-top: 10px;
}
.photo-username,
.photo-timestamp,
.photo-caption {
  margin: 5px 0;
}
</style>
