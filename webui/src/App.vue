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
			  <span>General</span>
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
  .profile-page {
	max-width: 800px; /* Container width */
	margin: 0 auto; /* Center container horizontally */
	font-family: 'Arial', sans-serif;
	padding: 20px; /* Padding around the container */
	display: flex;
	flex-direction: column;
	align-items: center; /* Center content horizontally */
	background-color: #d0d0d0; /* Light gray background color */
	border: 1px solid #a0a0a0; /* Optional: subtle border for a more classic look */
  }
  
  .profile-header {
	width: 100%; /* Full width of the container */
	max-width: 800px; /* Ensure it doesn't exceed the container's width */
	display: flex;
	flex-direction: column; /* Stack profile details vertically */
	align-items: center; /* Center profile details horizontally */
	gap: 20px; /* Spacing between elements */
	padding: 10px;
	border: 1px solid #0033cc;
	border-radius: 8px;
	background: #e0e0e0; /* Slightly lighter gray for contrast */
	box-shadow: 2px 2px 5px #888;
  }
  
  .profile-details {
	width: 100%; /* Full width within the profile header */
	max-width: 600px; /* Limit width for better alignment */
	text-align: center; /* Center text inside the profile details */
  }
  
  .profile-name-unselected,
  .profile-name-selected {
	font-size: 18px; /* Font size for profile name */
	font-weight: bold;
	border: 1px solid #0033cc;
	padding: 5px;
	border-radius: 3px;
	display: inline-block; /* Fit content width */
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
	font-size: 14px; /* Font size for stats */
	color: #000;
	margin-top: 10px;
	text-align: center; /* Center text inside stats */
  }
  
  .upload-form {
	width: 100%; /* Full width within the profile header */
	max-width: 600px; /* Limit width for better alignment */
	margin-top: 20px;
	background: #e0e0e0; /* Slightly lighter gray for contrast */
	padding: 10px;
	border: 1px solid #0033cc;
	border-radius: 8px;
	box-shadow: 2px 2px 5px #888;
  }
  
  .upload-form input[type="file"],
  .upload-form textarea,
  .upload-form button {
	display: block;
	width: calc(100% - 22px); /* Full width minus padding */
	margin: 0 auto 10px auto; /* Center elements horizontally */
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
	width: 100%; /* Ensure photos are responsive */
	max-width: 100%; /* Prevent overflow */
  }
  
  .photo:hover {
	filter: brightness(1.2);
  }
  </style>
  
  
  