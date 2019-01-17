<template>
  <div class="form-group" :class="{'input-group': hasIcon}">
    <slot name="label">
      <label v-if="label" class="control-label">
        {{label}}
      </label>
    </slot>
    <slot name="addonLeft">
      <span v-if="addonLeftIcon" class="input-group-prepend">
        <i :class="addonLeftIcon" class="input-group-text"></i>
      </span>
    </slot>
    <input
      :value="value"
      @input="$emit('input',$event.target.value)"
      @keyup="$emit('keyup', $event)"
      v-bind="$attrs"
      class="form-control"
      aria-describedby="addon-right addon-left">
    <slot></slot>
    <slot name="addonRight">
      <span v-if="addonRightIcon" class="input-group-append">
        <i :class="addonRightIcon" class="input-group-text"></i>
      </span>
    </slot>
    <slot name="errorMessage">
      <small
        v-if="errorMessage"
        class="form-text error-message"
      >{{ errorMessage }}</small>
    </slot>
  </div>
</template>
<script>
export default {
  inheritAttrs: false,
  name: "fg-input",
  props: {
    label: String,
    value: [String, Number],
    errorMessage: String,
    addonRightIcon: String,
    addonLeftIcon: String
  },
  computed: {
    hasIcon() {
      const { addonRight, addonLeft } = this.$slots;
      return (
        addonRight !== undefined ||
        addonLeft !== undefined ||
        this.addonRightIcon !== undefined ||
        this.addonLeftIcon !== undefined
      );
    }
  }
};
</script>
<style lang="scss" scoped>
.has-error {
  .error-message {
    color: #eb5e28;
  }
}

.form-control {
  background: #F7F7F7 !important;
  border: 1px solid #E9E9E9 !important;

  &:focus {
    color: #495057;
    background-color: #fff;
    border-color: #F2BF00;
    outline: none;
    box-shadow: 0px 0px 2px #F2BF00;
  }
}
</style>
