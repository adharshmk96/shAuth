import { NavBar } from "@/blocks/auth/NavBar";
import { UI_ROUTES } from "@/constants/ui_routes";

export default function Login() {

    document.title = "Login";

    return (
        <>
            <section class="hero is-fullheight">
                <div class="hero-head">
                    <NavBar/>
                </div>

                <div class="hero-body">
                    <div class="container">
                        <div class="columns is-centered">
                            <div class="column is-two-fifths">
                                <form id="login-form" class="box">
                                    <h1 class="title">Login</h1>
                                    <h2 class="subtitle">Enter your login details</h2>
                                    <div class="field">
                                        <label for="login-email" class="label">Email</label>
                                        <p class="control">
                                            <input id="login-email" class="input" name="email" type="email" placeholder="Email" />
                                        </p>
                                        <p class="help">Enter your email here</p>
                                    </div>
                                    <div class="field">
                                        <label for="login-password" class="label">Password</label>
                                        <p class="control">
                                            <input id="login-password" class="input" name="password" type="password" placeholder="Password" />
                                        </p>
                                        <p class="help">Enter your password here</p>
                                    </div>
                                    <div class="field">
                                        <p>
                                            Don't have an account? <a href={UI_ROUTES.REGISTER}>Register here</a>
                                        </p>
                                    </div>
                                    <div class="field">
                                        <p class="control">
                                            <button id="login-submit" class="button is-success" type="submit">
                                                Login
                                            </button>
                                        </p>
                                    </div>
                                </form>
                            </div>
                        </div>
                    </div>
                </div>
            </section>
        </>
    );
}