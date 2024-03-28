import { Route, Router } from "@solidjs/router";

import Login from "@/pages/auth/login.tsx";
import CreateAccount from "@/pages/auth/register.tsx";
import Home from "@/pages/home.tsx";
import { UI_ROUTES } from "../constants/ui_routes.ts";
import Profile from "../pages/profile";

function NotFOund() {
    return (
        <>
            <h1>404</h1>
            <p>Page not found</p>
        </>
    );

}

export default function Routes() {
    return (
        <>
            <Router>
                <Route path={UI_ROUTES.HOME} component={Home} />
                <Route path={UI_ROUTES.PROFILE} component={Profile} />
                <Route path={UI_ROUTES.LOGIN} component={Login} />
                <Route path={UI_ROUTES.REGISTER} component={CreateAccount} />

                <Route path="*" component={NotFOund} />
            </Router>
        </>
    );
}