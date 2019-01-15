<template>
  <li class="nav-item">
    <a class="nav-link" @click="logout">
      <i class="fa fa-sign-out"></i>
      <p>Logout</p>
    </a>
  </li>
</template>
<script>
export default {
  methods: {
    logout() {
      this.$store.dispatch("auth/logout");
      this.$router.push({ name: "login" });
    }
  }
};
</script>

<style lang="scss" scoped>
a {
  cursor: pointer;
}
</style>
