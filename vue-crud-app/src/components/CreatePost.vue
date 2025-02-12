<template>
  <div>
    <h1>Create Post</h1>
    <form @submit.prevent="createPost">
      <input v-model="title" placeholder="Title" required>
      <textarea v-model="content" placeholder="Content" required></textarea>
      <button type="submit">Submit</button>
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
  methods: {
    createPost() {
      this.axios.post('http://localhost:8080/posts', {
        title: this.title,
        content: this.content
      })
      .then(() => {
        this.$router.push({ name: 'Home' }) // Navigate back to home after creation
      })
      .catch(error => {
        console.error('Error creating post:', error)
      })
    }
  }
}
</script>
