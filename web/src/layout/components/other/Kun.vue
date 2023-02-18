<template>
  <div class="kun">
    <!-- ???小黑子是吧??? -->
    <div class="chick">
      <div class="feather"></div>
      <div class="blush"></div>
      <div class="wing"></div>
      <div class="beak"></div>
      <div class="eye"></div>

      <div class="left-leg"></div>
      <div class="right-leg"></div>
    </div>
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
  height: 42.25%;
  width: 66.8%;
  background-color: #ffec02;
  border-radius: 100px 100px 0 0;
  top: 0%;
  left: 34%;
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
  height: 63.25%;
  width: 70.5%;
  background-color: #ffce00;
  position: absolute;
  top: 86.75%;
  left: -41%;
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
  height: 48%;
  width: 41%;
  border-radius: 5px;
  top: 58.5%;
  left: 70%;
  transform: skewX(30deg) rotate(-25deg);
  z-index: -1;
}

.left-leg,
.right-leg {
  height: 73.5%;
  width: 13.25%;
  background-color: #f17c00;
  position: absolute;
  top: 180.5%;
  left: 28.75%;
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
  height: 30%;
  width: 21%;
  position: absolute;
  top: 45.75%;
  left: 64.25%;
  border-radius: 20px;
  transform-origin: bottom;
  animation: blink 2s infinite;
  @keyframes blink {
    30% {
      height: 6%;
      top: 55%;
    }
  }
}

.blush {
  background-color: rgba(241, 124, 0, 0.5);
  position: absolute;
  height: 28.75%;
  width: 18.75%;
  border-radius: 50%;
  top: 64%;
  left: 40%;
}
.feather {
  height: 68.75%;
  width: 13%;
  border-radius: 35px;
  background-color: #ffec02;
  position: absolute;
  left: 30.5%;
  top: -28%;
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
