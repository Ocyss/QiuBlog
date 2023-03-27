<template>
  <div
    class="login_main"
    :style="`background-image: url(${config.global.randomImgApi})`"
  >
    <div class="login">
      <n-icon size="120" v-html="icon"></n-icon>
      <n-space vertical>
        <n-space vertical v-if="loginty">
          <n-input v-model:value="loginData.username" placeholder="用户名">
            <template #prefix>
              <n-icon :component="PersonCircle"></n-icon>
            </template>
          </n-input>
          <n-input
            v-model:value="loginData.password"
            type="password"
            placeholder="密码"
          >
            <template #prefix>
              <n-icon :component="LockClosed"></n-icon>
            </template>
          </n-input>
        </n-space>
        <n-space vertical v-else>
          <n-input v-model:value="registerData.username" placeholder="用户名">
            <template #prefix>
              <n-icon :component="PersonCircle"></n-icon>
            </template>
          </n-input>
          <n-input
            v-model:value="registerData.password"
            type="password"
            placeholder="密码"
          >
            <template #prefix>
              <n-icon :component="LockClosed"></n-icon>
            </template>
          </n-input>
          <n-input
            v-model:value="registerData.password2"
            type="password"
            placeholder="再次密码"
          >
            <template #prefix>
              <n-icon :component="LockClosed"></n-icon>
            </template>
          </n-input>
        </n-space>
        <n-space justify="space-around">
          <n-button type="info" @click="register">注册</n-button>
          <n-button type="info" @click="login">登陆</n-button>
        </n-space>
      </n-space>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, inject, Ref } from "vue";
import { PersonCircle, LockClosed } from "@vicons/ionicons5";
import { useMessage } from "naive-ui";
import { useRouter } from "vue-router";
import api from "@/api";
import type { Config } from "@/types";

const config: Ref<Config> = inject("config");
const router = useRouter();
const message = useMessage();
const loginty = ref(true);
const icon = ref(
  `<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 1024 1024" height="1024" width="1024">
<rect fill="url(#paint0_linear_469_1621)" rx="45.3955" height="144.013" width="90.791" y="143.361" x="465"></rect>
<path fill="url(#paint1_linear_469_1621)" d="M517.198 231.046C410.673 230.773 205.411 241.559 146.727 367.515C82.5471 505.266 95.07 821.47 209.341 862.169C269.225 883.498 351.41 899.089 490.067 902.27V902.869C499.33 902.869 508.371 902.82 517.198 902.725C526.025 902.82 535.066 902.869 544.329 902.869V902.27C682.986 899.089 765.171 883.498 825.054 862.169C939.326 821.47 951.849 505.266 887.669 367.515C828.985 241.559 623.723 230.773 517.198 231.046Z"></path>
<path fill="url(#paint2_linear_469_1621)" d="M146.727 367.515C205.411 241.559 410.673 230.773 517.198 231.046C623.723 230.773 828.985 241.559 887.669 367.515C951.849 505.266 939.326 821.47 825.054 862.169C765.171 883.498 682.986 899.089 544.329 902.27V902.869C535.066 902.869 526.025 902.82 517.198 902.725C508.371 902.82 499.33 902.869 490.067 902.869V902.27C351.41 899.089 269.225 883.498 209.341 862.169C95.07 821.47 82.5471 505.266 146.727 367.515ZM166.525 407.005C221.99 293.723 415.995 284.022 516.679 284.268C617.362 284.022 811.367 293.723 866.832 407.005C927.493 530.897 915.656 815.284 807.652 851.889C751.052 871.071 673.374 885.094 542.322 887.955V888.493C533.566 888.493 525.021 888.449 516.678 888.364C508.336 888.449 499.791 888.493 491.036 888.493V887.955C359.983 885.094 282.305 871.071 225.705 851.889C117.701 815.284 105.865 530.897 166.525 407.005Z" clip-rule="evenodd" fill-rule="evenodd"></path>
<path fill="url(#paint3_linear_469_1621)" d="M516.659 265.446C420.867 265.213 236.287 274.411 183.517 381.824C125.803 499.296 137.065 768.948 239.822 803.656C293.672 821.844 367.576 835.141 492.262 837.854V838.364C500.591 838.364 508.721 838.322 516.659 838.241C524.596 838.322 532.726 838.364 541.056 838.364V837.854C665.742 835.141 739.646 821.844 793.496 803.656C896.253 768.948 907.514 499.296 849.801 381.824C797.03 274.411 612.451 265.213 516.659 265.446Z"></path>
<path fill="#FFB629" d="M516.704 428.244C447.816 428.117 315.075 433.144 277.125 491.839C235.621 556.031 243.719 703.383 317.617 722.349C356.343 732.288 409.491 739.553 499.159 741.036V741.315C505.149 741.315 510.996 741.292 516.704 741.248C522.412 741.292 528.259 741.315 534.249 741.315V741.036C623.917 739.553 677.065 732.288 715.791 722.349C789.689 703.383 797.787 556.031 756.283 491.839C718.333 433.144 585.593 428.117 516.704 428.244Z"></path>
<path fill="#643800" d="M603.507 670.238C601.832 669.879 600.156 670.478 598.84 671.674C577.062 692.136 543.317 703.145 517.949 703.145H513.641H509.334C483.965 703.145 450.221 692.136 428.443 671.674C427.246 670.478 425.451 669.999 423.776 670.238C421.981 670.597 421.143 672.512 421.981 674.067C440.648 706.376 475.11 726.359 512.564 726.359H513.641H514.718C552.172 726.359 586.635 706.376 605.302 674.067C606.259 672.632 605.302 670.597 603.507 670.238Z"></path>
<path fill="url(#paint4_linear_469_1621)" d="M603.498 670.238C601.822 669.879 600.147 670.477 598.831 671.674C577.052 692.136 543.308 703.145 517.94 703.145H513.632H509.324C483.956 703.145 450.212 692.136 428.433 671.674C427.237 670.477 425.442 669.999 423.766 670.238C421.972 670.597 421.134 672.512 421.972 674.067C422.211 674.426 422.331 674.666 422.57 675.025C424.365 673.948 426.16 672.99 428.194 675.264C448.776 698.598 479.529 712.479 512.435 712.479H513.632H514.829C547.975 712.479 578.967 698.359 599.668 674.785C601.583 672.631 603.378 673.708 605.053 674.785C605.173 674.546 605.292 674.307 605.412 674.067C606.25 672.631 605.292 670.597 603.498 670.238Z"></path>
<path fill="#E19808" d="M280.246 494.997C318.196 436.302 450.937 431.275 519.825 431.403C588.714 431.275 721.454 436.302 759.404 494.997C777.285 522.652 785.959 565.74 784.132 607.348C782.637 573.03 774.152 539.721 759.404 516.912C721.454 458.217 588.714 453.191 519.825 453.318C450.937 453.191 318.196 458.217 280.246 516.912C265.499 539.721 257.014 573.03 255.518 607.348C253.692 565.74 262.366 522.652 280.246 494.997Z"></path>
<g filter="url(#filter0_ddd_469_1621)">
<path fill="url(#paint5_linear_469_1621)" d="M234.557 467.057C279.244 393.688 435.548 387.404 516.665 387.563C597.782 387.404 754.086 393.688 798.773 467.057C847.645 547.297 838.109 731.486 751.093 755.193C705.493 767.617 642.91 776.699 537.325 778.552V778.901C530.271 778.901 523.387 778.873 516.665 778.817C509.944 778.873 503.059 778.901 496.005 778.901V778.552C390.42 776.699 327.838 767.617 282.237 755.193C195.222 731.486 185.686 547.297 234.557 467.057ZM277.086 491.857C315.036 433.162 447.777 428.135 516.665 428.262C585.554 428.135 718.294 433.162 756.244 491.857C797.748 556.049 789.65 703.4 715.752 722.366C677.026 732.305 623.878 739.571 534.21 741.053V741.332C528.22 741.332 522.373 741.31 516.665 741.265C510.957 741.31 505.11 741.332 499.12 741.332V741.053C409.452 739.571 356.304 732.305 317.578 722.366C243.68 703.4 235.582 556.049 277.086 491.857Z" clip-rule="evenodd" fill-rule="evenodd"></path>
</g>
<rect fill="white" rx="45.9366" height="146.997" width="91.8733" y="480.24" x="348.119"></rect>
<rect fill="black" rx="43.8301" height="90.791" width="87.6603" y="535.279" x="350.879"></rect>
<rect stroke-width="6.26145" stroke="url(#paint6_linear_469_1621)" rx="45.9366" height="146.997" width="91.8733" y="480.24" x="348.119"></rect>
<rect fill="white" rx="45.9366" height="146.997" width="91.8733" y="480.24" x="587"></rect>
<rect fill="black" rx="43.8301" height="90.791" width="87.6603" y="535.279" x="589.76"></rect>
<rect stroke-width="6.26145" stroke="url(#paint7_linear_469_1621)" rx="45.9366" height="146.997" width="91.8733" y="480.24" x="587"></rect>
<ellipse fill-opacity="0.16" fill="url(#paint8_linear_469_1621)" ry="78.2681" rx="378.818" cy="311.749" cx="505.777"></ellipse>
<ellipse fill-opacity="0.4" fill="white" ry="21.9151" rx="165.928" cy="289.757" cx="524.567"></ellipse>
<path fill-opacity="0.13" fill="white" d="M205.239 593.911C205.218 582.378 205.795 570.841 206.953 559.529L255.268 553.724C253.466 564.942 252.447 576.579 252.239 588.264L205.239 593.911ZM809.25 487.168C813.476 497.294 817.021 508.368 819.86 520.069L771.821 525.84C768.124 513.783 763.371 502.757 757.608 493.372L809.25 487.168ZM686.953 402.529C733.884 413.483 775.645 432 796.671 463.134L736.193 470.4C683.74 431.377 576.346 427.629 516.745 427.739C503.113 427.714 486.981 427.891 469.578 428.645L686.953 402.529ZM242.205 455.962C239.416 459.321 236.886 462.843 234.638 466.534C223.633 484.602 215.589 507.941 210.726 533.531L261.178 527.47C265.206 513.885 270.558 501.555 277.166 491.334C294.047 465.226 329.683 449.736 370.111 440.595L242.205 455.962Z" clip-rule="evenodd" fill-rule="evenodd"></path>
<rect fill="url(#paint9_linear_469_1621)" transform="rotate(-13.8027 36.1191 200.322)" rx="6.26145" height="289.736" width="12.5229" y="200.322" x="36.1191"></rect>
<circle fill="url(#paint10_linear_469_1621)" r="21.9151" cy="190.437" cx="42.3936"></circle>
<g filter="url(#filter1_i_469_1621)">
<rect fill="url(#paint11_linear_469_1621)" rx="15.6536" height="106.445" width="56.353" y="465.922" x="26.7578"></rect>
</g>
<g filter="url(#filter2_i_469_1621)">
<rect fill="url(#paint12_linear_469_1621)" rx="15.6536" height="169.059" width="56.353" y="434.682" x="73.7188"></rect>
</g>
<rect fill-opacity="0.3" fill="url(#paint13_linear_469_1621)" height="28.1765" width="56.353" y="456.561" x="73.7188"></rect>
<rect fill-opacity="0.2" fill="url(#paint14_linear_469_1621)" height="28.1765" width="46.9608" y="481.561" x="26.7578"></rect>
<rect fill-opacity="0.15" fill="white" height="9.39217" width="56.353" y="497.281" x="73.7188"></rect>
<rect fill-opacity="0.15" fill="white" height="3.13072" width="56.353" y="512.922" x="73.7188"></rect>
<rect fill="url(#paint15_linear_469_1621)" transform="rotate(7.48049 954.318 196.162)" rx="6.26145" height="289.736" width="12.5229" y="196.162" x="954.318"></rect>
<circle fill="url(#paint16_linear_469_1621)" transform="rotate(21.2832 963.743 189.297)" r="21.9151" cy="189.297" cx="963.743"></circle>
<g filter="url(#filter3_i_469_1621)">
<rect fill="url(#paint17_linear_469_1621)" transform="matrix(-1 0 0 1 1003.56 465.723)" rx="15.6536" height="106.445" width="56.353"></rect>
</g>
<g filter="url(#filter4_i_469_1621)">
<rect fill="url(#paint18_linear_469_1621)" transform="matrix(-1 0 0 1 959.678 434.441)" rx="15.6536" height="169.059" width="56.353"></rect>
</g>
<rect fill-opacity="0.2" fill="url(#paint19_linear_469_1621)" height="28.1765" width="43.8301" y="481.281" x="959.678"></rect>
<rect fill-opacity="0.3" fill="url(#paint20_linear_469_1621)" height="28.1765" width="56.353" y="456.361" x="903.318"></rect>
<rect fill-opacity="0.15" fill="white" height="9.39217" width="56.353" y="497.041" x="903.318"></rect>
<rect fill-opacity="0.15" fill="white" height="3.13072" width="56.353" y="512.762" x="903.318"></rect>
<defs>
<filter color-interpolation-filters="sRGB" filterUnits="userSpaceOnUse" height="419.516" width="648.059" y="384.43" x="198.897" id="filter0_ddd_469_1621">
<feFlood result="BackgroundImageFix" flood-opacity="0"></feFlood>
<feColorMatrix result="hardAlpha" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0" type="matrix" in="SourceAlpha"></feColorMatrix>
<feOffset dy="3.13072" dx="3.13072"></feOffset>
<feComposite operator="out" in2="hardAlpha"></feComposite>
<feColorMatrix values="0 0 0 0 0.236163 0 0 0 0 0.539093 0 0 0 0 0.929167 0 0 0 1 0" type="matrix"></feColorMatrix>
<feBlend result="effect1_dropShadow_469_1621" in2="BackgroundImageFix" mode="normal"></feBlend>
<feColorMatrix result="hardAlpha" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0" type="matrix" in="SourceAlpha"></feColorMatrix>
<feOffset dy="-3.13072" dx="-3.13072"></feOffset>
<feComposite operator="out" in2="hardAlpha"></feComposite>
<feColorMatrix values="0 0 0 0 0.745833 0 0 0 0 0.862451 0 0 0 0 1 0 0 0 1 0" type="matrix"></feColorMatrix>
<feBlend result="effect2_dropShadow_469_1621" in2="effect1_dropShadow_469_1621" mode="normal"></feBlend>
<feColorMatrix result="hardAlpha" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0" type="matrix" in="SourceAlpha"></feColorMatrix>
<feOffset dy="12.5229" dx="6.26145"></feOffset>
<feGaussianBlur stdDeviation="6.26145"></feGaussianBlur>
<feComposite operator="out" in2="hardAlpha"></feComposite>
<feColorMatrix values="0 0 0 0 0 0 0 0 0 0.12064 0 0 0 0 0.25 0 0 0 0.12 0" type="matrix"></feColorMatrix>
<feBlend result="effect3_dropShadow_469_1621" in2="effect2_dropShadow_469_1621" mode="normal"></feBlend>
<feBlend result="shape" in2="effect3_dropShadow_469_1621" in="SourceGraphic" mode="normal"></feBlend>
</filter>
<filter color-interpolation-filters="sRGB" filterUnits="userSpaceOnUse" height="109.576" width="59.4842" y="465.922" x="26.7578" id="filter1_i_469_1621">
<feFlood result="BackgroundImageFix" flood-opacity="0"></feFlood>
<feBlend result="shape" in2="BackgroundImageFix" in="SourceGraphic" mode="normal"></feBlend>
<feColorMatrix result="hardAlpha" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0" type="matrix" in="SourceAlpha"></feColorMatrix>
<feOffset dy="3.13072" dx="3.13072"></feOffset>
<feGaussianBlur stdDeviation="3.13072"></feGaussianBlur>
<feComposite k3="1" k2="-1" operator="arithmetic" in2="hardAlpha"></feComposite>
<feColorMatrix values="0 0 0 0 1 0 0 0 0 1 0 0 0 0 1 0 0 0 0.25 0" type="matrix"></feColorMatrix>
<feBlend result="effect1_innerShadow_469_1621" in2="shape" mode="normal"></feBlend>
</filter>
<filter color-interpolation-filters="sRGB" filterUnits="userSpaceOnUse" height="172.189" width="59.4842" y="434.682" x="73.7188" id="filter2_i_469_1621">
<feFlood result="BackgroundImageFix" flood-opacity="0"></feFlood>
<feBlend result="shape" in2="BackgroundImageFix" in="SourceGraphic" mode="normal"></feBlend>
<feColorMatrix result="hardAlpha" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0" type="matrix" in="SourceAlpha"></feColorMatrix>
<feOffset dy="3.13072" dx="3.13072"></feOffset>
<feGaussianBlur stdDeviation="3.13072"></feGaussianBlur>
<feComposite k3="1" k2="-1" operator="arithmetic" in2="hardAlpha"></feComposite>
<feColorMatrix values="0 0 0 0 1 0 0 0 0 1 0 0 0 0 1 0 0 0 0.25 0" type="matrix"></feColorMatrix>
<feBlend result="effect1_innerShadow_469_1621" in2="shape" mode="normal"></feBlend>
</filter>
<filter color-interpolation-filters="sRGB" filterUnits="userSpaceOnUse" height="109.576" width="59.4842" y="465.723" x="947.205" id="filter3_i_469_1621">
<feFlood result="BackgroundImageFix" flood-opacity="0"></feFlood>
<feBlend result="shape" in2="BackgroundImageFix" in="SourceGraphic" mode="normal"></feBlend>
<feColorMatrix result="hardAlpha" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0" type="matrix" in="SourceAlpha"></feColorMatrix>
<feOffset dy="3.13072" dx="3.13072"></feOffset>
<feGaussianBlur stdDeviation="3.13072"></feGaussianBlur>
<feComposite k3="1" k2="-1" operator="arithmetic" in2="hardAlpha"></feComposite>
<feColorMatrix values="0 0 0 0 1 0 0 0 0 1 0 0 0 0 1 0 0 0 0.25 0" type="matrix"></feColorMatrix>
<feBlend result="effect1_innerShadow_469_1621" in2="shape" mode="normal"></feBlend>
</filter>
<filter color-interpolation-filters="sRGB" filterUnits="userSpaceOnUse" height="172.189" width="59.4842" y="434.441" x="903.324" id="filter4_i_469_1621">
<feFlood result="BackgroundImageFix" flood-opacity="0"></feFlood>
<feBlend result="shape" in2="BackgroundImageFix" in="SourceGraphic" mode="normal"></feBlend>
<feColorMatrix result="hardAlpha" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0" type="matrix" in="SourceAlpha"></feColorMatrix>
<feOffset dy="3.13072" dx="3.13072"></feOffset>
<feGaussianBlur stdDeviation="3.13072"></feGaussianBlur>
<feComposite k3="1" k2="-1" operator="arithmetic" in2="hardAlpha"></feComposite>
<feColorMatrix values="0 0 0 0 1 0 0 0 0 1 0 0 0 0 1 0 0 0 0.25 0" type="matrix"></feColorMatrix>
<feBlend result="effect1_innerShadow_469_1621" in2="shape" mode="normal"></feBlend>
</filter>
<linearGradient gradientUnits="userSpaceOnUse" y2="251.371" x2="565.183" y1="188.757" x1="454.042" id="paint0_linear_469_1621">
<stop stop-color="#FFD177"></stop>
<stop stop-color="#FFCA62" offset="0.267053"></stop>
<stop stop-color="#FFB629" offset="0.558713"></stop>
<stop stop-color="#FFCF72" offset="0.926149"></stop>
</linearGradient>
<linearGradient gradientUnits="userSpaceOnUse" y2="902.869" x2="517.198" y1="231.041" x1="517.198" id="paint1_linear_469_1621">
<stop stop-color="#FFCC68"></stop>
<stop stop-color="#F9A01A" offset="1"></stop>
</linearGradient>
<linearGradient gradientUnits="userSpaceOnUse" y2="902.869" x2="517.198" y1="231.041" x1="517.198" id="paint2_linear_469_1621">
<stop stop-color="#FFCC68"></stop>
<stop stop-color="#F9921A" offset="1"></stop>
</linearGradient>
<linearGradient gradientUnits="userSpaceOnUse" y2="838.364" x2="516.659" y1="265.441" x1="516.659" id="paint3_linear_469_1621">
<stop stop-opacity="0" stop-color="#FFB629"></stop>
<stop stop-color="#FFB629" offset="0.53125"></stop>
</linearGradient>
<linearGradient gradientUnits="userSpaceOnUse" y2="715.007" x2="513.634" y1="684.638" x1="513.634" id="paint4_linear_469_1621">
<stop stop-color="#3C2200" offset="0.00132565"></stop>
<stop stop-color="#512D00" offset="1"></stop>
</linearGradient>
<linearGradient gradientUnits="userSpaceOnUse" y2="778.901" x2="516.665" y1="387.561" x1="516.665" id="paint5_linear_469_1621">
<stop stop-color="#4E93F3"></stop>
<stop stop-color="#41BBFF" offset="1"></stop>
</linearGradient>
<linearGradient gradientUnits="userSpaceOnUse" y2="634.865" x2="400.626" y1="486.687" x1="354.682" id="paint6_linear_469_1621">
<stop stop-color="#CE8800"></stop>
<stop stop-color="#FFD789" offset="1"></stop>
</linearGradient>
<linearGradient gradientUnits="userSpaceOnUse" y2="634.865" x2="639.506" y1="486.687" x1="593.562" id="paint7_linear_469_1621">
<stop stop-color="#CE8800"></stop>
<stop stop-color="#FFD789" offset="1"></stop>
</linearGradient>
<linearGradient gradientUnits="userSpaceOnUse" y2="390.017" x2="505.776" y1="288.268" x1="505.776" id="paint8_linear_469_1621">
<stop stop-opacity="0" stop-color="white"></stop>
<stop stop-color="white" offset="1"></stop>
</linearGradient>
<linearGradient gradientUnits="userSpaceOnUse" y2="321.407" x2="55.3156" y1="321.01" x1="21.0528" id="paint9_linear_469_1621">
<stop stop-color="#F4F4F4"></stop>
<stop stop-color="#C8C8C8" offset="0.651042"></stop>
<stop stop-color="#DDDDDD" offset="1"></stop>
</linearGradient>
<linearGradient gradientUnits="userSpaceOnUse" y2="212.785" x2="65.2365" y1="182.417" x1="18.1283" id="paint10_linear_469_1621">
<stop stop-color="#FFDC97"></stop>
<stop stop-color="#FFB932" offset="0.46875"></stop>
<stop stop-color="#FFA800" offset="0.729167"></stop>
<stop stop-color="#FFE8BB" offset="0.994792"></stop>
</linearGradient>
<linearGradient gradientUnits="userSpaceOnUse" y2="572.366" x2="54.9343" y1="465.922" x1="54.9343" id="paint11_linear_469_1621">
<stop stop-color="#FFDF6F"></stop>
<stop stop-color="#FFBC3A" offset="0.833333"></stop>
<stop stop-color="#FFD178" offset="1"></stop>
</linearGradient>
<linearGradient gradientUnits="userSpaceOnUse" y2="603.741" x2="101.895" y1="434.682" x1="101.895" id="paint12_linear_469_1621">
<stop stop-color="#FFDF6F"></stop>
<stop stop-color="#FFBC3A" offset="0.833333"></stop>
<stop stop-color="#FFD178" offset="1"></stop>
</linearGradient>
<linearGradient gradientUnits="userSpaceOnUse" y2="484.737" x2="101.895" y1="456.561" x1="101.895" id="paint13_linear_469_1621">
<stop stop-color="white"></stop>
<stop stop-opacity="0" stop-color="white" offset="1"></stop>
</linearGradient>
<linearGradient gradientUnits="userSpaceOnUse" y2="509.737" x2="50.2382" y1="481.561" x1="50.2382" id="paint14_linear_469_1621">
<stop stop-color="white"></stop>
<stop stop-opacity="0" stop-color="white" offset="1"></stop>
</linearGradient>
<linearGradient gradientUnits="userSpaceOnUse" y2="317.17" x2="972.483" y1="317.327" x1="947.243" id="paint15_linear_469_1621">
<stop stop-color="#F4F4F4"></stop>
<stop stop-color="#C8C8C8" offset="0.651042"></stop>
<stop stop-color="#DDDDDD" offset="1"></stop>
</linearGradient>
<linearGradient gradientUnits="userSpaceOnUse" y2="211.645" x2="986.586" y1="181.278" x1="939.478" id="paint16_linear_469_1621">
<stop stop-color="#FFDC97"></stop>
<stop stop-color="#FFB932" offset="0.46875"></stop>
<stop stop-color="#FFA800" offset="0.729167"></stop>
<stop stop-color="#FFE8BB" offset="0.994792"></stop>
</linearGradient>
<linearGradient gradientUnits="userSpaceOnUse" y2="106.445" x2="28.1765" y1="0" x1="28.1765" id="paint17_linear_469_1621">
<stop stop-color="#FFDF6F"></stop>
<stop stop-color="#FFBC3A" offset="0.833333"></stop>
<stop stop-color="#FFD178" offset="1"></stop>
</linearGradient>
<linearGradient gradientUnits="userSpaceOnUse" y2="169.059" x2="28.1765" y1="0" x1="28.1765" id="paint18_linear_469_1621">
<stop stop-color="#FFDF6F"></stop>
<stop stop-color="#FFBC3A" offset="0.833333"></stop>
<stop stop-color="#FFD178" offset="1"></stop>
</linearGradient>
<linearGradient gradientUnits="userSpaceOnUse" y2="509.458" x2="981.593" y1="481.281" x1="981.593" id="paint19_linear_469_1621">
<stop stop-color="white"></stop>
<stop stop-opacity="0" stop-color="white" offset="1"></stop>
</linearGradient>
<linearGradient gradientUnits="userSpaceOnUse" y2="484.538" x2="931.495" y1="456.361" x1="931.495" id="paint20_linear_469_1621">
<stop stop-color="white"></stop>
<stop stop-opacity="0" stop-color="white" offset="1"></stop>
</linearGradient>
</defs>
</svg>
`
);
const registerData = ref({ username: "", password: "", password2: "" });
const loginData = ref({ username: "", password: "" });

function register() {
  if (loginty.value) {
    loginty.value = false;
  } else {
    if (
      registerData.value.username == "" ||
      registerData.value.password == "" ||
      registerData.value.password2 == ""
    ) {
      message.error("必写");
      return;
    } else if (registerData.value.password != registerData.value.password2) {
      message.error("两次密码不一致");
      return;
    } else {
      api.user.register(registerData.value).then((res) => {
        message.success("注册成功！！");
      });
    }
  }
}
function login() {
  if (loginty.value) {
    if (loginData.value.username == "" || loginData.value.password == "") {
      message.error("必写");
      return;
    } else {
      api.user.login(loginData.value).then((res) => {
        message.success("登陆成功！！");
        router.push({ name: "admin" });
      });
    }
  } else {
    loginty.value = true;
  }
}
</script>

<style scoped lang="scss">
.login_main {
  width: 100vw;
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  background-size: cover;
  background-attachment: fixed;
  background-position: center center;
  background-repeat: no-repeat;
}

.login {
  min-width: 350px;
  width: 30vw;
  height: 40vh;
  background-color: rgba(10, 10, 10, 0.6);
  display: flex;
  align-items: center;
  flex-direction: column;
  border-radius: 30px;
}
</style>
