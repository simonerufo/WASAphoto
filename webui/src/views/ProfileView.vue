<script>
import axios, { getId } from '../services/axios'; // Adjust the path if needed

export default {
  data() {
    return {
      user: null,
      username: null,
      inputValue: '',
      id: this.$route.params.user_id,
      currentUserId: Number(getId()), // Store the current user ID
      isFollowing: false,
      isBanned: false,
      postsCount: 0,
      followersCount: null,
      followingCount: null,
      isInputEnabled: false,
      selectedPost: null,
      createPost: false,
      posts: [],
      caption: '',
      file: null,
      isOwner: false, // Track if the current user is the owner
      errorMessage: '', // Error message for display
      successMessage: '' // Success message for displaying username change
    };
  },
  methods: {

  async checkIfFollowing() {
    const profileUserId = this.id;
    const currentUserId = getId();

    try {
      const response = await axios.get(`/profiles/${currentUserId}/following`);

      if (response.status === 200) {
        const followedUsers = response.data || [];
        const followedUserIds = followedUsers.map(user => user.user_id);
      
          this.isFollowing = followedUserIds.includes(Number(profileUserId));

        console.log(`Following status for user ${profileUserId}: ${this.isFollowing}`);
      } else {
        console.warn(`Unexpected response status ${response.status}`);
      }
    } catch (error) {
      console.error("Error checking if the user is following the profile:", error);
    }
  },

    async checkIfBanned() {
      const profileUserId = this.id;
      const currentUserId = getId();

      try {
        const response = await axios.get(`profiles/${currentUserId}/ban`);

        if (response.status === 200) {
          const bannedUserIDs = response.data.banned_user_ids;

          console.log("Response Data:", bannedUserIDs);

          // Ensure bannedUserIDs is an array before using .some()
          if (Array.isArray(bannedUserIDs)) {
            this.isBanned = bannedUserIDs.some(bannedId => Number(bannedId) === Number(profileUserId));
          } else {
            console.warn("bannedUserIDs is not an array:", bannedUserIDs);
            this.isBanned = false;
          }

          console.log("isBanned:", this.isBanned);
        } else {
          console.warn(`Unexpected response status ${response.status}`);
        }
      } catch (error) {
        console.error("Error checking if the user is banned:", error);
      }
    },


  
    async getUser() {
      const path = `/profiles/${this.id}/profile`;
      console.log(`Fetching user profile from: ${path}`);
      try {
        const response = await axios.get(path);
        console.log(`Response status: ${response.status}`);

        if (response.status === 200) {
          console.log('User data:', response.data);

          const { user, photos = [], followers = 0, following = 0, isFollowing = false, isBanned = false } = response.data;
          this.user = user;
          this.username = user.username;
          this.inputValue = this.username;
          this.followersCount = followers;
          this.followingCount = following;
          this.isFollowing = isFollowing;
          this.isBanned = isBanned;

          // Check if the current user is the owner
          this.isOwner = this.currentUserId === user.user_id;

          // Only sort photos if the array has elements
          if (Array.isArray(photos) && photos.length > 0) {
            this.posts = photos.sort((a, b) => new Date(b.timestamp) - new Date(a.timestamp));
          } else {
            this.posts = []; // Handle the case when there are no photos
          }

          this.postsCount = this.posts.length;

          console.log('Updated component data:', {
            user: this.user,
            username: this.username,
            followersCount: this.followersCount,
            followingCount: this.followingCount,
            posts: this.posts,
            postsCount: this.postsCount
          });

          // Fetch follow and ban status after getting user data
          await this.checkIfFollowing();
          await this.checkIfBanned();
        }
      } catch (e) {
        console.error('Failed to fetch user:', e);
        this.errorMessage = 'Failed to fetch user profile. Please try again later.';
      }
    },


    

    enableInput() {
      this.isInputEnabled = true;
      this.successMessage = '';
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
          localStorage.setItem("username", this.username);
          this.errorMessage = '';
          this.successMessage = 'Username updated successfully!';
          this.isInputEnabled = false;
        }
      } catch (e) {
        console.error('Failed to update profile:', e);
        this.errorMessage = "Failed to update profile: the username must be long 3 to 16 chars and must have only lower case chars";
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
        this.errorMessage = 'Invalid post ID. Please try again.';
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
      if (!this.file) {
        this.errorMessage = 'Please select a file to upload.';
        return;
      }

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
          const { photo_id } = response.data;

          const post = {
            photo_id,
            image: URL.createObjectURL(this.file),
            caption: this.caption,
            timestamp: new Date().toISOString()
          };

          this.posts.unshift(post);
          this.postsCount++;
          this.caption = '';
          this.file = null;
          this.errorMessage = ''; // Clear any previous error
        }
      } catch (e) {
        console.error('Failed to upload photo:', e);
        this.errorMessage = 'Failed to upload photo. Please try again later.';
      }
    },

    // async toggleFollow() {
    //   const path = `/profiles/${this.currentUserId}/following/${this.id}`;
    //   const method = this.isFollowing ? 'DELETE' : 'PUT';

    //   try {
    //     const response = await axios({
    //       method,
    //       url: path
    //     });

    //     if (response.status === (this.isFollowing ? 200 : 201)) {
    //       this.isFollowing = !this.isFollowing;
    //       this.followersCount += this.isFollowing ? 1 : -1;
    //       this.errorMessage = '';

    //       // Reload the user profile to update UI
    //       await this.getUser();
    //     }
    //   } catch (e) {
    //     console.error(`Failed to ${this.isFollowing ? 'unfollow' : 'follow'} user:`, e);
    //     this.errorMessage = `Failed to ${this.isFollowing ? 'unfollow' : 'follow'} user. Please try again later.`;
    //   }
    // },

  async toggleFollow() {
    const path = `/profiles/${this.currentUserId}/following/${this.id}`;
    const method = this.isFollowing ? 'DELETE' : 'PUT';

    try {
      const response = await axios({
        method,
        url: path
      });

    if (response.status === 201 || response.status === 200) {
        this.isFollowing = !this.isFollowing;
    
        this.followersCount += this.isFollowing ? 1 : -1;

        console.log(`User ${this.isFollowing ? 'followed' : 'unfollowed'} successfully`);
      } else {
        console.warn(`Unexpected response status ${response.status}`);
      }
    } catch (e) {
      console.error(`Failed to ${this.isFollowing ? 'unfollow' : 'follow'} user:`, e);
      this.errorMessage = `Failed to ${this.isFollowing ? 'unfollow' : 'follow'} user. Please try again later.`;
    }
  },


    // async toggleBan() {

    //   const path = `/profiles/${this.currentUserId}/bans/${this.id}`;
    //   const method = this.isBanned ? 'DELETE' : 'PUT';
    //   const followPath = `/profiles/${this.id}/following/${this.currentUserId}`;
    //   const followMethod = 'DELETE';

    //   try {
    //     // Make the ban request
    //     console.log(`Ban API Path: ${path}`);
    //     console.log(`Unfollow API Path: ${followPath}`);
    //     console.log(`Current User ID: ${this.currentUserId}, Target User ID: ${this.id}`);

    //     const banResponse = await axios({
    //       method,
    //       url: path
    //     });
    //     console.log(banResponse.status);
    //     if (banResponse.status === 200) {
    //       // Update the banned status immediately
          
    //       this.isBanned = !this.isBanned;
    //       this.errorMessage = '';

    //       // If the user was banned, remove the follow relationship
    //       if (this.isBanned) {
    //         try {
    //           const followResponse = await axios({
    //             method: followMethod,
    //             url: followPath
    //           }); 
    //           if (followResponse.status === 200) {
    //             // Update the following status and count
    //             this.followersCount--;
    //             console.log('Successfully unfollowed the user');
    //           } else {
    //             console.warn(`Unexpected response status from unfollow API: ${followResponse.status}`);
    //           }
    //         } catch (error) {
    //           console.error('Error removing follow relationship:', error);
    //           this.errorMessage = 'Failed to unfollow user after banning.';
    //         }
    //       }
    //     }
    //   } catch (e) {
    //     console.error(`Failed to ${this.isBanned ? 'unban' : 'ban'} user:`, e);
    //     this.errorMessage = `Failed to ${this.isBanned ? 'unban' : 'ban'} user. Please try again later.`;
    //   }
    // }
  // async toggleBan() {
  //   const path = `/profiles/${this.currentUserId}/bans/${this.id}`;
  //   const method = this.isBanned ? 'DELETE' : 'PUT';
  //   const followPath = `/profiles/${this.id}/following/${this.currentUserId}`;
  //   const followMethod = 'DELETE';

  //   console.log(`Ban request: ${method} ${path}`);
  //   console.log(`Unfollow request: ${followMethod} ${followPath}`);

  //   try {
  //     const banResponse = await axios({ method, url: path });

  //     console.log(`Ban Response Status: ${banResponse.status}`);

  //     if (banResponse.status === 200) {
  //       this.isBanned = !this.isBanned;

  //     if (this.isBanned) {
  //         try {
  //           const followResponse = await axios({ method: followMethod, url: followPath });
  //           console.log(`Unfollow Response Status: ${followResponse.status}`);
  //           if (followResponse.status === 200) {
  //             this.followersCount--;
  //           }
  //         } catch (error) {
  //           console.error('Error removing follow relationship:', error.response?.data || error.message);
  //         }
  //       }
  //     }
  //   } catch (e) {
  //     console.error(`Failed to ${this.isBanned ? 'unban' : 'ban'} user:`, e.response?.data || e.message);
  //   }
  // }
  async toggleBan() {
    const path = `/profiles/${this.currentUserId}/bans/${this.id}`;
    const method = this.isBanned ? 'DELETE' : 'PUT';
    const followPath = `/profiles/${this.id}/following/${this.currentUserId}`;
    const followMethod = 'DELETE';

    console.log(`Ban request: ${method} ${path}`);
    console.log(`Unfollow request: ${followMethod} ${followPath}`);

    try {
      const banResponse = await axios({ method, url: path });

      console.log(`Ban Response Status: ${banResponse.status}`);

      if (banResponse.status === 200 || banResponse.status == 204) {
        console.log("Banned before:", this.isBanned);
        this.isBanned = !this.isBanned;
        console.log("Banned after:", this.isBanned);


        if (this.isBanned && this.isFollowing) {
          try {
            const followResponse = await axios({ method: followMethod, url: followPath });
            console.log(`Unfollow Response Status: ${followResponse.status}`);

            if (followResponse.status === 200) {
              this.followersCount--;
            }
          } catch (error) {
            console.error('Error removing follow relationship:', error.response?.data || error.message);
          }
        }
      }
    } catch (e) {
      console.error(`Failed to ${this.isBanned ? 'unban' : 'ban'} user:`, e.response?.data || e.message);
    }
  }

  },
  mounted() {
    this.getUser();
  }
};
</script>

<template>
  <div class="container profile-page text-center">
    <div v-if="errorMessage" class="alert alert-danger" role="alert">
      {{ errorMessage }}
    </div>
    <div v-if="successMessage" class="alert alert-success" role="alert">
      {{ successMessage }}
    </div>

    <div class="profile-header mt-4">
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

        <template v-if="isOwner">
          <button v-if="isInputEnabled" class="btn btn-primary" @click="editProfile">Save Username</button>
          <button v-else class="btn btn-outline-primary" @click="enableInput">Edit Username</button>

          <div class="upload-form mt-4">
            <input type="file" @change="onFileChange" />
            <textarea v-model="caption" placeholder="Add a caption..."></textarea>
            <button class="btn btn-primary" @click="uploadPhoto">Upload Photo</button>
          </div>
        </template>

        <template v-else>
          <button v-if="!isFollowing" class="btn btn-outline-primary follow-btn" @click="toggleFollow">Follow</button>
          <button v-else class="btn btn-outline-secondary follow-btn" @click="toggleFollow">Unfollow</button>

          <button v-if="!isBanned" class="btn btn-outline-primary ban-btn" @click="toggleBan">Ban</button>
          <button v-else class="btn btn-outline-secondary ban-btn" @click="toggleBan">Unban</button>
        </template>

        <div class="stats mt-3">
          <span class="mr-3">Posts: <strong>{{ postsCount }}</strong></span>
          <br>
          <span class="mr-3">Followers: <strong>{{ followersCount }}</strong></span>
          <br>
          <span class="mr-3">Following: <strong>{{ followingCount }}</strong></span>
        </div>
      </div>
    </div>

    <div class="tab-content" id="profileTabsContent">
      <div class="tab-pane fade show active" id="posts" role="tabpanel" aria-labelledby="posts-tab">
        <div class="photo-grid mt-4">
          <div v-for="post in posts" :key="post.photo_id" class="photo-item">
            <img :src="post.image" class="img-fluid photo" :alt="post.caption" @click="toggleModal(post)">
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.profile-page {
  background-color: #000;
  color: #fff;
  padding: 30px;
  border-radius: 15px;
  text-align: center;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}

.profile-header {
  background-color: #111;
  padding: 20px;
  border-radius: 15px;
  width: 100%;
  max-width: 600px;
  text-align: center;
}

.profile-details input {
  background-color: #222;
  color: #fff;
  border: 2px solid #444;
  padding: 10px;
  border-radius: 10px;
  text-align: center;
  width: 100%;
}

.upload-form textarea {
  background-color: #222;
  color: #fff;
  border: 2px solid #444;
  padding: 10px;
  border-radius: 10px;
  width: 100%;
  min-height: 100px;
  resize: none;
}

.btn {
  background-color: #333;
  color: #fff;
  border: 2px solid #555;
  padding: 10px 20px;
  border-radius: 10px;
  transition: 0.3s;
  display: block;
  width: 100%;
  margin-top: 10px;
}

.btn:hover {
  background-color: #555;
}

.photo-grid {
  display: flex;
  flex-wrap: wrap;
  gap: 15px;
  justify-content: center;
}

.photo-item {
  flex: 1 1 calc(33.33% - 15px);
  max-width: calc(33.33% - 15px);
}

.photo {
  width: 100%;
  height: auto;
  border-radius: 10px;
  transition: transform 0.3s ease;
  border: 3px solid #fff;
}

.photo:hover {
  transform: scale(1.1);
}
</style>

