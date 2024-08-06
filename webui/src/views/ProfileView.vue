<script>
import axios from '../services/axios';

export default {
  data() {
    return {
      user: null,
      username: null,
      inputValue: '',
      id: this.$route.params.user_id,
      isFollowing: false,
      isBanned: false,
      postsCount: 0,
      followersCount: null,
      followingCount: null,
      isInputEnabled: false,
      selectedPost: null,
      followPath: null,
      createPost: false,
      posts: [],
      caption: '',
      file: null
    };
  },
  methods: {
    async getUser() {
      const path = `/profiles/${this.id}/profile`;
      console.log(`Fetching user profile from: ${path}`);

      try {
        const response = await axios.get(path);
        console.log(`Response status: ${response.status}`);

        if (response.status === 200) {
          console.log('User data:', response.data);

          const { user, photos, followers, following, isFollowing } = response.data;

          this.user = user;
          this.username = user.username;
          this.inputValue = this.username;
          this.followersCount = followers;
          this.followingCount = following;
          this.isFollowing = isFollowing;
          // Sort photos by timestamp (newest first)
          this.posts = photos.sort((a, b) => new Date(b.timestamp) - new Date(a.timestamp));
          this.postsCount = this.posts.length;

          console.log('Updated component data:', {
            user: this.user,
            username: this.username,
            followersCount: this.followersCount,
            followingCount: this.followingCount,
            posts: this.posts,
            postsCount: this.postsCount
          });
        }
      } catch (e) {
        console.error('Failed to fetch user:', e);
      }
    },

    enableInput() {
      this.isInputEnabled = true;
    },

    onBlur() {
      if (this.isInputEnabled) {
        this.editProfile();
      }
    },

    onEnter() {
      if (this.isInputEnabled) {
        this.editProfile();
      }
    },

    async editProfile() {
      if (!this.isInputEnabled) return;

      try {
        const response = await axios.put(`/profiles/${this.id}/username`, { username: this.inputValue });
        if (response.status === 200) {
          this.username = this.inputValue;
          this.user.username = this.inputValue;
          localStorage["username"] = this.username;
        }
      } catch (e) {
        console.error('Failed to update profile:', e);
        this.inputValue = this.username;
      } finally {
        this.isInputEnabled = false;
      }
    },

    toggleModal(post) {
      if (post && post.photo_id) {
        this.$router.push({ path: `/profiles/${this.id}/photos/${post.photo_id}` });
      } else {
        console.error('Post ID is invalid', post);
      }
    },

    toggleCreateModal(reload) {
      this.createPost = !this.createPost;
      if (reload) {
        this.getUser();
      }
    },

    onFileChange(event) {
      this.file = event.target.files[0];
    },

    async uploadPhoto() {
      const formData = new FormData();
      formData.append('image', this.file);
      formData.append('caption', this.caption);

      try {
        const response = await axios.post(`/profiles/${this.id}/profile`, formData, {
          headers: {
            'Content-Type': 'multipart/form-data'
          }
        });

        if (response.status === 201) {
          console.log(response.data);
          const { photo_id } = response.data; // Ensure this field is returned from the backend

          const post = {
            photo_id, // Use the correct field name
            image: URL.createObjectURL(this.file),
            caption: this.caption,
            timestamp: new Date().toISOString()
          };

          // Insert the new photo at the beginning of the array
          this.posts.unshift(post);
          this.postsCount++;
          this.caption = '';
          this.file = null;
        }
      } catch (e) {
        console.error('Failed to upload photo:', e);
      }
    },
    async toggleFollow() {
      const path = `/profiles/${this.id}/following/${this.user.id}`;
      const method = this.isFollowing ? 'DELETE' : 'PUT';
    
      try {
        const response = await axios({
          method,
          url: path
        });

        if (response.status === (this.isFollowing ? 200 : 201)) {
          this.isFollowing = !this.isFollowing;
          this.followersCount += this.isFollowing ? 1 : -1;
        }
      } catch (e) {
        console.error(`Failed to ${this.isFollowing ? 'unfollow' : 'follow'} user:`, e);
      }
    },
  },
  mounted() {
    this.getUser();
  }
};
</script>


<template>
  <div class="container profile-page">
    <div class="row profile-header mt-4">
      <div class="col-md-8">
        <div class="profile-details">
          <div class="form-group">
            <input type="text"
                   id="username"
                   :class="{'profile-name-unselected': !isInputEnabled, 'profile-name-selected': isInputEnabled}"
                   v-model="inputValue"
                   @click="enableInput"
                   @blur="onBlur"
                   @keydown.enter="onEnter">
          </div>
          <button class="btn btn-outline-primary follow-btn" 
                  v-if="!isFollowing" 
                  @click="toggleFollow">Follow</button>
          <button class="btn btn-outline-secondary follow-btn" 
                  v-else 
                  @click="toggleFollow">Unfollow</button>
          <!--<button class="btn btn-outline-primary follow-btn" v-if="!isFollowing" @click="toggleFollow">Follow</button>-->
          <button class="btn btn-outline-secondary follow-btn" v-else @click="toggleFollow">Unfollow</button>
          <button class="btn btn-outline-primary ban-btn" v-if="!isBanned" @click="toggleBan">Ban</button>
          <button class="btn btn-outline-secondary ban-btn" v-else @click="toggleBan">Unban</button>
          <div class="stats mt-3">
            <span class="mr-3">posts: <strong>{{ postsCount }}</strong></span>
            <span class="mr-3">followers: <strong>{{ followersCount }}</strong></span>
            <span class="mr-3">following: <strong>{{ followingCount }}</strong></span>
          </div>
          <div class="upload-form mt-4">
            <input type="file" @change="onFileChange" />
            <textarea v-model="caption" placeholder="Add a caption..."></textarea>
            <button class="btn btn-primary" @click="uploadPhoto">Upload Photo</button>
          </div>
        </div>
      </div>
    </div>
    <div class="tab-content" id="profileTabsContent">
      <div class="tab-pane fade show active" id="posts" role="tabpanel" aria-labelledby="posts-tab">
        <div class="row mt-4">
          <div v-for="post in posts" :key="post.id" class="col-md-4 mb-4">
            <img :src="post.image" class="img-fluid photo" :alt="post.caption" @click="toggleModal(post)">
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.profile-page {
  max-width: 900px;
  margin: 0 auto;
}

.profile-header {
  display: flex;
  align-items: center;
}

.profile-details {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 10px;
  padding: 10px;
  border: 1px solid #000;
  border-radius: 10px;
}

.profile-name-unselected {
  font-size: 20px;
  font-weight: bold;
  border: 2px solid red;
  background-color: #3CBC8D;
  color: white;
}

.profile-name-selected {
  font-size: 20px;
  font-weight: bold;
  border: 2px solid green;
  background-color: #3CBC8D;
  color: white;
}

.stats {
  font-size: 16px;
}

.upload-form {
  margin-top: 20px;
}

.upload-form input[type="file"] {
  margin-bottom: 10px;
}

.upload-form textarea {
  width: 100%;
  margin-bottom: 10px;
  padding: 5px;
}

.upload-form button {
  display: block;
  width: 100%;
}

.photo {
  transition: filter 0.3s ease;
}

.photo:hover {
  filter: brightness(1.5);
}
</style>
