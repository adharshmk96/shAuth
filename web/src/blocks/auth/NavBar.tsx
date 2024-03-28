import { UI_ROUTES } from "@/constants/ui_routes";
import { Component, createSignal } from "solid-js";

export const NavBar: Component<{}> = () => {

    const [isActive, setIsActive] = createSignal<boolean>(false);

    return (
        <header class="navbar py-2">
            <div class="container">
                <div class="navbar-brand">
                    <a class="navbar-item">
                        <img src="/icons/android-chrome-192x192.png" alt="Logo" />
                    </a>
                    <span
                        class="navbar-burger"
                        classList={{ 'is-active': isActive() }}
                        data-target="navbarMenuHeroC"
                        onclick={() => setIsActive(!isActive())}
                    >
                        <span></span>
                        <span></span>
                        <span></span>
                        <span></span>
                    </span>
                </div>
                <div id="navbarMenuHeroC" class="navbar-menu" classList={{ 'is-active': isActive() }}>
                    <div class="navbar-end">
                        <a class="navbar-item" href={UI_ROUTES.HOME} > Home </a>
                        <span class="navbar-item">
                            <a class="button is-success is-inverted" href={UI_ROUTES.LOGIN}>
                                <span>Login</span>
                            </a>
                        </span>
                        <span class="navbar-item">
                            <a class="button is-success is-inverted" href={UI_ROUTES.REGISTER}>
                                <span>Register</span>
                            </a>
                        </span>
                    </div>
                </div>
            </div>
        </header>
    )
}