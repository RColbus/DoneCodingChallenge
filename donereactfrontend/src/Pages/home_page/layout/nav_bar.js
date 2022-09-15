import Logo from "./logo";
import classes from './layout_styles/nav_bar.module.css'
function MainNavigation(){
    return(
        <header className={classes.header}>
            <a href='/'>
                <Logo/>
            </a>
            <nav>
                <ul>
                    <li>
                        <a href="/Login">Login</a>
                    </li>
                    <li>
                        <a href="/Register">Register</a>
                    </li>
                </ul>
            </nav>
        </header>
    )

}

export default MainNavigation;