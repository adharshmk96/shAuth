/* @refresh reload */
import { render } from 'solid-js/web'
import { Provider } from './providers'
import Routes from './router/routes'

const root = document.getElementById('root')

render(
    () => <>
        <Provider>
            <Routes />
        </Provider>
    </>,root!)
