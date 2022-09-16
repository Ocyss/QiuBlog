<template>
  <div
    class="main"
    v-for="(item, index) in 8"
    :class="index % 2 === 0 ? 'left' : 'right'"
  >
    <div class="bgimg">
      <img
        :src="`https://api.ixiaowai.cn/gqapi/gqapi.php?sadhujiasdbjkasdkl=${index}`"
        alt=""
      />
    </div>
    <div class="img">
      <img
        :src="`https://api.ixiaowai.cn/gqapi/gqapi.php?sadhujiasdbjkasdkl=${index}`"
        alt=""
      />
    </div>
    <div class="content">
      <div class="information">
        <div>
          <n-icon-wrapper :size="18" :border-radius="15">
            <n-icon :size="13" :component="Calendar" />
          </n-icon-wrapper>
          <n-time time-zone="Asia/Shanghai" :time="0" format="yyyy-MM-dd" />
        </div>
        <div class="rightInfo">
          <div style="margin-right: 0.5rem">
            <n-icon-wrapper
              :size="18"
              :border-radius="15"
              color="rgb(192,202,51)"
            >
              <n-icon :size="13" :component="Book" />
            </n-icon-wrapper>
            1.5K阅读
          </div>
          <div>
            <n-icon-wrapper
              color="rgb(67,160,71)"
              :size="18"
              :border-radius="15"
            >
              <n-icon :size="13" :component="PricetagsSharp" />
            </n-icon-wrapper>
            技术
          </div>
        </div>
      </div>
      <div class="title"><a>我是标题hhhhhhh</a></div>
      <div class="contentMain">
        <a>
          我是内容床前明月光疑是地上霜举头望明月低头思故乡鹅鹅鹅曲项向天歌白毛浮绿水红掌拨清波
        </a>
      </div>
      <div class="tags">
        <n-space :wrap="false">
          <n-tag v-for="index in 8" size="small" round>不该</n-tag>
        </n-space>
      </div>
    </div>
  </div>
</template>

<script setup>
import { Calendar, Book, PricetagsSharp } from "@vicons/ionicons5";
</script>

<style lang="scss" scoped>
.content {
  color: #fff;
  z-index: 10;
  padding: 24px 18px;
  width: 66.6666666%;
  text-shadow: 0 0.1875rem 0.3125rem #0000004d;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  letter-spacing: 0.06rem;
  .information {
    display: flex;
    align-items: center;
    justify-content: space-between;
    font-size: 0.3rem;
    > div {
      display: flex;
      align-items: center;
    }
  }
  .rightInfo > div {
    display: flex;

    align-items: center;
  }
  .title {
    font-size: 1.5rem;
    font-weight: 600;
  }
  .contentMain {
    font-size: 1rem;
    line-height: 1.8rem;
    -webkit-line-clamp: 2;
    overflow: hidden;
    display: -webkit-box;
    -webkit-box-orient: vertical;
    text-indent: 2rem;
  }
  .tags {
    overflow: hidden;
  }
}
.main {
  margin: 20px 0px;
  display: flex;
  width: 100%;
  height: 14.5rem;
  background-color: #fff;
  border-radius: 0.5rem;
  overflow: hidden;
  position: relative;

  ~ .right {
    flex-direction: row-reverse;
    .img {
      clip-path: polygon(6% 0, 100% 0, 100% 100%, 0 100%);
      -webkit-clip-path: polygon(6% 0, 100% 0, 100% 100%, 0 100%);
    }
  }
}
.img {
  width: 33.333333%;
  -webkit-clip-path: polygon(0 0, 100% 0, 94% 100%, 0 100%);
  clip-path: polygon(0 0, 100% 0, 94% 100%, 0 100%);
  img {
    width: 100%;
    height: 100%;
    object-fit: cover;
  }
  z-index: 2;
}
.bgimg {
  z-index: 0;
  box-sizing: border-box;
  border: 0 solid #e5e7eb;
  position: absolute;
  width: 100%;
  height: 14.5rem;
  overflow: hidden;
  img {
    object-fit: cover;
    background-position: 50%;
    height: 100%;
    width: 100%;
    filter: blur(1.875rem) brightness(0.75);
    transform: translate(0, 0) rotate(0) skewX(0) skewY(0) scaleX(1.25)
      scaleY(1.25);
  }
}
</style>
