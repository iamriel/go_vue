<template>
  <div class="contact-us full-screen">
    <div class="wrapper wrapper-full-page section content">
      <div class="container h-100">
        <div class="row h-100 justify-content-center align-items-center">
          <div class="logo-container col-md-5 text-center ">
            <div class="row">
              <div class="col title text-center">
                <h1
                  class="text-hide"
                >datacue</h1>
                <h2>Automatically personalize your<br/>website & email to each visitor.</h2>
              </div>
            </div>

          </div>

          <div class="col-md-7 px-md-5">
            <div class="row px-md-3 mb-4">
              <div class="col-12 register text-center mt-4">
                <p class="float-md-right">
                  Don't have an account?
                  <a href="" style="color: #F3BB45;">Sign up</a>
                </p>
              </div>
              <div class="col-12">
                <div class="row px-3 py-2">
                  <div class="header col text-center">
                    <img
                      src="@/assets/img/app/plants-left.svg"
                      style="width: 100px;"
                      alt=""
                    />
                    <h3>Sign in to DataCue</h3>
                  </div>
                </div>
              </div>
            </div>

            <login-form />

          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import LoginForm from "./Login/LoginForm";

export default {
  components: {
    LoginForm
  }
};
</script>

<style lang="scss" scoped>
.logo-container {
  background: url("../assets/img/app/background-with-squares.svg") no-repeat;
  background-size: 100%;
  font-size: 16px;
  height: 350px;

  .title {
    h1 {
      background: url("../assets/img/app/logo.png") no-repeat center center;
      width: 100%;
      height: 75px;
    }
  }

  h2 {
    font-size: 16px;
    margin-bottom: 0;
    margin-top: 0;
    padding-bottom: 60px;
  }

  .row {
    -ms-flex-align: flex-end !important;
    align-items: flex-end !important;
    -ms-flex-pack: flex-end !important;
    justify-content: flex-end !important;
    height: 100%;
  }

  @media (min-width: 768px) {
    background: url("../assets/img/app/squares-background.svg") no-repeat center
      center;
    height: 100%;

    .row {
      -ms-flex-align: center !important;
      align-items: center !important;
      -ms-flex-pack: center !important;
      justify-content: center !important;
      height: 100%;
    }

    h2 {
      padding-bottom: 80px;
    }
  }
}

.header {
  border-bottom: 1px solid #ddd;
  margin-top: 30px;

  h3 {
    font-size: 20px;
    margin-bottom: 7px;
  }

  img {
    position: absolute;
    left: 0;
    bottom: 5px;

    @media (max-width: 767px) {
      display: none;
    }
  }
}
</style>
