<script>
import axios from "../services/axios.js"
export default {
  data() {
    return {
      errormsg: null,
      loading: false,
      currentUsername: null,
      searchUsername: null,
      searched: null,
    };
  },
  methods: {
    async refresh() {
      this.loading = true;
      this.errormsg = null;
      try {
        let response = await axios.get("/");
        this.some_data = response.data;
      } catch (e) {
        this.errormsg = e.toString();
      }
      this.loading = false;
    },
    
    async searchUsers() {
      this.loading = true;
      this.errormsg = null;
      try {
        if (!this.searchUsername) {
          throw new Error("Please enter a username.");
        }

        const userID = localStorage.getItem('id');
        if (!userID) {
          throw new Error("User ID not found in local storage");
        }

        const path = `/profiles/${userID}/username`;
		console.log(`Making API call to: ${path}`);
        // Make the API call
        let response = await axios.get(path, {
          params: { username: this.searchUsername }
        });

        if (response.status === 200) {
          // Update the searched result
          this.searched = response.data;
        } else {
          // Handle non-200 responses
		  console.log("A");
          this.errormsg = `Unexpected response status: ${response.status}`;
        }
      } catch (e) {
        // Handle errors from the API or other sources
		console.log("A");
        this.errormsg = e.response && e.response.data ? e.response.data : e.toString();
      }
      this.loading = false;
    }
  },
  mounted() {
    this.currentUsername = localStorage.getItem('username');
  }
}
</script>

<!-- <template>
	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">Home page</h1>
			<div class="btn-toolbar mb-2 mb-md-0">
				<div class="btn-group me-2">
					<button type="button" class="btn btn-sm btn-outline-secondary" @click="refresh">
						Refresh
					</button>
					<button type="button" class="btn btn-sm btn-outline-secondary" @click="exportList">
						Export
					</button>
				</div>
				<div class="btn-group me-2">
					<button type="button" class="btn btn-sm btn-outline-primary" @click="newItem">
						New
					</button>
				</div>
			</div>
		</div>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
	</div>
</template>
-->
<template>
	<div class="stream-page">
	  <div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
		<h2>{{ currentUsername }}'s Stream</h2>
	  </div>
	  
	  <div>
		<input v-model="searchUsername" placeholder="Enter username" />
		<button @click="searchUsers">Search</button>
	  </div>
	  
	  <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
	  
	  <div v-if="searched">
		<h2>User Profile</h2>
		<p><strong>Username:</strong> {{ searched.username }}</p>
	  </div>
	  
	  <div v-if="loading">Loading...</div>
	</div>
  </template>
<!--  
	  <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
	  <div v-else>
		<div v-if="stream.length > 0">
		  <div class="post-container" v-for="post in stream" :key="post.id">
			<div class="user-info d-flex justify-content-between align-items-center">
				<div>
					<RouterLink :to="$pathToProfile(post.username)" class="user-link">
						<span class="post-name">{{ post.username }}</span>
					</RouterLink>
				</div>
				<div class="post-stats">
					{{ post.timestamp }}
					<text class="post-stats-heart"> {{ post.commentCount }} <svg class="feather" style="width: 23px; height: 23px;"><use href="/feather-sprite-v4.29.0.svg#message-circle"/></svg></text> 
					<text>{{ post.likeCount }} <svg class="feather" style="width: 23px; height: 23px;"><use href="/feather-sprite-v4.29.0.svg#heart"/></svg> </text>
				</div>
			</div>
			<div class="post-content">
			  <img role="button" :src="post.image" @click="toggleModal(post)" class="post-image"/>
			  <p v-if="post.caption" class="post-caption">{{ post.caption }}</p>
			</div>
		  </div>
		</div>
		<div v-else>
		  <p class="empty-message">There is nothing to display here. Begin by following someone!</p>
		</div>
	  </div>
	</div>
	<Modal v-if="selectedPost" @close="toggleModal(null)" :photo="selectedPost"></Modal>
-->	

<style>
</style>
