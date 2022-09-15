import classes from './login_styles/login.module.css'


function Login_page(){
    return (
        <div>
            <h1 className={classes.title}>Login</h1>
            <div className={classes.centerouter}>
                <div>
                    Username: <input type={"text"} disabled={true} defaultValue={'Disabled'}/>
                </div>
                <div>
                    Password: <input type={'password'} disabled={true} defaultValue={'disabled'}/>
                </div>

                <a href="/Login/Admin">
                    <button className={classes.submitbutton}>Submit</button>
                </a>


            </div>
        </div>

    )
}

export default Login_page;