<template>
  <div>
    <h1>Edit Post</h1>
    <form @submit.prevent="updatePost">
      <input v-model="title" placeholder="Title" required>
      <textarea v-model="content" placeholder="Content" required></textarea>
      <button type="submit">Update</button>
    </form>
  </div>
</template>

<script>
export default {
  data() {
    return {
      title: '',
      content: ''
    }
  },
  mounted() {
    this.fetchPost()
  },
  methods: {
    fetchPost() {
      this.axios.get(`http://localhost:8080/posts/${this.$route.params.id}`)
        .then(response => {
          this.title = response.data.title
          this.content = response.data.content
        })
        .catch(error => {
          console.error('Error fetching post:', error)
        })
    },
    updatePost() {
      this.axios.put(`http://localhost:8080/posts/${this.$route.params.id}`, {
        title: this.title,
        content: this.content
      })
      .then(() => {
        this.$router.push({ name: 'Home' }) // Navigate back to home after update
      })
      .catch(error => {
        console.error('Error updating post:', error)
      })
    }
  }
}
</script>