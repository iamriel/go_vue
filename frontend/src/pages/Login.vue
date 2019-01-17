<template>
  <div class="container-fluid px-0">
    <div class="row no-gutters">
      <div class="col-lg-4 col-md-4 col-sm-12">
        <div class="main-logo-container">
          <div class="squares-top">
            <svg-squares-top/>
          </div>

          <div class="logo-container">
            <svg-logo/>
            <h2 class="description">
              Automatically personalize your website & email to each visitor.
            </h2>
          </div>

          <div class="squares-bottom">
            <svg-squares-bottom/>
          </div>
        </div>
      </div>

      <div class="col-md-8 col-xs-12 right-container">
        <div class="container">
          <div class="col-md-10 col-lg-7 col-xs-12 login-form-container">
            <div class="login-form-header">
              <p class="account-link float-md-right">Don't have an account? <a href="#">Signup</a></p>
              <div class="plants-left">
                <svg-plants-left/>
              </div>
              <h1>Sign in to DataCue</h1>
            </div>

            <login-form/>

          </div>
        </div>
        <div class="footer">
          <p>2019 All Rights Reserved to DataCue</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import SvgLogo from "@/components/svgs/SvgLogo";
import SvgSquaresTop from "@/components/svgs/SvgSquaresTop";
import SvgSquaresBottom from "@/components/svgs/SvgSquaresBottom";
import SvgPlantsLeft from "@/components/svgs/SvgPlantsLeft";
import LoginForm from "./Login/LoginForm";

export default {
  components: {
    LoginForm,
    SvgLogo,
    SvgSquaresTop,
    SvgSquaresBottom,
    SvgPlantsLeft
  }
};
</script>

<style lang="scss" scoped>
@import "./../assets/sass/components/login.scss";
</style>
