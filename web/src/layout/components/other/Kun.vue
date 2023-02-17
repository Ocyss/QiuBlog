<template>
  <div class="kun">
    <!-- ???小黑子是吧??? -->
    <div class="chick"></div>
    <div class="wing"></div>
    <div class="beak"></div>
    <div class="eye"></div>
    <div class="blush"></div>
    <div class="feather"></div>
    <div class="left-leg"></div>
    <div class="right-leg"></div>
  </div>
</template>

<style scoped lang="scss">
.kun {
  position: relative;
  display: flex;
  justify-content: center;
  align-content: center;
  z-index: 300;
}

.chick {
  position: absolute;
  height: 16.25%;
  width: 25%;
  background-color: #ffec02;
  border-radius: 100px 100px 0 0;
  top: 2.5%;
  left: -1%;
  animation: chick-walk 1.3s infinite;
  &:before {
    content: "";
    position: absolute;
    height: 123%;
    width: 150%;
    background-color: #ffec02;
    top: 100%;
    right: 0;
    border-radius: 15px 0 150px 150px;
  }
  @keyframes chick-walk {
    50% {
      top: 5%;
    }
  }
}

.wing {
  height: 9.25%;
  width: 18.5%;
  background-color: #ffce00;
  position: absolute;
  top: 18.75%;
  left: -10%;
  border-radius: 10px 10px 74px 74px;
  transform-origin: right;
  animation: wing 0.8s infinite;
  @keyframes wing {
    50% {
      transform: rotate(10deg);
    }
  }
}

.beak {
  position: absolute;
  background-color: #f17c00;
  height: 10%;
  width: 10%;
  border-radius: 5px;
  top: 12.5%;
  left: 17%;
  transform: skewX(30deg) rotate(-25deg);
  z-index: -1;
}

.left-leg,
.right-leg {
  height: 12.5%;
  width: 3.25%;
  background-color: #f17c00;
  position: absolute;
  top: 31.5%;
  left: 6.75%;
  transform-origin: top;
  transform: rotate(-45deg);
  animation: walk 1s infinite;
  @keyframes walk {
    33.33% {
      transform: translateY(5%) rotate(0deg);
    }
    66.66% {
      transform: translate(-6.25%, 2.5%) rotate(30deg);
    }
  }
  z-index: -1;
  &:before {
    content: "";
    position: absolute;
    height: 30%;
    width: 300%;
    background-color: #f17c00;
    right: -130%;
    top: 90%;
    border-radius: 40px;
  }
}

.right-leg {
  animation-delay: 0.5s;
}

.eye {
  background-color: #1c1c1c;
  height: 5%;
  width: 5%;
  position: absolute;
  top: 13.75%;
  left: 16.25%;
  border-radius: 20px;
  transform-origin: bottom;
  animation: blink 2s infinite;
  @keyframes blink {
    30% {
      height: 2px;
      top: 15%;
    }
  }
}

.blush {
  background-color: rgba(241, 124, 0, 0.5);
  position: absolute;
  height: 3.75%;
  width: 3.75%;
  border-radius: 50%;
  top: 17%;
  left: 11.5%;
}
.feather {
  height: 8.75%;
  width: 3%;
  border-radius: 35px;
  background-color: #ffec02;
  position: absolute;
  left: 7.5%;
  transform: rotate(-15deg);
  &:before {
    content: "";
    position: absolute;
    height: 130%;
    width: 95%;
    background-color: #ffec02;
    top: 5px;
    left: 10px;
    border-radius: 25px;
    transform: rotate(25deg);
  }
}
</style>
