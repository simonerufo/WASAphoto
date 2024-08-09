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
	// Check if the current route is the session route
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
	<header class="navbar navbar-dark sticky-top bg-dark flex-md-nowrap p-0 shadow">
	  <a class="navbar-brand col-md-3 col-lg-2 me-0 px-3 fs-6" href="#/">WASAphoto</a>
	  <button class="navbar-toggler position-absolute d-md-none collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#sidebarMenu" aria-controls="sidebarMenu" aria-expanded="false" aria-label="Toggle navigation">
		<span class="navbar-toggler-icon"></span>
	  </button>
	</header>
	<div class="container-fluid">
	  <div class="row">
		<!-- Conditionally render the sidebar based on the route -->
		<nav id="sidebarMenu" v-if="!isSessionPage" class="col-md-3 col-lg-2 d-md-block bg-light sidebar collapse">
		  <div class="position-sticky pt-3 sidebar-sticky">
			<h6 class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-4 mb-1 text-muted text-uppercase">
			  <!-- <span>** MENU **</span> -->
			</h6> 
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
				<div class="nav-link logout-button" role="button" @click="logout" style="color: red;">
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
  /* General Container */
  .container-fluid {
	background-color: #f0f0f0; /* Light grey background to match profile page */
	font-family: 'Arial', sans-serif;
  }
  
  /* Sidebar Styling */
  #sidebarMenu {
	background: #e0e0e0; /* Match profile page background */
	border-right: 1px solid #0033cc; /* Consistent border color */
	box-shadow: 2px 0 5px #888; /* Subtle shadow for sidebar */
	padding: 10px;
  }
  
  .sidebar-sticky {
	position: relative;
  }
  
  .sidebar-heading {
	font-size: 14px;
	font-weight: bold;
	color: #0033cc;
	text-transform: uppercase;
	padding: 10px;
	background-color: #d0d0d0; /* Light gray for the sidebar heading */
	border-bottom: 1px solid #0033cc;
	border-radius: 3px;
	text-align: center;
  }
  
  .nav-item .nav-link {
	color: #000;
	font-weight: bold;
	background-color: #d0d0d0;
	margin: 5px 0;
	padding: 8px 15px;
	border-radius: 3px;
	text-align: left;
	transition: background-color 0.3s ease, box-shadow 0.3s ease;
	border: 1px solid #0033cc;
	box-shadow: 2px 2px 5px #888;
  }
  
  .nav-item .nav-link:hover {
	background-color: #a0a0a0;
	box-shadow: 2px 2px 5px #666;
  }
  
  .nav-item .logout-button {
	color: #fff;
	font-weight: bold;
	background-color: #d9534f; /* Bootstrap danger color */
	margin: 5px 0;
	padding: 8px 15px;
	border-radius: 3px;
	text-align: left;
	transition: background-color 0.3s ease, box-shadow 0.3s ease;
	border: 1px solid #d9534f;
	box-shadow: 2px 2px 5px #888;
  }
  
  .nav-item .logout-button:hover {
	background-color: #c9302c; /* Darker shade on hover */
	box-shadow: 2px 2px 5px #666;
  }
  
  /* Header Styling */
  .navbar {
	background-color: #0033cc; /* Dark blue for the header */
	border-bottom: 1px solid #0033cc;
	box-shadow: 0 2px 5px #888;
	padding: 10px 15px;
  }
  
  .navbar-brand {
	color: #fff;
	font-family: 'Arial', sans-serif;
	font-size: 18px;
	font-weight: bold;
  }
  
  .navbar-toggler-icon {
	filter: invert(100%);
  }
  
  /* Main Content Area */
  .main {
	background-color: #f0f0f0; /* Light grey background for consistency */
	padding: 20px;
	border-left: 1px solid #0033cc;
  }
  
  /* Profile Page Style (unchanged) */
  .profile-page {
	max-width: 800px;
	margin: 0 auto;
	font-family: 'Arial', sans-serif;
	padding: 20px;
	display: flex;
	flex-direction: column;
	align-items: center;
  }
  
  .profile-header {
	width: 100%;
	max-width: 800px;
	display: flex;
	flex-direction: column;
	align-items: center;
	gap: 20px;
	padding: 10px;
	border: 1px solid #0033cc;
	border-radius: 8px;
	background: #e0e0e0;
	box-shadow: 2px 2px 5px #888;
  }
  
  .profile-details {
	width: 100%;
	max-width: 600px;
	text-align: center;
  }
  
  .profile-name-unselected,
  .profile-name-selected {
	font-size: 18px;
	font-weight: bold;
	border: 1px solid #0033cc;
	padding: 5px;
	border-radius: 3px;
	display: inline-block;
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
	font-size: 14px;
	color: #000;
	margin-top: 10px;
	text-align: center;
  }
  
  .upload-form {
	width: 100%;
	max-width: 600px;
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
	width: calc(100% - 22px);
	margin: 0 auto 10px auto;
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
	border: 1px solid #0033cc;
	border-radius: 3px;
	transition: filter 0.3s ease;
	width: 100%;
	max-width: 100%;
  }
  
  .photo:hover {
	filter: brightness(1.2);
  }
  </style>
  