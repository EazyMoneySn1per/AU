<template>
  <div class="card" @click="lineToTwitter(url)">
    <div v-if="time" class="card-time">{{time}}</div>
    <van-image :src="picUrl"></van-image>
    <div class="card-title">
        <p>{{title}}</p>
    </div>

  </div>
</template>

<script>
export default {
  name: 'TwitterGroup',
  props: {
    picUrl: String,
    title: String,
    time: String,
    url: String
  },
  methods: {
    lineToTwitter (e) {
      window.location.href = e
    }
  }
}
</script>

<style>
.card{
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  /*box-shadow: 10px 10px 10px #F4F3F477;*/
  width: 90%;
  margin: 10px auto;
}
.card-title{
  width: 100%;
  background-color: white;
  text-align: center;
}
.card-title p{
  padding: 20px 10px;
  margin: 0;
  font-size: 16px;
  font-weight: 420;
}
.card-time{
  font-size: 14px;
  color: #888888;
  margin: 1vh;
}

</style>
