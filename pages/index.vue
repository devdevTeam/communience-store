<template>
  <v-app>
    <v-content class="bg" fluid >     
      <v-row
      align="center"
      justify="center"
      height="100%"
      class="center"
    >
      <v-col
        class="text-center"
        cols="12"
      >
      
        <h1 class="txt anim-typewriter text-h4 font-weight-thin mb-4">
          Welcome to
        </h1>
        <h1 class="txt anim-typewriter_2 text-h4 font-weight-bold mb-4">
          コミュニエンスストア
        </h1>

      </v-col>
    </v-row>
    </v-content>
  </v-app>
</template>

<style>

  .bg {
    position: absolute;
    top: 0;
    left: 0;
    background-size: cover;
    background-image: url("data:image/svg+xml,%3csvg xmlns='http://www.w3.org/2000/svg' version='1.1' xmlns:xlink='http://www.w3.org/1999/xlink' xmlns:svgjs='http://svgjs.com/svgjs' width='1600' height='600' preserveAspectRatio='none' viewBox='0 0 1600 600'%3e%3cg mask='url(%26quot%3b%23SvgjsMask1059%26quot%3b)' fill='none'%3e%3cpath d='M544.858%2c205.807C580.48%2c207.999%2c619.099%2c199.003%2c637.493%2c168.419C656.313%2c137.127%2c648.117%2c97.45%2c628.582%2c66.598C610.462%2c37.981%2c578.726%2c21.657%2c544.858%2c22.117C511.803%2c22.566%2c480.908%2c39.851%2c464.978%2c68.818C449.58%2c96.818%2c453.97%2c130.512%2c470.067%2c158.117C486.03%2c185.493%2c513.228%2c203.861%2c544.858%2c205.807' fill='rgba(253%2c 228%2c 0%2c 0.4)' class='triangle-float1'%3e%3c/path%3e%3cpath d='M639.539%2c86.814C661.714%2c86.207%2c684.13%2c80.24%2c696.333%2c61.715C709.792%2c41.284%2c713.602%2c14.413%2c701.076%2c-6.603C688.786%2c-27.222%2c663.54%2c-33.934%2c639.539%2c-33.574C616.277%2c-33.225%2c593.572%2c-24.455%2c581.12%2c-4.803C567.753%2c16.292%2c563.664%2c43.924%2c577.091%2c64.98C589.798%2c84.908%2c615.913%2c87.461%2c639.539%2c86.814' fill='rgba(253%2c 228%2c 0%2c 0.4)' class='triangle-float3'%3e%3c/path%3e%3cpath d='M1106.617%2c216.881C1158.073%2c219.823%2c1210.824%2c201.732%2c1238.401%2c158.19C1267.801%2c111.769%2c1270.132%2c51.116%2c1241.182%2c4.413C1213.538%2c-40.182%2c1159.074%2c-53.311%2c1106.617%2c-54.407C1051.585%2c-55.557%2c988.216%2c-49.792%2c961.942%2c-1.424C936.288%2c45.802%2c968.89%2c99.68%2c998.579%2c144.479C1024.247%2c183.21%2c1060.229%2c214.229%2c1106.617%2c216.881' fill='rgba(253%2c 228%2c 0%2c 0.4)' class='triangle-float1'%3e%3c/path%3e%3cpath d='M1122.3378529873469 100.36286410968843L1043.892000076697 157.35711239492872 1100.8862483619373 235.80296530557865 1179.3321012725871 178.80871702033835z' fill='rgba(253%2c 228%2c 0%2c 0.4)' class='triangle-float1'%3e%3c/path%3e%3cpath d='M941.7875349047282 289.96406013957125L909.0270229252601 239.51729554576036 849.7371320242778 313.88144581186793z' fill='rgba(253%2c 228%2c 0%2c 0.4)' class='triangle-float1'%3e%3c/path%3e%3cpath d='M1168.3843079581704 331.6701225205118L1182.6793593200994 434.09286059107893 1254.8771373618838 379.6879325392698z' fill='rgba(253%2c 228%2c 0%2c 0.4)' class='triangle-float3'%3e%3c/path%3e%3c/g%3e%3cdefs%3e%3cmask id='SvgjsMask1059'%3e%3crect width='1600' height='600' fill='white'%3e%3c/rect%3e%3c/mask%3e%3cstyle%3e %40keyframes float1 %7b 0%25%7btransform: translate(0%2c 0)%7d 50%25%7btransform: translate(-10px%2c 0)%7d 100%25%7btransform: translate(0%2c 0)%7d %7d .triangle-float1 %7b animation: float1 5s infinite%3b %7d %40keyframes float2 %7b 0%25%7btransform: translate(0%2c 0)%7d 50%25%7btransform: translate(-5px%2c -5px)%7d 100%25%7btransform: translate(0%2c 0)%7d %7d .triangle-float2 %7b animation: float2 4s infinite%3b %7d %40keyframes float3 %7b 0%25%7btransform: translate(0%2c 0)%7d 50%25%7btransform: translate(0%2c -10px)%7d 100%25%7btransform: translate(0%2c 0)%7d %7d .triangle-float3 %7b animation: float3 6s infinite%3b %7d %3c/style%3e%3c/defs%3e%3c/svg%3e");
    width: 100%;
    height: 100%;
  }
  .center{
    display: flex;
    justify-content: center;
    align-items:flex-end;
  }
  .text-center{
    margin-top:200px;
  }
  .txt{
    position: relative;
    top: 50%;  
    width: 24em;
    margin: 0 auto;
    border-right: 2px solid rgba(255,255,255,.75);
    font-size: 180%;
    text-align: center;
    white-space: nowrap;
    overflow: hidden;
    transform: translateY(-50%); 
  }
  .anim-typewriter{/*アニメーションセッティング*/
  animation: typewriter 4s steps(44) 1s 1 normal both,
 blinkTextCursor 500ms steps(44) infinite normal;
 width: 30%;
}
@keyframes typewriter{/*タイプライターライクなアニメーション*/
  from{width: 0;}
  to{width: 24em;}
}
@keyframes blinkTextCursor{/*点滅するカーソルのアニメーション*/
  from{border-right-color: rgba(255,255,255,.75);}
  to{border-right-color: transparent;}
}
  .anim-typewriter_2{/*アニメーションセッティング*/
  animation: typewriter_2 8s steps(44) 1s 1 normal both,
 blinkTextCursor_2 500ms steps(44) infinite normal;
}
@keyframes typewriter_2{/*タイプライターライクなアニメーション*/
  from{width: 0;}
  to{width: 24em;}
}
@keyframes blinkTextCursor_2{/*点滅するカーソルのアニメーション*/
  from{border-right-color: rgba(255,255,255,.75);}
  to{border-right-color: transparent;}
}
  




</style>