<script>
import { RouterLink, RouterView } from 'vue-router'

export default {
  data() {
    return {
      logged: null,
      link: "",
      username: ""
    }
  },
  computed: {
    isSessionPage() {
      return this.$route.path === '/session';
    }
  },
  watch: {
    $route(to, from) {
      this.updateUserData()
    }
  },
  mounted() {
    this.updateUserData()
  },
  methods: {
    updateUserData() {
      this.logged = localStorage.getItem("id")
      this.username = localStorage.getItem("username")
      if (this.logged) {
        this.link = `/profiles/${this.logged}/profile`
      }
    },
    logout() {
      localStorage.removeItem("id");
      localStorage.removeItem("username");
      this.$router.push("/");
    }
  }
}
</script>

<template>
  <header class="navbar navbar-expand-lg navbar-dark bg-dark shadow-sm">
    <a class="navbar-brand" href="#/">WASAphoto</a>
    <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#sidebarMenu" aria-controls="sidebarMenu" aria-expanded="false" aria-label="Toggle navigation">
      <span class="navbar-toggler-icon"></span>
    </button>
  </header>
  <div class="container-fluid">
    <div class="row">
      <nav id="sidebarMenu" v-if="!isSessionPage" class="col-md-3 col-lg-2 d-md-block bg-light sidebar collapse">
        <div class="position-sticky pt-3">
          <ul class="nav flex-column">
            <li class="nav-item">
              <RouterLink to="/home" class="nav-link">
                <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#home"/></svg>
                Home
              </RouterLink>
            </li>
            <li class="nav-item">
              <RouterLink to="/search" class="nav-link">
                <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#compass"/></svg>
                Search
              </RouterLink>
            </li>
            <li class="nav-item">
              <RouterLink :to="link" class="nav-link">
                <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#image"/></svg>
                Profile: {{ username }}
              </RouterLink>
            </li>
            <li class="nav-item">
              <div class="nav-link logout-button" role="button" @click="logout">
                <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#key"/></svg>
                Logout
              </div>
            </li>
          </ul>
        </div>
      </nav>

      <main class="col-md-9 ms-sm-auto col-lg-10 px-md-4">
        <RouterView />
      </main>
    </div>
  </div>
</template>

<style scoped>
/* General Styles */
body {
  font-family: 'Roboto', sans-serif;
  background-color: #f8f9fa;
  color: #495057;
  margin: 0;
  padding: 0;
}

.container-fluid {
  background-color: #ffffff;
}

.navbar {
  background-color: #343a40;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.navbar-brand {
  color: #fff;
  font-weight: bold;
  font-size: 24px;
}

.navbar-toggler-icon {
  filter: invert(100%);
}

/* Sidebar Styles */
#sidebarMenu {
  background-color: #ffffff;
  border-right: 1px solid #e9ecef;
}

.nav-item .nav-link {
  color: #495057;
  font-weight: 600;
  padding: 10px 20px;
  margin-bottom: 8px;
  border-radius: 4px;
  transition: background-color 0.3s ease, box-shadow 0.3s ease;
}

.nav-item .nav-link:hover {
  background-color: #e9ecef;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.nav-item .logout-button {
  background-color: #dc3545;
  color: white;
  padding: 10px 20px;
  border-radius: 4px;
  cursor: pointer;
  transition: background-color 0.3s ease, box-shadow 0.3s ease;
  margin-top: 20px;
}

.nav-item .logout-button:hover {
  background-color: #c82333;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

/* Main Content */
main {
  padding-top: 20px;
}

/* Profile Page Style */
.profile-page {
  max-width: 800px;
  margin: 0 auto;
  padding: 20px;
}

.profile-header {
  background-color: #ffffff;
  border: 1px solid #e9ecef;
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.1);
  padding: 20px;
  border-radius: 8px;
}

.profile-details {
  text-align: center;
}

.stats {
  font-size: 14px;
  margin-top: 10px;
}

.upload-form {
  background-color: #ffffff;
  border: 1px solid #e9ecef;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.1);
}

.upload-form input[type="file"],
.upload-form textarea,
.upload-form button {
  width: 100%;
  margin-bottom: 15px;
  border: 1px solid #e9ecef;
  border-radius: 4px;
}

.upload-form button {
  background-color: #007bff;
  color: white;
  font-weight: bold;
  padding: 10px;
  cursor: pointer;
  border: none;
}

.upload-form button:hover {
  background-color: #0056b3;
}

/* Styling for photos */
.photo {
  border-radius: 5px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  width: 100%;
  max-width: 100%;
}

.photo:hover {
  filter: brightness(1.1);
}
</style>
