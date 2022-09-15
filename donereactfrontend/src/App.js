import './App.css';
import Title from "./Pages/home_page/Title";
import Login_page from "./Pages/login_display/login_page";
import NewPatientPage from "./Pages/new_patient/new_patient_detail_page";
import Admin_page from "./Pages/admin_page/admin_page";

function App() {


    let component
    if (window.location.pathname === '/') {
        component = <Title/>
    } else if (window.location.pathname === '/Login') {
        component = <Login_page/>
    } else if (window.location.pathname === '/Register') {
        component = <NewPatientPage/>
    } else if (window.location.pathname === '/Login/Admin') {
        component = <Admin_page />
    }
    return (
        <div>
            {component}
        </div>

    )

}

export default App;
