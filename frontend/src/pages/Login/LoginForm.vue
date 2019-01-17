<template>
  <form class="login-form" @submit.prevent>
    <div class="form-group">
      <fg-input
        v-model="formData.email"
        :error-message="emailError.message"
        :class="{'has-error': emailError.message}"
        type="email"
        label="Email"
      ></fg-input>
    </div>

    <div class="form-group form-check show-password d-none d-sm-block">
      <i class="fa fa-eye fa-lg"></i>
      <label class="form-check-label" for="show-password">Show password</label>
    </div>

    <div class="form-group">
      <fg-input
        v-model="formData.password"
        :error-message="passwordError.message"
        :class="{'has-error': passwordError.message}"
        type="password"
        label="Password"
      ></fg-input>
    </div>

    <div class="form-group form-check show-password d-sm-none">
      <i class="fa fa-eye fa-lg"></i>
      <label class="form-check-label" for="show-password">Show password</label>
    </div>

    <div class="px-0 py-5">
      <p-button
        @click.native.prevent="login"
        class="btn-login"
        size="sm"
      >Login</p-button>

      <div class="plants-right">
        <svg-plants-right/>
      </div>
    </div>
  </form>
</template>

<script>
import SvgPlantsRight from "@/components/svgs/SvgPlantsRight";

export default {
  components: {
    SvgPlantsRight
  },
  data() {
    return {
      loading: false,
      formData: {
        email: "",
        password: ""
      },
      errors: []
    };
  },
  computed: {
    disableSubmit() {
      return !this.formData.email || !this.formData.password || this.loading;
    },
    emailError() {
      return this.errors.find(error => error.field === "email") || {};
    },
    passwordError() {
      return this.errors.find(error => error.field === "password") || {};
    }
  },
  methods: {
    login() {
      this.errors = [];
      this.loading = true;
      this.$store
        .dispatch("auth/login", this.formData)
        .then(response => {
          this.$router.push({ name: "dashboard" });
        })
        .catch(err => {
          this.errors = err.response.data;
        })
        .finally(() => {
          this.loading = false;
        });
    }
  }
};
</script>

<style lang="scss" scoped>
@import "./../../assets/sass/components/login.scss";
</style>
