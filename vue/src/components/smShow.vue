<template>
	<div class="container">
		<el-tree :data="data" :props="defaultProps" @node-click="handleNodeClick"></el-tree>
	</div>
</template>

<script>
export default {
	name:"smShow",
	data() {
      return {
        data: [],
        defaultProps: {
          children: 'Folder',
          label: 'Name'
        }
      };
    },
		
	mounted(){
		this.axios.patch('/api').then((res)=>{
			this.data=res.data.Folder
			console.log(res)
		})
	},
    methods: {
      handleNodeClick(data) {
        
      }
    }
}
</script>

<style lang="stylus" scoped>
	.container
		width:100%
</style>
