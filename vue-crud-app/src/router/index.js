import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../components/HomeView.vue'
import CreatePost from '../components/CreatePost.vue'
import EditPost from '../components/EditPost.vue'

const routes = [
  {
    path: '/',
    name: 'HomeView',
    component: HomeView
  },
  {
    path: '/create',
    name: 'CreatePost',
    component: CreatePost
  },
  {
    path: '/edit/:id',
    name: 'EditPost',
    component: EditPost
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
