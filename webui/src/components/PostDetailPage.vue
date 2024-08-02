<script>
import axios from '../services/axios';

export default {
  data() {
    return {
      post: null,
      likeCount: 0,
      liked: false,
      commentText: '',
      comments: []
    };
  },
  async mounted() {
    const postId = this.$route.params.photo_id;
    const userId = this.$route.params.user_id;

    try {
      await this.fetchPostDetails(userId, postId);
      await this.fetchComments(userId, postId);
    } catch (e) {
      console.error('Failed to fetch data:', e);
    }
  },
  methods: {
    async fetchPostDetails(userId, postId) {
      try {
        const response = await axios.get(`/profiles/${userId}/photos/${postId}`);
        if (response.status === 200) {
          const postData = response.data;
          this.post = postData;
          this.likeCount = postData.likeCount || 0;
          this.liked = postData.liked || false;
        }
      } catch (e) {
        console.error('Failed to fetch post details:', e);
      }
    },

    async fetchComments(userId, postId) {
      try {
        const response = await axios.get(`/profiles/${userId}/photos/${postId}/comments`);
        if (response.status === 200) {
          this.comments = response.data || [];
        }
      } catch (e) {
        console.error('Failed to fetch comments:', e);
      }
    },

    async toggleLikePhoto() {
      const postId = this.$route.params.photo_id;
      const userId = this.$route.params.user_id;

      try {
        if (this.liked) {
          const response = await axios.delete(`/profiles/${userId}/likes/${postId}`);
          if (response.status === 200) {
            this.likeCount -= 1;
            this.liked = false;
          }
        } else {
          const response = await axios.put(`/profiles/${userId}/likes/${postId}`);
          if (response.status === 204) {
            this.likeCount += 1;
            this.liked = true;
          }
        }
      } catch (e) {
        console.error('Failed to like/unlike post:', e);
        alert('Failed to like/unlike post');
      }
    },

    async deletePost() {
      const postId = this.$route.params.photo_id;
      const userId = this.$route.params.user_id;

      try {
        const response = await axios.delete(`/profiles/${userId}/photos/${postId}`);
        if (response.status === 204) {
          alert('Post deleted successfully');
          this.$router.push(`/profiles/${userId}/profile`);
        }
      } catch (e) {
        console.error('Failed to delete post:', e);
        alert('Failed to delete post');
      }
    },

    async addComment() {
      const postId = this.$route.params.photo_id;
      const userId = this.$route.params.user_id;

      try {
        const response = await axios.post(`/profiles/${userId}/comments/${postId}`, {
          comment_text: this.commentText
        });

        if (response.status === 200) {
          const newComment = response.data;
          this.comments.push(newComment);
          this.commentText = '';
        }
      } catch (e) {
        console.error('Failed to post comment:', e);
        alert('Failed to post comment');
      }
    },

    async removeComment(commentId) {
      const userId = this.$route.params.user_id;

      try {
        const response = await axios.delete(`/profiles/${userId}/comments/${commentId}`);
        if (response.status === 200) {
          // Remove the comment from the list
          this.comments = this.comments.filter(comment => comment.comment_id !== commentId);
        }
      } catch (e) {
        console.error('Failed to remove comment:', e);
        alert('Failed to remove comment');
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
      <button @click="toggleLikePhoto">
        {{ liked ? 'Unlike' : 'Like' }} ({{ likeCount }})
      </button>
      <button @click="deletePost">Delete Post</button>
    </div>
    <div>
      <h2>Comments</h2>
    </div>
    <textarea v-model="commentText" placeholder="Add a comment"></textarea>
    <button @click="addComment">Post Comment</button>
    <div v-for="comment in comments" :key="comment.comment_id">
      <p><strong>{{ comment.username }}:</strong> {{ comment.comment_text }} <em>{{ comment.timestamp }}</em></p>
      <button @click="removeComment(comment.comment_id)">Remove</button>
    </div>
  </div>
</template>


<style scoped>
.post-details {
  text-align: center;
}

.post-details img {
  max-width: 100%;
  height: auto;
}

textarea {
  width: 100%;
  height: 100px;
}
</style>
