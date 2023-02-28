<template>
  <div class="main">
    <n-gradient-text type="info">哎呀糟糕!</n-gradient-text>
    <n-gradient-text type="info">网页去外太空旅行了</n-gradient-text>
    <section class="error-container">
      <span>4</span>
      <span></span>
      <span>4</span>
    </section>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue";
</script>

<style scoped lang="scss">
.main {
  width: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  font-size: 32px;
}
.animation {
  height: 70vh;
}

.error-container {
  position: relative;
  font-family: Arial, sans-serif;
  font-size: 240px;
  font-weight: 800;
  span {
    position: relative;
    display: inline-block;
    vertical-align: middle;
    &:nth-child(1) {
      color: #d1f2a5;
      animation: colordancing 4s infinite;
      @keyframes colordancing {
        0% {
          color: #d1f2a5;
        }
        25% {
          color: #f56991;
        }
        50% {
          color: #ffc48c;
        }
        75% {
          color: #effab4;
        }
        100% {
          color: #d1f2a5;
        }
      }
    }
    &:nth-child(2) {
      width: 140px;
      height: 140px;
      &:before {
        -webkit-transform: rotate(45deg);
        -moz-transform: rotate(45deg);
        transform: rotate(45deg);
      }
      &:before,
      &:after {
        content: "";
        position: absolute;
        top: 0;
        left: 0;
        width: inherit;
        height: inherit;
        border-radius: 50%;
        animation: shadowdancing 4s infinite;
        @keyframes shadowdancing {
          0% {
            box-shadow: inset 30px 0 0 rgba(209, 242, 165, 0.4),
              inset 0 30px 0 rgba(239, 250, 180, 0.4),
              inset -30px 0 0 rgba(255, 196, 140, 0.4),
              inset 0 -30px 0 rgba(245, 105, 145, 0.4);
          }
          25% {
            box-shadow: inset 30px 0 0 rgba(245, 105, 145, 0.4),
              inset 0 30px 0 rgba(209, 242, 165, 0.4),
              inset -30px 0 0 rgba(239, 250, 180, 0.4),
              inset 0 -30px 0 rgba(255, 196, 140, 0.4);
          }
          50% {
            box-shadow: inset 30px 0 0 rgba(255, 196, 140, 0.4),
              inset 0 30px 0 rgba(245, 105, 145, 0.4),
              inset -30px 0 0 rgba(209, 242, 165, 0.4),
              inset 0 -30px 0 rgba(239, 250, 180, 0.4);
          }
          75% {
            box-shadow: inset 30px 0 0 rgba(239, 250, 180, 0.4),
              inset 0 30px 0 rgba(255, 196, 140, 0.4),
              inset -30px 0 0 rgba(245, 105, 145, 0.4),
              inset 0 -30px 0 rgba(209, 242, 165, 0.4);
          }
          100% {
            box-shadow: inset 30px 0 0 rgba(209, 242, 165, 0.4),
              inset 0 30px 0 rgba(239, 250, 180, 0.4),
              inset -30px 0 0 rgba(255, 196, 140, 0.4),
              inset 0 -30px 0 rgba(245, 105, 145, 0.4);
          }
        }
      }
    }
    &:nth-child(3) {
      color: #f56991;
      animation: colordancing2 4s infinite;
      @keyframes colordancing2 {
        0% {
          color: #ffc48c;
        }
        25% {
          color: #effab4;
        }
        50% {
          color: #d1f2a5;
        }
        75% {
          color: #f56991;
        }
        100% {
          color: #ffc48c;
        }
      }
    }
  }
}

@media screen and (max-width: 700px) {
  .error-container {
    font-size: 120px;
    top: 10px;
    transition: all 1s;
    span:nth-child(2) {
      width: 72px;
      height: 72px;
      &:before,
      &:after {
        content: "";
        position: absolute;
        top: 0;
        left: 0;
        width: inherit;
        height: inherit;
        border-radius: 50%;
        animation: shadowdancing2 4s infinite;
        @keyframes shadowdancing2 {
          0% {
            box-shadow: inset 20px 0 0 rgba(209, 242, 165, 0.4),
              inset 0 20px 0 rgba(239, 250, 180, 0.4),
              inset -20px 0 0 rgba(255, 196, 140, 0.4),
              inset 0 -20px 0 rgba(245, 105, 145, 0.4);
          }
          25% {
            box-shadow: inset 20px 0 0 rgba(245, 105, 145, 0.4),
              inset 0 20px 0 rgba(209, 242, 165, 0.4),
              inset -20px 0 0 rgba(239, 250, 180, 0.4),
              inset 0 -20px 0 rgba(255, 196, 140, 0.4);
          }
          50% {
            box-shadow: inset 20px 0 0 rgba(255, 196, 140, 0.4),
              inset 0 20px 0 rgba(245, 105, 145, 0.4),
              inset -20px 0 0 rgba(209, 242, 165, 0.4),
              inset 0 -20px 0 rgba(239, 250, 180, 0.4);
          }
          75% {
            box-shadow: inset 20px 0 0 rgba(239, 250, 180, 0.4),
              inset 0 20px 0 rgba(255, 196, 140, 0.4),
              inset -20px 0 0 rgba(245, 105, 145, 0.4),
              inset 0 -20px 0 rgba(209, 242, 165, 0.4);
          }
          100% {
            box-shadow: inset 20px 0 0 rgba(209, 242, 165, 0.4),
              inset 0 20px 0 rgba(239, 250, 180, 0.4),
              inset -20px 0 0 rgba(255, 196, 140, 0.4),
              inset 0 -20px 0 rgba(245, 105, 145, 0.4);
          }
        }
      }
    }
  }
}
</style>
