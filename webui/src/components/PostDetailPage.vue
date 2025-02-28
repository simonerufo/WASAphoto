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
        if (response.status === 200 && response.data) {
        console.log("likes", response.data);
        this.likeCount = Array.isArray(response.data) ? response.data.length : 0;
      } else {
        console.warn(`Unexpected response or empty data: ${response.status}`);
        this.likeCount = 0;
      }
    } catch (e) {
      console.error('Failed to fetch like count:', e);
      this.likeCount = 0;
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
      if (response.status === 200 && response.data) {
        const likes = Array.isArray(response.data) ? response.data : [];
        const liked = likes.map(user => user.user_id);
        this.isLiked = liked.includes(Number(userId)); // Check if current user ID is in the list of likes
        console.log(this.isLiked);
        console.log(liked);
      } else {
        this.isLiked = false;
      }
    } catch (e) {
      console.error('Failed to check if photo is liked:', e);
      this.isLiked = false;
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
        this.comments = Array.isArray(response.data) ? response.data : [];
      } else {
        this.comments = [];
      }
    } catch (e) {
      console.error('Failed to fetch comments:', e);
      this.comments = [];
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
      for (const comment of this.comments) {
        const commentId = comment.comment_id; 
        await this.removeComment(commentId);
    }

      const response = await axios.delete(`/profiles/${userId}/photos/${postId}`);
      if (response.status === 204) {
        this.likeCount = 0;
        this.comments = [];
        alert('Post and associated comments deleted successfully');
        this.$router.push(`/profiles/${userId}/profile`);
      }
    } catch (e) {
      console.error('Failed to delete post or comments:', e);
      alert('Failed to delete post or comments');
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
      <img :src="post.image" alt="Post image" class="post-image" />
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
    <div v-for="comment in comments" :key="comment.comment_id" class="comment">
      <p><strong>{{ comment.username }} <em> [{{ formatTimestamp(comment.timestamp) }}]: </em></strong> {{ comment.comment_text }} </p>
      <button v-if="isOwner || comment.user_id === getID()" @click="removeComment(comment.comment_id)">Remove</button>
    </div>
  </div>
</template>

<style scoped>

.post-details {
  max-width: 800px; /* Larghezza massima del contenitore */
  margin: 0 auto; /* Centrare il contenitore orizzontalmente */
  padding: 20px;
  background-color: #fffbe1; /* Giallo chiaro vintage per lo sfondo */
  border: 2px solid #f8a300; /* Arancione per il bordo */
  border-radius: 10px; /* Angoli arrotondati per il contenitore */
  font-family: 'Courier New', Courier, monospace; /* Font retro */
  text-align: center;
  box-shadow: 3px 3px 5px rgba(0, 0, 0, 0.2); /* Ombra per effetto 3D */
}

.post-details h1 {
  font-size: 24px;
  color: #f8a300; /* Arancione caldo */
  margin-bottom: 20px;
}

.photo-container {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  height: auto;
  margin-bottom: 20px;
}

.post-details img {
  max-width: 400px;
  min-width: 200px;
  max-height: 300px;
  min-height: 150px;
  width: 100%;
  height: auto;
  object-fit: cover;
  border: 2px solid #f8a300; /* Bordo arancione intorno all'immagine */
  border-radius: 5px;
  transition: filter 0.3s ease;
}

button {
  background-color: #f8a300; /* Arancione per i pulsanti */
  color: white;
  border: 1px solid #ff6f00; /* Arancione più scuro per il bordo */
  border-radius: 5px;
  padding: 10px;
  cursor: pointer;
  font-size: 16px;
  margin: 10px;
  transition: background-color 0.3s ease;
  font-family: 'Courier New', Courier, monospace; /* Font retro */
}

button:hover {
  background-color: #ffcc00; /* Giallo per hover */
  color: #000; /* Colore del testo cambia a nero */
}

textarea {
  width: 100%;
  height: 100px;
  padding: 8px;
  border: 2px solid #f8a300;
  border-radius: 5px;
  background-color: #fffbe1; /* Giallo chiaro per la textarea */
  color: #000;
  font-family: 'Courier New', Courier, monospace;
}

.comment {
  background-color: #fffbe1; /* Giallo chiaro per i commenti */
  border: 1px solid #f8a300;
  border-radius: 5px;
  padding: 10px;
  margin-top: 10px;
  text-align: left;
}

.comment p {
  margin: 0;
  color: #f8a300; /* Arancione per il testo dei commenti */
}

.comment button {
  background-color: #ff6f00; /* Arancione scuro per il pulsante di rimozione */
  border: 1px solid #e65100;
  color: white;
}

.comment button:hover {
  background-color: #ff3d00; /* Arancione chiaro su hover */
}
</style>


