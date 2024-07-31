<script>
export default {
  name: 'PostModal',
  props: ["post"],
  data() {
    return {
      postLocal: this.post,
      errormsg: null,
      loading: false,
      comments: [],
      likes: [],
      currentUser: this.$store.state.currentUser,
      commentText: null,
      postLiked: false,
    };
  },
  watch: {
    post: {
      immediate: true,
      handler(newVal, oldVal) {
        if (newVal !== oldVal) {
          this.postLocal = newVal;
          this.loadComments();
          this.loadLikes();
        }
      }
    }
  },
  methods: {
    close() {
      this.$emit('close');
    },
    async loadComments() {
      this.loading = true;
      this.errormsg = null;
      let path = `/api/posts/${this.postLocal.post_id}/comments`;
      try {
        let response = await this.$axios.get(path);
        if (response.status === 200) {
          this.comments = response.data;
        }
      } catch (e) {
        this.errormsg = e.toString();
      }
      this.loading = false;
    },
    async loadLikes() {
      this.loading = true;
      this.errormsg = null;
      let path = `/api/posts/${this.postLocal.post_id}/likes`;
      try {
        let response = await this.$axios.get(path);
        if (response.status === 200) {
          this.likes = response.data;
          this.postLiked = this.likes.some(like => like.user_id === this.currentUser.user_id);
        }
      } catch (e) {
        this.errormsg = e.toString();
      }
      this.loading = false;
    },
    async sendComment() {
      this.loading = true;
      this.errormsg = null;
      let path = `/api/posts/${this.postLocal.post_id}/comments`;
      let payload = { text: this.commentText };
      try {
        let response = await this.$axios.post(path, payload);
        if (response.status === 200) {
          this.commentText = null;
          this.loadComments();
        }
      } catch (e) {
        this.errormsg = e.toString();
      }
      this.loading = false;
    },
    async deleteComment(commentID, index) {
      this.loading = true;
      this.errormsg = null;
      let path = `/api/posts/${this.postLocal.post_id}/comments/${commentID}`;
      try {
        let response = await this.$axios.delete(path);
        if (response.status === 200) {
          this.comments.splice(index, 1);
        }
      } catch (e) {
        this.errormsg = e.toString();
      }
      this.loading = false;
    },
    async toggleLike() {
      this.currentUser = this.$store.state.currentUser;
      let path = `/api/posts/${this.postLocal.post_id}/likes/${this.currentUser.user_id}`;
      this.loading = true;
      this.errormsg = null;
      try {
        if (this.postLiked) {
          let response = await this.$axios.delete(path);
          if (response.status === 200) {
            this.postLiked = false;
            this.postLocal.likesNum--;
            this.loadLikes();
          }
        } else {
          let response = await this.$axios.put(path);
          if (response.status === 200) {
            this.postLiked = true;
            this.postLocal.likesNum++;
            this.loadLikes();
          }
        }
      } catch (e) {
        this.errormsg = e.toString();
      }
      this.loading = false;
    },
    pathToProfile(username) {
      return `/profile/${username}`;
    },
    formatDate(date) {
      return new Date(date).toLocaleString();
    }
  }
};
</script>

<template>
  <div class="modal" @click="close">
    <div class="modal-inner" @click.stop="">
      <div class="container">
        <div class="row">
          <div class="col-md-6">
            <!-- Mostra l'immagine del post -->
            <img :src="post.imageURL" alt="Photo" class="img-fluid" ref="image">
          </div>
          <div class="col-md-6">
            <div class="photo-details-container" ref="detailsContainer">
              <!-- Dettagli del post -->
              <div class="d-flex justify-content-between">
                <RouterLink :to="pathToProfile(post.user.username)" class="user-link">
                  <h4>{{ post.user.username }}</h4>
                </RouterLink>
                <svg class="feather" role="button" @click="close"><use href="/feather-sprite-v4.29.0.svg#x"/></svg>
              </div>
              <p v-if="post.caption">{{ post.caption }}</p>
              <div class="d-flex justify-content-between align-items-center" style="margin-bottom: 5px;">
                <div class="d-flex justify-content-between align-items-center" role="button">
                  <!-- Toggle like per il post -->
                  <svg @click="toggleLike" xmlns="http://www.w3.org/2000/svg" width="22" height="22" fill="currentColor" class="bi bi-heart-fill text-danger" viewBox="0 0 16 16">
                    <path v-if="postLiked" fill-rule="evenodd" d="M8 1.314C12.438-3.248 23.534 4.735 8 15-7.534 4.736 3.562-3.248 8 1.314"/>
                    <path v-else d="m8 2.748-.717-.737C5.6.281 2.514.878 1.4 3.053c-.523 1.023-.641 2.5.314 4.385.92 1.815 2.834 3.989 6.286 6.357 3.452-2.368 5.365-4.542 6.286-6.357.955-1.886.838-3.362.314-4.385C13.486.878 10.4.28 8.717 2.01zM8 15C-7.333 4.868 3.279-3.04 7.824 1.143c.06.055.119.112.176.171a3.12 3.12 0 0 1 .176-.17C12.72-3.042 23.333 4.867 8 15"/>
                  </svg>
                </div>
                <small>{{ formatDate(post.timestamp) }}</small>
              </div>
              <div>
                <!-- Tab per commenti e like -->
                <ul class="nav nav-tabs nav-justified">
                  <li class="nav-item" role="presentation">
                    <button class="nav-link active" id="comment-tab" data-bs-toggle="tab" data-bs-target="#comments" type="button" role="tab" aria-controls="home" aria-selected="true">
                      <svg class="feather" style="width: 23px; height: 23px;"><use href="/feather-sprite-v4.29.0.svg#message-circle"/></svg>
                      {{ post.commentsNum }} comments
                    </button>
                  </li>
                  <li class="nav-item" role="presentation">
                    <button class="nav-link" id="likes-tab" data-bs-toggle="tab" data-bs-target="#likes" type="button" role="tab" aria-controls="home" aria-selected="true">
                      <svg class="feather" style="width: 23px; height: 23px;"><use href="/feather-sprite-v4.29.0.svg#heart"/></svg>
                      {{ post.likesNum }} Likes
                    </button>
                  </li>
                </ul>
              </div>
              <div class="tab-content" style="overflow-y: auto;">
                <!-- Contenuto dei commenti -->
                <div id="comments" class="tab-pane fade show active" role="tabpanel">
                  <div class="input-group">
                    <input v-model="commentText" type="text" class="form-control input-comment" placeholder="Write a comment..." aria-label="Write a comment..." aria-describedby="send">
                    <div class="input-group-append">
                      <span @click="sendComment" class="input-group-text text-primary send-button" id="send">
                        <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#send"/></svg>
                      </span>
                    </div>
                  </div>
                  <ul class="list-group">
                    <li class="list-group-item" v-for="(comment, index) in comments" :key="comment.comment_id">
                      <div class="comment-container">
                        <div class="d-flex justify-content-between align-items-center">
                          <RouterLink :to="pathToProfile(comment.user.username)" class="user-link">
                            <strong>{{ comment.user.username }}</strong><br>
                          </RouterLink>
                          <svg v-if="comment.owner_id === currentUser.user_id" class="feather text-danger" role="button" @click="deleteComment(comment.comment_id, index)"><use href="/feather-sprite-v4.29.0.svg#trash-2"/></svg>
                        </div>
                        {{ comment.text }}<br>
                        <small class="form-text font-italic">{{ formatDate(comment.timestamp) }}</small>
                      </div>
                    </li>
                  </ul>
                </div>
                <!-- Contenuto dei like -->
                <div id="likes" class="tab-pane fade" role="tabpanel">
                  <ul class="list-group">
                    <li class="list-group-item" v-for="like in likes" :key="like.user_id">
                      <div class="comment-container">
                        <div class="d-flex">
                          <RouterLink :to="pathToProfile(like.username)" class="user-link">
                            <strong>{{ like.username }}</strong><br>
                          </RouterLink>
                        </div>
                      </div>
                    </li>
                  </ul>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>


<style>
.modal {
  display: flex;
  align-items: center;
  justify-content: center;
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  z-index: 1050;
}

.modal-inner {
  background: white;
  padding: 20px;
  border-radius: 5px;
  max-width: 90%;
  max-height: 90%;
  overflow-y: auto;
}

.photo-details-container {
  display: flex;
  flex-direction: column;
  height: 100%;
}

.input-comment {
  font-size: 0.85rem;
}

.comment-container {
  max-width: 100%;
  word-wrap: break-word;
}
</style>
