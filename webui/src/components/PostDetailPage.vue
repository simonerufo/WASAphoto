<script>
import axios, { getId } from '../services/axios';

export default {
  data() {
    return {
      post: null,
      likeCount: 0,
      liked: false,
      commentText: '',
      comments: [],
      isOwner: false // To track if the current user owns the profile
    };
  },
  async mounted() {
    const postId = this.$route.params.photo_id;
    const userId = this.$route.params.user_id;

    try {
      await this.checkOwnership(userId); // Check if the current user is the owner
      await this.fetchPostDetails(userId, postId);
      await this.fetchComments(userId, postId);
    } catch (e) {
      console.error('Failed to fetch data:', e);
    }
  },
  methods: {
    async checkOwnership(profileUserId) {
      this.isOwner = getId() === profileUserId;
      console.log(this.isOwner);
      console.log(Number(getId()),profileUserId);
    },

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

      if (!this.isOwner) {
        alert('You do not have permission to delete this post');
        return;
      }

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
      const userId = localStorage.getItem("id");

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
          this.comments = this.comments.filter(comment => comment.comment_id !== commentId);
        }
      } catch (e) {
        console.error('Failed to remove comment:', e);
        alert('Failed to remove comment');
      }
    },
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
      <button v-if="isOwner" @click="deletePost">Delete Post</button>
    </div>
    <div>
      <h2>Comments</h2>
      <textarea v-model="commentText" placeholder="Add a comment"></textarea>
      <button @click="addComment">Post Comment</button>
    </div>
    <div v-for="comment in comments" :key="comment.comment_id">
      <p><strong>{{ comment.username }}:</strong> {{ comment.comment_text }} <em>{{ comment.timestamp }}</em></p>
      <button v-if="isOwner" @click="removeComment(comment.comment_id)">Remove</button>
    </div>
  </div>
</template>

<style scoped>
.post-details {
  max-width: 800px; /* Container width */
  margin: 0 auto; /* Center container horizontally */
  padding: 20px;
  background-color: #d0d0d0; /* Light gray background color */
  border: 1px solid #a0a0a0; /* Subtle border for classic look */
  border-radius: 5px;
  font-family: 'Arial', sans-serif;
  text-align: center;
}

.post-details h1 {
  font-size: 24px;
  color: #000;
  margin-bottom: 20px;
}

.post-details img {
  max-width: 100%;
  height: auto;
  border: 1px solid #a0a0a0;
  border-radius: 5px;
  margin-bottom: 10px;
}

.post-details p {
  font-size: 16px;
  color: #000;
}

button {
  background-color: #0033cc; /* Classic blue button color */
  color: #fff;
  border: 1px solid #002a80; /* Darker blue border */
  border-radius: 3px;
  padding: 10px;
  cursor: pointer;
  font-size: 16px;
  margin: 5px;
  transition: background-color 0.3s; /* Smooth transition on hover */
}

button:hover {
  background-color: #0055ff; /* Lighter blue on hover */
}

textarea {
  width: 100%;
  height: 100px;
  padding: 5px;
  border: 1px solid #a0a0a0;
  border-radius: 3px;
  background-color: #fff;
  color: #000;
}

.comment {
  background-color: #e0e0e0; /* Slightly lighter gray for comments */
  border: 1px solid #a0a0a0;
  border-radius: 3px;
  padding: 10px;
  margin-top: 10px;
  text-align: left;
}

.comment p {
  margin: 0;
}

.comment button {
  background-color: #cc0000; /* Red color for remove button */
  border: 1px solid #990000; /* Darker red border */
}

.comment button:hover {
  background-color: #ff0000; /* Brighter red on hover */
}
</style>

