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
      errorMessage: '' // Error message for display
    };
  },
  methods: {
    async checkIfFollowing() {
      // The ID of the profile the current user is visiting
      const profileUserId = this.id

      // Get the current user's ID from localStorage
      const currentUserId = getId();

      try {
        // Make a request to fetch the list of users that the current user follows
        const response = await axios.get(`/profiles/${currentUserId}/following`);

        if (response.status === 200) {
          const followedUsers = response.data || [];

          // Extract only the user_ids from the followed users
          const followedUserIds = followedUsers.map(user => user.user_id);

          // Check if the profileUserId is in the list of followed user_ids
          this.isFollowing = followedUserIds.includes(Number(profileUserId));

          console.log(`Is following: ${this.isFollowing}`);
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

          const { user, photos, followers, following, isFollowing, isBanned } = response.data;
          this.user = user;
          this.username = user.username;
          this.inputValue = this.username;
          this.followersCount = followers;
          this.followingCount = following;
          this.isFollowing = isFollowing;
          this.isBanned = isBanned;

          // Check if the current user is the owner
          this.isOwner = this.currentUserId === user.user_id;
          this.checkIfFollowing();
          this.checkIfBanned();
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
        this.errorMessage = 'Failed to fetch user profile. Please try again later.';
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
          localStorage.setItem("username", this.username);
          this.errorMessage = '';
        }
      } catch (e) {
        console.error('Failed to update profile:', e);
        this.errorMessage = "Failed to update profile: the username must be long 3 to 16 chars";
        this.inputValue = this.username; // Revert to old username
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

    async toggleFollow() {
      const path = `/profiles/${this.currentUserId}/following/${this.id}`;
      const method = this.isFollowing ? 'DELETE' : 'PUT';

      try {
        const response = await axios({
          method,
          url: path
        });

        if (response.status === (this.isFollowing ? 200 : 201)) {
          this.isFollowing = !this.isFollowing;
          this.followersCount += this.isFollowing ? 1 : -1;
          this.errorMessage = '';
        }
      } catch (e) {
        console.error(`Failed to ${this.isFollowing ? 'unfollow' : 'follow'} user:`, e);
        this.errorMessage = `Failed to ${this.isFollowing ? 'unfollow' : 'follow'} user. Please try again later.`;
      }
    },

    // async toggleBan() {
    //   const path = `/profiles/${this.currentUserId}/bans/${this.id}`;
    //   const method = this.isBanned ? 'DELETE' : 'PUT';
    //   const followPath = `/profiles/${this.currentUserId}/following/${this.id}`;
    //   const followMethod = 'DELETE';

    //   try {
    //     // Make the ban request
    //     const banResponse = await axios({
    //       method,
    //       url: path
    //     });
    
    //     if (banResponse.status === (this.isBanned ? 200 : 201)) {
    //       this.isBanned = !this.isBanned;
    //       this.errorMessage = '';

    //       // If the user was banned, remove the follow relationship
    //       if (this.isBanned) {
    //         try {
    //           console.log("BANNED");
    //           const followResponse = await axios({
    //             method: followMethod,
    //             url: followPath
    //           });

    //           if (followResponse.status === 200) {
    //             console.log("UNFOLLOWED")
    //             this.isFollowing = false;
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
    async toggleBan() {
      
      const path = `/profiles/${this.currentUserId}/bans/${this.id}`;
      const method = this.isBanned ? 'DELETE' : 'PUT';
      const followPath = `/profiles/${this.id}/following/${this.currentUserId}`;
      const followMethod = 'DELETE';

      try {
        // Make the ban request
        const banResponse = await axios({
          method,
          url: path
        });
        console.log(banResponse.status);
        if (banResponse.status === 200) {
          // Update the banned status immediately
          
          this.isBanned = !this.isBanned;
          this.errorMessage = '';

          // If the user was banned, remove the follow relationship
          if (this.isBanned) {
            try {
              const followResponse = await axios({
                method: followMethod,
                url: followPath
              }); 
              if (followResponse.status === 200) {
                // Update the following status and count
                this.followersCount--;
                console.log('Successfully unfollowed the user');
              } else {
                console.warn(`Unexpected response status from unfollow API: ${followResponse.status}`);
              }
            } catch (error) {
              console.error('Error removing follow relationship:', error);
              this.errorMessage = 'Failed to unfollow user after banning.';
            }
          }
        }
      } catch (e) {
        console.error(`Failed to ${this.isBanned ? 'unban' : 'ban'} user:`, e);
        this.errorMessage = `Failed to ${this.isBanned ? 'unban' : 'ban'} user. Please try again later.`;
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
    <!-- Display Error Message -->
    <div v-if="errorMessage" class="alert alert-danger" role="alert">
      {{ errorMessage }}
    </div>

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

          <!-- Conditional rendering based on ownership -->
          <template v-if="isOwner">
            <!-- Upload photo form -->
            <div class="upload-form mt-4">
              <input type="file" @change="onFileChange" />
              <textarea v-model="caption" placeholder="Add a caption..."></textarea>
              <button class="btn btn-primary" @click="uploadPhoto">Upload Photo</button>
            </div>
          </template>

          <template v-else>
            <!-- Follow and Ban buttons -->
            <button v-if="!isFollowing" class="btn btn-outline-primary follow-btn" @click="toggleFollow">Follow</button>
            <button v-else class="btn btn-outline-secondary follow-btn" @click="toggleFollow">Unfollow</button>

            <button v-if="!isBanned" class="btn btn-outline-primary ban-btn" @click="toggleBan">Ban</button>
            <button v-else class="btn btn-outline-secondary ban-btn" @click="toggleBan">Unban</button>
          </template>

          <div class="stats mt-3">
            <span class="mr-3">posts: <strong>{{ postsCount }}</strong></span>
            <br>
            <span class="mr-3">followers: <strong>{{ followersCount }}</strong></span>
            <br>
            <span class="mr-3">following: <strong>{{ followingCount }}</strong></span>
          </div>
        </div>
      </div>
    </div>

    <div class="tab-content" id="profileTabsContent">
      <div class="tab-pane fade show active" id="posts" role="tabpanel" aria-labelledby="posts-tab">
        <div class="row mt-4">
          <div v-for="post in posts" :key="post.photo_id" class="col-md-4 mb-4">
           <img :src="post.image" class="img-fluid photo" :alt="post.caption" @click="toggleModal(post)">
          </div>

        </div>
      </div>
    </div>
  </div>
</template>


<style scoped>
.profile-page {
  max-width: 800px; /* Container width */
  margin: 0 auto; /* Center container horizontally */
  font-family: 'Arial', sans-serif;
  padding: 20px; /* Padding around the container */
  display: flex;
  flex-direction: column;
  align-items: center; /* Center content horizontally */
}

.profile-header {
  width: 100%; /* Full width of the container */
  max-width: 800px; /* Ensure it doesn't exceed the container's width */
  display: flex;
  flex-direction: column; /* Stack profile details vertically */
  align-items: center; /* Center profile details horizontally */
  gap: 20px; /* Spacing between elements */
  padding: 10px;
  border: 1px solid #0033cc;
  border-radius: 8px;
  background: #e0e0e0;
  box-shadow: 2px 2px 5px #888;
}

.profile-details {
  width: 100%; /* Full width within the profile header */
  max-width: 600px; /* Limit width for better alignment */
  text-align: center; /* Center text inside the profile details */
}

.profile-name-unselected,
.profile-name-selected {
  font-size: 18px; /* Font size for profile name */
  font-weight: bold;
  border: 1px solid #0033cc;
  padding: 5px;
  border-radius: 3px;
  display: inline-block; /* Fit content width */
}

.profile-name-unselected {
  background-color: #d0d0d0;
  color: #000;
}

.profile-name-selected {
  background-color: #a0a0a0;
  color: #000;
}

.stats {
  font-size: 14px; /* Font size for stats */
  color: #000;
  margin-top: 10px;
  text-align: center; /* Center text inside stats */
}

.upload-form {
  width: 100%; /* Full width within the profile header */
  max-width: 600px; /* Limit width for better alignment */
  margin-top: 20px;
  background: #e0e0e0;
  padding: 10px;
  border: 1px solid #0033cc;
  border-radius: 8px;
  box-shadow: 2px 2px 5px #888;
}

.upload-form input[type="file"],
.upload-form textarea,
.upload-form button {
  display: block;
  width: calc(100% - 22px); /* Full width minus padding */
  margin: 0 auto 10px auto; /* Center elements horizontally */
  border: 1px solid #0033cc;
  border-radius: 3px;
}

.upload-form textarea {
  padding: 5px;
  background: #fff;
  color: #000;
}

.upload-form button {
  background: #0033cc;
  color: #fff;
  border: none;
  padding: 5px;
  cursor: pointer;
  font-weight: bold;
}

.upload-form button:hover {
  background: #0055ff;
}

.photo {
  width: 600px;
  height: 240px;
  border: 1px solid #0033cc;
  border-radius: 3px;
  transition: filter 0.3s ease;
  display: block;
}


.photo:hover {
  filter: brightness(1.2);
}
</style>



