import Vue from 'vue'
import App from './App.vue'
import './registerServiceWorker'
import router from './router'
import store from './store'

require('@/assets/styles/main.scss')
require('@/assets/styles/icons.scss')
require('@/assets/styles/navigation.scss')
require('@/assets/styles/forms.scss')
require('@/assets/styles/buttons.scss')
require('@/assets/styles/modals.scss')
require('@/assets/styles/animations.scss')

// Auth0 configuration
import { domain, clientId } from "../auth_config.json";
import { Auth0Plugin } from "./auth";

Vue.use(Auth0Plugin, {
  domain,
  clientId,
  onRedirectCallback: appState => {
    router.push(
      appState && appState.targetUrl
        ? appState.targetUrl
        : window.location.pathname
    );
  }
});

Vue.config.productionTip = false

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
