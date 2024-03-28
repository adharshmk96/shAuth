import { COOKIE_KEY } from '@/constants/config';
import { Accessor, Component, JSXElement, createContext, createEffect, createSignal, useContext } from 'solid-js';

type ThemeType = 'light' | 'dark' | 'auto';

interface ThemeContextType {
    theme: Accessor<ThemeType>;
    switchTheme: (theme: ThemeType) => void;
  }

const ThemeContext = createContext<ThemeContextType>();

export const useTheme = () => {
    const context = useContext(ThemeContext);
    if (!context) {
        throw new Error('useTheme must be used within a ThemeProvider');
    }
    return context;
};

export const ThemeProvider: Component<{children: JSXElement}> = (props) => {
    const [theme, setTheme] = createSignal<ThemeType>('auto');

    const switchTheme = (theme: ThemeType) => {
        let themeToApply = theme
        if (theme === 'auto') {
            const darkQuery = window.matchMedia('(prefers-color-scheme: dark)');
            themeToApply = darkQuery.matches ? 'dark' : 'light';
        }
        document.documentElement.setAttribute('data-theme', themeToApply);
        localStorage.setItem(COOKIE_KEY.THEME, theme);
        setTheme(theme);
    }; 

    createEffect(() => {
        const localTheme = localStorage.getItem(COOKIE_KEY.THEME) as ThemeType;
        if (localTheme) {
            switchTheme(localTheme);
        } else {
            switchTheme('auto');
        }
    });


    return (
        <ThemeContext.Provider value={{ theme, switchTheme }}>
            {props.children}
        </ThemeContext.Provider>
    );
};
