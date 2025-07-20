// Vuetify
//@ts-ignore
import 'vuetify/styles'
import { createVuetify } from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'



const vuetify = createVuetify({
    theme : {
        defaultTheme: 'light',
    },
    icons : {
        defaultSet: 'mdi'
    },
    components,
    directives,
})

export default vuetify