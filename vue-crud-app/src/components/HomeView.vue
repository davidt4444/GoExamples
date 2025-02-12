<template>
  <div>
    <h1>Posts</h1>
    <router-link :to="{ name: 'CreatePost' }">Create New Post</router-link>
    <ul>
      <li v-for="post in posts" :key="post.id">
        <p>{{ post.title }}</p>
        <router-link :to="{ name: 'EditPost', params: { id: post.id }}">Edit</router-link>
        <button @click="deletePost(post.id)">Delete</button>
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  data() {
    return {
      posts: []
    }
  },
  mounted() {
    this.fetchPosts()
  },
  methods: {
    fetchPosts() {
      this.axios.get('http://localhost:8080/posts')
        .then(response => {
          this.posts = response.data
        })
        .catch(error => {
          console.error('Error fetching posts:', error)
        })
    },
    deletePost(id) {
      this.axios.delete(`http://localhost:8080/posts/${id}`)
        .then(() => {
          this.fetchPosts() // Refresh posts after deletion
        })
        .catch(error => {
          console.error('Error deleting post:', error)
        })
    }
  }
}
</script>
