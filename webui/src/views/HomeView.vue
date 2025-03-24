<script>
import axios from "../services/axios.js";

export default {
  data() {
    return {
      errormsg: null,
      loading: false,
      currentUsername: null,
      userID: 0,
      userPhotos: [],
      followedUsers: [],
      followedUsernames: {},
    };
  },
  methods: {
    async refresh() {
      this.loading = true;
      this.errormsg = null;
      try {
        await this.fetchUserInfo();
        await this.fetchUserPhotos();
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
        console.log(this.userID);
        await this.fetchFollowedUsers(this.userID);
      } catch (e) {
        this.errormsg = e.toString();
      }
    },
    async fetchFollowedUsers(userId) {
      try {
        console.log("userID");
        console.log(userId);
        const response = await axios.get(`/profiles/${userId}/following`);
        this.followedUsers = response.data;
      } catch (e) {
        this.errormsg = e.toString();
      }
    },
    async fetchBlockedUsers() {
      try {
       const response = await axios.get(`/profiles/${this.userID}/banlist`);
       console.log(response.data.blockedUsers);
       return response.data.blockedUsers || []; // Assumendo che l'API restituisca un array di ID utente
     } catch (e) {
       console.error("Errore nel recuperare gli utenti bloccati:", e);
       return []; // In caso di errore, restituisce un array vuoto
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
      this.userPhotos = [];
      this.followedUsernames = {};

      const path = `/profiles/${this.userID}/stream`;
      const response = await axios.get(path);

      if (response.data) {
        let allPhotos = response.data.map((photo) => {
          if (photo.image && !photo.image.startsWith("data:image")) {
            photo.image = `data:image/jpeg;base64,${photo.image}`;
          }
          return photo;
        });

        const blockedUsers = await this.fetchBlockedUsers();

        this.userPhotos = allPhotos.filter((photo) => !blockedUsers.includes(photo.user_id));

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
      this.userPhotos = this.userPhotos.map((photo) => {
        if (photo.photo_id === photoId) {
          photo.image = "https://loremflickr.com/320/240";
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
        <RouterLink :to="`/profiles/${photo.user_id}/photos/${photo.photo_id}`">
          <img
            :src="photo.image"
            alt="Photo"
            class="photo"
            @error="onImageError(photo.photo_id)"
          />
        </RouterLink>

        <div class="photo-info">
          <RouterLink :to="`/profiles/${photo.user_id}/profile`" class="photo-username">
            <strong>{{ followedUsernames[photo.user_id] || "Loading..." }}</strong>
          </RouterLink>

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
  font-family: 'Courier New', Courier, monospace;
  max-width: 800px;
  margin: 0 auto;
  background-color: #F4E1C1;
  border-radius: 15px;
  box-shadow: 0 6px 15px rgba(0, 0, 0, 0.2);
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
  background: #FFF8DC;
  border: 2px solid #D67F00;
  border-radius: 12px;
  box-shadow: 0 6px 12px rgba(0, 0, 0, 0.15);
  padding: 15px;
  text-align: center;
  transition: transform 0.3s ease;
}

.photo-card:hover {
  transform: scale(1.05);
}

.photo {
  width: 100%;
  height: auto;
  object-fit: cover;
  border-radius: 8px;
  border: 2px solid #D67F00;
  transition: transform 0.3s ease;
}

.photo:hover {
  transform: scale(1.1);
}

.photo-info {
  margin-top: 10px;
}

.photo-username,
.photo-timestamp,
.photo-caption {
  margin: 5px 0;
  font-style: italic;
  font-size: 14px;
  color: #3E2723;
}

.photo-username {
  display: block;
  color: #D67F00;
  text-decoration: none;
  font-weight: bold;
  cursor: pointer;
  transition: color 0.3s ease;
}

.photo-username:hover {
  color: #F2A541;
  text-decoration: underline;
}

.photo-timestamp,
.photo-caption {
  font-size: 12px;
  color: #8B4513;
}

.photo-timestamp strong,
.photo-caption strong {
  font-weight: 700;
  color: #3E2723;
}
</style>
