import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import {defineComponent} from 'vue'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import {library} from "@fortawesome/fontawesome-svg-core";
import {faLock} from "@fortawesome/free-solid-svg-icons/faLock";
import {faAt} from "@fortawesome/free-solid-svg-icons/faAt";
import {faCheck} from "@fortawesome/free-solid-svg-icons/faCheck";


library.add(faLock, faAt, faCheck)

createApp(App)
    .component('font-awesome-icon', FontAwesomeIcon)
    .mount('#app')
