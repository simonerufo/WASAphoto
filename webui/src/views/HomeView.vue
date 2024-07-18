<script>
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			currentUsername: null,
		}
	},
	methods: {
		async refresh() {
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.get("/");
				this.some_data = response.data;
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},
	},
	mounted() {
		//this.refresh();
		this.currentUsername = localStorage["username"];
		console.log(localStorage["id"]);
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
	</div>
</template>

<style>
</style>
