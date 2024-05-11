<template>
  <div class="auth">
    <div class="auth__container">
      <div class="auth__up">
        <div class="up__logo">
          <img src="../../images/logo.svg" alt="logo" class="up__logo-img" />
        </div>
        <div class="up__title">
          <div class="up__title-place"></div>
          <div class="up__title-name">VULNERABILITY EXPLORER & MANAGEMENT</div>
        </div>
      </div>
      <div class="auth__form">
        <Form class="form" name="form" v-slot="{ meta }" @submit="onSubmit">
          <p class="error__message" :class="{ invalidCred: !isActive }">
            Неверный логин или пароль
          </p>
          <ErrorMessage name="login" class="error__message" />
          <Field
            name="login"
            type="text"
            class="form__input form__input-login"
            placeholder="Логин"
            :rules="validateLogin"
          />
          <ErrorMessage name="password" class="error__message" />
          <Field
            name="password"
            type="password"
            class="form__input form__input-password"
            placeholder="Пароль"
            :rules="validatePassword"
          />
          <button class="form__button" @click="Auth" :disabled="!meta.valid">
            Войти
          </button>
        </Form>
      </div>
    </div>
  </div>
</template>

<script>
import { defineComponent } from "vue";
import { Form, Field, ErrorMessage } from "vee-validate";
import * as httpClient from "../../httpClient";
import * as storage from "../../storage";

export default defineComponent({
  name: "auth",
  components: {
    Form,
    Field,
    ErrorMessage,
  },
  data: function () {
    return {
      user: {
        login: null,
        password: null,
      },
      isActive: false,
    };
  },
  methods: {
    onSubmit: function (values) {
      this.user.login = values.login;
      this.user.password = values.password;
    },
    validateLogin: function (value) {
      this.isActive = false;
      if (!value) {
        return "Поле должно быть заполнено";
      }
      const regex = /^[A-Za-z0-9]/i;
      if (!regex.test(value)) {
        return "Логин может содержать только латинские буквы и цифры";
      }
      this.user.login = value;
      return true;
    },
    validatePassword: function (value) {
      this.isActive = false;
      if (!value) {
        return "Поле должно быть заполнено";
      }
      this.user.password = value;
      return true;
    },
    Auth: async function () {
      let sendUrl = "/login";

      httpClient
        .Post(sendUrl, this.user)
        .then((response) => {
          storage.set("token", response.data.token);
          storage.set("accFIO", response.data.fio);
          this.$router.push("/dashboard/main");
        })
        .catch((err) => {
          this.isActive = !this.isActive;
        });
    },
  },
});
</script>

<style scoped lang="scss">
@use "auth";
</style>
