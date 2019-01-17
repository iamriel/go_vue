<template>
  <form class="login-form" @submit.prevent>
    <div class="form-group">
      <fg-input
        id="email"
        v-model="formData.email"
        :error-message="emailError.message"
        :class="{'has-error': emailError.message}"
        @keyup="keyup"
        type="email"
        label="Email"
      ></fg-input>
    </div>

    <div class="form-group form-check show-password d-none d-sm-block">
      <i :class="showPasswordIconClass"></i>
      <label class="form-check-label" for="show-password" @click="toggleShowPassword">
        {{ showPasswordLabel }}
      </label>
    </div>

    <div class="form-group">
      <fg-input
        id="password"
        v-model="formData.password"
        :error-message="passwordError.message"
        :class="{'has-error': passwordError.message}"
        :type="showPassword ? 'text' : 'password'"
        @keyup="keyup"
        label="Password"
      ></fg-input>
    </div>

    <div class="form-group form-check show-password d-sm-none">
      <i :class="showPasswordIconClass"></i>
      <label class="form-check-label" for="show-password" @click="toggleShowPassword">
        {{ showPasswordLabel }}
      </label>
    </div>

    <div class="px-0 py-5">
      <p-button
        :disabled="loading"
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
      showPassword: false,
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
    },
    showPasswordIconClass() {
      return {
        fa: true,
        "fa-lg": true,
        "fa-eye": !this.showPassword,
        "fa-eye-slash": this.showPassword
      };
    },
    showPasswordLabel() {
      return this.showPassword ? "Hide password" : "Show password";
    }
  },
  methods: {
    clearErrors() {
      this.errors = [];
    },
    appendError(field, message) {
      this.errors.push({
        "field": field,
        "message": message
      });
    },
    isValid() {
      let isValid = true
      if (!this.formData.email) {
        isValid = false
        this.appendError("email", "Email is required")
      }
      if (!this.formData.password) {
        isValid = false
        this.appendError("password", "Password is required")
      }
      return isValid
    },
    login() {
      this.clearErrors();
      if (!this.isValid()) {
        return
      }
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
    },
    keyup(event) {
      if (event.keyCode === 13) {
        this.login();
      } else {
        this.clearErrors();
      }
    },
    toggleShowPassword() {
      this.showPassword = !this.showPassword;
    }
  }
};
</script>

<style lang="scss" scoped>
@import "./../../assets/sass/components/login.scss";
</style>
