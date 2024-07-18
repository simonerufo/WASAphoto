<script>
import UploadPhoto from '../components/UploadPhoto.vue';
export default {
    components: {
    },
    data() {
        return {}
    },
    methods: {},

    mounted() {
        
    },
}

</script>


<template>
	<LoadingSpinner :loading=isLoading />
    <ErrorMsg v-if="errorMsg" :msg="errorMsg" @close-error="errorMsg = ''"></ErrorMsg>
    <UploadPhoto v-if="isEditingPropic" :photoType="'proPic'" @exit-upload-form="isEditingPropic = false"
        @refresh-data="updateProfile" @error-occurred="(value) => { errorMsg = value }" />
    <div class="top-profile-container">
        <div class="top-profile-picture" @mouseover="showEditPropic = isOwner" @mouseleave="showEditPropic = false">
            <div class="edit-propic" v-if="showEditPropic">
                <button class="reset-propic-button">
                    <font-awesome-icon icon="fa-solid fa-xmark" size="lg" color="#fff" @click="resetProfilePic" />
                </button>
                <button class="edit-propic-button" @click="isEditingPropic = true">
                    <font-awesome-icon icon="fa-regular fa-pen-to-square" size="lg" color="#fff" />
                </button>
            </div>
            <img :src="`data:image/jpg;base64,${proPic64}`">
        </div>
        <div class="top-body-profile-container">
            <div class="profile-options-button-container">
                <button class="profile-options-button" @click="showOptions = true" @focusout="closeOptions">
                    <font-awesome-icon icon="fa-solid fa-ellipsis" />
                </button>
                <div v-if="showOptions && isOwner" class="profile-options-menu">
                    <div class="options-menu">
                        <div class="options-menu-item" @click="getBans">
                            <span>Bans list</span>
                        </div>
                        <div class="options-menu-item" @click="deleteProfile">
                            <span>Delete profile</span>
                        </div>
                        <div class="options-menu-item" @click="doLogout">
                            <span>Logout</span>
                        </div>
                    </div>
                </div>
                <div v-else-if="showOptions" class="profile-options-menu">
                    <div class="options-menu-item" @click="banUser">
                        <span>Ban this user</span>
                    </div>
                </div>
            </div>
            <input :readonly="!isOwner" v-model="username" class="top-body-profile-username" @focusin="editingUsername"
                @focusout="saveChangeUsername" @input="checkUsername" maxlength="13" spellcheck="false">
            <div class="top-body-profile-bio-container">
                <span class="top-body-profile-bio-text-counter">{{ textCounter }}/100</span>
                <span class="top-body-profile-bio-label">Bio</span>
                <textarea :readonly="!isOwner" @focusin="editingBio" @focusout="saveChangeBio" @input="countChar"
                    v-model="bio" class="top-body-profile-bio-text" spellcheck="false" maxlength="100" rows="2"></textarea>
            </div>
            <div class="top-body-profile-stats-container">
                <div class="profile-stats" @click="goToPost">
                    <span class="profile-stats-text">posts</span>
                    <span class="profile-stats-number">{{ postsCount }}</span>
                </div>
                <div class="profile-stats" @click="getFollowers">
                    <span class="profile-stats-text">followers</span>
                    <span class="profile-stats-number">{{ followersCount }}</span>
                </div>
                <div class="profile-stats" @click="getFollowings">
                    <span class="profile-stats-text">followings</span>
                    <span class="profile-stats-number">{{ followingsCount }}</span>
                </div>
            </div>
            <div class="top-body-profile-actions" v-if="!isOwner">
                <button class="profile-actions-button follow-button" @click="follow()"> {{ followTextButton }} </button>
            </div>
        </div>
    </div>

    <ProfilesList v-if="showList" :dataGetter="dataGetter" :textHeader="textHeader" :typeList="typeList"
        @exit-list="freeLists" />
        <div v-if="showPost" class="post-view" @click.self="exitPost">
            <Post :postData="postViewData" @update-like="updateLike" @delete-post="deletePost" @error-occurred="(value) => { errorMsg = value }" />
        </div>

    <div class="posts-container">
        <span v-if="(posts.length === 0)" class="posts-grid-nopost-text"> There are no posts yet </span>
        <div class="posts-grid-container" v-if="posts.length > 0">
            <div v-for="post in posts" :key="post.postID" class="posts-grid-post" @click="openPost(post)">
                <img :src="`data:image/jpg;base64,${post.image}`" loading="lazy" class="posts-grid-post-image"
                    :id="post.postID">
            </div>
        </div>

    </div>
</template>