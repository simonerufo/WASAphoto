<script>
import axios,{ getId , getUsername } from '../services/axios';

export default {
  data() {
    return {
      owner: null,
      post: null,
      likeCount: 0,
      isLiked: false,
      commentText: '',
      comments: [],
      isOwner: false 
    };
  },
  async mounted() {
    const postId = this.$route.params.photo_id;
    const userId = this.$route.params.user_id;

    try {
      await this.getUsername();
      await this.checkOwnership(userId);
      await this.checkIfLikedPhoto();
      await this.fetchPostDetails(userId, postId);
      await this.fetchComments(userId, postId);
    } catch (e) {
      console.error('Failed to fetch data:', e);
    }
  },
  methods: {
  async fetchLikeCount(postId) {
    try {
      const response = await axios.get(`/photos/${postId}/likes`);
      if (response.status === 200) {
        console.log("likes",response.data);
        this.likeCount = response.data.length;  // Set the like count
      } else {
        console.warn(`Unexpected response status ${response.status}`);
      }
    } catch (e) {
      console.error('Failed to fetch like count:', e);
    }
  },

  async checkOwnership(profileUserId) {
    this.isOwner = getId() === profileUserId;
  },

  async checkIfLikedPhoto() {
    const postId = this.$route.params.photo_id;
    const userId = getId(); // Get the current user ID

    try {
      const response = await axios.get(`/photos/${postId}/likes`);
      if (response.status === 200) {
        const likes = response.data;
        const liked = likes.map(user => user.user_id);
        this.isLiked = liked.includes(Number(userId)); // Check if current user ID is in the list of likes
        console.log(this.isLiked);
        console.log(liked);
      } else {
        console.warn(`Unexpected response status ${response.status}`);
      }
    } catch (e) {
      console.error('Failed to check if photo is liked:', e);
    }
  },

  async fetchPostDetails(userId, postId) {
    try {
      const response = await axios.get(`/profiles/${userId}/photos/${postId}`);
      if (response.status === 200) {
        const postData = response.data;
        this.post = postData;
        console.log(this.post.image);
        this.checkIfLikedPhoto(); // Check if the photo is liked by the current user
        await this.fetchLikeCount(postId); // Fetch the like count
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
      } else {
        console.warn(`Unexpected response status ${response.status}`);
      }
    } catch (e) {
      console.error('Failed to fetch comments:', e);
    }
  },

  async toggleLikePhoto() {
    const postId = this.$route.params.photo_id;
    const userId = getId();

    try {
      if (this.isLiked) {
        const response = await axios.delete(`/profiles/${userId}/likes/${postId}`);
        if (response.status === 200) {
          this.likeCount -= 1;
          this.isLiked = false;
          console.log("Unliked");
        }
      } else {
        const response = await axios.put(`/profiles/${userId}/likes/${postId}`);
        if (response.status === 204) {
          this.likeCount += 1;
          this.isLiked = true;
          console.log("liked");
        }
      }
      // Update like count after liking/unliking
      await this.fetchLikeCount(postId);
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
        this.likeCount = 0;
        this.comments = [];
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

  // async removeComment(commentId) {
  //   const userId = this.$route.params.user_id;

  //   try {
  //     const response = await axios.delete(`/profiles/${userId}/comments/${commentId}`);
  //     if (response.status === 200) {
  //       this.comments = this.comments.filter(comment => comment.comment_id !== commentId);
  //     }
  //   } catch (e) {
  //     console.error('Failed to remove comment:', e);
  //     alert('Failed to remove comment');
  //   }
  // },
  getID(){
    getId();
  },

  async removeComment(commentId) {
    const username = getUsername();
    const userId = getId();
    const comment = this.comments.find(comment => comment.comment_id === commentId); // Find the comment
    console.log("comment_id:",commentId)
    // Check if the logged-in user is the author of the comment or the author of the post
    if (comment.username !== username  && !this.isOwner) {
      alert('You do not have permission to delete this comment');
      return;
    }
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
  },


  formatTimestamp(timestamp) {
    const date = new Date(timestamp);
    const options = {
      year: "numeric",
      month: "short",
      day: "2-digit",
      hour: "2-digit",
      minute: "2-digit",
      second: "2-digit",
    };
    return date.toLocaleDateString("en-US", options);
  },

  async getUsername() {
      const userID = this.$route.params.user_id;
      const path = `/profiles/${userID}/profile`;
      console.log("Fetching username for userID:", userID);
      try {
        const response = await axios.get(path);
        console.log("API response received for userID:", userID);
        if (response.status === 200) {
          this.owner = response.data.user.username;
        } else {
          console.warn("Unexpected response status ${response.status} for userID ${userID}");
        }
      } catch (error) {
        console.error("Error fetching username for userID:", userID, error);
      }
    },

},
}
</script>

<template>
  <div class="post-details">
    <div v-if="post">
      <img :src="post.image" alt="Post image" />
      <p><strong>{{ owner }}:</strong> {{ post.caption }}</p>
      <p>Posted on: {{ formatTimestamp(post.timestamp) }}</p>
      <button @click="toggleLikePhoto">
        {{ isLiked ? 'Unlike' : 'Like' }} ({{ likeCount }})
      </button>
      <button v-if="isOwner" @click="deletePost">Delete Post</button>
    </div>
    <div>
      <h2>Comments</h2>
      <textarea v-model="commentText" placeholder="Add a comment"></textarea>
      <button @click="addComment">Post Comment</button>
    </div>
    <div v-for="comment in comments" :key="comment.comment_id">
      <p><strong>{{ comment.username }} <em> [{{ formatTimestamp(comment.timestamp) }}]: </em></strong> {{ comment.comment_text }} </p>
      <button v-if="isOwner || comment.user_id === getID()" @click="removeComment(comment.comment_id)">Remove</button>
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

.photo-container {
  display: flex;
  justify-content: center; /* Center horizontally */
  align-items: center; /* Center vertically */
  width: 100%;
  height: auto;
  margin-bottom: 20px;
}

.post-details img {
  max-width: 400px;  /* Maximum width */
  min-width: 200px;  /* Minimum width */
  max-height: 300px; /* Maximum height */
  min-height: 150px; /* Minimum height */
  width: 100%; /* Ensure photo scales with the container */
  height: auto; /* Maintain aspect ratio */
  object-fit: cover; /* Ensures image fills the box, cropping if necessary */
  border: 1px solid #0033cc;
  border-radius: 3px;
  transition: filter 0.3s ease;
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
