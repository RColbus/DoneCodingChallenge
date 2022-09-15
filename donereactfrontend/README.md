# Done front end in Reactjs  
  

This is the front end for the coding challenge done in Reactjs.  
Home page is a landing page with a home page landing element.  
Navigation bar at the top is how to navigate to the different pages.  
To navigate back home click on the HealthFirst Link in the upper right.  
  
On the login page you will notice the input fields are disabled.  
Login functionality is disabled for the purposes of this application  
as login authentication would add too much overhead in the context of this challenge  
  
The registration page is where a new patient would input their data.  
The validation here on the front end is simple in that required elements  
are flagged as required on the front end and marked as required on the backend.  
  
The admin page is where after authenticated login the administrator would view all patient records.  
these entries on this page are children of the Patient Entry function and upon API response  
the response object is iterated through and for every patient in the response an instance  
of Patient Entry function is created. All patients are then passed to patient grid  
for containerization and formatting and Patient Grid is then passed to admin page to display to the admin  

# Requirements  
nodejs *the latest version*  
verify npm *the latest version* should be included with nodejs  
