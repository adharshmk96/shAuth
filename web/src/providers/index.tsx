import { Component, JSXElement } from "solid-js"
import { ThemeProvider } from "./ThemeProvider"

export const Provider: Component<{ children: JSXElement }> = ({ children }) => {
    return (
        <>
            <ThemeProvider>
                {children}
            </ThemeProvider>
        </>
    )
}
