import classes from './layout/layout_styles/Title.module.css'
function Title() {
    return <section className={classes.hero}>
        <img src={'./images/home_page/happy_patient_concerned_dad.jpg'} width={300} height={300} alt={'happy pt'}/>
        <h1>Welcome to HealthFirst</h1>
        <p>Here at HealthFirst our patients needs come first. How may our family help yours?</p>
    </section>
}

export default Title;