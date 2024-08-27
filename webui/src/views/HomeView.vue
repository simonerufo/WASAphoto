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
      followedUsernames: {}, // Mapping of user IDs to usernames
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
        year: "numeric",
        month: "short",
        day: "2-digit",
        hour: "2-digit",
        minute: "2-digit",
        second: "2-digit",
      };
      return date.toLocaleDateString("en-US", options);
    },
    sortPhotosByTimestamp() {
      this.userPhotos.sort(
        (a, b) => new Date(b.timestamp) - new Date(a.timestamp)
      );
    },
    async fetchUserInfo() {
      try {
        this.userID = localStorage.getItem("id");
        this.currentUsername = localStorage.getItem("username");
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
        const path = `/profiles/${this.userID}/stream`;
        const response = await axios.get(path);

        if (response.data) {
          this.userPhotos = response.data.map((photo) => {
            if (photo.image && !photo.image.startsWith("data:image")) {
              photo.image = `data:image/jpeg;base64,${photo.image}`;
              console.log(photo.image);
            }
            return photo;
          });

          // Fetch usernames for each photo
          for (let photo of this.userPhotos) {
            const username = await this.fetchUsername(photo.user_id);
            this.followedUsernames = {
              ...this.followedUsernames,
              [photo.user_id]: username,
            };
          }
        } else {
          console.warn(
            `Unexpected data format for the main user ${this.currentUsername}:`,
            response.data
          );
        }
      } catch (e) {
        this.errormsg = e.toString();
      }
    },
    onImageError(photoId) {
      console.error(`Failed to load image for photo ID: ${photoId}`);
      // Use a placeholder image from Bootstrap as filler
      this.userPhotos = this.userPhotos.map((photo) => {
        if (photo.photo_id === photoId) {
          photo.image =
            //"https://via.placeholder.com/150?text=No+Image"; // Bootstrap placeholder image
            "https://loremflickr.com/320/240"
        }
        return photo;
      });
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

    <div v-if="userPhotos.length > 0" class="photo-grid">
      <div v-for="photo in userPhotos" :key="photo.photo_id" class="photo-card">
        <!-- RouterLink for the image -->
        <RouterLink :to="`/profiles/${photo.user_id}/photos/${photo.photo_id}`">
          <img
            :src="photo.image"
            alt="Photo"
            class="photo"
            @error="onImageError(photo.photo_id)"
          />
        </RouterLink>
        
        <!-- Photo information -->
        <div class="photo-info">
          <!-- RouterLink for the username -->
          <RouterLink :to="`/profiles/${photo.user_id}/profile`" class="photo-username">
            <strong>{{ followedUsernames[photo.user_id] || "Loading..." }}</strong>
          </RouterLink>
          
          <!-- Timestamp and caption -->
          <p class="photo-timestamp">
            <strong>Timestamp:</strong> {{ formatTimestamp(photo.timestamp) }}
          </p>
          <p class="photo-caption">
            <strong>Caption:</strong> {{ photo.caption }}
          </p>
        </div>
      </div>
    </div>
    <div v-else>No photos available.</div>
  </div>
</template>


<style scoped>
.stream-page {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 20px;
  font-family: "Arial", sans-serif;
  max-width: 800px;
  margin: 0 auto;
}

.photo-grid {
  display: flex;
  flex-wrap: wrap;
  gap: 20px;
  justify-content: center;
  width: 100%;
}

.photo-card {
  width: 300px;
  background: #e0e0e0;
  border: 1px solid #0033cc;
  border-radius: 8px;
  box-shadow: 2px 2px 5px #888;
  padding: 10px;
  text-align: center;
}

.photo {
  width: 100%;
  height: auto;
  object-fit: cover;
  border-radius: 5px;
  border: 1px solid #0033cc;
}

.photo-info {
  margin-top: 10px;
}

.photo-username,
.photo-timestamp,
.photo-caption {
  margin: 5px 0;
}

.photo-username {
  display: block;
  color: #0033cc;
  text-decoration: none;
  font-weight: bold;
  cursor: pointer;
}

.photo-username:hover {
  text-decoration: underline;
}
</style>
