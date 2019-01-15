<template>
  <form @submit.prevent>
    <div class="row px-md-5">
      <div class="col-12">
        <fg-input
          v-model="formData.email"
          :error-message="emailError.message"
          :class="{'has-error': emailError.message}"
          type="text"
          label="Email"
        ></fg-input>
      </div>
      <div class="col-12">
        <fg-input
          v-model="formData.password"
          :error-message="passwordError.message"
          :class="{'has-error': passwordError.message}"
          type="password"
          label="Password"
        ></fg-input>
      </div>
    </div>

    <div class="row px-md-5 mb-4">
      <div class="col">
        <div class="row py-md-5">
          <div class="action col py-1">
            <p-button
              @click.native.prevent="login"
              :disabled="disableSubmit"
              class="btn-login"
              type="warning"
              size="sm"
            >Login</p-button>

            <img src="@/assets/img/app/plants-right.svg" alt="" style="width: 120px;"/>
          </div>
        </div>
      </div>
    </div>
  </form>
</template>

<script>
export default {
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
.action {
  img {
    position: absolute;
    right: 20px;
    bottom: -20px;

    @media (max-width: 767px) {
      display: none;
    }
  }
}

.btn-login {
  padding: 4px 30px;
  border-radius: 5px;

  @media (max-width: 767px) {
    width: 100%;
  }
}
</style>
