<template>
  <div>
    <div class="panel panel-default">
      <div class="panel-heading text-center">New Post</div>
      <div class="panel-body">
        <textarea class="form-control" rows="3" placeholder="What's going on" v-model="text"></textarea>
       </div>
    </div>
    <div class="text-center" v-if="text && text.length <= 140">
			<button type="button" @click="save" class="btn btn-primary">Save</button>
		</div>
    <div class="text-center" v-if="text && text.length > 140">
			<h3>Please enter text less than 140 characters!!!</h3>
		</div>
  </div>
</template>

<script>
import auth from '../auth'
import service from '../services/post'
import router from '../main'

export default {
  data () {
    return {
      username: '',
      text: '',
      model: {
        "text": '',
        "username": '',
        "isSave": false
      },
      error: ''
    }
  },
  mounted () {
    this.getUserInfo()
  },
  methods: {
    async getUserInfo () {
      this.$http.get('/api/user/info', { headers: auth.getAuthHeader() }).then(
        response => {
          this.username = response.body.username
        },
        response => {
          if (response.status === 401) {
            auth.logout(this)
          }
          console.log(response)
        }
      )
    },
    async save () {
      this.model.text = this.text
      this.model.username = this.username
		  if (this.$route.params.id) {
				this.model.id  = this.$route.params.id;
      }
      let result = await service.save(this.model);
      this.text = ''
		}
  }
}
</script>