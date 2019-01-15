<template>
  <form @submit.prevent>
    <div class="row px-md-5">
      <div class="col-12">
        <fg-input
          type="text"
          label="Email"
        ></fg-input>
      </div>
      <div class="col-12">
        <fg-input
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
export default {};
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
