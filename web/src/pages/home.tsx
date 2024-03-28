import { NavBar } from "@/blocks/auth/NavBar";


export default function Home() {

    document.title = "Home";

    return (
        <>
            <section class="hero is-fullheight">
                <div class="hero-head">
                    <NavBar />
                </div>

                <div class="hero-body">
                    <div class="container">
                        <div class="columns is-centered">
                            <div class="column is-two-fifths">
                                <h1>
                                    Hello, Welcome to Service Hub
                                </h1>
                            </div>
                        </div>
                    </div>
                </div>
            </section>
        </>
    );
}