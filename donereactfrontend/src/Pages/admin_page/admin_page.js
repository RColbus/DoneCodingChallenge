import classes from './admin_styles/admin_page.module.css'
import Patient_display_grid from "./patient_display_grid";
import {useEffect, useState} from "react";

function Admin_page() {
    const [AllPatients, setAllPatients] = useState([])
    function getPatients(){
        fetch('http://localhost:8000/api/v1/registrations?page_id=1&page_size=1000')
            .then((response) => response.json())
            .then((data) => setAllPatients(data));

    }

    useEffect(() => {

        getPatients()
    }, []);


    return (
        <div>
            <h1 className={classes.admintitle}>Admin display page</h1>
            <Patient_display_grid props={AllPatients}/>


        </div>

    )
}
export default Admin_page;

