import Patient_Entry from "./patient_display_item";
import classes from './admin_styles/patient_display_grid.module.css'

//this page accepts the API get request response object as props. Iterates through the response object and for
// every patient creates a child of patient entry and passes the props to patient entry for assignment and formatting


function Patient_display_grid({props}){

    return (
        <div>
            <h1 className={classes.title}>patient grid</h1>
            <ul className={classes.grid}>
                {props.map((post) => (
                    <Patient_Entry key={post.id} post={post}/>))}
            </ul>

        </div>)

}
export default Patient_display_grid