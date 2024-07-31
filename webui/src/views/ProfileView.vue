<script>
import axios from '../services/axios';
import Modal from '../components/PostModal.vue';
//import UsersListModal from '../components/UsersListModal.vue';
//import NewPostModal from '../components/NewPostModal.vue';

export default {
  components: {
    Modal,
    //UsersListModal,
    //NewPostModal
  },
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

      // Destructure response data
      const { user, photos, followers, following } = response.data;
      
      // Update component data
      this.user = user;
      this.username = user.username;
      this.inputValue = this.username;
      this.followersCount = followers;
      this.followingCount = following;
      this.posts = photos; 
      
      console.log('Updated component data:', {
        user: this.user,
        username: this.username,
        followersCount: this.followersCount,
        followingCount: this.followingCount,
        posts: this.posts
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
          localStorage["username"] = this.username
        }
      } catch (e) {
        console.error('Failed to update profile:', e);
        this.inputValue = this.username;
      } finally {
        this.isInputEnabled = false;
      }
    },

    async toggleFollow() {
      if (this.isFollowing) {
        await this.unfollowUser();
      } else {
        await this.followUser();
      }
    },

    async followUser() {
      try {
        const response = await axios.put(`/profiles/${this.id}/followed/${this.user.id}`);
        if (response.status === 200) {
          this.isFollowing = true;
          this.followersCount++;
        }
      } catch (e) {
        console.error('Failed to follow user:', e);
      }
    },

    async unfollowUser() {
      try {
        const response = await axios.delete(`/profiles/${this.id}/followed/${this.user.id}`);
        if (response.status === 200) {
          this.isFollowing = false;
          this.followersCount--;
        }
      } catch (e) {
        console.error('Failed to unfollow user:', e);
      }
    },

    async toggleBan() {
      if (this.isBanned) {
        await this.unbanUser();
      } else {
        await this.banUser();
      }
    },

    async banUser() {
      try {
        const response = await axios.put(`/profiles/${this.id}/bans/${this.user.id}`);
        if (response.status === 200) {
          this.isBanned = true;
          this.isFollowing = false;
          this.followersCount--;
        }
      } catch (e) {
        console.error('Failed to ban user:', e);
      }
    },

    async unbanUser() {
      try {
        const response = await axios.delete(`/profiles/${this.id}/bans/${this.user.id}`);
        if (response.status === 200) {
          this.isBanned = false;
        }
      } catch (e) {
        console.error('Failed to unban user:', e);
      }
    },

    toggleModal(post) {
      this.selectedPost = post;
    },

    toggleFollowModal(index) {
      this.followPath = index;
    },

    toggleCreateModal(reload) {
      this.createPost = !this.createPost;
      if (reload) {
        this.getPhotos(); // Assuming you have this method to refresh photos
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
          // Update posts list and reset form
          this.posts.push({ 
            id: response.data.pid, 
            image: URL.createObjectURL(this.file), 
            caption: this.caption 
          });
          this.caption = '';
          this.file = null;
        }
      } catch (e) {
        console.error('Failed to upload photo:', e);
      }
    }
  },
  mounted() {
    this.getUser();
  }
};
</script>


<template>
  <div class="container profile-page">
    <!-- Profile Header -->
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
          <button class="btn btn-outline-primary follow-btn" v-if="!isFollowing" @click="toggleFollow">Follow</button>
          <button class="btn btn-outline-secondary follow-btn" v-else @click="toggleFollow">Unfollow</button>
          <button class="btn btn-outline-primary ban-btn" v-if="!isBanned" @click="toggleBan">Ban</button>
          <button class="btn btn-outline-secondary ban-btn" v-else @click="toggleBan">Unban</button>
          <div class="stats mt-3">
            <span class="mr-3">posts: <strong>{{ postsCount }}</strong></span>
            <span class="mr-3">followers: <strong>{{ followersCount }}</strong></span>
            <span class="mr-3">following: <strong>{{ followingCount }}</strong></span>
          </div>
          <!-- Upload Photo Form -->
          <div class="upload-form mt-4">
            <input type="file" @change="onFileChange" />
            <textarea v-model="caption" placeholder="Add a caption..."></textarea>
            <button class="btn btn-primary" @click="uploadPhoto">Upload Photo</button>
          </div>
        </div>
      </div>
    </div>

    <!-- Tabs -->
    <div class="tab-content" id="profileTabsContent">
      <div class="tab-pane fade show active" id="posts" role="tabpanel" aria-labelledby="posts-tab">
        <!-- Posts Grid -->
        <div class="row mt-4">
          <div v-for="post in posts" :key="post.id" class="col-md-4 mb-4">
            <img :src="post.image" class="img-fluid" :alt="post.caption" @click="toggleModal(post)">
          </div>
        </div>
      </div>
    </div>
  </div>

  <Modal v-if="selectedPost" @close="toggleModal(null)" :photo="selectedPost"></Modal>
  <!--<UsersListModal v-if="followPath" @close="toggleFollowModal(null)" :list_mode="followPath" :username="user.username" :user_id="user.id" :bans="bans"></UsersListModal>
  <NewPostModal v-if="createPost" @close="toggleCreateModal(false)" @closeReload="toggleCreateModal(true)" :username="user.username"></NewPostModal>
  -->
</template>






<style scoped>
/*
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
*/
</style>