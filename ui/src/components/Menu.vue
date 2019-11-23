<template>
  <aside class="main-menu" v-bind:class="{ hidden }">
    <div class="main-menu-logo">
      <span>GMW</span>
    </div>
    <div class="main-menu-icon" v-on:click="hidden=!hidden" v-bind:class="{ flying: hidden }">
      <span>&lt;</span>
    </div>
    <nav>
    </nav>
  </aside>
</template>

<script>
export default {
  name: "Menu",
  data: function() {
    return {
      hidden: false
    };
  },
  props: {},
  methods: {}
};
</script>

<style scoped>
.main-menu {
  grid-column: 1 / 2;
  background-color: #e3e3e3;
  height: 100vh;
  width: 18vw;
  min-width: 16rem;
  max-width: 22rem;
  position: absolute;
  display: grid;
  grid-template-columns: 4rem 1fr 4rem;
  grid-template-rows: 4rem 1fr;
  transition-duration: 1s;
}

.main-menu > nav {
  grid-column: 1 / 4;
  grid-row: 2 / 3;
}

.main-menu-icon {
  position: absolute;
  grid-column: 3 / 4;
  grid-row: 1 / 2;
  width: 3rem;
  height: 3rem;
  background-color: #1da1f2;
  border-radius: 100%;
  margin: 0.5rem;
  transition-duration: 0.5s;
  color: #ffffff;
  font-size: 1.5rem;
  cursor: pointer;
}

.main-menu-logo > span,
.main-menu-icon > span {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  transform-origin: 50% 50%;
  transition-duration: 0.5s;
}

.main-menu-icon[class~="flying"] {
  transform: translateX(calc(100% + 1rem));
}

.main-menu-icon[class~="flying"] > span {
  transform: translate(-50%, -50%) rotateZ(-180deg);
}

.main-menu-logo {
  position: absolute;
  grid-column: 1 / 2;
  grid-row: 1 / 2;
  width: 3rem;
  height: 3rem;
  background-color: #e3000f;
  border-radius: 100%;
  margin: 0.5rem;
  color: #ffffff;
  font-size: 1.1rem;
}

.hidden {
  transform: translateX(-100%);
}

@media (orientation: portrait) {
  .main-menu {
    display: grid;
    width: 100%;
    max-width: unset;
  }

  .main-menu-icon {
    transition-duration: 0.5s;
    transform: scale(0);
  }
}
</style>
