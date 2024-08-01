<script>
import axios from '../services/axios';

export default {
  data() {
    return {
      post: null,
      likeCount: 0,
      liked: false
    };
  },
  async mounted() {
    const postId = this.$route.params.photo_id;
    const userId = this.$route.params.user_id;

    try {
      await this.fetchPostDetails(userId, postId);
      await this.fetchLikeCount(postId);
    } catch (e) {
      console.error('Failed to fetch data:', e);
    }
  },
  methods: {
    async fetchPostDetails(userId, postId) {
      try {
        const response = await axios.get(`/profiles/${userId}/photos/${postId}`);
        if (response.status === 200) {
          this.post = response.data;
          this.likeCount = response.data.like_count || 0;
          this.liked = response.data.liked_by_user || false;
        }
      } catch (e) {
        console.error('Failed to fetch post details:', e);
      }
    },
    
    async fetchLikeCount(postId) {
      try {
        const response = await axios.get(`/photos/${postId}/likes`);
        if (response.status === 200) {
          this.likeCount = response.data.length || 0;
        }
      } catch (e) {
        console.error('Failed to fetch like count:', e);
      }
    },
    
    async likePhoto() {
      const postID = this.$route.params.photo_id;
      const userID = this.$route.params.user_id;

      try {
        if (this.liked) {
          await axios.delete(`/profiles/${userID}/likes/${postID}`);
          this.likeCount -= 1;
          this.liked = false;
        } else {
          // Like the post
          await axios.put(`/profiles/${userID}/likes/${postID}`);
          this.likeCount += 1;
          this.liked = true;
        }
      } catch (e) {
        console.error('Failed to like/unlike post:', e);
        alert('Failed to like/unlike post');
      }
    },
    
    async deletePost() {
      const postID = this.$route.params.photo_id;
      const userID = this.$route.params.user_id;

      try {
        const response = await axios.delete(`/profiles/${userID}/photos/${postID}`);
        if (response.status === 204) {
          alert('Post deleted successfully');
          this.$router.push(`/profiles/${userID}/profile`);
        }
      } catch (e) {
        console.error('Failed to delete post:', e);
        alert('Failed to delete post');
      }
    }
  }
};
</script>

<template>
  <div class="post-details">
    <div v-if="post">
      <h1>{{ post.username }}</h1>
      <img :src="post.image" alt="Post image" />
      <p>{{ post.caption }}</p>
      <p>Posted on: {{ post.timestamp }}</p>
      <button @click="likePhoto">{{ liked ? 'Unlike' : 'Like' }} ({{ likeCount }})</button>
      <button @click="deletePost">Delete Post</button>
    </div>
  </div>
</template>

<style></style>
