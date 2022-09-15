import {Fragment, useState} from "react";
import classes from './new_patient_style/new_patient.module.css'


function NewPatientPage(){
    const [enteredFirstName, setEnteredFirstName] = useState('');
    const [enteredMiddleName, setEnteredMiddleName] = useState('');
    const [enteredLastName, setEnteredLastName] = useState('');
    const [enteredDateOfBirth, setEnteredDateOfBirth] = useState('');
    const [enteredPhoneNumber, setEnteredPhoneNumber] = useState('');
    const [enteredEmailAddress, setEnteredEmailAddress] = useState('');
    const [enteredAddressLine1, setEnteredAddressLine1] = useState('');
    const [enteredAddressLine2, setEnteredAddressLine2] = useState('');
    const [enteredCity, setEnteredCity] = useState('');
    const [enteredZipCode, setEnteredZipCode] = useState('');
    const [enteredState, setEnteredState] = useState('');
    const [enteredPhotoPath, setEnteredPhotopath] = useState('');
    const [enteredApptTime, setEnteredApptTime] = useState('');
    const [enteredApptDate, setEnteredApptDate] = useState('');


    const postNewPatient = (event) => {
        event.preventDefault();

        fetch(`http://localhost:8000/api/v1/registrations`, {
            method: 'POST',
            body: JSON.stringify({
                first_name: enteredFirstName,
                middle_name: enteredMiddleName,
                last_name: enteredLastName,
                phone_number: enteredPhoneNumber,
                email_address: enteredEmailAddress,
                street_address_line1: enteredAddressLine1,
                street_address_line2: enteredAddressLine2,
                city: enteredCity,
                state: enteredState,
                zip_code: enteredZipCode,
                photo_path: enteredPhotoPath,
                birth_date: enteredDateOfBirth,
                appointment_date: enteredApptDate,
                appointment_time: enteredApptTime
            }),
        }).then(response => {
            console.log('response', response)
            if(response.status === 200){
                alert(`Thank you ${enteredFirstName} your appointment on ${enteredApptDate} at ${enteredApptTime} has been saved. Confirmation sent to ${enteredEmailAddress}`)
            }
        }).catch(e=>{
            console.log('e', e)
        })

    }
    return (
        <Fragment>
            <h1 className={classes.title}>Register</h1>
            <form onSubmit={postNewPatient}>
                <div className={classes.centerouter}>
                    <div>
                        <label form="FirstName">First Name:</label>
                        <input type={"text"} required className={classes.inputfield} value={enteredFirstName} onChange={(event) => setEnteredFirstName(event.target.value)} id='FirstName'/>
                        <label form="MiddleName">Middle Initial:</label>
                        <input type={"text"} className={classes.inputfield} value={enteredMiddleName} onChange={(event) => setEnteredMiddleName(event.target.value)} id='MiddleName'/>
                        <label form="LastName">Last Name:</label>
                        <input type={'text'} required className={classes.inputfield} value={enteredLastName} onChange={(event) => setEnteredLastName(event.target.value)} id='LastName'/>

                    </div>
                    <div>
                        <label form="DateOfBirth">Date of Birth:</label>
                        <input type={"date"} required className={classes.inputfield} value={enteredDateOfBirth} onChange={(event) => setEnteredDateOfBirth(event.target.value)} id='DateOfBirth'/>

                        <label form="PhoneNumber">Phone Number:</label>
                        <input type={'number'} required className={classes.inputfield} value={enteredPhoneNumber} onChange={(event) => setEnteredPhoneNumber(event.target.value)} id='PhoneNumber'/>
                        <label form="EmailAddress">Email Address:</label>
                         <input type={"email"} required className={classes.inputfield} value={enteredEmailAddress} onChange={(event) => setEnteredEmailAddress(event.target.value)} id='EmailAddress'/>
                    </div>
                    <div>
                        <label form="AddressLine1">Street Address 1:</label>
                         <input type={'text'} required className={classes.inputfield} value={enteredAddressLine1} onChange={(event) => setEnteredAddressLine1(event.target.value)} id='AddressLine1'/>
                        <label form="AddressLine2">Street Address 2:</label>
                        <input type={'text'}  className={classes.inputfield} value={enteredAddressLine2} onChange={(event) => setEnteredAddressLine2(event.target.value)} id='AddressLine2'/>
                        <label form="City">City:</label>
                        <input type={'text'} required className={classes.inputfield} value={enteredCity} onChange={(event) => setEnteredCity(event.target.value)} id='City'/>

                    </div>
                    <div>


                        <label form="State">State:</label>
                        <input type={'text'} required className={classes.inputfield} value={enteredState} onChange={(event) => setEnteredState(event.target.value)} id='State'/>
                        <label form="ZipCode">Zip Code:</label>
                        <input type={'number'} required className={classes.inputfield} value={enteredZipCode} onChange={(event) => setEnteredZipCode(event.target.value)} id='ZipCode'/>
                        <label form="PhotoPath">DL Photo:</label>
                        <input type={"file"} required className={classes.inputfield} value={enteredPhotoPath} onChange={(event) => setEnteredPhotopath(event.target.value)} id='PhotoPath'/>
                    </div>
                    <div>

                        <label form="ApptDate">Desired Appointment Date:</label>
                         <input type={'date'} required className={classes.inputfield} value={enteredApptDate} onChange={(event) => setEnteredApptDate(event.target.value)} id='ApptDate'/>
                        <label form="ApptTime">Desired Appointment Time:</label>
                         <input type={'time'} required className={classes.inputfield} value={enteredApptTime} onChange={(event) => setEnteredApptTime(event.target.value)} id='ApptTime'/>
                    </div>
                    <div className={classes.div}>
                        <button type={"submit"}>Submit</button>
                    </div>

                </div>

            </form>


        </Fragment>
    )
}


export default NewPatientPage;