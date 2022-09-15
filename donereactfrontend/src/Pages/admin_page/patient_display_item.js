import classes from './admin_styles/patient_item.module.css'

function Patient_Entry(props){
    let {first_name, middle_name, last_name, birth_date, phone_number, email_address, street_address_line1, street_address_line2,
        city, state, zip_code, photo_path, appointment_date, appointment_time} = props.post;



    const formattedBirthDate = new Date(birth_date).toLocaleDateString('en-US', {
        day: 'numeric',
        month: 'long',
        year:  'numeric'
    });
    const formattedApptDate = new Date(appointment_date).toLocaleDateString('en-US', {
        day: 'numeric',
        month: 'long',
        year:  'numeric',

    });




    return(
        <div>
            <table className={classes.table}>
                <tr>
                    <th>First Name:</th>
                    <th>Middle Name:</th>
                    <th>Last Name:</th>
                    <th>Date of Birth:</th>
                    <th>Phone #:</th>
                    <th>Email Address:</th>
                    <th>Street Address 1:</th>
                    <th>Street Address 2:</th>
                    <th>City:</th>
                    <th>State:</th>
                    <th>Zip:</th>
                    <th>Photo Path:</th>
                    <th>Appt Date: </th>
                    <th>Appt Time: </th>
                </tr>
                <tr>
                    <td>{first_name}</td>
                    <td>{middle_name}</td>
                    <td>{last_name}</td>
                    <td>{formattedBirthDate}</td>
                    <td>{phone_number}</td>
                    <td>{email_address}</td>
                    <td>{street_address_line1}</td>
                    <td>{street_address_line2}</td>
                    <td>{city}</td>
                    <td>{state}</td>
                    <td>{zip_code}</td>
                    <td>{photo_path}</td>
                    <td>{formattedApptDate}</td>
                    <td>{appointment_time}</td>


                </tr>
            </table>

        </div>

    )


}



export default Patient_Entry;