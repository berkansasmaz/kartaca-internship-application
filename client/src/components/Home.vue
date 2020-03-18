<template>
  <div>
    <h2 class="text-center" v-if="posts">Time line</h2>
    <h2 class="text-center" v-if="!posts">Please login and add a new post!!!</h2>
    <div v-if="posts">
      <div class="col-sm-6 col-sm-offset-3"  v-for="(post, index) in posts" :key="index">
        <div class="list-group">
          <a type="button" class="list-group-item"  @click="getId(post.ID)" data-toggle="modal" data-target="#detailModal">{{post.Text}}</a>
        </div>
      </div>
  
        <!-- Modal -->
        <div class="modal fade" id="detailModal" tabindex="-1" role="dialog" aria-labelledby="detailModalLabel" aria-hidden="true">
          <div class="modal-dialog" role="document">
            <div class="modal-content">
              <div class="modal-header">
                <h5 class="modal-title text-center" id="detailModalLabel"><strong>USER:&nbsp;</strong> {{posts[this.id-1].User}}</h5>
                <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                </button>
              </div>
              <div class="modal-body">
                <strong>Context:&nbsp;</strong> {{posts[this.id-1].Text}}
                <br>
                <span><strong>Time:&nbsp;</strong>{{ posts[this.id-1].UpdatedAt | moment("from", "now", true) }} Ago</span>
  
              </div>
              <div class="modal-footer">
                <button type="button" class="btn btn-secondary" data-dismiss="modal">Close</button>
              </div>
            </div>
          </div>
        </div>
    </div>
    </div>
</template>

<script>
export default {
  data () {
    return {
      posts: '',
      id: 1
    }
  },
 mounted(){
   this.$http.get('/api/post/getposts').then(response => {
        this.posts = response.body
      }, response => {
        console.log(response)
      })
 },
 methods:{
   getId(id){
     this.id = id
   }
 }
}
</script>
