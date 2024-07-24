<script>
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
      isInputEnabled: true, 
      posts: [
        { id: 1, image: '../public/marco1.jpg', caption: 'Caption 1' },
        { id: 1, image: '../public/marco1.jpg', caption: 'Caption 1' },
        { id: 1, image: '../public/marco1.jpg', caption: 'Caption 1' },
        { id: 2, image: '../public/marco1.jpg', caption: 'Caption 2' },
      ],
    };
  },

  methods: {
    async getUser() {
      var path = `/profiles/${this.id}/profile`;
      try {
        var response = await this.$axios.get(path);
        if (response.status === 200) {
          this.user = response.data.user;
          this.username = response.data.user.username;
          this.inputValue = this.username; 
          this.followersCount = response.data.followers;
          this.followingCount = response.data.following;
        }
      } catch (e) {
        console.error('Failed to fetch user:', e);
      }
    },

    enableInput() {
      console.log('Input enabled');
      this.isInputEnabled = true; 
    },

    onBlur() {
      if (this.isInputEnabled) {
        console.log('Input blurred, saving...');
        this.editProfile(); 
      }
    },

    onEnter() {
      if (this.isInputEnabled) {
        console.log('Enter pressed, saving...');
        this.editProfile(); 
      }
    },

    async editProfile() {
      if (!this.isInputEnabled) return; 

      this.loading = true;
      this.errormsg = null;
      var path = `/profiles/${this.id}/profile`;
      try {
        let response = await this.$axios.put(path, {
          username: this.inputValue, // Send the updated username
        });
        if (response.status === 200) {
          this.username = this.inputValue; // Update local username
          this.user.username = this.inputValue; // Update user object
        }
      } catch (e) {
        this.errormsg = 'Failed to update profile: ' + e.toString();
        this.inputValue = this.username; // Revert changes on error
      } finally {
        this.loading = false;
        this.isInputEnabled = false; // Disable input after saving
      }
    },
  },

  mounted() {
    this.getUser();
  }
}
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
                   :disabled="!isInputEnabled"
                   @click="enableInput"
                   @blur="onBlur"
                   @keydown.enter="onEnter">
          </div>
          <button class="btn btn-outline-primary follow-btn" v-if="!isFollowing">Follow</button>
          <button class="btn btn-outline-secondary follow-btn" v-else>Unfollow</button>
          <button class="btn btn-outline-primary ban-btn" v-if="!isBanned">Ban</button>
          <button class="btn btn-outline-secondary ban-btn" v-else>Unban</button>
          <div class="stats mt-3">
            <span class="mr-3">  posts:  <strong> {{ postsCount }} </strong></span>
            <span class="mr-3">  followers:  <strong> {{ followersCount }} </strong></span>
            <span class="mr-3">  following:  <strong> {{ followingCount }} </strong></span>
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
            <img :src="post.image" class="img-fluid" :alt="post.caption">
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
  align-items: top;
  background-color: #3CBC8D;
  color: white;
}

.profile-name-unselected::placeholder {
  opacity:1
}

.profile-name-selected {
  font-size: 20px;
  font-weight: bold;
  border: 2px solid green;
  align-items: top;
  background-color: #3CBC8D;
  color: white;
}

.profile-name-selected::placeholder {
  opacity:1
}

.follow-btn {
  
}

.stats {
  font-size: 16px;
}

</style>